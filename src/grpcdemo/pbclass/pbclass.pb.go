// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pbclass.proto

package pbclass

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

type UserStatus int32

const (
	UserStatus_OFFLINE UserStatus = 0
	UserStatus_ONLINE  UserStatus = 1
)

var UserStatus_name = map[int32]string{
	0: "OFFLINE",
	1: "ONLINE",
}

var UserStatus_value = map[string]int32{
	"OFFLINE": 0,
	"ONLINE":  1,
}

func (x UserStatus) String() string {
	return proto.EnumName(UserStatus_name, int32(x))
}

func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b86db13ad72dcc18, []int{0}
}

type UserInfo struct {
	Id                   int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               UserStatus `protobuf:"varint,3,opt,name=status,proto3,enum=UserStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b86db13ad72dcc18, []int{0}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetStatus() UserStatus {
	if m != nil {
		return m.Status
	}
	return UserStatus_OFFLINE
}

type UserID struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserID) Reset()         { *m = UserID{} }
func (m *UserID) String() string { return proto.CompactTextString(m) }
func (*UserID) ProtoMessage()    {}
func (*UserID) Descriptor() ([]byte, []int) {
	return fileDescriptor_b86db13ad72dcc18, []int{1}
}

func (m *UserID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserID.Unmarshal(m, b)
}
func (m *UserID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserID.Marshal(b, m, deterministic)
}
func (m *UserID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserID.Merge(m, src)
}
func (m *UserID) XXX_Size() int {
	return xxx_messageInfo_UserID.Size(m)
}
func (m *UserID) XXX_DiscardUnknown() {
	xxx_messageInfo_UserID.DiscardUnknown(m)
}

var xxx_messageInfo_UserID proto.InternalMessageInfo

func (m *UserID) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FuncResponse struct {
	Reply                string   `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FuncResponse) Reset()         { *m = FuncResponse{} }
func (m *FuncResponse) String() string { return proto.CompactTextString(m) }
func (*FuncResponse) ProtoMessage()    {}
func (*FuncResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b86db13ad72dcc18, []int{2}
}

func (m *FuncResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FuncResponse.Unmarshal(m, b)
}
func (m *FuncResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FuncResponse.Marshal(b, m, deterministic)
}
func (m *FuncResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FuncResponse.Merge(m, src)
}
func (m *FuncResponse) XXX_Size() int {
	return xxx_messageInfo_FuncResponse.Size(m)
}
func (m *FuncResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FuncResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FuncResponse proto.InternalMessageInfo

func (m *FuncResponse) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func init() {
	proto.RegisterEnum("UserStatus", UserStatus_name, UserStatus_value)
	proto.RegisterType((*UserInfo)(nil), "UserInfo")
	proto.RegisterType((*UserID)(nil), "UserID")
	proto.RegisterType((*FuncResponse)(nil), "funcResponse")
}

func init() { proto.RegisterFile("pbclass.proto", fileDescriptor_b86db13ad72dcc18) }

var fileDescriptor_b86db13ad72dcc18 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x9b, 0xd5, 0x4d, 0xdd, 0x59, 0x5b, 0x64, 0xf0, 0x10, 0x3c, 0x2d, 0x51, 0x61, 0xf1,
	0x90, 0x43, 0x7d, 0x85, 0x5a, 0x28, 0x94, 0x16, 0x52, 0xf4, 0xbe, 0xdd, 0xa6, 0x25, 0xb0, 0x26,
	0x21, 0xc9, 0x0a, 0xbe, 0xbd, 0x34, 0xab, 0x28, 0xf6, 0xf6, 0xcd, 0xcc, 0x6f, 0xbe, 0xf9, 0x03,
	0x13, 0xb7, 0x6b, 0xbb, 0x26, 0x04, 0xe1, 0xbc, 0x8d, 0x96, 0x6f, 0xe1, 0xea, 0x35, 0x28, 0xbf,
	0x34, 0x07, 0x8b, 0x53, 0xc8, 0xf4, 0x9e, 0x91, 0x8a, 0xd4, 0xb9, 0xcc, 0xf4, 0x1e, 0x11, 0x2e,
	0x4d, 0xf3, 0xae, 0x58, 0x56, 0x91, 0xba, 0x90, 0x49, 0xe3, 0x3d, 0xd0, 0x10, 0x9b, 0xd8, 0x07,
	0x76, 0x51, 0x91, 0x7a, 0x3a, 0x2b, 0xc5, 0xa9, 0x7d, 0x9b, 0x52, 0xf2, 0xbb, 0xc4, 0x19, 0xd0,
	0x64, 0x3a, 0xff, 0x6f, 0xc9, 0x1f, 0xe0, 0xfa, 0xd0, 0x9b, 0x56, 0xaa, 0xe0, 0xac, 0x09, 0x0a,
	0x6f, 0x21, 0xf7, 0xca, 0x75, 0x9f, 0x09, 0x29, 0xe4, 0x10, 0x3c, 0x3d, 0x02, 0xfc, 0xba, 0x62,
	0x09, 0xe3, 0xcd, 0x62, 0xb1, 0x5a, 0xae, 0x5f, 0x6e, 0x46, 0x08, 0x40, 0x37, 0xeb, 0xa4, 0xc9,
	0xec, 0x0d, 0xca, 0x84, 0x29, 0xff, 0xa1, 0xdb, 0xd3, 0x6a, 0xf9, 0xca, 0x1e, 0xb5, 0xc1, 0x42,
	0xfc, 0x9c, 0x74, 0x37, 0x11, 0x7f, 0xc7, 0xf1, 0x11, 0x72, 0xa0, 0x9d, 0x3d, 0xda, 0x3e, 0xe2,
	0x78, 0xa0, 0xe6, 0x67, 0xcc, 0x8e, 0xa6, 0xd7, 0x3c, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xa3,
	0x27, 0x08, 0x77, 0x2b, 0x01, 0x00, 0x00,
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
	Login(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*FuncResponse, error)
	Logout(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*FuncResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Login(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*FuncResponse, error) {
	out := new(FuncResponse)
	err := c.cc.Invoke(ctx, "/UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Logout(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*FuncResponse, error) {
	out := new(FuncResponse)
	err := c.cc.Invoke(ctx, "/UserService/logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Login(context.Context, *UserInfo) (*FuncResponse, error)
	Logout(context.Context, *UserID) (*FuncResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Login(ctx context.Context, req *UserInfo) (*FuncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedUserServiceServer) Logout(ctx context.Context, req *UserID) (*FuncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Logout(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "logout",
			Handler:    _UserService_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pbclass.proto",
}