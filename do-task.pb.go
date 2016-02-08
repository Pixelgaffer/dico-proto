// Code generated by protoc-gen-go.
// source: do-task.proto
// DO NOT EDIT!

/*
Package dicoprotos is a generated protocol buffer package.

It is generated from these files:
	do-task.proto
	handshake.proto
	job-class-created.proto
	self-describing-message.proto
	submit-code.proto
	submit-task.proto
	task-result.proto
	task-status.proto
	worker-status.proto

It has these top-level messages:
	DoTask
	Handshake
	JobClassCreated
	SelfDescribingMessage
	SubmitCode
	SubmitTask
	TaskResult
	TaskStatus
	WorkerStatus
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
	Code             *string `protobuf:"bytes,2,req,name=code" json:"code,omitempty"`
	Options          *string `protobuf:"bytes,3,req,name=options" json:"options,omitempty"`
	JobType          *string `protobuf:"bytes,4,req,name=job_type" json:"job_type,omitempty"`
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

func (m *DoTask) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
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
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xc9, 0xd7, 0x2d,
	0x49, 0x2c, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0xc9, 0x4c, 0xce, 0x07,
	0x33, 0x8b, 0x95, 0x5c, 0xb9, 0xd8, 0x5c, 0xf2, 0x43, 0x80, 0x72, 0x42, 0x5c, 0x5c, 0x4c, 0x99,
	0x29, 0x12, 0x8c, 0x0a, 0x4c, 0x1a, 0xcc, 0x42, 0x3c, 0x5c, 0x2c, 0xc9, 0xf9, 0x29, 0xa9, 0x12,
	0x4c, 0x40, 0x1e, 0xa7, 0x10, 0x3f, 0x17, 0x7b, 0x7e, 0x41, 0x49, 0x66, 0x7e, 0x5e, 0xb1, 0x04,
	0x33, 0x58, 0x40, 0x80, 0x8b, 0x23, 0x2b, 0x3f, 0x29, 0xbe, 0xa4, 0xb2, 0x20, 0x55, 0x82, 0x05,
	0x24, 0x02, 0x08, 0x00, 0x00, 0xff, 0xff, 0x07, 0x90, 0x8d, 0x85, 0x62, 0x00, 0x00, 0x00,
}
