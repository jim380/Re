// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: re/fix/account_registration.proto

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

type AccountRegistration struct {
	Address          string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	CompanyName      string `protobuf:"bytes,2,opt,name=companyName,proto3" json:"companyName,omitempty"`
	Website          string `protobuf:"bytes,3,opt,name=website,proto3" json:"website,omitempty"`
	SocialMediaLinks string `protobuf:"bytes,4,opt,name=socialMediaLinks,proto3" json:"socialMediaLinks,omitempty"`
	CreatedAt        string `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (m *AccountRegistration) Reset()         { *m = AccountRegistration{} }
func (m *AccountRegistration) String() string { return proto.CompactTextString(m) }
func (*AccountRegistration) ProtoMessage()    {}
func (*AccountRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1216fa4b32a656c, []int{0}
}
func (m *AccountRegistration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountRegistration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRegistration.Merge(m, src)
}
func (m *AccountRegistration) XXX_Size() int {
	return m.Size()
}
func (m *AccountRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRegistration proto.InternalMessageInfo

func (m *AccountRegistration) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AccountRegistration) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *AccountRegistration) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *AccountRegistration) GetSocialMediaLinks() string {
	if m != nil {
		return m.SocialMediaLinks
	}
	return ""
}

func (m *AccountRegistration) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*AccountRegistration)(nil), "jim380.re.fix.AccountRegistration")
}

func init() { proto.RegisterFile("re/fix/account_registration.proto", fileDescriptor_d1216fa4b32a656c) }

var fileDescriptor_d1216fa4b32a656c = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x4a, 0xd5, 0x4f,
	0xcb, 0xac, 0xd0, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x89, 0x2f, 0x4a, 0x4d, 0xcf, 0x2c,
	0x2e, 0x29, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcd,
	0xca, 0xcc, 0x35, 0xb6, 0x30, 0xd0, 0x2b, 0x4a, 0xd5, 0x4b, 0xcb, 0xac, 0x50, 0xda, 0xca, 0xc8,
	0x25, 0xec, 0x08, 0x51, 0x1d, 0x84, 0xa4, 0x58, 0x48, 0x82, 0x8b, 0x3d, 0x31, 0x25, 0xa5, 0x28,
	0xb5, 0xb8, 0x58, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6, 0x15, 0x52, 0xe0, 0xe2, 0x4e,
	0xce, 0xcf, 0x2d, 0x48, 0xcc, 0xab, 0xf4, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x02, 0xcb, 0x22, 0x0b,
	0x81, 0xf4, 0x96, 0xa7, 0x26, 0x15, 0x67, 0x96, 0xa4, 0x4a, 0x30, 0x43, 0xf4, 0x42, 0xb9, 0x42,
	0x5a, 0x5c, 0x02, 0xc5, 0xf9, 0xc9, 0x99, 0x89, 0x39, 0xbe, 0xa9, 0x29, 0x99, 0x89, 0x3e, 0x99,
	0x79, 0xd9, 0xc5, 0x12, 0x2c, 0x60, 0x25, 0x18, 0xe2, 0x42, 0x32, 0x5c, 0x9c, 0xc9, 0x45, 0xa9,
	0x89, 0x25, 0xa9, 0x29, 0x8e, 0x25, 0x12, 0xac, 0x60, 0x45, 0x08, 0x01, 0x27, 0xab, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b,
	0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x52, 0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0x87, 0xf8, 0x55, 0x3f, 0x28, 0x55, 0xbf, 0x02, 0x1c, 0x2e, 0x25, 0x95,
	0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0x90, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x59, 0x77,
	0xfa, 0xe0, 0x2e, 0x01, 0x00, 0x00,
}

func (m *AccountRegistration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountRegistration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountRegistration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CreatedAt) > 0 {
		i -= len(m.CreatedAt)
		copy(dAtA[i:], m.CreatedAt)
		i = encodeVarintAccountRegistration(dAtA, i, uint64(len(m.CreatedAt)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SocialMediaLinks) > 0 {
		i -= len(m.SocialMediaLinks)
		copy(dAtA[i:], m.SocialMediaLinks)
		i = encodeVarintAccountRegistration(dAtA, i, uint64(len(m.SocialMediaLinks)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintAccountRegistration(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CompanyName) > 0 {
		i -= len(m.CompanyName)
		copy(dAtA[i:], m.CompanyName)
		i = encodeVarintAccountRegistration(dAtA, i, uint64(len(m.CompanyName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAccountRegistration(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccountRegistration(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccountRegistration(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccountRegistration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAccountRegistration(uint64(l))
	}
	l = len(m.CompanyName)
	if l > 0 {
		n += 1 + l + sovAccountRegistration(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovAccountRegistration(uint64(l))
	}
	l = len(m.SocialMediaLinks)
	if l > 0 {
		n += 1 + l + sovAccountRegistration(uint64(l))
	}
	l = len(m.CreatedAt)
	if l > 0 {
		n += 1 + l + sovAccountRegistration(uint64(l))
	}
	return n
}

func sovAccountRegistration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccountRegistration(x uint64) (n int) {
	return sovAccountRegistration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountRegistration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccountRegistration
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
			return fmt.Errorf("proto: AccountRegistration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountRegistration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountRegistration
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
				return ErrInvalidLengthAccountRegistration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompanyName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountRegistration
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
				return ErrInvalidLengthAccountRegistration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompanyName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountRegistration
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
				return ErrInvalidLengthAccountRegistration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SocialMediaLinks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountRegistration
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
				return ErrInvalidLengthAccountRegistration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SocialMediaLinks = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountRegistration
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
				return ErrInvalidLengthAccountRegistration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccountRegistration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccountRegistration
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
func skipAccountRegistration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccountRegistration
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
					return 0, ErrIntOverflowAccountRegistration
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
					return 0, ErrIntOverflowAccountRegistration
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
				return 0, ErrInvalidLengthAccountRegistration
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccountRegistration
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccountRegistration
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccountRegistration        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccountRegistration          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccountRegistration = fmt.Errorf("proto: unexpected end of group")
)