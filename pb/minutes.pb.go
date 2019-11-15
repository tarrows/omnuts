// Code generated by protoc-gen-go. DO NOT EDIT.
// source: minutes.proto

package omnutspb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	date "google.golang.org/genproto/googleapis/type/date"
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

type Person struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_e18dc4b95dc969e9, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Group struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Member               []*Person `protobuf:"bytes,2,rep,name=member,proto3" json:"member,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_e18dc4b95dc969e9, []int{1}
}

func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Group) GetMember() []*Person {
	if m != nil {
		return m.Member
	}
	return nil
}

type Minutes struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Date                 *date.Date           `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Attendees            []*Person            `protobuf:"bytes,4,rep,name=attendees,proto3" json:"attendees,omitempty"`
	Agenda               []*Minutes_Agendum   `protobuf:"bytes,5,rep,name=agenda,proto3" json:"agenda,omitempty"`
	Tags                 []string             `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Minutes) Reset()         { *m = Minutes{} }
func (m *Minutes) String() string { return proto.CompactTextString(m) }
func (*Minutes) ProtoMessage()    {}
func (*Minutes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e18dc4b95dc969e9, []int{2}
}

func (m *Minutes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Minutes.Unmarshal(m, b)
}
func (m *Minutes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Minutes.Marshal(b, m, deterministic)
}
func (m *Minutes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minutes.Merge(m, src)
}
func (m *Minutes) XXX_Size() int {
	return xxx_messageInfo_Minutes.Size(m)
}
func (m *Minutes) XXX_DiscardUnknown() {
	xxx_messageInfo_Minutes.DiscardUnknown(m)
}

var xxx_messageInfo_Minutes proto.InternalMessageInfo

func (m *Minutes) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Minutes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Minutes) GetDate() *date.Date {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Minutes) GetAttendees() []*Person {
	if m != nil {
		return m.Attendees
	}
	return nil
}

func (m *Minutes) GetAgenda() []*Minutes_Agendum {
	if m != nil {
		return m.Agenda
	}
	return nil
}

func (m *Minutes) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Minutes) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Minutes) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type Minutes_Agendum struct {
	No                   int32    `protobuf:"varint,1,opt,name=no,proto3" json:"no,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Minutes_Agendum) Reset()         { *m = Minutes_Agendum{} }
func (m *Minutes_Agendum) String() string { return proto.CompactTextString(m) }
func (*Minutes_Agendum) ProtoMessage()    {}
func (*Minutes_Agendum) Descriptor() ([]byte, []int) {
	return fileDescriptor_e18dc4b95dc969e9, []int{2, 0}
}

func (m *Minutes_Agendum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Minutes_Agendum.Unmarshal(m, b)
}
func (m *Minutes_Agendum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Minutes_Agendum.Marshal(b, m, deterministic)
}
func (m *Minutes_Agendum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minutes_Agendum.Merge(m, src)
}
func (m *Minutes_Agendum) XXX_Size() int {
	return xxx_messageInfo_Minutes_Agendum.Size(m)
}
func (m *Minutes_Agendum) XXX_DiscardUnknown() {
	xxx_messageInfo_Minutes_Agendum.DiscardUnknown(m)
}

var xxx_messageInfo_Minutes_Agendum proto.InternalMessageInfo

func (m *Minutes_Agendum) GetNo() int32 {
	if m != nil {
		return m.No
	}
	return 0
}

func (m *Minutes_Agendum) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Minutes_Agendum) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "omnutspb.Person")
	proto.RegisterType((*Group)(nil), "omnutspb.Group")
	proto.RegisterType((*Minutes)(nil), "omnutspb.Minutes")
	proto.RegisterType((*Minutes_Agendum)(nil), "omnutspb.Minutes.Agendum")
}

func init() { proto.RegisterFile("minutes.proto", fileDescriptor_e18dc4b95dc969e9) }

var fileDescriptor_e18dc4b95dc969e9 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x6b, 0xe3, 0x30,
	0x10, 0x85, 0xb1, 0x1d, 0x3b, 0xf1, 0x84, 0x5d, 0x76, 0xc5, 0xb2, 0x68, 0x7d, 0x59, 0x13, 0x28,
	0xf8, 0x50, 0x14, 0x9a, 0x9e, 0x7a, 0x0c, 0xb4, 0xf4, 0x54, 0x68, 0x4d, 0xef, 0x45, 0x8e, 0xa6,
	0xc6, 0x10, 0x49, 0xc6, 0x1e, 0x1f, 0xfa, 0x3f, 0xfa, 0x83, 0x8b, 0x64, 0xbb, 0xe9, 0xa5, 0xb4,
	0x37, 0x69, 0xde, 0x1b, 0xe6, 0x9b, 0x37, 0xf0, 0x43, 0x37, 0x66, 0x20, 0xec, 0x45, 0xdb, 0x59,
	0xb2, 0x6c, 0x65, 0xb5, 0x19, 0xa8, 0x6f, 0xab, 0xec, 0x7f, 0x6d, 0x6d, 0x7d, 0xc4, 0xad, 0xaf,
	0x57, 0xc3, 0xf3, 0x96, 0x1a, 0x8d, 0x3d, 0x49, 0xdd, 0x8e, 0xd6, 0xec, 0xef, 0x64, 0xa0, 0x97,
	0x16, 0xb7, 0x4a, 0x12, 0x8e, 0xf5, 0xcd, 0x39, 0x24, 0xf7, 0xd8, 0xf5, 0xd6, 0xb0, 0x9f, 0x10,
	0x36, 0x8a, 0x07, 0x79, 0x50, 0xa4, 0x65, 0xd8, 0x28, 0xc6, 0x60, 0x61, 0xa4, 0x46, 0x1e, 0xfa,
	0x8a, 0x7f, 0x6f, 0x6e, 0x20, 0xbe, 0xed, 0xec, 0xd0, 0xbe, 0x8b, 0xc1, 0x49, 0x64, 0x05, 0x24,
	0x1a, 0x75, 0x85, 0x1d, 0x0f, 0xf3, 0xa8, 0x58, 0xef, 0x7e, 0x89, 0x19, 0x4f, 0x8c, 0x23, 0xca,
	0x49, 0xdf, 0xbc, 0x46, 0xb0, 0xbc, 0x1b, 0x37, 0xf9, 0xce, 0x58, 0x76, 0x06, 0x0b, 0x87, 0xcc,
	0xa3, 0x3c, 0x28, 0xd6, 0xbb, 0xdf, 0x62, 0xdc, 0x45, 0xb8, 0x5d, 0xc4, 0xb5, 0x24, 0x2c, 0xbd,
	0xcc, 0x04, 0xa4, 0x92, 0x08, 0x8d, 0x42, 0xec, 0xf9, 0xe2, 0x13, 0x86, 0x93, 0x85, 0x5d, 0x40,
	0x22, 0x6b, 0x34, 0x4a, 0xf2, 0xd8, 0x9b, 0xff, 0x9d, 0xcc, 0x13, 0x9d, 0xd8, 0x3b, 0x7d, 0xd0,
	0xe5, 0x64, 0x74, 0x74, 0x24, 0xeb, 0x9e, 0x27, 0x79, 0xe4, 0xe8, 0xdc, 0x9b, 0x5d, 0x01, 0x1c,
	0x3a, 0x94, 0x84, 0xea, 0x49, 0x12, 0x5f, 0x7a, 0xc6, 0x6c, 0x66, 0x9c, 0x0f, 0x22, 0x1e, 0xe7,
	0x83, 0x94, 0xe9, 0xe4, 0xde, 0x93, 0x6b, 0x1d, 0x5a, 0x35, 0xb7, 0xae, 0xbe, 0x6e, 0x9d, 0xdc,
	0x7b, 0xca, 0x1e, 0x60, 0x39, 0xc1, 0xb9, 0x08, 0x8d, 0xf5, 0x11, 0xc6, 0x65, 0x68, 0x2c, 0xfb,
	0x03, 0x31, 0x35, 0x74, 0x9c, 0x33, 0x1c, 0x3f, 0x2c, 0x87, 0xb5, 0xc2, 0xfe, 0xd0, 0x35, 0x2d,
	0x35, 0xd6, 0xf8, 0x2c, 0xd3, 0xf2, 0x63, 0xa9, 0x4a, 0xfc, 0xc4, 0xcb, 0xb7, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xb6, 0xf6, 0x3d, 0x33, 0x66, 0x02, 0x00, 0x00,
}