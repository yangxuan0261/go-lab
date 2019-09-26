// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld222.proto

package helloworld

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

type HelloRequest222 struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest222) Reset()         { *m = HelloRequest222{} }
func (m *HelloRequest222) String() string { return proto.CompactTextString(m) }
func (*HelloRequest222) ProtoMessage()    {}
func (*HelloRequest222) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aefce65b0a12182, []int{0}
}

func (m *HelloRequest222) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest222.Unmarshal(m, b)
}
func (m *HelloRequest222) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest222.Marshal(b, m, deterministic)
}
func (m *HelloRequest222) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest222.Merge(m, src)
}
func (m *HelloRequest222) XXX_Size() int {
	return xxx_messageInfo_HelloRequest222.Size(m)
}
func (m *HelloRequest222) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest222.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest222 proto.InternalMessageInfo

func (m *HelloRequest222) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloReply222 struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply222) Reset()         { *m = HelloReply222{} }
func (m *HelloReply222) String() string { return proto.CompactTextString(m) }
func (*HelloReply222) ProtoMessage()    {}
func (*HelloReply222) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aefce65b0a12182, []int{1}
}

func (m *HelloReply222) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply222.Unmarshal(m, b)
}
func (m *HelloReply222) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply222.Marshal(b, m, deterministic)
}
func (m *HelloReply222) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply222.Merge(m, src)
}
func (m *HelloReply222) XXX_Size() int {
	return xxx_messageInfo_HelloReply222.Size(m)
}
func (m *HelloReply222) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply222.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply222 proto.InternalMessageInfo

func (m *HelloReply222) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest222)(nil), "helloworld.HelloRequest222")
	proto.RegisterType((*HelloReply222)(nil), "helloworld.HelloReply222")
}

func init() { proto.RegisterFile("helloworld222.proto", fileDescriptor_3aefce65b0a12182) }

var fileDescriptor_3aefce65b0a12182 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0x31, 0x32, 0x32, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x42, 0x08, 0x2a, 0xa9, 0x72, 0xf1, 0x7b, 0x80, 0x78, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5,
	0x25, 0x46, 0x46, 0x46, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x26, 0x17, 0x2f, 0x54, 0x59, 0x41, 0x4e, 0x25, 0x48, 0x91,
	0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x3a, 0x4c, 0x1d, 0x8c, 0x6b, 0x14, 0xc4, 0xc5,
	0xe5, 0x5e, 0x94, 0x9a, 0x5a, 0x92, 0x5a, 0x04, 0x52, 0xe7, 0xc2, 0xc5, 0x11, 0x9c, 0x58, 0x09,
	0xd6, 0x2b, 0x24, 0xad, 0x87, 0xb0, 0x58, 0x0f, 0xcd, 0x56, 0x29, 0x49, 0x2c, 0x92, 0x10, 0xbb,
	0x94, 0x18, 0x92, 0xd8, 0xc0, 0x0e, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xea, 0xe3, 0x60,
	0x3e, 0xcf, 0x00, 0x00, 0x00,
}
