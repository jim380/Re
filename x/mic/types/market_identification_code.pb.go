// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: re/mic/market_identification_code.proto

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

// Definition of MarketIdentificationCode message
type MarketIdentificationCode struct {
	// Market Identifier Code allocated to the market named in ‘Market
	// Name-Institution Description’.
	MIC string `protobuf:"bytes,1,opt,name=MIC,proto3" json:"MIC,omitempty"`
	// Entity operating an exchange/market/trade reporting facility in a specific
	// market/country.
	Operating_MIC string `protobuf:"bytes,2,opt,name=operating_MIC,json=operatingMIC,proto3" json:"operating_MIC,omitempty"`
	// Indicates whether the MIC is an operating MIC or a market segment MIC.
	OPRT_SGMT string `protobuf:"bytes,3,opt,name=OPRT_SGMT,json=OPRTSGMT,proto3" json:"OPRT_SGMT,omitempty"`
	// Institution Description: name of the market.
	MarketName string `protobuf:"bytes,4,opt,name=market_name,json=marketName,proto3" json:"market_name,omitempty"`
	// Legal name of the entity owning the market.
	LegalEntityName string `protobuf:"bytes,5,opt,name=legal_entity_name,json=legalEntityName,proto3" json:"legal_entity_name,omitempty"`
	// Legal Entity Identifier (LEI) see ISO 17442-1.
	LegalEntityIdentifier string `protobuf:"bytes,6,opt,name=legal_entity_identifier,json=legalEntityIdentifier,proto3" json:"legal_entity_identifier,omitempty"`
	// Specifies the type of market. The list of market types is predefined (1).
	// The list can be updated upon request to the RA, which will validate the
	// request.
	MarketCategory string `protobuf:"bytes,7,opt,name=market_category,json=marketCategory,proto3" json:"market_category,omitempty"`
	// Known acronym of the market.
	Acronym string `protobuf:"bytes,8,opt,name=acronym,proto3" json:"acronym,omitempty"`
	// Alpha-2 code of the country where the market is registered.
	ISOCountryCode string `protobuf:"bytes,9,opt,name=ISO_country_code,json=ISOCountryCode,proto3" json:"ISO_country_code,omitempty"`
	// City where the market is located.
	City string `protobuf:"bytes,10,opt,name=city,proto3" json:"city,omitempty"`
	// Website of the market.
	Website string `protobuf:"bytes,11,opt,name=website,proto3" json:"website,omitempty"`
	// Active, updated (since last publication), expired (= deactivated).
	Status string `protobuf:"bytes,12,opt,name=status,proto3" json:"status,omitempty"`
	// Date indicating when the MIC was originally created.
	CreationDate string `protobuf:"bytes,13,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	// Date indicating when the MIC was last modified.
	LastUpdateDate string `protobuf:"bytes,14,opt,name=last_update_date,json=lastUpdateDate,proto3" json:"last_update_date,omitempty"`
	// Date indicating when the MIC was last reviewed for correctness.
	LastValidationDate string `protobuf:"bytes,15,opt,name=last_validation_date,json=lastValidationDate,proto3" json:"last_validation_date,omitempty"`
	// The expiry date is populated when the MIC is deactivated; upon request from
	// the MIC owner; following market research (user request) or maintenance. The
	// expiry date field is left blank when a MIC is created.
	ExpiryDate string `protobuf:"bytes,16,opt,name=expiry_date,json=expiryDate,proto3" json:"expiry_date,omitempty"`
	// Any additional information worth mentioning to help users with identifying
	// the exchange or understanding a modification.
	Comments string `protobuf:"bytes,17,opt,name=comments,proto3" json:"comments,omitempty"`
	// Address for MIC creator.
	Creator string `protobuf:"bytes,18,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MarketIdentificationCode) Reset()         { *m = MarketIdentificationCode{} }
func (m *MarketIdentificationCode) String() string { return proto.CompactTextString(m) }
func (*MarketIdentificationCode) ProtoMessage()    {}
func (*MarketIdentificationCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdccc5d6811e94e3, []int{0}
}
func (m *MarketIdentificationCode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketIdentificationCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketIdentificationCode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketIdentificationCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketIdentificationCode.Merge(m, src)
}
func (m *MarketIdentificationCode) XXX_Size() int {
	return m.Size()
}
func (m *MarketIdentificationCode) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketIdentificationCode.DiscardUnknown(m)
}

var xxx_messageInfo_MarketIdentificationCode proto.InternalMessageInfo

func (m *MarketIdentificationCode) GetMIC() string {
	if m != nil {
		return m.MIC
	}
	return ""
}

func (m *MarketIdentificationCode) GetOperating_MIC() string {
	if m != nil {
		return m.Operating_MIC
	}
	return ""
}

func (m *MarketIdentificationCode) GetOPRT_SGMT() string {
	if m != nil {
		return m.OPRT_SGMT
	}
	return ""
}

func (m *MarketIdentificationCode) GetMarketName() string {
	if m != nil {
		return m.MarketName
	}
	return ""
}

func (m *MarketIdentificationCode) GetLegalEntityName() string {
	if m != nil {
		return m.LegalEntityName
	}
	return ""
}

func (m *MarketIdentificationCode) GetLegalEntityIdentifier() string {
	if m != nil {
		return m.LegalEntityIdentifier
	}
	return ""
}

func (m *MarketIdentificationCode) GetMarketCategory() string {
	if m != nil {
		return m.MarketCategory
	}
	return ""
}

func (m *MarketIdentificationCode) GetAcronym() string {
	if m != nil {
		return m.Acronym
	}
	return ""
}

func (m *MarketIdentificationCode) GetISOCountryCode() string {
	if m != nil {
		return m.ISOCountryCode
	}
	return ""
}

func (m *MarketIdentificationCode) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *MarketIdentificationCode) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *MarketIdentificationCode) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *MarketIdentificationCode) GetCreationDate() string {
	if m != nil {
		return m.CreationDate
	}
	return ""
}

func (m *MarketIdentificationCode) GetLastUpdateDate() string {
	if m != nil {
		return m.LastUpdateDate
	}
	return ""
}

func (m *MarketIdentificationCode) GetLastValidationDate() string {
	if m != nil {
		return m.LastValidationDate
	}
	return ""
}

func (m *MarketIdentificationCode) GetExpiryDate() string {
	if m != nil {
		return m.ExpiryDate
	}
	return ""
}

func (m *MarketIdentificationCode) GetComments() string {
	if m != nil {
		return m.Comments
	}
	return ""
}

func (m *MarketIdentificationCode) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*MarketIdentificationCode)(nil), "re.mic.MarketIdentificationCode")
}

func init() {
	proto.RegisterFile("re/mic/market_identification_code.proto", fileDescriptor_fdccc5d6811e94e3)
}

var fileDescriptor_fdccc5d6811e94e3 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x4f, 0x6f, 0x13, 0x3d,
	0x10, 0xc6, 0x93, 0xb7, 0x79, 0xd3, 0xc4, 0x69, 0x9b, 0xd4, 0xe2, 0x8f, 0x05, 0xd2, 0xb6, 0x82,
	0x43, 0x2b, 0x0e, 0xd9, 0x4a, 0x95, 0x10, 0xe2, 0x48, 0x40, 0x68, 0x0f, 0x21, 0x28, 0x29, 0x1c,
	0xb8, 0xac, 0x1c, 0xef, 0x10, 0x0c, 0xf1, 0x7a, 0xe5, 0x9d, 0x40, 0xf7, 0x5b, 0xf0, 0xb1, 0x38,
	0xf6, 0xc8, 0x11, 0x25, 0xdf, 0x03, 0x21, 0x8f, 0x37, 0x51, 0x7b, 0xf3, 0x3c, 0xcf, 0xcf, 0xa3,
	0xc7, 0x33, 0x66, 0x67, 0x0e, 0x62, 0xa3, 0x55, 0x6c, 0xa4, 0xfb, 0x06, 0x98, 0xea, 0x0c, 0x72,
	0xd4, 0x9f, 0xb5, 0x92, 0xa8, 0x6d, 0x9e, 0x2a, 0x9b, 0xc1, 0xb0, 0x70, 0x16, 0x2d, 0x6f, 0x3b,
	0x18, 0x1a, 0xad, 0x9e, 0xfc, 0x6d, 0x31, 0x31, 0x26, 0x38, 0xb9, 0xc3, 0x8e, 0x6c, 0x06, 0x7c,
	0xc0, 0xf6, 0xc6, 0xc9, 0x48, 0x34, 0x4f, 0x9b, 0xe7, 0xdd, 0xa9, 0x3f, 0xf2, 0xa7, 0xec, 0xd0,
	0x16, 0xe0, 0x24, 0xea, 0x7c, 0x91, 0x7a, 0xef, 0x3f, 0xf2, 0x0e, 0x76, 0xa2, 0x87, 0x1e, 0xb3,
	0xee, 0xe4, 0xfd, 0xf4, 0x2a, 0x9d, 0xbd, 0x1d, 0x5f, 0x89, 0x3d, 0x02, 0x3a, 0x5e, 0xf0, 0x35,
	0x3f, 0x61, 0xbd, 0x3a, 0x5c, 0x2e, 0x0d, 0x88, 0x16, 0xd9, 0x2c, 0x48, 0xef, 0xa4, 0x01, 0xfe,
	0x8c, 0x1d, 0x2f, 0x61, 0x21, 0x97, 0xa9, 0x8f, 0x83, 0x55, 0xc0, 0xfe, 0x27, 0xac, 0x4f, 0xc6,
	0x1b, 0xd2, 0x89, 0x7d, 0xce, 0x1e, 0xde, 0x61, 0xb7, 0xef, 0x05, 0x27, 0xda, 0x74, 0xe3, 0xfe,
	0xad, 0x1b, 0xc9, 0xce, 0xe4, 0x67, 0xac, 0x5f, 0x87, 0x50, 0x12, 0x61, 0x61, 0x5d, 0x25, 0xf6,
	0x89, 0x3f, 0x0a, 0xf2, 0xa8, 0x56, 0xb9, 0x60, 0xfb, 0x52, 0x39, 0x9b, 0x57, 0x46, 0x74, 0x08,
	0xd8, 0x96, 0xfc, 0x9c, 0x0d, 0x92, 0xd9, 0x24, 0x55, 0x76, 0x95, 0xa3, 0xab, 0x68, 0xb4, 0xa2,
	0x1b, 0x7a, 0x24, 0xb3, 0xc9, 0x28, 0xc8, 0x34, 0x45, 0xce, 0x5a, 0x4a, 0x63, 0x25, 0x18, 0xb9,
	0x74, 0xf6, 0x7d, 0x7f, 0xc0, 0xbc, 0xd4, 0x08, 0xa2, 0x17, 0xfa, 0xd6, 0x25, 0x7f, 0xc0, 0xda,
	0x25, 0x4a, 0x5c, 0x95, 0xe2, 0x80, 0x8c, 0xba, 0xf2, 0x93, 0x57, 0x0e, 0xc2, 0x1e, 0x33, 0x89,
	0x20, 0x0e, 0xc3, 0xe4, 0xb7, 0xe2, 0x6b, 0x89, 0xe0, 0x43, 0x2d, 0x65, 0x89, 0xe9, 0xaa, 0xf0,
	0x48, 0xe0, 0x8e, 0x42, 0x28, 0xaf, 0x7f, 0x20, 0x99, 0xc8, 0x0b, 0x76, 0x8f, 0xc8, 0xef, 0x72,
	0xa9, 0xb3, 0x5b, 0x5d, 0xfb, 0x44, 0x73, 0xef, 0x7d, 0xdc, 0x59, 0x74, 0xe3, 0x84, 0xf5, 0xe0,
	0xba, 0xd0, 0xae, 0x0a, 0xe0, 0x20, 0x2c, 0x2e, 0x48, 0x04, 0x3c, 0x62, 0x1d, 0x65, 0x8d, 0x81,
	0x1c, 0x4b, 0x71, 0x1c, 0xb6, 0xbe, 0xad, 0xfd, 0x7b, 0x29, 0xa8, 0x75, 0x82, 0x87, 0xf7, 0xd6,
	0xe5, 0xab, 0x97, 0xbf, 0xd6, 0x51, 0xf3, 0x66, 0x1d, 0x35, 0xff, 0xac, 0xa3, 0xe6, 0xcf, 0x4d,
	0xd4, 0xb8, 0xd9, 0x44, 0x8d, 0xdf, 0x9b, 0xa8, 0xf1, 0xe9, 0x74, 0xa1, 0xf1, 0xcb, 0x6a, 0x3e,
	0x54, 0xd6, 0xc4, 0x5f, 0xb5, 0xb9, 0x7c, 0x71, 0x11, 0x4f, 0x21, 0xbe, 0xa6, 0xff, 0x8d, 0x55,
	0x01, 0xe5, 0xbc, 0x4d, 0x7f, 0xf9, 0xf2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75, 0x93, 0xdb,
	0x9a, 0xf6, 0x02, 0x00, 0x00,
}

func (m *MarketIdentificationCode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketIdentificationCode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketIdentificationCode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x92
	}
	if len(m.Comments) > 0 {
		i -= len(m.Comments)
		copy(dAtA[i:], m.Comments)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.Comments)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if len(m.ExpiryDate) > 0 {
		i -= len(m.ExpiryDate)
		copy(dAtA[i:], m.ExpiryDate)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.ExpiryDate)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if len(m.LastValidationDate) > 0 {
		i -= len(m.LastValidationDate)
		copy(dAtA[i:], m.LastValidationDate)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.LastValidationDate)))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.LastUpdateDate) > 0 {
		i -= len(m.LastUpdateDate)
		copy(dAtA[i:], m.LastUpdateDate)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.LastUpdateDate)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.CreationDate) > 0 {
		i -= len(m.CreationDate)
		copy(dAtA[i:], m.CreationDate)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.CreationDate)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.City) > 0 {
		i -= len(m.City)
		copy(dAtA[i:], m.City)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.City)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.ISOCountryCode) > 0 {
		i -= len(m.ISOCountryCode)
		copy(dAtA[i:], m.ISOCountryCode)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.ISOCountryCode)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Acronym) > 0 {
		i -= len(m.Acronym)
		copy(dAtA[i:], m.Acronym)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.Acronym)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.MarketCategory) > 0 {
		i -= len(m.MarketCategory)
		copy(dAtA[i:], m.MarketCategory)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.MarketCategory)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.LegalEntityIdentifier) > 0 {
		i -= len(m.LegalEntityIdentifier)
		copy(dAtA[i:], m.LegalEntityIdentifier)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.LegalEntityIdentifier)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.LegalEntityName) > 0 {
		i -= len(m.LegalEntityName)
		copy(dAtA[i:], m.LegalEntityName)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.LegalEntityName)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.MarketName) > 0 {
		i -= len(m.MarketName)
		copy(dAtA[i:], m.MarketName)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.MarketName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OPRT_SGMT) > 0 {
		i -= len(m.OPRT_SGMT)
		copy(dAtA[i:], m.OPRT_SGMT)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.OPRT_SGMT)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Operating_MIC) > 0 {
		i -= len(m.Operating_MIC)
		copy(dAtA[i:], m.Operating_MIC)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.Operating_MIC)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MIC) > 0 {
		i -= len(m.MIC)
		copy(dAtA[i:], m.MIC)
		i = encodeVarintMarketIdentificationCode(dAtA, i, uint64(len(m.MIC)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMarketIdentificationCode(dAtA []byte, offset int, v uint64) int {
	offset -= sovMarketIdentificationCode(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MarketIdentificationCode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MIC)
	if l > 0 {
		n += 1 + l + sovMarketIdentificationCode(uint64(l))
	}
	l = len(m.Operating_MIC)
	if l > 0 {
		n += 1 + l + sovMarketIdentificationCode(uint64(l))
	}
	l = len(m.OPRT_SGMT)
	if l > 0 {
		n += 1 + l + sovMarketIdentificationCode(uint64(l))
	}
	l = len(m.MarketName)
	if l > 0 {
		n += 1 + l + sovMarketIdentificationCode(uint64(l))
	}
	l = len(m.LegalEntityName)
	if l > 0 {
		n += 1 + l + sovMarketIdentificationCode(uint64(l))
	}
	l = len(m.LegalEntityIdentifier)
	if l > 0 {
		n += 1 + l + sovMarketIdentificationCode(uint64(l))
	}
	l = len(m.MarketCategory)
	if l > 0 {
		n += 1 + l + sovMarketIdentificationCode(uint64(l))
	}
	l = len(m.Acronym)
	if l > 0 {
		n += 1 + l + sovMarketIdentificationCode(uint64(l))
	}
	l = len(m.ISOCountryCode)
	if l > 0 {
		n += 1 + l + sovMarketIdentificationCode(uint64(l))
	}
	l = len(m.City)
	if l > 0 {
		n += 1 + l + sovMarketIdentificationCode(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovMarketIdentificationCode(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovMarketIdentificationCode(uint64(l))
	}
	l = len(m.CreationDate)
	if l > 0 {
		n += 1 + l + sovMarketIdentificationCode(uint64(l))
	}
	l = len(m.LastUpdateDate)
	if l > 0 {
		n += 1 + l + sovMarketIdentificationCode(uint64(l))
	}
	l = len(m.LastValidationDate)
	if l > 0 {
		n += 1 + l + sovMarketIdentificationCode(uint64(l))
	}
	l = len(m.ExpiryDate)
	if l > 0 {
		n += 2 + l + sovMarketIdentificationCode(uint64(l))
	}
	l = len(m.Comments)
	if l > 0 {
		n += 2 + l + sovMarketIdentificationCode(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 2 + l + sovMarketIdentificationCode(uint64(l))
	}
	return n
}

func sovMarketIdentificationCode(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMarketIdentificationCode(x uint64) (n int) {
	return sovMarketIdentificationCode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MarketIdentificationCode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarketIdentificationCode
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
			return fmt.Errorf("proto: MarketIdentificationCode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketIdentificationCode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MIC", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MIC = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operating_MIC", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operating_MIC = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OPRT_SGMT", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OPRT_SGMT = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LegalEntityName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LegalEntityName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LegalEntityIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LegalEntityIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketCategory", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketCategory = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Acronym", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Acronym = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ISOCountryCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ISOCountryCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field City", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.City = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreationDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdateDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastUpdateDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastValidationDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastValidationDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiryDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExpiryDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comments", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Comments = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketIdentificationCode
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
				return ErrInvalidLengthMarketIdentificationCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketIdentificationCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMarketIdentificationCode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMarketIdentificationCode
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
func skipMarketIdentificationCode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMarketIdentificationCode
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
					return 0, ErrIntOverflowMarketIdentificationCode
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
					return 0, ErrIntOverflowMarketIdentificationCode
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
				return 0, ErrInvalidLengthMarketIdentificationCode
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMarketIdentificationCode
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMarketIdentificationCode
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMarketIdentificationCode        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMarketIdentificationCode          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMarketIdentificationCode = fmt.Errorf("proto: unexpected end of group")
)
