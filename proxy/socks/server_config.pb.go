// Code generated by protoc-gen-go.
// source: v2ray.com/core/proxy/socks/server_config.proto
// DO NOT EDIT!

/*
Package socks is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/proxy/socks/server_config.proto

It has these top-level messages:
	ServerConfig
*/
package socks

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import com_v2ray_core_common_net "v2ray.com/core/common/net"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServerConfig_AuthType int32

const (
	ServerConfig_NO_AUTH  ServerConfig_AuthType = 0
	ServerConfig_PASSWORD ServerConfig_AuthType = 1
)

var ServerConfig_AuthType_name = map[int32]string{
	0: "NO_AUTH",
	1: "PASSWORD",
}
var ServerConfig_AuthType_value = map[string]int32{
	"NO_AUTH":  0,
	"PASSWORD": 1,
}

func (x ServerConfig_AuthType) String() string {
	return proto.EnumName(ServerConfig_AuthType_name, int32(x))
}
func (ServerConfig_AuthType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type ServerConfig struct {
	AuthType   ServerConfig_AuthType                `protobuf:"varint,1,opt,name=auth_type,json=authType,enum=com.v2ray.core.proxy.socks.ServerConfig_AuthType" json:"auth_type,omitempty"`
	Accounts   map[string]string                    `protobuf:"bytes,2,rep,name=accounts" json:"accounts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Address    *com_v2ray_core_common_net.AddressPB `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	UdpEnabled bool                                 `protobuf:"varint,4,opt,name=udp_enabled,json=udpEnabled" json:"udp_enabled,omitempty"`
	Timeout    uint32                               `protobuf:"varint,5,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *ServerConfig) Reset()                    { *m = ServerConfig{} }
func (m *ServerConfig) String() string            { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()               {}
func (*ServerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServerConfig) GetAccounts() map[string]string {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *ServerConfig) GetAddress() *com_v2ray_core_common_net.AddressPB {
	if m != nil {
		return m.Address
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerConfig)(nil), "com.v2ray.core.proxy.socks.ServerConfig")
	proto.RegisterEnum("com.v2ray.core.proxy.socks.ServerConfig_AuthType", ServerConfig_AuthType_name, ServerConfig_AuthType_value)
}

func init() { proto.RegisterFile("v2ray.com/core/proxy/socks/server_config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0xd1, 0x4b, 0xeb, 0x30,
	0x14, 0xc6, 0x6f, 0xd7, 0xbb, 0xdb, 0x2e, 0xdd, 0x2e, 0x23, 0xf8, 0x50, 0xf6, 0x62, 0x19, 0x8a,
	0x7d, 0x4a, 0xb1, 0x82, 0x88, 0x82, 0xd0, 0xe9, 0xc0, 0xa7, 0x6d, 0x64, 0x13, 0xc1, 0x97, 0xd2,
	0xa5, 0x47, 0x37, 0xb6, 0x36, 0x25, 0x4d, 0x86, 0xfd, 0x33, 0xfc, 0x8f, 0xc5, 0x74, 0x1b, 0x3a,
	0x10, 0x7c, 0x3b, 0xe7, 0xf0, 0x7d, 0xbf, 0xe4, 0x7c, 0x07, 0x91, 0x4d, 0x28, 0x92, 0x8a, 0x30,
	0x9e, 0x05, 0x8c, 0x0b, 0x08, 0x0a, 0xc1, 0xdf, 0xaa, 0xa0, 0xe4, 0x6c, 0x55, 0x06, 0x25, 0x88,
	0x0d, 0x88, 0x98, 0xf1, 0xfc, 0x65, 0xf9, 0x4a, 0x0a, 0xc1, 0x25, 0xc7, 0x3d, 0xc6, 0xb3, 0xbd,
	0x47, 0x00, 0xd1, 0x7a, 0xa2, 0xf5, 0xbd, 0xb3, 0x03, 0x16, 0xe3, 0x59, 0xc6, 0xf3, 0x20, 0x07,
	0x19, 0x24, 0x69, 0x2a, 0xa0, 0x2c, 0x6b, 0x48, 0xff, 0xdd, 0x44, 0xed, 0xa9, 0x86, 0xdf, 0x69,
	0x36, 0x1e, 0xa1, 0x56, 0xa2, 0xe4, 0x22, 0x96, 0x55, 0x01, 0xae, 0xe1, 0x19, 0xfe, 0xff, 0xf0,
	0x9c, 0xfc, 0xfc, 0x12, 0xf9, 0x6a, 0x26, 0x91, 0x92, 0x8b, 0x59, 0x55, 0x00, 0xb5, 0x93, 0x6d,
	0x85, 0x29, 0xb2, 0x13, 0xc6, 0xb8, 0xca, 0x65, 0xe9, 0x36, 0x3c, 0xd3, 0x77, 0xc2, 0xcb, 0xdf,
	0xe3, 0xb6, 0xc6, 0x61, 0x2e, 0x45, 0x45, 0xf7, 0x1c, 0x7c, 0x8b, 0xac, 0xed, 0x16, 0xae, 0xe9,
	0x19, 0xbe, 0x13, 0x9e, 0x1c, 0x22, 0xeb, 0x7d, 0x49, 0x0e, 0x92, 0x44, 0xb5, 0x72, 0x32, 0xa0,
	0x3b, 0x13, 0x3e, 0x46, 0x8e, 0x4a, 0x8b, 0x18, 0xf2, 0x64, 0xbe, 0x86, 0xd4, 0xfd, 0xeb, 0x19,
	0xbe, 0x4d, 0x91, 0x4a, 0x8b, 0x61, 0x3d, 0xc1, 0x2e, 0xb2, 0xe4, 0x32, 0x03, 0xae, 0xa4, 0xdb,
	0xf4, 0x0c, 0xbf, 0x43, 0x77, 0x6d, 0xef, 0x06, 0x75, 0xbe, 0xfd, 0x0a, 0x77, 0x91, 0xb9, 0x82,
	0x4a, 0x27, 0xd5, 0xa2, 0x9f, 0x25, 0x3e, 0x42, 0xcd, 0x4d, 0xb2, 0x56, 0xe0, 0x36, 0xf4, 0xac,
	0x6e, 0xae, 0x1b, 0x57, 0x46, 0xff, 0x14, 0xd9, 0xbb, 0x84, 0xb0, 0x83, 0xac, 0xd1, 0x38, 0x8e,
	0x1e, 0x67, 0x0f, 0xdd, 0x3f, 0xb8, 0x8d, 0xec, 0x49, 0x34, 0x9d, 0x3e, 0x8d, 0xe9, 0x7d, 0xd7,
	0x18, 0x58, 0xcf, 0x4d, 0x1d, 0xc6, 0xfc, 0x9f, 0xbe, 0xd1, 0xc5, 0x47, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x1f, 0xec, 0x2f, 0xc7, 0x1a, 0x02, 0x00, 0x00,
}
