// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: re/fix/v1beta1/session_reject.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// This message is used to reject a FIX session
type SessionReject struct {
	// A string that identifies the session to which the rejected message belongs.
	SessionID string `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	// A FIX protocol header containing standard message fields such as
	// beginString, bodyLength, msgType, etc
	Header *Header `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	// A free-form text field that provides information about the rejection, such
	// as the reason for the rejection or a description of the error
	Text string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	// A FIX protocol trailer containing standard message fields such as checkSum
	Trailer *Trailer `protobuf:"bytes,4,opt,name=trailer,proto3" json:"trailer,omitempty"`
}

func (m *SessionReject) Reset()         { *m = SessionReject{} }
func (m *SessionReject) String() string { return proto.CompactTextString(m) }
func (*SessionReject) ProtoMessage()    {}
func (*SessionReject) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b4355740de74819, []int{0}
}
func (m *SessionReject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SessionReject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SessionReject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SessionReject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionReject.Merge(m, src)
}
func (m *SessionReject) XXX_Size() int {
	return m.Size()
}
func (m *SessionReject) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionReject.DiscardUnknown(m)
}

var xxx_messageInfo_SessionReject proto.InternalMessageInfo

func (m *SessionReject) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *SessionReject) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SessionReject) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *SessionReject) GetTrailer() *Trailer {
	if m != nil {
		return m.Trailer
	}
	return nil
}

func init() {
	proto.RegisterType((*SessionReject)(nil), "re.fix.SessionReject")
}

func init() {
	proto.RegisterFile("re/fix/v1beta1/session_reject.proto", fileDescriptor_4b4355740de74819)
}

var fileDescriptor_4b4355740de74819 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x4a, 0xd5, 0x4f,
	0xcb, 0xac, 0xd0, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x4e, 0x2d, 0x2e, 0xce,
	0xcc, 0xcf, 0x8b, 0x2f, 0x4a, 0xcd, 0x4a, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2b, 0x4a, 0xd5, 0x4b, 0xcb, 0xac, 0x90, 0x52, 0x45, 0x57, 0x5c, 0x92, 0x98, 0x97, 0x92,
	0x58, 0x94, 0x12, 0x9f, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x0a, 0x51, 0xae, 0x34, 0x85, 0x91,
	0x8b, 0x37, 0x18, 0x62, 0x4e, 0x10, 0xd8, 0x18, 0x21, 0x19, 0x2e, 0x4e, 0xa8, 0xc1, 0x9e, 0x2e,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x08, 0x01, 0x21, 0x35, 0x2e, 0xb6, 0x8c, 0xd4, 0xc4,
	0x94, 0xd4, 0x22, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x3e, 0x3d, 0x88, 0x7d, 0x7a, 0x1e,
	0x60, 0xd1, 0x20, 0xa8, 0xac, 0x90, 0x10, 0x17, 0x4b, 0x49, 0x6a, 0x45, 0x89, 0x04, 0x33, 0xd8,
	0x00, 0x30, 0x5b, 0x48, 0x93, 0x8b, 0xbd, 0xa4, 0x28, 0x31, 0x33, 0x27, 0xb5, 0x48, 0x82, 0x05,
	0xac, 0x99, 0x1f, 0xa6, 0x39, 0x04, 0x22, 0x1c, 0x04, 0x93, 0x77, 0xb2, 0x3a, 0xf1, 0x48, 0x8e,
	0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58,
	0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x85, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4,
	0xfc, 0x5c, 0xfd, 0xac, 0xcc, 0x5c, 0x63, 0x0b, 0x03, 0xfd, 0xa0, 0x54, 0xfd, 0x0a, 0xb0, 0x5f,
	0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x3e, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x45, 0xad, 0x9e, 0x92, 0x2f, 0x01, 0x00, 0x00,
}

func (m *SessionReject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionReject) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SessionReject) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintSessionReject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintSessionReject(dAtA, i, uint64(len(m.Text)))
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
			i = encodeVarintSessionReject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.SessionID) > 0 {
		i -= len(m.SessionID)
		copy(dAtA[i:], m.SessionID)
		i = encodeVarintSessionReject(dAtA, i, uint64(len(m.SessionID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSessionReject(dAtA []byte, offset int, v uint64) int {
	offset -= sovSessionReject(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SessionReject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SessionID)
	if l > 0 {
		n += 1 + l + sovSessionReject(uint64(l))
	}
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovSessionReject(uint64(l))
	}
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovSessionReject(uint64(l))
	}
	if m.Trailer != nil {
		l = m.Trailer.Size()
		n += 1 + l + sovSessionReject(uint64(l))
	}
	return n
}

func sovSessionReject(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSessionReject(x uint64) (n int) {
	return sovSessionReject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SessionReject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSessionReject
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
			return fmt.Errorf("proto: SessionReject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionReject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionReject
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
				return ErrInvalidLengthSessionReject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSessionReject
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
					return ErrIntOverflowSessionReject
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
				return ErrInvalidLengthSessionReject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSessionReject
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
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionReject
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
				return ErrInvalidLengthSessionReject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSessionReject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trailer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionReject
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
				return ErrInvalidLengthSessionReject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSessionReject
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
			skippy, err := skipSessionReject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSessionReject
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
func skipSessionReject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSessionReject
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
					return 0, ErrIntOverflowSessionReject
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
					return 0, ErrIntOverflowSessionReject
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
				return 0, ErrInvalidLengthSessionReject
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSessionReject
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSessionReject
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSessionReject        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSessionReject          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSessionReject = fmt.Errorf("proto: unexpected end of group")
)
