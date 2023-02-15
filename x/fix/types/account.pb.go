// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: re/fix/account.proto

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

type Account struct {
	Id                 uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyName        string   `protobuf:"bytes,2,opt,name=companyName,proto3" json:"companyName,omitempty"`
	Address            string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	EmailAddress       string   `protobuf:"bytes,4,opt,name=emailAddress,proto3" json:"emailAddress,omitempty"`
	PhoneNumber        int32    `protobuf:"varint,5,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Website            string   `protobuf:"bytes,6,opt,name=website,proto3" json:"website,omitempty"`
	SocialMediaLinks   []string `protobuf:"bytes,7,rep,name=socialMediaLinks,proto3" json:"socialMediaLinks,omitempty"`
	GovernmentIssuedId string   `protobuf:"bytes,8,opt,name=governmentIssuedId,proto3" json:"governmentIssuedId,omitempty"`
	CreatedAt          int32    `protobuf:"varint,9,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Creator            string   `protobuf:"bytes,10,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_b50bb661c38375e3, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Account.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Account) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *Account) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Account) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *Account) GetPhoneNumber() int32 {
	if m != nil {
		return m.PhoneNumber
	}
	return 0
}

func (m *Account) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Account) GetSocialMediaLinks() []string {
	if m != nil {
		return m.SocialMediaLinks
	}
	return nil
}

func (m *Account) GetGovernmentIssuedId() string {
	if m != nil {
		return m.GovernmentIssuedId
	}
	return ""
}

func (m *Account) GetCreatedAt() int32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Account) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Account)(nil), "jim380.re.fix.Account")
}

func init() { proto.RegisterFile("re/fix/account.proto", fileDescriptor_b50bb661c38375e3) }

var fileDescriptor_b50bb661c38375e3 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbd, 0x4e, 0xf3, 0x30,
	0x18, 0x85, 0xeb, 0xf4, 0xef, 0x8b, 0x3f, 0x40, 0xc8, 0x62, 0xf0, 0x80, 0xa2, 0xa8, 0x53, 0xc4,
	0x90, 0x20, 0x75, 0x41, 0x6c, 0x65, 0xab, 0x04, 0x1d, 0x32, 0xb2, 0x39, 0xf1, 0xdb, 0xd6, 0x50,
	0xdb, 0x91, 0xed, 0x40, 0x7b, 0x17, 0x5c, 0x16, 0x6c, 0x1d, 0x19, 0x51, 0x7b, 0x23, 0xa8, 0x4e,
	0x11, 0x45, 0x30, 0x9e, 0xe7, 0x58, 0x8f, 0x8e, 0xfc, 0xe2, 0x33, 0x03, 0xd9, 0x54, 0x2c, 0x33,
	0x56, 0x96, 0xba, 0x56, 0x2e, 0xad, 0x8c, 0x76, 0x9a, 0x1c, 0x3f, 0x08, 0x39, 0xbc, 0xba, 0x4c,
	0x0d, 0xa4, 0x53, 0xb1, 0x1c, 0xbc, 0x05, 0xb8, 0x3f, 0x6a, 0x1e, 0x90, 0x13, 0x1c, 0x08, 0x4e,
	0x51, 0x8c, 0x92, 0x4e, 0x1e, 0x08, 0x4e, 0x62, 0xfc, 0xbf, 0xd4, 0xb2, 0x62, 0x6a, 0x35, 0x61,
	0x12, 0x68, 0x10, 0xa3, 0x24, 0xcc, 0x0f, 0x11, 0xa1, 0xb8, 0xcf, 0x38, 0x37, 0x60, 0x2d, 0x6d,
	0xfb, 0xf6, 0x2b, 0x92, 0x01, 0x3e, 0x02, 0xc9, 0xc4, 0x62, 0xb4, 0xaf, 0x3b, 0xbe, 0xfe, 0xc1,
	0x76, 0xfe, 0x6a, 0xae, 0x15, 0x4c, 0x6a, 0x59, 0x80, 0xa1, 0xdd, 0x18, 0x25, 0xdd, 0xfc, 0x10,
	0xed, 0xfc, 0xcf, 0x50, 0x58, 0xe1, 0x80, 0xf6, 0x1a, 0xff, 0x3e, 0x92, 0x0b, 0x7c, 0x6a, 0x75,
	0x29, 0xd8, 0xe2, 0x0e, 0xb8, 0x60, 0xb7, 0x42, 0x3d, 0x5a, 0xda, 0x8f, 0xdb, 0x49, 0x98, 0xff,
	0xe2, 0x24, 0xc5, 0x64, 0xa6, 0x9f, 0xc0, 0x28, 0x09, 0xca, 0x8d, 0xad, 0xad, 0x81, 0x8f, 0x39,
	0xfd, 0xe7, 0x85, 0x7f, 0x34, 0xe4, 0x1c, 0x87, 0xa5, 0x01, 0xe6, 0x80, 0x8f, 0x1c, 0x0d, 0xfd,
	0xaa, 0x6f, 0xb0, 0xdb, 0xe4, 0x83, 0x36, 0x14, 0x37, 0x9b, 0xf6, 0xf1, 0xe6, 0xfa, 0x75, 0x13,
	0xa1, 0xf5, 0x26, 0x42, 0x1f, 0x9b, 0x08, 0xbd, 0x6c, 0xa3, 0xd6, 0x7a, 0x1b, 0xb5, 0xde, 0xb7,
	0x51, 0xeb, 0x3e, 0x9e, 0x09, 0x37, 0xaf, 0x8b, 0xb4, 0xd4, 0x32, 0x6b, 0xfe, 0x3f, 0xcb, 0x21,
	0x5b, 0xfa, 0xf3, 0xb8, 0x55, 0x05, 0xb6, 0xe8, 0xf9, 0xeb, 0x0c, 0x3f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x16, 0x0b, 0x0a, 0xbf, 0xb5, 0x01, 0x00, 0x00,
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x52
	}
	if m.CreatedAt != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x48
	}
	if len(m.GovernmentIssuedId) > 0 {
		i -= len(m.GovernmentIssuedId)
		copy(dAtA[i:], m.GovernmentIssuedId)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.GovernmentIssuedId)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.SocialMediaLinks) > 0 {
		for iNdEx := len(m.SocialMediaLinks) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SocialMediaLinks[iNdEx])
			copy(dAtA[i:], m.SocialMediaLinks[iNdEx])
			i = encodeVarintAccount(dAtA, i, uint64(len(m.SocialMediaLinks[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x32
	}
	if m.PhoneNumber != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.PhoneNumber))
		i--
		dAtA[i] = 0x28
	}
	if len(m.EmailAddress) > 0 {
		i -= len(m.EmailAddress)
		copy(dAtA[i:], m.EmailAddress)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.EmailAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CompanyName) > 0 {
		i -= len(m.CompanyName)
		copy(dAtA[i:], m.CompanyName)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.CompanyName)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAccount(uint64(m.Id))
	}
	l = len(m.CompanyName)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.EmailAddress)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	if m.PhoneNumber != 0 {
		n += 1 + sovAccount(uint64(m.PhoneNumber))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	if len(m.SocialMediaLinks) > 0 {
		for _, s := range m.SocialMediaLinks {
			l = len(s)
			n += 1 + l + sovAccount(uint64(l))
		}
	}
	l = len(m.GovernmentIssuedId)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovAccount(uint64(m.CreatedAt))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	return n
}

func sovAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccount(x uint64) (n int) {
	return sovAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Account) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
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
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return fmt.Errorf("proto: wrong wireType = %d for field CompanyName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompanyName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmailAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EmailAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PhoneNumber", wireType)
			}
			m.PhoneNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PhoneNumber |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SocialMediaLinks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SocialMediaLinks = append(m.SocialMediaLinks, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GovernmentIssuedId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GovernmentIssuedId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
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
func skipAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
				return 0, ErrInvalidLengthAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccount = fmt.Errorf("proto: unexpected end of group")
)
