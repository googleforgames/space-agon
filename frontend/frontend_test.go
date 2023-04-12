package main

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/googleforgames/space-agon/game/protostream"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/websocket"
	"google.golang.org/grpc/codes"
	ompb "open-match.dev/open-match/pkg/pb"
)

func TestMatchmake(t *testing.T) {
	ch := make(chan ompb.Assignment)
	defer close(ch)
	s := httptest.NewServer(websocket.Handler(matchmake))
	defer s.Close()

	wsUrl := "ws" + strings.TrimPrefix(s.URL, "http") + "/matchmake"
	ws, err := websocket.Dial(wsUrl, "", s.URL)
	if err != nil {
		t.Fatalf("Connect to matchmake websocket failed: %v", err)
	}
	defer ws.Close()

	stream := protostream.NewProtoStream(ws)
	a := &ompb.Assignment{}
	err = stream.Recv(a)
	if err != nil {
		t.Errorf("error receiving assignment: %v", err)
	}

	// Assert
	assert.NoError(t, err)
}

func TestStreamAssignments(t *testing.T) {
	ch := make(chan *ompb.Assignment)
	defer close(ch)
	errs := make(chan error)
	defer close(errs)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	go streamAssignments(ctx, ch, errs)
	assert.NotEqual(t, codes.Unimplemented, <-errs)
	cancel()
}
