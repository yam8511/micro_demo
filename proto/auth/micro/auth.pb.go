// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	proto/auth/auth.proto

It has these top-level messages:
	PingRequest
	PongResponse
	User
	RegisterRequest
	RegisterResponse
	LoginRequest
	LoginResponse
	CheckRequest
	CheckResponse
	LogoutRequest
	LogoutResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 性別群集
type Sex int32

const (
	Sex_MALE   Sex = 0
	Sex_FEMALE Sex = 1
)

var Sex_name = map[int32]string{
	0: "MALE",
	1: "FEMALE",
}
var Sex_value = map[string]int32{
	"MALE":   0,
	"FEMALE": 1,
}

func (x Sex) String() string {
	return proto.EnumName(Sex_name, int32(x))
}
func (Sex) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 登入方式群集
type LoginMethod int32

const (
	LoginMethod_PHONE LoginMethod = 0
	LoginMethod_EMAIL LoginMethod = 1
)

var LoginMethod_name = map[int32]string{
	0: "PHONE",
	1: "EMAIL",
}
var LoginMethod_value = map[string]int32{
	"PHONE": 0,
	"EMAIL": 1,
}

func (x LoginMethod) String() string {
	return proto.EnumName(LoginMethod_name, int32(x))
}
func (LoginMethod) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type PingRequest struct {
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PongResponse struct {
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	Environment string `protobuf:"bytes,2,opt,name=environment" json:"environment,omitempty"`
}

func (m *PongResponse) Reset()                    { *m = PongResponse{} }
func (m *PongResponse) String() string            { return proto.CompactTextString(m) }
func (*PongResponse) ProtoMessage()               {}
func (*PongResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PongResponse) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *PongResponse) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

// User 使用者資料
type User struct {
	Id        int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Nickname  string `protobuf:"bytes,4,opt,name=nickname" json:"nickname,omitempty"`
	Address   string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	Phone     string `protobuf:"bytes,6,opt,name=phone" json:"phone,omitempty"`
	Email     string `protobuf:"bytes,7,opt,name=email" json:"email,omitempty"`
	Sex       Sex    `protobuf:"varint,8,opt,name=sex,enum=pb.Sex" json:"sex,omitempty"`
	AuthToken string `protobuf:"bytes,9,opt,name=auth_token,json=authToken" json:"auth_token,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetSex() Sex {
	if m != nil {
		return m.Sex
	}
	return Sex_MALE
}

func (m *User) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

// 註冊請求
type RegisterRequest struct {
	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Nickname  string `protobuf:"bytes,3,opt,name=nickname" json:"nickname,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
	Phone     string `protobuf:"bytes,5,opt,name=phone" json:"phone,omitempty"`
	Email     string `protobuf:"bytes,6,opt,name=email" json:"email,omitempty"`
	Address   string `protobuf:"bytes,7,opt,name=address" json:"address,omitempty"`
	Sex       Sex    `protobuf:"varint,8,opt,name=sex,enum=pb.Sex" json:"sex,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RegisterRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *RegisterRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *RegisterRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RegisterRequest) GetSex() Sex {
	if m != nil {
		return m.Sex
	}
	return Sex_MALE
}

// 回傳註冊狀態
type RegisterResponse struct {
	ErrorCode string `protobuf:"bytes,1,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	ErrorText string `protobuf:"bytes,2,opt,name=error_text,json=errorText" json:"error_text,omitempty"`
	User      *User  `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RegisterResponse) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *RegisterResponse) GetErrorText() string {
	if m != nil {
		return m.ErrorText
	}
	return ""
}

func (m *RegisterResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// 登入請求
type LoginRequest struct {
	LoginData   string      `protobuf:"bytes,1,opt,name=login_data,json=loginData" json:"login_data,omitempty"`
	Password    string      `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	LoginMethod LoginMethod `protobuf:"varint,3,opt,name=login_method,json=loginMethod,enum=pb.LoginMethod" json:"login_method,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LoginRequest) GetLoginData() string {
	if m != nil {
		return m.LoginData
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetLoginMethod() LoginMethod {
	if m != nil {
		return m.LoginMethod
	}
	return LoginMethod_PHONE
}

// 回傳登入狀態
type LoginResponse struct {
	ErrorCode  string `protobuf:"bytes,1,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	ErrorText  string `protobuf:"bytes,2,opt,name=error_text,json=errorText" json:"error_text,omitempty"`
	LoginCheck bool   `protobuf:"varint,3,opt,name=login_check,json=loginCheck" json:"login_check,omitempty"`
	User       *User  `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LoginResponse) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *LoginResponse) GetErrorText() string {
	if m != nil {
		return m.ErrorText
	}
	return ""
}

func (m *LoginResponse) GetLoginCheck() bool {
	if m != nil {
		return m.LoginCheck
	}
	return false
}

func (m *LoginResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// 確認登入請求
type CheckRequest struct {
	UserId    int64  `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	AuthToken string `protobuf:"bytes,2,opt,name=auth_token,json=authToken" json:"auth_token,omitempty"`
}

func (m *CheckRequest) Reset()                    { *m = CheckRequest{} }
func (m *CheckRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()               {}
func (*CheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CheckRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CheckRequest) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

// 回傳登出狀態
type CheckResponse struct {
	ErrorCode string `protobuf:"bytes,1,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	ErrorText string `protobuf:"bytes,2,opt,name=error_text,json=errorText" json:"error_text,omitempty"`
	User      *User  `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
}

func (m *CheckResponse) Reset()                    { *m = CheckResponse{} }
func (m *CheckResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()               {}
func (*CheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CheckResponse) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *CheckResponse) GetErrorText() string {
	if m != nil {
		return m.ErrorText
	}
	return ""
}

func (m *CheckResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// 登出請求
type LogoutRequest struct {
	UserId    int64  `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	AuthToken string `protobuf:"bytes,2,opt,name=auth_token,json=authToken" json:"auth_token,omitempty"`
}

func (m *LogoutRequest) Reset()                    { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string            { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()               {}
func (*LogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *LogoutRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *LogoutRequest) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

// 回傳登出狀態
type LogoutResponse struct {
	ErrorCode string `protobuf:"bytes,1,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	ErrorText string `protobuf:"bytes,2,opt,name=error_text,json=errorText" json:"error_text,omitempty"`
}

func (m *LogoutResponse) Reset()                    { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string            { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()               {}
func (*LogoutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *LogoutResponse) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *LogoutResponse) GetErrorText() string {
	if m != nil {
		return m.ErrorText
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "pb.PingRequest")
	proto.RegisterType((*PongResponse)(nil), "pb.PongResponse")
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*RegisterRequest)(nil), "pb.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "pb.RegisterResponse")
	proto.RegisterType((*LoginRequest)(nil), "pb.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "pb.LoginResponse")
	proto.RegisterType((*CheckRequest)(nil), "pb.CheckRequest")
	proto.RegisterType((*CheckResponse)(nil), "pb.CheckResponse")
	proto.RegisterType((*LogoutRequest)(nil), "pb.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "pb.LogoutResponse")
	proto.RegisterEnum("pb.Sex", Sex_name, Sex_value)
	proto.RegisterEnum("pb.LoginMethod", LoginMethod_name, LoginMethod_value)
}

func init() { proto.RegisterFile("proto/auth/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 644 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0xad, 0x1d, 0x27, 0x71, 0x26, 0x69, 0x9a, 0x6f, 0xbf, 0x22, 0x4c, 0x4a, 0x45, 0x31, 0x97,
	0xaa, 0x48, 0x41, 0x84, 0x03, 0xe7, 0xaa, 0xb4, 0x50, 0xa9, 0x2d, 0x95, 0x5b, 0xce, 0x91, 0x1b,
	0x0f, 0x89, 0x95, 0xc4, 0x1b, 0xbc, 0x9b, 0x92, 0x0b, 0xff, 0x80, 0xdf, 0xc9, 0x91, 0xbf, 0x00,
	0xda, 0xd9, 0x6c, 0xb2, 0xb1, 0x4a, 0x2f, 0x54, 0x5c, 0xa2, 0xcc, 0x9b, 0xd9, 0x99, 0x79, 0xcf,
	0xcf, 0x6b, 0x78, 0x34, 0xcd, 0xb9, 0xe4, 0xaf, 0xe2, 0x99, 0x1c, 0xd2, 0x4f, 0x87, 0x62, 0xe6,
	0x4e, 0x6f, 0xc2, 0x4d, 0xa8, 0x5f, 0xa6, 0xd9, 0x20, 0xc2, 0x2f, 0x33, 0x14, 0x32, 0xbc, 0x82,
	0xc6, 0x25, 0x57, 0xa1, 0x98, 0xf2, 0x4c, 0x20, 0x7b, 0x0e, 0x0d, 0x81, 0xf9, 0x6d, 0xda, 0xc7,
	0x5e, 0x16, 0x4f, 0x30, 0x70, 0xf6, 0x9c, 0xfd, 0x5a, 0x54, 0x5f, 0x60, 0x17, 0xf1, 0x04, 0xd9,
	0x1e, 0xd4, 0x31, 0xbb, 0x4d, 0x73, 0x9e, 0x4d, 0x30, 0x93, 0x81, 0xab, 0x2b, 0x2c, 0x28, 0xfc,
	0xe9, 0x80, 0xf7, 0x49, 0x60, 0xce, 0x9a, 0xe0, 0xa6, 0x09, 0xf5, 0x28, 0x45, 0x6e, 0x9a, 0xb0,
	0x5d, 0x80, 0xcf, 0x69, 0x2e, 0xa4, 0xee, 0xad, 0x4f, 0xd6, 0x08, 0xa1, 0xce, 0x3b, 0x50, 0x1b,
	0xc7, 0x26, 0x5b, 0xa2, 0xac, 0xaf, 0x00, 0x4a, 0xb6, 0xc1, 0xcf, 0xd2, 0xfe, 0x88, 0x72, 0x9e,
	0xce, 0x99, 0x98, 0x05, 0x50, 0x8d, 0x93, 0x24, 0x47, 0x21, 0x82, 0x32, 0xa5, 0x4c, 0xc8, 0xb6,
	0xa1, 0x3c, 0x1d, 0xf2, 0x0c, 0x83, 0x0a, 0xe1, 0x3a, 0x50, 0x28, 0x4e, 0xe2, 0x74, 0x1c, 0x54,
	0x35, 0x4a, 0x01, 0x7b, 0x02, 0x25, 0x81, 0xf3, 0xc0, 0xdf, 0x73, 0xf6, 0x9b, 0xdd, 0x6a, 0x67,
	0x7a, 0xd3, 0xb9, 0xc2, 0x79, 0xa4, 0x30, 0xb5, 0xb8, 0xd2, 0xb1, 0x27, 0xf9, 0x08, 0xb3, 0xa0,
	0xa6, 0x17, 0x57, 0xc8, 0xb5, 0x02, 0xc2, 0x1f, 0x0e, 0x6c, 0x45, 0x38, 0x48, 0x85, 0xc4, 0x7c,
	0xa1, 0x6c, 0x81, 0xab, 0x73, 0x2f, 0x57, 0xf7, 0x1e, 0xae, 0xa5, 0x02, 0xd7, 0x36, 0xf8, 0xd3,
	0x58, 0x88, 0xaf, 0x3c, 0x4f, 0x8c, 0x0e, 0x26, 0x5e, 0xb1, 0x2d, 0xdf, 0xc9, 0xb6, 0x62, 0xb3,
	0xb5, 0x34, 0xab, 0xae, 0x6b, 0xf6, 0x67, 0x1d, 0xc2, 0x0c, 0x5a, 0x2b, 0x9e, 0x0b, 0xcb, 0xec,
	0x02, 0x60, 0x9e, 0xf3, 0xbc, 0xd7, 0xe7, 0xc9, 0x92, 0x28, 0x21, 0x47, 0x3c, 0xb1, 0xd2, 0x12,
	0xe7, 0xc6, 0x2d, 0x3a, 0x7d, 0x8d, 0x73, 0xc9, 0x9e, 0x82, 0x37, 0x13, 0x98, 0x13, 0xcd, 0x7a,
	0xd7, 0x57, 0xd3, 0x94, 0x75, 0x22, 0x42, 0xc3, 0x6f, 0xd0, 0x38, 0xe3, 0x83, 0x34, 0xb3, 0x44,
	0x1d, 0xab, 0xb8, 0x97, 0xc4, 0x32, 0x36, 0xb3, 0x08, 0x79, 0x17, 0xcb, 0x78, 0x4d, 0x1b, 0xb7,
	0xa0, 0x4d, 0x17, 0x1a, 0xfa, 0xe8, 0x04, 0xe5, 0x90, 0x27, 0x34, 0xb0, 0xd9, 0xdd, 0x52, 0x03,
	0x69, 0xc4, 0x39, 0xc1, 0x51, 0x7d, 0xbc, 0x0a, 0xc2, 0xef, 0x0e, 0x6c, 0x2e, 0xe6, 0x3f, 0x08,
	0xd9, 0x67, 0xa0, 0xdb, 0xf7, 0xfa, 0x43, 0xec, 0x8f, 0x68, 0x05, 0x3f, 0xd2, 0x8c, 0x8e, 0x14,
	0xb2, 0x54, 0xc3, 0xbb, 0x53, 0x8d, 0x13, 0x68, 0x50, 0x99, 0x51, 0xe3, 0x31, 0x54, 0x15, 0xde,
	0x5b, 0xbe, 0x63, 0x15, 0x15, 0x9e, 0x26, 0x05, 0xbb, 0xba, 0x45, 0xbb, 0x8e, 0x60, 0x73, 0xd1,
	0xe7, 0x1f, 0x3c, 0xc2, 0xf7, 0x24, 0x21, 0x9f, 0xc9, 0xbf, 0xdd, 0xfa, 0x02, 0x9a, 0xa6, 0xd1,
	0x43, 0xac, 0x7d, 0xb0, 0x03, 0xa5, 0x2b, 0x9c, 0x33, 0x1f, 0xbc, 0xf3, 0xc3, 0xb3, 0xe3, 0xd6,
	0x06, 0x03, 0xa8, 0x9c, 0x1c, 0xd3, 0x7f, 0xe7, 0xe0, 0x05, 0xd4, 0x2d, 0x57, 0xb0, 0x1a, 0x94,
	0x2f, 0x3f, 0x7c, 0xbc, 0x50, 0x55, 0x35, 0x28, 0x1f, 0x9f, 0x1f, 0x9e, 0x9e, 0xb5, 0x9c, 0xee,
	0x2f, 0x07, 0xbc, 0xc3, 0x99, 0x1c, 0xb2, 0x97, 0xe0, 0xa9, 0x4b, 0x95, 0x91, 0x9b, 0xac, 0xeb,
	0xb5, 0xdd, 0x22, 0xc0, 0xba, 0x60, 0xc3, 0x0d, 0xf6, 0x16, 0x7c, 0xf3, 0x0e, 0xb1, 0xff, 0x55,
	0xbe, 0x70, 0x73, 0xb4, 0xb7, 0xd7, 0xc1, 0xe5, 0xc1, 0x0e, 0x94, 0x69, 0x27, 0xd6, 0x5a, 0x9a,
	0xd6, 0x1c, 0xf9, 0xcf, 0x42, 0xec, 0x7a, 0xed, 0x2a, 0xaa, 0xb7, 0x9d, 0xa3, 0xeb, 0xd7, 0x3c,
	0x10, 0x6e, 0xb0, 0xd7, 0x50, 0xd1, 0x02, 0x33, 0xd3, 0x6e, 0xf5, 0xd4, 0xda, 0xcc, 0x86, 0xcc,
	0x91, 0x9b, 0x0a, 0x7d, 0x58, 0xde, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x16, 0x26, 0x34, 0xcb,
	0x71, 0x06, 0x00, 0x00,
}
