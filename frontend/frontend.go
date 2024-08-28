// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	pb2 "github.com/googleforgames/open-match2/v2/pkg/pb"
	"github.com/googleforgames/space-agon/game/protostream"
	"golang.org/x/net/websocket"
	"google.golang.org/grpc"
	"open-match.dev/open-match/pkg/pb"
)

const (
	defaultFrontendAddress = "https://om-core-976869741551.asia-northeast1.run.app:443"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/app/static/"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			log.Println("Serving index page.")
			http.ServeFile(w, r, "/app/static/index.html")
		} else {
			log.Println("404 on", r.URL.Path)
			http.NotFound(w, r)
		}
	})

	http.Handle("/matchmake/", websocket.Handler(matchmake))

	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func matchmake(ws *websocket.Conn) {
	ws.PayloadType = 2 // Sets sent payloads to binary
	stream := protostream.NewProtoStream(ws)

	ctx, cancel := context.WithCancel(ws.Request().Context())
	defer cancel()
	assignments := make(chan *pb2.Assignment)
	errs := make(chan error)

	go streamAssignments(ctx, assignments, errs)

	for {
		select {
		case err := <-errs:
			log.Println("Error getting assignment:", err)
			//err = stream.Send(&pb.Assignment{Error: status.Convert(err).Proto()})
			err = stream.Send(&pb.Assignment{})
			if err != nil {
				log.Println("Error sending error:", err)
			}
			return
		case assigment := <-assignments:
			err := stream.Send(assigment)
			if err != nil {
				log.Println("Error sending updated assignment:", err)
				cancel()
				return
			}
		}
	}
}

func streamAssignments(ctx context.Context, assignments chan *pb2.Assignment, errs chan error) {
	log.Println("streaming assignments...")
	conn, err := connectFrontendServer()
	if err != nil {
		errs <- err
	}
	defer conn.Close()
	fe := pb2.NewOpenMatchServiceClient(conn)

	var ticketId string
	crReq := &pb2.CreateTicketRequest{
		Ticket: &pb2.Ticket{},
	}

	log.Println("creating a ticket vvvv...")
	resp, err := fe.CreateTicket(ctx, crReq)
	if err != nil {
		errs <- fmt.Errorf("error creating open match ticket: %w", err)
		return
	}
	if resp == nil {
		errs <- fmt.Errorf("create ticket resp is nullptr: %w", err)
		return
	}
	ticketId = resp.TicketId
	log.Println("ticket created, ticket id is: ", ticketId)

	log.Println("activating a ticket: ", ticketId)
	acReq := &pb2.ActivateTicketsRequest{
		TicketIds: []string{ticketId},
	}
	if _, err := fe.ActivateTickets(ctx, acReq); err != nil {
		errs <- fmt.Errorf("error activating open match ticket: %w", err)
		return
	}
	log.Println("ticket activated: ", ticketId)

	waReq := &pb2.WatchAssignmentsRequest{
		TicketIds: []string{ticketId},
	}

	log.Println("watching assignments: ", waReq)
	stream, err := fe.WatchAssignments(ctx, waReq)
	if err != nil {
		errs <- fmt.Errorf("error getting assignment stream: %w", err)
		return
	}

	var assignment *pb2.Assignment
	for assignment.GetConnection() == "" {
		resp, err := stream.Recv()
		if err != nil {
			errs <- fmt.Errorf("error streaming assignment: %w", err)
			return
		}
		assignment = resp.Assignment
	}

	assignments <- assignment
	log.Printf("Got assignment: %v", assignment)
}

func connectFrontendServer() (*grpc.ClientConn, error) {
	frontendAddr := os.Getenv("FRONTEND_ADDR")
	if frontendAddr == "" {
		frontendAddr = defaultFrontendAddress
	}
	// conn, err := grpc.Dial(frontendAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// auth, _ := oauth.NewApplicationDefault(context.Background(), "")
	conn, err := grpc.Dial(
		frontendAddr,
		grpc.WithAuthority(frontendAddr),
	)
	if err != nil {
		return nil, fmt.Errorf("error dialing open match: %w", err)
	}
	return conn, nil
}
