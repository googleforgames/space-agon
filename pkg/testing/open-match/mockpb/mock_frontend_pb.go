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
// Source: /open-match/pkg/pb/frontend_grpc.pb.go

// Package mockpb is a generated GoMock package.
package mockpb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pb "open-match.dev/open-match/pkg/pb"
)

// MockFrontendServiceClient is a mock of FrontendServiceClient interface.
type MockFrontendServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockFrontendServiceClientMockRecorder
}

// MockFrontendServiceClientMockRecorder is the mock recorder for MockFrontendServiceClient.
type MockFrontendServiceClientMockRecorder struct {
	mock *MockFrontendServiceClient
}

// NewMockFrontendServiceClient creates a new mock instance.
func NewMockFrontendServiceClient(ctrl *gomock.Controller) *MockFrontendServiceClient {
	mock := &MockFrontendServiceClient{ctrl: ctrl}
	mock.recorder = &MockFrontendServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFrontendServiceClient) EXPECT() *MockFrontendServiceClientMockRecorder {
	return m.recorder
}

// AcknowledgeBackfill mocks base method.
func (m *MockFrontendServiceClient) AcknowledgeBackfill(ctx context.Context, in *pb.AcknowledgeBackfillRequest, opts ...grpc.CallOption) (*pb.AcknowledgeBackfillResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AcknowledgeBackfill", varargs...)
	ret0, _ := ret[0].(*pb.AcknowledgeBackfillResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcknowledgeBackfill indicates an expected call of AcknowledgeBackfill.
func (mr *MockFrontendServiceClientMockRecorder) AcknowledgeBackfill(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcknowledgeBackfill", reflect.TypeOf((*MockFrontendServiceClient)(nil).AcknowledgeBackfill), varargs...)
}

// CreateBackfill mocks base method.
func (m *MockFrontendServiceClient) CreateBackfill(ctx context.Context, in *pb.CreateBackfillRequest, opts ...grpc.CallOption) (*pb.Backfill, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateBackfill", varargs...)
	ret0, _ := ret[0].(*pb.Backfill)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBackfill indicates an expected call of CreateBackfill.
func (mr *MockFrontendServiceClientMockRecorder) CreateBackfill(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBackfill", reflect.TypeOf((*MockFrontendServiceClient)(nil).CreateBackfill), varargs...)
}

// CreateTicket mocks base method.
func (m *MockFrontendServiceClient) CreateTicket(ctx context.Context, in *pb.CreateTicketRequest, opts ...grpc.CallOption) (*pb.Ticket, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTicket", varargs...)
	ret0, _ := ret[0].(*pb.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTicket indicates an expected call of CreateTicket.
func (mr *MockFrontendServiceClientMockRecorder) CreateTicket(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTicket", reflect.TypeOf((*MockFrontendServiceClient)(nil).CreateTicket), varargs...)
}

// DeleteBackfill mocks base method.
func (m *MockFrontendServiceClient) DeleteBackfill(ctx context.Context, in *pb.DeleteBackfillRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteBackfill", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBackfill indicates an expected call of DeleteBackfill.
func (mr *MockFrontendServiceClientMockRecorder) DeleteBackfill(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBackfill", reflect.TypeOf((*MockFrontendServiceClient)(nil).DeleteBackfill), varargs...)
}

// DeleteTicket mocks base method.
func (m *MockFrontendServiceClient) DeleteTicket(ctx context.Context, in *pb.DeleteTicketRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTicket", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTicket indicates an expected call of DeleteTicket.
func (mr *MockFrontendServiceClientMockRecorder) DeleteTicket(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTicket", reflect.TypeOf((*MockFrontendServiceClient)(nil).DeleteTicket), varargs...)
}

// GetBackfill mocks base method.
func (m *MockFrontendServiceClient) GetBackfill(ctx context.Context, in *pb.GetBackfillRequest, opts ...grpc.CallOption) (*pb.Backfill, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBackfill", varargs...)
	ret0, _ := ret[0].(*pb.Backfill)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackfill indicates an expected call of GetBackfill.
func (mr *MockFrontendServiceClientMockRecorder) GetBackfill(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackfill", reflect.TypeOf((*MockFrontendServiceClient)(nil).GetBackfill), varargs...)
}

// GetTicket mocks base method.
func (m *MockFrontendServiceClient) GetTicket(ctx context.Context, in *pb.GetTicketRequest, opts ...grpc.CallOption) (*pb.Ticket, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTicket", varargs...)
	ret0, _ := ret[0].(*pb.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTicket indicates an expected call of GetTicket.
func (mr *MockFrontendServiceClientMockRecorder) GetTicket(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicket", reflect.TypeOf((*MockFrontendServiceClient)(nil).GetTicket), varargs...)
}

// UpdateBackfill mocks base method.
func (m *MockFrontendServiceClient) UpdateBackfill(ctx context.Context, in *pb.UpdateBackfillRequest, opts ...grpc.CallOption) (*pb.Backfill, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateBackfill", varargs...)
	ret0, _ := ret[0].(*pb.Backfill)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBackfill indicates an expected call of UpdateBackfill.
func (mr *MockFrontendServiceClientMockRecorder) UpdateBackfill(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBackfill", reflect.TypeOf((*MockFrontendServiceClient)(nil).UpdateBackfill), varargs...)
}

// WatchAssignments mocks base method.
func (m *MockFrontendServiceClient) WatchAssignments(ctx context.Context, in *pb.WatchAssignmentsRequest, opts ...grpc.CallOption) (pb.FrontendService_WatchAssignmentsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WatchAssignments", varargs...)
	ret0, _ := ret[0].(pb.FrontendService_WatchAssignmentsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchAssignments indicates an expected call of WatchAssignments.
func (mr *MockFrontendServiceClientMockRecorder) WatchAssignments(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchAssignments", reflect.TypeOf((*MockFrontendServiceClient)(nil).WatchAssignments), varargs...)
}

// MockFrontendService_WatchAssignmentsClient is a mock of FrontendService_WatchAssignmentsClient interface.
type MockFrontendService_WatchAssignmentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockFrontendService_WatchAssignmentsClientMockRecorder
}

// MockFrontendService_WatchAssignmentsClientMockRecorder is the mock recorder for MockFrontendService_WatchAssignmentsClient.
type MockFrontendService_WatchAssignmentsClientMockRecorder struct {
	mock *MockFrontendService_WatchAssignmentsClient
}

// NewMockFrontendService_WatchAssignmentsClient creates a new mock instance.
func NewMockFrontendService_WatchAssignmentsClient(ctrl *gomock.Controller) *MockFrontendService_WatchAssignmentsClient {
	mock := &MockFrontendService_WatchAssignmentsClient{ctrl: ctrl}
	mock.recorder = &MockFrontendService_WatchAssignmentsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFrontendService_WatchAssignmentsClient) EXPECT() *MockFrontendService_WatchAssignmentsClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockFrontendService_WatchAssignmentsClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockFrontendService_WatchAssignmentsClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockFrontendService_WatchAssignmentsClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockFrontendService_WatchAssignmentsClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockFrontendService_WatchAssignmentsClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockFrontendService_WatchAssignmentsClient)(nil).Context))
}

// Header mocks base method.
func (m *MockFrontendService_WatchAssignmentsClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockFrontendService_WatchAssignmentsClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockFrontendService_WatchAssignmentsClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockFrontendService_WatchAssignmentsClient) Recv() (*pb.WatchAssignmentsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*pb.WatchAssignmentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockFrontendService_WatchAssignmentsClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockFrontendService_WatchAssignmentsClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockFrontendService_WatchAssignmentsClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockFrontendService_WatchAssignmentsClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockFrontendService_WatchAssignmentsClient)(nil).RecvMsg), m)
}

// SendMsg mocks base method.
func (m_2 *MockFrontendService_WatchAssignmentsClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockFrontendService_WatchAssignmentsClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockFrontendService_WatchAssignmentsClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockFrontendService_WatchAssignmentsClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockFrontendService_WatchAssignmentsClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockFrontendService_WatchAssignmentsClient)(nil).Trailer))
}

// MockFrontendServiceServer is a mock of FrontendServiceServer interface.
type MockFrontendServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockFrontendServiceServerMockRecorder
}

// MockFrontendServiceServerMockRecorder is the mock recorder for MockFrontendServiceServer.
type MockFrontendServiceServerMockRecorder struct {
	mock *MockFrontendServiceServer
}

// NewMockFrontendServiceServer creates a new mock instance.
func NewMockFrontendServiceServer(ctrl *gomock.Controller) *MockFrontendServiceServer {
	mock := &MockFrontendServiceServer{ctrl: ctrl}
	mock.recorder = &MockFrontendServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFrontendServiceServer) EXPECT() *MockFrontendServiceServerMockRecorder {
	return m.recorder
}

// AcknowledgeBackfill mocks base method.
func (m *MockFrontendServiceServer) AcknowledgeBackfill(arg0 context.Context, arg1 *pb.AcknowledgeBackfillRequest) (*pb.AcknowledgeBackfillResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcknowledgeBackfill", arg0, arg1)
	ret0, _ := ret[0].(*pb.AcknowledgeBackfillResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcknowledgeBackfill indicates an expected call of AcknowledgeBackfill.
func (mr *MockFrontendServiceServerMockRecorder) AcknowledgeBackfill(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcknowledgeBackfill", reflect.TypeOf((*MockFrontendServiceServer)(nil).AcknowledgeBackfill), arg0, arg1)
}

// CreateBackfill mocks base method.
func (m *MockFrontendServiceServer) CreateBackfill(arg0 context.Context, arg1 *pb.CreateBackfillRequest) (*pb.Backfill, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBackfill", arg0, arg1)
	ret0, _ := ret[0].(*pb.Backfill)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBackfill indicates an expected call of CreateBackfill.
func (mr *MockFrontendServiceServerMockRecorder) CreateBackfill(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBackfill", reflect.TypeOf((*MockFrontendServiceServer)(nil).CreateBackfill), arg0, arg1)
}

// CreateTicket mocks base method.
func (m *MockFrontendServiceServer) CreateTicket(arg0 context.Context, arg1 *pb.CreateTicketRequest) (*pb.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTicket", arg0, arg1)
	ret0, _ := ret[0].(*pb.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTicket indicates an expected call of CreateTicket.
func (mr *MockFrontendServiceServerMockRecorder) CreateTicket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTicket", reflect.TypeOf((*MockFrontendServiceServer)(nil).CreateTicket), arg0, arg1)
}

// DeleteBackfill mocks base method.
func (m *MockFrontendServiceServer) DeleteBackfill(arg0 context.Context, arg1 *pb.DeleteBackfillRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBackfill", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBackfill indicates an expected call of DeleteBackfill.
func (mr *MockFrontendServiceServerMockRecorder) DeleteBackfill(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBackfill", reflect.TypeOf((*MockFrontendServiceServer)(nil).DeleteBackfill), arg0, arg1)
}

// DeleteTicket mocks base method.
func (m *MockFrontendServiceServer) DeleteTicket(arg0 context.Context, arg1 *pb.DeleteTicketRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTicket", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTicket indicates an expected call of DeleteTicket.
func (mr *MockFrontendServiceServerMockRecorder) DeleteTicket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTicket", reflect.TypeOf((*MockFrontendServiceServer)(nil).DeleteTicket), arg0, arg1)
}

// GetBackfill mocks base method.
func (m *MockFrontendServiceServer) GetBackfill(arg0 context.Context, arg1 *pb.GetBackfillRequest) (*pb.Backfill, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackfill", arg0, arg1)
	ret0, _ := ret[0].(*pb.Backfill)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackfill indicates an expected call of GetBackfill.
func (mr *MockFrontendServiceServerMockRecorder) GetBackfill(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackfill", reflect.TypeOf((*MockFrontendServiceServer)(nil).GetBackfill), arg0, arg1)
}

// GetTicket mocks base method.
func (m *MockFrontendServiceServer) GetTicket(arg0 context.Context, arg1 *pb.GetTicketRequest) (*pb.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicket", arg0, arg1)
	ret0, _ := ret[0].(*pb.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTicket indicates an expected call of GetTicket.
func (mr *MockFrontendServiceServerMockRecorder) GetTicket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicket", reflect.TypeOf((*MockFrontendServiceServer)(nil).GetTicket), arg0, arg1)
}

// UpdateBackfill mocks base method.
func (m *MockFrontendServiceServer) UpdateBackfill(arg0 context.Context, arg1 *pb.UpdateBackfillRequest) (*pb.Backfill, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBackfill", arg0, arg1)
	ret0, _ := ret[0].(*pb.Backfill)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBackfill indicates an expected call of UpdateBackfill.
func (mr *MockFrontendServiceServerMockRecorder) UpdateBackfill(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBackfill", reflect.TypeOf((*MockFrontendServiceServer)(nil).UpdateBackfill), arg0, arg1)
}

// WatchAssignments mocks base method.
func (m *MockFrontendServiceServer) WatchAssignments(arg0 *pb.WatchAssignmentsRequest, arg1 pb.FrontendService_WatchAssignmentsServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchAssignments", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WatchAssignments indicates an expected call of WatchAssignments.
func (mr *MockFrontendServiceServerMockRecorder) WatchAssignments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchAssignments", reflect.TypeOf((*MockFrontendServiceServer)(nil).WatchAssignments), arg0, arg1)
}

// MockUnsafeFrontendServiceServer is a mock of UnsafeFrontendServiceServer interface.
type MockUnsafeFrontendServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeFrontendServiceServerMockRecorder
}

// MockUnsafeFrontendServiceServerMockRecorder is the mock recorder for MockUnsafeFrontendServiceServer.
type MockUnsafeFrontendServiceServerMockRecorder struct {
	mock *MockUnsafeFrontendServiceServer
}

// NewMockUnsafeFrontendServiceServer creates a new mock instance.
func NewMockUnsafeFrontendServiceServer(ctrl *gomock.Controller) *MockUnsafeFrontendServiceServer {
	mock := &MockUnsafeFrontendServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeFrontendServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeFrontendServiceServer) EXPECT() *MockUnsafeFrontendServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedFrontendServiceServer mocks base method.
func (m *MockUnsafeFrontendServiceServer) mustEmbedUnimplementedFrontendServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedFrontendServiceServer")
}

// mustEmbedUnimplementedFrontendServiceServer indicates an expected call of mustEmbedUnimplementedFrontendServiceServer.
func (mr *MockUnsafeFrontendServiceServerMockRecorder) mustEmbedUnimplementedFrontendServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedFrontendServiceServer", reflect.TypeOf((*MockUnsafeFrontendServiceServer)(nil).mustEmbedUnimplementedFrontendServiceServer))
}

// MockFrontendService_WatchAssignmentsServer is a mock of FrontendService_WatchAssignmentsServer interface.
type MockFrontendService_WatchAssignmentsServer struct {
	ctrl     *gomock.Controller
	recorder *MockFrontendService_WatchAssignmentsServerMockRecorder
}

// MockFrontendService_WatchAssignmentsServerMockRecorder is the mock recorder for MockFrontendService_WatchAssignmentsServer.
type MockFrontendService_WatchAssignmentsServerMockRecorder struct {
	mock *MockFrontendService_WatchAssignmentsServer
}

// NewMockFrontendService_WatchAssignmentsServer creates a new mock instance.
func NewMockFrontendService_WatchAssignmentsServer(ctrl *gomock.Controller) *MockFrontendService_WatchAssignmentsServer {
	mock := &MockFrontendService_WatchAssignmentsServer{ctrl: ctrl}
	mock.recorder = &MockFrontendService_WatchAssignmentsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFrontendService_WatchAssignmentsServer) EXPECT() *MockFrontendService_WatchAssignmentsServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockFrontendService_WatchAssignmentsServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockFrontendService_WatchAssignmentsServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockFrontendService_WatchAssignmentsServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockFrontendService_WatchAssignmentsServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockFrontendService_WatchAssignmentsServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockFrontendService_WatchAssignmentsServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockFrontendService_WatchAssignmentsServer) Send(arg0 *pb.WatchAssignmentsResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockFrontendService_WatchAssignmentsServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockFrontendService_WatchAssignmentsServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockFrontendService_WatchAssignmentsServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockFrontendService_WatchAssignmentsServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockFrontendService_WatchAssignmentsServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockFrontendService_WatchAssignmentsServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockFrontendService_WatchAssignmentsServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockFrontendService_WatchAssignmentsServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockFrontendService_WatchAssignmentsServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockFrontendService_WatchAssignmentsServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockFrontendService_WatchAssignmentsServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockFrontendService_WatchAssignmentsServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockFrontendService_WatchAssignmentsServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockFrontendService_WatchAssignmentsServer)(nil).SetTrailer), arg0)
}
