// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: specy/specy/task.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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

type Task struct {
	Owner             string     `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name              string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Hash              string     `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	ConnectionId      string     `protobuf:"bytes,4,opt,name=connectionId,proto3" json:"connectionId,omitempty"`
	Msg               *types.Any `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
	RuleFiles         string     `protobuf:"bytes,6,opt,name=ruleFiles,proto3" json:"ruleFiles,omitempty"`
	TaskType          uint64     `protobuf:"varint,7,opt,name=taskType,proto3" json:"taskType,omitempty"`
	ScheduleType      *Condition `protobuf:"bytes,8,opt,name=scheduleType,proto3" json:"scheduleType,omitempty"`
	UpdateTime        uint64     `protobuf:"varint,9,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	UpdateBlockHeight uint64     `protobuf:"varint,10,opt,name=updateBlockHeight,proto3" json:"updateBlockHeight,omitempty"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6cd4bd76707bef, []int{0}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Task.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return m.Size()
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Task) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *Task) GetMsg() *types.Any {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *Task) GetRuleFiles() string {
	if m != nil {
		return m.RuleFiles
	}
	return ""
}

func (m *Task) GetTaskType() uint64 {
	if m != nil {
		return m.TaskType
	}
	return 0
}

func (m *Task) GetScheduleType() *Condition {
	if m != nil {
		return m.ScheduleType
	}
	return nil
}

func (m *Task) GetUpdateTime() uint64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *Task) GetUpdateBlockHeight() uint64 {
	if m != nil {
		return m.UpdateBlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Task)(nil), "specynetwork.specy.specy.Task")
}

func init() { proto.RegisterFile("specy/specy/task.proto", fileDescriptor_8e6cd4bd76707bef) }

var fileDescriptor_8e6cd4bd76707bef = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0xa5, 0x50, 0x10, 0x56, 0x2e, 0x6e, 0x88, 0x59, 0xd1, 0x34, 0x04, 0x13, 0xc3, 0x01, 0x97,
	0x44, 0xbf, 0x40, 0x4c, 0x50, 0xaf, 0x0d, 0x27, 0x6f, 0x65, 0x3b, 0xb6, 0x4d, 0xdb, 0xdd, 0xa6,
	0xbb, 0x0d, 0xf6, 0x2f, 0xfc, 0x2c, 0x8f, 0x1c, 0x3d, 0x2a, 0xfc, 0x88, 0xe9, 0x2e, 0x28, 0xc4,
	0x78, 0x99, 0xcc, 0x7b, 0x6f, 0x66, 0x67, 0xf2, 0x66, 0xd1, 0xa9, 0xcc, 0x80, 0x95, 0x13, 0x13,
	0x95, 0x27, 0x63, 0x9a, 0xe5, 0x42, 0x09, 0x4c, 0x34, 0xc3, 0x41, 0x2d, 0x45, 0x1e, 0x53, 0x0d,
	0x4c, 0xec, 0x9f, 0xef, 0x77, 0x30, 0xc1, 0xfd, 0x48, 0x45, 0x82, 0x9b, 0xb6, 0xfe, 0x59, 0x20,
	0x44, 0x90, 0xc0, 0x44, 0xa3, 0x45, 0xf1, 0x32, 0xf1, 0x78, 0x69, 0xa4, 0xe1, 0x57, 0x1d, 0xd9,
	0x73, 0x4f, 0xc6, 0xb8, 0x87, 0x9a, 0x62, 0xc9, 0x21, 0x27, 0xd6, 0xc0, 0x1a, 0x75, 0x5c, 0x03,
	0x30, 0x46, 0x36, 0xf7, 0x52, 0x20, 0x75, 0x4d, 0xea, 0xbc, 0xe2, 0x42, 0x4f, 0x86, 0xa4, 0x61,
	0xb8, 0x2a, 0xc7, 0x43, 0xd4, 0x65, 0x82, 0x73, 0x60, 0xd5, 0xd4, 0x27, 0x9f, 0xd8, 0x5a, 0x3b,
	0xe0, 0xf0, 0x15, 0x6a, 0xa4, 0x32, 0x20, 0xcd, 0x81, 0x35, 0x3a, 0xbe, 0xe9, 0x51, 0xb3, 0x13,
	0xdd, 0xed, 0x44, 0xef, 0x78, 0xe9, 0x56, 0x05, 0xf8, 0x02, 0x75, 0xf2, 0x22, 0x81, 0x59, 0x94,
	0x80, 0x24, 0x2d, 0xfd, 0xd0, 0x2f, 0x81, 0xfb, 0xa8, 0x5d, 0x19, 0x32, 0x2f, 0x33, 0x20, 0x47,
	0x03, 0x6b, 0x64, 0xbb, 0x3f, 0x18, 0x3f, 0xa0, 0xae, 0x64, 0x21, 0xf8, 0x45, 0x02, 0x5a, 0x6f,
	0xeb, 0x51, 0x97, 0xf4, 0x3f, 0xd7, 0xe8, 0xfd, 0xce, 0x28, 0xf7, 0xa0, 0x11, 0x3b, 0x08, 0x15,
	0x99, 0xef, 0x29, 0x98, 0x47, 0x29, 0x90, 0x8e, 0x1e, 0xb3, 0xc7, 0xe0, 0x31, 0x3a, 0x31, 0x68,
	0x9a, 0x08, 0x16, 0x3f, 0x42, 0x14, 0x84, 0x8a, 0x20, 0x5d, 0xf6, 0x57, 0x98, 0xce, 0xde, 0xd7,
	0x8e, 0xb5, 0x5a, 0x3b, 0xd6, 0xe7, 0xda, 0xb1, 0xde, 0x36, 0x4e, 0x6d, 0xb5, 0x71, 0x6a, 0x1f,
	0x1b, 0xa7, 0xf6, 0x3c, 0x0e, 0x22, 0x15, 0x16, 0x0b, 0xca, 0x44, 0x6a, 0x4e, 0x77, 0xbd, 0xdd,
	0x72, 0x7b, 0xc8, 0xd7, 0xdd, 0x17, 0x28, 0x33, 0x90, 0x8b, 0x96, 0xf6, 0xea, 0xf6, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x7a, 0xaa, 0xad, 0x2f, 0x1e, 0x02, 0x00, 0x00,
}

func (m *Task) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Task) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Task) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdateBlockHeight != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.UpdateBlockHeight))
		i--
		dAtA[i] = 0x50
	}
	if m.UpdateTime != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.UpdateTime))
		i--
		dAtA[i] = 0x48
	}
	if m.ScheduleType != nil {
		{
			size, err := m.ScheduleType.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTask(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.TaskType != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.TaskType))
		i--
		dAtA[i] = 0x38
	}
	if len(m.RuleFiles) > 0 {
		i -= len(m.RuleFiles)
		copy(dAtA[i:], m.RuleFiles)
		i = encodeVarintTask(dAtA, i, uint64(len(m.RuleFiles)))
		i--
		dAtA[i] = 0x32
	}
	if m.Msg != nil {
		{
			size, err := m.Msg.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTask(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintTask(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTask(dAtA []byte, offset int, v uint64) int {
	offset -= sovTask(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Task) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.Msg != nil {
		l = m.Msg.Size()
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.RuleFiles)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.TaskType != 0 {
		n += 1 + sovTask(uint64(m.TaskType))
	}
	if m.ScheduleType != nil {
		l = m.ScheduleType.Size()
		n += 1 + l + sovTask(uint64(l))
	}
	if m.UpdateTime != 0 {
		n += 1 + sovTask(uint64(m.UpdateTime))
	}
	if m.UpdateBlockHeight != 0 {
		n += 1 + sovTask(uint64(m.UpdateBlockHeight))
	}
	return n
}

func sovTask(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTask(x uint64) (n int) {
	return sovTask(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Task) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Task: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Task: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Msg == nil {
				m.Msg = &types.Any{}
			}
			if err := m.Msg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuleFiles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RuleFiles = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskType", wireType)
			}
			m.TaskType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TaskType |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduleType", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ScheduleType == nil {
				m.ScheduleType = &Condition{}
			}
			if err := m.ScheduleType.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateTime", wireType)
			}
			m.UpdateTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateBlockHeight", wireType)
			}
			m.UpdateBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTask
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTask(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTask
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTask
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTask
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTask
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTask
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTask
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTask        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTask          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTask = fmt.Errorf("proto: unexpected end of group")
)
