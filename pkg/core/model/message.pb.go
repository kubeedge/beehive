// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/core/model/message.proto

package model

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MessageHeader struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ParentID             string   `protobuf:"bytes,2,opt,name=ParentID,proto3" json:"ParentID,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Sync                 bool     `protobuf:"varint,4,opt,name=Sync,proto3" json:"Sync,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageHeader) Reset()         { *m = MessageHeader{} }
func (m *MessageHeader) String() string { return proto.CompactTextString(m) }
func (*MessageHeader) ProtoMessage()    {}
func (*MessageHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d87fc56ef8153b7, []int{0}
}
func (m *MessageHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageHeader.Unmarshal(m, b)
}
func (m *MessageHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageHeader.Marshal(b, m, deterministic)
}
func (m *MessageHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageHeader.Merge(m, src)
}
func (m *MessageHeader) XXX_Size() int {
	return xxx_messageInfo_MessageHeader.Size(m)
}
func (m *MessageHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageHeader.DiscardUnknown(m)
}

var xxx_messageInfo_MessageHeader proto.InternalMessageInfo

func (m *MessageHeader) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *MessageHeader) GetParentID() string {
	if m != nil {
		return m.ParentID
	}
	return ""
}

func (m *MessageHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MessageHeader) GetSync() bool {
	if m != nil {
		return m.Sync
	}
	return false
}

type MessageRoute struct {
	Source               string     `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
	Group                string     `protobuf:"bytes,2,opt,name=Group,proto3" json:"Group,omitempty"`
	Operation            string     `protobuf:"bytes,3,opt,name=Operation,proto3" json:"Operation,omitempty"`
	Resource             *types.Any `protobuf:"bytes,4,opt,name=Resource,proto3" json:"Resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MessageRoute) Reset()         { *m = MessageRoute{} }
func (m *MessageRoute) String() string { return proto.CompactTextString(m) }
func (*MessageRoute) ProtoMessage()    {}
func (*MessageRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d87fc56ef8153b7, []int{1}
}
func (m *MessageRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageRoute.Unmarshal(m, b)
}
func (m *MessageRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageRoute.Marshal(b, m, deterministic)
}
func (m *MessageRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRoute.Merge(m, src)
}
func (m *MessageRoute) XXX_Size() int {
	return xxx_messageInfo_MessageRoute.Size(m)
}
func (m *MessageRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRoute.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRoute proto.InternalMessageInfo

func (m *MessageRoute) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *MessageRoute) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *MessageRoute) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *MessageRoute) GetResource() *types.Any {
	if m != nil {
		return m.Resource
	}
	return nil
}

type Message struct {
	Header               *MessageHeader `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	Router               *MessageRoute  `protobuf:"bytes,2,opt,name=Router,proto3" json:"Router,omitempty"`
	Data                 *types.Any     `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d87fc56ef8153b7, []int{2}
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

func (m *Message) GetHeader() *MessageHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Message) GetRouter() *MessageRoute {
	if m != nil {
		return m.Router
	}
	return nil
}

func (m *Message) GetData() *types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageHeader)(nil), "model.MessageHeader")
	proto.RegisterType((*MessageRoute)(nil), "model.MessageRoute")
	proto.RegisterType((*Message)(nil), "model.Message")
}

func init() { proto.RegisterFile("pkg/core/model/message.proto", fileDescriptor_1d87fc56ef8153b7) }

var fileDescriptor_1d87fc56ef8153b7 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x99, 0xfe, 0xc9, 0xd7, 0xde, 0x7e, 0xba, 0xb8, 0x16, 0x89, 0xa5, 0x8b, 0xd2, 0x55,
	0x40, 0x49, 0xa4, 0x3e, 0x81, 0x10, 0xd0, 0x2e, 0x44, 0x99, 0xfa, 0x02, 0xd3, 0xf4, 0x1a, 0x8a,
	0x4d, 0x26, 0x4c, 0x26, 0x8b, 0xbc, 0x83, 0xf8, 0xcc, 0x92, 0x3b, 0x63, 0x4b, 0x17, 0xee, 0x72,
	0x4f, 0xce, 0x9d, 0xdf, 0xb9, 0x07, 0xe6, 0xd5, 0x67, 0x9e, 0x64, 0xda, 0x50, 0x52, 0xe8, 0x1d,
	0x1d, 0x92, 0x82, 0xea, 0x5a, 0xe5, 0x14, 0x57, 0x46, 0x5b, 0x8d, 0x43, 0x16, 0x67, 0x37, 0xb9,
	0xd6, 0xf9, 0x81, 0x12, 0x16, 0xb7, 0xcd, 0x47, 0xa2, 0xca, 0xd6, 0x39, 0x96, 0x05, 0x5c, 0xbc,
	0xb8, 0x95, 0x67, 0x52, 0x3b, 0x32, 0x78, 0x09, 0xbd, 0x75, 0x1a, 0x8a, 0x85, 0x88, 0xc6, 0xb2,
	0xb7, 0x4e, 0x71, 0x06, 0xa3, 0x37, 0x65, 0xa8, 0xb4, 0xeb, 0x34, 0xec, 0xb1, 0x7a, 0x9c, 0x71,
	0x0e, 0xe3, 0xf7, 0x7d, 0x41, 0xb5, 0x55, 0x45, 0x15, 0xf6, 0x17, 0x22, 0xea, 0xcb, 0x93, 0x80,
	0x08, 0x83, 0x4d, 0x5b, 0x66, 0xe1, 0x60, 0x21, 0xa2, 0x91, 0xe4, 0xef, 0xe5, 0x97, 0x80, 0xff,
	0x9e, 0x27, 0x75, 0x63, 0x09, 0xaf, 0x21, 0xd8, 0xe8, 0xc6, 0x64, 0xe4, 0x91, 0x7e, 0xc2, 0x29,
	0x0c, 0x9f, 0x8c, 0x6e, 0x2a, 0xcf, 0x74, 0x43, 0x07, 0x7c, 0xad, 0xc8, 0x28, 0xbb, 0xd7, 0x25,
	0x03, 0xc7, 0xf2, 0x24, 0xe0, 0x3d, 0x8c, 0x24, 0xd5, 0xee, 0xb5, 0x0e, 0x3a, 0x59, 0x4d, 0x63,
	0x77, 0x79, 0xfc, 0x7b, 0x79, 0xfc, 0x58, 0xb6, 0xf2, 0xe8, 0x5a, 0x7e, 0x0b, 0xf8, 0xe7, 0xe3,
	0xe0, 0x1d, 0x04, 0xae, 0x02, 0x4e, 0xd2, 0xed, 0x72, 0x79, 0xf1, 0x59, 0x3d, 0xd2, 0x7b, 0xf0,
	0x16, 0x02, 0x3e, 0xc0, 0x70, 0xc0, 0xc9, 0xea, 0xea, 0xdc, 0xcd, 0xff, 0xa4, 0xb7, 0x60, 0x04,
	0x83, 0x54, 0x59, 0xc5, 0x89, 0xff, 0x0a, 0xc5, 0x8e, 0x6d, 0xc0, 0xda, 0xc3, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xf2, 0x3c, 0x7a, 0xde, 0xd7, 0x01, 0x00, 0x00,
}
