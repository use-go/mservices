// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/user/handler.proto

package user

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

var File_proto_user_handler_proto protoreflect.FileDescriptor

var file_proto_user_handler_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x1a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe2, 0x04, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x26, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x44, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x59, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x16, 0x53,
	0x65, 0x6e, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0d, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_user_handler_proto_goTypes = []interface{}{
	(*CreateReq)(nil),                 // 0: user.CreateReq
	(*ReadReq)(nil),                   // 1: user.ReadReq
	(*UpdateReq)(nil),                 // 2: user.UpdateReq
	(*DeleteReq)(nil),                 // 3: user.DeleteReq
	(*ListReq)(nil),                   // 4: user.ListReq
	(*UpdatePasswordReq)(nil),         // 5: user.UpdatePasswordReq
	(*VerifyEmailReq)(nil),            // 6: user.VerifyEmailReq
	(*SendVerificationEmailReq)(nil),  // 7: user.SendVerificationEmailReq
	(*SendPasswordResetEmailReq)(nil), // 8: user.SendPasswordResetEmailReq
	(*ResetPasswordReq)(nil),          // 9: user.ResetPasswordReq
	(*CreateRes)(nil),                 // 10: user.CreateRes
	(*ReadRes)(nil),                   // 11: user.ReadRes
	(*UpdateRes)(nil),                 // 12: user.UpdateRes
	(*DeleteRes)(nil),                 // 13: user.DeleteRes
	(*ListRes)(nil),                   // 14: user.ListRes
	(*UpdatePasswordRes)(nil),         // 15: user.UpdatePasswordRes
	(*VerifyEmailRes)(nil),            // 16: user.VerifyEmailRes
	(*SendVerificationEmailRes)(nil),  // 17: user.SendVerificationEmailRes
	(*SendPasswordResetEmailRes)(nil), // 18: user.SendPasswordResetEmailRes
	(*ResetPasswordRes)(nil),          // 19: user.ResetPasswordRes
}
var file_proto_user_handler_proto_depIdxs = []int32{
	0,  // 0: user.Account.Create:input_type -> user.CreateReq
	1,  // 1: user.Account.Read:input_type -> user.ReadReq
	2,  // 2: user.Account.Update:input_type -> user.UpdateReq
	3,  // 3: user.Account.Delete:input_type -> user.DeleteReq
	4,  // 4: user.Account.List:input_type -> user.ListReq
	5,  // 5: user.Account.UpdatePassword:input_type -> user.UpdatePasswordReq
	6,  // 6: user.Account.VerifyEmail:input_type -> user.VerifyEmailReq
	7,  // 7: user.Account.SendVerificationEmail:input_type -> user.SendVerificationEmailReq
	8,  // 8: user.Account.SendPasswordResetEmail:input_type -> user.SendPasswordResetEmailReq
	9,  // 9: user.Account.ResetPassword:input_type -> user.ResetPasswordReq
	10, // 10: user.Account.Create:output_type -> user.CreateRes
	11, // 11: user.Account.Read:output_type -> user.ReadRes
	12, // 12: user.Account.Update:output_type -> user.UpdateRes
	13, // 13: user.Account.Delete:output_type -> user.DeleteRes
	14, // 14: user.Account.List:output_type -> user.ListRes
	15, // 15: user.Account.UpdatePassword:output_type -> user.UpdatePasswordRes
	16, // 16: user.Account.VerifyEmail:output_type -> user.VerifyEmailRes
	17, // 17: user.Account.SendVerificationEmail:output_type -> user.SendVerificationEmailRes
	18, // 18: user.Account.SendPasswordResetEmail:output_type -> user.SendPasswordResetEmailRes
	19, // 19: user.Account.ResetPassword:output_type -> user.ResetPasswordRes
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_user_handler_proto_init() }
func file_proto_user_handler_proto_init() {
	if File_proto_user_handler_proto != nil {
		return
	}
	file_proto_user_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_user_handler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_handler_proto_goTypes,
		DependencyIndexes: file_proto_user_handler_proto_depIdxs,
	}.Build()
	File_proto_user_handler_proto = out.File
	file_proto_user_handler_proto_rawDesc = nil
	file_proto_user_handler_proto_goTypes = nil
	file_proto_user_handler_proto_depIdxs = nil
}