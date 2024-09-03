// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// package omclient is a thin client to do grpc-gateway RESTful HTTP gRPC calls
// with affordances for retry with exponential backoff + jitter.
package omclient

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	pb "github.com/googleforgames/open-match2/v2/pkg/pb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
	"google.golang.org/protobuf/encoding/protojson"
)

type RestfulOMGrpcClient struct {
	Client      *http.Client
	Log         *logrus.Logger
	Cfg         *viper.Viper
	tokenSource oauth2.TokenSource
}

// CreateTicket Example
// Your platform services layer should take the matchmaking request (in this
// example, we assume it is from a game client, but it could come via another
// of your game platform services as well), add attributes your platform
// services have authority over (ex: MMR, ELO, inventory, etc), then call Open
// Match core on the client's behalf to make a matchmaking ticket. In this sense,
// your platform services layer acts as a 'proxy' for the player's game client
// from the viewpoint of Open Match.
func (rc *RestfulOMGrpcClient) CreateTicket(ctx context.Context, ticket *pb.Ticket) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := rc.Log.WithFields(logrus.Fields{
		"component": "open_match_client",
		"operation": "proxy_CreateTicket",
	})

	// Update metrics
	//createdTicketCounter.Add(ctx, 1)

	// Put the ticket into open match CreateTicket request protobuf message.
	reqPb := &pb.CreateTicketRequest{Ticket: ticket}
	resPb := &pb.CreateTicketResponse{}
	var err error
	var buf []byte

	// Marshal request into json
	buf, err = protojson.Marshal(reqPb)
	if err != nil {
		// Mark as a permanent error so the backoff library doesn't retry this (invalid input)
		err := backoff.Permanent(err)
		logger.WithFields(logrus.Fields{
			"pb_message": "CreateTicketsRequest",
		}).Errorf("cannot marshal proto to json")
		return "", err
	}

	// HTTP version of the gRPC CreateTicket() call
	resp, err := rc.Post(ctx, logger, rc.Cfg.GetString("OM_CORE_ADDR"), "/tickets", buf)

	// Include headers in structured loggerging when debugging.
	headerFields := logrus.Fields{}
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		for key, value := range resp.Header {
			headerFields[key] = value
		}
	}
	logger = logger.WithFields(headerFields)

	if resp != nil && resp.StatusCode != http.StatusOK { // HTTP error code
		err := fmt.Errorf("%s (%d)", http.StatusText(resp.StatusCode), resp.StatusCode)
		logger.WithFields(logrus.Fields{
			"CreateTicketsRequest": reqPb,
		}).Errorf("CreateTicket failed: %v", err)
		//TODO: switch resp.Header?
		return "", err
	}

	if err != nil { // HTTP library error
		logger.WithFields(logrus.Fields{
			"CreateTicketsRequest": reqPb,
		}).Errorf("CreateTicket failed: %v", err)
		return "", err
	}

	// Read HTTP response body
	body, err := readAllBody(*resp, logger)
	if err != nil {
		// Mark as a permanent error so the backoff library doesn't retry this REST call
		err := backoff.Permanent(err)
		logger.WithFields(logrus.Fields{
			"pb_message": "CreateTicketsResponse",
			"error":      err,
		}).Errorf("cannot read http response body")
		return "", err
	}

	// Success, unmarshal json back into protobuf
	err = protojson.Unmarshal(body, resPb)
	if err != nil {
		// Mark as a permanent error so the backoff library doesn't retry this REST call
		err := backoff.Permanent(err)
		logger.WithFields(headerFields).WithFields(logrus.Fields{
			"response_body": string(body),
			"pb_message":    "CreateTicketsResponse",
			"error":         err,
		}).Errorf("cannot unmarshal http response body back into protobuf")
		return "", err
	}

	if resPb == nil {
		// Mark as a permanent error so the backoff library doesn't retry this REST call
		err := backoff.Permanent(errors.New("CreateTicket returned empty result"))
		logger.Error(err)
		return "", err
	}

	// Successful ticket creation
	logger.Debugf("CreateTicket %v complete", resPb.TicketId)
	return resPb.TicketId, err
}

// ActivateTickets takes a list of ticket ids for tickets awaiting activation,
// and breaks them into batches of the size defined in
// rc.cfg.GetInt("OM_CORE_MAX_UPDATES_PER_ACTIVATION_CALL") (see
// open-match.dev/core/internal/config for limits on how many actions you can
// request in a single API call, this needs to use the same value)
func (rc *RestfulOMGrpcClient) ActivateTickets(ctx context.Context, ticketIdsToActivate chan string) {

	logger := rc.Log.WithFields(logrus.Fields{
		"component": "open_match_client",
		"operation": "proxy_ActivateTickets",
	})

	// Activate all new tickets
	done := false
	ticketsAwaitingActivation := make([]string, 0)

	var activationWg sync.WaitGroup
	for len(ticketsAwaitingActivation) == rc.Cfg.GetInt("OM_CORE_MAX_UPDATES_PER_ACTIVATION_CALL") || !done {

		// Collect tickets from the channel.
		ticketsAwaitingActivation = nil
		for !done {
			select {
			case tid := <-ticketIdsToActivate:
				ticketsAwaitingActivation = append(ticketsAwaitingActivation, tid)
				if len(ticketsAwaitingActivation) == rc.Cfg.GetInt("OM_CORE_MAX_UPDATES_PER_ACTIVATION_CALL") {
					// maximum updates allowed per api call, go ahead and activate these,
					// then loop to grab what is left in the channel.
					done = true
				}
			default:
				done = true
			}
		}

		// We've got tickets to activate
		if len(ticketsAwaitingActivation) > 0 {
			logger.Debugf("ActivateTicket call with %v tickets: %v...", len(ticketsAwaitingActivation), ticketsAwaitingActivation[0])
			if len(ticketsAwaitingActivation) == rc.Cfg.GetInt("OM_CORE_MAX_UPDATES_PER_ACTIVATION_CALL") {
				done = false
			}

			// Kick off activation in a goroutine in case we had to split on
			// rc.cfg.GetInt("OM_CORE_MAX_UPDATES_PER_ACTIVATION_CALL"), this way the calls are made concurrently.
			go func(ticketsAwaitingActivation []string) {

				// Quick in-line function to grab information
				// the calling function put into the context when
				// logging errors (to aid debugging)
				callerFromContext := func() string {
					ts, ok := ctx.Value("type").(string)
					if !ok {
						logger.Error("unable to get caller type from context")
						return "undefined"
					}
					return ts
				}

				// Track this goroutine
				activationWg.Add(1)
				defer activationWg.Done()

				// Activate tickets
				reqPb := &pb.ActivateTicketsRequest{TicketIds: ticketsAwaitingActivation}
				var err error

				// Marshal request into json
				req, err := protojson.Marshal(reqPb)
				if err != nil {
					logger.WithFields(logrus.Fields{
						"caller":     callerFromContext,
						"pb_message": "ActivateTicketsRequest",
					}).Errorf("cannot marshal protobuf to json")
				}

				// HTTP version of the gRPC ActivateTickets() call that we want to retry with exponential backoff and jitter
				activateTicketCall := func() error {
					logger.Tracef("!! %v Tix awaiting activation", len(ticketsAwaitingActivation))
					resp, err := rc.Post(ctx, logger, rc.Cfg.GetString("OM_CORE_ADDR"), "/tickets:activate", req)
					// TODO: In reality, the error has details fields, telling us which ticket couldn't
					// be activated, but we're not processing those or passing them on yet, we just act as
					// though it is all-or-nothing.
					if err != nil {
						logger.WithFields(logrus.Fields{
							"caller": callerFromContext,
						}).Errorf("ActivateTickets attempt failed: %w", err)

						return err
					} else if resp != nil && resp.StatusCode != http.StatusOK { // HTTP error code
						logger.WithFields(logrus.Fields{
							"caller": callerFromContext,
						}).Errorf("ActivateTickets attempt failed: %w", fmt.Errorf("%s (%d)", http.StatusText(resp.StatusCode), resp.StatusCode))
						return fmt.Errorf("%s (%d)", http.StatusText(resp.StatusCode), resp.StatusCode)
					}
					return nil
				}

				//err = activateTicketCall()
				err = backoff.RetryNotify(
					activateTicketCall,
					backoff.NewExponentialBackOff(backoff.WithMaxElapsedTime(5*time.Second)),
					func(err error, bo time.Duration) {
						// The function that gets called to notify us when there's an activateTicketCall failure, before it is retried.
						logger.Errorf("ActivateTicket temporary failure (backoff for %v): %v", err, bo)
					})
				if err != nil {
					logger.WithFields(logrus.Fields{
						"caller": callerFromContext,
					}).Errorf("ActivateTickets failed: %w", err)
				}
				logger.WithFields(logrus.Fields{
					"caller": callerFromContext,
				}).Debug("ActivateTickets complete")
			}(ticketsAwaitingActivation)
		}
	}

	// After gathering all the activations and making all the necessary
	// calls, wait on the waitgroup to finish.
	activationWg.Wait()
}

func (rc *RestfulOMGrpcClient) InvokeMatchmakingFunctions(ctx context.Context, reqPb *pb.MmfRequest, respChan chan *pb.StreamedMmfResponse) {
	// Make sure the output channel always gets closed when this function is finished.
	defer close(respChan)

	logger := rc.Log.WithFields(logrus.Fields{
		"component": "open_match_client",
		"operation": "proxy_InvokeMatchmakingFunctions",
	})

	// Have to cancel the context to tell the om-core server we're done reading from the stream.
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Marshal protobuf request to JSON for grpc-gateway HTTP call
	req, err := protojson.Marshal(reqPb)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"pb_message": "MmfRequest",
		}).Errorf("cannot marshal protobuf to json")
	} else {
		logger.Trace("Marshalled MmfRequest protobuf message to JSON")
	}

	// HTTP version of gRPC InvokeMatchmakingFunctions() call
	logger.Trace("Calling InvokeMatchmakingFunctions()")
	resp, err := rc.Post(ctx, logger, rc.Cfg.GetString("OM_CORE_ADDR"), "/matches:fetch", req)
	if err != nil {
		logger.Errorf("InvokeMatchmakingFunction call failed: %v", err)
	} else {
		logger.Tracef("InvokeMatchmakingFunction call successful: %v", resp.StatusCode)

		// Check for a successful HTTP status code
		if resp.StatusCode != http.StatusOK {
			logger.Fatalf("Request failed with status: %s", resp.Status)
		}

		// Process the result stream
		for {
			var resPb *pb.StreamedMmfResponse
			// Retrieve JSON-formatted protobuf from response body
			var msg string
			_, err := fmt.Fscanf(resp.Body, "%s\n", &msg)

			// Validate response before processing
			if err == io.EOF {
				break // End of stream; done!
			}
			if err != nil {
				logger.Errorf("Error reading stream, closing: %v", err)
				break
			}

			// Easy to miss: grpc-gateway returns your result inside a
			// overarching JSON enclosure under the key 'result', so
			// the actual text we need to pass to protojson.Unmarshal
			// needs to omit the top level JSON object. Rather than
			// marshal the text to json (which is slow), we just use
			// string trimming functions.
			trimmedMsg := strings.TrimSuffix(strings.TrimPrefix(msg, `{"result":`), "}")
			logger.WithFields(logrus.Fields{
				"pb_message": "StreamedMmfResponse",
			}).Tracef("HTTP response JSON body version of StreamedMmfResponse received")

			// Unmarshal json back into protobuf
			resPb = &pb.StreamedMmfResponse{}
			err = protojson.Unmarshal([]byte(trimmedMsg), resPb)
			if err != nil {
				logger.WithFields(logrus.Fields{
					"pb_message":    "StreamedMmfResponse",
					"http_response": trimmedMsg,
					"error":         err,
				}).Errorf("cannot unmarshal http response body back into protobuf")
				continue
			} else {
				logger.Trace("Successfully unmarshalled HTTP response JSON body back into StreamedMmfResponse protobuf message")
			}

			if resPb == nil {
				logger.Trace("StreamedMmfResponse protobuf was nil!")
				continue // Loop again to get the next streamed response
			}

			// Send back the streamed responses as we get them
			logger.Trace("StreamedMmfResponse protobuf exists")
			respChan <- resPb

		}

		// Close response body
		resp.Body.Close()
	}
}

// type idTokenSource struct {
// 	TokenSource oauth2.TokenSource
// }

// func (s *idTokenSource) Token() (*oauth2.Token, error) {
// 	token, err := s.TokenSource.Token()
// 	if err != nil {
// 		return nil, err
// 	}

// 	idToken, ok := token.Extra("id_token").(string)
// 	if !ok {
// 		return nil, fmt.Errorf("token did not contain an id_token")
// 	}

// 	return &oauth2.Token{
// 		AccessToken: idToken,
// 		TokenType:   "Bearer",
// 		Expiry:      token.Expiry,
// 	}, nil
// }

// Post Makes an HTTP request at the given url+path, marshalling the
// provided protobuf in the pbReq argument into the HTTP request JSON body (for
// use with grpc gateway). It attempts to transparently handle TLS if the target
// server requires it.
func (rc *RestfulOMGrpcClient) Post(ctx context.Context, logger *logrus.Entry, url string, path string, reqBuf []byte) (*http.Response, error) {
	// Add the url as a structured logging field when debug logging is enabled
	postLogger := logger.WithFields(logrus.Fields{
		"url": url + path,
	})

	// Set up our request parameters
	req, err := http.NewRequestWithContext(
		ctx,                     // Context
		http.MethodPost,         // HTTP verb
		url+path,                // RESTful OM2 path
		bytes.NewReader(reqBuf), // JSON-marshalled protobuf request message
	)
	if err != nil {
		postLogger.Errorf("cannot create http request with context")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	postLogger.Trace("Header set, request created")

	// if os.Getenv("K_REVISION") != "" { // Running on Cloud Run, get GCP SA token
	if rc.tokenSource == nil {
		// Create a TokenSource if none exists.
		rc.tokenSource, err = idtoken.NewTokenSource(ctx, url)
		if err != nil {
			fmt.Println("Cloud Run service account authentication requires token source, but couldn't get one. Trying to continue without SA auth. idtoken.NewTokenSource: %w", err)
		}
	}

	if rc.tokenSource == nil {
		fmt.Println("rc.tokenSource is nil")
	}

	// Retrieve an identity token. Will reuse tokens until refresh needed.
	token, err := rc.tokenSource.Token()
	if err != nil {
		fmt.Println("Cloud Run service account authentication requires a token, but couldn't get one. Trying to continue without SA auth. TokenSource.Token: %w", err)
	}
	token.SetAuthHeader(req)
	// }

	// client, err := idtoken.NewClient(ctx, url)
	// if err != nil {
	// 	fmt.Println("Error connecting to client: ", err)
	// }

	// Log all request headers if debug logging (slow)
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		headerFields := logrus.Fields{}
		for key, value := range req.Header {
			headerFields[key] = value
		}
		postLogger.WithFields(headerFields).Debug("Prepared Headers")
	}

	// Send request
	resp, err := rc.Client.Do(req)
	if err != nil {
		postLogger.WithFields(logrus.Fields{
			"error": err,
		}).Error("cannot execute http request")
		return nil, err
	}
	return resp, err
}

// readAllBody is a simple helper function to make sure an HTTP body is completely read and closed.
func readAllBody(resp http.Response, logger *logrus.Entry) ([]byte, error) {
	// Get results
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		logger.Errorf("cannot read bytes from http response body")
		return nil, err
	}

	return body, err
}

func syncMapDump(sm *sync.Map) map[string]interface{} {
	out := map[string]interface{}{}
	sm.Range(func(key, value interface{}) bool {
		out[fmt.Sprint(key)] = value
		return true
	})
	return out
}
