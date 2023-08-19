// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: re/fix/orders_cancel_request.proto

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

type OrdersCancelRequest struct {
	// A string field that specifies the FIX session ID for the message
	SessionID string `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	// A Header field that contains standard header information for the message,
	// such as the message type, sender and receiver identification, and sequence
	// number
	Header *Header `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	// A string field that contains the original client order ID for the order
	// that was cancelled
	OrigClOrdID string `protobuf:"bytes,3,opt,name=origClOrdID,proto3" json:"origClOrdID,omitempty"`
	// A string field that contains the client order ID for the order that was
	// cancelled or modified
	ClOrdID string `protobuf:"bytes,4,opt,name=clOrdID,proto3" json:"clOrdID,omitempty"`
	// A Trailer field that contains standard trailer information for the message,
	// such as the message checksum
	Trailer *Trailer `protobuf:"bytes,5,opt,name=trailer,proto3" json:"trailer,omitempty"`
}

func (m *OrdersCancelRequest) Reset()         { *m = OrdersCancelRequest{} }
func (m *OrdersCancelRequest) String() string { return proto.CompactTextString(m) }
func (*OrdersCancelRequest) ProtoMessage()    {}
func (*OrdersCancelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc47815a4d600d55, []int{0}
}
func (m *OrdersCancelRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrdersCancelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrdersCancelRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrdersCancelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrdersCancelRequest.Merge(m, src)
}
func (m *OrdersCancelRequest) XXX_Size() int {
	return m.Size()
}
func (m *OrdersCancelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrdersCancelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrdersCancelRequest proto.InternalMessageInfo

func (m *OrdersCancelRequest) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *OrdersCancelRequest) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *OrdersCancelRequest) GetOrigClOrdID() string {
	if m != nil {
		return m.OrigClOrdID
	}
	return ""
}

func (m *OrdersCancelRequest) GetClOrdID() string {
	if m != nil {
		return m.ClOrdID
	}
	return ""
}

func (m *OrdersCancelRequest) GetTrailer() *Trailer {
	if m != nil {
		return m.Trailer
	}
	return nil
}

func init() {
	proto.RegisterType((*OrdersCancelRequest)(nil), "re.fix.OrdersCancelRequest")
}

func init() {
	proto.RegisterFile("re/fix/orders_cancel_request.proto", fileDescriptor_fc47815a4d600d55)
}

var fileDescriptor_fc47815a4d600d55 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0x73, 0xfe, 0x49, 0xe9, 0x15, 0x14, 0xce, 0x25, 0x88, 0x1e, 0xa1, 0x83, 0xd4, 0x25,
	0x11, 0xbb, 0x88, 0xa3, 0xed, 0x60, 0xa7, 0x42, 0x70, 0x72, 0x09, 0xd7, 0xe4, 0x6d, 0x7a, 0x92,
	0xe4, 0xea, 0x7b, 0x57, 0x88, 0xdf, 0xc2, 0xcf, 0xe4, 0xe4, 0xd8, 0xd1, 0x51, 0x92, 0x2f, 0x22,
	0x5c, 0x12, 0xec, 0xf8, 0x3e, 0xbf, 0x97, 0xdf, 0x03, 0x0f, 0x1d, 0x23, 0x84, 0x6b, 0x59, 0x85,
	0x0a, 0x53, 0x40, 0x1d, 0x27, 0xa2, 0x4c, 0x20, 0x8f, 0x11, 0xde, 0x77, 0xa0, 0x4d, 0xb0, 0x45,
	0x65, 0x14, 0x73, 0x11, 0x82, 0xb5, 0xac, 0x2e, 0xaf, 0xbb, 0x5f, 0x6d, 0x44, 0x99, 0x0a, 0x4c,
	0xe3, 0x02, 0xb4, 0x16, 0x19, 0xb4, 0x6f, 0xe3, 0x2f, 0x42, 0x2f, 0x96, 0x56, 0x33, 0xb3, 0x96,
	0xa8, 0x95, 0xb0, 0x2b, 0x3a, 0xd4, 0xa0, 0xb5, 0x54, 0xe5, 0x62, 0xee, 0x11, 0x9f, 0x4c, 0x86,
	0xd1, 0x7f, 0xc0, 0x6e, 0xa8, 0xbb, 0x01, 0x91, 0x02, 0x7a, 0x47, 0x3e, 0x99, 0x8c, 0xee, 0xcf,
	0x82, 0xb6, 0x2d, 0x78, 0xb6, 0x69, 0xd4, 0x51, 0xe6, 0xd3, 0x91, 0x42, 0x99, 0xcd, 0xf2, 0x25,
	0xa6, 0x8b, 0xb9, 0x77, 0x6c, 0x3d, 0x87, 0x11, 0xf3, 0xe8, 0x20, 0xe9, 0xe8, 0x89, 0xa5, 0xfd,
	0xc9, 0x6e, 0xe9, 0xc0, 0xa0, 0x90, 0x39, 0xa0, 0x77, 0x6a, 0x4b, 0xce, 0xfb, 0x92, 0x97, 0x36,
	0x8e, 0x7a, 0xfe, 0xf4, 0xf8, 0x5d, 0x73, 0xb2, 0xaf, 0x39, 0xf9, 0xad, 0x39, 0xf9, 0x6c, 0xb8,
	0xb3, 0x6f, 0xb8, 0xf3, 0xd3, 0x70, 0xe7, 0xd5, 0xcf, 0xa4, 0xd9, 0xec, 0x56, 0x41, 0xa2, 0x8a,
	0xf0, 0x4d, 0x16, 0xd3, 0x87, 0xbb, 0x30, 0x82, 0xb0, 0xb2, 0x8b, 0x98, 0x8f, 0x2d, 0xe8, 0x95,
	0x6b, 0x77, 0x98, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x1d, 0x17, 0x9f, 0x54, 0x01, 0x00,
	0x00,
}

func (m *OrdersCancelRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrdersCancelRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrdersCancelRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Trailer != nil {
		{
			size, err := m.Trailer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOrdersCancelRequest(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ClOrdID) > 0 {
		i -= len(m.ClOrdID)
		copy(dAtA[i:], m.ClOrdID)
		i = encodeVarintOrdersCancelRequest(dAtA, i, uint64(len(m.ClOrdID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OrigClOrdID) > 0 {
		i -= len(m.OrigClOrdID)
		copy(dAtA[i:], m.OrigClOrdID)
		i = encodeVarintOrdersCancelRequest(dAtA, i, uint64(len(m.OrigClOrdID)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOrdersCancelRequest(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.SessionID) > 0 {
		i -= len(m.SessionID)
		copy(dAtA[i:], m.SessionID)
		i = encodeVarintOrdersCancelRequest(dAtA, i, uint64(len(m.SessionID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrdersCancelRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrdersCancelRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrdersCancelRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SessionID)
	if l > 0 {
		n += 1 + l + sovOrdersCancelRequest(uint64(l))
	}
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovOrdersCancelRequest(uint64(l))
	}
	l = len(m.OrigClOrdID)
	if l > 0 {
		n += 1 + l + sovOrdersCancelRequest(uint64(l))
	}
	l = len(m.ClOrdID)
	if l > 0 {
		n += 1 + l + sovOrdersCancelRequest(uint64(l))
	}
	if m.Trailer != nil {
		l = m.Trailer.Size()
		n += 1 + l + sovOrdersCancelRequest(uint64(l))
	}
	return n
}

func sovOrdersCancelRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrdersCancelRequest(x uint64) (n int) {
	return sovOrdersCancelRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrdersCancelRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrdersCancelRequest
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
			return fmt.Errorf("proto: OrdersCancelRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrdersCancelRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersCancelRequest
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
				return ErrInvalidLengthOrdersCancelRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersCancelRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersCancelRequest
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
				return ErrInvalidLengthOrdersCancelRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersCancelRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &Header{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrigClOrdID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersCancelRequest
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
				return ErrInvalidLengthOrdersCancelRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersCancelRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrigClOrdID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClOrdID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersCancelRequest
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
				return ErrInvalidLengthOrdersCancelRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersCancelRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClOrdID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trailer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersCancelRequest
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
				return ErrInvalidLengthOrdersCancelRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersCancelRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Trailer == nil {
				m.Trailer = &Trailer{}
			}
			if err := m.Trailer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrdersCancelRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrdersCancelRequest
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
func skipOrdersCancelRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrdersCancelRequest
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
					return 0, ErrIntOverflowOrdersCancelRequest
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
					return 0, ErrIntOverflowOrdersCancelRequest
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
				return 0, ErrInvalidLengthOrdersCancelRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrdersCancelRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrdersCancelRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrdersCancelRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrdersCancelRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrdersCancelRequest = fmt.Errorf("proto: unexpected end of group")
)
