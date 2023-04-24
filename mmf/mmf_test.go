// Copyright 2023 Google LLC
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
	"testing"

	"github.com/stretchr/testify/assert"
	"open-match.dev/open-match/pkg/pb"
)

func TestNewMatches(t *testing.T) {
	pool := &pb.Pool{
		Name: "everyone",
	}
	profile := &pb.MatchProfile{
		Name:  "test-profile",
		Pools: []*pb.Pool{pool},
	}

	poolTickets := map[string][]*pb.Ticket{
		pool.Name: {
			&pb.Ticket{Id: "ticket-1"},
			&pb.Ticket{Id: "ticket-2"},
		},
	}
	matches, err := makeMatches(profile, poolTickets)
	if err != nil {
		t.Errorf("Create new match proposal failed: %v", err)
	}

	assert.NoError(t, err)
	assert.Len(t, matches, 1)
	assert.Len(t, matches[0].Tickets, len(poolTickets[pool.Name]))
	assert.Contains(t, matches[0].MatchId, "profile-"+profile.Name+"-time")
	assert.Equal(t, matches[0].MatchProfile, profile.Name)
	assert.Equal(t, matches[0].MatchFunction, matchName)
}
