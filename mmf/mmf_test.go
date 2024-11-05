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

	pb2 "github.com/googleforgames/open-match2/v2/pkg/pb"
	"github.com/stretchr/testify/assert"
)

func TestNewMatches(t *testing.T) {
	pool := &pb2.Pool{
		Name: "everyone",
	}
	profile := &pb2.Profile{
		Name:  "test-profile",
		Pools: map[string]*pb2.Pool{"everyone": pool},
	}

	poolTickets := map[string][]*pb2.Ticket{
		pool.Name: {
			&pb2.Ticket{Id: "ticket-1"},
			&pb2.Ticket{Id: "ticket-2"},
		},
	}
	matches, err := makeMatches(profile, poolTickets)
	if err != nil {
		t.Errorf("Create new match proposal failed: %v", err)
	}

	assert.NoError(t, err)
	assert.Len(t, matches, 1)
	assert.Len(t, matches[0].Rosters[pool.Name].Tickets, len(poolTickets[pool.Name]))
	assert.Contains(t, matches[0].Id, "profile-"+profile.Name+"-time")
	assert.Equal(t, profile.Name, matches[0].Rosters[pool.Name].Name)
}
