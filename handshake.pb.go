// Code generated by protoc-gen-go.
// source: handshake.proto
// DO NOT EDIT!

package dicoprotos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Handshake struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	RunsTasks        *bool   `protobuf:"varint,2,opt,name=runs_tasks,def=0" json:"runs_tasks,omitempty"`
	ManagesTasks     *bool   `protobuf:"varint,3,opt,name=manages_tasks,def=0" json:"manages_tasks,omitempty"`
	RecievesStats    *bool   `protobuf:"varint,4,opt,name=recieves_stats,def=0" json:"recieves_stats,omitempty"`
	RecievesResults  *bool   `protobuf:"varint,5,opt,name=recieves_results,def=0" json:"recieves_results,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Handshake) Reset()                    { *m = Handshake{} }
func (m *Handshake) String() string            { return proto.CompactTextString(m) }
func (*Handshake) ProtoMessage()               {}
func (*Handshake) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

const Default_Handshake_RunsTasks bool = false
const Default_Handshake_ManagesTasks bool = false
const Default_Handshake_RecievesStats bool = false
const Default_Handshake_RecievesResults bool = false

func (m *Handshake) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Handshake) GetRunsTasks() bool {
	if m != nil && m.RunsTasks != nil {
		return *m.RunsTasks
	}
	return Default_Handshake_RunsTasks
}

func (m *Handshake) GetManagesTasks() bool {
	if m != nil && m.ManagesTasks != nil {
		return *m.ManagesTasks
	}
	return Default_Handshake_ManagesTasks
}

func (m *Handshake) GetRecievesStats() bool {
	if m != nil && m.RecievesStats != nil {
		return *m.RecievesStats
	}
	return Default_Handshake_RecievesStats
}

func (m *Handshake) GetRecievesResults() bool {
	if m != nil && m.RecievesResults != nil {
		return *m.RecievesResults
	}
	return Default_Handshake_RecievesResults
}

func init() {
	proto.RegisterType((*Handshake)(nil), "dicoprotos.Handshake")
}

var fileDescriptor1 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0xce, 0xc1, 0x0a, 0x82, 0x40,
	0x10, 0xc6, 0x71, 0x2c, 0x85, 0x1c, 0xca, 0x62, 0x4f, 0x16, 0x45, 0xd1, 0xa9, 0x43, 0xf8, 0x00,
	0x1d, 0x3d, 0x75, 0xec, 0x0d, 0x62, 0x58, 0xa7, 0x14, 0xdd, 0xdd, 0xd8, 0x59, 0x8b, 0x5e, 0xa3,
	0x27, 0x4e, 0x0b, 0xc5, 0xdb, 0x30, 0xbf, 0xef, 0xf0, 0x87, 0x79, 0x8e, 0x3a, 0xe3, 0x1c, 0x4b,
	0x4a, 0x1e, 0xd6, 0x38, 0x23, 0x20, 0x2b, 0xa4, 0xf9, 0x9d, 0xbc, 0xff, 0x78, 0x10, 0x9e, 0x3b,
	0x17, 0x53, 0xf0, 0x35, 0x2a, 0x8a, 0xbd, 0x9d, 0x77, 0x08, 0xc5, 0x12, 0xc0, 0xd6, 0x9a, 0xaf,
	0x0e, 0xb9, 0xe4, 0x78, 0xd4, 0xfc, 0x26, 0xa7, 0xe0, 0x86, 0x15, 0x93, 0x58, 0xc3, 0x4c, 0xa1,
	0xc6, 0x3b, 0x75, 0x3a, 0x1e, 0xea, 0x06, 0x22, 0x4b, 0xb2, 0xa0, 0x67, 0xc3, 0xec, 0xd0, 0x71,
	0xec, 0x0f, 0x79, 0x0b, 0x8b, 0x9e, 0x2d, 0x71, 0x5d, 0x35, 0x83, 0x60, 0x30, 0x48, 0x8f, 0xb0,
	0x92, 0x46, 0x25, 0x32, 0x27, 0x6b, 0xdf, 0x2f, 0x63, 0x55, 0xd2, 0x16, 0xff, 0xeb, 0x39, 0x8d,
	0xfa, 0xde, 0x4b, 0xfb, 0xf8, 0x06, 0x00, 0x00, 0xff, 0xff, 0x76, 0x81, 0xc1, 0xd9, 0xe0, 0x00,
	0x00, 0x00,
}
