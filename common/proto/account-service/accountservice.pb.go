// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: accountservice.proto

package account_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginProvider int32

const (
	LoginProvider_CREDENTIALS  LoginProvider = 0
	LoginProvider_GOOGLESIGNIN LoginProvider = 1
)

// Enum value maps for LoginProvider.
var (
	LoginProvider_name = map[int32]string{
		0: "CREDENTIALS",
		1: "GOOGLESIGNIN",
	}
	LoginProvider_value = map[string]int32{
		"CREDENTIALS":  0,
		"GOOGLESIGNIN": 1,
	}
)

func (x LoginProvider) Enum() *LoginProvider {
	p := new(LoginProvider)
	*p = x
	return p
}

func (x LoginProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_accountservice_proto_enumTypes[0].Descriptor()
}

func (LoginProvider) Type() protoreflect.EnumType {
	return &file_accountservice_proto_enumTypes[0]
}

func (x LoginProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginProvider.Descriptor instead.
func (LoginProvider) EnumDescriptor() ([]byte, []int) {
	return file_accountservice_proto_rawDescGZIP(), []int{0}
}

type GetPingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventLabel string `protobuf:"bytes,1,opt,name=eventLabel,proto3" json:"eventLabel,omitempty"`
}

func (x *GetPingRequest) Reset() {
	*x = GetPingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPingRequest) ProtoMessage() {}

func (x *GetPingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accountservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPingRequest.ProtoReflect.Descriptor instead.
func (*GetPingRequest) Descriptor() ([]byte, []int) {
	return file_accountservice_proto_rawDescGZIP(), []int{0}
}

func (x *GetPingRequest) GetEventLabel() string {
	if x != nil {
		return x.EventLabel
	}
	return ""
}

type GetPingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetPingResponse) Reset() {
	*x = GetPingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPingResponse) ProtoMessage() {}

func (x *GetPingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accountservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPingResponse.ProtoReflect.Descriptor instead.
func (*GetPingResponse) Descriptor() ([]byte, []int) {
	return file_accountservice_proto_rawDescGZIP(), []int{1}
}

func (x *GetPingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PostLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider   LoginProvider         `protobuf:"varint,1,opt,name=provider,proto3,enum=accountservice.LoginProvider" json:"provider,omitempty"`
	Properties map[string]*anypb.Any `protobuf:"bytes,2,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PostLoginRequest) Reset() {
	*x = PostLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostLoginRequest) ProtoMessage() {}

func (x *PostLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accountservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostLoginRequest.ProtoReflect.Descriptor instead.
func (*PostLoginRequest) Descriptor() ([]byte, []int) {
	return file_accountservice_proto_rawDescGZIP(), []int{2}
}

func (x *PostLoginRequest) GetProvider() LoginProvider {
	if x != nil {
		return x.Provider
	}
	return LoginProvider_CREDENTIALS
}

func (x *PostLoginRequest) GetProperties() map[string]*anypb.Any {
	if x != nil {
		return x.Properties
	}
	return nil
}

type TokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *TokenResponse) Reset() {
	*x = TokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenResponse) ProtoMessage() {}

func (x *TokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accountservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenResponse.ProtoReflect.Descriptor instead.
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return file_accountservice_proto_rawDescGZIP(), []int{3}
}

func (x *TokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *TokenResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type PostLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *TokenResponse `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *PostLoginResponse) Reset() {
	*x = PostLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostLoginResponse) ProtoMessage() {}

func (x *PostLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accountservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostLoginResponse.ProtoReflect.Descriptor instead.
func (*PostLoginResponse) Descriptor() ([]byte, []int) {
	return file_accountservice_proto_rawDescGZIP(), []int{4}
}

func (x *PostLoginResponse) GetToken() *TokenResponse {
	if x != nil {
		return x.Token
	}
	return nil
}

var File_accountservice_proto protoreflect.FileDescriptor

var file_accountservice_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x30, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x22, 0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0xf4, 0x01, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x12, 0x50, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x1a, 0x53, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x55, 0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x48,
	0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2a, 0x32, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x52, 0x45,
	0x44, 0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c, 0x53, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x4f,
	0x4f, 0x47, 0x4c, 0x45, 0x53, 0x49, 0x47, 0x4e, 0x49, 0x4e, 0x10, 0x01, 0x32, 0xab, 0x01, 0x0a,
	0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x49, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x05, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x20, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x65, 0x70, 0x68, 0x61, 0x6e,
	0x76, 0x65, 0x62, 0x72, 0x69, 0x61, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2d,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accountservice_proto_rawDescOnce sync.Once
	file_accountservice_proto_rawDescData = file_accountservice_proto_rawDesc
)

func file_accountservice_proto_rawDescGZIP() []byte {
	file_accountservice_proto_rawDescOnce.Do(func() {
		file_accountservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_accountservice_proto_rawDescData)
	})
	return file_accountservice_proto_rawDescData
}

var file_accountservice_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_accountservice_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_accountservice_proto_goTypes = []interface{}{
	(LoginProvider)(0),        // 0: accountservice.LoginProvider
	(*GetPingRequest)(nil),    // 1: accountservice.GetPingRequest
	(*GetPingResponse)(nil),   // 2: accountservice.GetPingResponse
	(*PostLoginRequest)(nil),  // 3: accountservice.PostLoginRequest
	(*TokenResponse)(nil),     // 4: accountservice.TokenResponse
	(*PostLoginResponse)(nil), // 5: accountservice.PostLoginResponse
	nil,                       // 6: accountservice.PostLoginRequest.PropertiesEntry
	(*anypb.Any)(nil),         // 7: google.protobuf.Any
}
var file_accountservice_proto_depIdxs = []int32{
	0, // 0: accountservice.PostLoginRequest.provider:type_name -> accountservice.LoginProvider
	6, // 1: accountservice.PostLoginRequest.properties:type_name -> accountservice.PostLoginRequest.PropertiesEntry
	4, // 2: accountservice.PostLoginResponse.token:type_name -> accountservice.TokenResponse
	7, // 3: accountservice.PostLoginRequest.PropertiesEntry.value:type_name -> google.protobuf.Any
	1, // 4: accountservice.AccountService.Ping:input_type -> accountservice.GetPingRequest
	3, // 5: accountservice.AccountService.Login:input_type -> accountservice.PostLoginRequest
	2, // 6: accountservice.AccountService.Ping:output_type -> accountservice.GetPingResponse
	5, // 7: accountservice.AccountService.Login:output_type -> accountservice.PostLoginResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_accountservice_proto_init() }
func file_accountservice_proto_init() {
	if File_accountservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accountservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPingRequest); i {
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
		file_accountservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPingResponse); i {
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
		file_accountservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostLoginRequest); i {
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
		file_accountservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenResponse); i {
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
		file_accountservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostLoginResponse); i {
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
			RawDescriptor: file_accountservice_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accountservice_proto_goTypes,
		DependencyIndexes: file_accountservice_proto_depIdxs,
		EnumInfos:         file_accountservice_proto_enumTypes,
		MessageInfos:      file_accountservice_proto_msgTypes,
	}.Build()
	File_accountservice_proto = out.File
	file_accountservice_proto_rawDesc = nil
	file_accountservice_proto_goTypes = nil
	file_accountservice_proto_depIdxs = nil
}
