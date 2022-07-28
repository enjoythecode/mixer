// Copyright 2022 Google LLC
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: v1/variables.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VariablesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity string `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *VariablesRequest) Reset() {
	*x = VariablesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_variables_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariablesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariablesRequest) ProtoMessage() {}

func (x *VariablesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_variables_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariablesRequest.ProtoReflect.Descriptor instead.
func (*VariablesRequest) Descriptor() ([]byte, []int) {
	return file_v1_variables_proto_rawDescGZIP(), []int{0}
}

func (x *VariablesRequest) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

type VariablesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity    string   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Variables []string `protobuf:"bytes,2,rep,name=variables,proto3" json:"variables,omitempty"`
}

func (x *VariablesResponse) Reset() {
	*x = VariablesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_variables_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariablesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariablesResponse) ProtoMessage() {}

func (x *VariablesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_variables_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariablesResponse.ProtoReflect.Descriptor instead.
func (*VariablesResponse) Descriptor() ([]byte, []int) {
	return file_v1_variables_proto_rawDescGZIP(), []int{1}
}

func (x *VariablesResponse) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *VariablesResponse) GetVariables() []string {
	if x != nil {
		return x.Variables
	}
	return nil
}

type BulkVariablesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entities []string `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
	// If set, return the union of all variables of all entities.
	Union bool `protobuf:"varint,2,opt,name=union,proto3" json:"union,omitempty"`
}

func (x *BulkVariablesRequest) Reset() {
	*x = BulkVariablesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_variables_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkVariablesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkVariablesRequest) ProtoMessage() {}

func (x *BulkVariablesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_variables_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkVariablesRequest.ProtoReflect.Descriptor instead.
func (*BulkVariablesRequest) Descriptor() ([]byte, []int) {
	return file_v1_variables_proto_rawDescGZIP(), []int{2}
}

func (x *BulkVariablesRequest) GetEntities() []string {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *BulkVariablesRequest) GetUnion() bool {
	if x != nil {
		return x.Union
	}
	return false
}

type BulkVariablesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*VariablesResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *BulkVariablesResponse) Reset() {
	*x = BulkVariablesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_variables_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkVariablesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkVariablesResponse) ProtoMessage() {}

func (x *BulkVariablesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_variables_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkVariablesResponse.ProtoReflect.Descriptor instead.
func (*BulkVariablesResponse) Descriptor() ([]byte, []int) {
	return file_v1_variables_proto_rawDescGZIP(), []int{3}
}

func (x *BulkVariablesResponse) GetData() []*VariablesResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_v1_variables_proto protoreflect.FileDescriptor

var file_v1_variables_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x22, 0x2a, 0x0a, 0x10, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x22, 0x49, 0x0a, 0x11, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x22, 0x48, 0x0a, 0x14, 0x42,
	0x75, 0x6c, 0x6b, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x22, 0x4e, 0x0a, 0x15, 0x42, 0x75, 0x6c, 0x6b, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_variables_proto_rawDescOnce sync.Once
	file_v1_variables_proto_rawDescData = file_v1_variables_proto_rawDesc
)

func file_v1_variables_proto_rawDescGZIP() []byte {
	file_v1_variables_proto_rawDescOnce.Do(func() {
		file_v1_variables_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_variables_proto_rawDescData)
	})
	return file_v1_variables_proto_rawDescData
}

var file_v1_variables_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_variables_proto_goTypes = []interface{}{
	(*VariablesRequest)(nil),      // 0: datacommons.v1.VariablesRequest
	(*VariablesResponse)(nil),     // 1: datacommons.v1.VariablesResponse
	(*BulkVariablesRequest)(nil),  // 2: datacommons.v1.BulkVariablesRequest
	(*BulkVariablesResponse)(nil), // 3: datacommons.v1.BulkVariablesResponse
}
var file_v1_variables_proto_depIdxs = []int32{
	1, // 0: datacommons.v1.BulkVariablesResponse.data:type_name -> datacommons.v1.VariablesResponse
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_variables_proto_init() }
func file_v1_variables_proto_init() {
	if File_v1_variables_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_variables_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VariablesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_variables_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VariablesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_variables_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkVariablesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_variables_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkVariablesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_variables_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_variables_proto_goTypes,
		DependencyIndexes: file_v1_variables_proto_depIdxs,
		MessageInfos:      file_v1_variables_proto_msgTypes,
	}.Build()
	File_v1_variables_proto = out.File
	file_v1_variables_proto_rawDesc = nil
	file_v1_variables_proto_goTypes = nil
	file_v1_variables_proto_depIdxs = nil
}
