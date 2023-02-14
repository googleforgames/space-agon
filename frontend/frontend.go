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

	"github.com/googleforgames/space-agon/game/protostream"
	"golang.org/x/net/websocket"
	"google.golang.org/grpc"
	"open-match.dev/open-match/pkg/pb"
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
	assignments := make(chan *pb.Assignment)
	errs := make(chan error)

	go streamAssignments(ctx, assignments, errs)

	for {
		select {
		case err := <-errs:
			log.Println("Error getting assgnment:", err)
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

func streamAssignments(ctx context.Context, assignments chan *pb.Assignment, errs chan error) {
	conn, err := grpc.Dial("open-match-frontend.open-match.svc.cluster.local:50504", grpc.WithInsecure())
	if err != nil {
		errs <- fmt.Errorf("Error dialing open match: %w", err)
	}
	defer conn.Close()
	fe := pb.NewFrontendServiceClient(conn)

	var ticketId string
	{
		req := &pb.CreateTicketRequest{
			Ticket: &pb.Ticket{},
		}

		resp, err := fe.CreateTicket(ctx, req)
		if err != nil {
			errs <- fmt.Errorf("Error creating open match ticket: %w", err)
			return
		}
		ticketId = resp.Id
	}

	log.Println("Ticket ID: ", ticketId)

	defer func() {
		_, err := fe.DeleteTicket(context.Background(), &pb.DeleteTicketRequest{TicketId: ticketId})
		if err != nil {
			log.Println("Error deleting ticket", ticketId, ":", err)
		}
	}()

	{
		req := &pb.WatchAssignmentsRequest{
			TicketId: ticketId,
		}

		log.Println("Assignment request: ", req)

		stream, err := fe.WatchAssignments(ctx, req)
		if err != nil {
			errs <- fmt.Errorf("Error getting assignment stream: %w", err)
			return
		}
		for {
			resp, err := stream.Recv()
			if err != nil {
				errs <- fmt.Errorf("Error streaming assignment: %w", err)
				return
			}
			assignments <- resp.Assignment

			log.Println("Assignments: ", assignments)
		}
	}
}
