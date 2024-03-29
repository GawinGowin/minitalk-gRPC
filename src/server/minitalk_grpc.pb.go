// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: minitalk.proto

package main

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

// MinitalkClient is the client API for Minitalk service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MinitalkClient interface {
	SendMsg(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgReply, error)
}

type minitalkClient struct {
	cc grpc.ClientConnInterface
}

func NewMinitalkClient(cc grpc.ClientConnInterface) MinitalkClient {
	return &minitalkClient{cc}
}

func (c *minitalkClient) SendMsg(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgReply, error) {
	out := new(MsgReply)
	err := c.cc.Invoke(ctx, "/minitalk.Minitalk/SendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MinitalkServer is the server API for Minitalk service.
// All implementations must embed UnimplementedMinitalkServer
// for forward compatibility
type MinitalkServer interface {
	SendMsg(context.Context, *MsgRequest) (*MsgReply, error)
	mustEmbedUnimplementedMinitalkServer()
}

// UnimplementedMinitalkServer must be embedded to have forward compatible implementations.
type UnimplementedMinitalkServer struct {
}

func (UnimplementedMinitalkServer) SendMsg(context.Context, *MsgRequest) (*MsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}
func (UnimplementedMinitalkServer) mustEmbedUnimplementedMinitalkServer() {}

// UnsafeMinitalkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MinitalkServer will
// result in compilation errors.
type UnsafeMinitalkServer interface {
	mustEmbedUnimplementedMinitalkServer()
}

func RegisterMinitalkServer(s grpc.ServiceRegistrar, srv MinitalkServer) {
	s.RegisterService(&Minitalk_ServiceDesc, srv)
}

func _Minitalk_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinitalkServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/minitalk.Minitalk/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinitalkServer).SendMsg(ctx, req.(*MsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Minitalk_ServiceDesc is the grpc.ServiceDesc for Minitalk service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Minitalk_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "minitalk.Minitalk",
	HandlerType: (*MinitalkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMsg",
			Handler:    _Minitalk_SendMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "minitalk.proto",
}
