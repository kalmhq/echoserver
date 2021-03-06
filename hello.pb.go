// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package main

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GreetingMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetingMessage) Reset()         { *m = GreetingMessage{} }
func (m *GreetingMessage) String() string { return proto.CompactTextString(m) }
func (*GreetingMessage) ProtoMessage()    {}
func (*GreetingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *GreetingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetingMessage.Unmarshal(m, b)
}
func (m *GreetingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetingMessage.Marshal(b, m, deterministic)
}
func (m *GreetingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingMessage.Merge(m, src)
}
func (m *GreetingMessage) XXX_Size() int {
	return xxx_messageInfo_GreetingMessage.Size(m)
}
func (m *GreetingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingMessage proto.InternalMessageInfo

type GreetingReply struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	ClientAddress        string   `protobuf:"bytes,2,opt,name=client_address,json=clientAddress,proto3" json:"client_address,omitempty"`
	AuthInfo             string   `protobuf:"bytes,3,opt,name=auth_info,json=authInfo,proto3" json:"auth_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetingReply) Reset()         { *m = GreetingReply{} }
func (m *GreetingReply) String() string { return proto.CompactTextString(m) }
func (*GreetingReply) ProtoMessage()    {}
func (*GreetingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *GreetingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetingReply.Unmarshal(m, b)
}
func (m *GreetingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetingReply.Marshal(b, m, deterministic)
}
func (m *GreetingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingReply.Merge(m, src)
}
func (m *GreetingReply) XXX_Size() int {
	return xxx_messageInfo_GreetingReply.Size(m)
}
func (m *GreetingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingReply.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingReply proto.InternalMessageInfo

func (m *GreetingReply) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *GreetingReply) GetClientAddress() string {
	if m != nil {
		return m.ClientAddress
	}
	return ""
}

func (m *GreetingReply) GetAuthInfo() string {
	if m != nil {
		return m.AuthInfo
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetingMessage)(nil), "main.GreetingMessage")
	proto.RegisterType((*GreetingReply)(nil), "main.GreetingReply")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce) }

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x12, 0xe4,
	0xe2, 0x77, 0x2f, 0x4a, 0x4d, 0x2d, 0xc9, 0xcc, 0x4b, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f,
	0x55, 0xca, 0xe7, 0xe2, 0x85, 0x09, 0x05, 0xa5, 0x16, 0xe4, 0x54, 0x0a, 0x49, 0x71, 0x71, 0x64,
	0xe4, 0x17, 0x97, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9,
	0x42, 0xaa, 0x5c, 0x7c, 0xc9, 0x39, 0x99, 0xa9, 0x79, 0x25, 0xf1, 0x89, 0x29, 0x29, 0x45, 0xa9,
	0xc5, 0xc5, 0x12, 0x4c, 0x60, 0x15, 0xbc, 0x10, 0x51, 0x47, 0x88, 0xa0, 0x90, 0x34, 0x17, 0x67,
	0x62, 0x69, 0x49, 0x46, 0x7c, 0x66, 0x5e, 0x5a, 0xbe, 0x04, 0x33, 0xc4, 0x0c, 0x90, 0x80, 0x67,
	0x5e, 0x5a, 0xbe, 0x91, 0x1b, 0x17, 0x97, 0x07, 0xc8, 0x61, 0xe1, 0xf9, 0x45, 0x39, 0x29, 0x42,
	0x16, 0x5c, 0x1c, 0x30, 0xeb, 0x85, 0x44, 0xf5, 0x40, 0x8e, 0xd4, 0x43, 0x73, 0xa1, 0x94, 0x30,
	0xaa, 0x30, 0xd8, 0x95, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x8f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x27, 0xfc, 0x54, 0x28, 0xe7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloWorldClient is the client API for HelloWorld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloWorldClient interface {
	Greeting(ctx context.Context, in *GreetingMessage, opts ...grpc.CallOption) (*GreetingReply, error)
}

type helloWorldClient struct {
	cc *grpc.ClientConn
}

func NewHelloWorldClient(cc *grpc.ClientConn) HelloWorldClient {
	return &helloWorldClient{cc}
}

func (c *helloWorldClient) Greeting(ctx context.Context, in *GreetingMessage, opts ...grpc.CallOption) (*GreetingReply, error) {
	out := new(GreetingReply)
	err := c.cc.Invoke(ctx, "/main.HelloWorld/Greeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloWorldServer is the server API for HelloWorld service.
type HelloWorldServer interface {
	Greeting(context.Context, *GreetingMessage) (*GreetingReply, error)
}

// UnimplementedHelloWorldServer can be embedded to have forward compatible implementations.
type UnimplementedHelloWorldServer struct {
}

func (*UnimplementedHelloWorldServer) Greeting(ctx context.Context, req *GreetingMessage) (*GreetingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greeting not implemented")
}

func RegisterHelloWorldServer(s *grpc.Server, srv HelloWorldServer) {
	s.RegisterService(&_HelloWorld_serviceDesc, srv)
}

func _HelloWorld_Greeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetingMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).Greeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.HelloWorld/Greeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).Greeting(ctx, req.(*GreetingMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloWorld_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.HelloWorld",
	HandlerType: (*HelloWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greeting",
			Handler:    _HelloWorld_Greeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
