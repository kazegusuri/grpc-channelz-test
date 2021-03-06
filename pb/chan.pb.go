// Code generated by protoc-gen-go.
// source: pb/chan.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/chan.proto

It has these top-level messages:
	EchoMessage
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EchoMessage struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *EchoMessage) Reset()                    { *m = EchoMessage{} }
func (m *EchoMessage) String() string            { return proto.CompactTextString(m) }
func (*EchoMessage) ProtoMessage()               {}
func (*EchoMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EchoMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoMessage)(nil), "kazegusuri.test.EchoMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Echo service

type EchoClient interface {
	Echo(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Echo(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error) {
	out := new(EchoMessage)
	err := grpc.Invoke(ctx, "/kazegusuri.test.Echo/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Echo service

type EchoServer interface {
	Echo(context.Context, *EchoMessage) (*EchoMessage, error)
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kazegusuri.test.Echo/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Echo(ctx, req.(*EchoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kazegusuri.test.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Echo_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/chan.proto",
}

func init() { proto.RegisterFile("pb/chan.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 120 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x48, 0xd2, 0x4f,
	0xce, 0x48, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcf, 0x4e, 0xac, 0x4a, 0x4d,
	0x2f, 0x2d, 0x2e, 0x2d, 0xca, 0xd4, 0x2b, 0x49, 0x2d, 0x2e, 0x51, 0x52, 0xe7, 0xe2, 0x76, 0x4d,
	0xce, 0xc8, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x92, 0xe0, 0x62, 0xcf, 0x85, 0x30,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x23, 0x1f, 0x2e, 0x16, 0x90, 0x42, 0x21,
	0x17, 0x28, 0x2d, 0xa3, 0x87, 0x66, 0x94, 0x1e, 0x92, 0x39, 0x52, 0x78, 0x65, 0x95, 0x18, 0x9c,
	0x58, 0xa2, 0x98, 0x0a, 0x92, 0x92, 0xd8, 0xc0, 0x8e, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xf5, 0x33, 0xb7, 0xa8, 0xa5, 0x00, 0x00, 0x00,
}
