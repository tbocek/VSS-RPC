// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.proto

package schema

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

type AMessage struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AMessage) Reset()         { *m = AMessage{} }
func (m *AMessage) String() string { return proto.CompactTextString(m) }
func (*AMessage) ProtoMessage()    {}
func (*AMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0}
}

func (m *AMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AMessage.Unmarshal(m, b)
}
func (m *AMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AMessage.Marshal(b, m, deterministic)
}
func (m *AMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AMessage.Merge(m, src)
}
func (m *AMessage) XXX_Size() int {
	return xxx_messageInfo_AMessage.Size(m)
}
func (m *AMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AMessage proto.InternalMessageInfo

func (m *AMessage) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*AMessage)(nil), "AMessage")
}

func init() { proto.RegisterFile("schema.proto", fileDescriptor_1c5fb4d8cc22d66a) }

var fileDescriptor_1c5fb4d8cc22d66a = []byte{
	// 88 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xb2, 0xe0, 0xe2, 0x70, 0xf4, 0x4d, 0x2d,
	0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe2, 0x62, 0x49, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0d, 0x02, 0xb3, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x21, 0xd2, 0x12, 0x4c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x30, 0x6e, 0x12, 0x1b, 0xd8, 0x00, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x3c, 0x48, 0x6a, 0x37, 0x50, 0x00, 0x00, 0x00,
}
