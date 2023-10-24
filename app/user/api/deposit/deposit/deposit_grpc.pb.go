// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: deposit/deposit/deposit.proto

package deposit

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
	DepositSvc_ListDepositAmountItems_FullMethodName = "/deposit.DepositSvc/ListDepositAmountItems"
)

// DepositSvcClient is the client API for DepositSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DepositSvcClient interface {
	ListDepositAmountItems(ctx context.Context, in *DepositAmountItemsReq, opts ...grpc.CallOption) (*DepositAmountItemsRes, error)
}

type depositSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewDepositSvcClient(cc grpc.ClientConnInterface) DepositSvcClient {
	return &depositSvcClient{cc}
}

func (c *depositSvcClient) ListDepositAmountItems(ctx context.Context, in *DepositAmountItemsReq, opts ...grpc.CallOption) (*DepositAmountItemsRes, error) {
	out := new(DepositAmountItemsRes)
	err := c.cc.Invoke(ctx, DepositSvc_ListDepositAmountItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DepositSvcServer is the server API for DepositSvc service.
// All implementations must embed UnimplementedDepositSvcServer
// for forward compatibility
type DepositSvcServer interface {
	ListDepositAmountItems(context.Context, *DepositAmountItemsReq) (*DepositAmountItemsRes, error)
	mustEmbedUnimplementedDepositSvcServer()
}

// UnimplementedDepositSvcServer must be embedded to have forward compatible implementations.
type UnimplementedDepositSvcServer struct {
}

func (UnimplementedDepositSvcServer) ListDepositAmountItems(context.Context, *DepositAmountItemsReq) (*DepositAmountItemsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDepositAmountItems not implemented")
}
func (UnimplementedDepositSvcServer) mustEmbedUnimplementedDepositSvcServer() {}

// UnsafeDepositSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DepositSvcServer will
// result in compilation errors.
type UnsafeDepositSvcServer interface {
	mustEmbedUnimplementedDepositSvcServer()
}

func RegisterDepositSvcServer(s grpc.ServiceRegistrar, srv DepositSvcServer) {
	s.RegisterService(&DepositSvc_ServiceDesc, srv)
}

func _DepositSvc_ListDepositAmountItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositAmountItemsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepositSvcServer).ListDepositAmountItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepositSvc_ListDepositAmountItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepositSvcServer).ListDepositAmountItems(ctx, req.(*DepositAmountItemsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DepositSvc_ServiceDesc is the grpc.ServiceDesc for DepositSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DepositSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "deposit.DepositSvc",
	HandlerType: (*DepositSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDepositAmountItems",
			Handler:    _DepositSvc_ListDepositAmountItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deposit/deposit/deposit.proto",
}