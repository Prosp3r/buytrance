// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

//void - a blank parameter for empty parameter entry purpose only
type Emptyentry struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Emptyentry) Reset()         { *m = Emptyentry{} }
func (m *Emptyentry) String() string { return proto.CompactTextString(m) }
func (*Emptyentry) ProtoMessage()    {}
func (*Emptyentry) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *Emptyentry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Emptyentry.Unmarshal(m, b)
}
func (m *Emptyentry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Emptyentry.Marshal(b, m, deterministic)
}
func (m *Emptyentry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Emptyentry.Merge(m, src)
}
func (m *Emptyentry) XXX_Size() int {
	return xxx_messageInfo_Emptyentry.Size(m)
}
func (m *Emptyentry) XXX_DiscardUnknown() {
	xxx_messageInfo_Emptyentry.DiscardUnknown(m)
}

var xxx_messageInfo_Emptyentry proto.InternalMessageInfo

type IDString struct {
	Info                 string   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDString) Reset()         { *m = IDString{} }
func (m *IDString) String() string { return proto.CompactTextString(m) }
func (*IDString) ProtoMessage()    {}
func (*IDString) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *IDString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDString.Unmarshal(m, b)
}
func (m *IDString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDString.Marshal(b, m, deterministic)
}
func (m *IDString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDString.Merge(m, src)
}
func (m *IDString) XXX_Size() int {
	return xxx_messageInfo_IDString.Size(m)
}
func (m *IDString) XXX_DiscardUnknown() {
	xxx_messageInfo_IDString.DiscardUnknown(m)
}

var xxx_messageInfo_IDString proto.InternalMessageInfo

func (m *IDString) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type BOOLValue struct {
	Info                 bool     `protobuf:"varint,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BOOLValue) Reset()         { *m = BOOLValue{} }
func (m *BOOLValue) String() string { return proto.CompactTextString(m) }
func (*BOOLValue) ProtoMessage()    {}
func (*BOOLValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *BOOLValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BOOLValue.Unmarshal(m, b)
}
func (m *BOOLValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BOOLValue.Marshal(b, m, deterministic)
}
func (m *BOOLValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BOOLValue.Merge(m, src)
}
func (m *BOOLValue) XXX_Size() int {
	return xxx_messageInfo_BOOLValue.Size(m)
}
func (m *BOOLValue) XXX_DiscardUnknown() {
	xxx_messageInfo_BOOLValue.DiscardUnknown(m)
}

var xxx_messageInfo_BOOLValue proto.InternalMessageInfo

func (m *BOOLValue) GetInfo() bool {
	if m != nil {
		return m.Info
	}
	return false
}

//User - models user account information
type User struct {
	Uuid                 string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Accountstatus        string               `protobuf:"bytes,4,opt,name=accountstatus,proto3" json:"accountstatus,omitempty"`
	Verificationcode     string               `protobuf:"bytes,5,opt,name=verificationcode,proto3" json:"verificationcode,omitempty"`
	Codestatus           string               `protobuf:"bytes,6,opt,name=codestatus,proto3" json:"codestatus,omitempty"`
	Datejoined           *timestamp.Timestamp `protobuf:"bytes,8,opt,name=datejoined,proto3" json:"datejoined,omitempty"`
	Lastupdated          *timestamp.Timestamp `protobuf:"bytes,9,opt,name=lastupdated,proto3" json:"lastupdated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetAccountstatus() string {
	if m != nil {
		return m.Accountstatus
	}
	return ""
}

func (m *User) GetVerificationcode() string {
	if m != nil {
		return m.Verificationcode
	}
	return ""
}

func (m *User) GetCodestatus() string {
	if m != nil {
		return m.Codestatus
	}
	return ""
}

func (m *User) GetDatejoined() *timestamp.Timestamp {
	if m != nil {
		return m.Datejoined
	}
	return nil
}

func (m *User) GetLastupdated() *timestamp.Timestamp {
	if m != nil {
		return m.Lastupdated
	}
	return nil
}

type UserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

//Users - with an (s)plural structure will take a list/array of user profiles.
type Users struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (m *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(m, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*Emptyentry)(nil), "user.Emptyentry")
	proto.RegisterType((*IDString)(nil), "user.IDString")
	proto.RegisterType((*BOOLValue)(nil), "user.BOOLValue")
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*UserResponse)(nil), "user.UserResponse")
	proto.RegisterType((*Users)(nil), "user.Users")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6b, 0xdb, 0x40,
	0x10, 0xc5, 0xad, 0xc4, 0x4e, 0xe4, 0x51, 0x9a, 0x86, 0xa1, 0x87, 0x45, 0x87, 0xc4, 0x2c, 0x3d,
	0xa4, 0x0d, 0xc8, 0xe0, 0xde, 0x4a, 0x2f, 0x2d, 0x09, 0xa5, 0x50, 0x08, 0xd8, 0x75, 0xef, 0x1b,
	0x69, 0x6c, 0xb6, 0x58, 0xbb, 0x42, 0xbb, 0x6b, 0xc8, 0x37, 0xeb, 0x57, 0xeb, 0xad, 0xec, 0xc8,
	0x4e, 0x94, 0xfe, 0xa1, 0x3d, 0x69, 0xf6, 0xcd, 0xef, 0x89, 0xe1, 0xcd, 0x00, 0x04, 0x47, 0x6d,
	0xd1, 0xb4, 0xd6, 0x5b, 0x1c, 0xc6, 0x3a, 0xbf, 0x58, 0x5b, 0xbb, 0xde, 0xd0, 0x94, 0xb5, 0xbb,
	0xb0, 0x9a, 0x7a, 0x5d, 0x93, 0xf3, 0xaa, 0x6e, 0x3a, 0x4c, 0x9e, 0x00, 0xdc, 0xd4, 0x8d, 0xbf,
	0x27, 0xe3, 0xdb, 0x7b, 0x79, 0x0e, 0xe9, 0xa7, 0xeb, 0x85, 0x6f, 0xb5, 0x59, 0x23, 0xc2, 0x50,
	0x9b, 0x95, 0x15, 0xc9, 0x24, 0xb9, 0x1c, 0xcf, 0xb9, 0x96, 0x17, 0x30, 0xfe, 0x70, 0x7b, 0xfb,
	0xf9, 0xab, 0xda, 0x04, 0x7a, 0x02, 0xa4, 0x3b, 0xe0, 0xfb, 0x01, 0x0c, 0x97, 0x8e, 0xda, 0xd8,
	0x0c, 0x41, 0x57, 0x7b, 0x77, 0xac, 0x31, 0x87, 0x34, 0x0e, 0x65, 0x54, 0x4d, 0xe2, 0x80, 0xf5,
	0x87, 0x37, 0xbe, 0x80, 0x11, 0xd5, 0x4a, 0x6f, 0xc4, 0x21, 0x37, 0xba, 0x07, 0xbe, 0x84, 0x67,
	0xaa, 0x2c, 0x6d, 0x30, 0xde, 0x79, 0xe5, 0x83, 0x13, 0x43, 0xee, 0x3e, 0x15, 0xf1, 0x35, 0x9c,
	0x6d, 0xa9, 0xd5, 0x2b, 0x5d, 0x2a, 0xaf, 0xad, 0x29, 0x6d, 0x45, 0x62, 0xc4, 0xe0, 0x6f, 0x3a,
	0x9e, 0x03, 0xc4, 0xef, 0xee, 0x77, 0x47, 0x4c, 0xf5, 0x14, 0x7c, 0x0b, 0x50, 0x29, 0x4f, 0xdf,
	0xac, 0x36, 0x54, 0x89, 0x74, 0x92, 0x5c, 0x66, 0xb3, 0xbc, 0xe8, 0x52, 0x2c, 0xf6, 0x29, 0x16,
	0x5f, 0xf6, 0x29, 0xce, 0x7b, 0x34, 0xbe, 0x83, 0x6c, 0xa3, 0x9c, 0x0f, 0x4d, 0xd4, 0x2a, 0x31,
	0xfe, 0xa7, 0xb9, 0x8f, 0xcb, 0x02, 0x4e, 0x62, 0x72, 0x73, 0x72, 0x8d, 0x35, 0x2e, 0x4e, 0xca,
	0x2b, 0xe4, 0x04, 0xb3, 0x19, 0x14, 0xbc, 0x5b, 0x26, 0x58, 0x97, 0xaf, 0x60, 0x14, 0x5f, 0x0e,
	0x27, 0x30, 0x8a, 0x82, 0x13, 0xc9, 0xe4, 0xf0, 0x17, 0xb2, 0x6b, 0xcc, 0x7e, 0x24, 0x90, 0x2d,
	0x17, 0x37, 0xf3, 0x05, 0xb5, 0x5b, 0x5d, 0x12, 0x5e, 0xc1, 0xf1, 0xfb, 0xaa, 0xe2, 0x3d, 0xf5,
	0xe8, 0x1c, 0x7b, 0xce, 0xdd, 0x14, 0x72, 0x80, 0x53, 0x38, 0xfe, 0x48, 0x9e, 0xe1, 0xd3, 0x0e,
	0xd8, 0x9f, 0xc8, 0x5f, 0x0c, 0x57, 0x90, 0xee, 0x0c, 0x0e, 0xcf, 0x3a, 0xe2, 0xf1, 0xc4, 0xf2,
	0xec, 0xd1, 0xe3, 0xe4, 0x00, 0x0b, 0x80, 0x25, 0x07, 0xf0, 0xdf, 0xd3, 0x9c, 0x5e, 0x93, 0x2a,
	0xbd, 0xde, 0xfe, 0xc9, 0xf3, 0xbc, 0xab, 0x1f, 0x6e, 0x54, 0x0e, 0xee, 0x8e, 0x38, 0xf7, 0x37,
	0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xfc, 0x0c, 0x73, 0x1c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// USERServiceClient is the client API for USERService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type USERServiceClient interface {
	AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error)
	GetUser(ctx context.Context, in *IDString, opts ...grpc.CallOption) (*UserResponse, error)
	GetUsers(ctx context.Context, in *Emptyentry, opts ...grpc.CallOption) (*Users, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error)
	DeactivateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*BOOLValue, error)
}

type uSERServiceClient struct {
	cc *grpc.ClientConn
}

func NewUSERServiceClient(cc *grpc.ClientConn) USERServiceClient {
	return &uSERServiceClient{cc}
}

func (c *uSERServiceClient) AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.USERService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uSERServiceClient) GetUser(ctx context.Context, in *IDString, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.USERService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uSERServiceClient) GetUsers(ctx context.Context, in *Emptyentry, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/user.USERService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uSERServiceClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.USERService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uSERServiceClient) DeactivateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*BOOLValue, error) {
	out := new(BOOLValue)
	err := c.cc.Invoke(ctx, "/user.USERService/DeactivateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// USERServiceServer is the server API for USERService service.
type USERServiceServer interface {
	AddUser(context.Context, *User) (*UserResponse, error)
	GetUser(context.Context, *IDString) (*UserResponse, error)
	GetUsers(context.Context, *Emptyentry) (*Users, error)
	UpdateUser(context.Context, *User) (*UserResponse, error)
	DeactivateUser(context.Context, *User) (*BOOLValue, error)
}

// UnimplementedUSERServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUSERServiceServer struct {
}

func (*UnimplementedUSERServiceServer) AddUser(ctx context.Context, req *User) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (*UnimplementedUSERServiceServer) GetUser(ctx context.Context, req *IDString) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUSERServiceServer) GetUsers(ctx context.Context, req *Emptyentry) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (*UnimplementedUSERServiceServer) UpdateUser(ctx context.Context, req *User) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUSERServiceServer) DeactivateUser(ctx context.Context, req *User) (*BOOLValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateUser not implemented")
}

func RegisterUSERServiceServer(s *grpc.Server, srv USERServiceServer) {
	s.RegisterService(&_USERService_serviceDesc, srv)
}

func _USERService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(USERServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.USERService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(USERServiceServer).AddUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _USERService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(USERServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.USERService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(USERServiceServer).GetUser(ctx, req.(*IDString))
	}
	return interceptor(ctx, in, info, handler)
}

func _USERService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Emptyentry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(USERServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.USERService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(USERServiceServer).GetUsers(ctx, req.(*Emptyentry))
	}
	return interceptor(ctx, in, info, handler)
}

func _USERService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(USERServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.USERService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(USERServiceServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _USERService_DeactivateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(USERServiceServer).DeactivateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.USERService/DeactivateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(USERServiceServer).DeactivateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _USERService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.USERService",
	HandlerType: (*USERServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _USERService_AddUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _USERService_GetUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _USERService_GetUsers_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _USERService_UpdateUser_Handler,
		},
		{
			MethodName: "DeactivateUser",
			Handler:    _USERService_DeactivateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
