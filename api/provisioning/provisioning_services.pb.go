// Code generated by protoc-gen-go. DO NOT EDIT.
// source: provisioning/provisioning_services.proto

package provisioning

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type ProvisioningResult struct {
	Subscription         string   `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProvisioningResult) Reset()         { *m = ProvisioningResult{} }
func (m *ProvisioningResult) String() string { return proto.CompactTextString(m) }
func (*ProvisioningResult) ProtoMessage()    {}
func (*ProvisioningResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_5051cf94832cfec3, []int{0}
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

func (m *ProvisioningResult) GetSubscription() string {
	if m != nil {
		return m.Subscription
	}
	return ""
}

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
	proto.RegisterFile("provisioning/provisioning_services.proto", fileDescriptor_5051cf94832cfec3)
}

var fileDescriptor_5051cf94832cfec3 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0x41, 0x0b, 0x82, 0x40,
	0x10, 0x85, 0xb1, 0x83, 0xd1, 0xe2, 0x21, 0x36, 0x08, 0xb1, 0x8b, 0x78, 0xf2, 0xb4, 0x42, 0xfd,
	0x03, 0xa1, 0x7b, 0x6c, 0xa7, 0xba, 0x84, 0xda, 0x26, 0x03, 0xe6, 0x2c, 0x3b, 0xab, 0xe0, 0xbf,
	0x8f, 0x34, 0x61, 0xa5, 0xe3, 0x9b, 0xf7, 0xf1, 0xf8, 0x86, 0xa5, 0xda, 0x60, 0x0f, 0x04, 0xd8,
	0x42, 0x5b, 0x67, 0x6e, 0x78, 0x90, 0x32, 0x3d, 0x54, 0x8a, 0x84, 0x36, 0x68, 0x91, 0x07, 0x6e,
	0x19, 0x1d, 0x6a, 0xc4, 0xba, 0x51, 0xd9, 0xd8, 0x95, 0xdd, 0x2b, 0x53, 0x6f, 0x6d, 0x87, 0x09,
	0x4d, 0x24, 0xe3, 0x17, 0x07, 0x96, 0x8a, 0xba, 0xc6, 0xf2, 0x84, 0x05, 0xd4, 0x95, 0x54, 0x19,
	0xd0, 0x16, 0xb0, 0x0d, 0xbd, 0xd8, 0x4b, 0x37, 0x72, 0x71, 0xe3, 0x21, 0x5b, 0xeb, 0x62, 0x68,
	0xb0, 0x78, 0x86, 0xab, 0xd8, 0x4b, 0x03, 0x39, 0xc7, 0xe3, 0x8d, 0xed, 0xdc, 0xcd, 0xeb, 0x24,
	0xc7, 0x73, 0xe6, 0xff, 0xe6, 0x63, 0xe1, 0x0a, 0x8a, 0x7f, 0x81, 0x68, 0x2f, 0x26, 0x69, 0x31,
	0x4b, 0x8b, 0xf3, 0x57, 0x3a, 0xe7, 0xf7, 0x6d, 0xa1, 0x61, 0xf1, 0x7c, 0xe9, 0x8f, 0xcc, 0xe9,
	0x13, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xaa, 0x3d, 0x8b, 0x20, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProvisioningServiceClient is the client API for ProvisioningService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProvisioningServiceClient interface {
	Result(ctx context.Context, in *ProvisioningResult, opts ...grpc.CallOption) (*empty.Empty, error)
}

type provisioningServiceClient struct {
	cc *grpc.ClientConn
}

func NewProvisioningServiceClient(cc *grpc.ClientConn) ProvisioningServiceClient {
	return &provisioningServiceClient{cc}
}

func (c *provisioningServiceClient) Result(ctx context.Context, in *ProvisioningResult, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/provisioning.ProvisioningService/Result", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProvisioningServiceServer is the server API for ProvisioningService service.
type ProvisioningServiceServer interface {
	Result(context.Context, *ProvisioningResult) (*empty.Empty, error)
}

// UnimplementedProvisioningServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProvisioningServiceServer struct {
}

func (*UnimplementedProvisioningServiceServer) Result(ctx context.Context, req *ProvisioningResult) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}

func RegisterProvisioningServiceServer(s *grpc.Server, srv ProvisioningServiceServer) {
	s.RegisterService(&_ProvisioningService_serviceDesc, srv)
}

func _ProvisioningService_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvisioningResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisioningServiceServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provisioning.ProvisioningService/Result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisioningServiceServer).Result(ctx, req.(*ProvisioningResult))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProvisioningService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "provisioning.ProvisioningService",
	HandlerType: (*ProvisioningServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Result",
			Handler:    _ProvisioningService_Result_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provisioning/provisioning_services.proto",
}