/*
 Copyright 2022 Google LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"golang.org/x/net/websocket"

	"github.com/googleforgames/space-agon/game/pb"
	"github.com/googleforgames/space-agon/game/protostream"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	ompb "open-match.dev/open-match/pkg/pb"
)

const (
	NAMESPACE = "default"
	SVC       = "frontend"
)

var (
	pip    string
	gsconn string
)

func TestMain(m *testing.M) {
	// Confirm LB PIP
	confirmLBPIP()
	// Execute Test Cases
	m.Run()

}

func TestPvPAssignment(t *testing.T) {
	url := "ws://" + pip + "/matchmake/"
	origin := "http://" + pip + "/"

	achan := getAssignment(url, origin)
	achan2 := getAssignment(url, origin)

	a1, a2 := <-achan, <-achan2

	if a1.Connection != a2.Connection {
		t.Errorf("Different Assignments !!")
	}
	gsconn = a1.Connection
}

func TestConnectGameServer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	url := "ws://" + gsconn + "/connect/"
	origin := "http://" + gsconn
	cid1 := connectGameServer(url, origin, ctx)
	cid2 := connectGameServer(url, origin, ctx)

	if cid1 != 1 {
		t.Errorf("First ClientID should be 1")
	}

	if cid2 != 2 {
		t.Errorf("First ClientID should be 2")
	}
}

func connectGameServer(url string, origin string, ctx context.Context) int64 {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		panic(err)
	}
	stream := protostream.NewProtoStream(ws)
	clientInitialize := &pb.ClientInitialize{}
	err = stream.Recv(clientInitialize)
	if err != nil {
		panic(err)
	}
	cid := clientInitialize.Cid

	go func() {
	LOOP:
		for {
			select {
			case <-ctx.Done():
				ws.Close()
				break LOOP
			}
		}
	}()
	return cid
}

func getAssignment(url string, origin string) <-chan ompb.Assignment {
	ch := make(chan ompb.Assignment)
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		panic(err)
	}
	stream := protostream.NewProtoStream(ws)
	as := &ompb.Assignment{}
	go func() {
		defer ws.Close()
		defer close(ch)
		for {
			err := stream.Recv(as)
			if err != nil {
				fmt.Printf("Failed to receive an assignment with error: %s. Wait for 1 second.\n", err.Error())
				time.Sleep(time.Second * 1)
				continue
			}
			fmt.Printf("Conntection: %v.\n", as.Connection)
			ch <- *as
			break
		}
	}()
	return ch
}

func confirmLBPIP() {
	var kubeconfig string
	ctx := context.Background()

	if home := os.Getenv("HOME"); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	svcClient := clientset.CoreV1().Services(NAMESPACE)
	for {
		result, err := svcClient.Get(ctx, SVC, metav1.GetOptions{})
		if err != nil {
			fmt.Printf("Unable to get public IP from svc: %s, error: %s\n", SVC, err.Error())
			time.Sleep(time.Second * 1)
			continue
		}
		if len(result.Status.LoadBalancer.Ingress) == 0 {
			fmt.Printf("Waiting for assignment of a PIP\n")
			time.Sleep(time.Second * 1)
			continue
		} else {
			pip = result.Status.LoadBalancer.Ingress[0].IP
			fmt.Printf("PIP is %v\n", pip)
			break
		}
	}
}
