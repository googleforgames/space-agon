// Copyright 2019 Google LLC
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
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	pb2 "github.com/googleforgames/open-match2/v2/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Initializing")

	mmf := &matchFunctionService{}

	server := grpc.NewServer()
	pb2.RegisterMatchMakingFunctionServiceServer(server, mmf)

	ln, err := net.Listen("tcp", ":50502")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Serving started v1")
	err = server.Serve(ln)
	log.Printf("Err status of servering grpc: %s", err.Error())
}

type matchFunctionService struct {
	pb2.UnimplementedMatchMakingFunctionServiceServer
}

func GetChunkedRequest(stream pb2.MatchMakingFunctionService_RunServer) (*pb2.Profile, error) {
	var req *pb2.Profile
	pools := make(map[string]*pb2.Pool)
	for i := int32(1); i >= 0; i++ {
		in, err := stream.Recv()
		if err != nil {
			// TODO: Check if we got any portion of a valid profile, if so, attempt a run
			return nil, err
		}

		fmt.Println("Processing chunk ", i, "/", in.GetNumChunks())
		for name, pool := range in.GetProfile().GetPools() {
			if _, ok := pools[name]; !ok {
				// First chunk containing this pool; initialize the local copy.
				pools[name] = &pb2.Pool{
					Name:                    pool.GetName(),
					TagPresentFilters:       pool.GetTagPresentFilters(),
					StringEqualsFilters:     pool.GetStringEqualsFilters(),
					DoubleRangeFilters:      pool.GetDoubleRangeFilters(),
					CreationTimeRangeFilter: pool.GetCreationTimeRangeFilter(),
					Extensions:              pool.GetExtensions(),
					Participants:            pool.GetParticipants(),
				}
			} else {
				// concate pools split amoung multiple chunks
				pools[name].Participants.Tickets = append(pools[name].Participants.Tickets, pool.GetParticipants().GetTickets()...)
			}
		}

		if in.GetNumChunks() == i {
			req = &pb2.Profile{
				Name:       in.GetProfile().GetName(),
				Pools:      pools,
				Extensions: in.GetProfile().GetExtensions(),
			}
			// fmt.Println("Finished receiving %v chunks of MMF profile %v", in.GetProfile().GetName(), i)
			break
		}
	}
	return req, nil
}

func (mmf *matchFunctionService) Run(stream pb2.MatchMakingFunctionService_RunServer) error {
	log.Printf("Running mmf")

	req, err := GetChunkedRequest(stream)
	if err != nil {
		fmt.Println("error getting chunked request: ", err)
	}
	fmt.Println("Generating matches for profile ", req.GetName())

	// Process tickets for the pools specified in the Match Profile.  In this
	// example fifo matching strategy, any player can be matched with any other
	// player, so just concatinate all the pools together.
	tickets := []*pb2.Ticket{}
	for pname, pool := range req.GetPools() {
		tickets = append(tickets, pool.GetParticipants().GetTickets()...)
		fmt.Printf("Found %v tickets in pool %v", len(tickets), pname)
	}
	fmt.Printf("Matching among %v tickets from %v provided pools", len(tickets), len(req.GetPools()))

	// Make match proposal
	poolTickets := make(map[string][]*pb2.Ticket)
	poolTickets["everyone"] = tickets
	proposals, err := makeMatches(req, poolTickets)
	if err != nil {
		return err
	}

	matchesFound := 0
	for _, proposal := range proposals {
		err = stream.Send(&pb2.StreamedMmfResponse{Match: proposal})
		if err != nil {
			return err
		}
		matchesFound++
	}
	log.Printf("MMF ran creating %d matches", matchesFound)

	return nil
}

func makeMatches(profile *pb2.Profile, poolTickets map[string][]*pb2.Ticket) ([]*pb2.Match, error) {
	var matches []*pb2.Match

	tickets, ok := poolTickets["everyone"]
	if !ok {
		return matches, errors.New("expected pool named everyone")
	}

	t := time.Now().Format("2006-01-02T15:04:05.00")

	for i := 0; i+1 < len(tickets); i += 2 {
		proposal := &pb2.Match{
			Id: fmt.Sprintf("profile-%s-time-%s-num-%d", profile.Name, t, i/2),
			Rosters: map[string]*pb2.Roster{
				"everyone": {
					Name:    profile.Name,
					Tickets: []*pb2.Ticket{tickets[i], tickets[i+1]},
				},
			},
		}
		matches = append(matches, proposal)
	}
	return matches, nil
}
