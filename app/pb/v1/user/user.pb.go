// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/user/user.proto

package user

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateUserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8fd1c8d22ccfa2e, []int{0}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateUserResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8fd1c8d22ccfa2e, []int{1}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateUserRequest)(nil), "user.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "user.CreateUserResponse")
}

func init() { proto.RegisterFile("proto/v1/user/user.proto", fileDescriptor_a8fd1c8d22ccfa2e) }

var fileDescriptor_a8fd1c8d22ccfa2e = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0xe9, 0xdc, 0xe6, 0x16, 0x11, 0x66, 0x14, 0x16, 0x8a, 0x17, 0xa3, 0x20, 0x88, 0xb0,
	0xc6, 0x3f, 0x77, 0x5e, 0x49, 0x7d, 0x83, 0xaa, 0x37, 0xe2, 0x4d, 0xda, 0x1d, 0xba, 0x40, 0xdb,
	0x13, 0x73, 0xd2, 0x7a, 0xef, 0x2b, 0xf8, 0x68, 0xbe, 0x82, 0x4f, 0xb1, 0x0b, 0x91, 0xa5, 0x54,
	0x85, 0xdd, 0x84, 0x2f, 0xf9, 0x7d, 0xf9, 0x38, 0xe7, 0x63, 0xc2, 0x58, 0x74, 0x28, 0xdb, 0x2b,
	0xd9, 0x10, 0x58, 0x7f, 0xc4, 0xfe, 0x89, 0x0f, 0xb7, 0x3a, 0x3c, 0x2d, 0x10, 0x8b, 0x12, 0xa4,
	0x32, 0x5a, 0xaa, 0xba, 0x46, 0xa7, 0x9c, 0xc6, 0x9a, 0x3a, 0x4f, 0x38, 0x6f, 0x55, 0xa9, 0x57,
	0xca, 0x81, 0xec, 0x45, 0x07, 0xa2, 0x17, 0x76, 0x74, 0x6f, 0x41, 0x39, 0x78, 0x22, 0xb0, 0x29,
	0xbc, 0x36, 0x40, 0x8e, 0x2f, 0xd8, 0x08, 0x2a, 0xa5, 0x4b, 0x11, 0x2c, 0x82, 0xf3, 0x69, 0xc2,
	0x36, 0xc9, 0xbe, 0x1d, 0xcd, 0xf6, 0xc4, 0x77, 0x90, 0x76, 0x80, 0x9f, 0xb1, 0x89, 0x51, 0x44,
	0x6f, 0x68, 0x57, 0x62, 0xe0, 0x4d, 0xd3, 0x4d, 0x32, 0xb6, 0xc3, 0xd9, 0x44, 0xdc, 0xa5, 0xbf,
	0x28, 0x3a, 0x61, 0xfc, 0x7f, 0x3a, 0x19, 0xac, 0x09, 0xae, 0x73, 0x76, 0xb0, 0xbd, 0x3f, 0x80,
	0x6d, 0x75, 0x0e, 0xfc, 0x91, 0xb1, 0x3f, 0x13, 0x9f, 0xc7, 0x7e, 0xb5, 0x9d, 0xa1, 0x42, 0xb1,
	0x0b, 0xba, 0xbc, 0xe8, 0xf8, 0xfd, 0xf3, 0xeb, 0x63, 0x70, 0x18, 0x4d, 0xfa, 0x76, 0x6e, 0x83,
	0x8b, 0xe4, 0xf2, 0x39, 0x2e, 0xb4, 0x5b, 0x37, 0x59, 0x9c, 0x63, 0x25, 0x35, 0x39, 0x5a, 0xcb,
	0x02, 0x97, 0x85, 0x35, 0xf9, 0x92, 0x54, 0x65, 0x7c, 0x55, 0x46, 0x9a, 0xac, 0xff, 0x94, 0x8d,
	0x7d, 0x23, 0x37, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x55, 0xce, 0x42, 0x90, 0x6a, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/user/user.proto",
}