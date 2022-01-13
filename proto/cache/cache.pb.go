// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/cache/cache.proto

package cache

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

type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,json=name,proto3" json:"key,omitempty"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cache_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cache_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_proto_cache_cache_proto_rawDescGZIP(), []int{0}
}

func (x *GetReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetRes) Reset() {
	*x = GetRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cache_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRes) ProtoMessage() {}

func (x *GetRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cache_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRes.ProtoReflect.Descriptor instead.
func (*GetRes) Descriptor() ([]byte, []int) {
	return file_proto_cache_cache_proto_rawDescGZIP(), []int{1}
}

func (x *GetRes) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type SetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,json=name,proto3" json:"key,omitempty"`
	Value  []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Expire int64  `protobuf:"varint,3,opt,name=expire,proto3" json:"expire,omitempty"`
}

func (x *SetReq) Reset() {
	*x = SetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cache_cache_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetReq) ProtoMessage() {}

func (x *SetReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cache_cache_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetReq.ProtoReflect.Descriptor instead.
func (*SetReq) Descriptor() ([]byte, []int) {
	return file_proto_cache_cache_proto_rawDescGZIP(), []int{2}
}

func (x *SetReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetReq) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *SetReq) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

type SetRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetRes) Reset() {
	*x = SetRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cache_cache_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRes) ProtoMessage() {}

func (x *SetRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cache_cache_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRes.ProtoReflect.Descriptor instead.
func (*SetRes) Descriptor() ([]byte, []int) {
	return file_proto_cache_cache_proto_rawDescGZIP(), []int{3}
}

type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,json=name,proto3" json:"key,omitempty"`
	Value  []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Expire int64  `protobuf:"varint,3,opt,name=expire,proto3" json:"expire,omitempty"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cache_cache_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cache_cache_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_proto_cache_cache_proto_rawDescGZIP(), []int{4}
}

func (x *AddReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AddReq) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *AddReq) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

type AddRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddRes) Reset() {
	*x = AddRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cache_cache_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRes) ProtoMessage() {}

func (x *AddRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cache_cache_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRes.ProtoReflect.Descriptor instead.
func (*AddRes) Descriptor() ([]byte, []int) {
	return file_proto_cache_cache_proto_rawDescGZIP(), []int{5}
}

type ReplaceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,json=name,proto3" json:"key,omitempty"`
	Value  []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Expire int64  `protobuf:"varint,3,opt,name=expire,proto3" json:"expire,omitempty"`
}

func (x *ReplaceReq) Reset() {
	*x = ReplaceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cache_cache_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplaceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceReq) ProtoMessage() {}

func (x *ReplaceReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cache_cache_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceReq.ProtoReflect.Descriptor instead.
func (*ReplaceReq) Descriptor() ([]byte, []int) {
	return file_proto_cache_cache_proto_rawDescGZIP(), []int{6}
}

func (x *ReplaceReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ReplaceReq) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ReplaceReq) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

type ReplaceRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReplaceRes) Reset() {
	*x = ReplaceRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cache_cache_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplaceRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceRes) ProtoMessage() {}

func (x *ReplaceRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cache_cache_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceRes.ProtoReflect.Descriptor instead.
func (*ReplaceRes) Descriptor() ([]byte, []int) {
	return file_proto_cache_cache_proto_rawDescGZIP(), []int{7}
}

type DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,json=name,proto3" json:"key,omitempty"`
}

func (x *DeleteReq) Reset() {
	*x = DeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cache_cache_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cache_cache_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReq.ProtoReflect.Descriptor instead.
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return file_proto_cache_cache_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRes) Reset() {
	*x = DeleteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cache_cache_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRes) ProtoMessage() {}

func (x *DeleteRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cache_cache_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRes.ProtoReflect.Descriptor instead.
func (*DeleteRes) Descriptor() ([]byte, []int) {
	return file_proto_cache_cache_proto_rawDescGZIP(), []int{9}
}

type IncrementReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,json=name,proto3" json:"key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IncrementReq) Reset() {
	*x = IncrementReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cache_cache_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrementReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrementReq) ProtoMessage() {}

func (x *IncrementReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cache_cache_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrementReq.ProtoReflect.Descriptor instead.
func (*IncrementReq) Descriptor() ([]byte, []int) {
	return file_proto_cache_cache_proto_rawDescGZIP(), []int{10}
}

func (x *IncrementReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *IncrementReq) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type IncrementRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IncrementRes) Reset() {
	*x = IncrementRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cache_cache_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrementRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrementRes) ProtoMessage() {}

func (x *IncrementRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cache_cache_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrementRes.ProtoReflect.Descriptor instead.
func (*IncrementRes) Descriptor() ([]byte, []int) {
	return file_proto_cache_cache_proto_rawDescGZIP(), []int{11}
}

func (x *IncrementRes) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type DecrementReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,json=name,proto3" json:"key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DecrementReq) Reset() {
	*x = DecrementReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cache_cache_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecrementReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecrementReq) ProtoMessage() {}

func (x *DecrementReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cache_cache_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecrementReq.ProtoReflect.Descriptor instead.
func (*DecrementReq) Descriptor() ([]byte, []int) {
	return file_proto_cache_cache_proto_rawDescGZIP(), []int{12}
}

func (x *DecrementReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *DecrementReq) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type DecrementRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DecrementRes) Reset() {
	*x = DecrementRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cache_cache_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecrementRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecrementRes) ProtoMessage() {}

func (x *DecrementRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cache_cache_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecrementRes.ProtoReflect.Descriptor instead.
func (*DecrementRes) Descriptor() ([]byte, []int) {
	return file_proto_cache_cache_proto_rawDescGZIP(), []int{13}
}

func (x *DecrementRes) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type FlushReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FlushReq) Reset() {
	*x = FlushReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cache_cache_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlushReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlushReq) ProtoMessage() {}

func (x *FlushReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cache_cache_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlushReq.ProtoReflect.Descriptor instead.
func (*FlushReq) Descriptor() ([]byte, []int) {
	return file_proto_cache_cache_proto_rawDescGZIP(), []int{14}
}

type FlushRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FlushRes) Reset() {
	*x = FlushRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cache_cache_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlushRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlushRes) ProtoMessage() {}

func (x *FlushRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cache_cache_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlushRes.ProtoReflect.Descriptor instead.
func (*FlushRes) Descriptor() ([]byte, []int) {
	return file_proto_cache_cache_proto_rawDescGZIP(), []int{15}
}

var File_proto_cache_cache_proto protoreflect.FileDescriptor

var file_proto_cache_cache_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x22, 0x1b, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x11, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1e, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x49, 0x0a,
	0x06, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x11, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x08, 0x0a, 0x06, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x22, 0x49, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x11, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x08, 0x0a,
	0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x22, 0x4d, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x11, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x22, 0x1e, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x11, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x22, 0x37, 0x0a, 0x0c, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x11, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x24, 0x0a, 0x0c, 0x49, 0x6e,
	0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x37, 0x0a, 0x0c, 0x44, 0x65, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x11, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x24, 0x0a, 0x0c, 0x44, 0x65, 0x63,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x0a, 0x0a, 0x08, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x22, 0x0a, 0x0a, 0x08, 0x46,
	0x6c, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cache_cache_proto_rawDescOnce sync.Once
	file_proto_cache_cache_proto_rawDescData = file_proto_cache_cache_proto_rawDesc
)

func file_proto_cache_cache_proto_rawDescGZIP() []byte {
	file_proto_cache_cache_proto_rawDescOnce.Do(func() {
		file_proto_cache_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cache_cache_proto_rawDescData)
	})
	return file_proto_cache_cache_proto_rawDescData
}

var file_proto_cache_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_proto_cache_cache_proto_goTypes = []interface{}{
	(*GetReq)(nil),       // 0: cache.GetReq
	(*GetRes)(nil),       // 1: cache.GetRes
	(*SetReq)(nil),       // 2: cache.SetReq
	(*SetRes)(nil),       // 3: cache.SetRes
	(*AddReq)(nil),       // 4: cache.AddReq
	(*AddRes)(nil),       // 5: cache.AddRes
	(*ReplaceReq)(nil),   // 6: cache.ReplaceReq
	(*ReplaceRes)(nil),   // 7: cache.ReplaceRes
	(*DeleteReq)(nil),    // 8: cache.DeleteReq
	(*DeleteRes)(nil),    // 9: cache.DeleteRes
	(*IncrementReq)(nil), // 10: cache.IncrementReq
	(*IncrementRes)(nil), // 11: cache.IncrementRes
	(*DecrementReq)(nil), // 12: cache.DecrementReq
	(*DecrementRes)(nil), // 13: cache.DecrementRes
	(*FlushReq)(nil),     // 14: cache.FlushReq
	(*FlushRes)(nil),     // 15: cache.FlushRes
}
var file_proto_cache_cache_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_cache_cache_proto_init() }
func file_proto_cache_cache_proto_init() {
	if File_proto_cache_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cache_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReq); i {
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
		file_proto_cache_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRes); i {
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
		file_proto_cache_cache_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetReq); i {
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
		file_proto_cache_cache_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRes); i {
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
		file_proto_cache_cache_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReq); i {
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
		file_proto_cache_cache_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRes); i {
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
		file_proto_cache_cache_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplaceReq); i {
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
		file_proto_cache_cache_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplaceRes); i {
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
		file_proto_cache_cache_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReq); i {
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
		file_proto_cache_cache_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRes); i {
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
		file_proto_cache_cache_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrementReq); i {
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
		file_proto_cache_cache_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrementRes); i {
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
		file_proto_cache_cache_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecrementReq); i {
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
		file_proto_cache_cache_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecrementRes); i {
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
		file_proto_cache_cache_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlushReq); i {
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
		file_proto_cache_cache_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlushRes); i {
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
			RawDescriptor: file_proto_cache_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_cache_cache_proto_goTypes,
		DependencyIndexes: file_proto_cache_cache_proto_depIdxs,
		MessageInfos:      file_proto_cache_cache_proto_msgTypes,
	}.Build()
	File_proto_cache_cache_proto = out.File
	file_proto_cache_cache_proto_rawDesc = nil
	file_proto_cache_cache_proto_goTypes = nil
	file_proto_cache_cache_proto_depIdxs = nil
}