// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: sso/sso.proto

package sso_contract

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

type SignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`       // Email of the user to register.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` // Password of the user to register.
}

func (x *SignupRequest) Reset() {
	*x = SignupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_sso_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupRequest) ProtoMessage() {}

func (x *SignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupRequest.ProtoReflect.Descriptor instead.
func (*SignupRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{0}
}

func (x *SignupRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignupRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // User ID of the registered user.
}

func (x *SignupResponse) Reset() {
	*x = SignupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_sso_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupResponse) ProtoMessage() {}

func (x *SignupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupResponse.ProtoReflect.Descriptor instead.
func (*SignupResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{1}
}

func (x *SignupResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type SigninRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`               // Email of the user to login.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`         // Password of the user to login.
	AppId    int32  `protobuf:"varint,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"` // ID of the app to login to.
}

func (x *SigninRequest) Reset() {
	*x = SigninRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_sso_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigninRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigninRequest) ProtoMessage() {}

func (x *SigninRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigninRequest.ProtoReflect.Descriptor instead.
func (*SigninRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{2}
}

func (x *SigninRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SigninRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SigninRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type SigninResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` // Auth token of the logged in user.
}

func (x *SigninResponse) Reset() {
	*x = SigninResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_sso_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigninResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigninResponse) ProtoMessage() {}

func (x *SigninResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigninResponse.ProtoReflect.Descriptor instead.
func (*SigninResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{3}
}

func (x *SigninResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type IsAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // User ID to validate.
}

func (x *IsAdminRequest) Reset() {
	*x = IsAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_sso_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAdminRequest) ProtoMessage() {}

func (x *IsAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAdminRequest.ProtoReflect.Descriptor instead.
func (*IsAdminRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{4}
}

func (x *IsAdminRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type IsAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAdmin bool `protobuf:"varint,1,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"` // Indicate whether the user is admin.
}

func (x *IsAdminResponse) Reset() {
	*x = IsAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_sso_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAdminResponse) ProtoMessage() {}

func (x *IsAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAdminResponse.ProtoReflect.Descriptor instead.
func (*IsAdminResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{5}
}

func (x *IsAdminResponse) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

var File_sso_sso_proto protoreflect.FileDescriptor

var file_sso_sso_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x73, 0x6f, 0x2f, 0x73, 0x73, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x73, 0x73, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x22, 0x41, 0x0a,
	0x0d, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x29, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x0d, 0x53,
	0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x29, 0x0a,
	0x0e, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0f, 0x49, 0x73, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x32, 0xd8, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x43, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x1b, 0x2e, 0x73, 0x73, 0x6f, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x73, 0x6f, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x12, 0x1b,
	0x2e, 0x73, 0x73, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x73,
	0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x49, 0x73, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x73, 0x73, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x2e, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x73, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x3b, 0x73, 0x73, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sso_sso_proto_rawDescOnce sync.Once
	file_sso_sso_proto_rawDescData = file_sso_sso_proto_rawDesc
)

func file_sso_sso_proto_rawDescGZIP() []byte {
	file_sso_sso_proto_rawDescOnce.Do(func() {
		file_sso_sso_proto_rawDescData = protoimpl.X.CompressGZIP(file_sso_sso_proto_rawDescData)
	})
	return file_sso_sso_proto_rawDescData
}

var file_sso_sso_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sso_sso_proto_goTypes = []interface{}{
	(*SignupRequest)(nil),   // 0: sso_contract.SignupRequest
	(*SignupResponse)(nil),  // 1: sso_contract.SignupResponse
	(*SigninRequest)(nil),   // 2: sso_contract.SigninRequest
	(*SigninResponse)(nil),  // 3: sso_contract.SigninResponse
	(*IsAdminRequest)(nil),  // 4: sso_contract.IsAdminRequest
	(*IsAdminResponse)(nil), // 5: sso_contract.IsAdminResponse
}
var file_sso_sso_proto_depIdxs = []int32{
	0, // 0: sso_contract.Auth.Signup:input_type -> sso_contract.SignupRequest
	2, // 1: sso_contract.Auth.Signin:input_type -> sso_contract.SigninRequest
	4, // 2: sso_contract.Auth.IsAdmin:input_type -> sso_contract.IsAdminRequest
	1, // 3: sso_contract.Auth.Signup:output_type -> sso_contract.SignupResponse
	3, // 4: sso_contract.Auth.Signin:output_type -> sso_contract.SigninResponse
	5, // 5: sso_contract.Auth.IsAdmin:output_type -> sso_contract.IsAdminResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sso_sso_proto_init() }
func file_sso_sso_proto_init() {
	if File_sso_sso_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sso_sso_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupRequest); i {
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
		file_sso_sso_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupResponse); i {
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
		file_sso_sso_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigninRequest); i {
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
		file_sso_sso_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigninResponse); i {
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
		file_sso_sso_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsAdminRequest); i {
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
		file_sso_sso_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsAdminResponse); i {
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
			RawDescriptor: file_sso_sso_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sso_sso_proto_goTypes,
		DependencyIndexes: file_sso_sso_proto_depIdxs,
		MessageInfos:      file_sso_sso_proto_msgTypes,
	}.Build()
	File_sso_sso_proto = out.File
	file_sso_sso_proto_rawDesc = nil
	file_sso_sso_proto_goTypes = nil
	file_sso_sso_proto_depIdxs = nil
}
