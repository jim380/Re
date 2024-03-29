// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: re/fix/v1beta1/sessions.proto

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

type Sessions struct {
	SessionID      string          `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	LogonInitiator *LogonInitiator `protobuf:"bytes,2,opt,name=logonInitiator,proto3" json:"logonInitiator,omitempty"`
	LogonAcceptor  *LogonAcceptor  `protobuf:"bytes,3,opt,name=logonAcceptor,proto3" json:"logonAcceptor,omitempty"`
	Status         string          `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	IsAccepted     bool            `protobuf:"varint,5,opt,name=IsAccepted,proto3" json:"IsAccepted,omitempty"`
}

func (m *Sessions) Reset()         { *m = Sessions{} }
func (m *Sessions) String() string { return proto.CompactTextString(m) }
func (*Sessions) ProtoMessage()    {}
func (*Sessions) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c6d8ee69d8468a9, []int{0}
}
func (m *Sessions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Sessions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sessions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Sessions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sessions.Merge(m, src)
}
func (m *Sessions) XXX_Size() int {
	return m.Size()
}
func (m *Sessions) XXX_DiscardUnknown() {
	xxx_messageInfo_Sessions.DiscardUnknown(m)
}

var xxx_messageInfo_Sessions proto.InternalMessageInfo

func (m *Sessions) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *Sessions) GetLogonInitiator() *LogonInitiator {
	if m != nil {
		return m.LogonInitiator
	}
	return nil
}

func (m *Sessions) GetLogonAcceptor() *LogonAcceptor {
	if m != nil {
		return m.LogonAcceptor
	}
	return nil
}

func (m *Sessions) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Sessions) GetIsAccepted() bool {
	if m != nil {
		return m.IsAccepted
	}
	return false
}

// This message is used by the initiator of a FIX session to initiate a
// connection with the acceptor
type LogonInitiator struct {
	// The FIX message header contains standard metadata about the message,
	// including the sender and receiver IDs and the message type
	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// An integer value indicating the encryption method to be used for the
	// session
	EncryptMethod int64 `protobuf:"varint,2,opt,name=encryptMethod,proto3" json:"encryptMethod,omitempty"`
	// An integer value indicating the number of seconds between heartbeats that
	// will be sent during the session. This is used to monitor the connection
	// between the two parties
	HeartBtInt int64 `protobuf:"varint,3,opt,name=heartBtInt,proto3" json:"heartBtInt,omitempty"`
	// The FIX message trailer contains standard metadata about the message,
	// including the message checksum
	Trailer *Trailer `protobuf:"bytes,4,opt,name=trailer,proto3" json:"trailer,omitempty"`
}

func (m *LogonInitiator) Reset()         { *m = LogonInitiator{} }
func (m *LogonInitiator) String() string { return proto.CompactTextString(m) }
func (*LogonInitiator) ProtoMessage()    {}
func (*LogonInitiator) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c6d8ee69d8468a9, []int{1}
}
func (m *LogonInitiator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LogonInitiator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LogonInitiator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LogonInitiator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogonInitiator.Merge(m, src)
}
func (m *LogonInitiator) XXX_Size() int {
	return m.Size()
}
func (m *LogonInitiator) XXX_DiscardUnknown() {
	xxx_messageInfo_LogonInitiator.DiscardUnknown(m)
}

var xxx_messageInfo_LogonInitiator proto.InternalMessageInfo

func (m *LogonInitiator) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *LogonInitiator) GetEncryptMethod() int64 {
	if m != nil {
		return m.EncryptMethod
	}
	return 0
}

func (m *LogonInitiator) GetHeartBtInt() int64 {
	if m != nil {
		return m.HeartBtInt
	}
	return 0
}

func (m *LogonInitiator) GetTrailer() *Trailer {
	if m != nil {
		return m.Trailer
	}
	return nil
}

// This message is used by the acceptor of a FIX session to acknowledge the
// initiation of the connection by the initiator
type LogonAcceptor struct {
	// The FIX message header contains standard metadata about the message,
	// including the sender and receiver IDs and the message type
	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// An integer value indicating the encryption method to be used for the
	// session
	EncryptMethod int64 `protobuf:"varint,2,opt,name=encryptMethod,proto3" json:"encryptMethod,omitempty"`
	// An integer value indicating the number of seconds between heartbeats that
	// will be sent during the session. This is used to monitor the connection
	// between the two parties
	HeartBtInt int64 `protobuf:"varint,3,opt,name=heartBtInt,proto3" json:"heartBtInt,omitempty"`
	// The FIX message trailer contains standard metadata about the message,
	// including the message checksum
	Trailer *Trailer `protobuf:"bytes,4,opt,name=trailer,proto3" json:"trailer,omitempty"`
}

func (m *LogonAcceptor) Reset()         { *m = LogonAcceptor{} }
func (m *LogonAcceptor) String() string { return proto.CompactTextString(m) }
func (*LogonAcceptor) ProtoMessage()    {}
func (*LogonAcceptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c6d8ee69d8468a9, []int{2}
}
func (m *LogonAcceptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LogonAcceptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LogonAcceptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LogonAcceptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogonAcceptor.Merge(m, src)
}
func (m *LogonAcceptor) XXX_Size() int {
	return m.Size()
}
func (m *LogonAcceptor) XXX_DiscardUnknown() {
	xxx_messageInfo_LogonAcceptor.DiscardUnknown(m)
}

var xxx_messageInfo_LogonAcceptor proto.InternalMessageInfo

func (m *LogonAcceptor) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *LogonAcceptor) GetEncryptMethod() int64 {
	if m != nil {
		return m.EncryptMethod
	}
	return 0
}

func (m *LogonAcceptor) GetHeartBtInt() int64 {
	if m != nil {
		return m.HeartBtInt
	}
	return 0
}

func (m *LogonAcceptor) GetTrailer() *Trailer {
	if m != nil {
		return m.Trailer
	}
	return nil
}

func init() {
	proto.RegisterType((*Sessions)(nil), "re.fix.Sessions")
	proto.RegisterType((*LogonInitiator)(nil), "re.fix.LogonInitiator")
	proto.RegisterType((*LogonAcceptor)(nil), "re.fix.LogonAcceptor")
}

func init() { proto.RegisterFile("re/fix/v1beta1/sessions.proto", fileDescriptor_7c6d8ee69d8468a9) }

var fileDescriptor_7c6d8ee69d8468a9 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x86, 0xbd, 0x75, 0xab, 0xda, 0x6b, 0xec, 0xc2, 0x42, 0x8d, 0x28, 0xad, 0x10, 0xa6, 0x2d,
	0xee, 0x45, 0xaa, 0xed, 0x4b, 0x48, 0x20, 0x10, 0x93, 0x43, 0x04, 0xc9, 0x65, 0x93, 0x53, 0x2e,
	0x61, 0x2d, 0x8d, 0xad, 0x0d, 0xb6, 0x56, 0xec, 0xae, 0x83, 0xfd, 0x16, 0x79, 0x8c, 0xe4, 0x4d,
	0x72, 0xf4, 0x31, 0x97, 0x40, 0xb0, 0x5f, 0x24, 0x78, 0x25, 0x11, 0xcb, 0x4f, 0x90, 0xe3, 0x7e,
	0xff, 0x3f, 0x33, 0xff, 0x2c, 0x83, 0x7f, 0x49, 0xf0, 0xc7, 0x7c, 0xe1, 0xdf, 0xf5, 0x46, 0xa0,
	0x59, 0xcf, 0x57, 0xa0, 0x14, 0x17, 0x89, 0xf2, 0x52, 0x29, 0xb4, 0x20, 0x96, 0x04, 0x6f, 0xcc,
	0x17, 0x3f, 0xfe, 0xec, 0xdb, 0x34, 0x4b, 0x22, 0x26, 0xa3, 0x9b, 0x19, 0x28, 0xc5, 0x26, 0x90,
	0xd9, 0x3b, 0x2f, 0x08, 0xd7, 0x2e, 0xf3, 0x0e, 0xe4, 0x27, 0xae, 0xe7, 0xdd, 0x82, 0x53, 0x1b,
	0xb9, 0xa8, 0x5b, 0xa7, 0xef, 0x80, 0x1c, 0xe3, 0xd6, 0x54, 0x4c, 0x44, 0x12, 0x24, 0x5c, 0x73,
	0xa6, 0x85, 0xb4, 0x3f, 0xb9, 0xa8, 0xdb, 0xe8, 0xb7, 0xbd, 0x6c, 0xa4, 0x77, 0x5e, 0x52, 0xe9,
	0x9e, 0x9b, 0x1c, 0xe1, 0xa6, 0x21, 0x27, 0x61, 0x08, 0xe9, 0xb6, 0xbc, 0x6a, 0xca, 0xbf, 0x97,
	0xca, 0x0b, 0x91, 0x96, 0xbd, 0xa4, 0x8d, 0x2d, 0xa5, 0x99, 0x9e, 0x2b, 0xfb, 0xb3, 0xc9, 0x95,
	0xbf, 0x88, 0x83, 0x71, 0xa0, 0x32, 0x17, 0x44, 0xf6, 0x17, 0x17, 0x75, 0x6b, 0x74, 0x87, 0x74,
	0x1e, 0x11, 0x6e, 0x95, 0x73, 0x91, 0xbf, 0xd8, 0x8a, 0x81, 0x45, 0x20, 0xcd, 0x8a, 0x8d, 0x7e,
	0xab, 0x08, 0x70, 0x66, 0x28, 0xcd, 0x55, 0xf2, 0x1b, 0x37, 0x21, 0x09, 0xe5, 0x32, 0xd5, 0x17,
	0xa0, 0x63, 0x11, 0x99, 0x75, 0xab, 0xb4, 0x0c, 0xb7, 0x01, 0x62, 0x60, 0x52, 0x0f, 0x75, 0x90,
	0x68, 0xb3, 0x52, 0x95, 0xee, 0x10, 0xf2, 0x0f, 0x7f, 0xd5, 0x92, 0xf1, 0x29, 0x48, 0x93, 0xbc,
	0xd1, 0xff, 0x56, 0x8c, 0xbb, 0xca, 0x30, 0x2d, 0xf4, 0xce, 0x03, 0xc2, 0xcd, 0xd2, 0x27, 0x7c,
	0xd8, 0xa8, 0xc3, 0xc3, 0xa7, 0xb5, 0x83, 0x56, 0x6b, 0x07, 0xbd, 0xae, 0x1d, 0x74, 0xbf, 0x71,
	0x2a, 0xab, 0x8d, 0x53, 0x79, 0xde, 0x38, 0x95, 0x6b, 0x77, 0xc2, 0x75, 0x3c, 0x1f, 0x79, 0xa1,
	0x98, 0xf9, 0xb7, 0x7c, 0x36, 0x38, 0xf8, 0xef, 0x53, 0xf0, 0x17, 0xe6, 0x16, 0xf5, 0x32, 0x05,
	0x35, 0xb2, 0xcc, 0xe5, 0x0d, 0xde, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x19, 0x6b, 0xd4, 0xc9,
	0x02, 0x00, 0x00,
}

func (m *Sessions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sessions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Sessions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsAccepted {
		i--
		if m.IsAccepted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintSessions(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x22
	}
	if m.LogonAcceptor != nil {
		{
			size, err := m.LogonAcceptor.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSessions(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.LogonInitiator != nil {
		{
			size, err := m.LogonInitiator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSessions(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.SessionID) > 0 {
		i -= len(m.SessionID)
		copy(dAtA[i:], m.SessionID)
		i = encodeVarintSessions(dAtA, i, uint64(len(m.SessionID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LogonInitiator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogonInitiator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LogonInitiator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintSessions(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.HeartBtInt != 0 {
		i = encodeVarintSessions(dAtA, i, uint64(m.HeartBtInt))
		i--
		dAtA[i] = 0x18
	}
	if m.EncryptMethod != 0 {
		i = encodeVarintSessions(dAtA, i, uint64(m.EncryptMethod))
		i--
		dAtA[i] = 0x10
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSessions(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LogonAcceptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogonAcceptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LogonAcceptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintSessions(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.HeartBtInt != 0 {
		i = encodeVarintSessions(dAtA, i, uint64(m.HeartBtInt))
		i--
		dAtA[i] = 0x18
	}
	if m.EncryptMethod != 0 {
		i = encodeVarintSessions(dAtA, i, uint64(m.EncryptMethod))
		i--
		dAtA[i] = 0x10
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSessions(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSessions(dAtA []byte, offset int, v uint64) int {
	offset -= sovSessions(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Sessions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SessionID)
	if l > 0 {
		n += 1 + l + sovSessions(uint64(l))
	}
	if m.LogonInitiator != nil {
		l = m.LogonInitiator.Size()
		n += 1 + l + sovSessions(uint64(l))
	}
	if m.LogonAcceptor != nil {
		l = m.LogonAcceptor.Size()
		n += 1 + l + sovSessions(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovSessions(uint64(l))
	}
	if m.IsAccepted {
		n += 2
	}
	return n
}

func (m *LogonInitiator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovSessions(uint64(l))
	}
	if m.EncryptMethod != 0 {
		n += 1 + sovSessions(uint64(m.EncryptMethod))
	}
	if m.HeartBtInt != 0 {
		n += 1 + sovSessions(uint64(m.HeartBtInt))
	}
	if m.Trailer != nil {
		l = m.Trailer.Size()
		n += 1 + l + sovSessions(uint64(l))
	}
	return n
}

func (m *LogonAcceptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovSessions(uint64(l))
	}
	if m.EncryptMethod != 0 {
		n += 1 + sovSessions(uint64(m.EncryptMethod))
	}
	if m.HeartBtInt != 0 {
		n += 1 + sovSessions(uint64(m.HeartBtInt))
	}
	if m.Trailer != nil {
		l = m.Trailer.Size()
		n += 1 + l + sovSessions(uint64(l))
	}
	return n
}

func sovSessions(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSessions(x uint64) (n int) {
	return sovSessions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Sessions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSessions
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
			return fmt.Errorf("proto: Sessions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sessions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessions
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
				return ErrInvalidLengthSessions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSessions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogonInitiator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessions
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
				return ErrInvalidLengthSessions
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSessions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LogonInitiator == nil {
				m.LogonInitiator = &LogonInitiator{}
			}
			if err := m.LogonInitiator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogonAcceptor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessions
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
				return ErrInvalidLengthSessions
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSessions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LogonAcceptor == nil {
				m.LogonAcceptor = &LogonAcceptor{}
			}
			if err := m.LogonAcceptor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessions
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
				return ErrInvalidLengthSessions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSessions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsAccepted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsAccepted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSessions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSessions
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
func (m *LogonInitiator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSessions
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
			return fmt.Errorf("proto: LogonInitiator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogonInitiator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessions
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
				return ErrInvalidLengthSessions
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSessions
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptMethod", wireType)
			}
			m.EncryptMethod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EncryptMethod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeartBtInt", wireType)
			}
			m.HeartBtInt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HeartBtInt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trailer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessions
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
				return ErrInvalidLengthSessions
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSessions
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
			skippy, err := skipSessions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSessions
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
func (m *LogonAcceptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSessions
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
			return fmt.Errorf("proto: LogonAcceptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogonAcceptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessions
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
				return ErrInvalidLengthSessions
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSessions
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptMethod", wireType)
			}
			m.EncryptMethod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EncryptMethod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeartBtInt", wireType)
			}
			m.HeartBtInt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HeartBtInt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trailer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessions
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
				return ErrInvalidLengthSessions
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSessions
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
			skippy, err := skipSessions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSessions
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
func skipSessions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSessions
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
					return 0, ErrIntOverflowSessions
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
					return 0, ErrIntOverflowSessions
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
				return 0, ErrInvalidLengthSessions
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSessions
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSessions
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSessions        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSessions          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSessions = fmt.Errorf("proto: unexpected end of group")
)
