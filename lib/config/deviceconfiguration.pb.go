// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/config/deviceconfiguration.proto

package config

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	github_com_syncthing_syncthing_lib_protocol "github.com/syncthing/syncthing/lib/protocol"
	protocol "github.com/syncthing/syncthing/lib/protocol"
	_ "github.com/syncthing/syncthing/proto/ext"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DeviceConfiguration struct {
	DeviceID                 github_com_syncthing_syncthing_lib_protocol.DeviceID `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3,customtype=github.com/syncthing/syncthing/lib/protocol.DeviceID" json:"deviceID" xml:"id,attr"`
	Name                     string                                               `protobuf:"bytes,2,opt,name=name,proto3" json:"name" xml:"name,attr,omitempty"`
	Addresses                []string                                             `protobuf:"bytes,3,rep,name=addresses,proto3" json:"addresses" xml:"address,omitempty" default:"dynamic"`
	Compression              protocol.Compression                                 `protobuf:"varint,4,opt,name=compression,proto3,enum=protocol.Compression" json:"compression" xml:"compression,attr"`
	CertName                 string                                               `protobuf:"bytes,5,opt,name=cert_name,json=certName,proto3" json:"certName" xml:"certName,attr,omitempty"`
	Introducer               bool                                                 `protobuf:"varint,6,opt,name=introducer,proto3" json:"introducer" xml:"introducer,attr"`
	SkipIntroductionRemovals bool                                                 `protobuf:"varint,7,opt,name=skip_introduction_removals,json=skipIntroductionRemovals,proto3" json:"skipIntroductionRemovals" xml:"skipIntroductionRemovals,attr"`
	IntroducedBy             github_com_syncthing_syncthing_lib_protocol.DeviceID `protobuf:"bytes,8,opt,name=introduced_by,json=introducedBy,proto3,customtype=github.com/syncthing/syncthing/lib/protocol.DeviceID" json:"introducedBy" xml:"introducedBy,attr"`
	Paused                   bool                                                 `protobuf:"varint,9,opt,name=paused,proto3" json:"paused" xml:"paused"`
	AllowedNetworks          []string                                             `protobuf:"bytes,10,rep,name=allowed_networks,json=allowedNetworks,proto3" json:"allowedNetworks" xml:"allowedNetwork,omitempty"`
	AutoAcceptFolders        bool                                                 `protobuf:"varint,11,opt,name=auto_accept_folders,json=autoAcceptFolders,proto3" json:"autoAcceptFolders" xml:"autoAcceptFolders"`
	MaxSendKbps              int                                                  `protobuf:"varint,12,opt,name=max_send_kbps,json=maxSendKbps,proto3,casttype=int" json:"maxSendKbps" xml:"maxSendKbps"`
	MaxRecvKbps              int                                                  `protobuf:"varint,13,opt,name=max_recv_kbps,json=maxRecvKbps,proto3,casttype=int" json:"maxRecvKbps" xml:"maxRecvKbps"`
	IgnoredFolders           []ObservedFolder                                     `protobuf:"bytes,14,rep,name=ignored_folders,json=ignoredFolders,proto3" json:"ignoredFolders" xml:"ignoredFolder"`
	PendingFolders           []ObservedFolder                                     `protobuf:"bytes,15,rep,name=pending_folders,json=pendingFolders,proto3" json:"pendingFolders" xml:"pendingFolder"`
	MaxRequestKiB            int                                                  `protobuf:"varint,16,opt,name=max_request_kib,json=maxRequestKib,proto3,casttype=int" json:"maxRequestKiB" xml:"maxRequestKiB"`
	Untrusted                bool                                                 `protobuf:"varint,17,opt,name=untrusted,proto3" json:"untrusted" xml:"untrusted"`
}

func (m *DeviceConfiguration) Reset()         { *m = DeviceConfiguration{} }
func (m *DeviceConfiguration) String() string { return proto.CompactTextString(m) }
func (*DeviceConfiguration) ProtoMessage()    {}
func (*DeviceConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_744b782bd13071dd, []int{0}
}
func (m *DeviceConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceConfiguration.Unmarshal(m, b)
}
func (m *DeviceConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceConfiguration.Marshal(b, m, deterministic)
}
func (m *DeviceConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceConfiguration.Merge(m, src)
}
func (m *DeviceConfiguration) XXX_Size() int {
	return xxx_messageInfo_DeviceConfiguration.Size(m)
}
func (m *DeviceConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceConfiguration proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DeviceConfiguration)(nil), "config.DeviceConfiguration")
}

func init() {
	proto.RegisterFile("lib/config/deviceconfiguration.proto", fileDescriptor_744b782bd13071dd)
}

var fileDescriptor_744b782bd13071dd = []byte{
	// 929 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xbf, 0x6f, 0xdb, 0x46,
	0x03, 0x15, 0x3f, 0x27, 0xb6, 0x75, 0xb6, 0x2c, 0x8b, 0x46, 0x1c, 0xc6, 0x40, 0x74, 0x04, 0x3f,
	0x0d, 0x0a, 0x9a, 0xca, 0x85, 0xdb, 0xc9, 0x68, 0x0b, 0x94, 0x09, 0xda, 0x06, 0x41, 0x93, 0xf6,
	0xba, 0x79, 0x61, 0x49, 0xde, 0x59, 0x39, 0x58, 0xfc, 0x51, 0xf2, 0xa8, 0x48, 0x40, 0x87, 0x8e,
	0x05, 0xda, 0xa1, 0xc8, 0xda, 0xa5, 0xe8, 0xd0, 0xa1, 0x7f, 0x49, 0x36, 0x6b, 0x2c, 0x3a, 0x1c,
	0x10, 0x7b, 0x29, 0x38, 0x72, 0xcc, 0x54, 0xf0, 0x8e, 0xa2, 0x48, 0x3a, 0x2e, 0x02, 0x74, 0xbb,
	0x7b, 0xef, 0xdd, 0x7b, 0xc7, 0xa7, 0x3b, 0x1d, 0x18, 0x4c, 0xa8, 0x73, 0xe8, 0x06, 0xfe, 0x29,
	0x1d, 0x1f, 0x62, 0x32, 0xa5, 0x2e, 0x91, 0x93, 0x24, 0xb2, 0x19, 0x0d, 0xfc, 0x51, 0x18, 0x05,
	0x2c, 0x50, 0xd7, 0x25, 0x78, 0xb0, 0x9f, 0xab, 0x05, 0xe4, 0x06, 0x93, 0x43, 0x87, 0x84, 0x92,
	0x3f, 0xb8, 0x53, 0x71, 0x09, 0x9c, 0x98, 0x44, 0x53, 0x82, 0x0b, 0xaa, 0x4d, 0x66, 0x4c, 0x0e,
	0x8d, 0x1f, 0xbb, 0x60, 0xef, 0xa1, 0xc8, 0x78, 0x50, 0xcd, 0x50, 0xff, 0x50, 0x40, 0x5b, 0x66,
	0x5b, 0x14, 0x6b, 0x8a, 0xae, 0x0c, 0xb7, 0xcd, 0x9f, 0x94, 0x97, 0x1c, 0xb6, 0xfe, 0xe2, 0xf0,
	0x83, 0x31, 0x65, 0xcf, 0x12, 0x67, 0xe4, 0x06, 0xde, 0x61, 0x3c, 0xf7, 0x5d, 0xf6, 0x8c, 0xfa,
	0xe3, 0xca, 0xa8, 0xba, 0xa3, 0x91, 0x74, 0x7f, 0xf4, 0xf0, 0x82, 0xc3, 0xcd, 0xe5, 0x38, 0xe5,
	0x70, 0x13, 0x17, 0xe3, 0x8c, 0xc3, 0xce, 0xcc, 0x9b, 0x1c, 0x1b, 0x14, 0xdf, 0xb7, 0x19, 0x8b,
	0x8c, 0xf4, 0x7c, 0xb0, 0x51, 0x8c, 0xb3, 0xf3, 0x41, 0xa9, 0xfb, 0x61, 0x31, 0x50, 0x5e, 0x2c,
	0x06, 0xa5, 0x07, 0x5a, 0x32, 0x58, 0xfd, 0x12, 0xdc, 0xf0, 0x6d, 0x8f, 0x68, 0xff, 0xd3, 0x95,
	0x61, 0xdb, 0xfc, 0x30, 0xe5, 0x50, 0xcc, 0x33, 0x0e, 0xef, 0x08, 0xe7, 0x7c, 0x22, 0xfc, 0xee,
	0x07, 0x1e, 0x65, 0xc4, 0x0b, 0xd9, 0x3c, 0x4f, 0xd9, 0x7b, 0x03, 0x8e, 0xc4, 0x4a, 0x75, 0x06,
	0xda, 0x36, 0xc6, 0x11, 0x89, 0x63, 0x12, 0x6b, 0x6b, 0xfa, 0xda, 0xb0, 0x6d, 0x9e, 0xa4, 0x1c,
	0xae, 0xc0, 0x8c, 0xc3, 0x7b, 0xc2, 0xbb, 0x40, 0x2a, 0xce, 0x3a, 0x26, 0xa7, 0x76, 0x32, 0x61,
	0xc7, 0x06, 0x9e, 0xfb, 0xb6, 0x47, 0xdd, 0x3c, 0xab, 0x77, 0x45, 0xf7, 0xfa, 0x7c, 0xb0, 0x51,
	0x08, 0xd0, 0xca, 0x57, 0x9d, 0x82, 0x2d, 0x37, 0xf0, 0xc2, 0x7c, 0x46, 0x03, 0x5f, 0xbb, 0xa1,
	0x2b, 0xc3, 0x9d, 0xa3, 0x5b, 0xa3, 0xb2, 0xce, 0x07, 0x2b, 0xd2, 0xfc, 0x28, 0xe5, 0xb0, 0xaa,
	0xce, 0x38, 0xdc, 0x17, 0x9b, 0xaa, 0x60, 0x65, 0xa7, 0xbb, 0x4d, 0x10, 0x55, 0x97, 0xaa, 0x04,
	0xb4, 0x5d, 0x12, 0x31, 0x4b, 0x14, 0x79, 0x53, 0x14, 0xf9, 0x79, 0xfe, 0x33, 0xe5, 0xe0, 0x13,
	0x59, 0xe6, 0x5d, 0xe9, 0x5d, 0x00, 0x6f, 0x28, 0xf4, 0xf6, 0x35, 0x1c, 0x2a, 0x5d, 0xd4, 0x13,
	0x00, 0xa8, 0xcf, 0xa2, 0x00, 0x27, 0x2e, 0x89, 0xb4, 0x75, 0x5d, 0x19, 0x6e, 0x9a, 0xc7, 0x29,
	0x87, 0x15, 0x34, 0xe3, 0xf0, 0x96, 0x3c, 0x10, 0x25, 0x54, 0x7e, 0x44, 0xb7, 0x81, 0xa1, 0xca,
	0x3a, 0xf5, 0x37, 0x05, 0x1c, 0xc4, 0x67, 0x34, 0xb4, 0x96, 0x58, 0x7e, 0x92, 0xad, 0x88, 0x78,
	0xc1, 0xd4, 0x9e, 0xc4, 0xda, 0x86, 0x08, 0xc3, 0x29, 0x87, 0x5a, 0xae, 0x7a, 0x54, 0x11, 0xa1,
	0x42, 0x93, 0x71, 0xf8, 0x7f, 0x11, 0x7d, 0x9d, 0xa0, 0xdc, 0xc8, 0xdd, 0x7f, 0x55, 0xa0, 0x6b,
	0x13, 0xd4, 0xdf, 0x15, 0xd0, 0x29, 0xf7, 0x8c, 0x2d, 0x67, 0xae, 0x6d, 0x8a, 0xcb, 0xf5, 0xfd,
	0x7f, 0xba, 0x5c, 0x29, 0x87, 0xdb, 0x2b, 0x57, 0x73, 0x9e, 0x71, 0x78, 0xbb, 0xde, 0x21, 0x36,
	0xe7, 0xe5, 0xe6, 0x7b, 0x57, 0xd0, 0xfc, 0x72, 0xa1, 0x9a, 0x83, 0x7a, 0x04, 0xd6, 0x43, 0x3b,
	0x89, 0x09, 0xd6, 0xda, 0xa2, 0xb8, 0x83, 0x94, 0xc3, 0x02, 0xc9, 0x38, 0xdc, 0x16, 0xee, 0x72,
	0x6a, 0xa0, 0x02, 0x57, 0xbf, 0x03, 0xbb, 0xf6, 0x64, 0x12, 0x3c, 0x27, 0xd8, 0xf2, 0x09, 0x7b,
	0x1e, 0x44, 0x67, 0xb1, 0x06, 0xc4, 0xed, 0xf9, 0x2a, 0xe5, 0xb0, 0x5b, 0x70, 0x4f, 0x0a, 0x2a,
	0xe3, 0xb0, 0x2f, 0xef, 0x50, 0x0d, 0xaf, 0x9f, 0x29, 0xed, 0x3a, 0x12, 0x35, 0xed, 0xd4, 0x6f,
	0xc0, 0x9e, 0x9d, 0xb0, 0xc0, 0xb2, 0x5d, 0x97, 0x84, 0xcc, 0x3a, 0x0d, 0x26, 0x98, 0x44, 0xb1,
	0xb6, 0x25, 0xb6, 0xff, 0x5e, 0xca, 0x61, 0x2f, 0xa7, 0x3f, 0x11, 0xec, 0xa7, 0x92, 0x2c, 0x7b,
	0xba, 0xc2, 0x18, 0xe8, 0xaa, 0x5a, 0x7d, 0x0a, 0x3a, 0x9e, 0x3d, 0xb3, 0x62, 0xe2, 0x63, 0xeb,
	0xcc, 0x09, 0x63, 0x6d, 0x5b, 0x57, 0x86, 0x37, 0xcd, 0x77, 0xf2, 0x7b, 0xe8, 0xd9, 0xb3, 0xaf,
	0x89, 0x8f, 0x1f, 0x3b, 0x61, 0xee, 0xda, 0x13, 0xae, 0x15, 0xcc, 0x78, 0xcd, 0xe1, 0x1a, 0xf5,
	0x19, 0xaa, 0x0a, 0x97, 0x86, 0x11, 0x71, 0xa7, 0xd2, 0xb0, 0x53, 0x33, 0x44, 0xc4, 0x9d, 0x36,
	0x0d, 0x97, 0x58, 0xcd, 0x70, 0x09, 0xaa, 0x3e, 0xe8, 0xd2, 0xb1, 0x1f, 0x44, 0x04, 0x97, 0xdf,
	0xbf, 0xa3, 0xaf, 0x0d, 0xb7, 0x8e, 0xf6, 0x47, 0xf2, 0x2d, 0x18, 0x3d, 0x2d, 0xde, 0x02, 0xf9,
	0x4d, 0xe6, 0xbb, 0xf9, 0xb1, 0x4b, 0x39, 0xdc, 0x29, 0x96, 0xad, 0x8a, 0xd9, 0x93, 0x07, 0xa8,
	0x0a, 0x1b, 0xa8, 0x21, 0xcb, 0xf3, 0x42, 0xe2, 0x63, 0xea, 0x8f, 0xcb, 0xbc, 0xee, 0xdb, 0xe5,
	0x15, 0xcb, 0x9a, 0x79, 0x35, 0xd8, 0x40, 0x0d, 0x99, 0xfa, 0x8b, 0x02, 0xba, 0xb2, 0xb1, 0x6f,
	0x13, 0x12, 0x33, 0xeb, 0x8c, 0x3a, 0xda, 0xae, 0xe8, 0x2c, 0xbe, 0xe0, 0xb0, 0xf3, 0x45, 0x5e,
	0x85, 0x60, 0x1e, 0x53, 0x33, 0xe5, 0xb0, 0xe3, 0x55, 0x81, 0x32, 0xa4, 0x86, 0x2e, 0x8b, 0x4c,
	0xcf, 0x07, 0x0d, 0x79, 0x13, 0x78, 0xb1, 0x18, 0xd4, 0x13, 0x50, 0x8d, 0x77, 0xd4, 0x8f, 0x41,
	0x3b, 0xf1, 0x59, 0x94, 0xc4, 0x8c, 0x60, 0xad, 0x27, 0xce, 0x9d, 0x9e, 0x3f, 0x1b, 0x25, 0x98,
	0x71, 0xd8, 0x15, 0x3b, 0x28, 0x11, 0x03, 0xad, 0x58, 0xf3, 0xb3, 0x97, 0xaf, 0xfa, 0xad, 0xc5,
	0xab, 0x7e, 0xeb, 0xef, 0x8b, 0x7e, 0xeb, 0xe7, 0xcb, 0x7e, 0xeb, 0xd7, 0xcb, 0xbe, 0xb2, 0xb8,
	0xec, 0xb7, 0xfe, 0xbc, 0xec, 0xb7, 0x4e, 0xee, 0xbd, 0xc5, 0x3f, 0x83, 0x2c, 0xdd, 0x59, 0x17,
	0xff, 0x10, 0xef, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x7a, 0x4b, 0x6c, 0x4b, 0x08, 0x00,
	0x00,
}

func (m *DeviceConfiguration) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DeviceID.ProtoSize()
	n += 1 + l + sovDeviceconfiguration(uint64(l))
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDeviceconfiguration(uint64(l))
	}
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovDeviceconfiguration(uint64(l))
		}
	}
	if m.Compression != 0 {
		n += 1 + sovDeviceconfiguration(uint64(m.Compression))
	}
	l = len(m.CertName)
	if l > 0 {
		n += 1 + l + sovDeviceconfiguration(uint64(l))
	}
	if m.Introducer {
		n += 2
	}
	if m.SkipIntroductionRemovals {
		n += 2
	}
	l = m.IntroducedBy.ProtoSize()
	n += 1 + l + sovDeviceconfiguration(uint64(l))
	if m.Paused {
		n += 2
	}
	if len(m.AllowedNetworks) > 0 {
		for _, s := range m.AllowedNetworks {
			l = len(s)
			n += 1 + l + sovDeviceconfiguration(uint64(l))
		}
	}
	if m.AutoAcceptFolders {
		n += 2
	}
	if m.MaxSendKbps != 0 {
		n += 1 + sovDeviceconfiguration(uint64(m.MaxSendKbps))
	}
	if m.MaxRecvKbps != 0 {
		n += 1 + sovDeviceconfiguration(uint64(m.MaxRecvKbps))
	}
	if len(m.IgnoredFolders) > 0 {
		for _, e := range m.IgnoredFolders {
			l = e.ProtoSize()
			n += 1 + l + sovDeviceconfiguration(uint64(l))
		}
	}
	if len(m.PendingFolders) > 0 {
		for _, e := range m.PendingFolders {
			l = e.ProtoSize()
			n += 1 + l + sovDeviceconfiguration(uint64(l))
		}
	}
	if m.MaxRequestKiB != 0 {
		n += 2 + sovDeviceconfiguration(uint64(m.MaxRequestKiB))
	}
	if m.Untrusted {
		n += 3
	}
	return n
}

func sovDeviceconfiguration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeviceconfiguration(x uint64) (n int) {
	return sovDeviceconfiguration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}