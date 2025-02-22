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

// REST API URL from the proto in this file:
// ========================================
//    /node/property-labels
//    /node/property-values
//    /node/triples
// ========================================

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: node.proto

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

// Triple.
type Triple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubjectId    string   `protobuf:"bytes,1,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	Predicate    string   `protobuf:"bytes,2,opt,name=predicate,proto3" json:"predicate,omitempty"`
	ObjectId     string   `protobuf:"bytes,3,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	ObjectName   string   `protobuf:"bytes,4,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"` // Only when object_id is set.
	ObjectValue  string   `protobuf:"bytes,5,opt,name=object_value,json=objectValue,proto3" json:"object_value,omitempty"`
	ProvenanceId string   `protobuf:"bytes,6,opt,name=provenance_id,json=provenanceId,proto3" json:"provenance_id,omitempty"`
	SubjectName  string   `protobuf:"bytes,7,opt,name=subject_name,json=subjectName,proto3" json:"subject_name,omitempty"`    // Only for in arcs.
	SubjectTypes []string `protobuf:"bytes,8,rep,name=subject_types,json=subjectTypes,proto3" json:"subject_types,omitempty"` // Only for in arcs.
	ObjectTypes  []string `protobuf:"bytes,10,rep,name=object_types,json=objectTypes,proto3" json:"object_types,omitempty"`   // Only when object_id is set.
}

func (x *Triple) Reset() {
	*x = Triple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Triple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Triple) ProtoMessage() {}

func (x *Triple) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Triple.ProtoReflect.Descriptor instead.
func (*Triple) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{0}
}

func (x *Triple) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *Triple) GetPredicate() string {
	if x != nil {
		return x.Predicate
	}
	return ""
}

func (x *Triple) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *Triple) GetObjectName() string {
	if x != nil {
		return x.ObjectName
	}
	return ""
}

func (x *Triple) GetObjectValue() string {
	if x != nil {
		return x.ObjectValue
	}
	return ""
}

func (x *Triple) GetProvenanceId() string {
	if x != nil {
		return x.ProvenanceId
	}
	return ""
}

func (x *Triple) GetSubjectName() string {
	if x != nil {
		return x.SubjectName
	}
	return ""
}

func (x *Triple) GetSubjectTypes() []string {
	if x != nil {
		return x.SubjectTypes
	}
	return nil
}

func (x *Triple) GetObjectTypes() []string {
	if x != nil {
		return x.ObjectTypes
	}
	return nil
}

// A collection of triples.
type Triples struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Triples []*Triple `protobuf:"bytes,1,rep,name=triples,proto3" json:"triples,omitempty"`
	// Before potential truncation.
	Total float64 `protobuf:"fixed64,6,opt,name=total,proto3" json:"total,omitempty"`
	// A predicate could link to multiple neighbors.
	InNodes   map[string]*EntityInfoCollection `protobuf:"bytes,3,rep,name=in_nodes,json=inNodes,proto3" json:"in_nodes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`    // Key is predicate.
	OutNodes  map[string]*EntityInfoCollection `protobuf:"bytes,4,rep,name=out_nodes,json=outNodes,proto3" json:"out_nodes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Key is predicate.
	SubjectId string                           `protobuf:"bytes,5,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
}

func (x *Triples) Reset() {
	*x = Triples{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Triples) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Triples) ProtoMessage() {}

func (x *Triples) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Triples.ProtoReflect.Descriptor instead.
func (*Triples) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{1}
}

func (x *Triples) GetTriples() []*Triple {
	if x != nil {
		return x.Triples
	}
	return nil
}

func (x *Triples) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Triples) GetInNodes() map[string]*EntityInfoCollection {
	if x != nil {
		return x.InNodes
	}
	return nil
}

func (x *Triples) GetOutNodes() map[string]*EntityInfoCollection {
	if x != nil {
		return x.OutNodes
	}
	return nil
}

func (x *Triples) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

// Full lists of in and out labels for a given node ID.
type PropertyLabels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InLabels  []string `protobuf:"bytes,1,rep,name=in_labels,json=inLabels,proto3" json:"in_labels,omitempty"`
	OutLabels []string `protobuf:"bytes,2,rep,name=out_labels,json=outLabels,proto3" json:"out_labels,omitempty"`
}

func (x *PropertyLabels) Reset() {
	*x = PropertyLabels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropertyLabels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropertyLabels) ProtoMessage() {}

func (x *PropertyLabels) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropertyLabels.ProtoReflect.Descriptor instead.
func (*PropertyLabels) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{2}
}

func (x *PropertyLabels) GetInLabels() []string {
	if x != nil {
		return x.InLabels
	}
	return nil
}

func (x *PropertyLabels) GetOutLabels() []string {
	if x != nil {
		return x.OutLabels
	}
	return nil
}

// Request to get property labels.
type GetPropertyLabelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The dcids of nodes.
	Dcids []string `protobuf:"bytes,1,rep,name=dcids,proto3" json:"dcids,omitempty"`
}

func (x *GetPropertyLabelsRequest) Reset() {
	*x = GetPropertyLabelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPropertyLabelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPropertyLabelsRequest) ProtoMessage() {}

func (x *GetPropertyLabelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPropertyLabelsRequest.ProtoReflect.Descriptor instead.
func (*GetPropertyLabelsRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{3}
}

func (x *GetPropertyLabelsRequest) GetDcids() []string {
	if x != nil {
		return x.Dcids
	}
	return nil
}

// Response returned by GetPropertyLabels.
type GetPropertyLabelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[string]*PropertyLabels `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetPropertyLabelsResponse) Reset() {
	*x = GetPropertyLabelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPropertyLabelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPropertyLabelsResponse) ProtoMessage() {}

func (x *GetPropertyLabelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPropertyLabelsResponse.ProtoReflect.Descriptor instead.
func (*GetPropertyLabelsResponse) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{4}
}

func (x *GetPropertyLabelsResponse) GetData() map[string]*PropertyLabels {
	if x != nil {
		return x.Data
	}
	return nil
}

// Request to get all neighboring nodes along an edge labeled by a property.
type GetPropertyValuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The dcids of nodes to query for.
	Dcids []string `protobuf:"bytes,1,rep,name=dcids,proto3" json:"dcids,omitempty"`
	// The type of the neighboring node to query for.
	ValueType string `protobuf:"bytes,2,opt,name=value_type,json=valueType,proto3" json:"value_type,omitempty"`
	// The property to get adjacent nodes for.
	Property string `protobuf:"bytes,3,opt,name=property,proto3" json:"property,omitempty"`
	// Maximum number of nodes to query for. When set to 0 or non-specified,
	// no limit is applied.
	Limit int32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// Direction, "in" or "out". When not specified, return results for both.
	Direction string `protobuf:"bytes,5,opt,name=direction,proto3" json:"direction,omitempty"`
}

func (x *GetPropertyValuesRequest) Reset() {
	*x = GetPropertyValuesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPropertyValuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPropertyValuesRequest) ProtoMessage() {}

func (x *GetPropertyValuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPropertyValuesRequest.ProtoReflect.Descriptor instead.
func (*GetPropertyValuesRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{5}
}

func (x *GetPropertyValuesRequest) GetDcids() []string {
	if x != nil {
		return x.Dcids
	}
	return nil
}

func (x *GetPropertyValuesRequest) GetValueType() string {
	if x != nil {
		return x.ValueType
	}
	return ""
}

func (x *GetPropertyValuesRequest) GetProperty() string {
	if x != nil {
		return x.Property
	}
	return ""
}

func (x *GetPropertyValuesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetPropertyValuesRequest) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

// ArcNodes represents property and linked nodes.
type ArcNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key is property
	In  []*EntityInfo `protobuf:"bytes,1,rep,name=in,proto3" json:"in,omitempty"`
	Out []*EntityInfo `protobuf:"bytes,2,rep,name=out,proto3" json:"out,omitempty"`
}

func (x *ArcNodes) Reset() {
	*x = ArcNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArcNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArcNodes) ProtoMessage() {}

func (x *ArcNodes) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArcNodes.ProtoReflect.Descriptor instead.
func (*ArcNodes) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{6}
}

func (x *ArcNodes) GetIn() []*EntityInfo {
	if x != nil {
		return x.In
	}
	return nil
}

func (x *ArcNodes) GetOut() []*EntityInfo {
	if x != nil {
		return x.Out
	}
	return nil
}

// Response returned by GetPropertyValues.
type GetPropertyValuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key is dcid.
	Data map[string]*ArcNodes `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetPropertyValuesResponse) Reset() {
	*x = GetPropertyValuesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPropertyValuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPropertyValuesResponse) ProtoMessage() {}

func (x *GetPropertyValuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPropertyValuesResponse.ProtoReflect.Descriptor instead.
func (*GetPropertyValuesResponse) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{7}
}

func (x *GetPropertyValuesResponse) GetData() map[string]*ArcNodes {
	if x != nil {
		return x.Data
	}
	return nil
}

// Request to get all triples linking to the given nodes.
type GetTriplesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The dcids of the nodes to query for.
	Dcids []string `protobuf:"bytes,1,rep,name=dcids,proto3" json:"dcids,omitempty"`
	// Maximum number of triples for each property and type of the neighbor.
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetTriplesRequest) Reset() {
	*x = GetTriplesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTriplesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTriplesRequest) ProtoMessage() {}

func (x *GetTriplesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTriplesRequest.ProtoReflect.Descriptor instead.
func (*GetTriplesRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{8}
}

func (x *GetTriplesRequest) GetDcids() []string {
	if x != nil {
		return x.Dcids
	}
	return nil
}

func (x *GetTriplesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

// Response returned by GetTriples.
type GetTriplesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key is entity dcid.
	Triples map[string]*Triples `protobuf:"bytes,1,rep,name=triples,proto3" json:"triples,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetTriplesResponse) Reset() {
	*x = GetTriplesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTriplesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTriplesResponse) ProtoMessage() {}

func (x *GetTriplesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTriplesResponse.ProtoReflect.Descriptor instead.
func (*GetTriplesResponse) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{9}
}

func (x *GetTriplesResponse) GetTriples() map[string]*Triples {
	if x != nil {
		return x.Triples
	}
	return nil
}

var File_node_proto protoreflect.FileDescriptor

var file_node_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x1a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x02, 0x0a, 0x06, 0x54, 0x72, 0x69, 0x70,
	0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x22, 0xb1, 0x03, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x07,
	0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x72, 0x69, 0x70,
	0x6c, 0x65, 0x52, 0x07, 0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x3c, 0x0a, 0x08, 0x69, 0x6e, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x69, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12,
	0x3f, 0x0a, 0x09, 0x6f, 0x75, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6f, 0x75, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x1a,
	0x5d, 0x0a, 0x0c, 0x49, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x37, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5e,
	0x0a, 0x0d, 0x4f, 0x75, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x37, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04,
	0x08, 0x02, 0x10, 0x03, 0x22, 0x4c, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x5f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x75, 0x74, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x22, 0x30, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x64, 0x63, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x64,
	0x63, 0x69, 0x64, 0x73, 0x22, 0xb7, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x44, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x54, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9f,
	0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x64,
	0x63, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x64, 0x63, 0x69, 0x64,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x5e, 0x0a, 0x08, 0x41, 0x72, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x02,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x02, 0x69, 0x6e, 0x12, 0x29, 0x0a, 0x03, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x6f, 0x75, 0x74,
	0x22, 0xb1, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x4e, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x41, 0x72, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x63, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x64, 0x63, 0x69, 0x64, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xae, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69,
	0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x07,
	0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54,
	0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x74, 0x72, 0x69,
	0x70, 0x6c, 0x65, 0x73, 0x1a, 0x50, 0x0a, 0x0c, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_proto_rawDescOnce sync.Once
	file_node_proto_rawDescData = file_node_proto_rawDesc
)

func file_node_proto_rawDescGZIP() []byte {
	file_node_proto_rawDescOnce.Do(func() {
		file_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_proto_rawDescData)
	})
	return file_node_proto_rawDescData
}

var file_node_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_node_proto_goTypes = []interface{}{
	(*Triple)(nil),                    // 0: datacommons.Triple
	(*Triples)(nil),                   // 1: datacommons.Triples
	(*PropertyLabels)(nil),            // 2: datacommons.PropertyLabels
	(*GetPropertyLabelsRequest)(nil),  // 3: datacommons.GetPropertyLabelsRequest
	(*GetPropertyLabelsResponse)(nil), // 4: datacommons.GetPropertyLabelsResponse
	(*GetPropertyValuesRequest)(nil),  // 5: datacommons.GetPropertyValuesRequest
	(*ArcNodes)(nil),                  // 6: datacommons.ArcNodes
	(*GetPropertyValuesResponse)(nil), // 7: datacommons.GetPropertyValuesResponse
	(*GetTriplesRequest)(nil),         // 8: datacommons.GetTriplesRequest
	(*GetTriplesResponse)(nil),        // 9: datacommons.GetTriplesResponse
	nil,                               // 10: datacommons.Triples.InNodesEntry
	nil,                               // 11: datacommons.Triples.OutNodesEntry
	nil,                               // 12: datacommons.GetPropertyLabelsResponse.DataEntry
	nil,                               // 13: datacommons.GetPropertyValuesResponse.DataEntry
	nil,                               // 14: datacommons.GetTriplesResponse.TriplesEntry
	(*EntityInfo)(nil),                // 15: datacommons.EntityInfo
	(*EntityInfoCollection)(nil),      // 16: datacommons.EntityInfoCollection
}
var file_node_proto_depIdxs = []int32{
	0,  // 0: datacommons.Triples.triples:type_name -> datacommons.Triple
	10, // 1: datacommons.Triples.in_nodes:type_name -> datacommons.Triples.InNodesEntry
	11, // 2: datacommons.Triples.out_nodes:type_name -> datacommons.Triples.OutNodesEntry
	12, // 3: datacommons.GetPropertyLabelsResponse.data:type_name -> datacommons.GetPropertyLabelsResponse.DataEntry
	15, // 4: datacommons.ArcNodes.in:type_name -> datacommons.EntityInfo
	15, // 5: datacommons.ArcNodes.out:type_name -> datacommons.EntityInfo
	13, // 6: datacommons.GetPropertyValuesResponse.data:type_name -> datacommons.GetPropertyValuesResponse.DataEntry
	14, // 7: datacommons.GetTriplesResponse.triples:type_name -> datacommons.GetTriplesResponse.TriplesEntry
	16, // 8: datacommons.Triples.InNodesEntry.value:type_name -> datacommons.EntityInfoCollection
	16, // 9: datacommons.Triples.OutNodesEntry.value:type_name -> datacommons.EntityInfoCollection
	2,  // 10: datacommons.GetPropertyLabelsResponse.DataEntry.value:type_name -> datacommons.PropertyLabels
	6,  // 11: datacommons.GetPropertyValuesResponse.DataEntry.value:type_name -> datacommons.ArcNodes
	1,  // 12: datacommons.GetTriplesResponse.TriplesEntry.value:type_name -> datacommons.Triples
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_node_proto_init() }
func file_node_proto_init() {
	if File_node_proto != nil {
		return
	}
	file_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Triple); i {
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
		file_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Triples); i {
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
		file_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertyLabels); i {
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
		file_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPropertyLabelsRequest); i {
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
		file_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPropertyLabelsResponse); i {
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
		file_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPropertyValuesRequest); i {
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
		file_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArcNodes); i {
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
		file_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPropertyValuesResponse); i {
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
		file_node_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTriplesRequest); i {
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
		file_node_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTriplesResponse); i {
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
			RawDescriptor: file_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_node_proto_goTypes,
		DependencyIndexes: file_node_proto_depIdxs,
		MessageInfos:      file_node_proto_msgTypes,
	}.Build()
	File_node_proto = out.File
	file_node_proto_rawDesc = nil
	file_node_proto_goTypes = nil
	file_node_proto_depIdxs = nil
}
