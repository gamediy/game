// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: user/wallet/wallet.proto

package wallet

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
	WalletService_ListChangeLog_FullMethodName = "/wallet.WalletService/ListChangeLog"
	WalletService_ListTransType_FullMethodName = "/wallet.WalletService/ListTransType"
)

// WalletServiceClient is the client API for WalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletServiceClient interface {
	ListChangeLog(ctx context.Context, in *ListChangeLogReq, opts ...grpc.CallOption) (*ListChangeLogRes, error)
	ListTransType(ctx context.Context, in *ListTransTypeReq, opts ...grpc.CallOption) (*ListTransTypeRes, error)
}

type walletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
	return &walletServiceClient{cc}
}

func (c *walletServiceClient) ListChangeLog(ctx context.Context, in *ListChangeLogReq, opts ...grpc.CallOption) (*ListChangeLogRes, error) {
	out := new(ListChangeLogRes)
	err := c.cc.Invoke(ctx, WalletService_ListChangeLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) ListTransType(ctx context.Context, in *ListTransTypeReq, opts ...grpc.CallOption) (*ListTransTypeRes, error) {
	out := new(ListTransTypeRes)
	err := c.cc.Invoke(ctx, WalletService_ListTransType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServiceServer is the server API for WalletService service.
// All implementations must embed UnimplementedWalletServiceServer
// for forward compatibility
type WalletServiceServer interface {
	ListChangeLog(context.Context, *ListChangeLogReq) (*ListChangeLogRes, error)
	ListTransType(context.Context, *ListTransTypeReq) (*ListTransTypeRes, error)
	mustEmbedUnimplementedWalletServiceServer()
}

// UnimplementedWalletServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServiceServer struct {
}

func (UnimplementedWalletServiceServer) ListChangeLog(context.Context, *ListChangeLogReq) (*ListChangeLogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChangeLog not implemented")
}
func (UnimplementedWalletServiceServer) ListTransType(context.Context, *ListTransTypeReq) (*ListTransTypeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransType not implemented")
}
func (UnimplementedWalletServiceServer) mustEmbedUnimplementedWalletServiceServer() {}

// UnsafeWalletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServiceServer will
// result in compilation errors.
type UnsafeWalletServiceServer interface {
	mustEmbedUnimplementedWalletServiceServer()
}

func RegisterWalletServiceServer(s grpc.ServiceRegistrar, srv WalletServiceServer) {
	s.RegisterService(&WalletService_ServiceDesc, srv)
}

func _WalletService_ListChangeLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChangeLogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).ListChangeLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_ListChangeLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).ListChangeLog(ctx, req.(*ListChangeLogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_ListTransType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransTypeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).ListTransType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_ListTransType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).ListTransType(ctx, req.(*ListTransTypeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletService_ServiceDesc is the grpc.ServiceDesc for WalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wallet.WalletService",
	HandlerType: (*WalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListChangeLog",
			Handler:    _WalletService_ListChangeLog_Handler,
		},
		{
			MethodName: "ListTransType",
			Handler:    _WalletService_ListTransType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/wallet/wallet.proto",
}
