// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login.proto

package pb

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

//登录类型
type LoginType int32

const (
	//默认无效
	LoginType_Default LoginType = 0
	//code登录
	LoginType_Code LoginType = 1
	//账户登录
	LoginType_Account LoginType = 2
)

var LoginType_name = map[int32]string{
	0: "Default",
	1: "Code",
	2: "Account",
}

var LoginType_value = map[string]int32{
	"Default": 0,
	"Code":    1,
	"Account": 2,
}

func (x LoginType) String() string {
	return proto.EnumName(LoginType_name, int32(x))
}

func (LoginType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{0}
}

type LoginRequest struct {
	//登录类型
	LoginType LoginType `protobuf:"varint,1,opt,name=login_type,json=loginType,proto3,enum=user.LoginType" json:"login_type,omitempty"`
	//用户名
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	//设备号
	DeviceId string `protobuf:"bytes,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	//密码
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	//验证码
	SmsCode              string   `protobuf:"bytes,5,opt,name=sms_code,json=smsCode,proto3" json:"sms_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetLoginType() LoginType {
	if m != nil {
		return m.LoginType
	}
	return LoginType_Default
}

func (m *LoginRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetSmsCode() string {
	if m != nil {
		return m.SmsCode
	}
	return ""
}

//通证类型
type LoginResponse struct {
	//登录结果
	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	//用户名
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *LoginResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterEnum("user.LoginType", LoginType_name, LoginType_value)
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "user.LoginResponse")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor_67c21677aa7f4e4f) }

var fileDescriptor_67c21677aa7f4e4f = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xcd, 0x4e, 0x84, 0x30,
	0x14, 0x85, 0x65, 0x44, 0x28, 0x77, 0xfc, 0x21, 0x77, 0xa1, 0x55, 0x37, 0x93, 0x59, 0x4d, 0x5c,
	0x60, 0xa2, 0x2f, 0xa0, 0xd1, 0xcd, 0x24, 0xc6, 0x05, 0x71, 0x4f, 0x90, 0x5e, 0x0d, 0x09, 0xd0,
	0xca, 0x6d, 0x35, 0xf3, 0x54, 0xbe, 0xa2, 0x69, 0x61, 0xdc, 0xf5, 0x9c, 0xef, 0x24, 0xfd, 0x5a,
	0x58, 0x76, 0xfa, 0xb3, 0x1d, 0x0a, 0x33, 0x6a, 0xab, 0x31, 0x76, 0x4c, 0xe3, 0xfa, 0x37, 0x82,
	0xe3, 0x17, 0xdf, 0x96, 0xf4, 0xe5, 0x88, 0x2d, 0x16, 0x00, 0x61, 0x55, 0xd9, 0x9d, 0x21, 0x19,
	0xad, 0xa2, 0xcd, 0xe9, 0xdd, 0x59, 0xe1, 0xb7, 0x45, 0xd8, 0xbd, 0xed, 0x0c, 0x95, 0x59, 0xb7,
	0x3f, 0xe2, 0x35, 0x64, 0x1e, 0x56, 0x43, 0xdd, 0x93, 0x5c, 0xac, 0xa2, 0x4d, 0x56, 0x0a, 0x5f,
	0xbc, 0xd6, 0x7d, 0x80, 0x8a, 0xbe, 0xdb, 0x86, 0xaa, 0x56, 0xc9, 0xc3, 0x09, 0x4e, 0xc5, 0x56,
	0xe1, 0x15, 0x08, 0x53, 0x33, 0xff, 0xe8, 0x51, 0xc9, 0x78, 0x62, 0xfb, 0x8c, 0x97, 0x20, 0xb8,
	0xe7, 0xaa, 0xd1, 0x8a, 0xe4, 0x51, 0x60, 0x29, 0xf7, 0xfc, 0xa4, 0x15, 0xad, 0x1f, 0xe0, 0x64,
	0x16, 0x66, 0xa3, 0x07, 0x26, 0x3c, 0x87, 0x64, 0x24, 0x76, 0x9d, 0x0d, 0xb6, 0xa2, 0x9c, 0x13,
	0x5e, 0x40, 0x1a, 0xcc, 0x5a, 0x35, 0x7b, 0x25, 0x3e, 0x6e, 0xd5, 0xcd, 0x2d, 0x64, 0xff, 0x4f,
	0xc1, 0x25, 0xa4, 0xcf, 0xf4, 0x51, 0xbb, 0xce, 0xe6, 0x07, 0x28, 0x20, 0xf6, 0x77, 0xe4, 0x91,
	0xaf, 0x1f, 0x9b, 0x46, 0xbb, 0xc1, 0xe6, 0x8b, 0xf7, 0x24, 0xfc, 0xd8, 0xfd, 0x5f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x09, 0xd1, 0x89, 0x1d, 0x40, 0x01, 0x00, 0x00,
}
