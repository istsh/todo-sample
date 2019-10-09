// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/proto/v1/todo/todo.proto

package todo

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

type ScheduledDate struct {
	Year                 int32    `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month                int32    `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Day                  int32    `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScheduledDate) Reset()         { *m = ScheduledDate{} }
func (m *ScheduledDate) String() string { return proto.CompactTextString(m) }
func (*ScheduledDate) ProtoMessage()    {}
func (*ScheduledDate) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec8192e4acf1e0e, []int{0}
}

func (m *ScheduledDate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduledDate.Unmarshal(m, b)
}
func (m *ScheduledDate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduledDate.Marshal(b, m, deterministic)
}
func (m *ScheduledDate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduledDate.Merge(m, src)
}
func (m *ScheduledDate) XXX_Size() int {
	return xxx_messageInfo_ScheduledDate.Size(m)
}
func (m *ScheduledDate) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduledDate.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduledDate proto.InternalMessageInfo

func (m *ScheduledDate) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *ScheduledDate) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *ScheduledDate) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

type Todo struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string         `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Project              string         `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	Label                string         `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	Description          string         `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	ScheduledDate        *ScheduledDate `protobuf:"bytes,6,opt,name=scheduled_date,json=scheduledDate,proto3" json:"scheduled_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec8192e4acf1e0e, []int{1}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Todo) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Todo) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Todo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Todo) GetScheduledDate() *ScheduledDate {
	if m != nil {
		return m.ScheduledDate
	}
	return nil
}

type CreateTodoRequest struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoRequest) Reset()         { *m = CreateTodoRequest{} }
func (m *CreateTodoRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTodoRequest) ProtoMessage()    {}
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec8192e4acf1e0e, []int{2}
}

func (m *CreateTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoRequest.Unmarshal(m, b)
}
func (m *CreateTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoRequest.Marshal(b, m, deterministic)
}
func (m *CreateTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoRequest.Merge(m, src)
}
func (m *CreateTodoRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTodoRequest.Size(m)
}
func (m *CreateTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoRequest proto.InternalMessageInfo

func (m *CreateTodoRequest) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type CreateTodoResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoResponse) Reset()         { *m = CreateTodoResponse{} }
func (m *CreateTodoResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTodoResponse) ProtoMessage()    {}
func (*CreateTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec8192e4acf1e0e, []int{3}
}

func (m *CreateTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoResponse.Unmarshal(m, b)
}
func (m *CreateTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoResponse.Marshal(b, m, deterministic)
}
func (m *CreateTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoResponse.Merge(m, src)
}
func (m *CreateTodoResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodoResponse.Size(m)
}
func (m *CreateTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoResponse proto.InternalMessageInfo

type UpdateTodoRequest struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodoRequest) Reset()         { *m = UpdateTodoRequest{} }
func (m *UpdateTodoRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoRequest) ProtoMessage()    {}
func (*UpdateTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec8192e4acf1e0e, []int{4}
}

func (m *UpdateTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoRequest.Unmarshal(m, b)
}
func (m *UpdateTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoRequest.Merge(m, src)
}
func (m *UpdateTodoRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoRequest.Size(m)
}
func (m *UpdateTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoRequest proto.InternalMessageInfo

func (m *UpdateTodoRequest) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type UpdateTodoResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodoResponse) Reset()         { *m = UpdateTodoResponse{} }
func (m *UpdateTodoResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoResponse) ProtoMessage()    {}
func (*UpdateTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec8192e4acf1e0e, []int{5}
}

func (m *UpdateTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoResponse.Unmarshal(m, b)
}
func (m *UpdateTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoResponse.Marshal(b, m, deterministic)
}
func (m *UpdateTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoResponse.Merge(m, src)
}
func (m *UpdateTodoResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoResponse.Size(m)
}
func (m *UpdateTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoResponse proto.InternalMessageInfo

type DeleteTodoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoRequest) Reset()         { *m = DeleteTodoRequest{} }
func (m *DeleteTodoRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoRequest) ProtoMessage()    {}
func (*DeleteTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec8192e4acf1e0e, []int{6}
}

func (m *DeleteTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoRequest.Unmarshal(m, b)
}
func (m *DeleteTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoRequest.Merge(m, src)
}
func (m *DeleteTodoRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoRequest.Size(m)
}
func (m *DeleteTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoRequest proto.InternalMessageInfo

func (m *DeleteTodoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteTodoResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoResponse) Reset()         { *m = DeleteTodoResponse{} }
func (m *DeleteTodoResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoResponse) ProtoMessage()    {}
func (*DeleteTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec8192e4acf1e0e, []int{7}
}

func (m *DeleteTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoResponse.Unmarshal(m, b)
}
func (m *DeleteTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoResponse.Merge(m, src)
}
func (m *DeleteTodoResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoResponse.Size(m)
}
func (m *DeleteTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoResponse proto.InternalMessageInfo

func (m *DeleteTodoResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTodosRequest struct {
	StartDate            *ScheduledDate `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate              *ScheduledDate `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetTodosRequest) Reset()         { *m = GetTodosRequest{} }
func (m *GetTodosRequest) String() string { return proto.CompactTextString(m) }
func (*GetTodosRequest) ProtoMessage()    {}
func (*GetTodosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec8192e4acf1e0e, []int{8}
}

func (m *GetTodosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodosRequest.Unmarshal(m, b)
}
func (m *GetTodosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodosRequest.Marshal(b, m, deterministic)
}
func (m *GetTodosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodosRequest.Merge(m, src)
}
func (m *GetTodosRequest) XXX_Size() int {
	return xxx_messageInfo_GetTodosRequest.Size(m)
}
func (m *GetTodosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodosRequest proto.InternalMessageInfo

func (m *GetTodosRequest) GetStartDate() *ScheduledDate {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *GetTodosRequest) GetEndDate() *ScheduledDate {
	if m != nil {
		return m.EndDate
	}
	return nil
}

type GetTodosResponse struct {
	Todos                []*Todo  `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodosResponse) Reset()         { *m = GetTodosResponse{} }
func (m *GetTodosResponse) String() string { return proto.CompactTextString(m) }
func (*GetTodosResponse) ProtoMessage()    {}
func (*GetTodosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec8192e4acf1e0e, []int{9}
}

func (m *GetTodosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodosResponse.Unmarshal(m, b)
}
func (m *GetTodosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodosResponse.Marshal(b, m, deterministic)
}
func (m *GetTodosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodosResponse.Merge(m, src)
}
func (m *GetTodosResponse) XXX_Size() int {
	return xxx_messageInfo_GetTodosResponse.Size(m)
}
func (m *GetTodosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodosResponse proto.InternalMessageInfo

func (m *GetTodosResponse) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

func init() {
	proto.RegisterType((*ScheduledDate)(nil), "todo.ScheduledDate")
	proto.RegisterType((*Todo)(nil), "todo.Todo")
	proto.RegisterType((*CreateTodoRequest)(nil), "todo.CreateTodoRequest")
	proto.RegisterType((*CreateTodoResponse)(nil), "todo.CreateTodoResponse")
	proto.RegisterType((*UpdateTodoRequest)(nil), "todo.UpdateTodoRequest")
	proto.RegisterType((*UpdateTodoResponse)(nil), "todo.UpdateTodoResponse")
	proto.RegisterType((*DeleteTodoRequest)(nil), "todo.DeleteTodoRequest")
	proto.RegisterType((*DeleteTodoResponse)(nil), "todo.DeleteTodoResponse")
	proto.RegisterType((*GetTodosRequest)(nil), "todo.GetTodosRequest")
	proto.RegisterType((*GetTodosResponse)(nil), "todo.GetTodosResponse")
}

func init() { proto.RegisterFile("app/proto/v1/todo/todo.proto", fileDescriptor_5ec8192e4acf1e0e) }

var fileDescriptor_5ec8192e4acf1e0e = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xc1, 0x6e, 0xd3, 0x4e,
	0x10, 0xc6, 0x65, 0x27, 0x69, 0xe3, 0xb1, 0xd2, 0x36, 0xdb, 0xf6, 0xff, 0x5f, 0x45, 0x15, 0xb2,
	0x0c, 0x87, 0xaa, 0x87, 0x58, 0xa4, 0x9c, 0x7a, 0xa5, 0x12, 0x07, 0x38, 0xb9, 0x85, 0x2b, 0xda,
	0x66, 0x47, 0xad, 0x91, 0xf1, 0x1a, 0xef, 0xa6, 0x52, 0x85, 0xb8, 0xf0, 0x0a, 0xbc, 0x0d, 0xaf,
	0xc1, 0x03, 0x70, 0xe1, 0x41, 0xd0, 0xce, 0x7a, 0x89, 0x13, 0xd3, 0x03, 0x97, 0x68, 0x67, 0xf6,
	0x9b, 0x5f, 0x32, 0xdf, 0x17, 0x1b, 0x4e, 0x44, 0x5d, 0x67, 0x75, 0xa3, 0x8c, 0xca, 0xee, 0x9f,
	0x67, 0x46, 0x49, 0x45, 0x1f, 0x73, 0x6a, 0xb1, 0xa1, 0x3d, 0xcf, 0x4e, 0x6e, 0x95, 0xba, 0x2d,
	0x31, 0x13, 0x75, 0x91, 0x89, 0xaa, 0x52, 0x46, 0x98, 0x42, 0x55, 0xda, 0x69, 0xd2, 0xd7, 0x30,
	0xb9, 0x5a, 0xde, 0xa1, 0x5c, 0x95, 0x28, 0x2f, 0x85, 0x41, 0xc6, 0x60, 0xf8, 0x80, 0xa2, 0xe1,
	0x41, 0x12, 0x9c, 0x8e, 0x72, 0x3a, 0xb3, 0x23, 0x18, 0x7d, 0x54, 0x95, 0xb9, 0xe3, 0x21, 0x35,
	0x5d, 0xc1, 0x0e, 0x60, 0x20, 0xc5, 0x03, 0x1f, 0x50, 0xcf, 0x1e, 0xd3, 0xef, 0x01, 0x0c, 0xaf,
	0x95, 0x54, 0x6c, 0x0f, 0xc2, 0x42, 0x12, 0x22, 0xca, 0xc3, 0x42, 0x5a, 0x80, 0x29, 0x4c, 0x89,
	0x04, 0x88, 0x72, 0x57, 0x30, 0x0e, 0xbb, 0x75, 0xa3, 0x3e, 0xe0, 0xd2, 0x10, 0x24, 0xca, 0x7d,
	0x69, 0xf5, 0xa5, 0xb8, 0xc1, 0x92, 0x0f, 0x9d, 0x9e, 0x0a, 0x96, 0x40, 0x2c, 0x51, 0x2f, 0x9b,
	0xa2, 0xb6, 0x1b, 0xf0, 0x11, 0xdd, 0x75, 0x5b, 0xec, 0x02, 0xf6, 0xb4, 0xdf, 0xe6, 0xbd, 0x14,
	0x06, 0xf9, 0x4e, 0x12, 0x9c, 0xc6, 0x8b, 0xc3, 0x39, 0xd9, 0xb2, 0xb1, 0x69, 0x3e, 0xd1, 0xdd,
	0x32, 0x3d, 0x87, 0xe9, 0xcb, 0x06, 0x85, 0x41, 0xbb, 0x41, 0x8e, 0x9f, 0x56, 0xa8, 0x0d, 0x7b,
	0x02, 0x64, 0x22, 0xad, 0x12, 0x2f, 0xc0, 0x61, 0x48, 0x40, 0xfd, 0xf4, 0x08, 0x58, 0x77, 0x48,
	0xd7, 0xaa, 0xd2, 0x84, 0x7a, 0x5b, 0xcb, 0x7f, 0x47, 0x75, 0x87, 0x5a, 0xd4, 0x53, 0x98, 0x5e,
	0x62, 0x89, 0x9b, 0xa8, 0x2d, 0x7b, 0xd3, 0x67, 0xc0, 0xba, 0x22, 0x37, 0xda, 0x53, 0xad, 0x60,
	0xff, 0x15, 0x1a, 0x2b, 0xd1, 0x1e, 0xb4, 0x00, 0xd0, 0x46, 0x34, 0xc6, 0x79, 0x15, 0x3c, 0xee,
	0x55, 0x44, 0x32, 0xfa, 0x83, 0xcc, 0x61, 0x8c, 0x55, 0xeb, 0x6e, 0xf8, 0xf8, 0xc4, 0x2e, 0x56,
	0xce, 0xd7, 0x17, 0x70, 0xb0, 0xfe, 0xda, 0xf6, 0xa7, 0x25, 0x30, 0xb2, 0x23, 0x9a, 0x07, 0xc9,
	0x60, 0xcb, 0x0c, 0x77, 0xb1, 0xf8, 0x19, 0x42, 0x6c, 0xeb, 0x2b, 0x6c, 0xee, 0x8b, 0x25, 0xb2,
	0x6b, 0x80, 0xb5, 0xd1, 0xec, 0x7f, 0x37, 0xd0, 0xcb, 0x6b, 0xc6, 0xfb, 0x17, 0xad, 0x91, 0x87,
	0x5f, 0x7f, 0xfc, 0xfa, 0x16, 0x4e, 0xd2, 0xb1, 0x7f, 0x52, 0x2e, 0x82, 0x33, 0x4b, 0x5d, 0x7b,
	0xee, 0xa9, 0xbd, 0xe8, 0x3c, 0xf5, 0x2f, 0xf1, 0xb4, 0xd4, 0xd9, 0x06, 0xf5, 0x1d, 0xc0, 0x3a,
	0x0e, 0x4f, 0xed, 0xa5, 0xe8, 0xa9, 0xfd, 0xe4, 0xd2, 0x63, 0xa2, 0xee, 0x9f, 0x4d, 0xfe, 0x3c,
	0xd5, 0x9f, 0x0b, 0xf9, 0x85, 0xbd, 0x81, 0xb1, 0x77, 0x92, 0x1d, 0xbb, 0xe1, 0xad, 0x40, 0x67,
	0xff, 0x6d, 0xb7, 0x5b, 0xe2, 0x94, 0x88, 0x31, 0x8b, 0x3c, 0x51, 0xdf, 0xec, 0xd0, 0x0b, 0xe0,
	0xfc, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xbd, 0xa5, 0xa1, 0x44, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error)
	UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*UpdateTodoResponse, error)
	DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error)
	GetTodos(ctx context.Context, in *GetTodosRequest, opts ...grpc.CallOption) (*GetTodosResponse, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error) {
	out := new(CreateTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.TodoService/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*UpdateTodoResponse, error) {
	out := new(UpdateTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.TodoService/UpdateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error) {
	out := new(DeleteTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.TodoService/DeleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetTodos(ctx context.Context, in *GetTodosRequest, opts ...grpc.CallOption) (*GetTodosResponse, error) {
	out := new(GetTodosResponse)
	err := c.cc.Invoke(ctx, "/todo.TodoService/GetTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	CreateTodo(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error)
	UpdateTodo(context.Context, *UpdateTodoRequest) (*UpdateTodoResponse, error)
	DeleteTodo(context.Context, *DeleteTodoRequest) (*DeleteTodoResponse, error)
	GetTodos(context.Context, *GetTodosRequest) (*GetTodosResponse, error)
}

// UnimplementedTodoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (*UnimplementedTodoServiceServer) CreateTodo(ctx context.Context, req *CreateTodoRequest) (*CreateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (*UnimplementedTodoServiceServer) UpdateTodo(ctx context.Context, req *UpdateTodoRequest) (*UpdateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (*UnimplementedTodoServiceServer) DeleteTodo(ctx context.Context, req *DeleteTodoRequest) (*DeleteTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}
func (*UnimplementedTodoServiceServer) GetTodos(ctx context.Context, req *GetTodosRequest) (*GetTodosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodos not implemented")
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodo(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/UpdateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateTodo(ctx, req.(*UpdateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodo(ctx, req.(*DeleteTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/GetTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetTodos(ctx, req.(*GetTodosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoService_CreateTodo_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _TodoService_UpdateTodo_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _TodoService_DeleteTodo_Handler,
		},
		{
			MethodName: "GetTodos",
			Handler:    _TodoService_GetTodos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/proto/v1/todo/todo.proto",
}
