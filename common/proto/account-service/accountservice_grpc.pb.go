// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: accountservice.proto

package account_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AccountService_Ping_FullMethodName  = "/accountservice.AccountService/Ping"
	AccountService_Login_FullMethodName = "/accountservice.AccountService/Login"
)

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	Ping(ctx context.Context, in *GetPingRequest, opts ...grpc.CallOption) (*GetPingResponse, error)
	Login(ctx context.Context, in *PostLoginRequest, opts ...grpc.CallOption) (*PostLoginResponse, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) Ping(ctx context.Context, in *GetPingRequest, opts ...grpc.CallOption) (*GetPingResponse, error) {
	out := new(GetPingResponse)
	err := c.cc.Invoke(ctx, AccountService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Login(ctx context.Context, in *PostLoginRequest, opts ...grpc.CallOption) (*PostLoginResponse, error) {
	out := new(PostLoginResponse)
	err := c.cc.Invoke(ctx, AccountService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations should embed UnimplementedAccountServiceServer
// for forward compatibility
type AccountServiceServer interface {
	Ping(context.Context, *GetPingRequest) (*GetPingResponse, error)
	Login(context.Context, *PostLoginRequest) (*PostLoginResponse, error)
}

// UnimplementedAccountServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) Ping(context.Context, *GetPingRequest) (*GetPingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAccountServiceServer) Login(context.Context, *PostLoginRequest) (*PostLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Ping(ctx, req.(*GetPingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Login(ctx, req.(*PostLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accountservice.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _AccountService_Ping_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AccountService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accountservice.proto",
}
