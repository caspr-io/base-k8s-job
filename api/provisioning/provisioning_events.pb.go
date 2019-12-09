// Code generated by protoc-gen-go. DO NOT EDIT.
// source: provisioning/provisioning_events.proto

package provisioning

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
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

type ProvisioningResult struct {
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProvisioningResult) Reset()         { *m = ProvisioningResult{} }
func (m *ProvisioningResult) String() string { return proto.CompactTextString(m) }
func (*ProvisioningResult) ProtoMessage()    {}
func (*ProvisioningResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b05476ba6149fe7e, []int{0}
}

func (m *ProvisioningResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProvisioningResult.Unmarshal(m, b)
}
func (m *ProvisioningResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProvisioningResult.Marshal(b, m, deterministic)
}
func (m *ProvisioningResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProvisioningResult.Merge(m, src)
}
func (m *ProvisioningResult) XXX_Size() int {
	return xxx_messageInfo_ProvisioningResult.Size(m)
}
func (m *ProvisioningResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ProvisioningResult.DiscardUnknown(m)
}

var xxx_messageInfo_ProvisioningResult proto.InternalMessageInfo

func (m *ProvisioningResult) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*ProvisioningResult)(nil), "provisioning.ProvisioningResult")
}

func init() {
	proto.RegisterFile("provisioning/provisioning_events.proto", fileDescriptor_b05476ba6149fe7e)
}

var fileDescriptor_b05476ba6149fe7e = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x28, 0xca, 0x2f,
	0xcb, 0x2c, 0xce, 0xcc, 0xcf, 0xcb, 0xcc, 0x4b, 0xd7, 0x47, 0xe6, 0xc4, 0xa7, 0x96, 0xa5, 0xe6,
	0x95, 0x14, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0x20, 0x4b, 0x49, 0x49, 0xa6, 0xe7,
	0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0xe5, 0x92, 0x4a, 0xd3, 0xf4, 0x13, 0xf3, 0x2a, 0x21, 0x0a,
	0x95, 0xf4, 0xb8, 0x84, 0x02, 0x90, 0x94, 0x06, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0x08, 0x49, 0x70,
	0xb1, 0x17, 0x24, 0x56, 0xe6, 0xe4, 0x27, 0xa6, 0x48, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x04, 0xc1,
	0xb8, 0x4e, 0x42, 0x51, 0x02, 0x89, 0x05, 0x99, 0x28, 0x36, 0x27, 0xb1, 0x81, 0x8d, 0x32, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x08, 0xb8, 0xc8, 0x35, 0x9d, 0x00, 0x00, 0x00,
}