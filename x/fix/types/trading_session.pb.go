// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: re/fix/trading_session.proto

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

type TradingSession struct {
	SessionID                  string                      `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	TradingSessionStatusReport *TradingSessionStatusReport `protobuf:"bytes,2,opt,name=TradingSessionStatusReport,proto3" json:"TradingSessionStatusReport,omitempty"`
}

func (m *TradingSession) Reset()         { *m = TradingSession{} }
func (m *TradingSession) String() string { return proto.CompactTextString(m) }
func (*TradingSession) ProtoMessage()    {}
func (*TradingSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0d8ceb9122a9133, []int{0}
}
func (m *TradingSession) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TradingSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TradingSession.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TradingSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradingSession.Merge(m, src)
}
func (m *TradingSession) XXX_Size() int {
	return m.Size()
}
func (m *TradingSession) XXX_DiscardUnknown() {
	xxx_messageInfo_TradingSession.DiscardUnknown(m)
}

var xxx_messageInfo_TradingSession proto.InternalMessageInfo

func (m *TradingSession) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *TradingSession) GetTradingSessionStatusReport() *TradingSessionStatusReport {
	if m != nil {
		return m.TradingSessionStatusReport
	}
	return nil
}

// TradingSessionStatusReport with fields, tags and descriptions.
type TradingSessionStatusReport struct {
	// (336)	Identifies the specific trading session for which status is requested
	TradingSessionID string `protobuf:"bytes,1,opt,name=tradingSessionID,proto3" json:"tradingSessionID,omitempty"`
	// (625)	Identifies a specific sub-session or sub-segment within a trading session
	TradingSessionSubID string `protobuf:"bytes,2,opt,name=tradingSessionSubID,proto3" json:"tradingSessionSubID,omitempty"`
	// (335)	Unique identifier assigned by the party generating the request
	TradSesReqID string `protobuf:"bytes,3,opt,name=tradSesReqID,proto3" json:"tradSesReqID,omitempty"`
	// (1301)	Identifies the market or exchange for which the trading session status is requested
	MarketID string `protobuf:"bytes,4,opt,name=marketID,proto3" json:"marketID,omitempty"`
	// (263) Indicates whether the request is for a one-time snapshot or subscription for updates
	SubscriptionRequest string `protobuf:"bytes,5,opt,name=subscriptionRequest,proto3" json:"subscriptionRequest,omitempty"`
	// (48)	Identifier for the security being traded in the session
	SecurityID string `protobuf:"bytes,6,opt,name=securityID,proto3" json:"securityID,omitempty"`
	// (22)	Identifies the source of the security identifier
	SecurityIDSource string `protobuf:"bytes,7,opt,name=securityIDSource,proto3" json:"securityIDSource,omitempty"`
	// (55)	Symbol representing the security being traded in the session
	Symbol string `protobuf:"bytes,8,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// 207	Exchange where the security is listed
	SecurityExchange string `protobuf:"bytes,9,opt,name=securityExchange,proto3" json:"securityExchange,omitempty"`
	// (1300)	Identifies the specific market segment or sector within the trading session
	MarketSegmentID string `protobuf:"bytes,10,opt,name=marketSegmentID,proto3" json:"marketSegmentID,omitempty"`
	// (263) Type of trading session status request
	TradSesReqType int64 `protobuf:"varint,11,opt,name=tradSesReqType,proto3" json:"tradSesReqType,omitempty"`
	// (123)	Sub-type of trading session status request
	TradSesSubReqType int64 `protobuf:"varint,12,opt,name=tradSesSubReqType,proto3" json:"tradSesSubReqType,omitempty"`
	// (338)	Method used to determine the trading session
	TradSesMethod int64 `protobuf:"varint,13,opt,name=tradSesMethod,proto3" json:"tradSesMethod,omitempty"`
	// (339)	Mode of trading for the trading session
	TradSesMode int64 `protobuf:"varint,14,opt,name=tradSesMode,proto3" json:"tradSesMode,omitempty"`
	// (336) Date of the trading session for which status is requested
	TradingSessionDate string `protobuf:"bytes,15,opt,name=tradingSessionDate,proto3" json:"tradingSessionDate,omitempty"`
	// (338)Time of the trading session for which status is requested
	TradingSessionTime string `protobuf:"bytes,16,opt,name=tradingSessionTime,proto3" json:"tradingSessionTime,omitempty"`
	// (1147) Time of the sub-session or sub-segment within a trading session
	TradingSessionSubTime string `protobuf:"bytes,17,opt,name=tradingSessionSubTime,proto3" json:"tradingSessionSubTime,omitempty"`
	// (432) Date and time when the trading session status request is no longer valid and should be disregarded
	ExpirationDate string `protobuf:"bytes,18,opt,name=expirationDate,proto3" json:"expirationDate,omitempty"`
	// owner
	Creator string `protobuf:"bytes,19,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *TradingSessionStatusReport) Reset()         { *m = TradingSessionStatusReport{} }
func (m *TradingSessionStatusReport) String() string { return proto.CompactTextString(m) }
func (*TradingSessionStatusReport) ProtoMessage()    {}
func (*TradingSessionStatusReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0d8ceb9122a9133, []int{1}
}
func (m *TradingSessionStatusReport) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TradingSessionStatusReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TradingSessionStatusReport.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TradingSessionStatusReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradingSessionStatusReport.Merge(m, src)
}
func (m *TradingSessionStatusReport) XXX_Size() int {
	return m.Size()
}
func (m *TradingSessionStatusReport) XXX_DiscardUnknown() {
	xxx_messageInfo_TradingSessionStatusReport.DiscardUnknown(m)
}

var xxx_messageInfo_TradingSessionStatusReport proto.InternalMessageInfo

func (m *TradingSessionStatusReport) GetTradingSessionID() string {
	if m != nil {
		return m.TradingSessionID
	}
	return ""
}

func (m *TradingSessionStatusReport) GetTradingSessionSubID() string {
	if m != nil {
		return m.TradingSessionSubID
	}
	return ""
}

func (m *TradingSessionStatusReport) GetTradSesReqID() string {
	if m != nil {
		return m.TradSesReqID
	}
	return ""
}

func (m *TradingSessionStatusReport) GetMarketID() string {
	if m != nil {
		return m.MarketID
	}
	return ""
}

func (m *TradingSessionStatusReport) GetSubscriptionRequest() string {
	if m != nil {
		return m.SubscriptionRequest
	}
	return ""
}

func (m *TradingSessionStatusReport) GetSecurityID() string {
	if m != nil {
		return m.SecurityID
	}
	return ""
}

func (m *TradingSessionStatusReport) GetSecurityIDSource() string {
	if m != nil {
		return m.SecurityIDSource
	}
	return ""
}

func (m *TradingSessionStatusReport) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *TradingSessionStatusReport) GetSecurityExchange() string {
	if m != nil {
		return m.SecurityExchange
	}
	return ""
}

func (m *TradingSessionStatusReport) GetMarketSegmentID() string {
	if m != nil {
		return m.MarketSegmentID
	}
	return ""
}

func (m *TradingSessionStatusReport) GetTradSesReqType() int64 {
	if m != nil {
		return m.TradSesReqType
	}
	return 0
}

func (m *TradingSessionStatusReport) GetTradSesSubReqType() int64 {
	if m != nil {
		return m.TradSesSubReqType
	}
	return 0
}

func (m *TradingSessionStatusReport) GetTradSesMethod() int64 {
	if m != nil {
		return m.TradSesMethod
	}
	return 0
}

func (m *TradingSessionStatusReport) GetTradSesMode() int64 {
	if m != nil {
		return m.TradSesMode
	}
	return 0
}

func (m *TradingSessionStatusReport) GetTradingSessionDate() string {
	if m != nil {
		return m.TradingSessionDate
	}
	return ""
}

func (m *TradingSessionStatusReport) GetTradingSessionTime() string {
	if m != nil {
		return m.TradingSessionTime
	}
	return ""
}

func (m *TradingSessionStatusReport) GetTradingSessionSubTime() string {
	if m != nil {
		return m.TradingSessionSubTime
	}
	return ""
}

func (m *TradingSessionStatusReport) GetExpirationDate() string {
	if m != nil {
		return m.ExpirationDate
	}
	return ""
}

func (m *TradingSessionStatusReport) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*TradingSession)(nil), "jim380.re.fix.TradingSession")
	proto.RegisterType((*TradingSessionStatusReport)(nil), "jim380.re.fix.TradingSessionStatusReport")
}

func init() { proto.RegisterFile("re/fix/trading_session.proto", fileDescriptor_e0d8ceb9122a9133) }

var fileDescriptor_e0d8ceb9122a9133 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x63, 0x0a, 0x69, 0xf3, 0xd2, 0xa4, 0xed, 0x55, 0xa0, 0x53, 0x55, 0x59, 0x56, 0x84,
	0x50, 0x40, 0xc8, 0xae, 0x28, 0x03, 0x62, 0x44, 0x66, 0xf0, 0xc0, 0x62, 0x67, 0x62, 0x41, 0xb6,
	0xf3, 0x9a, 0x1c, 0x60, 0x9f, 0x7b, 0x77, 0x96, 0x9c, 0x6f, 0xc1, 0x8a, 0xc4, 0x07, 0x62, 0xec,
	0xc8, 0x88, 0x92, 0x2f, 0x82, 0x7c, 0x8e, 0x9b, 0x38, 0x31, 0x1d, 0xdf, 0xef, 0xff, 0x7b, 0xf7,
	0xee, 0x39, 0xb1, 0xe1, 0x52, 0xa0, 0x73, 0xc3, 0x0a, 0x47, 0x89, 0x70, 0xca, 0xd2, 0xd9, 0x17,
	0x89, 0x52, 0x32, 0x9e, 0xda, 0x99, 0xe0, 0x8a, 0x93, 0xc1, 0x57, 0x96, 0x5c, 0xbf, 0xbb, 0xb2,
	0x05, 0xda, 0x37, 0xac, 0x18, 0xfd, 0x34, 0x60, 0x38, 0xa9, 0xc4, 0xa0, 0xf2, 0xc8, 0x25, 0xf4,
	0xd6, 0x2d, 0x9e, 0x4b, 0x0d, 0xcb, 0x18, 0xf7, 0xfc, 0x0d, 0x20, 0x0c, 0x2e, 0x9a, 0x7e, 0xa0,
	0x42, 0x95, 0x4b, 0x1f, 0x33, 0x2e, 0x14, 0x7d, 0x64, 0x19, 0xe3, 0xfe, 0x9b, 0x97, 0x76, 0x63,
	0x88, 0xfd, 0xff, 0x06, 0xff, 0x81, 0xc3, 0x46, 0xbf, 0xba, 0x0f, 0xcd, 0x22, 0xaf, 0xe0, 0x54,
	0x35, 0xd2, 0xfb, 0xeb, 0xee, 0x71, 0x72, 0x05, 0xe7, 0x4d, 0x16, 0xe4, 0x91, 0xe7, 0xea, 0xeb,
	0xf6, 0xfc, 0xb6, 0x88, 0x8c, 0xe0, 0xb8, 0xc4, 0x01, 0x4a, 0x1f, 0x6f, 0x3d, 0x97, 0x1e, 0x68,
	0xb5, 0xc1, 0xc8, 0x05, 0x1c, 0x25, 0xa1, 0xf8, 0x86, 0xca, 0x73, 0xe9, 0x63, 0x9d, 0xdf, 0xd7,
	0xe5, 0x44, 0x99, 0x47, 0x32, 0x16, 0x2c, 0x53, 0x8c, 0xa7, 0x3e, 0xde, 0xe6, 0x28, 0x15, 0x7d,
	0x52, 0x4d, 0x6c, 0x89, 0x88, 0x09, 0x20, 0x31, 0xce, 0x05, 0x53, 0x0b, 0xcf, 0xa5, 0x5d, 0x2d,
	0x6e, 0x91, 0x72, 0xdf, 0x4d, 0x15, 0xf0, 0x5c, 0xc4, 0x48, 0x0f, 0xab, 0x7d, 0x77, 0x39, 0x79,
	0x06, 0x5d, 0xb9, 0x48, 0x22, 0xfe, 0x9d, 0x1e, 0x69, 0x63, 0x5d, 0x6d, 0x9f, 0xf1, 0xb1, 0x88,
	0xe7, 0x61, 0x3a, 0x43, 0xda, 0x6b, 0x9e, 0x51, 0x73, 0x32, 0x86, 0x93, 0x6a, 0x9b, 0x00, 0x67,
	0x09, 0xa6, 0xe5, 0x92, 0xa0, 0xd5, 0x5d, 0x4c, 0x5e, 0xc0, 0x70, 0xf3, 0x5c, 0x26, 0x8b, 0x0c,
	0x69, 0xdf, 0x32, 0xc6, 0x07, 0xfe, 0x0e, 0x25, 0xaf, 0xe1, 0x6c, 0x4d, 0x82, 0x3c, 0xaa, 0xd5,
	0x63, 0xad, 0xee, 0x07, 0xe4, 0x39, 0x0c, 0xd6, 0xf0, 0x13, 0xaa, 0x39, 0x9f, 0xd2, 0x81, 0x36,
	0x9b, 0x90, 0x58, 0xd0, 0xaf, 0x01, 0x9f, 0x22, 0x1d, 0x6a, 0x67, 0x1b, 0x11, 0x1b, 0x48, 0xf3,
	0x07, 0x76, 0x43, 0x85, 0xf4, 0x44, 0xaf, 0xd2, 0x92, 0xec, 0xfb, 0x13, 0x96, 0x20, 0x3d, 0x6d,
	0xf3, 0xcb, 0x84, 0xbc, 0x85, 0xa7, 0x7b, 0x7f, 0x20, 0xdd, 0x72, 0xa6, 0x5b, 0xda, 0xc3, 0xf2,
	0x99, 0x61, 0x91, 0x31, 0x11, 0xaa, 0xfa, 0x46, 0x44, 0xeb, 0x3b, 0x94, 0x50, 0x38, 0x8c, 0x05,
	0x86, 0x8a, 0x0b, 0x7a, 0xae, 0x85, 0xba, 0xfc, 0xf0, 0xfe, 0xf7, 0xd2, 0x34, 0xee, 0x96, 0xa6,
	0xf1, 0x77, 0x69, 0x1a, 0x3f, 0x56, 0x66, 0xe7, 0x6e, 0x65, 0x76, 0xfe, 0xac, 0xcc, 0xce, 0x67,
	0x6b, 0xc6, 0xd4, 0x3c, 0x8f, 0xec, 0x98, 0x27, 0x4e, 0xf5, 0x26, 0x3a, 0x3e, 0x3a, 0x45, 0xf5,
	0x55, 0x58, 0x64, 0x28, 0xa3, 0xae, 0xfe, 0x18, 0x5c, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x8f,
	0x28, 0x37, 0xfd, 0x2c, 0x04, 0x00, 0x00,
}

func (m *TradingSession) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TradingSession) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TradingSession) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TradingSessionStatusReport != nil {
		{
			size, err := m.TradingSessionStatusReport.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTradingSession(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.SessionID) > 0 {
		i -= len(m.SessionID)
		copy(dAtA[i:], m.SessionID)
		i = encodeVarintTradingSession(dAtA, i, uint64(len(m.SessionID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TradingSessionStatusReport) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TradingSessionStatusReport) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TradingSessionStatusReport) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTradingSession(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x9a
	}
	if len(m.ExpirationDate) > 0 {
		i -= len(m.ExpirationDate)
		copy(dAtA[i:], m.ExpirationDate)
		i = encodeVarintTradingSession(dAtA, i, uint64(len(m.ExpirationDate)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x92
	}
	if len(m.TradingSessionSubTime) > 0 {
		i -= len(m.TradingSessionSubTime)
		copy(dAtA[i:], m.TradingSessionSubTime)
		i = encodeVarintTradingSession(dAtA, i, uint64(len(m.TradingSessionSubTime)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if len(m.TradingSessionTime) > 0 {
		i -= len(m.TradingSessionTime)
		copy(dAtA[i:], m.TradingSessionTime)
		i = encodeVarintTradingSession(dAtA, i, uint64(len(m.TradingSessionTime)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if len(m.TradingSessionDate) > 0 {
		i -= len(m.TradingSessionDate)
		copy(dAtA[i:], m.TradingSessionDate)
		i = encodeVarintTradingSession(dAtA, i, uint64(len(m.TradingSessionDate)))
		i--
		dAtA[i] = 0x7a
	}
	if m.TradSesMode != 0 {
		i = encodeVarintTradingSession(dAtA, i, uint64(m.TradSesMode))
		i--
		dAtA[i] = 0x70
	}
	if m.TradSesMethod != 0 {
		i = encodeVarintTradingSession(dAtA, i, uint64(m.TradSesMethod))
		i--
		dAtA[i] = 0x68
	}
	if m.TradSesSubReqType != 0 {
		i = encodeVarintTradingSession(dAtA, i, uint64(m.TradSesSubReqType))
		i--
		dAtA[i] = 0x60
	}
	if m.TradSesReqType != 0 {
		i = encodeVarintTradingSession(dAtA, i, uint64(m.TradSesReqType))
		i--
		dAtA[i] = 0x58
	}
	if len(m.MarketSegmentID) > 0 {
		i -= len(m.MarketSegmentID)
		copy(dAtA[i:], m.MarketSegmentID)
		i = encodeVarintTradingSession(dAtA, i, uint64(len(m.MarketSegmentID)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.SecurityExchange) > 0 {
		i -= len(m.SecurityExchange)
		copy(dAtA[i:], m.SecurityExchange)
		i = encodeVarintTradingSession(dAtA, i, uint64(len(m.SecurityExchange)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTradingSession(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.SecurityIDSource) > 0 {
		i -= len(m.SecurityIDSource)
		copy(dAtA[i:], m.SecurityIDSource)
		i = encodeVarintTradingSession(dAtA, i, uint64(len(m.SecurityIDSource)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.SecurityID) > 0 {
		i -= len(m.SecurityID)
		copy(dAtA[i:], m.SecurityID)
		i = encodeVarintTradingSession(dAtA, i, uint64(len(m.SecurityID)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.SubscriptionRequest) > 0 {
		i -= len(m.SubscriptionRequest)
		copy(dAtA[i:], m.SubscriptionRequest)
		i = encodeVarintTradingSession(dAtA, i, uint64(len(m.SubscriptionRequest)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.MarketID) > 0 {
		i -= len(m.MarketID)
		copy(dAtA[i:], m.MarketID)
		i = encodeVarintTradingSession(dAtA, i, uint64(len(m.MarketID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TradSesReqID) > 0 {
		i -= len(m.TradSesReqID)
		copy(dAtA[i:], m.TradSesReqID)
		i = encodeVarintTradingSession(dAtA, i, uint64(len(m.TradSesReqID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TradingSessionSubID) > 0 {
		i -= len(m.TradingSessionSubID)
		copy(dAtA[i:], m.TradingSessionSubID)
		i = encodeVarintTradingSession(dAtA, i, uint64(len(m.TradingSessionSubID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TradingSessionID) > 0 {
		i -= len(m.TradingSessionID)
		copy(dAtA[i:], m.TradingSessionID)
		i = encodeVarintTradingSession(dAtA, i, uint64(len(m.TradingSessionID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTradingSession(dAtA []byte, offset int, v uint64) int {
	offset -= sovTradingSession(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TradingSession) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SessionID)
	if l > 0 {
		n += 1 + l + sovTradingSession(uint64(l))
	}
	if m.TradingSessionStatusReport != nil {
		l = m.TradingSessionStatusReport.Size()
		n += 1 + l + sovTradingSession(uint64(l))
	}
	return n
}

func (m *TradingSessionStatusReport) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TradingSessionID)
	if l > 0 {
		n += 1 + l + sovTradingSession(uint64(l))
	}
	l = len(m.TradingSessionSubID)
	if l > 0 {
		n += 1 + l + sovTradingSession(uint64(l))
	}
	l = len(m.TradSesReqID)
	if l > 0 {
		n += 1 + l + sovTradingSession(uint64(l))
	}
	l = len(m.MarketID)
	if l > 0 {
		n += 1 + l + sovTradingSession(uint64(l))
	}
	l = len(m.SubscriptionRequest)
	if l > 0 {
		n += 1 + l + sovTradingSession(uint64(l))
	}
	l = len(m.SecurityID)
	if l > 0 {
		n += 1 + l + sovTradingSession(uint64(l))
	}
	l = len(m.SecurityIDSource)
	if l > 0 {
		n += 1 + l + sovTradingSession(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTradingSession(uint64(l))
	}
	l = len(m.SecurityExchange)
	if l > 0 {
		n += 1 + l + sovTradingSession(uint64(l))
	}
	l = len(m.MarketSegmentID)
	if l > 0 {
		n += 1 + l + sovTradingSession(uint64(l))
	}
	if m.TradSesReqType != 0 {
		n += 1 + sovTradingSession(uint64(m.TradSesReqType))
	}
	if m.TradSesSubReqType != 0 {
		n += 1 + sovTradingSession(uint64(m.TradSesSubReqType))
	}
	if m.TradSesMethod != 0 {
		n += 1 + sovTradingSession(uint64(m.TradSesMethod))
	}
	if m.TradSesMode != 0 {
		n += 1 + sovTradingSession(uint64(m.TradSesMode))
	}
	l = len(m.TradingSessionDate)
	if l > 0 {
		n += 1 + l + sovTradingSession(uint64(l))
	}
	l = len(m.TradingSessionTime)
	if l > 0 {
		n += 2 + l + sovTradingSession(uint64(l))
	}
	l = len(m.TradingSessionSubTime)
	if l > 0 {
		n += 2 + l + sovTradingSession(uint64(l))
	}
	l = len(m.ExpirationDate)
	if l > 0 {
		n += 2 + l + sovTradingSession(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 2 + l + sovTradingSession(uint64(l))
	}
	return n
}

func sovTradingSession(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTradingSession(x uint64) (n int) {
	return sovTradingSession(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TradingSession) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTradingSession
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
			return fmt.Errorf("proto: TradingSession: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TradingSession: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
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
				return ErrInvalidLengthTradingSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradingSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradingSessionStatusReport", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
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
				return ErrInvalidLengthTradingSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTradingSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TradingSessionStatusReport == nil {
				m.TradingSessionStatusReport = &TradingSessionStatusReport{}
			}
			if err := m.TradingSessionStatusReport.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTradingSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTradingSession
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
func (m *TradingSessionStatusReport) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTradingSession
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
			return fmt.Errorf("proto: TradingSessionStatusReport: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TradingSessionStatusReport: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradingSessionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
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
				return ErrInvalidLengthTradingSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradingSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradingSessionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradingSessionSubID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
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
				return ErrInvalidLengthTradingSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradingSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradingSessionSubID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradSesReqID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
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
				return ErrInvalidLengthTradingSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradingSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradSesReqID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
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
				return ErrInvalidLengthTradingSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradingSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriptionRequest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
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
				return ErrInvalidLengthTradingSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradingSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubscriptionRequest = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
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
				return ErrInvalidLengthTradingSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradingSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecurityID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityIDSource", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
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
				return ErrInvalidLengthTradingSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradingSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecurityIDSource = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
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
				return ErrInvalidLengthTradingSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradingSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityExchange", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
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
				return ErrInvalidLengthTradingSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradingSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecurityExchange = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketSegmentID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
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
				return ErrInvalidLengthTradingSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradingSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketSegmentID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradSesReqType", wireType)
			}
			m.TradSesReqType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TradSesReqType |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradSesSubReqType", wireType)
			}
			m.TradSesSubReqType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TradSesSubReqType |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradSesMethod", wireType)
			}
			m.TradSesMethod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TradSesMethod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradSesMode", wireType)
			}
			m.TradSesMode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TradSesMode |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradingSessionDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
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
				return ErrInvalidLengthTradingSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradingSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradingSessionDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradingSessionTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
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
				return ErrInvalidLengthTradingSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradingSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradingSessionTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradingSessionSubTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
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
				return ErrInvalidLengthTradingSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradingSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradingSessionSubTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
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
				return ErrInvalidLengthTradingSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradingSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExpirationDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradingSession
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
				return ErrInvalidLengthTradingSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradingSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTradingSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTradingSession
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
func skipTradingSession(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTradingSession
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
					return 0, ErrIntOverflowTradingSession
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
					return 0, ErrIntOverflowTradingSession
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
				return 0, ErrInvalidLengthTradingSession
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTradingSession
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTradingSession
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTradingSession        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTradingSession          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTradingSession = fmt.Errorf("proto: unexpected end of group")
)
