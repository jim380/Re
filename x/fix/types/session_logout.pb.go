// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: re/fix/session_logout.proto

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

// This message is used to initiate a logout session between two parties
type SessionLogout struct {
	// A string that identifies the session being logged out of
	SessionID string `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	// The message sent by the party that initiates the logout
	SessionLogoutInitiator *SessionLogoutInitiator `protobuf:"bytes,2,opt,name=sessionLogoutInitiator,proto3" json:"sessionLogoutInitiator,omitempty"`
	// The message sent by the party that confirms the logout
	SessionLogoutAcceptor *SessionLogoutAcceptor `protobuf:"bytes,3,opt,name=sessionLogoutAcceptor,proto3" json:"sessionLogoutAcceptor,omitempty"`
}

func (m *SessionLogout) Reset()         { *m = SessionLogout{} }
func (m *SessionLogout) String() string { return proto.CompactTextString(m) }
func (*SessionLogout) ProtoMessage()    {}
func (*SessionLogout) Descriptor() ([]byte, []int) {
	return fileDescriptor_997949e9dc500e6d, []int{0}
}
func (m *SessionLogout) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SessionLogout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SessionLogout.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SessionLogout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionLogout.Merge(m, src)
}
func (m *SessionLogout) XXX_Size() int {
	return m.Size()
}
func (m *SessionLogout) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionLogout.DiscardUnknown(m)
}

var xxx_messageInfo_SessionLogout proto.InternalMessageInfo

func (m *SessionLogout) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *SessionLogout) GetSessionLogoutInitiator() *SessionLogoutInitiator {
	if m != nil {
		return m.SessionLogoutInitiator
	}
	return nil
}

func (m *SessionLogout) GetSessionLogoutAcceptor() *SessionLogoutAcceptor {
	if m != nil {
		return m.SessionLogoutAcceptor
	}
	return nil
}

// This message is sent by the party that initiates the logout
type SessionLogoutInitiator struct {
	// A FIX protocol header containing standard message fields such as
	// beginString, bodyLength, msgType, etc
	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// A free-form text field that can be used to provide additional information
	// about the logout
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	// A FIX protocol trailer containing standard message fields such as checkSum
	Trailer *Trailer `protobuf:"bytes,3,opt,name=trailer,proto3" json:"trailer,omitempty"`
}

func (m *SessionLogoutInitiator) Reset()         { *m = SessionLogoutInitiator{} }
func (m *SessionLogoutInitiator) String() string { return proto.CompactTextString(m) }
func (*SessionLogoutInitiator) ProtoMessage()    {}
func (*SessionLogoutInitiator) Descriptor() ([]byte, []int) {
	return fileDescriptor_997949e9dc500e6d, []int{1}
}
func (m *SessionLogoutInitiator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SessionLogoutInitiator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SessionLogoutInitiator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SessionLogoutInitiator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionLogoutInitiator.Merge(m, src)
}
func (m *SessionLogoutInitiator) XXX_Size() int {
	return m.Size()
}
func (m *SessionLogoutInitiator) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionLogoutInitiator.DiscardUnknown(m)
}

var xxx_messageInfo_SessionLogoutInitiator proto.InternalMessageInfo

func (m *SessionLogoutInitiator) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SessionLogoutInitiator) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *SessionLogoutInitiator) GetTrailer() *Trailer {
	if m != nil {
		return m.Trailer
	}
	return nil
}

// This message is sent by the party that confirms the logout
type SessionLogoutAcceptor struct {
	// A FIX protocol header containing standard message fields such as
	// beginString, bodyLength, msgType, etc
	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// A free-form text field that can be used to provide additional information
	// about the logout
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	// A FIX protocol trailer containing standard message fields such as checkSum
	Trailer *Trailer `protobuf:"bytes,3,opt,name=trailer,proto3" json:"trailer,omitempty"`
}

func (m *SessionLogoutAcceptor) Reset()         { *m = SessionLogoutAcceptor{} }
func (m *SessionLogoutAcceptor) String() string { return proto.CompactTextString(m) }
func (*SessionLogoutAcceptor) ProtoMessage()    {}
func (*SessionLogoutAcceptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_997949e9dc500e6d, []int{2}
}
func (m *SessionLogoutAcceptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SessionLogoutAcceptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SessionLogoutAcceptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SessionLogoutAcceptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionLogoutAcceptor.Merge(m, src)
}
func (m *SessionLogoutAcceptor) XXX_Size() int {
	return m.Size()
}
func (m *SessionLogoutAcceptor) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionLogoutAcceptor.DiscardUnknown(m)
}

var xxx_messageInfo_SessionLogoutAcceptor proto.InternalMessageInfo

func (m *SessionLogoutAcceptor) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SessionLogoutAcceptor) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *SessionLogoutAcceptor) GetTrailer() *Trailer {
	if m != nil {
		return m.Trailer
	}
	return nil
}

func init() {
	proto.RegisterType((*SessionLogout)(nil), "re.fix.SessionLogout")
	proto.RegisterType((*SessionLogoutInitiator)(nil), "re.fix.SessionLogoutInitiator")
	proto.RegisterType((*SessionLogoutAcceptor)(nil), "re.fix.SessionLogoutAcceptor")
}

func init() { proto.RegisterFile("re/fix/session_logout.proto", fileDescriptor_997949e9dc500e6d) }

var fileDescriptor_997949e9dc500e6d = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x4a, 0xd5, 0x4f,
	0xcb, 0xac, 0xd0, 0x2f, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0x8b, 0xcf, 0xc9, 0x4f, 0xcf, 0x2f,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0x4a, 0xd5, 0x4b, 0xcb, 0xac, 0x90,
	0x92, 0x85, 0x29, 0x2a, 0x49, 0xcc, 0x4b, 0x49, 0x2c, 0x4a, 0x89, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e,
	0x4c, 0x4f, 0x85, 0x28, 0x53, 0xba, 0xc5, 0xc8, 0xc5, 0x1b, 0x0c, 0xd1, 0xef, 0x03, 0xd6, 0x2e,
	0x24, 0xc3, 0xc5, 0x09, 0x35, 0xd0, 0xd3, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0x21,
	0x20, 0x14, 0xc6, 0x25, 0x56, 0x8c, 0xac, 0xdc, 0x33, 0x2f, 0xb3, 0x24, 0x33, 0xb1, 0x24, 0xbf,
	0x48, 0x82, 0x49, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x4e, 0x0f, 0x62, 0xaf, 0x5e, 0x30, 0x56, 0x55,
	0x41, 0x38, 0x74, 0x0b, 0x05, 0x73, 0x89, 0xa2, 0xc8, 0x38, 0x26, 0x27, 0xa7, 0x16, 0x80, 0x8c,
	0x65, 0x06, 0x1b, 0x2b, 0x8b, 0xd5, 0x58, 0x98, 0xa2, 0x20, 0xec, 0x7a, 0x95, 0xea, 0xb9, 0xc4,
	0xb0, 0x3b, 0x43, 0x48, 0x8d, 0x8b, 0x2d, 0x23, 0x35, 0x31, 0x25, 0xb5, 0x08, 0xec, 0x43, 0x6e,
	0x23, 0x3e, 0x98, 0xf9, 0x1e, 0x60, 0xd1, 0x20, 0xa8, 0xac, 0x90, 0x10, 0x17, 0x4b, 0x49, 0x6a,
	0x45, 0x09, 0xd8, 0x73, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x26, 0x17, 0x7b, 0x49, 0x51, 0x62, 0x66,
	0x4e, 0x2a, 0xcc, 0x71, 0xfc, 0x30, 0xcd, 0x21, 0x10, 0xe1, 0x20, 0x98, 0xbc, 0x52, 0x1d, 0x97,
	0x28, 0x56, 0x07, 0xd3, 0xc9, 0x7e, 0x27, 0xab, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63,
	0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96,
	0x63, 0x88, 0x52, 0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xca,
	0xcc, 0x35, 0xb6, 0x30, 0xd0, 0x0f, 0x4a, 0xd5, 0xaf, 0x00, 0x27, 0x95, 0x92, 0xca, 0x82, 0xd4,
	0xe2, 0x24, 0x36, 0x70, 0x02, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x35, 0x29, 0x31, 0xf9,
	0x66, 0x02, 0x00, 0x00,
}

func (m *SessionLogout) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionLogout) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SessionLogout) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SessionLogoutAcceptor != nil {
		{
			size, err := m.SessionLogoutAcceptor.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSessionLogout(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.SessionLogoutInitiator != nil {
		{
			size, err := m.SessionLogoutInitiator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSessionLogout(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.SessionID) > 0 {
		i -= len(m.SessionID)
		copy(dAtA[i:], m.SessionID)
		i = encodeVarintSessionLogout(dAtA, i, uint64(len(m.SessionID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SessionLogoutInitiator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionLogoutInitiator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SessionLogoutInitiator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintSessionLogout(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintSessionLogout(dAtA, i, uint64(len(m.Text)))
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSessionLogout(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SessionLogoutAcceptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionLogoutAcceptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SessionLogoutAcceptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintSessionLogout(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintSessionLogout(dAtA, i, uint64(len(m.Text)))
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSessionLogout(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSessionLogout(dAtA []byte, offset int, v uint64) int {
	offset -= sovSessionLogout(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SessionLogout) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SessionID)
	if l > 0 {
		n += 1 + l + sovSessionLogout(uint64(l))
	}
	if m.SessionLogoutInitiator != nil {
		l = m.SessionLogoutInitiator.Size()
		n += 1 + l + sovSessionLogout(uint64(l))
	}
	if m.SessionLogoutAcceptor != nil {
		l = m.SessionLogoutAcceptor.Size()
		n += 1 + l + sovSessionLogout(uint64(l))
	}
	return n
}

func (m *SessionLogoutInitiator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovSessionLogout(uint64(l))
	}
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovSessionLogout(uint64(l))
	}
	if m.Trailer != nil {
		l = m.Trailer.Size()
		n += 1 + l + sovSessionLogout(uint64(l))
	}
	return n
}

func (m *SessionLogoutAcceptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovSessionLogout(uint64(l))
	}
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovSessionLogout(uint64(l))
	}
	if m.Trailer != nil {
		l = m.Trailer.Size()
		n += 1 + l + sovSessionLogout(uint64(l))
	}
	return n
}

func sovSessionLogout(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSessionLogout(x uint64) (n int) {
	return sovSessionLogout(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SessionLogout) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSessionLogout
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
			return fmt.Errorf("proto: SessionLogout: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionLogout: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionLogout
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
				return ErrInvalidLengthSessionLogout
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSessionLogout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionLogoutInitiator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionLogout
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
				return ErrInvalidLengthSessionLogout
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSessionLogout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SessionLogoutInitiator == nil {
				m.SessionLogoutInitiator = &SessionLogoutInitiator{}
			}
			if err := m.SessionLogoutInitiator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionLogoutAcceptor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionLogout
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
				return ErrInvalidLengthSessionLogout
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSessionLogout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SessionLogoutAcceptor == nil {
				m.SessionLogoutAcceptor = &SessionLogoutAcceptor{}
			}
			if err := m.SessionLogoutAcceptor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSessionLogout(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSessionLogout
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
func (m *SessionLogoutInitiator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSessionLogout
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
			return fmt.Errorf("proto: SessionLogoutInitiator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionLogoutInitiator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionLogout
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
				return ErrInvalidLengthSessionLogout
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSessionLogout
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionLogout
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
				return ErrInvalidLengthSessionLogout
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSessionLogout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trailer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionLogout
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
				return ErrInvalidLengthSessionLogout
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSessionLogout
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
			skippy, err := skipSessionLogout(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSessionLogout
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
func (m *SessionLogoutAcceptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSessionLogout
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
			return fmt.Errorf("proto: SessionLogoutAcceptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionLogoutAcceptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionLogout
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
				return ErrInvalidLengthSessionLogout
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSessionLogout
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionLogout
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
				return ErrInvalidLengthSessionLogout
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSessionLogout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trailer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionLogout
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
				return ErrInvalidLengthSessionLogout
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSessionLogout
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
			skippy, err := skipSessionLogout(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSessionLogout
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
func skipSessionLogout(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSessionLogout
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
					return 0, ErrIntOverflowSessionLogout
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
					return 0, ErrIntOverflowSessionLogout
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
				return 0, ErrInvalidLengthSessionLogout
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSessionLogout
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSessionLogout
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSessionLogout        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSessionLogout          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSessionLogout = fmt.Errorf("proto: unexpected end of group")
)
