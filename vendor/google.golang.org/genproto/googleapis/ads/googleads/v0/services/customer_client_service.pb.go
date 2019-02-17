// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/customer_client_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [CustomerClientService.GetCustomerClient][google.ads.googleads.v0.services.CustomerClientService.GetCustomerClient].
type GetCustomerClientRequest struct {
	// The resource name of the customer client to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerClientRequest) Reset()         { *m = GetCustomerClientRequest{} }
func (m *GetCustomerClientRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerClientRequest) ProtoMessage()    {}
func (*GetCustomerClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_client_service_528cdd682e78fa39, []int{0}
}
func (m *GetCustomerClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerClientRequest.Unmarshal(m, b)
}
func (m *GetCustomerClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerClientRequest.Marshal(b, m, deterministic)
}
func (dst *GetCustomerClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerClientRequest.Merge(dst, src)
}
func (m *GetCustomerClientRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerClientRequest.Size(m)
}
func (m *GetCustomerClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerClientRequest proto.InternalMessageInfo

func (m *GetCustomerClientRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCustomerClientRequest)(nil), "google.ads.googleads.v0.services.GetCustomerClientRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerClientServiceClient is the client API for CustomerClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClientServiceClient interface {
	// Returns the requested customer client in full detail.
	GetCustomerClient(ctx context.Context, in *GetCustomerClientRequest, opts ...grpc.CallOption) (*resources.CustomerClient, error)
}

type customerClientServiceClient struct {
	cc *grpc.ClientConn
}

func NewCustomerClientServiceClient(cc *grpc.ClientConn) CustomerClientServiceClient {
	return &customerClientServiceClient{cc}
}

func (c *customerClientServiceClient) GetCustomerClient(ctx context.Context, in *GetCustomerClientRequest, opts ...grpc.CallOption) (*resources.CustomerClient, error) {
	out := new(resources.CustomerClient)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.CustomerClientService/GetCustomerClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerClientServiceServer is the server API for CustomerClientService service.
type CustomerClientServiceServer interface {
	// Returns the requested customer client in full detail.
	GetCustomerClient(context.Context, *GetCustomerClientRequest) (*resources.CustomerClient, error)
}

func RegisterCustomerClientServiceServer(s *grpc.Server, srv CustomerClientServiceServer) {
	s.RegisterService(&_CustomerClientService_serviceDesc, srv)
}

func _CustomerClientService_GetCustomerClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerClientServiceServer).GetCustomerClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.CustomerClientService/GetCustomerClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerClientServiceServer).GetCustomerClient(ctx, req.(*GetCustomerClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.CustomerClientService",
	HandlerType: (*CustomerClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerClient",
			Handler:    _CustomerClientService_GetCustomerClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/customer_client_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/customer_client_service.proto", fileDescriptor_customer_client_service_528cdd682e78fa39)
}

var fileDescriptor_customer_client_service_528cdd682e78fa39 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xe2, 0xd4,
	0xa2, 0xb2, 0xcc, 0xe4, 0xd4, 0x62, 0xfd, 0xe4, 0xd2, 0xe2, 0x92, 0xfc, 0xdc, 0xd4, 0xa2, 0xf8,
	0xe4, 0x9c, 0xcc, 0xd4, 0xbc, 0x92, 0x78, 0xa8, 0x84, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90,
	0x02, 0x44, 0x93, 0x5e, 0x62, 0x4a, 0xb1, 0x1e, 0x5c, 0xbf, 0x5e, 0x99, 0x81, 0x1e, 0x4c, 0xbf,
	0x94, 0x39, 0x2e, 0x1b, 0x8a, 0x52, 0x8b, 0xf3, 0x4b, 0x8b, 0xb0, 0x58, 0x01, 0x31, 0x5a, 0x4a,
	0x06, 0xa6, 0xb1, 0x20, 0x53, 0x3f, 0x31, 0x2f, 0x2f, 0xbf, 0x24, 0xb1, 0x24, 0x33, 0x3f, 0xaf,
	0x18, 0x22, 0xab, 0x64, 0xcf, 0x25, 0xe1, 0x9e, 0x5a, 0xe2, 0x0c, 0xd5, 0xe9, 0x0c, 0xd6, 0x18,
	0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0xa4, 0xcc, 0xc5, 0x0b, 0x33, 0x3c, 0x3e, 0x2f, 0x31,
	0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x07, 0x26, 0xe8, 0x97, 0x98, 0x9b, 0x6a,
	0x74, 0x9d, 0x91, 0x4b, 0x14, 0x55, 0x7b, 0x30, 0xc4, 0xc9, 0x42, 0x7b, 0x19, 0xb9, 0x04, 0x31,
	0xcc, 0x16, 0xb2, 0xd2, 0x23, 0xe4, 0x55, 0x3d, 0x5c, 0x0e, 0x92, 0x32, 0xc4, 0xa9, 0x17, 0x1e,
	0x08, 0x7a, 0xa8, 0x3a, 0x95, 0x2c, 0x9b, 0x2e, 0x3f, 0x99, 0xcc, 0x64, 0x2c, 0x64, 0x08, 0x0a,
	0xaa, 0x6a, 0x14, 0xef, 0xd8, 0xc2, 0xc2, 0xab, 0x58, 0x5f, 0x0b, 0x1e, 0x76, 0x10, 0x6d, 0xc5,
	0xfa, 0x5a, 0xb5, 0x4e, 0xb7, 0x18, 0xb9, 0x54, 0x92, 0xf3, 0x73, 0x09, 0xba, 0xd7, 0x49, 0x0a,
	0xab, 0xff, 0x03, 0x40, 0xe1, 0x1b, 0xc0, 0x18, 0xe5, 0x01, 0xd5, 0x9f, 0x9e, 0x9f, 0x93, 0x98,
	0x97, 0xae, 0x97, 0x5f, 0x94, 0xae, 0x9f, 0x9e, 0x9a, 0x07, 0x0e, 0x7d, 0x58, 0x44, 0x16, 0x64,
	0x16, 0xe3, 0x4e, 0x39, 0xd6, 0x30, 0xc6, 0x22, 0x26, 0x66, 0x77, 0x47, 0xc7, 0x55, 0x4c, 0x0a,
	0xee, 0x10, 0x03, 0x1d, 0x53, 0x8a, 0xf5, 0x20, 0x4c, 0x10, 0x2b, 0xcc, 0x40, 0x0f, 0x6a, 0x71,
	0xf1, 0x29, 0x98, 0x92, 0x18, 0xc7, 0x94, 0xe2, 0x18, 0xb8, 0x92, 0x98, 0x30, 0x83, 0x18, 0x98,
	0x92, 0x24, 0x36, 0xb0, 0x03, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xbb, 0x69, 0xb1,
	0xb9, 0x02, 0x00, 0x00,
}