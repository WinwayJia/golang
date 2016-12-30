// Code generated by protoc-gen-go.
// source: im.proto
// DO NOT EDIT!

/*
Package im is a generated protocol buffer package.

It is generated from these files:
	im.proto

It has these top-level messages:
	XMessageHeader
	XLoginReq
	XLoginResp
*/
package im

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import base "."

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type XMessageType int32

const (
	XMessageType_Login  XMessageType = 0
	XMessageType_Logout XMessageType = 1
	XMessageType_Ping   XMessageType = 2
	XMessageType_Pong   XMessageType = 3
	XMessageType_Chat   XMessageType = 4
)

var XMessageType_name = map[int32]string{
	0: "Login",
	1: "Logout",
	2: "Ping",
	3: "Pong",
	4: "Chat",
}
var XMessageType_value = map[string]int32{
	"Login":  0,
	"Logout": 1,
	"Ping":   2,
	"Pong":   3,
	"Chat":   4,
}

func (x XMessageType) Enum() *XMessageType {
	p := new(XMessageType)
	*p = x
	return p
}
func (x XMessageType) String() string {
	return proto.EnumName(XMessageType_name, int32(x))
}
func (x *XMessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(XMessageType_value, data, "XMessageType")
	if err != nil {
		return err
	}
	*x = XMessageType(value)
	return nil
}
func (XMessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type XMessageHeader struct {
	Type             *XMessageType `protobuf:"varint,1,opt,name=type,enum=im.XMessageType" json:"type,omitempty"`
	Token            *string       `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *XMessageHeader) Reset()                    { *m = XMessageHeader{} }
func (m *XMessageHeader) String() string            { return proto.CompactTextString(m) }
func (*XMessageHeader) ProtoMessage()               {}
func (*XMessageHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *XMessageHeader) GetType() XMessageType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return XMessageType_Login
}

func (m *XMessageHeader) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

type XLoginReq struct {
	UserID           *int64 `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *XLoginReq) Reset()                    { *m = XLoginReq{} }
func (m *XLoginReq) String() string            { return proto.CompactTextString(m) }
func (*XLoginReq) ProtoMessage()               {}
func (*XLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *XLoginReq) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

type XLoginResp struct {
	Token            *string       `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Errmsg           *base.ErrInfo `protobuf:"bytes,2,opt,name=errmsg" json:"errmsg,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *XLoginResp) Reset()                    { *m = XLoginResp{} }
func (m *XLoginResp) String() string            { return proto.CompactTextString(m) }
func (*XLoginResp) ProtoMessage()               {}
func (*XLoginResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *XLoginResp) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *XLoginResp) GetErrmsg() *base.ErrInfo {
	if m != nil {
		return m.Errmsg
	}
	return nil
}

func init() {
	proto.RegisterType((*XMessageHeader)(nil), "im.XMessageHeader")
	proto.RegisterType((*XLoginReq)(nil), "im.XLoginReq")
	proto.RegisterType((*XLoginResp)(nil), "im.XLoginResp")
	proto.RegisterEnum("im.XMessageType", XMessageType_name, XMessageType_value)
}

func init() { proto.RegisterFile("im.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x8e, 0x4b, 0x4b, 0xc4, 0x30,
	0x10, 0xc7, 0xed, 0x93, 0x66, 0xb4, 0x25, 0xe4, 0x54, 0x14, 0x45, 0x7a, 0x12, 0x0f, 0x3d, 0xf4,
	0xe8, 0xc5, 0x43, 0x15, 0x2c, 0x28, 0x88, 0x28, 0xf4, 0x1a, 0x71, 0x8c, 0x41, 0x9a, 0x74, 0x93,
	0xec, 0x61, 0xbf, 0xfd, 0xa6, 0x59, 0xca, 0xb2, 0xb7, 0xdf, 0xcc, 0xff, 0x31, 0x03, 0x85, 0x9c,
	0xda, 0xd9, 0x68, 0xa7, 0x59, 0x2c, 0xa7, 0x4b, 0xf8, 0xe6, 0x16, 0x0f, 0x73, 0xf3, 0x08, 0xd5,
	0xf8, 0x86, 0xd6, 0x72, 0x81, 0x2f, 0xc8, 0x7f, 0xd0, 0xb0, 0x1b, 0x48, 0xdd, 0x6e, 0xc6, 0x3a,
	0xba, 0x8d, 0xee, 0xaa, 0x8e, 0xb6, 0x3e, 0xba, 0x3a, 0x3e, 0xfd, 0x9e, 0x95, 0x90, 0x39, 0xfd,
	0x8f, 0xaa, 0x8e, 0xbd, 0x81, 0x34, 0x57, 0x40, 0xc6, 0x57, 0x2d, 0xa4, 0xfa, 0xc0, 0x0d, 0xab,
	0x20, 0xff, 0xb2, 0x68, 0x86, 0xa7, 0x90, 0x4e, 0x9a, 0x07, 0x80, 0x55, 0xb4, 0xf3, 0x31, 0xb9,
	0x88, 0x84, 0x5d, 0x43, 0x8e, 0xc6, 0x4c, 0x56, 0x84, 0xa6, 0xf3, 0xae, 0x6c, 0xc3, 0x5f, 0xcf,
	0xc6, 0x0c, 0xea, 0x57, 0xdf, 0xf7, 0x70, 0x71, 0x72, 0x97, 0x40, 0x16, 0xaa, 0xe8, 0x19, 0x03,
	0xc8, 0x3d, 0xea, 0xad, 0xa3, 0x11, 0x2b, 0x20, 0x7d, 0x97, 0x4a, 0xd0, 0x38, 0x90, 0xf6, 0x94,
	0x2c, 0xd4, 0xff, 0x71, 0x47, 0xd3, 0x7d, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x9a, 0x9d, 0x5d,
	0xf9, 0x00, 0x00, 0x00,
}
