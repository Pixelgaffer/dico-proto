// Code generated by protoc-gen-go.
// source: do-task.proto
// DO NOT EDIT!

/*
Package dicoprotos is a generated protocol buffer package.

It is generated from these files:
	do-task.proto
	handshake.proto
	submit-task.proto
	task-result.proto
	task-status.proto

It has these top-level messages:
	DoTask
	Handshake
	SubmitTask
	TaskResult
	TaskStatus
*/
package dicoprotos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type DoTask struct {
	Id               *int64  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Options          *string `protobuf:"bytes,2,req,name=options" json:"options,omitempty"`
	JobType          *string `protobuf:"bytes,3,req,name=job_type" json:"job_type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DoTask) Reset()                    { *m = DoTask{} }
func (m *DoTask) String() string            { return proto.CompactTextString(m) }
func (*DoTask) ProtoMessage()               {}
func (*DoTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DoTask) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *DoTask) GetOptions() string {
	if m != nil && m.Options != nil {
		return *m.Options
	}
	return ""
}

func (m *DoTask) GetJobType() string {
	if m != nil && m.JobType != nil {
		return *m.JobType
	}
	return ""
}

func init() {
	proto.RegisterType((*DoTask)(nil), "dicoprotos.DoTask")
}

var fileDescriptor0 = []byte{
	// 100 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xc9, 0xd7, 0x2d,
	0x49, 0x2c, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0xc9, 0x4c, 0xce, 0x07,
	0x33, 0x8b, 0x95, 0xcc, 0xb9, 0xd8, 0x5c, 0xf2, 0x43, 0x80, 0x72, 0x42, 0x5c, 0x5c, 0x4c, 0x99,
	0x29, 0x12, 0x8c, 0x0a, 0x4c, 0x1a, 0xcc, 0x42, 0xfc, 0x5c, 0xec, 0xf9, 0x05, 0x25, 0x99, 0xf9,
	0x79, 0xc5, 0x12, 0x4c, 0x40, 0x01, 0x4e, 0x21, 0x01, 0x2e, 0x8e, 0xac, 0xfc, 0xa4, 0xf8, 0x92,
	0xca, 0x82, 0x54, 0x09, 0x66, 0x90, 0x08, 0x20, 0x00, 0x00, 0xff, 0xff, 0x91, 0xec, 0xe2, 0x87,
	0x54, 0x00, 0x00, 0x00,
}
