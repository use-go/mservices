// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/helloworld/info.proto

package helloworld

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

type InsertInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *InsertInfoRequest) Reset() {
	*x = InsertInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertInfoRequest) ProtoMessage() {}

func (x *InsertInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertInfoRequest.ProtoReflect.Descriptor instead.
func (*InsertInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_info_proto_rawDescGZIP(), []int{0}
}

func (x *InsertInfoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type InsertInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InsertInfoResponse) Reset() {
	*x = InsertInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertInfoResponse) ProtoMessage() {}

func (x *InsertInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertInfoResponse.ProtoReflect.Descriptor instead.
func (*InsertInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_info_proto_rawDescGZIP(), []int{1}
}

func (x *InsertInfoResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteInfoRequest) Reset() {
	*x = DeleteInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInfoRequest) ProtoMessage() {}

func (x *DeleteInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInfoRequest.ProtoReflect.Descriptor instead.
func (*DeleteInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_info_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteInfoRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteInfoResponse) Reset() {
	*x = DeleteInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInfoResponse) ProtoMessage() {}

func (x *DeleteInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInfoResponse.ProtoReflect.Descriptor instead.
func (*DeleteInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_info_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteInfoResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateInfoRequest) Reset() {
	*x = UpdateInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_info_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInfoRequest) ProtoMessage() {}

func (x *UpdateInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_info_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_info_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateInfoRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateInfoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateInfoResponse) Reset() {
	*x = UpdateInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_info_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInfoResponse) ProtoMessage() {}

func (x *UpdateInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_info_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInfoResponse.ProtoReflect.Descriptor instead.
func (*UpdateInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_info_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateInfoResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type QueryInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Page  int32  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Size  int32  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Order int32  `protobuf:"varint,5,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *QueryInfoRequest) Reset() {
	*x = QueryInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_info_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryInfoRequest) ProtoMessage() {}

func (x *QueryInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_info_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryInfoRequest.ProtoReflect.Descriptor instead.
func (*QueryInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_info_proto_rawDescGZIP(), []int{6}
}

func (x *QueryInfoRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QueryInfoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QueryInfoRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *QueryInfoRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *QueryInfoRequest) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

type QueryInfoResponseItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *QueryInfoResponseItem) Reset() {
	*x = QueryInfoResponseItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_info_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryInfoResponseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryInfoResponseItem) ProtoMessage() {}

func (x *QueryInfoResponseItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_info_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryInfoResponseItem.ProtoReflect.Descriptor instead.
func (*QueryInfoResponseItem) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_info_proto_rawDescGZIP(), []int{7}
}

func (x *QueryInfoResponseItem) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QueryInfoResponseItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type QueryInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []*QueryInfoResponseItem `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Page       int32                    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size       int32                    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	TotalCount int32                    `protobuf:"varint,4,opt,name=total_count,proto3" json:"total_count,omitempty"`
}

func (x *QueryInfoResponse) Reset() {
	*x = QueryInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_info_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryInfoResponse) ProtoMessage() {}

func (x *QueryInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_info_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryInfoResponse.ProtoReflect.Descriptor instead.
func (*QueryInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_info_proto_rawDescGZIP(), []int{8}
}

func (x *QueryInfoResponse) GetData() []*QueryInfoResponseItem {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryInfoResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *QueryInfoResponse) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *QueryInfoResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type QueryInfoDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *QueryInfoDetailRequest) Reset() {
	*x = QueryInfoDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_info_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryInfoDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryInfoDetailRequest) ProtoMessage() {}

func (x *QueryInfoDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_info_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryInfoDetailRequest.ProtoReflect.Descriptor instead.
func (*QueryInfoDetailRequest) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_info_proto_rawDescGZIP(), []int{9}
}

func (x *QueryInfoDetailRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type QueryInfoDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *QueryInfoDetailResponse) Reset() {
	*x = QueryInfoDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_info_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryInfoDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryInfoDetailResponse) ProtoMessage() {}

func (x *QueryInfoDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_info_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryInfoDetailResponse.ProtoReflect.Descriptor instead.
func (*QueryInfoDetailResponse) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_info_proto_rawDescGZIP(), []int{10}
}

func (x *QueryInfoDetailResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QueryInfoDetailResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_helloworld_info_proto protoreflect.FileDescriptor

var file_proto_helloworld_info_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x27, 0x0a, 0x11, 0x49, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x74, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x28, 0x0a, 0x16,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x17, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_helloworld_info_proto_rawDescOnce sync.Once
	file_proto_helloworld_info_proto_rawDescData = file_proto_helloworld_info_proto_rawDesc
)

func file_proto_helloworld_info_proto_rawDescGZIP() []byte {
	file_proto_helloworld_info_proto_rawDescOnce.Do(func() {
		file_proto_helloworld_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_helloworld_info_proto_rawDescData)
	})
	return file_proto_helloworld_info_proto_rawDescData
}

var file_proto_helloworld_info_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_helloworld_info_proto_goTypes = []interface{}{
	(*InsertInfoRequest)(nil),       // 0: helloworld.InsertInfoRequest
	(*InsertInfoResponse)(nil),      // 1: helloworld.InsertInfoResponse
	(*DeleteInfoRequest)(nil),       // 2: helloworld.DeleteInfoRequest
	(*DeleteInfoResponse)(nil),      // 3: helloworld.DeleteInfoResponse
	(*UpdateInfoRequest)(nil),       // 4: helloworld.UpdateInfoRequest
	(*UpdateInfoResponse)(nil),      // 5: helloworld.UpdateInfoResponse
	(*QueryInfoRequest)(nil),        // 6: helloworld.QueryInfoRequest
	(*QueryInfoResponseItem)(nil),   // 7: helloworld.QueryInfoResponseItem
	(*QueryInfoResponse)(nil),       // 8: helloworld.QueryInfoResponse
	(*QueryInfoDetailRequest)(nil),  // 9: helloworld.QueryInfoDetailRequest
	(*QueryInfoDetailResponse)(nil), // 10: helloworld.QueryInfoDetailResponse
}
var file_proto_helloworld_info_proto_depIdxs = []int32{
	7, // 0: helloworld.QueryInfoResponse.data:type_name -> helloworld.QueryInfoResponseItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_helloworld_info_proto_init() }
func file_proto_helloworld_info_proto_init() {
	if File_proto_helloworld_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_helloworld_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertInfoRequest); i {
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
		file_proto_helloworld_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertInfoResponse); i {
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
		file_proto_helloworld_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInfoRequest); i {
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
		file_proto_helloworld_info_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInfoResponse); i {
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
		file_proto_helloworld_info_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInfoRequest); i {
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
		file_proto_helloworld_info_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInfoResponse); i {
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
		file_proto_helloworld_info_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryInfoRequest); i {
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
		file_proto_helloworld_info_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryInfoResponseItem); i {
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
		file_proto_helloworld_info_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryInfoResponse); i {
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
		file_proto_helloworld_info_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryInfoDetailRequest); i {
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
		file_proto_helloworld_info_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryInfoDetailResponse); i {
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
			RawDescriptor: file_proto_helloworld_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_helloworld_info_proto_goTypes,
		DependencyIndexes: file_proto_helloworld_info_proto_depIdxs,
		MessageInfos:      file_proto_helloworld_info_proto_msgTypes,
	}.Build()
	File_proto_helloworld_info_proto = out.File
	file_proto_helloworld_info_proto_rawDesc = nil
	file_proto_helloworld_info_proto_goTypes = nil
	file_proto_helloworld_info_proto_depIdxs = nil
}
