// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authority.proto

package authoritypb

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

type ServiceTypeFlag int32

const (
	ServiceTypeFlag_Unknown          ServiceTypeFlag = 0
	ServiceTypeFlag_ServiceProvider  ServiceTypeFlag = 1
	ServiceTypeFlag_IdentityProvider ServiceTypeFlag = 2
)

var ServiceTypeFlag_name = map[int32]string{
	0: "Unknown",
	1: "ServiceProvider",
	2: "IdentityProvider",
}

var ServiceTypeFlag_value = map[string]int32{
	"Unknown":          0,
	"ServiceProvider":  1,
	"IdentityProvider": 2,
}

func (x ServiceTypeFlag) String() string {
	return proto.EnumName(ServiceTypeFlag_name, int32(x))
}

func (ServiceTypeFlag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a5a0640cd66a638, []int{0}
}

type Application struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TypeFlags            ServiceTypeFlag `protobuf:"varint,2,opt,name=type_flags,json=typeFlags,proto3,enum=authoritypb.ServiceTypeFlag" json:"type_flags,omitempty"`
	Label                string          `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Description          string          `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CreateAt             int64           `protobuf:"varint,5,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5a0640cd66a638, []int{0}
}

func (m *Application) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Application.Unmarshal(m, b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Application.Marshal(b, m, deterministic)
}
func (m *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(m, src)
}
func (m *Application) XXX_Size() int {
	return xxx_messageInfo_Application.Size(m)
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Application) GetTypeFlags() ServiceTypeFlag {
	if m != nil {
		return m.TypeFlags
	}
	return ServiceTypeFlag_Unknown
}

func (m *Application) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Application) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Application) GetCreateAt() int64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

type Authentication struct {
	AppName              string   `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	Nonce                string   `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Challenge            string   `protobuf:"bytes,3,opt,name=challenge,proto3" json:"challenge,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Authentication) Reset()         { *m = Authentication{} }
func (m *Authentication) String() string { return proto.CompactTextString(m) }
func (*Authentication) ProtoMessage()    {}
func (*Authentication) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5a0640cd66a638, []int{1}
}

func (m *Authentication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authentication.Unmarshal(m, b)
}
func (m *Authentication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authentication.Marshal(b, m, deterministic)
}
func (m *Authentication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authentication.Merge(m, src)
}
func (m *Authentication) XXX_Size() int {
	return xxx_messageInfo_Authentication.Size(m)
}
func (m *Authentication) XXX_DiscardUnknown() {
	xxx_messageInfo_Authentication.DiscardUnknown(m)
}

var xxx_messageInfo_Authentication proto.InternalMessageInfo

func (m *Authentication) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *Authentication) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *Authentication) GetChallenge() string {
	if m != nil {
		return m.Challenge
	}
	return ""
}

type AppCredentials struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AccessKey            string   `protobuf:"bytes,2,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	AccessSecret         string   `protobuf:"bytes,3,opt,name=access_secret,json=accessSecret,proto3" json:"access_secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppCredentials) Reset()         { *m = AppCredentials{} }
func (m *AppCredentials) String() string { return proto.CompactTextString(m) }
func (*AppCredentials) ProtoMessage()    {}
func (*AppCredentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5a0640cd66a638, []int{2}
}

func (m *AppCredentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppCredentials.Unmarshal(m, b)
}
func (m *AppCredentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppCredentials.Marshal(b, m, deterministic)
}
func (m *AppCredentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppCredentials.Merge(m, src)
}
func (m *AppCredentials) XXX_Size() int {
	return xxx_messageInfo_AppCredentials.Size(m)
}
func (m *AppCredentials) XXX_DiscardUnknown() {
	xxx_messageInfo_AppCredentials.DiscardUnknown(m)
}

var xxx_messageInfo_AppCredentials proto.InternalMessageInfo

func (m *AppCredentials) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AppCredentials) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *AppCredentials) GetAccessSecret() string {
	if m != nil {
		return m.AccessSecret
	}
	return ""
}

type CertificateTemplate struct {
	Addresses            []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Domains              []string `protobuf:"bytes,2,rep,name=domains,proto3" json:"domains,omitempty"`
	PublicKey            []byte   `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	ServiceName          string   `protobuf:"bytes,4,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateTemplate) Reset()         { *m = CertificateTemplate{} }
func (m *CertificateTemplate) String() string { return proto.CompactTextString(m) }
func (*CertificateTemplate) ProtoMessage()    {}
func (*CertificateTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5a0640cd66a638, []int{3}
}

func (m *CertificateTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateTemplate.Unmarshal(m, b)
}
func (m *CertificateTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateTemplate.Marshal(b, m, deterministic)
}
func (m *CertificateTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateTemplate.Merge(m, src)
}
func (m *CertificateTemplate) XXX_Size() int {
	return xxx_messageInfo_CertificateTemplate.Size(m)
}
func (m *CertificateTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateTemplate proto.InternalMessageInfo

func (m *CertificateTemplate) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *CertificateTemplate) GetDomains() []string {
	if m != nil {
		return m.Domains
	}
	return nil
}

func (m *CertificateTemplate) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *CertificateTemplate) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

type RegisterApplicationRequest struct {
	Application          *Application `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RegisterApplicationRequest) Reset()         { *m = RegisterApplicationRequest{} }
func (m *RegisterApplicationRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterApplicationRequest) ProtoMessage()    {}
func (*RegisterApplicationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5a0640cd66a638, []int{4}
}

func (m *RegisterApplicationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterApplicationRequest.Unmarshal(m, b)
}
func (m *RegisterApplicationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterApplicationRequest.Marshal(b, m, deterministic)
}
func (m *RegisterApplicationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterApplicationRequest.Merge(m, src)
}
func (m *RegisterApplicationRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterApplicationRequest.Size(m)
}
func (m *RegisterApplicationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterApplicationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterApplicationRequest proto.InternalMessageInfo

func (m *RegisterApplicationRequest) GetApplication() *Application {
	if m != nil {
		return m.Application
	}
	return nil
}

type RegisterApplicationResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterApplicationResponse) Reset()         { *m = RegisterApplicationResponse{} }
func (m *RegisterApplicationResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterApplicationResponse) ProtoMessage()    {}
func (*RegisterApplicationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5a0640cd66a638, []int{5}
}

func (m *RegisterApplicationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterApplicationResponse.Unmarshal(m, b)
}
func (m *RegisterApplicationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterApplicationResponse.Marshal(b, m, deterministic)
}
func (m *RegisterApplicationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterApplicationResponse.Merge(m, src)
}
func (m *RegisterApplicationResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterApplicationResponse.Size(m)
}
func (m *RegisterApplicationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterApplicationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterApplicationResponse proto.InternalMessageInfo

type SignKeyRequest struct {
	KeyBytes             []byte   `protobuf:"bytes,1,opt,name=key_bytes,json=keyBytes,proto3" json:"key_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignKeyRequest) Reset()         { *m = SignKeyRequest{} }
func (m *SignKeyRequest) String() string { return proto.CompactTextString(m) }
func (*SignKeyRequest) ProtoMessage()    {}
func (*SignKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5a0640cd66a638, []int{6}
}

func (m *SignKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignKeyRequest.Unmarshal(m, b)
}
func (m *SignKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignKeyRequest.Marshal(b, m, deterministic)
}
func (m *SignKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignKeyRequest.Merge(m, src)
}
func (m *SignKeyRequest) XXX_Size() int {
	return xxx_messageInfo_SignKeyRequest.Size(m)
}
func (m *SignKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignKeyRequest proto.InternalMessageInfo

func (m *SignKeyRequest) GetKeyBytes() []byte {
	if m != nil {
		return m.KeyBytes
	}
	return nil
}

type SignKeyResponse struct {
	Signature            []byte   `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignKeyResponse) Reset()         { *m = SignKeyResponse{} }
func (m *SignKeyResponse) String() string { return proto.CompactTextString(m) }
func (*SignKeyResponse) ProtoMessage()    {}
func (*SignKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5a0640cd66a638, []int{7}
}

func (m *SignKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignKeyResponse.Unmarshal(m, b)
}
func (m *SignKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignKeyResponse.Marshal(b, m, deterministic)
}
func (m *SignKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignKeyResponse.Merge(m, src)
}
func (m *SignKeyResponse) XXX_Size() int {
	return xxx_messageInfo_SignKeyResponse.Size(m)
}
func (m *SignKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignKeyResponse proto.InternalMessageInfo

func (m *SignKeyResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type SignCertificateRequest struct {
	Template             *CertificateTemplate `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SignCertificateRequest) Reset()         { *m = SignCertificateRequest{} }
func (m *SignCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*SignCertificateRequest) ProtoMessage()    {}
func (*SignCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5a0640cd66a638, []int{8}
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

func (m *SignCertificateRequest) GetTemplate() *CertificateTemplate {
	if m != nil {
		return m.Template
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
	return fileDescriptor_6a5a0640cd66a638, []int{9}
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

type ValidateAuthRequest struct {
	Authentication       *Authentication `protobuf:"bytes,1,opt,name=authentication,proto3" json:"authentication,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ValidateAuthRequest) Reset()         { *m = ValidateAuthRequest{} }
func (m *ValidateAuthRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateAuthRequest) ProtoMessage()    {}
func (*ValidateAuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5a0640cd66a638, []int{10}
}

func (m *ValidateAuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateAuthRequest.Unmarshal(m, b)
}
func (m *ValidateAuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateAuthRequest.Marshal(b, m, deterministic)
}
func (m *ValidateAuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateAuthRequest.Merge(m, src)
}
func (m *ValidateAuthRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateAuthRequest.Size(m)
}
func (m *ValidateAuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateAuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateAuthRequest proto.InternalMessageInfo

func (m *ValidateAuthRequest) GetAuthentication() *Authentication {
	if m != nil {
		return m.Authentication
	}
	return nil
}

type ValidateAuthResponse struct {
	Nonce                string   `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateAuthResponse) Reset()         { *m = ValidateAuthResponse{} }
func (m *ValidateAuthResponse) String() string { return proto.CompactTextString(m) }
func (*ValidateAuthResponse) ProtoMessage()    {}
func (*ValidateAuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5a0640cd66a638, []int{11}
}

func (m *ValidateAuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateAuthResponse.Unmarshal(m, b)
}
func (m *ValidateAuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateAuthResponse.Marshal(b, m, deterministic)
}
func (m *ValidateAuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateAuthResponse.Merge(m, src)
}
func (m *ValidateAuthResponse) XXX_Size() int {
	return xxx_messageInfo_ValidateAuthResponse.Size(m)
}
func (m *ValidateAuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateAuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateAuthResponse proto.InternalMessageInfo

func (m *ValidateAuthResponse) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *ValidateAuthResponse) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterEnum("authoritypb.ServiceTypeFlag", ServiceTypeFlag_name, ServiceTypeFlag_value)
	proto.RegisterType((*Application)(nil), "authoritypb.Application")
	proto.RegisterType((*Authentication)(nil), "authoritypb.Authentication")
	proto.RegisterType((*AppCredentials)(nil), "authoritypb.AppCredentials")
	proto.RegisterType((*CertificateTemplate)(nil), "authoritypb.CertificateTemplate")
	proto.RegisterType((*RegisterApplicationRequest)(nil), "authoritypb.RegisterApplicationRequest")
	proto.RegisterType((*RegisterApplicationResponse)(nil), "authoritypb.RegisterApplicationResponse")
	proto.RegisterType((*SignKeyRequest)(nil), "authoritypb.SignKeyRequest")
	proto.RegisterType((*SignKeyResponse)(nil), "authoritypb.SignKeyResponse")
	proto.RegisterType((*SignCertificateRequest)(nil), "authoritypb.SignCertificateRequest")
	proto.RegisterType((*SignCertificateResponse)(nil), "authoritypb.SignCertificateResponse")
	proto.RegisterType((*ValidateAuthRequest)(nil), "authoritypb.ValidateAuthRequest")
	proto.RegisterType((*ValidateAuthResponse)(nil), "authoritypb.ValidateAuthResponse")
}

func init() { proto.RegisterFile("authority.proto", fileDescriptor_6a5a0640cd66a638) }

var fileDescriptor_6a5a0640cd66a638 = []byte{
	// 771 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xcd, 0x6e, 0xf3, 0x44,
	0x14, 0xc5, 0xcd, 0x17, 0x12, 0xdf, 0x84, 0x24, 0x9a, 0xe4, 0xe3, 0x73, 0x93, 0x54, 0xa4, 0x2e,
	0x52, 0xa3, 0x4a, 0x34, 0xa2, 0xec, 0x0a, 0x9b, 0xb4, 0x12, 0x52, 0x55, 0x09, 0x21, 0xb7, 0x54,
	0x88, 0x4d, 0x34, 0x71, 0x6e, 0x1d, 0x2b, 0xae, 0x3d, 0xcc, 0x4c, 0x52, 0x79, 0xc3, 0x82, 0x1d,
	0x2b, 0x16, 0x3c, 0x0a, 0x8f, 0xc2, 0x2b, 0xf0, 0x0a, 0xec, 0xd1, 0x8c, 0xc7, 0x89, 0xdd, 0x86,
	0xb2, 0xf3, 0x9c, 0xfb, 0x73, 0xce, 0xbd, 0x73, 0x26, 0x81, 0x36, 0x5d, 0xcb, 0x65, 0xc2, 0x43,
	0x99, 0x9e, 0x33, 0x9e, 0xc8, 0x84, 0x34, 0xb6, 0x00, 0x9b, 0xf7, 0x87, 0x41, 0x92, 0x04, 0x11,
	0x4e, 0x28, 0x0b, 0x27, 0x34, 0x8e, 0x13, 0x49, 0x65, 0x98, 0xc4, 0x22, 0x4b, 0x75, 0xff, 0xb4,
	0xa0, 0x31, 0x65, 0x2c, 0x0a, 0x7d, 0x0d, 0x13, 0x02, 0xef, 0x62, 0xfa, 0x84, 0x8e, 0x35, 0xb2,
	0xc6, 0xb6, 0xa7, 0xbf, 0xc9, 0xd7, 0x00, 0x32, 0x65, 0x38, 0x7b, 0x8c, 0x68, 0x20, 0x9c, 0x83,
	0x91, 0x35, 0x6e, 0x5d, 0x0c, 0xcf, 0x0b, 0x1c, 0xe7, 0x77, 0xc8, 0x37, 0xa1, 0x8f, 0xf7, 0x29,
	0xc3, 0x6f, 0x23, 0x1a, 0x78, 0xb6, 0x34, 0x5f, 0x82, 0xf4, 0xa0, 0x1a, 0xd1, 0x39, 0x46, 0x4e,
	0x45, 0x77, 0xcc, 0x0e, 0x64, 0x04, 0x8d, 0x05, 0x0a, 0x9f, 0x87, 0x4c, 0xb1, 0x3a, 0xef, 0x74,
	0xac, 0x08, 0x91, 0x01, 0xd8, 0x3e, 0x47, 0x2a, 0x71, 0x46, 0xa5, 0x53, 0x1d, 0x59, 0xe3, 0x8a,
	0x57, 0xcf, 0x80, 0xa9, 0x74, 0x67, 0xd0, 0x9a, 0xae, 0xe5, 0x12, 0x63, 0x99, 0xeb, 0x3e, 0x84,
	0x3a, 0x65, 0x6c, 0x56, 0xd0, 0x5e, 0xa3, 0x8c, 0x7d, 0xa7, 0xe4, 0xf7, 0xa0, 0x1a, 0x27, 0xb1,
	0x8f, 0x5a, 0xb9, 0xed, 0x65, 0x07, 0x32, 0x04, 0xdb, 0x5f, 0xd2, 0x28, 0xc2, 0x38, 0x40, 0xa3,
	0x6d, 0x07, 0xb8, 0x4b, 0x68, 0x4d, 0x19, 0xbb, 0xe6, 0xb8, 0x50, 0x1c, 0x34, 0x12, 0x7b, 0x17,
	0x73, 0x04, 0x40, 0x7d, 0x1f, 0x85, 0x98, 0xad, 0x30, 0x35, 0xed, 0xed, 0x0c, 0xb9, 0xc5, 0x94,
	0x9c, 0xc0, 0x27, 0x26, 0x2c, 0xd0, 0xe7, 0x28, 0x0d, 0x4d, 0x33, 0x03, 0xef, 0x34, 0xe6, 0xfe,
	0x6e, 0x41, 0xf7, 0x1a, 0xb9, 0x0c, 0x1f, 0xd5, 0x28, 0x78, 0x8f, 0x4f, 0x2c, 0xa2, 0x52, 0xeb,
	0xa3, 0x8b, 0x05, 0x47, 0x21, 0x50, 0x38, 0xd6, 0xa8, 0xa2, 0x5b, 0xe7, 0x00, 0x71, 0xa0, 0xb6,
	0x48, 0x9e, 0x68, 0x18, 0xab, 0xfb, 0x50, 0xb1, 0xfc, 0xa8, 0x34, 0xb1, 0xf5, 0x3c, 0x0a, 0x7d,
	0xad, 0x49, 0x31, 0x36, 0x3d, 0x3b, 0x43, 0x94, 0xa6, 0x63, 0x68, 0x8a, 0xec, 0xb2, 0xb2, 0x5d,
	0x99, 0xcd, 0x1b, 0x4c, 0xed, 0xcb, 0xfd, 0x11, 0xfa, 0x1e, 0x06, 0xa1, 0x90, 0xc8, 0x0b, 0xce,
	0xf0, 0xf0, 0xe7, 0x35, 0x0a, 0x49, 0x2e, 0xa1, 0x41, 0x77, 0xa8, 0x5e, 0x47, 0xe3, 0xc2, 0x29,
	0xb9, 0xa1, 0x58, 0x55, 0x4c, 0x76, 0x8f, 0x60, 0xb0, 0xb7, 0xb3, 0x60, 0x49, 0x2c, 0xd0, 0xfd,
	0x02, 0x5a, 0x77, 0x61, 0x10, 0xdf, 0x62, 0x9a, 0x93, 0x0d, 0xc0, 0x5e, 0x61, 0x3a, 0x9b, 0xa7,
	0x52, 0x2f, 0x41, 0xcd, 0x52, 0x5f, 0x61, 0x7a, 0xa5, 0xce, 0xee, 0x04, 0xda, 0xdb, 0xf4, 0xac,
	0x83, 0x5a, 0x9a, 0x08, 0x83, 0x98, 0xca, 0x35, 0x47, 0x93, 0xbf, 0x03, 0xdc, 0x07, 0xf8, 0x54,
	0x15, 0x14, 0xb6, 0x9d, 0xf3, 0x7c, 0x03, 0x75, 0x69, 0x16, 0x6f, 0x26, 0x1a, 0x95, 0x26, 0xda,
	0x73, 0x41, 0xde, 0xb6, 0xc2, 0xbd, 0x82, 0x0f, 0xaf, 0xfa, 0x1a, 0x41, 0xa7, 0xd0, 0xe6, 0xf4,
	0x79, 0xe6, 0xef, 0x42, 0x46, 0x56, 0x8b, 0xd3, 0xe7, 0x42, 0x81, 0xfb, 0x13, 0x74, 0x1f, 0x68,
	0x14, 0x2e, 0x94, 0xbf, 0xd7, 0x72, 0x99, 0x0b, 0xbb, 0x86, 0x16, 0x2d, 0x19, 0xdd, 0xc8, 0x1b,
	0x94, 0x17, 0x5e, 0x4a, 0xf1, 0x5e, 0x94, 0xb8, 0x57, 0xd0, 0x2b, 0xf7, 0x36, 0xe2, 0xb6, 0x0f,
	0xc3, 0x2a, 0x3e, 0x8c, 0x1e, 0x54, 0x37, 0x2a, 0x5b, 0xfb, 0xb9, 0xee, 0x65, 0x87, 0xb3, 0x1b,
	0x68, 0xbf, 0x78, 0xe4, 0xa4, 0x01, 0xb5, 0x1f, 0xe2, 0x55, 0x9c, 0x3c, 0xc7, 0x9d, 0x8f, 0x48,
	0x77, 0x1b, 0xff, 0x9e, 0x27, 0x9b, 0x70, 0x81, 0xbc, 0x63, 0x91, 0x1e, 0x74, 0x6e, 0xf4, 0x03,
	0x92, 0xe9, 0x16, 0x3d, 0xb8, 0xf8, 0xa7, 0x02, 0x9d, 0x69, 0xae, 0xde, 0x14, 0x91, 0xdf, 0x2c,
	0xe8, 0xee, 0xf1, 0x06, 0x39, 0x2d, 0x0d, 0xfa, 0xdf, 0xbe, 0xec, 0x8f, 0xff, 0x3f, 0xd1, 0xd8,
	0xec, 0xb3, 0x5f, 0xff, 0xfa, 0xfb, 0x8f, 0x83, 0x43, 0xb7, 0xa7, 0x7f, 0x12, 0x37, 0x5f, 0x4e,
	0x28, 0x63, 0x13, 0x6e, 0x0a, 0x2e, 0xad, 0x33, 0x42, 0xa1, 0x66, 0x8c, 0x45, 0xca, 0x7b, 0x2e,
	0xbb, 0xb3, 0x3f, 0xdc, 0x1f, 0x34, 0x34, 0x03, 0x4d, 0xf3, 0xde, 0xed, 0xe4, 0x34, 0xca, 0x88,
	0x93, 0x15, 0xa6, 0x8a, 0xe2, 0x97, 0xcc, 0xbb, 0x05, 0x07, 0x90, 0x93, 0x57, 0xdd, 0x5e, 0x1b,
	0xb5, 0xff, 0xf9, 0xdb, 0x49, 0x86, 0x7a, 0xa4, 0xa9, 0xfb, 0xee, 0xfb, 0xe2, 0x84, 0x9a, 0x5e,
	0x19, 0x51, 0xf1, 0x73, 0x68, 0x16, 0x2d, 0x41, 0xca, 0x76, 0xdf, 0xe3, 0xc4, 0xfe, 0xf1, 0x1b,
	0x19, 0x86, 0xf6, 0x48, 0xd3, 0x7e, 0x20, 0x3b, 0xda, 0xb5, 0x5c, 0x4e, 0x36, 0x26, 0x75, 0xfe,
	0xb1, 0xfe, 0xc7, 0xf9, 0xea, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xa1, 0x41, 0x7f, 0xaf,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorityServiceClient is the client API for AuthorityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorityServiceClient interface {
	RegisterApplication(ctx context.Context, in *RegisterApplicationRequest, opts ...grpc.CallOption) (*RegisterApplicationResponse, error)
	SignKey(ctx context.Context, in *SignKeyRequest, opts ...grpc.CallOption) (*SignKeyResponse, error)
	SignCertificate(ctx context.Context, in *SignCertificateRequest, opts ...grpc.CallOption) (*SignCertificateResponse, error)
	ValidateAuth(ctx context.Context, in *ValidateAuthRequest, opts ...grpc.CallOption) (*ValidateAuthResponse, error)
}

type authorityServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthorityServiceClient(cc *grpc.ClientConn) AuthorityServiceClient {
	return &authorityServiceClient{cc}
}

func (c *authorityServiceClient) RegisterApplication(ctx context.Context, in *RegisterApplicationRequest, opts ...grpc.CallOption) (*RegisterApplicationResponse, error) {
	out := new(RegisterApplicationResponse)
	err := c.cc.Invoke(ctx, "/authoritypb.AuthorityService/RegisterApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityServiceClient) SignKey(ctx context.Context, in *SignKeyRequest, opts ...grpc.CallOption) (*SignKeyResponse, error) {
	out := new(SignKeyResponse)
	err := c.cc.Invoke(ctx, "/authoritypb.AuthorityService/SignKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityServiceClient) SignCertificate(ctx context.Context, in *SignCertificateRequest, opts ...grpc.CallOption) (*SignCertificateResponse, error) {
	out := new(SignCertificateResponse)
	err := c.cc.Invoke(ctx, "/authoritypb.AuthorityService/SignCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityServiceClient) ValidateAuth(ctx context.Context, in *ValidateAuthRequest, opts ...grpc.CallOption) (*ValidateAuthResponse, error) {
	out := new(ValidateAuthResponse)
	err := c.cc.Invoke(ctx, "/authoritypb.AuthorityService/ValidateAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorityServiceServer is the server API for AuthorityService service.
type AuthorityServiceServer interface {
	RegisterApplication(context.Context, *RegisterApplicationRequest) (*RegisterApplicationResponse, error)
	SignKey(context.Context, *SignKeyRequest) (*SignKeyResponse, error)
	SignCertificate(context.Context, *SignCertificateRequest) (*SignCertificateResponse, error)
	ValidateAuth(context.Context, *ValidateAuthRequest) (*ValidateAuthResponse, error)
}

// UnimplementedAuthorityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorityServiceServer struct {
}

func (*UnimplementedAuthorityServiceServer) RegisterApplication(ctx context.Context, req *RegisterApplicationRequest) (*RegisterApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterApplication not implemented")
}
func (*UnimplementedAuthorityServiceServer) SignKey(ctx context.Context, req *SignKeyRequest) (*SignKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignKey not implemented")
}
func (*UnimplementedAuthorityServiceServer) SignCertificate(ctx context.Context, req *SignCertificateRequest) (*SignCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignCertificate not implemented")
}
func (*UnimplementedAuthorityServiceServer) ValidateAuth(ctx context.Context, req *ValidateAuthRequest) (*ValidateAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAuth not implemented")
}

func RegisterAuthorityServiceServer(s *grpc.Server, srv AuthorityServiceServer) {
	s.RegisterService(&_AuthorityService_serviceDesc, srv)
}

func _AuthorityService_RegisterApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).RegisterApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authoritypb.AuthorityService/RegisterApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).RegisterApplication(ctx, req.(*RegisterApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityService_SignKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).SignKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authoritypb.AuthorityService/SignKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).SignKey(ctx, req.(*SignKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityService_SignCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).SignCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authoritypb.AuthorityService/SignCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).SignCertificate(ctx, req.(*SignCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityService_ValidateAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).ValidateAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authoritypb.AuthorityService/ValidateAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).ValidateAuth(ctx, req.(*ValidateAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthorityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authoritypb.AuthorityService",
	HandlerType: (*AuthorityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterApplication",
			Handler:    _AuthorityService_RegisterApplication_Handler,
		},
		{
			MethodName: "SignKey",
			Handler:    _AuthorityService_SignKey_Handler,
		},
		{
			MethodName: "SignCertificate",
			Handler:    _AuthorityService_SignCertificate_Handler,
		},
		{
			MethodName: "ValidateAuth",
			Handler:    _AuthorityService_ValidateAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authority.proto",
}
