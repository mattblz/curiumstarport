// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: synchronizer/source.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Source struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Url     string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *Source) Reset()         { *m = Source{} }
func (m *Source) String() string { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()    {}
func (*Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_9198c38abe4fb49d, []int{0}
}
func (m *Source) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Source.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Source.Merge(m, src)
}
func (m *Source) XXX_Size() int {
	return m.Size()
}
func (m *Source) XXX_DiscardUnknown() {
	xxx_messageInfo_Source.DiscardUnknown(m)
}

var xxx_messageInfo_Source proto.InternalMessageInfo

func (m *Source) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Source) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Source) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*Source)(nil), "bluzelle.curium.synchronizer.Source")
}

func init() { proto.RegisterFile("synchronizer/source.proto", fileDescriptor_9198c38abe4fb49d) }

var fileDescriptor_9198c38abe4fb49d = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0xae, 0xcc, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcb, 0xac, 0x4a, 0x2d, 0xd2, 0x2f, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x49, 0xca, 0x29, 0xad, 0x4a, 0xcd, 0xc9, 0x49, 0xd5,
	0x4b, 0x2e, 0x2d, 0xca, 0x2c, 0xcd, 0xd5, 0x43, 0x56, 0x2a, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f,
	0x56, 0xa8, 0x0f, 0x62, 0x41, 0xf4, 0x28, 0x79, 0x70, 0xb1, 0x05, 0x83, 0xcd, 0x10, 0x92, 0xe0,
	0x62, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82,
	0x71, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0xc0, 0xc2, 0x60, 0xb6, 0x90,
	0x00, 0x17, 0x73, 0x69, 0x51, 0x8e, 0x04, 0x33, 0x58, 0x08, 0xc4, 0x74, 0xf2, 0x3c, 0xf1, 0x48,
	0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0,
	0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xfd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd,
	0xe4, 0xfc, 0x5c, 0x7d, 0x98, 0x13, 0xf5, 0x21, 0x4e, 0xd4, 0xaf, 0xd0, 0x47, 0xf1, 0x4f, 0x49,
	0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x6d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16,
	0x2d, 0xab, 0xf6, 0xec, 0x00, 0x00, 0x00,
}

func (m *Source) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Source) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Source) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintSource(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSource(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSource(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSource(dAtA []byte, offset int, v uint64) int {
	offset -= sovSource(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Source) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSource(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSource(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovSource(uint64(l))
	}
	return n
}

func sovSource(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSource(x uint64) (n int) {
	return sovSource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Source) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSource
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
			return fmt.Errorf("proto: Source: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Source: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSource
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
				return ErrInvalidLengthSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSource
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
				return ErrInvalidLengthSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSource
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
				return ErrInvalidLengthSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSource
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
func skipSource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSource
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
					return 0, ErrIntOverflowSource
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
					return 0, ErrIntOverflowSource
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
				return 0, ErrInvalidLengthSource
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSource
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSource
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSource        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSource          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSource = fmt.Errorf("proto: unexpected end of group")
)
