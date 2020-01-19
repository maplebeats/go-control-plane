// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/listener/proxy_protocol/v2/proxy_protocol.proto

package envoy_config_filter_listener_proxy_protocol_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type ProxyProtocol struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyProtocol) Reset()         { *m = ProxyProtocol{} }
func (m *ProxyProtocol) String() string { return proto.CompactTextString(m) }
func (*ProxyProtocol) ProtoMessage()    {}
func (*ProxyProtocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_614eea5d50c8eea0, []int{0}
}

func (m *ProxyProtocol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyProtocol.Unmarshal(m, b)
}
func (m *ProxyProtocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyProtocol.Marshal(b, m, deterministic)
}
func (m *ProxyProtocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyProtocol.Merge(m, src)
}
func (m *ProxyProtocol) XXX_Size() int {
	return xxx_messageInfo_ProxyProtocol.Size(m)
}
func (m *ProxyProtocol) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyProtocol.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyProtocol proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProxyProtocol)(nil), "envoy.config.filter.listener.proxy_protocol.v2.ProxyProtocol")
}

func init() {
	proto.RegisterFile("envoy/config/filter/listener/proxy_protocol/v2/proxy_protocol.proto", fileDescriptor_614eea5d50c8eea0)
}

var fileDescriptor_614eea5d50c8eea0 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0xc9, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x2d, 0xd2, 0x2f, 0x28, 0xca, 0xaf, 0xa8, 0x8c, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x33, 0x42, 0x13, 0xd1, 0x03, 0x33, 0x84,
	0xf4, 0xc0, 0x86, 0xe8, 0x41, 0x0c, 0xd1, 0x83, 0x18, 0xa2, 0x07, 0x33, 0x44, 0x0f, 0x4d, 0x4b,
	0x99, 0x91, 0x94, 0x5c, 0x69, 0x4a, 0x41, 0xa2, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49, 0x62, 0x49,
	0x66, 0x7e, 0x5e, 0xb1, 0x7e, 0x6e, 0x66, 0x7a, 0x51, 0x62, 0x49, 0x2a, 0xc4, 0x3c, 0x25, 0x7e,
	0x2e, 0xde, 0x00, 0x90, 0xa6, 0x00, 0xa8, 0x1e, 0xa7, 0x7e, 0xc6, 0x4f, 0x33, 0xfe, 0xf5, 0xb3,
	0x9a, 0x0a, 0x19, 0x43, 0x6c, 0x4a, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0x06, 0xe9, 0x84, 0xda, 0x56,
	0x8c, 0xdb, 0x3a, 0x63, 0x2e, 0x9b, 0xcc, 0x7c, 0x88, 0x0b, 0xc1, 0x52, 0x24, 0x3a, 0xd6, 0x49,
	0x08, 0xc5, 0x29, 0x60, 0x3a, 0x80, 0x31, 0x89, 0x0d, 0xac, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x99, 0x2b, 0xe5, 0xfc, 0x40, 0x01, 0x00, 0x00,
}