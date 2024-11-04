// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.2
// source: v2.proto

// @formatter:off
// protoc --proto_path=proto --go_out=. --go-grpc_out=. proto/v2.proto

package v2

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

type Event_Kind int32

const (
	Event_UNKNOWN Event_Kind = 0
	Event_INSERT  Event_Kind = 1
	Event_UPDATE  Event_Kind = 2
	Event_DELETE  Event_Kind = 3
)

// Enum value maps for Event_Kind.
var (
	Event_Kind_name = map[int32]string{
		0: "UNKNOWN",
		1: "INSERT",
		2: "UPDATE",
		3: "DELETE",
	}
	Event_Kind_value = map[string]int32{
		"UNKNOWN": 0,
		"INSERT":  1,
		"UPDATE":  2,
		"DELETE":  3,
	}
)

func (x Event_Kind) Enum() *Event_Kind {
	p := new(Event_Kind)
	*p = x
	return p
}

func (x Event_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_v2_proto_enumTypes[0].Descriptor()
}

func (Event_Kind) Type() protoreflect.EnumType {
	return &file_v2_proto_enumTypes[0]
}

func (x Event_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_Kind.Descriptor instead.
func (Event_Kind) EnumDescriptor() ([]byte, []int) {
	return file_v2_proto_rawDescGZIP(), []int{5, 0}
}

// Subscription ...
type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ns   string `protobuf:"bytes,1,opt,name=ns,proto3" json:"ns,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	mi := &file_v2_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_v2_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_v2_proto_rawDescGZIP(), []int{0}
}

func (x *Subscription) GetNs() string {
	if x != nil {
		return x.Ns
	}
	return ""
}

func (x *Subscription) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// Schema ...
type Schema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ns string `protobuf:"bytes,1,opt,name=ns,proto3" json:"ns,omitempty"`
}

func (x *Schema) Reset() {
	*x = Schema{}
	mi := &file_v2_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Schema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema) ProtoMessage() {}

func (x *Schema) ProtoReflect() protoreflect.Message {
	mi := &file_v2_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema.ProtoReflect.Descriptor instead.
func (*Schema) Descriptor() ([]byte, []int) {
	return file_v2_proto_rawDescGZIP(), []int{1}
}

func (x *Schema) GetNs() string {
	if x != nil {
		return x.Ns
	}
	return ""
}

// Request ...
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ns        string      `protobuf:"bytes,1,opt,name=ns,proto3" json:"ns,omitempty"`
	Ttx       string      `protobuf:"bytes,2,opt,name=ttx,proto3" json:"ttx,omitempty"`
	Query     string      `protobuf:"bytes,5,opt,name=query,proto3" json:"query,omitempty"`
	Mutations []*Mutation `protobuf:"bytes,8,rep,name=mutations,proto3" json:"mutations,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_v2_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_v2_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_v2_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetNs() string {
	if x != nil {
		return x.Ns
	}
	return ""
}

func (x *Request) GetTtx() string {
	if x != nil {
		return x.Ttx
	}
	return ""
}

func (x *Request) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Request) GetMutations() []*Mutation {
	if x != nil {
		return x.Mutations
	}
	return nil
}

// Txn is the transaction context.
type Ttx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ns string `protobuf:"bytes,1,opt,name=ns,proto3" json:"ns,omitempty"`
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Ttx) Reset() {
	*x = Ttx{}
	mi := &file_v2_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ttx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ttx) ProtoMessage() {}

func (x *Ttx) ProtoReflect() protoreflect.Message {
	mi := &file_v2_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ttx.ProtoReflect.Descriptor instead.
func (*Ttx) Descriptor() ([]byte, []int) {
	return file_v2_proto_rawDescGZIP(), []int{3}
}

func (x *Ttx) GetNs() string {
	if x != nil {
		return x.Ns
	}
	return ""
}

func (x *Ttx) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Mutation ...
type Mutation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Set []byte `protobuf:"bytes,8,opt,name=set,proto3" json:"set,omitempty"`
	Del []byte `protobuf:"bytes,9,opt,name=del,proto3" json:"del,omitempty"`
}

func (x *Mutation) Reset() {
	*x = Mutation{}
	mi := &file_v2_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Mutation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mutation) ProtoMessage() {}

func (x *Mutation) ProtoReflect() protoreflect.Message {
	mi := &file_v2_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mutation.ProtoReflect.Descriptor instead.
func (*Mutation) Descriptor() ([]byte, []int) {
	return file_v2_proto_rawDescGZIP(), []int{4}
}

func (x *Mutation) GetSet() []byte {
	if x != nil {
		return x.Set
	}
	return nil
}

func (x *Mutation) GetDel() []byte {
	if x != nil {
		return x.Del
	}
	return nil
}

// Event ...
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind     Event_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=graph.Event_Kind" json:"kind,omitempty"`
	Response *Response  `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_v2_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_v2_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_v2_proto_rawDescGZIP(), []int{5}
}

func (x *Event) GetKind() Event_Kind {
	if x != nil {
		return x.Kind
	}
	return Event_UNKNOWN
}

func (x *Event) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

// Response ...
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ns      string `protobuf:"bytes,1,opt,name=ns,proto3" json:"ns,omitempty"`
	Ttx     string `protobuf:"bytes,2,opt,name=ttx,proto3" json:"ttx,omitempty"`
	Payload []byte `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_v2_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_v2_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_v2_proto_rawDescGZIP(), []int{6}
}

func (x *Response) GetNs() string {
	if x != nil {
		return x.Ns
	}
	return ""
}

func (x *Response) GetTtx() string {
	if x != nil {
		return x.Ttx
	}
	return ""
}

func (x *Response) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// Empty ...
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_v2_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_v2_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_v2_proto_rawDescGZIP(), []int{7}
}

var File_v2_proto protoreflect.FileDescriptor

var file_v2_proto_rawDesc = []byte{
	0x0a, 0x08, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x22, 0x32, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6e,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x18, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6e, 0x73, 0x22,
	0x70, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6e, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x74, 0x78, 0x12, 0x14, 0x0a, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x2d, 0x0a, 0x09, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x25, 0x0a, 0x03, 0x54, 0x74, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6e, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x08, 0x4d, 0x75, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x73, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65, 0x6c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x64, 0x65, 0x6c, 0x22, 0x94, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x49,
	0x4e, 0x53, 0x45, 0x52, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x22,
	0x46, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6e, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x74, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x74, 0x78, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x32, 0xd2, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x09, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0c, 0x2e, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x30, 0x01, 0x12, 0x25, 0x0a, 0x05, 0x41,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x1a, 0x0d, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x12, 0x28, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0e, 0x2e, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x0a, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x54,
	0x74, 0x78, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x12, 0x0a, 0x2e, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x54, 0x74, 0x78, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_proto_rawDescOnce sync.Once
	file_v2_proto_rawDescData = file_v2_proto_rawDesc
)

func file_v2_proto_rawDescGZIP() []byte {
	file_v2_proto_rawDescOnce.Do(func() {
		file_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_proto_rawDescData)
	})
	return file_v2_proto_rawDescData
}

var file_v2_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_v2_proto_goTypes = []any{
	(Event_Kind)(0),      // 0: graph.Event.Kind
	(*Subscription)(nil), // 1: graph.Subscription
	(*Schema)(nil),       // 2: graph.Schema
	(*Request)(nil),      // 3: graph.Request
	(*Ttx)(nil),          // 4: graph.Ttx
	(*Mutation)(nil),     // 5: graph.Mutation
	(*Event)(nil),        // 6: graph.Event
	(*Response)(nil),     // 7: graph.Response
	(*Empty)(nil),        // 8: graph.Empty
}
var file_v2_proto_depIdxs = []int32{
	5, // 0: graph.Request.mutations:type_name -> graph.Mutation
	0, // 1: graph.Event.kind:type_name -> graph.Event.Kind
	7, // 2: graph.Event.response:type_name -> graph.Response
	1, // 3: graph.Data.Subscribe:input_type -> graph.Subscription
	2, // 4: graph.Data.Alter:input_type -> graph.Schema
	3, // 5: graph.Data.Query:input_type -> graph.Request
	4, // 6: graph.Data.Commit:input_type -> graph.Ttx
	4, // 7: graph.Data.Discard:input_type -> graph.Ttx
	6, // 8: graph.Data.Subscribe:output_type -> graph.Event
	2, // 9: graph.Data.Alter:output_type -> graph.Schema
	7, // 10: graph.Data.Query:output_type -> graph.Response
	8, // 11: graph.Data.Commit:output_type -> graph.Empty
	8, // 12: graph.Data.Discard:output_type -> graph.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v2_proto_init() }
func file_v2_proto_init() {
	if File_v2_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v2_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v2_proto_goTypes,
		DependencyIndexes: file_v2_proto_depIdxs,
		EnumInfos:         file_v2_proto_enumTypes,
		MessageInfos:      file_v2_proto_msgTypes,
	}.Build()
	File_v2_proto = out.File
	file_v2_proto_rawDesc = nil
	file_v2_proto_goTypes = nil
	file_v2_proto_depIdxs = nil
}
