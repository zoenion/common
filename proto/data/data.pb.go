// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

package datapb

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

type Scope int32

const (
	Scope_Public            Scope = 0
	Scope_Personal          Scope = 1
	Scope_ContextRestricted Scope = 2
)

var Scope_name = map[int32]string{
	0: "Public",
	1: "Personal",
	2: "ContextRestricted",
}

var Scope_value = map[string]int32{
	"Public":            0,
	"Personal":          1,
	"ContextRestricted": 2,
}

func (x Scope) String() string {
	return proto.EnumName(Scope_name, int32(x))
}

func (Scope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

type Definition struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Mime                 string   `protobuf:"bytes,4,opt,name=mime,proto3" json:"mime,omitempty"`
	Context              string   `protobuf:"bytes,5,opt,name=context,proto3" json:"context,omitempty"`
	Scope                Scope    `protobuf:"varint,6,opt,name=scope,proto3,enum=datapb.Scope" json:"scope,omitempty"`
	Categories           []string `protobuf:"bytes,7,rep,name=categories,proto3" json:"categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Definition) Reset()         { *m = Definition{} }
func (m *Definition) String() string { return proto.CompactTextString(m) }
func (*Definition) ProtoMessage()    {}
func (*Definition) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

func (m *Definition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Definition.Unmarshal(m, b)
}
func (m *Definition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Definition.Marshal(b, m, deterministic)
}
func (m *Definition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Definition.Merge(m, src)
}
func (m *Definition) XXX_Size() int {
	return xxx_messageInfo_Definition.Size(m)
}
func (m *Definition) XXX_DiscardUnknown() {
	xxx_messageInfo_Definition.DiscardUnknown(m)
}

var xxx_messageInfo_Definition proto.InternalMessageInfo

func (m *Definition) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Definition) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Definition) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Definition) GetMime() string {
	if m != nil {
		return m.Mime
	}
	return ""
}

func (m *Definition) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func (m *Definition) GetScope() Scope {
	if m != nil {
		return m.Scope
	}
	return Scope_Public
}

func (m *Definition) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

type Data struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Context              string   `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	Uri                  string   `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{1}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Data) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func (m *Data) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type Query struct {
	Subject              string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Keys                 []string `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{2}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Query) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type SetDataRequest struct {
	Subject              string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Data                 []*Data  `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDataRequest) Reset()         { *m = SetDataRequest{} }
func (m *SetDataRequest) String() string { return proto.CompactTextString(m) }
func (*SetDataRequest) ProtoMessage()    {}
func (*SetDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{3}
}

func (m *SetDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDataRequest.Unmarshal(m, b)
}
func (m *SetDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDataRequest.Marshal(b, m, deterministic)
}
func (m *SetDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDataRequest.Merge(m, src)
}
func (m *SetDataRequest) XXX_Size() int {
	return xxx_messageInfo_SetDataRequest.Size(m)
}
func (m *SetDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDataRequest proto.InternalMessageInfo

func (m *SetDataRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *SetDataRequest) GetData() []*Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type SetDataResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDataResponse) Reset()         { *m = SetDataResponse{} }
func (m *SetDataResponse) String() string { return proto.CompactTextString(m) }
func (*SetDataResponse) ProtoMessage()    {}
func (*SetDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{4}
}

func (m *SetDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDataResponse.Unmarshal(m, b)
}
func (m *SetDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDataResponse.Marshal(b, m, deterministic)
}
func (m *SetDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDataResponse.Merge(m, src)
}
func (m *SetDataResponse) XXX_Size() int {
	return xxx_messageInfo_SetDataResponse.Size(m)
}
func (m *SetDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetDataResponse proto.InternalMessageInfo

type GetDataRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDataRequest) Reset()         { *m = GetDataRequest{} }
func (m *GetDataRequest) String() string { return proto.CompactTextString(m) }
func (*GetDataRequest) ProtoMessage()    {}
func (*GetDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{5}
}

func (m *GetDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataRequest.Unmarshal(m, b)
}
func (m *GetDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataRequest.Marshal(b, m, deterministic)
}
func (m *GetDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataRequest.Merge(m, src)
}
func (m *GetDataRequest) XXX_Size() int {
	return xxx_messageInfo_GetDataRequest.Size(m)
}
func (m *GetDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataRequest proto.InternalMessageInfo

func (m *GetDataRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetDataResponse struct {
	Definition           *Definition `protobuf:"bytes,1,opt,name=definition,proto3" json:"definition,omitempty"`
	Value                *Data       `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetDataResponse) Reset()         { *m = GetDataResponse{} }
func (m *GetDataResponse) String() string { return proto.CompactTextString(m) }
func (*GetDataResponse) ProtoMessage()    {}
func (*GetDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{6}
}

func (m *GetDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataResponse.Unmarshal(m, b)
}
func (m *GetDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataResponse.Marshal(b, m, deterministic)
}
func (m *GetDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataResponse.Merge(m, src)
}
func (m *GetDataResponse) XXX_Size() int {
	return xxx_messageInfo_GetDataResponse.Size(m)
}
func (m *GetDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataResponse proto.InternalMessageInfo

func (m *GetDataResponse) GetDefinition() *Definition {
	if m != nil {
		return m.Definition
	}
	return nil
}

func (m *GetDataResponse) GetValue() *Data {
	if m != nil {
		return m.Value
	}
	return nil
}

type DeleteDataRequest struct {
	Query                *Query   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDataRequest) Reset()         { *m = DeleteDataRequest{} }
func (m *DeleteDataRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDataRequest) ProtoMessage()    {}
func (*DeleteDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{7}
}

func (m *DeleteDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDataRequest.Unmarshal(m, b)
}
func (m *DeleteDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDataRequest.Marshal(b, m, deterministic)
}
func (m *DeleteDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDataRequest.Merge(m, src)
}
func (m *DeleteDataRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDataRequest.Size(m)
}
func (m *DeleteDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDataRequest proto.InternalMessageInfo

func (m *DeleteDataRequest) GetQuery() *Query {
	if m != nil {
		return m.Query
	}
	return nil
}

type DeleteDataResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDataResponse) Reset()         { *m = DeleteDataResponse{} }
func (m *DeleteDataResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteDataResponse) ProtoMessage()    {}
func (*DeleteDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{8}
}

func (m *DeleteDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDataResponse.Unmarshal(m, b)
}
func (m *DeleteDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDataResponse.Marshal(b, m, deterministic)
}
func (m *DeleteDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDataResponse.Merge(m, src)
}
func (m *DeleteDataResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteDataResponse.Size(m)
}
func (m *DeleteDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDataResponse proto.InternalMessageInfo

type SetDefinitionRequest struct {
	Definition           *Definition `protobuf:"bytes,1,opt,name=definition,proto3" json:"definition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SetDefinitionRequest) Reset()         { *m = SetDefinitionRequest{} }
func (m *SetDefinitionRequest) String() string { return proto.CompactTextString(m) }
func (*SetDefinitionRequest) ProtoMessage()    {}
func (*SetDefinitionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{9}
}

func (m *SetDefinitionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDefinitionRequest.Unmarshal(m, b)
}
func (m *SetDefinitionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDefinitionRequest.Marshal(b, m, deterministic)
}
func (m *SetDefinitionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDefinitionRequest.Merge(m, src)
}
func (m *SetDefinitionRequest) XXX_Size() int {
	return xxx_messageInfo_SetDefinitionRequest.Size(m)
}
func (m *SetDefinitionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDefinitionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDefinitionRequest proto.InternalMessageInfo

func (m *SetDefinitionRequest) GetDefinition() *Definition {
	if m != nil {
		return m.Definition
	}
	return nil
}

type SetDefinitionResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDefinitionResponse) Reset()         { *m = SetDefinitionResponse{} }
func (m *SetDefinitionResponse) String() string { return proto.CompactTextString(m) }
func (*SetDefinitionResponse) ProtoMessage()    {}
func (*SetDefinitionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{10}
}

func (m *SetDefinitionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDefinitionResponse.Unmarshal(m, b)
}
func (m *SetDefinitionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDefinitionResponse.Marshal(b, m, deterministic)
}
func (m *SetDefinitionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDefinitionResponse.Merge(m, src)
}
func (m *SetDefinitionResponse) XXX_Size() int {
	return xxx_messageInfo_SetDefinitionResponse.Size(m)
}
func (m *SetDefinitionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDefinitionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetDefinitionResponse proto.InternalMessageInfo

type GetDefinitionRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDefinitionRequest) Reset()         { *m = GetDefinitionRequest{} }
func (m *GetDefinitionRequest) String() string { return proto.CompactTextString(m) }
func (*GetDefinitionRequest) ProtoMessage()    {}
func (*GetDefinitionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{11}
}

func (m *GetDefinitionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDefinitionRequest.Unmarshal(m, b)
}
func (m *GetDefinitionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDefinitionRequest.Marshal(b, m, deterministic)
}
func (m *GetDefinitionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDefinitionRequest.Merge(m, src)
}
func (m *GetDefinitionRequest) XXX_Size() int {
	return xxx_messageInfo_GetDefinitionRequest.Size(m)
}
func (m *GetDefinitionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDefinitionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDefinitionRequest proto.InternalMessageInfo

func (m *GetDefinitionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetDefinitionResponse struct {
	Definition           *Definition `protobuf:"bytes,1,opt,name=definition,proto3" json:"definition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetDefinitionResponse) Reset()         { *m = GetDefinitionResponse{} }
func (m *GetDefinitionResponse) String() string { return proto.CompactTextString(m) }
func (*GetDefinitionResponse) ProtoMessage()    {}
func (*GetDefinitionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{12}
}

func (m *GetDefinitionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDefinitionResponse.Unmarshal(m, b)
}
func (m *GetDefinitionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDefinitionResponse.Marshal(b, m, deterministic)
}
func (m *GetDefinitionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDefinitionResponse.Merge(m, src)
}
func (m *GetDefinitionResponse) XXX_Size() int {
	return xxx_messageInfo_GetDefinitionResponse.Size(m)
}
func (m *GetDefinitionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDefinitionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDefinitionResponse proto.InternalMessageInfo

func (m *GetDefinitionResponse) GetDefinition() *Definition {
	if m != nil {
		return m.Definition
	}
	return nil
}

type DeleteDefinitionRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDefinitionRequest) Reset()         { *m = DeleteDefinitionRequest{} }
func (m *DeleteDefinitionRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDefinitionRequest) ProtoMessage()    {}
func (*DeleteDefinitionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{13}
}

func (m *DeleteDefinitionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDefinitionRequest.Unmarshal(m, b)
}
func (m *DeleteDefinitionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDefinitionRequest.Marshal(b, m, deterministic)
}
func (m *DeleteDefinitionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDefinitionRequest.Merge(m, src)
}
func (m *DeleteDefinitionRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDefinitionRequest.Size(m)
}
func (m *DeleteDefinitionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDefinitionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDefinitionRequest proto.InternalMessageInfo

func (m *DeleteDefinitionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteDefinitionResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDefinitionResponse) Reset()         { *m = DeleteDefinitionResponse{} }
func (m *DeleteDefinitionResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteDefinitionResponse) ProtoMessage()    {}
func (*DeleteDefinitionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{14}
}

func (m *DeleteDefinitionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDefinitionResponse.Unmarshal(m, b)
}
func (m *DeleteDefinitionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDefinitionResponse.Marshal(b, m, deterministic)
}
func (m *DeleteDefinitionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDefinitionResponse.Merge(m, src)
}
func (m *DeleteDefinitionResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteDefinitionResponse.Size(m)
}
func (m *DeleteDefinitionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDefinitionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDefinitionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("datapb.Scope", Scope_name, Scope_value)
	proto.RegisterType((*Definition)(nil), "datapb.Definition")
	proto.RegisterType((*Data)(nil), "datapb.Data")
	proto.RegisterType((*Query)(nil), "datapb.Query")
	proto.RegisterType((*SetDataRequest)(nil), "datapb.SetDataRequest")
	proto.RegisterType((*SetDataResponse)(nil), "datapb.SetDataResponse")
	proto.RegisterType((*GetDataRequest)(nil), "datapb.GetDataRequest")
	proto.RegisterType((*GetDataResponse)(nil), "datapb.GetDataResponse")
	proto.RegisterType((*DeleteDataRequest)(nil), "datapb.DeleteDataRequest")
	proto.RegisterType((*DeleteDataResponse)(nil), "datapb.DeleteDataResponse")
	proto.RegisterType((*SetDefinitionRequest)(nil), "datapb.SetDefinitionRequest")
	proto.RegisterType((*SetDefinitionResponse)(nil), "datapb.SetDefinitionResponse")
	proto.RegisterType((*GetDefinitionRequest)(nil), "datapb.GetDefinitionRequest")
	proto.RegisterType((*GetDefinitionResponse)(nil), "datapb.GetDefinitionResponse")
	proto.RegisterType((*DeleteDefinitionRequest)(nil), "datapb.DeleteDefinitionRequest")
	proto.RegisterType((*DeleteDefinitionResponse)(nil), "datapb.DeleteDefinitionResponse")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor_871986018790d2fd) }

var fileDescriptor_871986018790d2fd = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x26, 0x6e, 0x9c, 0x92, 0x69, 0x9b, 0x26, 0xa3, 0x94, 0x98, 0xa8, 0x80, 0xb5, 0xbd, 0x54,
	0x91, 0x68, 0xa4, 0x20, 0xa4, 0x8a, 0x23, 0x54, 0xb2, 0x04, 0x1c, 0x8a, 0x73, 0xe5, 0xe2, 0xd8,
	0x93, 0x68, 0xc1, 0xf5, 0xa6, 0xf6, 0x06, 0x51, 0x21, 0x2e, 0xbc, 0x02, 0x0f, 0x85, 0xc4, 0x95,
	0x57, 0xe0, 0x41, 0xd0, 0xae, 0x37, 0xa9, 0xed, 0xa4, 0x11, 0xea, 0x6d, 0x76, 0x7e, 0xbe, 0x6f,
	0xbe, 0x99, 0x71, 0x02, 0x10, 0x05, 0x32, 0x38, 0x9b, 0xa7, 0x42, 0x0a, 0x6c, 0x28, 0x7b, 0x3e,
	0xe9, 0x1f, 0xcf, 0x84, 0x98, 0xc5, 0x34, 0x0c, 0xe6, 0x7c, 0x18, 0x24, 0x89, 0x90, 0x81, 0xe4,
	0x22, 0xc9, 0xf2, 0x2c, 0xf6, 0xab, 0x06, 0x70, 0x41, 0x53, 0x9e, 0x70, 0xe5, 0xc5, 0x16, 0x58,
	0x3c, 0x72, 0x6a, 0x6e, 0xed, 0xb4, 0xe9, 0x5b, 0x3c, 0xc2, 0x2e, 0xd8, 0x71, 0x30, 0xa1, 0xd8,
	0xb1, 0xb4, 0x2b, 0x7f, 0xa0, 0x0b, 0x7b, 0x11, 0x65, 0x61, 0xca, 0xe7, 0xaa, 0xc8, 0xd9, 0xd1,
	0xb1, 0xa2, 0x0b, 0x11, 0xea, 0x57, 0xfc, 0x8a, 0x9c, 0xba, 0x0e, 0x69, 0x1b, 0x1d, 0xd8, 0x0d,
	0x45, 0x22, 0xe9, 0xab, 0x74, 0x6c, 0xed, 0x5e, 0x3e, 0xf1, 0x04, 0xec, 0x2c, 0x14, 0x73, 0x72,
	0x1a, 0x6e, 0xed, 0xb4, 0x35, 0x3a, 0x38, 0xcb, 0x5b, 0x3f, 0x1b, 0x2b, 0xa7, 0x9f, 0xc7, 0xf0,
	0x29, 0x40, 0x18, 0x48, 0x9a, 0x89, 0x94, 0x53, 0xe6, 0xec, 0xba, 0x3b, 0xa7, 0x4d, 0xbf, 0xe0,
	0x61, 0xaf, 0xa1, 0x7e, 0x11, 0xc8, 0x60, 0x4d, 0x42, 0x81, 0xd6, 0x2a, 0xd3, 0xb6, 0x61, 0x67,
	0x91, 0x72, 0xd3, 0xbe, 0x32, 0xd9, 0x4b, 0xb0, 0x3f, 0x2c, 0x28, 0xbd, 0x51, 0x45, 0xd9, 0x62,
	0xf2, 0x89, 0x42, 0x69, 0x90, 0x96, 0x4f, 0xa5, 0xec, 0x33, 0xdd, 0x64, 0x8e, 0xa5, 0x1b, 0xd0,
	0x36, 0x7b, 0x0f, 0xad, 0x31, 0x49, 0xc5, 0xee, 0xd3, 0xf5, 0x82, 0x32, 0xb9, 0xa5, 0xde, 0x85,
	0xba, 0x52, 0xa7, 0xeb, 0xf7, 0x46, 0xfb, 0x4b, 0xa9, 0xba, 0x58, 0x47, 0x58, 0x07, 0x0e, 0x57,
	0x68, 0xd9, 0x5c, 0x24, 0x19, 0x31, 0x17, 0x5a, 0x5e, 0x99, 0xa0, 0xa2, 0x92, 0x71, 0x38, 0xf4,
	0xca, 0x45, 0x38, 0x02, 0x88, 0x56, 0x9b, 0xd5, 0xa9, 0x7b, 0x23, 0x5c, 0xf1, 0xad, 0x22, 0x7e,
	0x21, 0x0b, 0x19, 0xd8, 0x5f, 0x82, 0x78, 0x41, 0x7a, 0x54, 0xd5, 0xf6, 0xf2, 0x10, 0x3b, 0x87,
	0xce, 0x05, 0xc5, 0x24, 0xa9, 0xd8, 0xcf, 0x09, 0xd8, 0xd7, 0x6a, 0x72, 0x86, 0x67, 0xb5, 0x42,
	0x3d, 0x4e, 0x3f, 0x8f, 0xb1, 0x2e, 0x60, 0xb1, 0xd2, 0x88, 0x7b, 0x0b, 0x5d, 0xa5, 0xf7, 0xb6,
	0x21, 0x03, 0x79, 0x8f, 0xfe, 0x59, 0x0f, 0x8e, 0x2a, 0x58, 0x86, 0x64, 0x00, 0x5d, 0x6f, 0x13,
	0x09, 0x42, 0x3d, 0x09, 0xae, 0xc8, 0x4c, 0x52, 0xdb, 0xec, 0x1d, 0x1c, 0x79, 0x9b, 0x40, 0xee,
	0xd5, 0xd1, 0x73, 0xe8, 0x19, 0xcd, 0xff, 0xc5, 0xdd, 0x07, 0x67, 0x3d, 0x3d, 0xa7, 0x1f, 0x9c,
	0x83, 0xad, 0xbf, 0x08, 0x04, 0x68, 0x5c, 0x2e, 0x26, 0x31, 0x0f, 0xdb, 0x0f, 0x70, 0x1f, 0x1e,
	0x5e, 0x52, 0x9a, 0x89, 0x24, 0x88, 0xdb, 0x35, 0x3c, 0x82, 0xce, 0x9b, 0xfc, 0xba, 0x7d, 0xca,
	0x64, 0xca, 0x43, 0x49, 0x51, 0xdb, 0x1a, 0xfd, 0xae, 0x43, 0x53, 0xcd, 0x7c, 0x2c, 0x45, 0x4a,
	0x78, 0x09, 0xbb, 0xe6, 0xc0, 0xf0, 0xd1, 0xea, 0x53, 0x2b, 0x9d, 0x57, 0xbf, 0xb7, 0xe6, 0x37,
	0x73, 0xec, 0xfe, 0xf8, 0xf3, 0xf7, 0xa7, 0xd5, 0x62, 0xcd, 0xa1, 0x4a, 0x18, 0x66, 0x24, 0x5f,
	0xd5, 0x06, 0x0a, 0xd1, 0xab, 0x22, 0x7a, 0x77, 0x20, 0x7a, 0xdb, 0x11, 0x67, 0x39, 0xe2, 0x47,
	0xf5, 0xb3, 0xb4, 0x3c, 0x15, 0x7c, 0x7c, 0x3b, 0xe4, 0xca, 0xe1, 0xf5, 0xfb, 0x9b, 0x42, 0x06,
	0xba, 0xa7, 0xa1, 0x3b, 0x6c, 0x3f, 0x87, 0x8e, 0x74, 0x86, 0x42, 0x9f, 0xc2, 0x41, 0x69, 0xc3,
	0x78, 0x5c, 0xec, 0xae, 0xba, 0xa8, 0xfe, 0x93, 0x3b, 0xa2, 0x86, 0xc6, 0xd1, 0x34, 0x88, 0xed,
	0x25, 0xcd, 0x74, 0xf8, 0x4d, 0x2d, 0xf3, 0x3b, 0x4e, 0xe0, 0x60, 0xbc, 0x99, 0x67, 0xbc, 0x95,
	0x67, 0xf3, 0x0d, 0x57, 0x26, 0x15, 0xd1, 0x54, 0x69, 0x11, 0xd0, 0xae, 0x5e, 0x0c, 0x3e, 0xab,
	0x0c, 0x65, 0x8d, 0xc9, 0xbd, 0x3b, 0xa1, 0x2c, 0x6a, 0xb0, 0x26, 0x6a, 0xd2, 0xd0, 0xff, 0x1c,
	0x2f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xd5, 0x7a, 0x0e, 0x6d, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataStoreClient is the client API for DataStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataStoreClient interface {
	SetData(ctx context.Context, in *SetDataRequest, opts ...grpc.CallOption) (*SetDataResponse, error)
	GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error)
	DeleteData(ctx context.Context, in *DeleteDataRequest, opts ...grpc.CallOption) (*DeleteDataResponse, error)
	GetDefinition(ctx context.Context, in *GetDefinitionRequest, opts ...grpc.CallOption) (*GetDefinitionResponse, error)
	SetDefinition(ctx context.Context, in *SetDefinitionRequest, opts ...grpc.CallOption) (*SetDefinitionResponse, error)
	DeleteDefinition(ctx context.Context, in *DeleteDefinitionRequest, opts ...grpc.CallOption) (*DeleteDefinitionResponse, error)
}

type dataStoreClient struct {
	cc *grpc.ClientConn
}

func NewDataStoreClient(cc *grpc.ClientConn) DataStoreClient {
	return &dataStoreClient{cc}
}

func (c *dataStoreClient) SetData(ctx context.Context, in *SetDataRequest, opts ...grpc.CallOption) (*SetDataResponse, error) {
	out := new(SetDataResponse)
	err := c.cc.Invoke(ctx, "/datapb.DataStore/SetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreClient) GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error) {
	out := new(GetDataResponse)
	err := c.cc.Invoke(ctx, "/datapb.DataStore/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreClient) DeleteData(ctx context.Context, in *DeleteDataRequest, opts ...grpc.CallOption) (*DeleteDataResponse, error) {
	out := new(DeleteDataResponse)
	err := c.cc.Invoke(ctx, "/datapb.DataStore/DeleteData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreClient) GetDefinition(ctx context.Context, in *GetDefinitionRequest, opts ...grpc.CallOption) (*GetDefinitionResponse, error) {
	out := new(GetDefinitionResponse)
	err := c.cc.Invoke(ctx, "/datapb.DataStore/GetDefinition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreClient) SetDefinition(ctx context.Context, in *SetDefinitionRequest, opts ...grpc.CallOption) (*SetDefinitionResponse, error) {
	out := new(SetDefinitionResponse)
	err := c.cc.Invoke(ctx, "/datapb.DataStore/SetDefinition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreClient) DeleteDefinition(ctx context.Context, in *DeleteDefinitionRequest, opts ...grpc.CallOption) (*DeleteDefinitionResponse, error) {
	out := new(DeleteDefinitionResponse)
	err := c.cc.Invoke(ctx, "/datapb.DataStore/DeleteDefinition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataStoreServer is the server API for DataStore service.
type DataStoreServer interface {
	SetData(context.Context, *SetDataRequest) (*SetDataResponse, error)
	GetData(context.Context, *GetDataRequest) (*GetDataResponse, error)
	DeleteData(context.Context, *DeleteDataRequest) (*DeleteDataResponse, error)
	GetDefinition(context.Context, *GetDefinitionRequest) (*GetDefinitionResponse, error)
	SetDefinition(context.Context, *SetDefinitionRequest) (*SetDefinitionResponse, error)
	DeleteDefinition(context.Context, *DeleteDefinitionRequest) (*DeleteDefinitionResponse, error)
}

// UnimplementedDataStoreServer can be embedded to have forward compatible implementations.
type UnimplementedDataStoreServer struct {
}

func (*UnimplementedDataStoreServer) SetData(ctx context.Context, req *SetDataRequest) (*SetDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetData not implemented")
}
func (*UnimplementedDataStoreServer) GetData(ctx context.Context, req *GetDataRequest) (*GetDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (*UnimplementedDataStoreServer) DeleteData(ctx context.Context, req *DeleteDataRequest) (*DeleteDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteData not implemented")
}
func (*UnimplementedDataStoreServer) GetDefinition(ctx context.Context, req *GetDefinitionRequest) (*GetDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefinition not implemented")
}
func (*UnimplementedDataStoreServer) SetDefinition(ctx context.Context, req *SetDefinitionRequest) (*SetDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDefinition not implemented")
}
func (*UnimplementedDataStoreServer) DeleteDefinition(ctx context.Context, req *DeleteDefinitionRequest) (*DeleteDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDefinition not implemented")
}

func RegisterDataStoreServer(s *grpc.Server, srv DataStoreServer) {
	s.RegisterService(&_DataStore_serviceDesc, srv)
}

func _DataStore_SetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataStoreServer).SetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datapb.DataStore/SetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataStoreServer).SetData(ctx, req.(*SetDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataStore_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataStoreServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datapb.DataStore/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataStoreServer).GetData(ctx, req.(*GetDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataStore_DeleteData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataStoreServer).DeleteData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datapb.DataStore/DeleteData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataStoreServer).DeleteData(ctx, req.(*DeleteDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataStore_GetDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataStoreServer).GetDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datapb.DataStore/GetDefinition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataStoreServer).GetDefinition(ctx, req.(*GetDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataStore_SetDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataStoreServer).SetDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datapb.DataStore/SetDefinition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataStoreServer).SetDefinition(ctx, req.(*SetDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataStore_DeleteDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataStoreServer).DeleteDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datapb.DataStore/DeleteDefinition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataStoreServer).DeleteDefinition(ctx, req.(*DeleteDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "datapb.DataStore",
	HandlerType: (*DataStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetData",
			Handler:    _DataStore_SetData_Handler,
		},
		{
			MethodName: "GetData",
			Handler:    _DataStore_GetData_Handler,
		},
		{
			MethodName: "DeleteData",
			Handler:    _DataStore_DeleteData_Handler,
		},
		{
			MethodName: "GetDefinition",
			Handler:    _DataStore_GetDefinition_Handler,
		},
		{
			MethodName: "SetDefinition",
			Handler:    _DataStore_SetDefinition_Handler,
		},
		{
			MethodName: "DeleteDefinition",
			Handler:    _DataStore_DeleteDefinition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}
