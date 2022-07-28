// Copyright 2021 Google LLC
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

//  **** IMPORTANT NOTE ****
//
//  The proto of BT data has to match exactly the g3 proto, including tag
//  number.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: common.proto

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

// The direction of node property (arc in the graph).
type PropertyDirection int32

const (
	PropertyDirection_DIRECTION_UNKNOWN PropertyDirection = 0
	PropertyDirection_DIRECTION_IN      PropertyDirection = 1
	PropertyDirection_DIRECTION_OUT     PropertyDirection = 2
)

// Enum value maps for PropertyDirection.
var (
	PropertyDirection_name = map[int32]string{
		0: "DIRECTION_UNKNOWN",
		1: "DIRECTION_IN",
		2: "DIRECTION_OUT",
	}
	PropertyDirection_value = map[string]int32{
		"DIRECTION_UNKNOWN": 0,
		"DIRECTION_IN":      1,
		"DIRECTION_OUT":     2,
	}
)

func (x PropertyDirection) Enum() *PropertyDirection {
	p := new(PropertyDirection)
	*p = x
	return p
}

func (x PropertyDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PropertyDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (PropertyDirection) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x PropertyDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PropertyDirection.Descriptor instead.
func (PropertyDirection) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

// Represent a node with a subgraph attached to it.
type GraphNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A group of neighbour nodes by predicate.
	Neighbors []*GraphNode_LinkedNodes `protobuf:"bytes,1,rep,name=neighbors,proto3" json:"neighbors,omitempty"`
	// Value of the node, could be dcid or node string value.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GraphNode) Reset() {
	*x = GraphNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphNode) ProtoMessage() {}

func (x *GraphNode) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphNode.ProtoReflect.Descriptor instead.
func (*GraphNode) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *GraphNode) GetNeighbors() []*GraphNode_LinkedNodes {
	if x != nil {
		return x.Neighbors
	}
	return nil
}

func (x *GraphNode) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// This is used in SubGraph cache.
// A single root ID might match multiple SubGraph configs.
// All `nodes` have the same root ID.
type GraphNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*GraphNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *GraphNodes) Reset() {
	*x = GraphNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphNodes) ProtoMessage() {}

func (x *GraphNodes) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphNodes.ProtoReflect.Descriptor instead.
func (*GraphNodes) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *GraphNodes) GetNodes() []*GraphNode {
	if x != nil {
		return x.Nodes
	}
	return nil
}

// Holds the import manifest information.
type MemdbConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the import
	ImportName string `protobuf:"bytes,1,opt,name=import_name,json=importName,proto3" json:"import_name,omitempty"`
	// The url of the data, ususally the source url link
	ProvenanceUrl string `protobuf:"bytes,2,opt,name=provenance_url,json=provenanceUrl,proto3" json:"provenance_url,omitempty"`
	// Data download link. For private import, this should be the GCS folder url.
	// Example:
	// https://pantheon.corp.google.com/storage/browser/datcom-public/food
	DataDownloadUrl string `protobuf:"bytes,3,opt,name=data_download_url,json=dataDownloadUrl,proto3" json:"data_download_url,omitempty"`
	// The root SVG for the import. This will be set as a child stat var group
	// of "dc/g/Root"
	RootSvg string `protobuf:"bytes,4,opt,name=root_svg,json=rootSvg,proto3" json:"root_svg,omitempty"`
	// Key is StatVarGroup ID.
	StatVarGroups map[string]*StatVarGroupNode `protobuf:"bytes,5,rep,name=stat_var_groups,json=statVarGroups,proto3" json:"stat_var_groups,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MemdbConfig) Reset() {
	*x = MemdbConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemdbConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemdbConfig) ProtoMessage() {}

func (x *MemdbConfig) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemdbConfig.ProtoReflect.Descriptor instead.
func (*MemdbConfig) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *MemdbConfig) GetImportName() string {
	if x != nil {
		return x.ImportName
	}
	return ""
}

func (x *MemdbConfig) GetProvenanceUrl() string {
	if x != nil {
		return x.ProvenanceUrl
	}
	return ""
}

func (x *MemdbConfig) GetDataDownloadUrl() string {
	if x != nil {
		return x.DataDownloadUrl
	}
	return ""
}

func (x *MemdbConfig) GetRootSvg() string {
	if x != nil {
		return x.RootSvg
	}
	return ""
}

func (x *MemdbConfig) GetStatVarGroups() map[string]*StatVarGroupNode {
	if x != nil {
		return x.StatVarGroups
	}
	return nil
}

// PayloadResponse represents API response with encoded payload string
type PayloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload string `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *PayloadResponse) Reset() {
	*x = PayloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadResponse) ProtoMessage() {}

func (x *PayloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadResponse.ProtoReflect.Descriptor instead.
func (*PayloadResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *PayloadResponse) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

// Message to hold a cohort of nodes that have the same predicate and
// direction.
type GraphNode_LinkedNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Property  string            `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
	Direction PropertyDirection `protobuf:"varint,2,opt,name=direction,proto3,enum=datacommons.PropertyDirection" json:"direction,omitempty"`
	Nodes     []*GraphNode      `protobuf:"bytes,3,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *GraphNode_LinkedNodes) Reset() {
	*x = GraphNode_LinkedNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphNode_LinkedNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphNode_LinkedNodes) ProtoMessage() {}

func (x *GraphNode_LinkedNodes) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphNode_LinkedNodes.ProtoReflect.Descriptor instead.
func (*GraphNode_LinkedNodes) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GraphNode_LinkedNodes) GetProperty() string {
	if x != nil {
		return x.Property
	}
	return ""
}

func (x *GraphNode_LinkedNodes) GetDirection() PropertyDirection {
	if x != nil {
		return x.Direction
	}
	return PropertyDirection_DIRECTION_UNKNOWN
}

func (x *GraphNode_LinkedNodes) GetNodes() []*GraphNode {
	if x != nil {
		return x.Nodes
	}
	return nil
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x1a, 0x0e, 0x73, 0x74, 0x61,
	0x74, 0x5f, 0x76, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x01, 0x0a, 0x09,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x6e, 0x65, 0x69,
	0x67, 0x68, 0x62, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x52, 0x09, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x1a, 0x95, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x3c, 0x0a,
	0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x3a, 0x0a, 0x0a, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0xd2, 0x02, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x64, 0x62, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x2a, 0x0a,
	0x11, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f,
	0x74, 0x5f, 0x73, 0x76, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f,
	0x74, 0x53, 0x76, 0x67, 0x12, 0x53, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x72,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x64,
	0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74,
	0x56, 0x61, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x1a, 0x5f, 0x0a, 0x12, 0x53, 0x74, 0x61,
	0x74, 0x56, 0x61, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2b, 0x0a, 0x0f, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2a, 0x4f, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x11,
	0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x02, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_proto_goTypes = []interface{}{
	(PropertyDirection)(0),        // 0: datacommons.PropertyDirection
	(*GraphNode)(nil),             // 1: datacommons.GraphNode
	(*GraphNodes)(nil),            // 2: datacommons.GraphNodes
	(*MemdbConfig)(nil),           // 3: datacommons.MemdbConfig
	(*PayloadResponse)(nil),       // 4: datacommons.PayloadResponse
	(*GraphNode_LinkedNodes)(nil), // 5: datacommons.GraphNode.LinkedNodes
	nil,                           // 6: datacommons.MemdbConfig.StatVarGroupsEntry
	(*StatVarGroupNode)(nil),      // 7: datacommons.StatVarGroupNode
}
var file_common_proto_depIdxs = []int32{
	5, // 0: datacommons.GraphNode.neighbors:type_name -> datacommons.GraphNode.LinkedNodes
	1, // 1: datacommons.GraphNodes.nodes:type_name -> datacommons.GraphNode
	6, // 2: datacommons.MemdbConfig.stat_var_groups:type_name -> datacommons.MemdbConfig.StatVarGroupsEntry
	0, // 3: datacommons.GraphNode.LinkedNodes.direction:type_name -> datacommons.PropertyDirection
	1, // 4: datacommons.GraphNode.LinkedNodes.nodes:type_name -> datacommons.GraphNode
	7, // 5: datacommons.MemdbConfig.StatVarGroupsEntry.value:type_name -> datacommons.StatVarGroupNode
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	file_stat_var_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphNode); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphNodes); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemdbConfig); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadResponse); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphNode_LinkedNodes); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
