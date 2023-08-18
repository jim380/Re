// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jim380/re/fix/standard_message.proto

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

// This message header contains fields that identify the message being sent and
// its sender and recipient
type Header struct {
	// This field specifies the version of the FIX protocol being used. For
	// example, "FIX.4.4" indicates version 4.4 of the protocol.
	BeginString string `protobuf:"bytes,1,opt,name=beginString,proto3" json:"beginString,omitempty"`
	// This field specifies the length of the message body in bytes.
	BodyLength int64 `protobuf:"varint,2,opt,name=bodyLength,proto3" json:"bodyLength,omitempty"`
	// This field identifies the type of message being sent. For example, "D"
	// indicates a New Order Single message
	MsgType string `protobuf:"bytes,3,opt,name=msgType,proto3" json:"msgType,omitempty"`
	// This field identifies the sender of the message. This is typically a unique
	// identifier assigned to each party that uses the FIX protocol
	SenderCompID string `protobuf:"bytes,4,opt,name=senderCompID,proto3" json:"senderCompID,omitempty"`
	// This field identifies the recipient of the message.
	TargetCompID string `protobuf:"bytes,5,opt,name=targetCompID,proto3" json:"targetCompID,omitempty"`
	// Is a field in the FIX protocol message header that is used to assign a
	// unique sequence number to each message
	MsgSeqNum int64 `protobuf:"varint,6,opt,name=msgSeqNum,proto3" json:"msgSeqNum,omitempty"`
	// This field specifies the time that the message was sent
	SendingTime string `protobuf:"bytes,7,opt,name=sendingTime,proto3" json:"sendingTime,omitempty"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ca24882ac2cd552, []int{0}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Header.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return m.Size()
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetBeginString() string {
	if m != nil {
		return m.BeginString
	}
	return ""
}

func (m *Header) GetBodyLength() int64 {
	if m != nil {
		return m.BodyLength
	}
	return 0
}

func (m *Header) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *Header) GetSenderCompID() string {
	if m != nil {
		return m.SenderCompID
	}
	return ""
}

func (m *Header) GetTargetCompID() string {
	if m != nil {
		return m.TargetCompID
	}
	return ""
}

func (m *Header) GetMsgSeqNum() int64 {
	if m != nil {
		return m.MsgSeqNum
	}
	return 0
}

func (m *Header) GetSendingTime() string {
	if m != nil {
		return m.SendingTime
	}
	return ""
}

// This message trailer contains a checksum field that is used to ensure message
// integrity
type Trailer struct {
	// The checksum value is calculated by summing the ASCII values of all
	// characters in the message (excluding the SOH (start of header) and ETX (end
	// of text) characters), taking the sum modulo 256, and converting the result
	// to a two-character hexadecimal string.
	CheckSum int64 `protobuf:"varint,1,opt,name=checkSum,proto3" json:"checkSum,omitempty"`
}

func (m *Trailer) Reset()         { *m = Trailer{} }
func (m *Trailer) String() string { return proto.CompactTextString(m) }
func (*Trailer) ProtoMessage()    {}
func (*Trailer) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ca24882ac2cd552, []int{1}
}
func (m *Trailer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Trailer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Trailer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Trailer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trailer.Merge(m, src)
}
func (m *Trailer) XXX_Size() int {
	return m.Size()
}
func (m *Trailer) XXX_DiscardUnknown() {
	xxx_messageInfo_Trailer.DiscardUnknown(m)
}

var xxx_messageInfo_Trailer proto.InternalMessageInfo

func (m *Trailer) GetCheckSum() int64 {
	if m != nil {
		return m.CheckSum
	}
	return 0
}

func init() {
	proto.RegisterType((*Header)(nil), "jim380.re.fix.Header")
	proto.RegisterType((*Trailer)(nil), "jim380.re.fix.Trailer")
}

func init() {
	proto.RegisterFile("jim380/re/fix/standard_message.proto", fileDescriptor_5ca24882ac2cd552)
}

var fileDescriptor_5ca24882ac2cd552 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0xeb, 0x30,
	0x18, 0x85, 0xeb, 0xdb, 0x4b, 0x4b, 0x0d, 0x2c, 0x9e, 0x2c, 0x84, 0xac, 0xa8, 0x02, 0xa9, 0x53,
	0x82, 0xd4, 0x05, 0x31, 0x02, 0x03, 0x48, 0x88, 0x21, 0xed, 0xc4, 0x82, 0x9c, 0xe6, 0xaf, 0x6b,
	0xc0, 0x4e, 0xb0, 0x5d, 0x29, 0x79, 0x0b, 0x1e, 0x8b, 0xb1, 0x23, 0x23, 0x4a, 0x56, 0x1e, 0x02,
	0xc5, 0x29, 0x10, 0xc6, 0xf3, 0xf9, 0xb3, 0xe5, 0x73, 0xf0, 0xf1, 0xa3, 0x54, 0xd3, 0xb3, 0xd3,
	0xc8, 0x40, 0xb4, 0x94, 0x45, 0x64, 0x1d, 0xd7, 0x29, 0x37, 0xe9, 0x83, 0x02, 0x6b, 0xb9, 0x80,
	0x30, 0x37, 0x99, 0xcb, 0xc8, 0x41, 0x6b, 0x85, 0x06, 0xc2, 0xa5, 0x2c, 0xc6, 0x9f, 0x08, 0x0f,
	0xae, 0x81, 0xa7, 0x60, 0x48, 0x80, 0xf7, 0x12, 0x10, 0x52, 0xcf, 0x9c, 0x91, 0x5a, 0x50, 0x14,
	0xa0, 0xc9, 0x28, 0xee, 0x22, 0xc2, 0x30, 0x4e, 0xb2, 0xb4, 0xbc, 0x05, 0x2d, 0xdc, 0x8a, 0xfe,
	0x0b, 0xd0, 0xa4, 0x1f, 0x77, 0x08, 0xa1, 0x78, 0xa8, 0xac, 0x98, 0x97, 0x39, 0xd0, 0xbe, 0xbf,
	0xfd, 0x1d, 0xc9, 0x18, 0xef, 0x5b, 0xd0, 0x29, 0x98, 0xcb, 0x4c, 0xe5, 0x37, 0x57, 0xf4, 0xbf,
	0x3f, 0xfe, 0xc3, 0x1a, 0xc7, 0x71, 0x23, 0xc0, 0x6d, 0x9d, 0x9d, 0xd6, 0xe9, 0x32, 0x72, 0x84,
	0x47, 0xca, 0x8a, 0x19, 0xbc, 0xdc, 0xad, 0x15, 0x1d, 0xf8, 0x0f, 0xfc, 0x82, 0xa6, 0x41, 0xf3,
	0xa2, 0xd4, 0x62, 0x2e, 0x15, 0xd0, 0x61, 0xdb, 0xa0, 0x83, 0xc6, 0x27, 0x78, 0x38, 0x37, 0x5c,
	0x3e, 0x83, 0x21, 0x87, 0x78, 0x77, 0xb1, 0x82, 0xc5, 0xd3, 0x6c, 0xad, 0x7c, 0xd7, 0x7e, 0xfc,
	0x93, 0x2f, 0xce, 0xdf, 0x2a, 0x86, 0x36, 0x15, 0x43, 0x1f, 0x15, 0x43, 0xaf, 0x35, 0xeb, 0x6d,
	0x6a, 0xd6, 0x7b, 0xaf, 0x59, 0xef, 0x3e, 0x10, 0xd2, 0xad, 0xd6, 0x49, 0xb8, 0xc8, 0x54, 0xb4,
	0xdd, 0x3b, 0x86, 0xa8, 0xf0, 0x8b, 0xbb, 0x32, 0x07, 0x9b, 0x0c, 0xfc, 0xce, 0xd3, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd1, 0x88, 0x1d, 0x33, 0x8f, 0x01, 0x00, 0x00,
}

func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Header) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SendingTime) > 0 {
		i -= len(m.SendingTime)
		copy(dAtA[i:], m.SendingTime)
		i = encodeVarintStandardMessage(dAtA, i, uint64(len(m.SendingTime)))
		i--
		dAtA[i] = 0x3a
	}
	if m.MsgSeqNum != 0 {
		i = encodeVarintStandardMessage(dAtA, i, uint64(m.MsgSeqNum))
		i--
		dAtA[i] = 0x30
	}
	if len(m.TargetCompID) > 0 {
		i -= len(m.TargetCompID)
		copy(dAtA[i:], m.TargetCompID)
		i = encodeVarintStandardMessage(dAtA, i, uint64(len(m.TargetCompID)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SenderCompID) > 0 {
		i -= len(m.SenderCompID)
		copy(dAtA[i:], m.SenderCompID)
		i = encodeVarintStandardMessage(dAtA, i, uint64(len(m.SenderCompID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.MsgType) > 0 {
		i -= len(m.MsgType)
		copy(dAtA[i:], m.MsgType)
		i = encodeVarintStandardMessage(dAtA, i, uint64(len(m.MsgType)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BodyLength != 0 {
		i = encodeVarintStandardMessage(dAtA, i, uint64(m.BodyLength))
		i--
		dAtA[i] = 0x10
	}
	if len(m.BeginString) > 0 {
		i -= len(m.BeginString)
		copy(dAtA[i:], m.BeginString)
		i = encodeVarintStandardMessage(dAtA, i, uint64(len(m.BeginString)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Trailer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Trailer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Trailer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CheckSum != 0 {
		i = encodeVarintStandardMessage(dAtA, i, uint64(m.CheckSum))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStandardMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovStandardMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BeginString)
	if l > 0 {
		n += 1 + l + sovStandardMessage(uint64(l))
	}
	if m.BodyLength != 0 {
		n += 1 + sovStandardMessage(uint64(m.BodyLength))
	}
	l = len(m.MsgType)
	if l > 0 {
		n += 1 + l + sovStandardMessage(uint64(l))
	}
	l = len(m.SenderCompID)
	if l > 0 {
		n += 1 + l + sovStandardMessage(uint64(l))
	}
	l = len(m.TargetCompID)
	if l > 0 {
		n += 1 + l + sovStandardMessage(uint64(l))
	}
	if m.MsgSeqNum != 0 {
		n += 1 + sovStandardMessage(uint64(m.MsgSeqNum))
	}
	l = len(m.SendingTime)
	if l > 0 {
		n += 1 + l + sovStandardMessage(uint64(l))
	}
	return n
}

func (m *Trailer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CheckSum != 0 {
		n += 1 + sovStandardMessage(uint64(m.CheckSum))
	}
	return n
}

func sovStandardMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStandardMessage(x uint64) (n int) {
	return sovStandardMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStandardMessage
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
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStandardMessage
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
				return ErrInvalidLengthStandardMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStandardMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BeginString = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BodyLength", wireType)
			}
			m.BodyLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStandardMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BodyLength |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStandardMessage
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
				return ErrInvalidLengthStandardMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStandardMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderCompID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStandardMessage
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
				return ErrInvalidLengthStandardMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStandardMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SenderCompID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetCompID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStandardMessage
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
				return ErrInvalidLengthStandardMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStandardMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetCompID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgSeqNum", wireType)
			}
			m.MsgSeqNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStandardMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgSeqNum |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendingTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStandardMessage
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
				return ErrInvalidLengthStandardMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStandardMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SendingTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStandardMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStandardMessage
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
func (m *Trailer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStandardMessage
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
			return fmt.Errorf("proto: Trailer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Trailer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckSum", wireType)
			}
			m.CheckSum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStandardMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CheckSum |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStandardMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStandardMessage
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
func skipStandardMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStandardMessage
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
					return 0, ErrIntOverflowStandardMessage
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
					return 0, ErrIntOverflowStandardMessage
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
				return 0, ErrInvalidLengthStandardMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStandardMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStandardMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStandardMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStandardMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStandardMessage = fmt.Errorf("proto: unexpected end of group")
)
