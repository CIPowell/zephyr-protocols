// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	auth.proto

It has these top-level messages:
	LoginRequest
	LoginResult
*/
package auth

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

type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Ip       string `protobuf:"bytes,3,opt,name=ip" json:"ip,omitempty"`
	Device   string `protobuf:"bytes,4,opt,name=device" json:"device,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *LoginRequest) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

type LoginResult struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Token   string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *LoginResult) Reset()                    { *m = LoginResult{} }
func (m *LoginResult) String() string            { return proto.CompactTextString(m) }
func (*LoginResult) ProtoMessage()               {}
func (*LoginResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LoginResult) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "auth.LoginRequest")
	proto.RegisterType((*LoginResult)(nil), "auth.LoginResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auth service

type AuthClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResult, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResult, error) {
	out := new(LoginResult)
	err := grpc.Invoke(ctx, "/auth.Auth/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	Login(context.Context, *LoginRequest) (*LoginResult, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x40, 0x6d, 0x6d, 0x6b, 0x1d, 0x45, 0x70, 0x10, 0x09, 0x3d, 0x49, 0x4f, 0x9e, 0x8a, 0xe8,
	0xc5, 0xab, 0x77, 0x4f, 0xf5, 0x0b, 0x6a, 0x3b, 0xd8, 0xe0, 0x6e, 0x92, 0xcd, 0x24, 0xbb, 0xbf,
	0xbf, 0x24, 0x4d, 0x97, 0xdd, 0x5b, 0xde, 0x1b, 0xc8, 0x63, 0x06, 0x60, 0xf0, 0x6e, 0xee, 0x8c,
	0xd5, 0x4e, 0x63, 0x11, 0xde, 0xad, 0x82, 0xfb, 0x6f, 0xfd, 0x27, 0x55, 0x4f, 0x3b, 0x4f, 0xec,
	0xb0, 0x81, 0xda, 0x33, 0x59, 0x35, 0x6c, 0x49, 0x64, 0x2f, 0xd9, 0xeb, 0x6d, 0x7f, 0xe2, 0x30,
	0x33, 0x03, 0xf3, 0x41, 0xdb, 0x49, 0xe4, 0xcb, 0x6c, 0x65, 0x7c, 0x80, 0x5c, 0x1a, 0x71, 0x1d,
	0x6d, 0x2e, 0x0d, 0x3e, 0x43, 0x35, 0xd1, 0x5e, 0x8e, 0x24, 0x8a, 0xe8, 0x12, 0xb5, 0x3f, 0x70,
	0x97, 0x7a, 0xec, 0x37, 0x0e, 0x05, 0xdc, 0xb0, 0x1f, 0x47, 0x62, 0x8e, 0xb5, 0xba, 0x5f, 0x11,
	0x9f, 0xa0, 0x74, 0xfa, 0x9f, 0x54, 0x2a, 0x2d, 0x10, 0x2c, 0x59, 0xab, 0x6d, 0x2a, 0x2d, 0xf0,
	0xfe, 0x09, 0xc5, 0x97, 0x77, 0x33, 0xbe, 0x41, 0x19, 0x3f, 0x47, 0xec, 0xe2, 0xa2, 0xe7, 0x9b,
	0x35, 0x8f, 0x17, 0x2e, 0xd4, 0xdb, 0xab, 0xdf, 0x2a, 0xde, 0xe2, 0xe3, 0x18, 0x00, 0x00, 0xff,
	0xff, 0xe9, 0x7a, 0xe3, 0xc6, 0x19, 0x01, 0x00, 0x00,
}
