// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/GetSmsCd/GetSmsCd.proto

package go_micro_srv_GetSmsCd

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_47b2b2b0bd6df588, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	Uuid                 string   `protobuf:"bytes,2,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
	Text                 string   `protobuf:"bytes,3,opt,name=Text,proto3" json:"Text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_47b2b2b0bd6df588, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Request) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Request) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Response struct {
	Error                string   `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_47b2b2b0bd6df588, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Response) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.GetSmsCd.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.GetSmsCd.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.GetSmsCd.Response")
}

func init() {
	proto.RegisterFile("proto/GetSmsCd/GetSmsCd.proto", fileDescriptor_47b2b2b0bd6df588)
}

var fileDescriptor_47b2b2b0bd6df588 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x77, 0x4f, 0x2d, 0x09, 0xce, 0x2d, 0x76, 0x4e, 0x81, 0x33, 0xf4, 0xc0, 0xe2, 0x42,
	0xa2, 0xe9, 0xf9, 0x7a, 0xb9, 0x99, 0xc9, 0x45, 0xf9, 0x7a, 0xc5, 0x45, 0x65, 0x7a, 0x30, 0x49,
	0x25, 0x69, 0x2e, 0x76, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x21, 0x01, 0x2e, 0xe6, 0xe2,
	0xc4, 0x4a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0xc9, 0x93, 0x8b, 0x3d, 0x28,
	0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x8c, 0x8b, 0xcd, 0x37, 0x3f, 0x29, 0x33, 0x27, 0x15,
	0x2a, 0x0f, 0xe5, 0x09, 0x09, 0x71, 0xb1, 0x84, 0x96, 0x66, 0xa6, 0x48, 0x30, 0x81, 0x45, 0xc1,
	0x6c, 0x90, 0x58, 0x48, 0x6a, 0x45, 0x89, 0x04, 0x33, 0x44, 0x0c, 0xc4, 0x56, 0xb2, 0xe0, 0xe2,
	0x08, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe1, 0x62, 0x75, 0x2d, 0x2a, 0xca,
	0x2f, 0x82, 0x1a, 0x05, 0xe1, 0x80, 0x6c, 0x70, 0x2d, 0x2a, 0xf2, 0x2d, 0x4e, 0x87, 0x9a, 0x05,
	0xe5, 0x19, 0xc5, 0x72, 0x71, 0xc0, 0x5c, 0x2b, 0x14, 0xc8, 0xc5, 0xe3, 0x9c, 0x98, 0x93, 0x03,
	0xe7, 0xcb, 0xe9, 0x61, 0xf5, 0x95, 0x1e, 0xd4, 0xd5, 0x52, 0xf2, 0x38, 0xe5, 0x21, 0x4e, 0x51,
	0x62, 0x48, 0x62, 0x03, 0x07, 0x8f, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xca, 0x92, 0x7f,
	0x3f, 0x01, 0x00, 0x00,
}
