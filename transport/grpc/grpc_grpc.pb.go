// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// WashServerServiceClient is the client API for WashServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WashServerServiceClient interface {
	VerifyClient(ctx context.Context, in *Verify, opts ...grpc.CallOption) (*VerifyAnswer, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (WashServerService_SendMessageClient, error)
	SendMessageToOtherClient(ctx context.Context, opts ...grpc.CallOption) (WashServerService_SendMessageToOtherClientClient, error)
}

type washServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWashServerServiceClient(cc grpc.ClientConnInterface) WashServerServiceClient {
	return &washServerServiceClient{cc}
}

func (c *washServerServiceClient) VerifyClient(ctx context.Context, in *Verify, opts ...grpc.CallOption) (*VerifyAnswer, error) {
	out := new(VerifyAnswer)
	err := c.cc.Invoke(ctx, "/xgrpc.WashServerService/VerifyClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *washServerServiceClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (WashServerService_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &WashServerService_ServiceDesc.Streams[0], "/xgrpc.WashServerService/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &washServerServiceSendMessageClient{stream}
	return x, nil
}

type WashServerService_SendMessageClient interface {
	Send(*Message) error
	Recv() (*MessageAnswer, error)
	grpc.ClientStream
}

type washServerServiceSendMessageClient struct {
	grpc.ClientStream
}

func (x *washServerServiceSendMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *washServerServiceSendMessageClient) Recv() (*MessageAnswer, error) {
	m := new(MessageAnswer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *washServerServiceClient) SendMessageToOtherClient(ctx context.Context, opts ...grpc.CallOption) (WashServerService_SendMessageToOtherClientClient, error) {
	stream, err := c.cc.NewStream(ctx, &WashServerService_ServiceDesc.Streams[1], "/xgrpc.WashServerService/SendMessageToOtherClient", opts...)
	if err != nil {
		return nil, err
	}
	x := &washServerServiceSendMessageToOtherClientClient{stream}
	return x, nil
}

type WashServerService_SendMessageToOtherClientClient interface {
	Send(*MessageToOther) error
	Recv() (*MessageToOtherAnswer, error)
	grpc.ClientStream
}

type washServerServiceSendMessageToOtherClientClient struct {
	grpc.ClientStream
}

func (x *washServerServiceSendMessageToOtherClientClient) Send(m *MessageToOther) error {
	return x.ClientStream.SendMsg(m)
}

func (x *washServerServiceSendMessageToOtherClientClient) Recv() (*MessageToOtherAnswer, error) {
	m := new(MessageToOtherAnswer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WashServerServiceServer is the server API for WashServerService service.
// All implementations must embed UnimplementedWashServerServiceServer
// for forward compatibility
type WashServerServiceServer interface {
	VerifyClient(context.Context, *Verify) (*VerifyAnswer, error)
	SendMessage(WashServerService_SendMessageServer) error
	SendMessageToOtherClient(WashServerService_SendMessageToOtherClientServer) error
	mustEmbedUnimplementedWashServerServiceServer()
}

// UnimplementedWashServerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWashServerServiceServer struct {
}

func (UnimplementedWashServerServiceServer) VerifyClient(context.Context, *Verify) (*VerifyAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyClient not implemented")
}
func (UnimplementedWashServerServiceServer) SendMessage(WashServerService_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedWashServerServiceServer) SendMessageToOtherClient(WashServerService_SendMessageToOtherClientServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessageToOtherClient not implemented")
}
func (UnimplementedWashServerServiceServer) mustEmbedUnimplementedWashServerServiceServer() {}

// UnsafeWashServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WashServerServiceServer will
// result in compilation errors.
type UnsafeWashServerServiceServer interface {
	mustEmbedUnimplementedWashServerServiceServer()
}

func RegisterWashServerServiceServer(s grpc.ServiceRegistrar, srv WashServerServiceServer) {
	s.RegisterService(&WashServerService_ServiceDesc, srv)
}

func _WashServerService_VerifyClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Verify)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WashServerServiceServer).VerifyClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xgrpc.WashServerService/VerifyClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WashServerServiceServer).VerifyClient(ctx, req.(*Verify))
	}
	return interceptor(ctx, in, info, handler)
}

func _WashServerService_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WashServerServiceServer).SendMessage(&washServerServiceSendMessageServer{stream})
}

type WashServerService_SendMessageServer interface {
	Send(*MessageAnswer) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type washServerServiceSendMessageServer struct {
	grpc.ServerStream
}

func (x *washServerServiceSendMessageServer) Send(m *MessageAnswer) error {
	return x.ServerStream.SendMsg(m)
}

func (x *washServerServiceSendMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _WashServerService_SendMessageToOtherClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WashServerServiceServer).SendMessageToOtherClient(&washServerServiceSendMessageToOtherClientServer{stream})
}

type WashServerService_SendMessageToOtherClientServer interface {
	Send(*MessageToOtherAnswer) error
	Recv() (*MessageToOther, error)
	grpc.ServerStream
}

type washServerServiceSendMessageToOtherClientServer struct {
	grpc.ServerStream
}

func (x *washServerServiceSendMessageToOtherClientServer) Send(m *MessageToOtherAnswer) error {
	return x.ServerStream.SendMsg(m)
}

func (x *washServerServiceSendMessageToOtherClientServer) Recv() (*MessageToOther, error) {
	m := new(MessageToOther)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WashServerService_ServiceDesc is the grpc.ServiceDesc for WashServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WashServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xgrpc.WashServerService",
	HandlerType: (*WashServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyClient",
			Handler:    _WashServerService_VerifyClient_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessage",
			Handler:       _WashServerService_SendMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SendMessageToOtherClient",
			Handler:       _WashServerService_SendMessageToOtherClient_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc.proto",
}