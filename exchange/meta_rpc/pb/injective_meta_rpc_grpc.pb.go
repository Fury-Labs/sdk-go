// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: kaiju_meta_rpc.proto

package kaiju_meta_rpcpb

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

// KaijuMetaRPCClient is the client API for KaijuMetaRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KaijuMetaRPCClient interface {
	// Endpoint for checking server health.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// Returns kaiju-exchange version.
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// Gets connection info
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	// Stream keepalive, if server exits, a shutdown event will be sent over this
	// channel.
	StreamKeepalive(ctx context.Context, in *StreamKeepaliveRequest, opts ...grpc.CallOption) (KaijuMetaRPC_StreamKeepaliveClient, error)
}

type kaijuMetaRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewKaijuMetaRPCClient(cc grpc.ClientConnInterface) KaijuMetaRPCClient {
	return &kaijuMetaRPCClient{cc}
}

func (c *kaijuMetaRPCClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/kaiju_meta_rpc.KaijuMetaRPC/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuMetaRPCClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/kaiju_meta_rpc.KaijuMetaRPC/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuMetaRPCClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/kaiju_meta_rpc.KaijuMetaRPC/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaijuMetaRPCClient) StreamKeepalive(ctx context.Context, in *StreamKeepaliveRequest, opts ...grpc.CallOption) (KaijuMetaRPC_StreamKeepaliveClient, error) {
	stream, err := c.cc.NewStream(ctx, &KaijuMetaRPC_ServiceDesc.Streams[0], "/kaiju_meta_rpc.KaijuMetaRPC/StreamKeepalive", opts...)
	if err != nil {
		return nil, err
	}
	x := &kaijuMetaRPCStreamKeepaliveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KaijuMetaRPC_StreamKeepaliveClient interface {
	Recv() (*StreamKeepaliveResponse, error)
	grpc.ClientStream
}

type kaijuMetaRPCStreamKeepaliveClient struct {
	grpc.ClientStream
}

func (x *kaijuMetaRPCStreamKeepaliveClient) Recv() (*StreamKeepaliveResponse, error) {
	m := new(StreamKeepaliveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KaijuMetaRPCServer is the server API for KaijuMetaRPC service.
// All implementations must embed UnimplementedKaijuMetaRPCServer
// for forward compatibility
type KaijuMetaRPCServer interface {
	// Endpoint for checking server health.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// Returns kaiju-exchange version.
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	// Gets connection info
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	// Stream keepalive, if server exits, a shutdown event will be sent over this
	// channel.
	StreamKeepalive(*StreamKeepaliveRequest, KaijuMetaRPC_StreamKeepaliveServer) error
	mustEmbedUnimplementedKaijuMetaRPCServer()
}

// UnimplementedKaijuMetaRPCServer must be embedded to have forward compatible implementations.
type UnimplementedKaijuMetaRPCServer struct {
}

func (UnimplementedKaijuMetaRPCServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedKaijuMetaRPCServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedKaijuMetaRPCServer) Info(context.Context, *InfoRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedKaijuMetaRPCServer) StreamKeepalive(*StreamKeepaliveRequest, KaijuMetaRPC_StreamKeepaliveServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamKeepalive not implemented")
}
func (UnimplementedKaijuMetaRPCServer) mustEmbedUnimplementedKaijuMetaRPCServer() {}

// UnsafeKaijuMetaRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KaijuMetaRPCServer will
// result in compilation errors.
type UnsafeKaijuMetaRPCServer interface {
	mustEmbedUnimplementedKaijuMetaRPCServer()
}

func RegisterKaijuMetaRPCServer(s grpc.ServiceRegistrar, srv KaijuMetaRPCServer) {
	s.RegisterService(&KaijuMetaRPC_ServiceDesc, srv)
}

func _KaijuMetaRPC_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuMetaRPCServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_meta_rpc.KaijuMetaRPC/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuMetaRPCServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuMetaRPC_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuMetaRPCServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_meta_rpc.KaijuMetaRPC/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuMetaRPCServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuMetaRPC_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaijuMetaRPCServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaiju_meta_rpc.KaijuMetaRPC/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaijuMetaRPCServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaijuMetaRPC_StreamKeepalive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamKeepaliveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KaijuMetaRPCServer).StreamKeepalive(m, &kaijuMetaRPCStreamKeepaliveServer{stream})
}

type KaijuMetaRPC_StreamKeepaliveServer interface {
	Send(*StreamKeepaliveResponse) error
	grpc.ServerStream
}

type kaijuMetaRPCStreamKeepaliveServer struct {
	grpc.ServerStream
}

func (x *kaijuMetaRPCStreamKeepaliveServer) Send(m *StreamKeepaliveResponse) error {
	return x.ServerStream.SendMsg(m)
}

// KaijuMetaRPC_ServiceDesc is the grpc.ServiceDesc for KaijuMetaRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KaijuMetaRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kaiju_meta_rpc.KaijuMetaRPC",
	HandlerType: (*KaijuMetaRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _KaijuMetaRPC_Ping_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _KaijuMetaRPC_Version_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _KaijuMetaRPC_Info_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamKeepalive",
			Handler:       _KaijuMetaRPC_StreamKeepalive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kaiju_meta_rpc.proto",
}
