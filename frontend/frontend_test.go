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
	"context"
	"fmt"
	"log"
	"net"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/googleforgames/space-agon/game/protostream"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	ompb "open-match.dev/open-match/pkg/pb"
	omtest "open-match.dev/open-match/testing"
)

func setupFrontendMock() (*grpc.Server, net.Listener, error) {
	var l net.Listener
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}
	_, err = grpc.Dial(l.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, l, fmt.Errorf("error dialiing to mock: %w", err)
	}
	gsrv := grpc.NewServer()
	ff := omtest.FakeFrontend{}
	ompb.RegisterFrontendServiceServer(gsrv, &ff)

	// Run grpc mock server
	go func() {
		log.Println("Mock server start:", l.Addr())
		if err = gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	return gsrv, l, nil
}

func TestMain(m *testing.M) {
	// Setup for test
	gsrv, l, err := setupFrontendMock()
	defer func() {
		gsrv.Stop()
		l.Close()
	}()
	if err != nil {
		log.Fatalf("Frontend Mockserver start failed: %v", err)
	}
	err = os.Setenv("FRONTEND_ADDR", l.Addr().String())
	if err != nil {
		log.Fatalf("Set FRONTEND_ADDR env failed: %v", err)
	}

	m.Run()
}

func TestMatchmake(t *testing.T) {
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

func TestConnectFrontendServer(t *testing.T) {
	conn, err := connectFrontendServer()
	defer func() {
		err = conn.Close()
		if err != nil {
			t.Errorf("Close grpc connection failed: %v", err)
		}
	}()
	assert.NoError(t, err)
	t.Logf("Connected to the frontend mockserver")
}
