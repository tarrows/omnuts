// Code generated by protoc-gen-go. DO NOT EDIT.
// source: minutes.proto

package omnuts_api_v1

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

type Backlog_BacklogType int32

const (
	Backlog_PRODUCT Backlog_BacklogType = 0
	Backlog_SPRINT  Backlog_BacklogType = 1
)

var Backlog_BacklogType_name = map[int32]string{
	0: "PRODUCT",
	1: "SPRINT",
}

var Backlog_BacklogType_value = map[string]int32{
	"PRODUCT": 0,
	"SPRINT":  1,
}

func (x Backlog_BacklogType) String() string {
	return proto.EnumName(Backlog_BacklogType_name, int32(x))
}

func (Backlog_BacklogType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e18dc4b95dc969e9, []int{0, 0}
}

type Backlog_Status int32

const (
	Backlog_NOT_STARTED Backlog_Status = 0
	Backlog_IN_PROGRESS Backlog_Status = 1
	Backlog_COMPLETED   Backlog_Status = 2
)

var Backlog_Status_name = map[int32]string{
	0: "NOT_STARTED",
	1: "IN_PROGRESS",
	2: "COMPLETED",
}

var Backlog_Status_value = map[string]int32{
	"NOT_STARTED": 0,
	"IN_PROGRESS": 1,
	"COMPLETED":   2,
}

func (x Backlog_Status) String() string {
	return proto.EnumName(Backlog_Status_name, int32(x))
}

func (Backlog_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e18dc4b95dc969e9, []int{0, 1}
}

type Backlog struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string              `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type                 Backlog_BacklogType `protobuf:"varint,3,opt,name=type,proto3,enum=omnuts.api.v1.Backlog_BacklogType" json:"type,omitempty"`
	Status               Backlog_Status      `protobuf:"varint,4,opt,name=status,proto3,enum=omnuts.api.v1.Backlog_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Backlog) Reset()         { *m = Backlog{} }
func (m *Backlog) String() string { return proto.CompactTextString(m) }
func (*Backlog) ProtoMessage()    {}
func (*Backlog) Descriptor() ([]byte, []int) {
	return fileDescriptor_e18dc4b95dc969e9, []int{0}
}

func (m *Backlog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Backlog.Unmarshal(m, b)
}
func (m *Backlog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Backlog.Marshal(b, m, deterministic)
}
func (m *Backlog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Backlog.Merge(m, src)
}
func (m *Backlog) XXX_Size() int {
	return xxx_messageInfo_Backlog.Size(m)
}
func (m *Backlog) XXX_DiscardUnknown() {
	xxx_messageInfo_Backlog.DiscardUnknown(m)
}

var xxx_messageInfo_Backlog proto.InternalMessageInfo

func (m *Backlog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Backlog) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Backlog) GetType() Backlog_BacklogType {
	if m != nil {
		return m.Type
	}
	return Backlog_PRODUCT
}

func (m *Backlog) GetStatus() Backlog_Status {
	if m != nil {
		return m.Status
	}
	return Backlog_NOT_STARTED
}

func init() {
	proto.RegisterEnum("omnuts.api.v1.Backlog_BacklogType", Backlog_BacklogType_name, Backlog_BacklogType_value)
	proto.RegisterEnum("omnuts.api.v1.Backlog_Status", Backlog_Status_name, Backlog_Status_value)
	proto.RegisterType((*Backlog)(nil), "omnuts.api.v1.Backlog")
}

func init() { proto.RegisterFile("minutes.proto", fileDescriptor_e18dc4b95dc969e9) }

var fileDescriptor_e18dc4b95dc969e9 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xcd, 0xcc, 0x2b,
	0x2d, 0x49, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcd, 0xcf, 0xcd, 0x2b, 0x2d,
	0x29, 0xd6, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x33, 0x54, 0x6a, 0x64, 0xe2, 0x62, 0x77, 0x4a, 0x4c,
	0xce, 0xce, 0xc9, 0x4f, 0x17, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0xc9, 0x2c, 0xc9, 0x49, 0x95, 0x60, 0x02,
	0x0b, 0x41, 0x38, 0x42, 0x66, 0x5c, 0x2c, 0x25, 0x95, 0x05, 0xa9, 0x12, 0xcc, 0x0a, 0x8c, 0x1a,
	0x7c, 0x46, 0x4a, 0x7a, 0x28, 0xe6, 0xe9, 0x41, 0xcd, 0x82, 0xd1, 0x21, 0x95, 0x05, 0xa9, 0x41,
	0x60, 0xf5, 0x42, 0xa6, 0x5c, 0x6c, 0xc5, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0x12, 0x2c, 0x60, 0x9d,
	0xb2, 0x38, 0x74, 0x06, 0x83, 0x15, 0x05, 0x41, 0x15, 0x2b, 0xa9, 0x71, 0x71, 0x23, 0x99, 0x25,
	0xc4, 0xcd, 0xc5, 0x1e, 0x10, 0xe4, 0xef, 0x12, 0xea, 0x1c, 0x22, 0xc0, 0x20, 0xc4, 0xc5, 0xc5,
	0x16, 0x1c, 0x10, 0xe4, 0xe9, 0x17, 0x22, 0xc0, 0xa8, 0x64, 0xc9, 0xc5, 0x06, 0xd1, 0x29, 0xc4,
	0xcf, 0xc5, 0xed, 0xe7, 0x1f, 0x12, 0x1f, 0x1c, 0xe2, 0x18, 0x14, 0xe2, 0xea, 0x22, 0xc0, 0x00,
	0x12, 0xf0, 0xf4, 0x8b, 0x0f, 0x08, 0xf2, 0x77, 0x0f, 0x72, 0x0d, 0x0e, 0x16, 0x60, 0x14, 0xe2,
	0xe5, 0xe2, 0x74, 0xf6, 0xf7, 0x0d, 0xf0, 0x71, 0x05, 0xc9, 0x33, 0x25, 0xb1, 0x81, 0x43, 0xc6,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xb7, 0x5e, 0x1e, 0x2a, 0x01, 0x00, 0x00,
}
