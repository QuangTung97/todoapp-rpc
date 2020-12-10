// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: todoapp.proto

package todoapp

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TodoGetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoGetRequest) Reset()         { *m = TodoGetRequest{} }
func (m *TodoGetRequest) String() string { return proto.CompactTextString(m) }
func (*TodoGetRequest) ProtoMessage()    {}
func (*TodoGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_576f159e2510ba3a, []int{0}
}
func (m *TodoGetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TodoGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TodoGetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TodoGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoGetRequest.Merge(m, src)
}
func (m *TodoGetRequest) XXX_Size() int {
	return m.Size()
}
func (m *TodoGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TodoGetRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TodoGetRequest)(nil), "todoapp.v1.TodoGetRequest")
}

func init() { proto.RegisterFile("todoapp.proto", fileDescriptor_576f159e2510ba3a) }

var fileDescriptor_576f159e2510ba3a = []byte{
	// 105 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0xc9, 0x4f, 0xc9,
	0x4f, 0x2c, 0x28, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0x71, 0xcb, 0x0c, 0x95,
	0x04, 0xb8, 0xf8, 0x42, 0xf2, 0x53, 0xf2, 0xdd, 0x53, 0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x9c, 0x4c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6,
	0x19, 0x8f, 0xe5, 0x18, 0xa2, 0x94, 0xa0, 0xea, 0x75, 0x8b, 0x0a, 0x92, 0xf5, 0x41, 0x18, 0xca,
	0xd7, 0x2f, 0x33, 0xb4, 0x86, 0x32, 0x93, 0xd8, 0xc0, 0x46, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xa9, 0xd5, 0x6f, 0xcf, 0x6b, 0x00, 0x00, 0x00,
}

func (m *TodoGetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TodoGetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TodoGetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintTodoapp(dAtA []byte, offset int, v uint64) int {
	offset -= sovTodoapp(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TodoGetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTodoapp(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTodoapp(x uint64) (n int) {
	return sovTodoapp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TodoGetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTodoapp
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
			return fmt.Errorf("proto: TodoGetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TodoGetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTodoapp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTodoapp
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTodoapp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTodoapp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTodoapp
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
					return 0, ErrIntOverflowTodoapp
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
					return 0, ErrIntOverflowTodoapp
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
				return 0, ErrInvalidLengthTodoapp
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTodoapp
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTodoapp
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTodoapp        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTodoapp          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTodoapp = fmt.Errorf("proto: unexpected end of group")
)
