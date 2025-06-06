//
//Copyright 2021 The Dapr Authors
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//http://www.apache.org/licenses/LICENSE-2.0
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.29.3
// source: dapr/proto/internals/v1/service_invocation.proto

package internals

import (
	v1 "github.com/dapr/dapr/pkg/proto/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Actor represents actor using actor_type and actor_id
type Actor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The type of actor.
	ActorType string `protobuf:"bytes,1,opt,name=actor_type,json=actorType,proto3" json:"actor_type,omitempty"`
	// Required. The ID of actor type (actor_type)
	ActorId string `protobuf:"bytes,2,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"`
}

func (x *Actor) Reset() {
	*x = Actor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Actor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Actor) ProtoMessage() {}

func (x *Actor) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Actor.ProtoReflect.Descriptor instead.
func (*Actor) Descriptor() ([]byte, []int) {
	return file_dapr_proto_internals_v1_service_invocation_proto_rawDescGZIP(), []int{0}
}

func (x *Actor) GetActorType() string {
	if x != nil {
		return x.ActorType
	}
	return ""
}

func (x *Actor) GetActorId() string {
	if x != nil {
		return x.ActorId
	}
	return ""
}

// InternalInvokeRequest is the message to transfer caller's data to callee
// for service invocation. This includes callee's app id and caller's request data.
type InternalInvokeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The version of Dapr runtime API.
	Ver APIVersion `protobuf:"varint,1,opt,name=ver,proto3,enum=dapr.proto.internals.v1.APIVersion" json:"ver,omitempty"`
	// Required. metadata holds caller's HTTP headers or gRPC metadata.
	Metadata map[string]*ListStringValue `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Required. message including caller's invocation request.
	Message *v1.InvokeRequest `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Actor type and id. This field is used only for
	// actor service invocation.
	Actor *Actor `protobuf:"bytes,4,opt,name=actor,proto3" json:"actor,omitempty"`
}

func (x *InternalInvokeRequest) Reset() {
	*x = InternalInvokeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalInvokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalInvokeRequest) ProtoMessage() {}

func (x *InternalInvokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalInvokeRequest.ProtoReflect.Descriptor instead.
func (*InternalInvokeRequest) Descriptor() ([]byte, []int) {
	return file_dapr_proto_internals_v1_service_invocation_proto_rawDescGZIP(), []int{1}
}

func (x *InternalInvokeRequest) GetVer() APIVersion {
	if x != nil {
		return x.Ver
	}
	return APIVersion_APIVERSION_UNSPECIFIED
}

func (x *InternalInvokeRequest) GetMetadata() map[string]*ListStringValue {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *InternalInvokeRequest) GetMessage() *v1.InvokeRequest {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *InternalInvokeRequest) GetActor() *Actor {
	if x != nil {
		return x.Actor
	}
	return nil
}

// InternalInvokeResponse is the message to transfer callee's response to caller
// for service invocation.
type InternalInvokeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. HTTP/gRPC status.
	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// Required. The app callback response headers.
	Headers map[string]*ListStringValue `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// App callback response trailers.
	// This will be used only for gRPC app callback
	Trailers map[string]*ListStringValue `protobuf:"bytes,3,rep,name=trailers,proto3" json:"trailers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Callee's invocation response message.
	Message *v1.InvokeResponse `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *InternalInvokeResponse) Reset() {
	*x = InternalInvokeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalInvokeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalInvokeResponse) ProtoMessage() {}

func (x *InternalInvokeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalInvokeResponse.ProtoReflect.Descriptor instead.
func (*InternalInvokeResponse) Descriptor() ([]byte, []int) {
	return file_dapr_proto_internals_v1_service_invocation_proto_rawDescGZIP(), []int{2}
}

func (x *InternalInvokeResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *InternalInvokeResponse) GetHeaders() map[string]*ListStringValue {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *InternalInvokeResponse) GetTrailers() map[string]*ListStringValue {
	if x != nil {
		return x.Trailers
	}
	return nil
}

func (x *InternalInvokeResponse) GetMessage() *v1.InvokeResponse {
	if x != nil {
		return x.Message
	}
	return nil
}

// InternalInvokeRequestStream is a variant of InternalInvokeRequest used in streaming RPCs.
type InternalInvokeRequestStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Request details.
	// This does not contain any data in message.data.
	Request *InternalInvokeRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	// Chunk of data.
	Payload *v1.StreamPayload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *InternalInvokeRequestStream) Reset() {
	*x = InternalInvokeRequestStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalInvokeRequestStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalInvokeRequestStream) ProtoMessage() {}

func (x *InternalInvokeRequestStream) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalInvokeRequestStream.ProtoReflect.Descriptor instead.
func (*InternalInvokeRequestStream) Descriptor() ([]byte, []int) {
	return file_dapr_proto_internals_v1_service_invocation_proto_rawDescGZIP(), []int{3}
}

func (x *InternalInvokeRequestStream) GetRequest() *InternalInvokeRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *InternalInvokeRequestStream) GetPayload() *v1.StreamPayload {
	if x != nil {
		return x.Payload
	}
	return nil
}

// InternalInvokeResponseStream is a variant of InternalInvokeResponse used in streaming RPCs.
type InternalInvokeResponseStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Response details.
	// This does not contain any data in message.data.
	Response *InternalInvokeResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	// Chunk of data.
	Payload *v1.StreamPayload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *InternalInvokeResponseStream) Reset() {
	*x = InternalInvokeResponseStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalInvokeResponseStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalInvokeResponseStream) ProtoMessage() {}

func (x *InternalInvokeResponseStream) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalInvokeResponseStream.ProtoReflect.Descriptor instead.
func (*InternalInvokeResponseStream) Descriptor() ([]byte, []int) {
	return file_dapr_proto_internals_v1_service_invocation_proto_rawDescGZIP(), []int{4}
}

func (x *InternalInvokeResponseStream) GetResponse() *InternalInvokeResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *InternalInvokeResponseStream) GetPayload() *v1.StreamPayload {
	if x != nil {
		return x.Payload
	}
	return nil
}

// ListStringValue represents string value array
type ListStringValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The array of string.
	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ListStringValue) Reset() {
	*x = ListStringValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStringValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStringValue) ProtoMessage() {}

func (x *ListStringValue) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStringValue.ProtoReflect.Descriptor instead.
func (*ListStringValue) Descriptor() ([]byte, []int) {
	return file_dapr_proto_internals_v1_service_invocation_proto_rawDescGZIP(), []int{5}
}

func (x *ListStringValue) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_dapr_proto_internals_v1_service_invocation_proto protoreflect.FileDescriptor

var file_dapr_proto_internals_v1_service_invocation_proto_rawDesc = []byte{
	0x0a, 0x30, 0x64, 0x61, 0x70, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x64, 0x61, 0x70,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27,
	0x64, 0x61, 0x70, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x64, 0x61, 0x70, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x24, 0x64, 0x61, 0x70, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x05, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x84, 0x03, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x35, 0x0a, 0x03, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23,
	0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x03, 0x76, 0x65, 0x72, 0x12, 0x58, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x64, 0x61, 0x70,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x76,
	0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x3d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x6b,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x34, 0x0a, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x65, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3e, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x61, 0x70, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x91,
	0x04, 0x0a, 0x16, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x6b,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x61, 0x70, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x56, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x59, 0x0a, 0x08, 0x74, 0x72,
	0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x64,
	0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49,
	0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x72,
	0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x74, 0x72, 0x61,
	0x69, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x3e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x64, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x65, 0x0a, 0x0d, 0x54,
	0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3e,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xa6, 0x01, 0x0a, 0x1b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49,
	0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x48, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xaa, 0x01, 0x0a, 0x1c,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x4b, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x61, 0x70,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x29, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x32, 0xc4, 0x04, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6e, 0x0a, 0x09, 0x43, 0x61, 0x6c,
	0x6c, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x2e, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x09, 0x43, 0x61, 0x6c,
	0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x2e, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x11, 0x43, 0x61, 0x6c,
	0x6c, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x21,
	0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65,
	0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x84, 0x01, 0x0a, 0x0f,
	0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x34, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x35, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x12, 0x76, 0x0a, 0x0f, 0x43, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2e, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x70, 0x72, 0x2f, 0x64, 0x61,
	0x70, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dapr_proto_internals_v1_service_invocation_proto_rawDescOnce sync.Once
	file_dapr_proto_internals_v1_service_invocation_proto_rawDescData = file_dapr_proto_internals_v1_service_invocation_proto_rawDesc
)

func file_dapr_proto_internals_v1_service_invocation_proto_rawDescGZIP() []byte {
	file_dapr_proto_internals_v1_service_invocation_proto_rawDescOnce.Do(func() {
		file_dapr_proto_internals_v1_service_invocation_proto_rawDescData = protoimpl.X.CompressGZIP(file_dapr_proto_internals_v1_service_invocation_proto_rawDescData)
	})
	return file_dapr_proto_internals_v1_service_invocation_proto_rawDescData
}

var file_dapr_proto_internals_v1_service_invocation_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_dapr_proto_internals_v1_service_invocation_proto_goTypes = []interface{}{
	(*Actor)(nil),                        // 0: dapr.proto.internals.v1.Actor
	(*InternalInvokeRequest)(nil),        // 1: dapr.proto.internals.v1.InternalInvokeRequest
	(*InternalInvokeResponse)(nil),       // 2: dapr.proto.internals.v1.InternalInvokeResponse
	(*InternalInvokeRequestStream)(nil),  // 3: dapr.proto.internals.v1.InternalInvokeRequestStream
	(*InternalInvokeResponseStream)(nil), // 4: dapr.proto.internals.v1.InternalInvokeResponseStream
	(*ListStringValue)(nil),              // 5: dapr.proto.internals.v1.ListStringValue
	nil,                                  // 6: dapr.proto.internals.v1.InternalInvokeRequest.MetadataEntry
	nil,                                  // 7: dapr.proto.internals.v1.InternalInvokeResponse.HeadersEntry
	nil,                                  // 8: dapr.proto.internals.v1.InternalInvokeResponse.TrailersEntry
	(APIVersion)(0),                      // 9: dapr.proto.internals.v1.APIVersion
	(*v1.InvokeRequest)(nil),             // 10: dapr.proto.common.v1.InvokeRequest
	(*Status)(nil),                       // 11: dapr.proto.internals.v1.Status
	(*v1.InvokeResponse)(nil),            // 12: dapr.proto.common.v1.InvokeResponse
	(*v1.StreamPayload)(nil),             // 13: dapr.proto.common.v1.StreamPayload
	(*Reminder)(nil),                     // 14: dapr.proto.internals.v1.Reminder
	(*emptypb.Empty)(nil),                // 15: google.protobuf.Empty
}
var file_dapr_proto_internals_v1_service_invocation_proto_depIdxs = []int32{
	9,  // 0: dapr.proto.internals.v1.InternalInvokeRequest.ver:type_name -> dapr.proto.internals.v1.APIVersion
	6,  // 1: dapr.proto.internals.v1.InternalInvokeRequest.metadata:type_name -> dapr.proto.internals.v1.InternalInvokeRequest.MetadataEntry
	10, // 2: dapr.proto.internals.v1.InternalInvokeRequest.message:type_name -> dapr.proto.common.v1.InvokeRequest
	0,  // 3: dapr.proto.internals.v1.InternalInvokeRequest.actor:type_name -> dapr.proto.internals.v1.Actor
	11, // 4: dapr.proto.internals.v1.InternalInvokeResponse.status:type_name -> dapr.proto.internals.v1.Status
	7,  // 5: dapr.proto.internals.v1.InternalInvokeResponse.headers:type_name -> dapr.proto.internals.v1.InternalInvokeResponse.HeadersEntry
	8,  // 6: dapr.proto.internals.v1.InternalInvokeResponse.trailers:type_name -> dapr.proto.internals.v1.InternalInvokeResponse.TrailersEntry
	12, // 7: dapr.proto.internals.v1.InternalInvokeResponse.message:type_name -> dapr.proto.common.v1.InvokeResponse
	1,  // 8: dapr.proto.internals.v1.InternalInvokeRequestStream.request:type_name -> dapr.proto.internals.v1.InternalInvokeRequest
	13, // 9: dapr.proto.internals.v1.InternalInvokeRequestStream.payload:type_name -> dapr.proto.common.v1.StreamPayload
	2,  // 10: dapr.proto.internals.v1.InternalInvokeResponseStream.response:type_name -> dapr.proto.internals.v1.InternalInvokeResponse
	13, // 11: dapr.proto.internals.v1.InternalInvokeResponseStream.payload:type_name -> dapr.proto.common.v1.StreamPayload
	5,  // 12: dapr.proto.internals.v1.InternalInvokeRequest.MetadataEntry.value:type_name -> dapr.proto.internals.v1.ListStringValue
	5,  // 13: dapr.proto.internals.v1.InternalInvokeResponse.HeadersEntry.value:type_name -> dapr.proto.internals.v1.ListStringValue
	5,  // 14: dapr.proto.internals.v1.InternalInvokeResponse.TrailersEntry.value:type_name -> dapr.proto.internals.v1.ListStringValue
	1,  // 15: dapr.proto.internals.v1.ServiceInvocation.CallActor:input_type -> dapr.proto.internals.v1.InternalInvokeRequest
	1,  // 16: dapr.proto.internals.v1.ServiceInvocation.CallLocal:input_type -> dapr.proto.internals.v1.InternalInvokeRequest
	14, // 17: dapr.proto.internals.v1.ServiceInvocation.CallActorReminder:input_type -> dapr.proto.internals.v1.Reminder
	3,  // 18: dapr.proto.internals.v1.ServiceInvocation.CallLocalStream:input_type -> dapr.proto.internals.v1.InternalInvokeRequestStream
	1,  // 19: dapr.proto.internals.v1.ServiceInvocation.CallActorStream:input_type -> dapr.proto.internals.v1.InternalInvokeRequest
	2,  // 20: dapr.proto.internals.v1.ServiceInvocation.CallActor:output_type -> dapr.proto.internals.v1.InternalInvokeResponse
	2,  // 21: dapr.proto.internals.v1.ServiceInvocation.CallLocal:output_type -> dapr.proto.internals.v1.InternalInvokeResponse
	15, // 22: dapr.proto.internals.v1.ServiceInvocation.CallActorReminder:output_type -> google.protobuf.Empty
	4,  // 23: dapr.proto.internals.v1.ServiceInvocation.CallLocalStream:output_type -> dapr.proto.internals.v1.InternalInvokeResponseStream
	2,  // 24: dapr.proto.internals.v1.ServiceInvocation.CallActorStream:output_type -> dapr.proto.internals.v1.InternalInvokeResponse
	20, // [20:25] is the sub-list for method output_type
	15, // [15:20] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_dapr_proto_internals_v1_service_invocation_proto_init() }
func file_dapr_proto_internals_v1_service_invocation_proto_init() {
	if File_dapr_proto_internals_v1_service_invocation_proto != nil {
		return
	}
	file_dapr_proto_internals_v1_reminders_proto_init()
	file_dapr_proto_internals_v1_apiversion_proto_init()
	file_dapr_proto_internals_v1_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Actor); i {
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
		file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalInvokeRequest); i {
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
		file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalInvokeResponse); i {
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
		file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalInvokeRequestStream); i {
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
		file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalInvokeResponseStream); i {
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
		file_dapr_proto_internals_v1_service_invocation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStringValue); i {
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
			RawDescriptor: file_dapr_proto_internals_v1_service_invocation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dapr_proto_internals_v1_service_invocation_proto_goTypes,
		DependencyIndexes: file_dapr_proto_internals_v1_service_invocation_proto_depIdxs,
		MessageInfos:      file_dapr_proto_internals_v1_service_invocation_proto_msgTypes,
	}.Build()
	File_dapr_proto_internals_v1_service_invocation_proto = out.File
	file_dapr_proto_internals_v1_service_invocation_proto_rawDesc = nil
	file_dapr_proto_internals_v1_service_invocation_proto_goTypes = nil
	file_dapr_proto_internals_v1_service_invocation_proto_depIdxs = nil
}
