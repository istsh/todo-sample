// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/proto/v1/user/user.proto

package user

import (
	context "context"
	fmt "fmt"
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
	return fileDescriptor_f9205c859dcaa697, []int{0}
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
	return fileDescriptor_f9205c859dcaa697, []int{1}
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

type ChangeEmailRequest struct {
	NewEmail             string   `protobuf:"bytes,1,opt,name=new_email,json=newEmail,proto3" json:"new_email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeEmailRequest) Reset()         { *m = ChangeEmailRequest{} }
func (m *ChangeEmailRequest) String() string { return proto.CompactTextString(m) }
func (*ChangeEmailRequest) ProtoMessage()    {}
func (*ChangeEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9205c859dcaa697, []int{2}
}

func (m *ChangeEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeEmailRequest.Unmarshal(m, b)
}
func (m *ChangeEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeEmailRequest.Marshal(b, m, deterministic)
}
func (m *ChangeEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeEmailRequest.Merge(m, src)
}
func (m *ChangeEmailRequest) XXX_Size() int {
	return xxx_messageInfo_ChangeEmailRequest.Size(m)
}
func (m *ChangeEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeEmailRequest proto.InternalMessageInfo

func (m *ChangeEmailRequest) GetNewEmail() string {
	if m != nil {
		return m.NewEmail
	}
	return ""
}

type ChangeEmailResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeEmailResponse) Reset()         { *m = ChangeEmailResponse{} }
func (m *ChangeEmailResponse) String() string { return proto.CompactTextString(m) }
func (*ChangeEmailResponse) ProtoMessage()    {}
func (*ChangeEmailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9205c859dcaa697, []int{3}
}

func (m *ChangeEmailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeEmailResponse.Unmarshal(m, b)
}
func (m *ChangeEmailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeEmailResponse.Marshal(b, m, deterministic)
}
func (m *ChangeEmailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeEmailResponse.Merge(m, src)
}
func (m *ChangeEmailResponse) XXX_Size() int {
	return xxx_messageInfo_ChangeEmailResponse.Size(m)
}
func (m *ChangeEmailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeEmailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeEmailResponse proto.InternalMessageInfo

type ChangePasswordRequest struct {
	CurrentPassword      string   `protobuf:"bytes,1,opt,name=current_password,json=currentPassword,proto3" json:"current_password,omitempty"`
	NewPassword          string   `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangePasswordRequest) Reset()         { *m = ChangePasswordRequest{} }
func (m *ChangePasswordRequest) String() string { return proto.CompactTextString(m) }
func (*ChangePasswordRequest) ProtoMessage()    {}
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9205c859dcaa697, []int{4}
}

func (m *ChangePasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangePasswordRequest.Unmarshal(m, b)
}
func (m *ChangePasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangePasswordRequest.Marshal(b, m, deterministic)
}
func (m *ChangePasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangePasswordRequest.Merge(m, src)
}
func (m *ChangePasswordRequest) XXX_Size() int {
	return xxx_messageInfo_ChangePasswordRequest.Size(m)
}
func (m *ChangePasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangePasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangePasswordRequest proto.InternalMessageInfo

func (m *ChangePasswordRequest) GetCurrentPassword() string {
	if m != nil {
		return m.CurrentPassword
	}
	return ""
}

func (m *ChangePasswordRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type ChangePasswordResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangePasswordResponse) Reset()         { *m = ChangePasswordResponse{} }
func (m *ChangePasswordResponse) String() string { return proto.CompactTextString(m) }
func (*ChangePasswordResponse) ProtoMessage()    {}
func (*ChangePasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9205c859dcaa697, []int{5}
}

func (m *ChangePasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangePasswordResponse.Unmarshal(m, b)
}
func (m *ChangePasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangePasswordResponse.Marshal(b, m, deterministic)
}
func (m *ChangePasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangePasswordResponse.Merge(m, src)
}
func (m *ChangePasswordResponse) XXX_Size() int {
	return xxx_messageInfo_ChangePasswordResponse.Size(m)
}
func (m *ChangePasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangePasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChangePasswordResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateUserRequest)(nil), "user.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "user.CreateUserResponse")
	proto.RegisterType((*ChangeEmailRequest)(nil), "user.ChangeEmailRequest")
	proto.RegisterType((*ChangeEmailResponse)(nil), "user.ChangeEmailResponse")
	proto.RegisterType((*ChangePasswordRequest)(nil), "user.ChangePasswordRequest")
	proto.RegisterType((*ChangePasswordResponse)(nil), "user.ChangePasswordResponse")
}

func init() { proto.RegisterFile("app/proto/v1/user/user.proto", fileDescriptor_f9205c859dcaa697) }

var fileDescriptor_f9205c859dcaa697 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x4f, 0xc2, 0x30,
	0x18, 0xc6, 0x03, 0x51, 0x03, 0x2f, 0x46, 0xb4, 0x82, 0xce, 0x82, 0x09, 0xce, 0x8b, 0x7a, 0x60,
	0x41, 0x6f, 0x5e, 0x09, 0x77, 0x83, 0xf1, 0xe4, 0x81, 0xd4, 0xf1, 0x06, 0x97, 0x60, 0x5b, 0xdb,
	0xc2, 0xee, 0x7e, 0x05, 0xbf, 0x91, 0x5f, 0xc1, 0xaf, 0xe0, 0x07, 0x31, 0xeb, 0xda, 0x31, 0xfe,
	0x5c, 0x96, 0xb5, 0xef, 0xd3, 0xdf, 0xd3, 0xe7, 0xd9, 0xa0, 0xcb, 0xa4, 0x8c, 0xa4, 0x12, 0x46,
	0x44, 0xcb, 0x41, 0xb4, 0xd0, 0xa8, 0xec, 0xa3, 0x6f, 0xb7, 0xc8, 0x5e, 0xf6, 0x4e, 0xbb, 0x33,
	0x21, 0x66, 0x73, 0x8c, 0x98, 0x4c, 0x22, 0xc6, 0xb9, 0x30, 0xcc, 0x24, 0x82, 0xeb, 0x5c, 0x13,
	0x8e, 0xe0, 0x64, 0xa8, 0x90, 0x19, 0x7c, 0xd1, 0xa8, 0xc6, 0xf8, 0xb9, 0x40, 0x6d, 0x48, 0x0b,
	0xf6, 0xf1, 0x83, 0x25, 0xf3, 0xa0, 0xd2, 0xab, 0xdc, 0xd4, 0xc7, 0xf9, 0x82, 0x50, 0xa8, 0x49,
	0xa6, 0x75, 0x2a, 0xd4, 0x34, 0xa8, 0xda, 0x41, 0xb1, 0x0e, 0x5b, 0x40, 0xca, 0x18, 0x2d, 0x05,
	0xd7, 0x18, 0x0e, 0x80, 0x0c, 0xdf, 0x19, 0x9f, 0xe1, 0x28, 0x03, 0x78, 0x7a, 0x07, 0xea, 0x1c,
	0xd3, 0x49, 0xd9, 0xa1, 0xc6, 0x31, 0xb5, 0x9a, 0xb0, 0x0d, 0xa7, 0x6b, 0x47, 0x1c, 0x09, 0xa1,
	0x9d, 0x6f, 0x3f, 0x39, 0x47, 0x0f, 0xbb, 0x85, 0xe3, 0x78, 0xa1, 0x14, 0x72, 0x33, 0x29, 0x2e,
	0x97, 0x33, 0x9b, 0x6e, 0xdf, 0x9f, 0x20, 0x57, 0x70, 0x98, 0xf9, 0x6e, 0x64, 0x68, 0x70, 0x4c,
	0xbd, 0x24, 0x0c, 0xe0, 0x6c, 0xd3, 0x26, 0xbf, 0xc0, 0xfd, 0x4f, 0x15, 0x1a, 0x59, 0xb6, 0x67,
	0x54, 0xcb, 0x24, 0x46, 0xf2, 0x0a, 0xb0, 0x0a, 0x4c, 0xce, 0xfb, 0xb6, 0xf6, 0xad, 0x26, 0x69,
	0xb0, 0x3d, 0x70, 0x89, 0xe8, 0xd7, 0xef, 0xdf, 0x77, 0xb5, 0x15, 0x36, 0x8b, 0x2f, 0x17, 0x5b,
	0xd1, 0x63, 0xe5, 0x8e, 0x4c, 0xa1, 0x51, 0x2a, 0x81, 0x78, 0xc8, 0x56, 0x95, 0xf4, 0x62, 0xc7,
	0xc4, 0xf1, 0x7b, 0x96, 0x4f, 0x69, 0xbb, 0xe0, 0xdb, 0xc2, 0xa3, 0xd8, 0x6a, 0x33, 0x17, 0x09,
	0x47, 0xeb, 0x61, 0x49, 0xa7, 0x8c, 0xdb, 0x68, 0x9a, 0x76, 0x77, 0x0f, 0x9d, 0xdd, 0xb5, 0xb5,
	0xbb, 0xa4, 0x41, 0x61, 0xe7, 0x7b, 0x5e, 0x39, 0xbe, 0x1d, 0xd8, 0x7f, 0xee, 0xe1, 0x3f, 0x00,
	0x00, 0xff, 0xff, 0xfd, 0xfd, 0x4a, 0x81, 0xb7, 0x02, 0x00, 0x00,
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
	ChangeEmail(ctx context.Context, in *ChangeEmailRequest, opts ...grpc.CallOption) (*ChangeEmailResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
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

func (c *userServiceClient) ChangeEmail(ctx context.Context, in *ChangeEmailRequest, opts ...grpc.CallOption) (*ChangeEmailResponse, error) {
	out := new(ChangeEmailResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/ChangeEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	ChangeEmail(context.Context, *ChangeEmailRequest) (*ChangeEmailResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServiceServer) ChangeEmail(ctx context.Context, req *ChangeEmailRequest) (*ChangeEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeEmail not implemented")
}
func (*UnimplementedUserServiceServer) ChangePassword(ctx context.Context, req *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
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

func _UserService_ChangeEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ChangeEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeEmail(ctx, req.(*ChangeEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
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
		{
			MethodName: "ChangeEmail",
			Handler:    _UserService_ChangeEmail_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UserService_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/proto/v1/user/user.proto",
}