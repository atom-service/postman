// Code generated by protoc-gen-go. DO NOT EDIT.
// source: standard.proto

package standard

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// 状态
type State int32

const (
	State_UNKNOWN              State = 0
	State_SUCCESS              State = 1
	State_FAILURE              State = 2
	State_REPEAT               State = 3
	State_WAITING              State = 4
	State_PROCEEDING           State = 5
	State_NOT_EXIST            State = 6
	State_UNDEFINED            State = 7
	State_PROHIBITED           State = 8
	State_DB_OPERATION_FATLURE State = 9
)

var State_name = map[int32]string{
	0: "UNKNOWN",
	1: "SUCCESS",
	2: "FAILURE",
	3: "REPEAT",
	4: "WAITING",
	5: "PROCEEDING",
	6: "NOT_EXIST",
	7: "UNDEFINED",
	8: "PROHIBITED",
	9: "DB_OPERATION_FATLURE",
}

var State_value = map[string]int32{
	"UNKNOWN":              0,
	"SUCCESS":              1,
	"FAILURE":              2,
	"REPEAT":               3,
	"WAITING":              4,
	"PROCEEDING":           5,
	"NOT_EXIST":            6,
	"UNDEFINED":            7,
	"PROHIBITED":           8,
	"DB_OPERATION_FATLURE": 9,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0b3e239150a6a10f, []int{0}
}

type CheckVerifyCodeRequest struct {
	HashCode             string   `protobuf:"bytes,1,opt,name=HashCode,proto3" json:"HashCode,omitempty"`
	VerifyCode           string   `protobuf:"bytes,2,opt,name=VerifyCode,proto3" json:"VerifyCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckVerifyCodeRequest) Reset()         { *m = CheckVerifyCodeRequest{} }
func (m *CheckVerifyCodeRequest) String() string { return proto.CompactTextString(m) }
func (*CheckVerifyCodeRequest) ProtoMessage()    {}
func (*CheckVerifyCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b3e239150a6a10f, []int{0}
}

func (m *CheckVerifyCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckVerifyCodeRequest.Unmarshal(m, b)
}
func (m *CheckVerifyCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckVerifyCodeRequest.Marshal(b, m, deterministic)
}
func (m *CheckVerifyCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckVerifyCodeRequest.Merge(m, src)
}
func (m *CheckVerifyCodeRequest) XXX_Size() int {
	return xxx_messageInfo_CheckVerifyCodeRequest.Size(m)
}
func (m *CheckVerifyCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckVerifyCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckVerifyCodeRequest proto.InternalMessageInfo

func (m *CheckVerifyCodeRequest) GetHashCode() string {
	if m != nil {
		return m.HashCode
	}
	return ""
}

func (m *CheckVerifyCodeRequest) GetVerifyCode() string {
	if m != nil {
		return m.VerifyCode
	}
	return ""
}

type CheckVerifyCodeeResponse struct {
	State                State    `protobuf:"varint,1,opt,name=State,proto3,enum=standard.State" json:"State,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckVerifyCodeeResponse) Reset()         { *m = CheckVerifyCodeeResponse{} }
func (m *CheckVerifyCodeeResponse) String() string { return proto.CompactTextString(m) }
func (*CheckVerifyCodeeResponse) ProtoMessage()    {}
func (*CheckVerifyCodeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b3e239150a6a10f, []int{1}
}

func (m *CheckVerifyCodeeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckVerifyCodeeResponse.Unmarshal(m, b)
}
func (m *CheckVerifyCodeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckVerifyCodeeResponse.Marshal(b, m, deterministic)
}
func (m *CheckVerifyCodeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckVerifyCodeeResponse.Merge(m, src)
}
func (m *CheckVerifyCodeeResponse) XXX_Size() int {
	return xxx_messageInfo_CheckVerifyCodeeResponse.Size(m)
}
func (m *CheckVerifyCodeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckVerifyCodeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckVerifyCodeeResponse proto.InternalMessageInfo

func (m *CheckVerifyCodeeResponse) GetState() State {
	if m != nil {
		return m.State
	}
	return State_UNKNOWN
}

func (m *CheckVerifyCodeeResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DiscardVerifyCodeRequest struct {
	HashCode             string   `protobuf:"bytes,1,opt,name=HashCode,proto3" json:"HashCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscardVerifyCodeRequest) Reset()         { *m = DiscardVerifyCodeRequest{} }
func (m *DiscardVerifyCodeRequest) String() string { return proto.CompactTextString(m) }
func (*DiscardVerifyCodeRequest) ProtoMessage()    {}
func (*DiscardVerifyCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b3e239150a6a10f, []int{2}
}

func (m *DiscardVerifyCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscardVerifyCodeRequest.Unmarshal(m, b)
}
func (m *DiscardVerifyCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscardVerifyCodeRequest.Marshal(b, m, deterministic)
}
func (m *DiscardVerifyCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscardVerifyCodeRequest.Merge(m, src)
}
func (m *DiscardVerifyCodeRequest) XXX_Size() int {
	return xxx_messageInfo_DiscardVerifyCodeRequest.Size(m)
}
func (m *DiscardVerifyCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscardVerifyCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscardVerifyCodeRequest proto.InternalMessageInfo

func (m *DiscardVerifyCodeRequest) GetHashCode() string {
	if m != nil {
		return m.HashCode
	}
	return ""
}

type DiscardVerifyCodeResponse struct {
	State                State    `protobuf:"varint,1,opt,name=State,proto3,enum=standard.State" json:"State,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscardVerifyCodeResponse) Reset()         { *m = DiscardVerifyCodeResponse{} }
func (m *DiscardVerifyCodeResponse) String() string { return proto.CompactTextString(m) }
func (*DiscardVerifyCodeResponse) ProtoMessage()    {}
func (*DiscardVerifyCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b3e239150a6a10f, []int{3}
}

func (m *DiscardVerifyCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscardVerifyCodeResponse.Unmarshal(m, b)
}
func (m *DiscardVerifyCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscardVerifyCodeResponse.Marshal(b, m, deterministic)
}
func (m *DiscardVerifyCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscardVerifyCodeResponse.Merge(m, src)
}
func (m *DiscardVerifyCodeResponse) XXX_Size() int {
	return xxx_messageInfo_DiscardVerifyCodeResponse.Size(m)
}
func (m *DiscardVerifyCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscardVerifyCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiscardVerifyCodeResponse proto.InternalMessageInfo

func (m *DiscardVerifyCodeResponse) GetState() State {
	if m != nil {
		return m.State
	}
	return State_UNKNOWN
}

func (m *DiscardVerifyCodeResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SendVerifyCodeByEmailRequest struct {
	Operation            string   `protobuf:"bytes,1,opt,name=Operation,proto3" json:"Operation,omitempty"`
	EmailAddress         string   `protobuf:"bytes,2,opt,name=EmailAddress,proto3" json:"EmailAddress,omitempty"`
	ValidityPeriod       int64    `protobuf:"varint,3,opt,name=ValidityPeriod,proto3" json:"ValidityPeriod,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendVerifyCodeByEmailRequest) Reset()         { *m = SendVerifyCodeByEmailRequest{} }
func (m *SendVerifyCodeByEmailRequest) String() string { return proto.CompactTextString(m) }
func (*SendVerifyCodeByEmailRequest) ProtoMessage()    {}
func (*SendVerifyCodeByEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b3e239150a6a10f, []int{4}
}

func (m *SendVerifyCodeByEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendVerifyCodeByEmailRequest.Unmarshal(m, b)
}
func (m *SendVerifyCodeByEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendVerifyCodeByEmailRequest.Marshal(b, m, deterministic)
}
func (m *SendVerifyCodeByEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendVerifyCodeByEmailRequest.Merge(m, src)
}
func (m *SendVerifyCodeByEmailRequest) XXX_Size() int {
	return xxx_messageInfo_SendVerifyCodeByEmailRequest.Size(m)
}
func (m *SendVerifyCodeByEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendVerifyCodeByEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendVerifyCodeByEmailRequest proto.InternalMessageInfo

func (m *SendVerifyCodeByEmailRequest) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *SendVerifyCodeByEmailRequest) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *SendVerifyCodeByEmailRequest) GetValidityPeriod() int64 {
	if m != nil {
		return m.ValidityPeriod
	}
	return 0
}

//发送成功失败状态
type SendVerifyCodeByEmailResponse struct {
	State                State    `protobuf:"varint,1,opt,name=State,proto3,enum=standard.State" json:"State,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	HashCode             string   `protobuf:"bytes,3,opt,name=HashCode,proto3" json:"HashCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendVerifyCodeByEmailResponse) Reset()         { *m = SendVerifyCodeByEmailResponse{} }
func (m *SendVerifyCodeByEmailResponse) String() string { return proto.CompactTextString(m) }
func (*SendVerifyCodeByEmailResponse) ProtoMessage()    {}
func (*SendVerifyCodeByEmailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b3e239150a6a10f, []int{5}
}

func (m *SendVerifyCodeByEmailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendVerifyCodeByEmailResponse.Unmarshal(m, b)
}
func (m *SendVerifyCodeByEmailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendVerifyCodeByEmailResponse.Marshal(b, m, deterministic)
}
func (m *SendVerifyCodeByEmailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendVerifyCodeByEmailResponse.Merge(m, src)
}
func (m *SendVerifyCodeByEmailResponse) XXX_Size() int {
	return xxx_messageInfo_SendVerifyCodeByEmailResponse.Size(m)
}
func (m *SendVerifyCodeByEmailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendVerifyCodeByEmailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendVerifyCodeByEmailResponse proto.InternalMessageInfo

func (m *SendVerifyCodeByEmailResponse) GetState() State {
	if m != nil {
		return m.State
	}
	return State_UNKNOWN
}

func (m *SendVerifyCodeByEmailResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SendVerifyCodeByEmailResponse) GetHashCode() string {
	if m != nil {
		return m.HashCode
	}
	return ""
}

type SendVerifyCodeBySmsRequest struct {
	Operation            string   `protobuf:"bytes,1,opt,name=Operation,proto3" json:"Operation,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,2,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	CountryCode          string   `protobuf:"bytes,3,opt,name=CountryCode,proto3" json:"CountryCode,omitempty"`
	ValidityPeriod       int64    `protobuf:"varint,4,opt,name=ValidityPeriod,proto3" json:"ValidityPeriod,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendVerifyCodeBySmsRequest) Reset()         { *m = SendVerifyCodeBySmsRequest{} }
func (m *SendVerifyCodeBySmsRequest) String() string { return proto.CompactTextString(m) }
func (*SendVerifyCodeBySmsRequest) ProtoMessage()    {}
func (*SendVerifyCodeBySmsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b3e239150a6a10f, []int{6}
}

func (m *SendVerifyCodeBySmsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendVerifyCodeBySmsRequest.Unmarshal(m, b)
}
func (m *SendVerifyCodeBySmsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendVerifyCodeBySmsRequest.Marshal(b, m, deterministic)
}
func (m *SendVerifyCodeBySmsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendVerifyCodeBySmsRequest.Merge(m, src)
}
func (m *SendVerifyCodeBySmsRequest) XXX_Size() int {
	return xxx_messageInfo_SendVerifyCodeBySmsRequest.Size(m)
}
func (m *SendVerifyCodeBySmsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendVerifyCodeBySmsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendVerifyCodeBySmsRequest proto.InternalMessageInfo

func (m *SendVerifyCodeBySmsRequest) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *SendVerifyCodeBySmsRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *SendVerifyCodeBySmsRequest) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *SendVerifyCodeBySmsRequest) GetValidityPeriod() int64 {
	if m != nil {
		return m.ValidityPeriod
	}
	return 0
}

//发送成功失败状态
type SendVerifyCodeBySmsResponse struct {
	State                State    `protobuf:"varint,1,opt,name=State,proto3,enum=standard.State" json:"State,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	HashCode             string   `protobuf:"bytes,3,opt,name=HashCode,proto3" json:"HashCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendVerifyCodeBySmsResponse) Reset()         { *m = SendVerifyCodeBySmsResponse{} }
func (m *SendVerifyCodeBySmsResponse) String() string { return proto.CompactTextString(m) }
func (*SendVerifyCodeBySmsResponse) ProtoMessage()    {}
func (*SendVerifyCodeBySmsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b3e239150a6a10f, []int{7}
}

func (m *SendVerifyCodeBySmsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendVerifyCodeBySmsResponse.Unmarshal(m, b)
}
func (m *SendVerifyCodeBySmsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendVerifyCodeBySmsResponse.Marshal(b, m, deterministic)
}
func (m *SendVerifyCodeBySmsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendVerifyCodeBySmsResponse.Merge(m, src)
}
func (m *SendVerifyCodeBySmsResponse) XXX_Size() int {
	return xxx_messageInfo_SendVerifyCodeBySmsResponse.Size(m)
}
func (m *SendVerifyCodeBySmsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendVerifyCodeBySmsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendVerifyCodeBySmsResponse proto.InternalMessageInfo

func (m *SendVerifyCodeBySmsResponse) GetState() State {
	if m != nil {
		return m.State
	}
	return State_UNKNOWN
}

func (m *SendVerifyCodeBySmsResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SendVerifyCodeBySmsResponse) GetHashCode() string {
	if m != nil {
		return m.HashCode
	}
	return ""
}

type SendVerifyCodeByCallPhoneRequest struct {
	Operation            string   `protobuf:"bytes,1,opt,name=Operation,proto3" json:"Operation,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,2,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	CountryCode          string   `protobuf:"bytes,3,opt,name=CountryCode,proto3" json:"CountryCode,omitempty"`
	ValidityPeriod       int64    `protobuf:"varint,4,opt,name=ValidityPeriod,proto3" json:"ValidityPeriod,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendVerifyCodeByCallPhoneRequest) Reset()         { *m = SendVerifyCodeByCallPhoneRequest{} }
func (m *SendVerifyCodeByCallPhoneRequest) String() string { return proto.CompactTextString(m) }
func (*SendVerifyCodeByCallPhoneRequest) ProtoMessage()    {}
func (*SendVerifyCodeByCallPhoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b3e239150a6a10f, []int{8}
}

func (m *SendVerifyCodeByCallPhoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendVerifyCodeByCallPhoneRequest.Unmarshal(m, b)
}
func (m *SendVerifyCodeByCallPhoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendVerifyCodeByCallPhoneRequest.Marshal(b, m, deterministic)
}
func (m *SendVerifyCodeByCallPhoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendVerifyCodeByCallPhoneRequest.Merge(m, src)
}
func (m *SendVerifyCodeByCallPhoneRequest) XXX_Size() int {
	return xxx_messageInfo_SendVerifyCodeByCallPhoneRequest.Size(m)
}
func (m *SendVerifyCodeByCallPhoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendVerifyCodeByCallPhoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendVerifyCodeByCallPhoneRequest proto.InternalMessageInfo

func (m *SendVerifyCodeByCallPhoneRequest) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *SendVerifyCodeByCallPhoneRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *SendVerifyCodeByCallPhoneRequest) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *SendVerifyCodeByCallPhoneRequest) GetValidityPeriod() int64 {
	if m != nil {
		return m.ValidityPeriod
	}
	return 0
}

//发送成功失败状态
type SendVerifyCodeByCallPhoneResponse struct {
	State                State    `protobuf:"varint,1,opt,name=State,proto3,enum=standard.State" json:"State,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	HashCode             string   `protobuf:"bytes,3,opt,name=HashCode,proto3" json:"HashCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendVerifyCodeByCallPhoneResponse) Reset()         { *m = SendVerifyCodeByCallPhoneResponse{} }
func (m *SendVerifyCodeByCallPhoneResponse) String() string { return proto.CompactTextString(m) }
func (*SendVerifyCodeByCallPhoneResponse) ProtoMessage()    {}
func (*SendVerifyCodeByCallPhoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b3e239150a6a10f, []int{9}
}

func (m *SendVerifyCodeByCallPhoneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendVerifyCodeByCallPhoneResponse.Unmarshal(m, b)
}
func (m *SendVerifyCodeByCallPhoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendVerifyCodeByCallPhoneResponse.Marshal(b, m, deterministic)
}
func (m *SendVerifyCodeByCallPhoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendVerifyCodeByCallPhoneResponse.Merge(m, src)
}
func (m *SendVerifyCodeByCallPhoneResponse) XXX_Size() int {
	return xxx_messageInfo_SendVerifyCodeByCallPhoneResponse.Size(m)
}
func (m *SendVerifyCodeByCallPhoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendVerifyCodeByCallPhoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendVerifyCodeByCallPhoneResponse proto.InternalMessageInfo

func (m *SendVerifyCodeByCallPhoneResponse) GetState() State {
	if m != nil {
		return m.State
	}
	return State_UNKNOWN
}

func (m *SendVerifyCodeByCallPhoneResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SendVerifyCodeByCallPhoneResponse) GetHashCode() string {
	if m != nil {
		return m.HashCode
	}
	return ""
}

func init() {
	proto.RegisterEnum("standard.State", State_name, State_value)
	proto.RegisterType((*CheckVerifyCodeRequest)(nil), "standard.CheckVerifyCodeRequest")
	proto.RegisterType((*CheckVerifyCodeeResponse)(nil), "standard.CheckVerifyCodeeResponse")
	proto.RegisterType((*DiscardVerifyCodeRequest)(nil), "standard.DiscardVerifyCodeRequest")
	proto.RegisterType((*DiscardVerifyCodeResponse)(nil), "standard.DiscardVerifyCodeResponse")
	proto.RegisterType((*SendVerifyCodeByEmailRequest)(nil), "standard.SendVerifyCodeByEmailRequest")
	proto.RegisterType((*SendVerifyCodeByEmailResponse)(nil), "standard.SendVerifyCodeByEmailResponse")
	proto.RegisterType((*SendVerifyCodeBySmsRequest)(nil), "standard.SendVerifyCodeBySmsRequest")
	proto.RegisterType((*SendVerifyCodeBySmsResponse)(nil), "standard.SendVerifyCodeBySmsResponse")
	proto.RegisterType((*SendVerifyCodeByCallPhoneRequest)(nil), "standard.SendVerifyCodeByCallPhoneRequest")
	proto.RegisterType((*SendVerifyCodeByCallPhoneResponse)(nil), "standard.SendVerifyCodeByCallPhoneResponse")
}

func init() { proto.RegisterFile("standard.proto", fileDescriptor_0b3e239150a6a10f) }

var fileDescriptor_0b3e239150a6a10f = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xdf, 0x6f, 0x93, 0x50,
	0x18, 0x1d, 0xeb, 0xd6, 0x1f, 0xdf, 0xb4, 0xc3, 0xeb, 0x8f, 0x30, 0x9c, 0xa6, 0xa2, 0x9b, 0xcb,
	0x4c, 0xf6, 0x30, 0x13, 0xdf, 0x29, 0xdc, 0x3a, 0xa2, 0x42, 0x03, 0x74, 0xd3, 0x68, 0x6c, 0xd8,
	0xb8, 0x5a, 0xb4, 0x85, 0x7a, 0x2f, 0x35, 0xa9, 0xf1, 0xc1, 0x47, 0x5f, 0xfc, 0x1b, 0x7c, 0xf3,
	0xc9, 0x3f, 0xd2, 0xc0, 0x0a, 0x65, 0x08, 0xdb, 0x4c, 0x96, 0xc5, 0xc7, 0xef, 0xdc, 0x73, 0xcf,
	0x39, 0x9c, 0xed, 0x7e, 0x85, 0x26, 0x0b, 0x1d, 0xdf, 0x75, 0xa8, 0xbb, 0x33, 0xa6, 0x41, 0x18,
	0xa0, 0x7a, 0x32, 0x4b, 0x36, 0xdc, 0x52, 0x06, 0xe4, 0xe8, 0xe3, 0x3e, 0xa1, 0xde, 0xbb, 0xa9,
	0x12, 0xb8, 0xc4, 0x24, 0x9f, 0x26, 0x84, 0x85, 0x48, 0x84, 0xfa, 0x9e, 0xc3, 0x06, 0x11, 0x24,
	0x70, 0x2d, 0x6e, 0xab, 0x61, 0xa6, 0x33, 0xba, 0x0b, 0x30, 0xbf, 0x20, 0x2c, 0xc6, 0xa7, 0x19,
	0x44, 0x7a, 0x0d, 0x42, 0x4e, 0x95, 0x98, 0x84, 0x8d, 0x03, 0x9f, 0x11, 0xb4, 0x01, 0xcb, 0x56,
	0xe8, 0x84, 0xc7, 0xa2, 0xcd, 0xdd, 0xd5, 0x9d, 0x34, 0x5b, 0x0c, 0x9b, 0xc7, 0xa7, 0x48, 0x80,
	0xda, 0x0b, 0xc2, 0x98, 0xf3, 0x3e, 0xd1, 0x4f, 0x46, 0xe9, 0x09, 0x08, 0xaa, 0xc7, 0x8e, 0x1c,
	0xea, 0xfe, 0x53, 0x68, 0xe9, 0x0d, 0xac, 0x15, 0xdc, 0xbb, 0xa8, 0x54, 0xdf, 0x39, 0x58, 0xb7,
	0x88, 0x9f, 0xd1, 0x6e, 0x4f, 0xf1, 0xc8, 0xf1, 0x86, 0x49, 0xb4, 0x75, 0x68, 0x18, 0x63, 0x42,
	0x9d, 0xd0, 0x0b, 0xfc, 0x59, 0xb6, 0x39, 0x80, 0x24, 0xb8, 0x12, 0xb3, 0x65, 0xd7, 0xa5, 0x84,
	0xb1, 0x99, 0xfa, 0x09, 0x0c, 0x6d, 0x42, 0x73, 0xdf, 0x19, 0x7a, 0xae, 0x17, 0x4e, 0xbb, 0x84,
	0x7a, 0x81, 0x2b, 0x54, 0x5a, 0xdc, 0x56, 0xc5, 0xcc, 0xa1, 0xd2, 0x57, 0xb8, 0x53, 0x92, 0xe4,
	0x82, 0x3e, 0xf6, 0x44, 0xcd, 0x95, 0x5c, 0xcd, 0xbf, 0x38, 0x10, 0xf3, 0xf6, 0xd6, 0x88, 0x9d,
	0xaf, 0x86, 0x16, 0xac, 0x74, 0x07, 0x81, 0x4f, 0xf4, 0xc9, 0xe8, 0x90, 0xd0, 0x99, 0x6d, 0x16,
	0x8a, 0x18, 0x4a, 0x30, 0xf1, 0x43, 0x3a, 0xcd, 0xb8, 0x67, 0xa1, 0x82, 0x9a, 0x96, 0x0a, 0x6b,
	0xfa, 0x02, 0xb7, 0x0b, 0x73, 0x5e, 0x46, 0x49, 0xbf, 0x39, 0x68, 0xe5, 0xcd, 0x15, 0x67, 0x38,
	0x8c, 0xbf, 0xf4, 0xff, 0xab, 0xea, 0x1b, 0x07, 0xf7, 0x4e, 0x89, 0x7b, 0x09, 0x8d, 0x6d, 0xff,
	0xe4, 0x66, 0xea, 0x68, 0x05, 0x6a, 0x3d, 0xfd, 0x99, 0x6e, 0x1c, 0xe8, 0xfc, 0x42, 0x34, 0x58,
	0x3d, 0x45, 0xc1, 0x96, 0xc5, 0x73, 0xd1, 0xd0, 0x91, 0xb5, 0xe7, 0x3d, 0x13, 0xf3, 0x8b, 0x08,
	0xa0, 0x6a, 0xe2, 0x2e, 0x96, 0x6d, 0xbe, 0x12, 0x1d, 0x1c, 0xc8, 0x9a, 0xad, 0xe9, 0x4f, 0xf9,
	0x25, 0xd4, 0x04, 0xe8, 0x9a, 0x86, 0x82, 0xb1, 0x1a, 0xcd, 0xcb, 0xe8, 0x2a, 0x34, 0x74, 0xc3,
	0xee, 0xe3, 0x97, 0x9a, 0x65, 0xf3, 0xd5, 0x68, 0xec, 0xe9, 0x2a, 0xee, 0x68, 0x3a, 0x56, 0xf9,
	0xda, 0x8c, 0xbd, 0xa7, 0xb5, 0x35, 0x1b, 0xab, 0x7c, 0x1d, 0x09, 0x70, 0x43, 0x6d, 0xf7, 0x8d,
	0x2e, 0x36, 0x65, 0x5b, 0x33, 0xf4, 0x7e, 0x47, 0xb6, 0x63, 0xc3, 0xc6, 0xee, 0x8f, 0x25, 0xa8,
	0x46, 0x25, 0x11, 0x8a, 0x5e, 0xc1, 0x6a, 0x6e, 0xff, 0xa1, 0xd6, 0xbc, 0x8d, 0xe2, 0x85, 0x2b,
	0x4a, 0xa5, 0x8c, 0xb4, 0x62, 0x69, 0x01, 0xbd, 0x85, 0x6b, 0x7f, 0x6d, 0x31, 0x94, 0xb9, 0x5a,
	0xb6, 0x1a, 0xc5, 0xfb, 0xa7, 0x72, 0x52, 0x7d, 0x17, 0xae, 0x17, 0xbc, 0x0a, 0xf4, 0x20, 0xf3,
	0xc7, 0x2c, 0x7d, 0xdc, 0xe2, 0xc6, 0x19, 0xac, 0xd4, 0xe5, 0x03, 0xdc, 0x2c, 0x5c, 0x51, 0x68,
	0xb3, 0x5c, 0x21, 0xbb, 0x4d, 0xc5, 0x87, 0x67, 0xf2, 0x52, 0xaf, 0xcf, 0xb0, 0x56, 0xfa, 0xbf,
	0x8b, 0xb6, 0xcb, 0x75, 0xf2, 0xef, 0x51, 0x7c, 0x74, 0x2e, 0x6e, 0xe2, 0x7b, 0x58, 0x8d, 0x7f,
	0x6b, 0x1f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x91, 0xa8, 0x3b, 0xe7, 0x7d, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SenderClient is the client API for Sender service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SenderClient interface {
	CheckVerifyCode(ctx context.Context, in *CheckVerifyCodeRequest, opts ...grpc.CallOption) (*CheckVerifyCodeeResponse, error)
	DiscardVerifyCode(ctx context.Context, in *DiscardVerifyCodeRequest, opts ...grpc.CallOption) (*DiscardVerifyCodeResponse, error)
	SendVerifyCodeBySms(ctx context.Context, in *SendVerifyCodeBySmsRequest, opts ...grpc.CallOption) (*SendVerifyCodeBySmsResponse, error)
	SendVerifyCodeByEmail(ctx context.Context, in *SendVerifyCodeByEmailRequest, opts ...grpc.CallOption) (*SendVerifyCodeByEmailResponse, error)
	SendVerifyCodeByCallPhone(ctx context.Context, in *SendVerifyCodeByCallPhoneRequest, opts ...grpc.CallOption) (*SendVerifyCodeByCallPhoneResponse, error)
}

type senderClient struct {
	cc *grpc.ClientConn
}

func NewSenderClient(cc *grpc.ClientConn) SenderClient {
	return &senderClient{cc}
}

func (c *senderClient) CheckVerifyCode(ctx context.Context, in *CheckVerifyCodeRequest, opts ...grpc.CallOption) (*CheckVerifyCodeeResponse, error) {
	out := new(CheckVerifyCodeeResponse)
	err := c.cc.Invoke(ctx, "/standard.Sender/CheckVerifyCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *senderClient) DiscardVerifyCode(ctx context.Context, in *DiscardVerifyCodeRequest, opts ...grpc.CallOption) (*DiscardVerifyCodeResponse, error) {
	out := new(DiscardVerifyCodeResponse)
	err := c.cc.Invoke(ctx, "/standard.Sender/DiscardVerifyCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *senderClient) SendVerifyCodeBySms(ctx context.Context, in *SendVerifyCodeBySmsRequest, opts ...grpc.CallOption) (*SendVerifyCodeBySmsResponse, error) {
	out := new(SendVerifyCodeBySmsResponse)
	err := c.cc.Invoke(ctx, "/standard.Sender/SendVerifyCodeBySms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *senderClient) SendVerifyCodeByEmail(ctx context.Context, in *SendVerifyCodeByEmailRequest, opts ...grpc.CallOption) (*SendVerifyCodeByEmailResponse, error) {
	out := new(SendVerifyCodeByEmailResponse)
	err := c.cc.Invoke(ctx, "/standard.Sender/SendVerifyCodeByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *senderClient) SendVerifyCodeByCallPhone(ctx context.Context, in *SendVerifyCodeByCallPhoneRequest, opts ...grpc.CallOption) (*SendVerifyCodeByCallPhoneResponse, error) {
	out := new(SendVerifyCodeByCallPhoneResponse)
	err := c.cc.Invoke(ctx, "/standard.Sender/SendVerifyCodeByCallPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SenderServer is the server API for Sender service.
type SenderServer interface {
	CheckVerifyCode(context.Context, *CheckVerifyCodeRequest) (*CheckVerifyCodeeResponse, error)
	DiscardVerifyCode(context.Context, *DiscardVerifyCodeRequest) (*DiscardVerifyCodeResponse, error)
	SendVerifyCodeBySms(context.Context, *SendVerifyCodeBySmsRequest) (*SendVerifyCodeBySmsResponse, error)
	SendVerifyCodeByEmail(context.Context, *SendVerifyCodeByEmailRequest) (*SendVerifyCodeByEmailResponse, error)
	SendVerifyCodeByCallPhone(context.Context, *SendVerifyCodeByCallPhoneRequest) (*SendVerifyCodeByCallPhoneResponse, error)
}

// UnimplementedSenderServer can be embedded to have forward compatible implementations.
type UnimplementedSenderServer struct {
}

func (*UnimplementedSenderServer) CheckVerifyCode(ctx context.Context, req *CheckVerifyCodeRequest) (*CheckVerifyCodeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckVerifyCode not implemented")
}
func (*UnimplementedSenderServer) DiscardVerifyCode(ctx context.Context, req *DiscardVerifyCodeRequest) (*DiscardVerifyCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscardVerifyCode not implemented")
}
func (*UnimplementedSenderServer) SendVerifyCodeBySms(ctx context.Context, req *SendVerifyCodeBySmsRequest) (*SendVerifyCodeBySmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerifyCodeBySms not implemented")
}
func (*UnimplementedSenderServer) SendVerifyCodeByEmail(ctx context.Context, req *SendVerifyCodeByEmailRequest) (*SendVerifyCodeByEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerifyCodeByEmail not implemented")
}
func (*UnimplementedSenderServer) SendVerifyCodeByCallPhone(ctx context.Context, req *SendVerifyCodeByCallPhoneRequest) (*SendVerifyCodeByCallPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerifyCodeByCallPhone not implemented")
}

func RegisterSenderServer(s *grpc.Server, srv SenderServer) {
	s.RegisterService(&_Sender_serviceDesc, srv)
}

func _Sender_CheckVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckVerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SenderServer).CheckVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/standard.Sender/CheckVerifyCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SenderServer).CheckVerifyCode(ctx, req.(*CheckVerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sender_DiscardVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscardVerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SenderServer).DiscardVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/standard.Sender/DiscardVerifyCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SenderServer).DiscardVerifyCode(ctx, req.(*DiscardVerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sender_SendVerifyCodeBySms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVerifyCodeBySmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SenderServer).SendVerifyCodeBySms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/standard.Sender/SendVerifyCodeBySms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SenderServer).SendVerifyCodeBySms(ctx, req.(*SendVerifyCodeBySmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sender_SendVerifyCodeByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVerifyCodeByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SenderServer).SendVerifyCodeByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/standard.Sender/SendVerifyCodeByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SenderServer).SendVerifyCodeByEmail(ctx, req.(*SendVerifyCodeByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sender_SendVerifyCodeByCallPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVerifyCodeByCallPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SenderServer).SendVerifyCodeByCallPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/standard.Sender/SendVerifyCodeByCallPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SenderServer).SendVerifyCodeByCallPhone(ctx, req.(*SendVerifyCodeByCallPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sender_serviceDesc = grpc.ServiceDesc{
	ServiceName: "standard.Sender",
	HandlerType: (*SenderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckVerifyCode",
			Handler:    _Sender_CheckVerifyCode_Handler,
		},
		{
			MethodName: "DiscardVerifyCode",
			Handler:    _Sender_DiscardVerifyCode_Handler,
		},
		{
			MethodName: "SendVerifyCodeBySms",
			Handler:    _Sender_SendVerifyCodeBySms_Handler,
		},
		{
			MethodName: "SendVerifyCodeByEmail",
			Handler:    _Sender_SendVerifyCodeByEmail_Handler,
		},
		{
			MethodName: "SendVerifyCodeByCallPhone",
			Handler:    _Sender_SendVerifyCodeByCallPhone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "standard.proto",
}
