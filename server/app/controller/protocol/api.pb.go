// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package app

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

type ApiReq struct {
	ReqName              *string  `protobuf:"bytes,1,req,name=ReqName" json:"ReqName,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,req,name=Data" json:"Data,omitempty"`
	Flag                 *int32   `protobuf:"varint,3,opt,name=Flag" json:"Flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiReq) Reset()         { *m = ApiReq{} }
func (m *ApiReq) String() string { return proto.CompactTextString(m) }
func (*ApiReq) ProtoMessage()    {}
func (*ApiReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *ApiReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiReq.Unmarshal(m, b)
}
func (m *ApiReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiReq.Marshal(b, m, deterministic)
}
func (m *ApiReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiReq.Merge(m, src)
}
func (m *ApiReq) XXX_Size() int {
	return xxx_messageInfo_ApiReq.Size(m)
}
func (m *ApiReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiReq.DiscardUnknown(m)
}

var xxx_messageInfo_ApiReq proto.InternalMessageInfo

func (m *ApiReq) GetReqName() string {
	if m != nil && m.ReqName != nil {
		return *m.ReqName
	}
	return ""
}

func (m *ApiReq) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ApiReq) GetFlag() int32 {
	if m != nil && m.Flag != nil {
		return *m.Flag
	}
	return 0
}

type ApiAck struct {
	AckName              *string  `protobuf:"bytes,1,req,name=AckName" json:"AckName,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,req,name=Data" json:"Data,omitempty"`
	Flag                 *int32   `protobuf:"varint,3,opt,name=Flag" json:"Flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiAck) Reset()         { *m = ApiAck{} }
func (m *ApiAck) String() string { return proto.CompactTextString(m) }
func (*ApiAck) ProtoMessage()    {}
func (*ApiAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *ApiAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiAck.Unmarshal(m, b)
}
func (m *ApiAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiAck.Marshal(b, m, deterministic)
}
func (m *ApiAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiAck.Merge(m, src)
}
func (m *ApiAck) XXX_Size() int {
	return xxx_messageInfo_ApiAck.Size(m)
}
func (m *ApiAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiAck.DiscardUnknown(m)
}

var xxx_messageInfo_ApiAck proto.InternalMessageInfo

func (m *ApiAck) GetAckName() string {
	if m != nil && m.AckName != nil {
		return *m.AckName
	}
	return ""
}

func (m *ApiAck) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ApiAck) GetFlag() int32 {
	if m != nil && m.Flag != nil {
		return *m.Flag
	}
	return 0
}

type Message struct {
	Data                 *string  `protobuf:"bytes,1,req,name=Data" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
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

func (m *Message) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*ApiReq)(nil), "app.ApiReq")
	proto.RegisterType((*ApiAck)(nil), "app.ApiAck")
	proto.RegisterType((*Message)(nil), "app.Message")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0x28, 0x50, 0xf2, 0xe2, 0x62, 0x73, 0x2c,
	0xc8, 0x0c, 0x4a, 0x2d, 0x14, 0x92, 0xe0, 0x62, 0x0f, 0x4a, 0x2d, 0xf4, 0x4b, 0xcc, 0x4d, 0x95,
	0x60, 0x54, 0x60, 0xd2, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0x84, 0xb8, 0x58, 0x5c, 0x12, 0x4b, 0x12,
	0x25, 0x98, 0x14, 0x98, 0x34, 0x78, 0x82, 0xc0, 0x6c, 0x90, 0x98, 0x5b, 0x4e, 0x62, 0xba, 0x04,
	0xb3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x98, 0x0d, 0x35, 0xcb, 0x31, 0x39, 0x1b, 0x64, 0x96, 0x63,
	0x72, 0x36, 0xb2, 0x59, 0x50, 0x2e, 0xd1, 0x66, 0xc9, 0x72, 0xb1, 0xfb, 0xa6, 0x16, 0x17, 0x27,
	0xa6, 0x23, 0xb4, 0x40, 0x4c, 0x02, 0xb3, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x89, 0xd9, 0x21,
	0xf5, 0xc7, 0x00, 0x00, 0x00,
}