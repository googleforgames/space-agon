// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: /open-match/pkg/pb/backend_grpc.pb.go

// Package mockpb is a generated GoMock package.
package mockpb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	pb "open-match.dev/open-match/pkg/pb"
)

// MockBackendServiceClient is a mock of BackendServiceClient interface.
type MockBackendServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockBackendServiceClientMockRecorder
}

// MockBackendServiceClientMockRecorder is the mock recorder for MockBackendServiceClient.
type MockBackendServiceClientMockRecorder struct {
	mock *MockBackendServiceClient
}

// NewMockBackendServiceClient creates a new mock instance.
func NewMockBackendServiceClient(ctrl *gomock.Controller) *MockBackendServiceClient {
	mock := &MockBackendServiceClient{ctrl: ctrl}
	mock.recorder = &MockBackendServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackendServiceClient) EXPECT() *MockBackendServiceClientMockRecorder {
	return m.recorder
}

// AssignTickets mocks base method.
func (m *MockBackendServiceClient) AssignTickets(ctx context.Context, in *pb.AssignTicketsRequest, opts ...grpc.CallOption) (*pb.AssignTicketsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AssignTickets", varargs...)
	ret0, _ := ret[0].(*pb.AssignTicketsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignTickets indicates an expected call of AssignTickets.
func (mr *MockBackendServiceClientMockRecorder) AssignTickets(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignTickets", reflect.TypeOf((*MockBackendServiceClient)(nil).AssignTickets), varargs...)
}

// FetchMatches mocks base method.
func (m *MockBackendServiceClient) FetchMatches(ctx context.Context, in *pb.FetchMatchesRequest, opts ...grpc.CallOption) (pb.BackendService_FetchMatchesClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchMatches", varargs...)
	ret0, _ := ret[0].(pb.BackendService_FetchMatchesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMatches indicates an expected call of FetchMatches.
func (mr *MockBackendServiceClientMockRecorder) FetchMatches(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMatches", reflect.TypeOf((*MockBackendServiceClient)(nil).FetchMatches), varargs...)
}

// ReleaseAllTickets mocks base method.
func (m *MockBackendServiceClient) ReleaseAllTickets(ctx context.Context, in *pb.ReleaseAllTicketsRequest, opts ...grpc.CallOption) (*pb.ReleaseAllTicketsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReleaseAllTickets", varargs...)
	ret0, _ := ret[0].(*pb.ReleaseAllTicketsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseAllTickets indicates an expected call of ReleaseAllTickets.
func (mr *MockBackendServiceClientMockRecorder) ReleaseAllTickets(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseAllTickets", reflect.TypeOf((*MockBackendServiceClient)(nil).ReleaseAllTickets), varargs...)
}

// ReleaseTickets mocks base method.
func (m *MockBackendServiceClient) ReleaseTickets(ctx context.Context, in *pb.ReleaseTicketsRequest, opts ...grpc.CallOption) (*pb.ReleaseTicketsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReleaseTickets", varargs...)
	ret0, _ := ret[0].(*pb.ReleaseTicketsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseTickets indicates an expected call of ReleaseTickets.
func (mr *MockBackendServiceClientMockRecorder) ReleaseTickets(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseTickets", reflect.TypeOf((*MockBackendServiceClient)(nil).ReleaseTickets), varargs...)
}

// MockBackendService_FetchMatchesClient is a mock of BackendService_FetchMatchesClient interface.
type MockBackendService_FetchMatchesClient struct {
	ctrl     *gomock.Controller
	recorder *MockBackendService_FetchMatchesClientMockRecorder
}

// MockBackendService_FetchMatchesClientMockRecorder is the mock recorder for MockBackendService_FetchMatchesClient.
type MockBackendService_FetchMatchesClientMockRecorder struct {
	mock *MockBackendService_FetchMatchesClient
}

// NewMockBackendService_FetchMatchesClient creates a new mock instance.
func NewMockBackendService_FetchMatchesClient(ctrl *gomock.Controller) *MockBackendService_FetchMatchesClient {
	mock := &MockBackendService_FetchMatchesClient{ctrl: ctrl}
	mock.recorder = &MockBackendService_FetchMatchesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackendService_FetchMatchesClient) EXPECT() *MockBackendService_FetchMatchesClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockBackendService_FetchMatchesClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockBackendService_FetchMatchesClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockBackendService_FetchMatchesClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockBackendService_FetchMatchesClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockBackendService_FetchMatchesClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBackendService_FetchMatchesClient)(nil).Context))
}

// Header mocks base method.
func (m *MockBackendService_FetchMatchesClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockBackendService_FetchMatchesClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockBackendService_FetchMatchesClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockBackendService_FetchMatchesClient) Recv() (*pb.FetchMatchesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*pb.FetchMatchesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockBackendService_FetchMatchesClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockBackendService_FetchMatchesClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockBackendService_FetchMatchesClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockBackendService_FetchMatchesClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBackendService_FetchMatchesClient)(nil).RecvMsg), m)
}

// SendMsg mocks base method.
func (m_2 *MockBackendService_FetchMatchesClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockBackendService_FetchMatchesClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBackendService_FetchMatchesClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockBackendService_FetchMatchesClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockBackendService_FetchMatchesClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockBackendService_FetchMatchesClient)(nil).Trailer))
}

// MockBackendServiceServer is a mock of BackendServiceServer interface.
type MockBackendServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockBackendServiceServerMockRecorder
}

// MockBackendServiceServerMockRecorder is the mock recorder for MockBackendServiceServer.
type MockBackendServiceServerMockRecorder struct {
	mock *MockBackendServiceServer
}

// NewMockBackendServiceServer creates a new mock instance.
func NewMockBackendServiceServer(ctrl *gomock.Controller) *MockBackendServiceServer {
	mock := &MockBackendServiceServer{ctrl: ctrl}
	mock.recorder = &MockBackendServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackendServiceServer) EXPECT() *MockBackendServiceServerMockRecorder {
	return m.recorder
}

// AssignTickets mocks base method.
func (m *MockBackendServiceServer) AssignTickets(arg0 context.Context, arg1 *pb.AssignTicketsRequest) (*pb.AssignTicketsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignTickets", arg0, arg1)
	ret0, _ := ret[0].(*pb.AssignTicketsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignTickets indicates an expected call of AssignTickets.
func (mr *MockBackendServiceServerMockRecorder) AssignTickets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignTickets", reflect.TypeOf((*MockBackendServiceServer)(nil).AssignTickets), arg0, arg1)
}

// FetchMatches mocks base method.
func (m *MockBackendServiceServer) FetchMatches(arg0 *pb.FetchMatchesRequest, arg1 pb.BackendService_FetchMatchesServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMatches", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchMatches indicates an expected call of FetchMatches.
func (mr *MockBackendServiceServerMockRecorder) FetchMatches(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMatches", reflect.TypeOf((*MockBackendServiceServer)(nil).FetchMatches), arg0, arg1)
}

// ReleaseAllTickets mocks base method.
func (m *MockBackendServiceServer) ReleaseAllTickets(arg0 context.Context, arg1 *pb.ReleaseAllTicketsRequest) (*pb.ReleaseAllTicketsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseAllTickets", arg0, arg1)
	ret0, _ := ret[0].(*pb.ReleaseAllTicketsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseAllTickets indicates an expected call of ReleaseAllTickets.
func (mr *MockBackendServiceServerMockRecorder) ReleaseAllTickets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseAllTickets", reflect.TypeOf((*MockBackendServiceServer)(nil).ReleaseAllTickets), arg0, arg1)
}

// ReleaseTickets mocks base method.
func (m *MockBackendServiceServer) ReleaseTickets(arg0 context.Context, arg1 *pb.ReleaseTicketsRequest) (*pb.ReleaseTicketsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseTickets", arg0, arg1)
	ret0, _ := ret[0].(*pb.ReleaseTicketsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseTickets indicates an expected call of ReleaseTickets.
func (mr *MockBackendServiceServerMockRecorder) ReleaseTickets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseTickets", reflect.TypeOf((*MockBackendServiceServer)(nil).ReleaseTickets), arg0, arg1)
}

// MockUnsafeBackendServiceServer is a mock of UnsafeBackendServiceServer interface.
type MockUnsafeBackendServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeBackendServiceServerMockRecorder
}

// MockUnsafeBackendServiceServerMockRecorder is the mock recorder for MockUnsafeBackendServiceServer.
type MockUnsafeBackendServiceServerMockRecorder struct {
	mock *MockUnsafeBackendServiceServer
}

// NewMockUnsafeBackendServiceServer creates a new mock instance.
func NewMockUnsafeBackendServiceServer(ctrl *gomock.Controller) *MockUnsafeBackendServiceServer {
	mock := &MockUnsafeBackendServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeBackendServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeBackendServiceServer) EXPECT() *MockUnsafeBackendServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedBackendServiceServer mocks base method.
func (m *MockUnsafeBackendServiceServer) mustEmbedUnimplementedBackendServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedBackendServiceServer")
}

// mustEmbedUnimplementedBackendServiceServer indicates an expected call of mustEmbedUnimplementedBackendServiceServer.
func (mr *MockUnsafeBackendServiceServerMockRecorder) mustEmbedUnimplementedBackendServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedBackendServiceServer", reflect.TypeOf((*MockUnsafeBackendServiceServer)(nil).mustEmbedUnimplementedBackendServiceServer))
}

// MockBackendService_FetchMatchesServer is a mock of BackendService_FetchMatchesServer interface.
type MockBackendService_FetchMatchesServer struct {
	ctrl     *gomock.Controller
	recorder *MockBackendService_FetchMatchesServerMockRecorder
}

// MockBackendService_FetchMatchesServerMockRecorder is the mock recorder for MockBackendService_FetchMatchesServer.
type MockBackendService_FetchMatchesServerMockRecorder struct {
	mock *MockBackendService_FetchMatchesServer
}

// NewMockBackendService_FetchMatchesServer creates a new mock instance.
func NewMockBackendService_FetchMatchesServer(ctrl *gomock.Controller) *MockBackendService_FetchMatchesServer {
	mock := &MockBackendService_FetchMatchesServer{ctrl: ctrl}
	mock.recorder = &MockBackendService_FetchMatchesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackendService_FetchMatchesServer) EXPECT() *MockBackendService_FetchMatchesServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockBackendService_FetchMatchesServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockBackendService_FetchMatchesServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBackendService_FetchMatchesServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockBackendService_FetchMatchesServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockBackendService_FetchMatchesServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBackendService_FetchMatchesServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockBackendService_FetchMatchesServer) Send(arg0 *pb.FetchMatchesResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockBackendService_FetchMatchesServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBackendService_FetchMatchesServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockBackendService_FetchMatchesServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockBackendService_FetchMatchesServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockBackendService_FetchMatchesServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockBackendService_FetchMatchesServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockBackendService_FetchMatchesServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBackendService_FetchMatchesServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockBackendService_FetchMatchesServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockBackendService_FetchMatchesServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockBackendService_FetchMatchesServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockBackendService_FetchMatchesServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockBackendService_FetchMatchesServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockBackendService_FetchMatchesServer)(nil).SetTrailer), arg0)
}
