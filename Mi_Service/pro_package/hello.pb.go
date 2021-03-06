// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package pro_package

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

//客户端发送给服务端
type HelloReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReq) Reset()         { *m = HelloReq{} }
func (m *HelloReq) String() string { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()    {}
func (*HelloReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *HelloReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReq.Unmarshal(m, b)
}
func (m *HelloReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReq.Marshal(b, m, deterministic)
}
func (m *HelloReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReq.Merge(m, src)
}
func (m *HelloReq) XXX_Size() int {
	return xxx_messageInfo_HelloReq.Size(m)
}
func (m *HelloReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReq.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReq proto.InternalMessageInfo

func (m *HelloReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

//服务端返回给客户端
type HelloRsp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRsp) Reset()         { *m = HelloRsp{} }
func (m *HelloRsp) String() string { return proto.CompactTextString(m) }
func (*HelloRsp) ProtoMessage()    {}
func (*HelloRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *HelloRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRsp.Unmarshal(m, b)
}
func (m *HelloRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRsp.Marshal(b, m, deterministic)
}
func (m *HelloRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRsp.Merge(m, src)
}
func (m *HelloRsp) XXX_Size() int {
	return xxx_messageInfo_HelloRsp.Size(m)
}
func (m *HelloRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRsp.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRsp proto.InternalMessageInfo

func (m *HelloRsp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type NameReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameReq) Reset()         { *m = NameReq{} }
func (m *NameReq) String() string { return proto.CompactTextString(m) }
func (*NameReq) ProtoMessage()    {}
func (*NameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{2}
}

func (m *NameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameReq.Unmarshal(m, b)
}
func (m *NameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameReq.Marshal(b, m, deterministic)
}
func (m *NameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameReq.Merge(m, src)
}
func (m *NameReq) XXX_Size() int {
	return xxx_messageInfo_NameReq.Size(m)
}
func (m *NameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NameReq.DiscardUnknown(m)
}

var xxx_messageInfo_NameReq proto.InternalMessageInfo

func (m *NameReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type NameRsp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameRsp) Reset()         { *m = NameRsp{} }
func (m *NameRsp) String() string { return proto.CompactTextString(m) }
func (*NameRsp) ProtoMessage()    {}
func (*NameRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{3}
}

func (m *NameRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameRsp.Unmarshal(m, b)
}
func (m *NameRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameRsp.Marshal(b, m, deterministic)
}
func (m *NameRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameRsp.Merge(m, src)
}
func (m *NameRsp) XXX_Size() int {
	return xxx_messageInfo_NameRsp.Size(m)
}
func (m *NameRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_NameRsp.DiscardUnknown(m)
}

var xxx_messageInfo_NameRsp proto.InternalMessageInfo

func (m *NameRsp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloReq)(nil), "HelloReq")
	proto.RegisterType((*HelloRsp)(nil), "HelloRsp")
	proto.RegisterType((*NameReq)(nil), "NameReq")
	proto.RegisterType((*NameRsp)(nil), "NameRsp")
}

func init() {
	proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce)
}

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe3, 0xe2, 0xf0, 0x00, 0x71, 0x83, 0x52,
	0x0b, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0xc0, 0x6c, 0x25, 0x19, 0x98, 0x7c, 0x71, 0x81, 0x90, 0x00, 0x17, 0x73, 0x6e, 0x71, 0x3a, 0x54,
	0x1a, 0xc4, 0x54, 0x92, 0xe5, 0x62, 0xf7, 0x4b, 0xcc, 0x4d, 0xc5, 0xa5, 0x59, 0x1a, 0x2a, 0x8d,
	0x4d, 0xaf, 0x51, 0x10, 0x17, 0x37, 0xd8, 0xe4, 0xe0, 0xd4, 0xa2, 0xb2, 0xd4, 0x22, 0x21, 0x25,
	0x2e, 0x8e, 0xe0, 0xc4, 0x4a, 0xb0, 0x88, 0x10, 0xa7, 0x1e, 0xcc, 0x4d, 0x52, 0x30, 0x66, 0x71,
	0x81, 0x12, 0x83, 0x90, 0x3c, 0x17, 0x7b, 0x70, 0x62, 0x25, 0xc8, 0x48, 0x21, 0x0e, 0x3d, 0xa8,
	0xc5, 0x52, 0x50, 0x16, 0x48, 0x81, 0x93, 0x71, 0x94, 0xa1, 0x9b, 0x95, 0x7e, 0x70, 0x7e, 0x5a,
	0x49, 0x79, 0x62, 0x51, 0xaa, 0x7e, 0x7a, 0x7e, 0x7c, 0x41, 0x62, 0x49, 0x86, 0x7e, 0x6e, 0x65,
	0x7c, 0x41, 0x51, 0xbe, 0xbe, 0x6f, 0x66, 0x3c, 0xc8, 0xb2, 0xcc, 0xe4, 0x54, 0xfd, 0x82, 0x22,
	0x90, 0x54, 0x72, 0x76, 0x62, 0x7a, 0x6a, 0x12, 0x1b, 0x38, 0x24, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xbb, 0xcb, 0x43, 0x2d, 0x18, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloServerClient is the client API for HelloServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServerClient interface {
	SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRsp, error)
	SayName(ctx context.Context, in *NameReq, opts ...grpc.CallOption) (*NameRsp, error)
}

type helloServerClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServerClient(cc grpc.ClientConnInterface) HelloServerClient {
	return &helloServerClient{cc}
}

func (c *helloServerClient) SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRsp, error) {
	out := new(HelloRsp)
	err := c.cc.Invoke(ctx, "/HelloServer/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServerClient) SayName(ctx context.Context, in *NameReq, opts ...grpc.CallOption) (*NameRsp, error) {
	out := new(NameRsp)
	err := c.cc.Invoke(ctx, "/HelloServer/SayName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServerServer is the server API for HelloServer service.
type HelloServerServer interface {
	SayHello(context.Context, *HelloReq) (*HelloRsp, error)
	SayName(context.Context, *NameReq) (*NameRsp, error)
}

// UnimplementedHelloServerServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServerServer struct {
}

func (*UnimplementedHelloServerServer) SayHello(ctx context.Context, req *HelloReq) (*HelloRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedHelloServerServer) SayName(ctx context.Context, req *NameReq) (*NameRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayName not implemented")
}

func RegisterHelloServerServer(s *grpc.Server, srv HelloServerServer) {
	s.RegisterService(&_HelloServer_serviceDesc, srv)
}

func _HelloServer_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServerServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HelloServer/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServerServer).SayHello(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloServer_SayName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServerServer).SayName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HelloServer/SayName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServerServer).SayName(ctx, req.(*NameReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "HelloServer",
	HandlerType: (*HelloServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloServer_SayHello_Handler,
		},
		{
			MethodName: "SayName",
			Handler:    _HelloServer_SayName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
