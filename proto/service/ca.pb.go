// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ca.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CSRData struct {
	Addresses            []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Domains              []string `protobuf:"bytes,2,rep,name=domains,proto3" json:"domains,omitempty"`
	PublicKey            []byte   `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Subject              string   `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CSRData) Reset()         { *m = CSRData{} }
func (m *CSRData) String() string { return proto.CompactTextString(m) }
func (*CSRData) ProtoMessage()    {}
func (*CSRData) Descriptor() ([]byte, []int) {
	return fileDescriptor_84034e4847c57e85, []int{0}
}

func (m *CSRData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CSRData.Unmarshal(m, b)
}
func (m *CSRData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CSRData.Marshal(b, m, deterministic)
}
func (m *CSRData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CSRData.Merge(m, src)
}
func (m *CSRData) XXX_Size() int {
	return xxx_messageInfo_CSRData.Size(m)
}
func (m *CSRData) XXX_DiscardUnknown() {
	xxx_messageInfo_CSRData.DiscardUnknown(m)
}

var xxx_messageInfo_CSRData proto.InternalMessageInfo

func (m *CSRData) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *CSRData) GetDomains() []string {
	if m != nil {
		return m.Domains
	}
	return nil
}

func (m *CSRData) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *CSRData) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

type SignCertificateRequest struct {
	Csr                  *CSRData `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignCertificateRequest) Reset()         { *m = SignCertificateRequest{} }
func (m *SignCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*SignCertificateRequest) ProtoMessage()    {}
func (*SignCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_84034e4847c57e85, []int{1}
}

func (m *SignCertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignCertificateRequest.Unmarshal(m, b)
}
func (m *SignCertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignCertificateRequest.Marshal(b, m, deterministic)
}
func (m *SignCertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignCertificateRequest.Merge(m, src)
}
func (m *SignCertificateRequest) XXX_Size() int {
	return xxx_messageInfo_SignCertificateRequest.Size(m)
}
func (m *SignCertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignCertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignCertificateRequest proto.InternalMessageInfo

func (m *SignCertificateRequest) GetCsr() *CSRData {
	if m != nil {
		return m.Csr
	}
	return nil
}

type SignCertificateResponse struct {
	RawCertificate       []byte   `protobuf:"bytes,1,opt,name=raw_certificate,json=rawCertificate,proto3" json:"raw_certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignCertificateResponse) Reset()         { *m = SignCertificateResponse{} }
func (m *SignCertificateResponse) String() string { return proto.CompactTextString(m) }
func (*SignCertificateResponse) ProtoMessage()    {}
func (*SignCertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84034e4847c57e85, []int{2}
}

func (m *SignCertificateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignCertificateResponse.Unmarshal(m, b)
}
func (m *SignCertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignCertificateResponse.Marshal(b, m, deterministic)
}
func (m *SignCertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignCertificateResponse.Merge(m, src)
}
func (m *SignCertificateResponse) XXX_Size() int {
	return xxx_messageInfo_SignCertificateResponse.Size(m)
}
func (m *SignCertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignCertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignCertificateResponse proto.InternalMessageInfo

func (m *SignCertificateResponse) GetRawCertificate() []byte {
	if m != nil {
		return m.RawCertificate
	}
	return nil
}

func init() {
	proto.RegisterType((*CSRData)(nil), "pb.CSRData")
	proto.RegisterType((*SignCertificateRequest)(nil), "pb.SignCertificateRequest")
	proto.RegisterType((*SignCertificateResponse)(nil), "pb.SignCertificateResponse")
}

func init() { proto.RegisterFile("ca.proto", fileDescriptor_84034e4847c57e85) }

var fileDescriptor_84034e4847c57e85 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4a, 0x03, 0x31,
	0x18, 0x85, 0x49, 0x47, 0xac, 0xfd, 0x5b, 0x2c, 0x04, 0xd4, 0x50, 0x5b, 0x18, 0x66, 0x63, 0x71,
	0xd1, 0x81, 0xba, 0x10, 0x5c, 0x5a, 0x77, 0xee, 0x32, 0x07, 0x28, 0xff, 0x64, 0xe2, 0x18, 0xad,
	0x49, 0x4c, 0x32, 0x94, 0xba, 0xf4, 0x0a, 0x1e, 0xcd, 0x2b, 0x78, 0x10, 0x99, 0x99, 0x4a, 0x45,
	0xbb, 0xfc, 0xdf, 0xcb, 0xf7, 0xde, 0x23, 0x70, 0x24, 0x70, 0x66, 0x9d, 0x09, 0x86, 0x76, 0x6c,
	0x3e, 0x1a, 0x97, 0xc6, 0x94, 0x2b, 0x99, 0xa2, 0x55, 0x29, 0x6a, 0x6d, 0x02, 0x06, 0x65, 0xb4,
	0x6f, 0x5f, 0x24, 0x6f, 0xd0, 0x5d, 0x64, 0xfc, 0x0e, 0x03, 0xd2, 0x31, 0xf4, 0xb0, 0x28, 0x9c,
	0xf4, 0x5e, 0x7a, 0x46, 0xe2, 0x68, 0xda, 0xe3, 0x3b, 0x81, 0x32, 0xe8, 0x16, 0xe6, 0x05, 0x95,
	0xf6, 0xac, 0xd3, 0x78, 0x3f, 0x27, 0x9d, 0x00, 0xd8, 0x2a, 0x5f, 0x29, 0xb1, 0x7c, 0x96, 0x1b,
	0x16, 0xc5, 0x64, 0x3a, 0xe0, 0xbd, 0x56, 0xb9, 0x97, 0x9b, 0x1a, 0xf4, 0x55, 0xfe, 0x24, 0x45,
	0x60, 0x07, 0x31, 0xa9, 0xc1, 0xed, 0x99, 0x5c, 0xc3, 0x69, 0xa6, 0x4a, 0xbd, 0x90, 0x2e, 0xa8,
	0x07, 0x25, 0x30, 0x48, 0x2e, 0x5f, 0x2b, 0xe9, 0x03, 0x9d, 0x40, 0x24, 0xbc, 0x63, 0x24, 0x26,
	0xd3, 0xfe, 0xbc, 0x3f, 0xb3, 0xf9, 0x6c, 0x3b, 0x92, 0xd7, 0x7a, 0x72, 0x0b, 0x67, 0xff, 0x40,
	0x6f, 0x8d, 0xf6, 0x92, 0x5e, 0xc0, 0xd0, 0xe1, 0x7a, 0x29, 0x76, 0x56, 0x93, 0x32, 0xe0, 0xc7,
	0x0e, 0xd7, 0xbf, 0x80, 0xf9, 0x23, 0x44, 0x8b, 0x8c, 0x53, 0x84, 0xe1, 0x9f, 0x28, 0x3a, 0xaa,
	0xfb, 0xf6, 0x0f, 0x1b, 0x9d, 0xef, 0xf5, 0xda, 0xee, 0xe4, 0xe4, 0xfd, 0xf3, 0xeb, 0xa3, 0x33,
	0x4c, 0x20, 0xad, 0xeb, 0x53, 0xaf, 0x4a, 0x7d, 0x43, 0x2e, 0xf3, 0xc3, 0xe6, 0xa7, 0xaf, 0xbe,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x9d, 0xed, 0x40, 0x97, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CSRClient is the client API for CSR service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CSRClient interface {
	SignCertificate(ctx context.Context, in *SignCertificateRequest, opts ...grpc.CallOption) (*SignCertificateResponse, error)
}

type cSRClient struct {
	cc grpc.ClientConnInterface
}

func NewCSRClient(cc grpc.ClientConnInterface) CSRClient {
	return &cSRClient{cc}
}

func (c *cSRClient) SignCertificate(ctx context.Context, in *SignCertificateRequest, opts ...grpc.CallOption) (*SignCertificateResponse, error) {
	out := new(SignCertificateResponse)
	err := c.cc.Invoke(ctx, "/pb.CSR/SignCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CSRServer is the server API for CSR service.
type CSRServer interface {
	SignCertificate(context.Context, *SignCertificateRequest) (*SignCertificateResponse, error)
}

// UnimplementedCSRServer can be embedded to have forward compatible implementations.
type UnimplementedCSRServer struct {
}

func (*UnimplementedCSRServer) SignCertificate(ctx context.Context, req *SignCertificateRequest) (*SignCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignCertificate not implemented")
}

func RegisterCSRServer(s *grpc.Server, srv CSRServer) {
	s.RegisterService(&_CSR_serviceDesc, srv)
}

func _CSR_SignCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CSRServer).SignCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CSR/SignCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CSRServer).SignCertificate(ctx, req.(*SignCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CSR_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CSR",
	HandlerType: (*CSRServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignCertificate",
			Handler:    _CSR_SignCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ca.proto",
}
