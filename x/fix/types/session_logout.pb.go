// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: re/fix/session_logout.proto

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

type SessionLogout struct {
	SessionName            string                  `protobuf:"bytes,2,opt,name=sessionName,proto3" json:"sessionName,omitempty"`
	InitiatorAddress       string                  `protobuf:"bytes,3,opt,name=initiatorAddress,proto3" json:"initiatorAddress,omitempty"`
	AcceptorAddress        string                  `protobuf:"bytes,4,opt,name=acceptorAddress,proto3" json:"acceptorAddress,omitempty"`
	SessionLogoutInitiator *SessionLogoutInitiator `protobuf:"bytes,5,opt,name=sessionLogoutInitiator,proto3" json:"sessionLogoutInitiator,omitempty"`
	SessionLogoutAcceptor  *SessionLogoutAcceptor  `protobuf:"bytes,6,opt,name=sessionLogoutAcceptor,proto3" json:"sessionLogoutAcceptor,omitempty"`
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

func (m *SessionLogout) GetSessionName() string {
	if m != nil {
		return m.SessionName
	}
	return ""
}

func (m *SessionLogout) GetInitiatorAddress() string {
	if m != nil {
		return m.InitiatorAddress
	}
	return ""
}

func (m *SessionLogout) GetAcceptorAddress() string {
	if m != nil {
		return m.AcceptorAddress
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

type SessionLogoutInitiator struct {
	Header  *Header  `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Text    string   `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Trailer *Trailer `protobuf:"bytes,4,opt,name=trailer,proto3" json:"trailer,omitempty"`
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

type SessionLogoutAcceptor struct {
	Header  *Header  `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Text    string   `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Trailer *Trailer `protobuf:"bytes,4,opt,name=trailer,proto3" json:"trailer,omitempty"`
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
	proto.RegisterType((*SessionLogout)(nil), "jim380.re.fix.SessionLogout")
	proto.RegisterType((*SessionLogoutInitiator)(nil), "jim380.re.fix.SessionLogoutInitiator")
	proto.RegisterType((*SessionLogoutAcceptor)(nil), "jim380.re.fix.SessionLogoutAcceptor")
}

func init() { proto.RegisterFile("re/fix/session_logout.proto", fileDescriptor_997949e9dc500e6d) }

var fileDescriptor_997949e9dc500e6d = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0x41, 0x4b, 0x02, 0x41,
	0x14, 0xc7, 0x1d, 0x33, 0xa3, 0x59, 0xa4, 0x18, 0x50, 0x96, 0xa2, 0x65, 0x91, 0x02, 0x09, 0xda,
	0x15, 0xbd, 0x44, 0x37, 0x3b, 0x15, 0x44, 0x87, 0xb1, 0x93, 0x10, 0x32, 0xba, 0x4f, 0x9d, 0x70,
	0x1d, 0x99, 0x19, 0x61, 0xfb, 0x0e, 0x05, 0x7d, 0xa6, 0x4e, 0x1d, 0x3d, 0x76, 0x0c, 0xfd, 0x22,
	0xc1, 0xec, 0x2e, 0xe8, 0xb6, 0x75, 0xed, 0x36, 0xbc, 0xf7, 0xe3, 0x37, 0xef, 0xfd, 0x79, 0xf8,
	0x58, 0x82, 0x3f, 0xe2, 0x91, 0xaf, 0x40, 0x29, 0x2e, 0x66, 0xfd, 0xa9, 0x18, 0x8b, 0x85, 0xf6,
	0xe6, 0x52, 0x68, 0x41, 0x2a, 0x4f, 0x3c, 0x6c, 0x5f, 0x36, 0x3d, 0x09, 0xde, 0x88, 0x47, 0x47,
	0x27, 0x29, 0xab, 0xd9, 0x2c, 0x60, 0x32, 0xe8, 0x87, 0xa0, 0x14, 0x1b, 0x43, 0x4c, 0xd7, 0xdf,
	0x8b, 0xb8, 0xd2, 0x8d, 0x35, 0x77, 0xc6, 0x42, 0x5c, 0x6c, 0x25, 0xde, 0x7b, 0x16, 0x82, 0x5d,
	0x74, 0x51, 0x63, 0x9f, 0x6e, 0x96, 0xc8, 0x39, 0x3e, 0xe4, 0x33, 0xae, 0x39, 0xd3, 0x42, 0x76,
	0x82, 0x40, 0x82, 0x52, 0xf6, 0x8e, 0xc1, 0x7e, 0xd4, 0x49, 0x03, 0x1f, 0xb0, 0xe1, 0x10, 0xe6,
	0x1b, 0x68, 0xc9, 0xa0, 0xd9, 0x32, 0x79, 0xc4, 0x35, 0xb5, 0x39, 0xc8, 0x6d, 0xaa, 0xb2, 0x77,
	0x5d, 0xd4, 0xb0, 0x5a, 0x67, 0xde, 0xd6, 0x62, 0x5e, 0x37, 0x17, 0xa6, 0xbf, 0x48, 0x48, 0x0f,
	0x57, 0xb7, 0x3a, 0x9d, 0xe4, 0x7b, 0xbb, 0x6c, 0xec, 0xa7, 0x7f, 0xd9, 0x53, 0x96, 0xe6, 0x2b,
	0xea, 0xaf, 0x08, 0xd7, 0xf2, 0xc7, 0x21, 0x17, 0xb8, 0x3c, 0x01, 0x16, 0x80, 0x34, 0x41, 0x5a,
	0xad, 0x6a, 0xe6, 0x9f, 0x1b, 0xd3, 0xa4, 0x09, 0x44, 0x08, 0x2e, 0x69, 0x88, 0x74, 0x12, 0xa7,
	0x79, 0x93, 0x26, 0xde, 0xd3, 0x92, 0xf1, 0x29, 0x48, 0x13, 0x9d, 0xd5, 0xaa, 0x65, 0x1c, 0x0f,
	0x71, 0x97, 0xa6, 0x58, 0xfd, 0x05, 0xe1, 0x6a, 0xee, 0x02, 0xff, 0x32, 0xce, 0xf5, 0xd5, 0xc7,
	0xca, 0x41, 0xcb, 0x95, 0x83, 0xbe, 0x56, 0x0e, 0x7a, 0x5b, 0x3b, 0x85, 0xe5, 0xda, 0x29, 0x7c,
	0xae, 0x9d, 0x42, 0xcf, 0x1d, 0x73, 0x3d, 0x59, 0x0c, 0xbc, 0xa1, 0x08, 0xfd, 0x58, 0xe2, 0x53,
	0xf0, 0x23, 0x73, 0xb0, 0xfa, 0x79, 0x0e, 0x6a, 0x50, 0x36, 0x67, 0xda, 0xfe, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x28, 0x42, 0x97, 0x56, 0xf3, 0x02, 0x00, 0x00,
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
		dAtA[i] = 0x32
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
		dAtA[i] = 0x2a
	}
	if len(m.AcceptorAddress) > 0 {
		i -= len(m.AcceptorAddress)
		copy(dAtA[i:], m.AcceptorAddress)
		i = encodeVarintSessionLogout(dAtA, i, uint64(len(m.AcceptorAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.InitiatorAddress) > 0 {
		i -= len(m.InitiatorAddress)
		copy(dAtA[i:], m.InitiatorAddress)
		i = encodeVarintSessionLogout(dAtA, i, uint64(len(m.InitiatorAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SessionName) > 0 {
		i -= len(m.SessionName)
		copy(dAtA[i:], m.SessionName)
		i = encodeVarintSessionLogout(dAtA, i, uint64(len(m.SessionName)))
		i--
		dAtA[i] = 0x12
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
		dAtA[i] = 0x22
	}
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintSessionLogout(dAtA, i, uint64(len(m.Text)))
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
			i = encodeVarintSessionLogout(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
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
		dAtA[i] = 0x22
	}
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintSessionLogout(dAtA, i, uint64(len(m.Text)))
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
			i = encodeVarintSessionLogout(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
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
	l = len(m.SessionName)
	if l > 0 {
		n += 1 + l + sovSessionLogout(uint64(l))
	}
	l = len(m.InitiatorAddress)
	if l > 0 {
		n += 1 + l + sovSessionLogout(uint64(l))
	}
	l = len(m.AcceptorAddress)
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionName", wireType)
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
			m.SessionName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitiatorAddress", wireType)
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
			m.InitiatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AcceptorAddress", wireType)
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
			m.AcceptorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
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
		case 6:
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
		case 2:
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
		case 3:
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
		case 4:
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
		case 2:
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
		case 3:
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
		case 4:
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
