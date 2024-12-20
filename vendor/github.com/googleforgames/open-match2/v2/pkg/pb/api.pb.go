// Copyright 2024 Google LLC
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
//
// If you're modifying this file, please follow the protobuf style guide:
//   https://protobuf.dev/programming-guides/style/
// and also the Google API design guide
//   https://cloud.google.com/apis/design/
// also see the comments in the http grpc source file:
//   https://github.com/googleapis/googleapis/blob/master/google/api/http.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: api.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6f, 0x70, 0x65,
	0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x32, 0x1a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xdf, 0x04, 0x0a, 0x10, 0x4f, 0x70, 0x65,
	0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x20, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x32, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x11, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x25, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x32, 0x2e, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x32, 0x2e, 0x44,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x0f, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x23, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x32, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x32,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x1a, 0x49, 0x6e,
	0x76, 0x6f, 0x6b, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x5f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x32, 0x2e, 0x4d, 0x6d, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x32, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x4d, 0x6d, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x64, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a,
	0x10, 0x57, 0x61, 0x74, 0x63, 0x68, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x24, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x32, 0x2e,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x32, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x3b, 0x5a, 0x2c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x66,
	0x6f, 0x72, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x2d, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0xaa, 0x02, 0x0a, 0x4f, 0x70, 0x65,
	0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_proto_goTypes = []any{
	(*CreateTicketRequest)(nil),              // 0: open_match2.CreateTicketRequest
	(*DeactivateTicketsRequest)(nil),         // 1: open_match2.DeactivateTicketsRequest
	(*ActivateTicketsRequest)(nil),           // 2: open_match2.ActivateTicketsRequest
	(*MmfRequest)(nil),                       // 3: open_match2.MmfRequest
	(*CreateAssignmentsRequest)(nil),         // 4: open_match2.CreateAssignmentsRequest
	(*WatchAssignmentsRequest)(nil),          // 5: open_match2.WatchAssignmentsRequest
	(*CreateTicketResponse)(nil),             // 6: open_match2.CreateTicketResponse
	(*DeactivateTicketsResponse)(nil),        // 7: open_match2.DeactivateTicketsResponse
	(*ActivateTicketsResponse)(nil),          // 8: open_match2.ActivateTicketsResponse
	(*StreamedMmfResponse)(nil),              // 9: open_match2.StreamedMmfResponse
	(*CreateAssignmentsResponse)(nil),        // 10: open_match2.CreateAssignmentsResponse
	(*StreamedWatchAssignmentsResponse)(nil), // 11: open_match2.StreamedWatchAssignmentsResponse
}
var file_api_proto_depIdxs = []int32{
	0,  // 0: open_match2.OpenMatchService.CreateTicket:input_type -> open_match2.CreateTicketRequest
	1,  // 1: open_match2.OpenMatchService.DeactivateTickets:input_type -> open_match2.DeactivateTicketsRequest
	2,  // 2: open_match2.OpenMatchService.ActivateTickets:input_type -> open_match2.ActivateTicketsRequest
	3,  // 3: open_match2.OpenMatchService.InvokeMatchmakingFunctions:input_type -> open_match2.MmfRequest
	4,  // 4: open_match2.OpenMatchService.CreateAssignments:input_type -> open_match2.CreateAssignmentsRequest
	5,  // 5: open_match2.OpenMatchService.WatchAssignments:input_type -> open_match2.WatchAssignmentsRequest
	6,  // 6: open_match2.OpenMatchService.CreateTicket:output_type -> open_match2.CreateTicketResponse
	7,  // 7: open_match2.OpenMatchService.DeactivateTickets:output_type -> open_match2.DeactivateTicketsResponse
	8,  // 8: open_match2.OpenMatchService.ActivateTickets:output_type -> open_match2.ActivateTicketsResponse
	9,  // 9: open_match2.OpenMatchService.InvokeMatchmakingFunctions:output_type -> open_match2.StreamedMmfResponse
	10, // 10: open_match2.OpenMatchService.CreateAssignments:output_type -> open_match2.CreateAssignmentsResponse
	11, // 11: open_match2.OpenMatchService.WatchAssignments:output_type -> open_match2.StreamedWatchAssignmentsResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	file_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
