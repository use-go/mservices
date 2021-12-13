// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.10.0
// source: proto/helloworld/helloworld.proto

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

type Helloworld struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Helloworld) Reset() {
	*x = Helloworld{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Helloworld) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Helloworld) ProtoMessage() {}

func (x *Helloworld) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Helloworld.ProtoReflect.Descriptor instead.
func (*Helloworld) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *Helloworld) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloworldList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int32         `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size       int32         `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	TotalCount int32         `protobuf:"varint,3,opt,name=total_count,proto3" json:"total_count,omitempty"`
	Data       []*Helloworld `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *HelloworldList) Reset() {
	*x = HelloworldList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloworldList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloworldList) ProtoMessage() {}

func (x *HelloworldList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloworldList.ProtoReflect.Descriptor instead.
func (*HelloworldList) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *HelloworldList) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *HelloworldList) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *HelloworldList) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *HelloworldList) GetData() []*Helloworld {
	if x != nil {
		return x.Data
	}
	return nil
}

type HelloworldFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size  int32  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Order int32  `protobuf:"varint,3,opt,name=order,proto3" json:"order,omitempty"`
	Id    int32  `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloworldFilter) Reset() {
	*x = HelloworldFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloworldFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloworldFilter) ProtoMessage() {}

func (x *HelloworldFilter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_helloworld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloworldFilter.ProtoReflect.Descriptor instead.
func (*HelloworldFilter) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_helloworld_proto_rawDescGZIP(), []int{2}
}

func (x *HelloworldFilter) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *HelloworldFilter) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *HelloworldFilter) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *HelloworldFilter) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HelloworldFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_helloworld_helloworld_proto protoreflect.FileDescriptor

var file_proto_helloworld_helloworld_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22,
	0x20, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x86, 0x01, 0x0a, 0x0e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x74, 0x0a, 0x10, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_helloworld_helloworld_proto_rawDescOnce sync.Once
	file_proto_helloworld_helloworld_proto_rawDescData = file_proto_helloworld_helloworld_proto_rawDesc
)

func file_proto_helloworld_helloworld_proto_rawDescGZIP() []byte {
	file_proto_helloworld_helloworld_proto_rawDescOnce.Do(func() {
		file_proto_helloworld_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_helloworld_helloworld_proto_rawDescData)
	})
	return file_proto_helloworld_helloworld_proto_rawDescData
}

var file_proto_helloworld_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_helloworld_helloworld_proto_goTypes = []interface{}{
	(*Helloworld)(nil),       // 0: helloworld.Helloworld
	(*HelloworldList)(nil),   // 1: helloworld.HelloworldList
	(*HelloworldFilter)(nil), // 2: helloworld.HelloworldFilter
}
var file_proto_helloworld_helloworld_proto_depIdxs = []int32{
	0, // 0: helloworld.HelloworldList.data:type_name -> helloworld.Helloworld
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_helloworld_helloworld_proto_init() }
func file_proto_helloworld_helloworld_proto_init() {
	if File_proto_helloworld_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_helloworld_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Helloworld); i {
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
		file_proto_helloworld_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloworldList); i {
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
		file_proto_helloworld_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloworldFilter); i {
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
			RawDescriptor: file_proto_helloworld_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_helloworld_helloworld_proto_goTypes,
		DependencyIndexes: file_proto_helloworld_helloworld_proto_depIdxs,
		MessageInfos:      file_proto_helloworld_helloworld_proto_msgTypes,
	}.Build()
	File_proto_helloworld_helloworld_proto = out.File
	file_proto_helloworld_helloworld_proto_rawDesc = nil
	file_proto_helloworld_helloworld_proto_goTypes = nil
	file_proto_helloworld_helloworld_proto_depIdxs = nil
}
