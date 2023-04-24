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
