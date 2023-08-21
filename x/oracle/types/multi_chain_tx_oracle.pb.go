// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: re/oracle/multi_chain_tx_oracle.proto

package types

import (
	fmt "fmt"
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

type MultiChainTxOracle struct {
	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OracleId  string `protobuf:"bytes,2,opt,name=oracleId,proto3" json:"oracleId,omitempty"`
	Address   string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Timestamp string `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *MultiChainTxOracle) Reset()         { *m = MultiChainTxOracle{} }
func (m *MultiChainTxOracle) String() string { return proto.CompactTextString(m) }
func (*MultiChainTxOracle) ProtoMessage()    {}
func (*MultiChainTxOracle) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2c534195cf0d16e, []int{0}
}
func (m *MultiChainTxOracle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MultiChainTxOracle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MultiChainTxOracle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MultiChainTxOracle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiChainTxOracle.Merge(m, src)
}
func (m *MultiChainTxOracle) XXX_Size() int {
	return m.Size()
}
func (m *MultiChainTxOracle) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiChainTxOracle.DiscardUnknown(m)
}

var xxx_messageInfo_MultiChainTxOracle proto.InternalMessageInfo

func (m *MultiChainTxOracle) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MultiChainTxOracle) GetOracleId() string {
	if m != nil {
		return m.OracleId
	}
	return ""
}

func (m *MultiChainTxOracle) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MultiChainTxOracle) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func init() {
	proto.RegisterType((*MultiChainTxOracle)(nil), "jim380.re.oracle.MultiChainTxOracle")
}

func init() {
	proto.RegisterFile("re/oracle/multi_chain_tx_oracle.proto", fileDescriptor_a2c534195cf0d16e)
}

var fileDescriptor_a2c534195cf0d16e = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x4a, 0xd5, 0xcf,
	0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0xcf, 0x2d, 0xcd, 0x29, 0xc9, 0x8c, 0x4f, 0xce, 0x48, 0xcc,
	0xcc, 0x8b, 0x2f, 0xa9, 0x88, 0x87, 0x88, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0x64,
	0x65, 0xe6, 0x1a, 0x5b, 0x18, 0xe8, 0x15, 0xa5, 0xea, 0x41, 0xc4, 0x95, 0x2a, 0xb8, 0x84, 0x7c,
	0x41, 0x1a, 0x9c, 0x41, 0xea, 0x43, 0x2a, 0xfc, 0xc1, 0xa2, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x52, 0x5c, 0x1c, 0x10, 0xf5,
	0x9e, 0x29, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x90, 0x04, 0x17, 0x7b, 0x62,
	0x4a, 0x4a, 0x51, 0x6a, 0x71, 0xb1, 0x04, 0x33, 0x58, 0x0a, 0xc6, 0x15, 0x92, 0xe1, 0xe2, 0x2c,
	0xc9, 0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0x90, 0x60, 0x01, 0xcb, 0x21, 0x04, 0x9c, 0x6c,
	0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5,
	0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x39, 0x3d, 0xb3, 0x24, 0xa3,
	0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xe2, 0x60, 0xfd, 0xa0, 0x54, 0xfd, 0x0a, 0x98, 0x07,
	0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x3e, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xb9, 0x28, 0xf0, 0xfe, 0xfa, 0x00, 0x00, 0x00,
}

func (m *MultiChainTxOracle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiChainTxOracle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MultiChainTxOracle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Timestamp) > 0 {
		i -= len(m.Timestamp)
		copy(dAtA[i:], m.Timestamp)
		i = encodeVarintMultiChainTxOracle(dAtA, i, uint64(len(m.Timestamp)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMultiChainTxOracle(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OracleId) > 0 {
		i -= len(m.OracleId)
		copy(dAtA[i:], m.OracleId)
		i = encodeVarintMultiChainTxOracle(dAtA, i, uint64(len(m.OracleId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintMultiChainTxOracle(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMultiChainTxOracle(dAtA []byte, offset int, v uint64) int {
	offset -= sovMultiChainTxOracle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MultiChainTxOracle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMultiChainTxOracle(uint64(m.Id))
	}
	l = len(m.OracleId)
	if l > 0 {
		n += 1 + l + sovMultiChainTxOracle(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMultiChainTxOracle(uint64(l))
	}
	l = len(m.Timestamp)
	if l > 0 {
		n += 1 + l + sovMultiChainTxOracle(uint64(l))
	}
	return n
}

func sovMultiChainTxOracle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMultiChainTxOracle(x uint64) (n int) {
	return sovMultiChainTxOracle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MultiChainTxOracle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMultiChainTxOracle
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
			return fmt.Errorf("proto: MultiChainTxOracle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiChainTxOracle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultiChainTxOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultiChainTxOracle
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
				return ErrInvalidLengthMultiChainTxOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultiChainTxOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultiChainTxOracle
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
				return ErrInvalidLengthMultiChainTxOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultiChainTxOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultiChainTxOracle
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
				return ErrInvalidLengthMultiChainTxOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultiChainTxOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timestamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMultiChainTxOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMultiChainTxOracle
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
func skipMultiChainTxOracle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMultiChainTxOracle
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
					return 0, ErrIntOverflowMultiChainTxOracle
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
					return 0, ErrIntOverflowMultiChainTxOracle
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
				return 0, ErrInvalidLengthMultiChainTxOracle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMultiChainTxOracle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMultiChainTxOracle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMultiChainTxOracle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMultiChainTxOracle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMultiChainTxOracle = fmt.Errorf("proto: unexpected end of group")
)
