// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api_proto/api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api_proto/api.proto

It has these top-level messages:
	Empty
	User
	MulRequest
	GetNameResponse
	SayHelloResponse
	MulResponse
*/
package api

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

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type User struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Age  int64  `protobuf:"varint,2,opt,name=Age" json:"Age,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

type MulRequest struct {
	A int64 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B int64 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
}

func (m *MulRequest) Reset()                    { *m = MulRequest{} }
func (m *MulRequest) String() string            { return proto.CompactTextString(m) }
func (*MulRequest) ProtoMessage()               {}
func (*MulRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MulRequest) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *MulRequest) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

type GetNameResponse struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetNameResponse) Reset()                    { *m = GetNameResponse{} }
func (m *GetNameResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNameResponse) ProtoMessage()               {}
func (*GetNameResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetNameResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SayHelloResponse struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *SayHelloResponse) Reset()                    { *m = SayHelloResponse{} }
func (m *SayHelloResponse) String() string            { return proto.CompactTextString(m) }
func (*SayHelloResponse) ProtoMessage()               {}
func (*SayHelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SayHelloResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type MulResponse struct {
	Result int64 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *MulResponse) Reset()                    { *m = MulResponse{} }
func (m *MulResponse) String() string            { return proto.CompactTextString(m) }
func (*MulResponse) ProtoMessage()               {}
func (*MulResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MulResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "api.Empty")
	proto.RegisterType((*User)(nil), "api.User")
	proto.RegisterType((*MulRequest)(nil), "api.MulRequest")
	proto.RegisterType((*GetNameResponse)(nil), "api.GetNameResponse")
	proto.RegisterType((*SayHelloResponse)(nil), "api.SayHelloResponse")
	proto.RegisterType((*MulResponse)(nil), "api.MulResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Serv service

type ServClient interface {
	GetName(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetNameResponse, error)
	SayHello(ctx context.Context, in *User, opts ...grpc.CallOption) (*SayHelloResponse, error)
	Mul(ctx context.Context, in *MulRequest, opts ...grpc.CallOption) (*MulResponse, error)
}

type servClient struct {
	cc *grpc.ClientConn
}

func NewServClient(cc *grpc.ClientConn) ServClient {
	return &servClient{cc}
}

func (c *servClient) GetName(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetNameResponse, error) {
	out := new(GetNameResponse)
	err := grpc.Invoke(ctx, "/api.Serv/GetName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servClient) SayHello(ctx context.Context, in *User, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	out := new(SayHelloResponse)
	err := grpc.Invoke(ctx, "/api.Serv/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servClient) Mul(ctx context.Context, in *MulRequest, opts ...grpc.CallOption) (*MulResponse, error) {
	out := new(MulResponse)
	err := grpc.Invoke(ctx, "/api.Serv/Mul", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Serv service

type ServServer interface {
	GetName(context.Context, *Empty) (*GetNameResponse, error)
	SayHello(context.Context, *User) (*SayHelloResponse, error)
	Mul(context.Context, *MulRequest) (*MulResponse, error)
}

func RegisterServServer(s *grpc.Server, srv ServServer) {
	s.RegisterService(&_Serv_serviceDesc, srv)
}

func _Serv_GetName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServServer).GetName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Serv/GetName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServServer).GetName(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Serv_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Serv/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServServer).SayHello(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Serv_Mul_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MulRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServServer).Mul(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Serv/Mul",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServServer).Mul(ctx, req.(*MulRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Serv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Serv",
	HandlerType: (*ServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetName",
			Handler:    _Serv_GetName_Handler,
		},
		{
			MethodName: "SayHello",
			Handler:    _Serv_SayHello_Handler,
		},
		{
			MethodName: "Mul",
			Handler:    _Serv_Mul_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_proto/api.proto",
}

func init() { proto.RegisterFile("api_proto/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x89, 0x59, 0x5b, 0xfb, 0x29, 0x74, 0x89, 0x7f, 0x28, 0x7b, 0x2a, 0x81, 0xc2, 0xa2,
	0xa5, 0x82, 0x3e, 0x81, 0x07, 0xd1, 0x4b, 0x3d, 0xa4, 0x78, 0x96, 0x2c, 0x0c, 0xb2, 0x90, 0x76,
	0xe3, 0x26, 0x2b, 0xf4, 0x19, 0x7c, 0x69, 0xc9, 0xb8, 0xb5, 0xa2, 0xde, 0x7e, 0x93, 0xfc, 0x92,
	0xf9, 0x66, 0x70, 0x6a, 0x7d, 0xfd, 0xe2, 0xdb, 0x26, 0x36, 0xd7, 0xd6, 0xd7, 0x0b, 0x26, 0x25,
	0xad, 0xaf, 0xf5, 0x10, 0x87, 0xf7, 0x6b, 0x1f, 0xb7, 0x7a, 0x8e, 0xec, 0x39, 0x50, 0xab, 0x14,
	0xb2, 0x27, 0xbb, 0xa6, 0x89, 0x98, 0x8a, 0x72, 0x64, 0x98, 0x55, 0x0e, 0x79, 0xf7, 0x4a, 0x93,
	0x83, 0xa9, 0x28, 0xa5, 0x49, 0xa8, 0x4b, 0x60, 0xd9, 0x39, 0x43, 0x6f, 0x1d, 0x85, 0xa8, 0x4e,
	0x20, 0x2c, 0x3f, 0x90, 0x46, 0xd8, 0x54, 0x55, 0xbd, 0x2b, 0x2a, 0x3d, 0xc3, 0xf8, 0x81, 0x62,
	0xfa, 0xc6, 0x50, 0xf0, 0xcd, 0x26, 0x50, 0x6a, 0xb1, 0xf9, 0xd1, 0x22, 0xb1, 0xbe, 0x44, 0xbe,
	0xb2, 0xdb, 0x47, 0x72, 0xae, 0xf9, 0xf6, 0x2e, 0x30, 0x68, 0x29, 0x74, 0x2e, 0xf6, 0x66, 0x5f,
	0xe9, 0x19, 0x8e, 0xb9, 0xf9, 0xbf, 0x9a, 0xdc, 0x69, 0x37, 0x1f, 0x02, 0xd9, 0x8a, 0xda, 0x77,
	0x75, 0x85, 0x61, 0x1f, 0x41, 0x61, 0x91, 0xe6, 0xe7, 0x89, 0x8b, 0x33, 0xe6, 0xdf, 0xe1, 0xe6,
	0x38, 0xda, 0x05, 0x51, 0x23, 0x36, 0xd2, 0x5a, 0x8a, 0x73, 0xc6, 0x3f, 0x11, 0x4b, 0xc8, 0x65,
	0xe7, 0xd4, 0x98, 0x6f, 0xf7, 0x1b, 0x29, 0xf2, 0xfd, 0xc1, 0x97, 0x59, 0x0d, 0x78, 0xe9, 0xb7,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xf4, 0x8a, 0x24, 0x8b, 0x01, 0x00, 0x00,
}
