// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: synchronizer/config.proto

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

type Config struct {
	SyncAddress string `protobuf:"bytes,1,opt,name=syncAddress,proto3" json:"syncAddress,omitempty"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_b465213be122c849, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Config.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return m.Size()
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetSyncAddress() string {
	if m != nil {
		return m.SyncAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "bluzelle.curium.synchronizer.Config")
}

func init() { proto.RegisterFile("synchronizer/config.proto", fileDescriptor_b465213be122c849) }

var fileDescriptor_b465213be122c849 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0xae, 0xcc, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcb, 0xac, 0x4a, 0x2d, 0xd2, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x49, 0xca, 0x29, 0xad, 0x4a, 0xcd, 0xc9, 0x49, 0xd5,
	0x4b, 0x2e, 0x2d, 0xca, 0x2c, 0xcd, 0xd5, 0x43, 0x56, 0x2a, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f,
	0x56, 0xa8, 0x0f, 0x62, 0x41, 0xf4, 0x28, 0x69, 0x71, 0xb1, 0x39, 0x83, 0xcd, 0x10, 0x52, 0xe0,
	0xe2, 0x06, 0xa9, 0x77, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x42, 0x16, 0x72, 0xf2, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28,
	0xfd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x98, 0x23, 0xf4, 0x21,
	0x8e, 0xd0, 0xaf, 0xd0, 0x47, 0x71, 0x71, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x76,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x0b, 0x61, 0xfb, 0xce, 0x00, 0x00, 0x00,
}

func (m *Config) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Config) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SyncAddress) > 0 {
		i -= len(m.SyncAddress)
		copy(dAtA[i:], m.SyncAddress)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.SyncAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Config) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SyncAddress)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Config) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: Config: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Config: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SyncAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SyncAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)
