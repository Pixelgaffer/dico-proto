// Code generated by protoc-gen-go.
// source: worker-status.proto
// DO NOT EDIT!

package dicoprotos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WorkerStatus_WorkerStatusUpdate int32

const (
	WorkerStatus_CONNECTED    WorkerStatus_WorkerStatusUpdate = 0
	WorkerStatus_DISCONNECTED WorkerStatus_WorkerStatusUpdate = 1
)

var WorkerStatus_WorkerStatusUpdate_name = map[int32]string{
	0: "CONNECTED",
	1: "DISCONNECTED",
}
var WorkerStatus_WorkerStatusUpdate_value = map[string]int32{
	"CONNECTED":    0,
	"DISCONNECTED": 1,
}

func (x WorkerStatus_WorkerStatusUpdate) Enum() *WorkerStatus_WorkerStatusUpdate {
	p := new(WorkerStatus_WorkerStatusUpdate)
	*p = x
	return p
}
func (x WorkerStatus_WorkerStatusUpdate) String() string {
	return proto.EnumName(WorkerStatus_WorkerStatusUpdate_name, int32(x))
}
func (x *WorkerStatus_WorkerStatusUpdate) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WorkerStatus_WorkerStatusUpdate_value, data, "WorkerStatus_WorkerStatusUpdate")
	if err != nil {
		return err
	}
	*x = WorkerStatus_WorkerStatusUpdate(value)
	return nil
}
func (WorkerStatus_WorkerStatusUpdate) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor8, []int{0, 0}
}

type WorkerStatus struct {
	Id               *string                          `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Update           *WorkerStatus_WorkerStatusUpdate `protobuf:"varint,2,req,name=update,enum=dicoprotos.WorkerStatus_WorkerStatusUpdate" json:"update,omitempty"`
	XXX_unrecognized []byte                           `json:"-"`
}

func (m *WorkerStatus) Reset()                    { *m = WorkerStatus{} }
func (m *WorkerStatus) String() string            { return proto.CompactTextString(m) }
func (*WorkerStatus) ProtoMessage()               {}
func (*WorkerStatus) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *WorkerStatus) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *WorkerStatus) GetUpdate() WorkerStatus_WorkerStatusUpdate {
	if m != nil && m.Update != nil {
		return *m.Update
	}
	return WorkerStatus_CONNECTED
}

func init() {
	proto.RegisterType((*WorkerStatus)(nil), "dicoprotos.WorkerStatus")
	proto.RegisterEnum("dicoprotos.WorkerStatus_WorkerStatusUpdate", WorkerStatus_WorkerStatusUpdate_name, WorkerStatus_WorkerStatusUpdate_value)
}

var fileDescriptor8 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xcf, 0x2f, 0xca,
	0x4e, 0x2d, 0xd2, 0x2d, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x4a, 0xc9, 0x4c, 0xce, 0x07, 0x33, 0x8b, 0x95, 0xfa, 0x18, 0xb9, 0x78, 0xc2, 0xc1, 0x6a,
	0x82, 0xc1, 0x4a, 0x84, 0xb8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x98, 0x34, 0x38, 0x85,
	0xac, 0xb9, 0xd8, 0x4a, 0x0b, 0x52, 0x12, 0x4b, 0x52, 0x25, 0x98, 0x80, 0x7c, 0x3e, 0x23, 0x6d,
	0x3d, 0x84, 0x4e, 0x3d, 0x64, 0x5d, 0x28, 0x9c, 0x50, 0xb0, 0x16, 0x25, 0x53, 0x2e, 0x21, 0x4c,
	0x51, 0x21, 0x5e, 0x2e, 0x4e, 0x67, 0x7f, 0x3f, 0x3f, 0x57, 0xe7, 0x10, 0x57, 0x17, 0x01, 0x06,
	0x21, 0x01, 0x2e, 0x1e, 0x17, 0xcf, 0x60, 0x84, 0x08, 0xa3, 0x93, 0x3e, 0x97, 0x54, 0x72, 0x7e,
	0xae, 0x5e, 0x72, 0x46, 0x6a, 0x51, 0x51, 0x25, 0xd0, 0xf5, 0xb9, 0x60, 0x3b, 0x21, 0x2e, 0x2f,
	0x76, 0x12, 0x44, 0x36, 0x32, 0x00, 0x24, 0x06, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x2a, 0x8a,
	0xa0, 0xe3, 0x00, 0x00, 0x00,
}
