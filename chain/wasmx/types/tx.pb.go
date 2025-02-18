// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kaiju/wasmx/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("kaiju/wasmx/v1/tx.proto", fileDescriptor_f7afe23baa925f70) }

var fileDescriptor_f7afe23baa925f70 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0xcc, 0xcb, 0x4a,
	0x4d, 0x2e, 0xc9, 0x2c, 0x4b, 0xd5, 0x2f, 0x4f, 0x2c, 0xce, 0xad, 0xd0, 0x2f, 0x33, 0xd4, 0x2f,
	0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x82, 0x4b, 0xea, 0x81, 0x25, 0xf5, 0xca,
	0x0c, 0x8d, 0x58, 0xb9, 0x98, 0x7d, 0x8b, 0xd3, 0x9d, 0x52, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0,
	0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8,
	0xf1, 0x58, 0x8e, 0x21, 0xca, 0x3b, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57,
	0xdf, 0x13, 0xa6, 0xdf, 0x27, 0x31, 0xa9, 0x58, 0x1f, 0x6e, 0x9a, 0x6e, 0x72, 0x7e, 0x51, 0x2a,
	0x32, 0x37, 0x23, 0x31, 0x33, 0x4f, 0x3f, 0x37, 0x3f, 0xa5, 0x34, 0x27, 0xb5, 0x18, 0xea, 0x8e,
	0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x43, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x76, 0x83, 0xd0, 0x89, 0xa7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kaiju.wasmx.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "kaiju/wasmx/v1/tx.proto",
}
