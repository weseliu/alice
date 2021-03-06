// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logicproto.proto

/*
Package logic is a generated protocol buffer package.

It is generated from these files:
	logicproto.proto

It has these top-level messages:
	LGC_PING_CS
	LGC_PING_SC
	LGC_KICKOFF_SC
	LGC_ACCOUNT_LOGIN_CS
	LGC_ACCOUNT_LOGIN_SC
	LGC_ACCOUNT_LOGOUT_CS
	LGC_ACCOUNT_LOGOUT_SC
*/
package logic

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

type LogicType int32

const (
	LogicType_LGC_ZERO_PLACEHOLDER LogicType = 0
	LogicType_LGC_PING             LogicType = 100
	LogicType_LGC_KICKOFF          LogicType = 101
	LogicType_LGC_ACCOUNT_LOGIN    LogicType = 200
	LogicType_LGC_ACCOUNT_LOGOUT   LogicType = 201
)

var LogicType_name = map[int32]string{
	0:   "LGC_ZERO_PLACEHOLDER",
	100: "LGC_PING",
	101: "LGC_KICKOFF",
	200: "LGC_ACCOUNT_LOGIN",
	201: "LGC_ACCOUNT_LOGOUT",
}
var LogicType_value = map[string]int32{
	"LGC_ZERO_PLACEHOLDER": 0,
	"LGC_PING":             100,
	"LGC_KICKOFF":          101,
	"LGC_ACCOUNT_LOGIN":    200,
	"LGC_ACCOUNT_LOGOUT":   201,
}

func (x LogicType) String() string {
	return proto.EnumName(LogicType_name, int32(x))
}
func (LogicType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LGC_PING_CS struct {
	ClientTick uint32 `protobuf:"varint,1,opt,name=client_tick,json=clientTick" json:"client_tick,omitempty"`
	LastDelay  uint32 `protobuf:"varint,2,opt,name=last_delay,json=lastDelay" json:"last_delay,omitempty"`
}

func (m *LGC_PING_CS) Reset()                    { *m = LGC_PING_CS{} }
func (m *LGC_PING_CS) String() string            { return proto.CompactTextString(m) }
func (*LGC_PING_CS) ProtoMessage()               {}
func (*LGC_PING_CS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LGC_PING_CS) GetClientTick() uint32 {
	if m != nil {
		return m.ClientTick
	}
	return 0
}

func (m *LGC_PING_CS) GetLastDelay() uint32 {
	if m != nil {
		return m.LastDelay
	}
	return 0
}

type LGC_PING_SC struct {
	ClientTick uint32 `protobuf:"varint,1,opt,name=client_tick,json=clientTick" json:"client_tick,omitempty"`
}

func (m *LGC_PING_SC) Reset()                    { *m = LGC_PING_SC{} }
func (m *LGC_PING_SC) String() string            { return proto.CompactTextString(m) }
func (*LGC_PING_SC) ProtoMessage()               {}
func (*LGC_PING_SC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LGC_PING_SC) GetClientTick() uint32 {
	if m != nil {
		return m.ClientTick
	}
	return 0
}

type LGC_KICKOFF_SC struct {
	Reason  string `protobuf:"bytes,1,opt,name=reason" json:"reason,omitempty"`
	HintMsg string `protobuf:"bytes,2,opt,name=hint_msg,json=hintMsg" json:"hint_msg,omitempty"`
}

func (m *LGC_KICKOFF_SC) Reset()                    { *m = LGC_KICKOFF_SC{} }
func (m *LGC_KICKOFF_SC) String() string            { return proto.CompactTextString(m) }
func (*LGC_KICKOFF_SC) ProtoMessage()               {}
func (*LGC_KICKOFF_SC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LGC_KICKOFF_SC) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *LGC_KICKOFF_SC) GetHintMsg() string {
	if m != nil {
		return m.HintMsg
	}
	return ""
}

type LGC_ACCOUNT_LOGIN_CS struct {
	PlatformType  int32  `protobuf:"varint,1,opt,name=platform_type,json=platformType" json:"platform_type,omitempty"`
	ClientVersion uint32 `protobuf:"varint,2,opt,name=client_version,json=clientVersion" json:"client_version,omitempty"`
	ClientChannel string `protobuf:"bytes,3,opt,name=client_channel,json=clientChannel" json:"client_channel,omitempty"`
	OperationType int32  `protobuf:"varint,4,opt,name=operation_type,json=operationType" json:"operation_type,omitempty"`
	Name          string `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	ClientPkgName string `protobuf:"bytes,6,opt,name=client_pkg_name,json=clientPkgName" json:"client_pkg_name,omitempty"`
	ClientAppName string `protobuf:"bytes,7,opt,name=client_app_name,json=clientAppName" json:"client_app_name,omitempty"`
	MobileType    string `protobuf:"bytes,8,opt,name=mobile_type,json=mobileType" json:"mobile_type,omitempty"`
	DeviceId      string `protobuf:"bytes,9,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
}

func (m *LGC_ACCOUNT_LOGIN_CS) Reset()                    { *m = LGC_ACCOUNT_LOGIN_CS{} }
func (m *LGC_ACCOUNT_LOGIN_CS) String() string            { return proto.CompactTextString(m) }
func (*LGC_ACCOUNT_LOGIN_CS) ProtoMessage()               {}
func (*LGC_ACCOUNT_LOGIN_CS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LGC_ACCOUNT_LOGIN_CS) GetPlatformType() int32 {
	if m != nil {
		return m.PlatformType
	}
	return 0
}

func (m *LGC_ACCOUNT_LOGIN_CS) GetClientVersion() uint32 {
	if m != nil {
		return m.ClientVersion
	}
	return 0
}

func (m *LGC_ACCOUNT_LOGIN_CS) GetClientChannel() string {
	if m != nil {
		return m.ClientChannel
	}
	return ""
}

func (m *LGC_ACCOUNT_LOGIN_CS) GetOperationType() int32 {
	if m != nil {
		return m.OperationType
	}
	return 0
}

func (m *LGC_ACCOUNT_LOGIN_CS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LGC_ACCOUNT_LOGIN_CS) GetClientPkgName() string {
	if m != nil {
		return m.ClientPkgName
	}
	return ""
}

func (m *LGC_ACCOUNT_LOGIN_CS) GetClientAppName() string {
	if m != nil {
		return m.ClientAppName
	}
	return ""
}

func (m *LGC_ACCOUNT_LOGIN_CS) GetMobileType() string {
	if m != nil {
		return m.MobileType
	}
	return ""
}

func (m *LGC_ACCOUNT_LOGIN_CS) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

type LGC_ACCOUNT_LOGIN_SC struct {
	RetCode     int32  `protobuf:"varint,1,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	RetCodeDesc string `protobuf:"bytes,2,opt,name=ret_code_desc,json=retCodeDesc" json:"ret_code_desc,omitempty"`
}

func (m *LGC_ACCOUNT_LOGIN_SC) Reset()                    { *m = LGC_ACCOUNT_LOGIN_SC{} }
func (m *LGC_ACCOUNT_LOGIN_SC) String() string            { return proto.CompactTextString(m) }
func (*LGC_ACCOUNT_LOGIN_SC) ProtoMessage()               {}
func (*LGC_ACCOUNT_LOGIN_SC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LGC_ACCOUNT_LOGIN_SC) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *LGC_ACCOUNT_LOGIN_SC) GetRetCodeDesc() string {
	if m != nil {
		return m.RetCodeDesc
	}
	return ""
}

type LGC_ACCOUNT_LOGOUT_CS struct {
}

func (m *LGC_ACCOUNT_LOGOUT_CS) Reset()                    { *m = LGC_ACCOUNT_LOGOUT_CS{} }
func (m *LGC_ACCOUNT_LOGOUT_CS) String() string            { return proto.CompactTextString(m) }
func (*LGC_ACCOUNT_LOGOUT_CS) ProtoMessage()               {}
func (*LGC_ACCOUNT_LOGOUT_CS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type LGC_ACCOUNT_LOGOUT_SC struct {
}

func (m *LGC_ACCOUNT_LOGOUT_SC) Reset()                    { *m = LGC_ACCOUNT_LOGOUT_SC{} }
func (m *LGC_ACCOUNT_LOGOUT_SC) String() string            { return proto.CompactTextString(m) }
func (*LGC_ACCOUNT_LOGOUT_SC) ProtoMessage()               {}
func (*LGC_ACCOUNT_LOGOUT_SC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*LGC_PING_CS)(nil), "logic.LGC_PING_CS")
	proto.RegisterType((*LGC_PING_SC)(nil), "logic.LGC_PING_SC")
	proto.RegisterType((*LGC_KICKOFF_SC)(nil), "logic.LGC_KICKOFF_SC")
	proto.RegisterType((*LGC_ACCOUNT_LOGIN_CS)(nil), "logic.LGC_ACCOUNT_LOGIN_CS")
	proto.RegisterType((*LGC_ACCOUNT_LOGIN_SC)(nil), "logic.LGC_ACCOUNT_LOGIN_SC")
	proto.RegisterType((*LGC_ACCOUNT_LOGOUT_CS)(nil), "logic.LGC_ACCOUNT_LOGOUT_CS")
	proto.RegisterType((*LGC_ACCOUNT_LOGOUT_SC)(nil), "logic.LGC_ACCOUNT_LOGOUT_SC")
	proto.RegisterEnum("logic.LogicType", LogicType_name, LogicType_value)
}

func init() { proto.RegisterFile("logicproto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0xa5, 0x63, 0xfd, 0xc9, 0xd7, 0x65, 0x0b, 0x16, 0x6c, 0x41, 0x08, 0x0d, 0x05, 0x81, 0x10,
	0x17, 0xbb, 0xe1, 0x09, 0x2a, 0xb7, 0x2b, 0xd5, 0xb2, 0xa6, 0x4a, 0x5a, 0x2e, 0xb8, 0xb1, 0xb2,
	0xc4, 0x64, 0x56, 0xd3, 0xd8, 0x4a, 0xcc, 0xa4, 0xbe, 0x21, 0xdc, 0xf2, 0x44, 0xc8, 0x9f, 0x53,
	0x54, 0x28, 0x88, 0x9b, 0x28, 0x3e, 0xe7, 0xe4, 0x9c, 0x63, 0xfb, 0x0b, 0x78, 0xa5, 0x2c, 0x44,
	0xa6, 0x6a, 0xa9, 0xe5, 0x15, 0x3e, 0x49, 0x17, 0x91, 0xe0, 0x16, 0x86, 0xe1, 0x94, 0xb2, 0xc5,
	0x6c, 0x3e, 0x65, 0x34, 0x21, 0x97, 0x30, 0xcc, 0x4a, 0xc1, 0x2b, 0xcd, 0xb4, 0xc8, 0xd6, 0x7e,
	0xe7, 0x55, 0xe7, 0x9d, 0x1b, 0x83, 0x85, 0x96, 0x22, 0x5b, 0x93, 0x97, 0x00, 0x65, 0xda, 0x68,
	0x96, 0xf3, 0x32, 0xdd, 0xfa, 0x47, 0xc8, 0x3b, 0x06, 0x19, 0x1b, 0x20, 0xb8, 0xda, 0xb3, 0x4b,
	0xe8, 0x7f, 0xed, 0x02, 0x0a, 0xa7, 0x46, 0x7f, 0x33, 0xa3, 0x37, 0xd1, 0xf5, 0xb5, 0xf9, 0xe4,
	0x1c, 0x7a, 0x35, 0x4f, 0x1b, 0x59, 0xa1, 0xda, 0x89, 0xdb, 0x15, 0x79, 0x0e, 0x83, 0x7b, 0x51,
	0x69, 0xb6, 0x69, 0x0a, 0x8c, 0x75, 0xe2, 0xbe, 0x59, 0xdf, 0x36, 0x45, 0xf0, 0xe3, 0x08, 0x9e,
	0x1a, 0x97, 0x11, 0xa5, 0xd1, 0x6a, 0xbe, 0x64, 0x61, 0x34, 0x9d, 0xcd, 0xcd, 0x6e, 0x5e, 0x83,
	0xab, 0xca, 0x54, 0x7f, 0x91, 0xf5, 0x86, 0xe9, 0xad, 0xe2, 0x68, 0xd9, 0x8d, 0x4f, 0x76, 0xe0,
	0x72, 0xab, 0x38, 0x79, 0x03, 0xa7, 0x6d, 0xc7, 0x07, 0x5e, 0x37, 0x42, 0x56, 0xed, 0xae, 0x5c,
	0x8b, 0x7e, 0xb2, 0xe0, 0x9e, 0x2c, 0xbb, 0x4f, 0xab, 0x8a, 0x97, 0xfe, 0x63, 0x6c, 0xd1, 0xca,
	0xa8, 0x05, 0x8d, 0x4c, 0x2a, 0x5e, 0xa7, 0x5a, 0xc8, 0xca, 0x66, 0x1e, 0x63, 0xa6, 0xfb, 0x0b,
	0xc5, 0x50, 0x02, 0xc7, 0x55, 0xba, 0xe1, 0x7e, 0x17, 0x3d, 0xf0, 0x9d, 0xbc, 0x85, 0xb3, 0x36,
	0x41, 0xad, 0x0b, 0x86, 0x74, 0x6f, 0x3f, 0x62, 0xb1, 0x2e, 0xe6, 0xbf, 0xeb, 0x52, 0xa5, 0xac,
	0xae, 0xbf, 0xaf, 0x1b, 0x29, 0x85, 0xba, 0x4b, 0x18, 0x6e, 0xe4, 0x9d, 0x28, 0xb9, 0xed, 0x31,
	0x40, 0x0d, 0x58, 0x08, 0x4b, 0xbc, 0x00, 0x27, 0xe7, 0x0f, 0x22, 0xe3, 0x4c, 0xe4, 0xbe, 0x83,
	0xf4, 0xc0, 0x02, 0xb3, 0x3c, 0x58, 0xfd, 0xed, 0x4c, 0x13, 0x6a, 0xee, 0xa1, 0xe6, 0x9a, 0x65,
	0x32, 0xdf, 0x1d, 0x67, 0xbf, 0xe6, 0x9a, 0xca, 0x9c, 0x93, 0x00, 0xdc, 0x1d, 0xc5, 0x72, 0xde,
	0x64, 0xed, 0x3d, 0x0d, 0x5b, 0x7e, 0xcc, 0x9b, 0x2c, 0xb8, 0x80, 0x67, 0x7f, 0xd8, 0x46, 0xab,
	0x25, 0xa3, 0xc9, 0x3f, 0x88, 0x84, 0xbe, 0xff, 0x0a, 0x4e, 0x68, 0x46, 0x15, 0x2b, 0xfb, 0xb6,
	0xd5, 0xe7, 0x49, 0x1c, 0xb1, 0x45, 0x38, 0xa2, 0x93, 0x8f, 0x51, 0x38, 0x9e, 0xc4, 0xde, 0x23,
	0x72, 0x02, 0x83, 0xdd, 0xe4, 0x79, 0x39, 0x39, 0xb3, 0x73, 0xd8, 0xce, 0x95, 0xc7, 0xc9, 0x39,
	0x3c, 0x39, 0xd8, 0x8e, 0xf7, 0xad, 0x43, 0x2e, 0x80, 0x1c, 0xc6, 0x7a, 0xdf, 0x3b, 0x77, 0x3d,
	0xfc, 0x4d, 0x3e, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x51, 0xaa, 0x6d, 0x22, 0x3a, 0x03, 0x00,
	0x00,
}
