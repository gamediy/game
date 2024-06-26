// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: user/user/user.proto

package user

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
	UserService_Reg_FullMethodName             = "/user.UserService/Reg"
	UserService_Login_FullMethodName           = "/user.UserService/Login"
	UserService_UserInfo_FullMethodName        = "/user.UserService/UserInfo"
	UserService_Wallet_FullMethodName          = "/user.UserService/Wallet"
	UserService_UpdateLoginPass_FullMethodName = "/user.UserService/UpdateLoginPass"
	UserService_SendMsgCode_FullMethodName     = "/user.UserService/SendMsgCode"
	UserService_BindPhone_FullMethodName       = "/user.UserService/BindPhone"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Reg(ctx context.Context, in *RegRequest, opts ...grpc.CallOption) (*RegReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error)
	Wallet(ctx context.Context, in *WalletRequest, opts ...grpc.CallOption) (*WalletReply, error)
	UpdateLoginPass(ctx context.Context, in *UpdateLoginPassReq, opts ...grpc.CallOption) (*UpdateLoginPassRes, error)
	SendMsgCode(ctx context.Context, in *SendMsgCodeReq, opts ...grpc.CallOption) (*SendMsgCodeRes, error)
	BindPhone(ctx context.Context, in *BindPhoneReq, opts ...grpc.CallOption) (*BindPhoneRes, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Reg(ctx context.Context, in *RegRequest, opts ...grpc.CallOption) (*RegReply, error) {
	out := new(RegReply)
	err := c.cc.Invoke(ctx, UserService_Reg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, UserService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, UserService_UserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Wallet(ctx context.Context, in *WalletRequest, opts ...grpc.CallOption) (*WalletReply, error) {
	out := new(WalletReply)
	err := c.cc.Invoke(ctx, UserService_Wallet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateLoginPass(ctx context.Context, in *UpdateLoginPassReq, opts ...grpc.CallOption) (*UpdateLoginPassRes, error) {
	out := new(UpdateLoginPassRes)
	err := c.cc.Invoke(ctx, UserService_UpdateLoginPass_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SendMsgCode(ctx context.Context, in *SendMsgCodeReq, opts ...grpc.CallOption) (*SendMsgCodeRes, error) {
	out := new(SendMsgCodeRes)
	err := c.cc.Invoke(ctx, UserService_SendMsgCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) BindPhone(ctx context.Context, in *BindPhoneReq, opts ...grpc.CallOption) (*BindPhoneRes, error) {
	out := new(BindPhoneRes)
	err := c.cc.Invoke(ctx, UserService_BindPhone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Reg(context.Context, *RegRequest) (*RegReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoReply, error)
	Wallet(context.Context, *WalletRequest) (*WalletReply, error)
	UpdateLoginPass(context.Context, *UpdateLoginPassReq) (*UpdateLoginPassRes, error)
	SendMsgCode(context.Context, *SendMsgCodeReq) (*SendMsgCodeRes, error)
	BindPhone(context.Context, *BindPhoneReq) (*BindPhoneRes, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Reg(context.Context, *RegRequest) (*RegReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reg not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) UserInfo(context.Context, *UserInfoRequest) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedUserServiceServer) Wallet(context.Context, *WalletRequest) (*WalletReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Wallet not implemented")
}
func (UnimplementedUserServiceServer) UpdateLoginPass(context.Context, *UpdateLoginPassReq) (*UpdateLoginPassRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLoginPass not implemented")
}
func (UnimplementedUserServiceServer) SendMsgCode(context.Context, *SendMsgCodeReq) (*SendMsgCodeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsgCode not implemented")
}
func (UnimplementedUserServiceServer) BindPhone(context.Context, *BindPhoneReq) (*BindPhoneRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindPhone not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Reg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Reg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Reg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Reg(ctx, req.(*RegRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Wallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Wallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Wallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Wallet(ctx, req.(*WalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateLoginPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLoginPassReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateLoginPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateLoginPass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateLoginPass(ctx, req.(*UpdateLoginPassReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SendMsgCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsgCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SendMsgCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SendMsgCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SendMsgCode(ctx, req.(*SendMsgCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_BindPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindPhoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).BindPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_BindPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).BindPhone(ctx, req.(*BindPhoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reg",
			Handler:    _UserService_Reg_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _UserService_UserInfo_Handler,
		},
		{
			MethodName: "Wallet",
			Handler:    _UserService_Wallet_Handler,
		},
		{
			MethodName: "UpdateLoginPass",
			Handler:    _UserService_UpdateLoginPass_Handler,
		},
		{
			MethodName: "SendMsgCode",
			Handler:    _UserService_SendMsgCode_Handler,
		},
		{
			MethodName: "BindPhone",
			Handler:    _UserService_BindPhone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user/user.proto",
}
