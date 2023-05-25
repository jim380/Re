// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: re/fix/orders_execution_report.proto

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

type OrdersExecutionReport struct {
	// A string field that specifies the FIX session ID for the message.
	SessionID string `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	// A Header field that contains standard header information for the message,
	// such as the message type, sender and receiver identification, and sequence
	// number.
	Header *Header `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	// A string field that contains the client order ID for the order being
	// reported.
	ClOrdID string `protobuf:"bytes,3,opt,name=clOrdID,proto3" json:"clOrdID,omitempty"`
	// A string field that contains the ID of the order being reported. OrderID is a string field that contains the unique identifier assigned to an order by the system or trading platform. It is typically generated by the system upon order submission and is used to uniquely identify the order within the system. The orderID is useful for internal system processes, such as order matching, order book management, and order status tracking
	OrderID string `protobuf:"bytes,4,opt,name=orderID,proto3" json:"orderID,omitempty"`
	// A string field that contains the ID of the execution being reported, if any. It is a string field that contains the identifier assigned to a particular execution of an order. Each execution of an order typically has a separate execID associated with it. The execID can be used to uniquely identify and track a specific execution in the order execution report. It is useful for monitoring the status, details, and timing of individual executions within the overall order lifecycle. The execID field in the FIX protocol can be auto-generated by the trading venue or the FIX protocol software when responding with an execution report. The trading venue's execID serves as a local identifier within their own system to track and reference executions. It allows the trading venue to uniquely identify each execution and associate it with the relevant order and trade details. When sending an execution report to clients or counterparties via the FIX protocol, the trading venue includes its own generated execID value in the execution report message. This enables the recipient to correlate the execution report with the specific execution on the trading venue's side
	ExecID string `protobuf:"bytes,5,opt,name=execID,proto3" json:"execID,omitempty"`
	// A string field that indicates the current status of the order, such as
	// "New", "Partially filled", "Filled", "Cancelled", "Pending Cancel",
	// "Rejected", etc.
	OrdStatus string `protobuf:"bytes,6,opt,name=ordStatus,proto3" json:"ordStatus,omitempty"`
	// A string field that indicates the type of the execution being reported,
	// such as "New", "Partial fill", "Fill", "Done for day", "Cancelled", etc.
	ExecType string `protobuf:"bytes,7,opt,name=execType,proto3" json:"execType,omitempty"`
	// A string field that identifies the security being traded.
	Symbol string `protobuf:"bytes,8,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// An integer field that specifies the side of the order (buy or sell).
	Side int64 `protobuf:"varint,9,opt,name=side,proto3" json:"side,omitempty"`
	// An integer field that specifies the quantity of the order.
	OrderQty string `protobuf:"bytes,10,opt,name=orderQty,proto3" json:"orderQty,omitempty"`
	// An integer field that specifies the price of the order.
	Price string `protobuf:"bytes,11,opt,name=price,proto3" json:"price,omitempty"`
	// An integer field that specifies the time-in-force value for the order, such
	// as "Day", "GTC" (Good 'Til Cancelled), "IOC" (Immediate or Cancel), etc.
	TimeInForce int64 `protobuf:"varint,12,opt,name=timeInForce,proto3" json:"timeInForce,omitempty"`
	// An integer field that specifies the price of the last execution, if any.
	LastPx int64 `protobuf:"varint,13,opt,name=lastPx,proto3" json:"lastPx,omitempty"`
	// An integer field that specifies the quantity of the last execution, if any.
	LastQty int64 `protobuf:"varint,14,opt,name=lastQty,proto3" json:"lastQty,omitempty"`
	// An integer field that specifies the quantity of the order that remains open
	// and has not yet been filled or cancelled.
	LeavesQty int64 `protobuf:"varint,15,opt,name=leavesQty,proto3" json:"leavesQty,omitempty"`
	// An integer field that specifies the total quantity of the order that has
	// been filled.
	CumQty int64 `protobuf:"varint,16,opt,name=cumQty,proto3" json:"cumQty,omitempty"`
	// An integer field that specifies the average price of all executions that
	// have occurred for the order.
	AvgPx int64 `protobuf:"varint,17,opt,name=avgPx,proto3" json:"avgPx,omitempty"`
	// A string field that provides additional information about the order status
	// or execution, such as an error message or reason for rejection.
	Text string `protobuf:"bytes,18,opt,name=text,proto3" json:"text,omitempty"`
	// A string field that specifies the time of the transaction.
	TransactTime string `protobuf:"bytes,19,opt,name=transactTime,proto3" json:"transactTime,omitempty"`
	// A Trailer field that contains standard trailer information for the message,
	// such as the message checksum.
	Trailer *Trailer `protobuf:"bytes,20,opt,name=trailer,proto3" json:"trailer,omitempty"`
	// A field that identifies the creator of the message.
	Creator string `protobuf:"bytes,21,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *OrdersExecutionReport) Reset()         { *m = OrdersExecutionReport{} }
func (m *OrdersExecutionReport) String() string { return proto.CompactTextString(m) }
func (*OrdersExecutionReport) ProtoMessage()    {}
func (*OrdersExecutionReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_612a8365f41aaf72, []int{0}
}
func (m *OrdersExecutionReport) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrdersExecutionReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrdersExecutionReport.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrdersExecutionReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrdersExecutionReport.Merge(m, src)
}
func (m *OrdersExecutionReport) XXX_Size() int {
	return m.Size()
}
func (m *OrdersExecutionReport) XXX_DiscardUnknown() {
	xxx_messageInfo_OrdersExecutionReport.DiscardUnknown(m)
}

var xxx_messageInfo_OrdersExecutionReport proto.InternalMessageInfo

func (m *OrdersExecutionReport) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *OrdersExecutionReport) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *OrdersExecutionReport) GetClOrdID() string {
	if m != nil {
		return m.ClOrdID
	}
	return ""
}

func (m *OrdersExecutionReport) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func (m *OrdersExecutionReport) GetExecID() string {
	if m != nil {
		return m.ExecID
	}
	return ""
}

func (m *OrdersExecutionReport) GetOrdStatus() string {
	if m != nil {
		return m.OrdStatus
	}
	return ""
}

func (m *OrdersExecutionReport) GetExecType() string {
	if m != nil {
		return m.ExecType
	}
	return ""
}

func (m *OrdersExecutionReport) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *OrdersExecutionReport) GetSide() int64 {
	if m != nil {
		return m.Side
	}
	return 0
}

func (m *OrdersExecutionReport) GetOrderQty() string {
	if m != nil {
		return m.OrderQty
	}
	return ""
}

func (m *OrdersExecutionReport) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *OrdersExecutionReport) GetTimeInForce() int64 {
	if m != nil {
		return m.TimeInForce
	}
	return 0
}

func (m *OrdersExecutionReport) GetLastPx() int64 {
	if m != nil {
		return m.LastPx
	}
	return 0
}

func (m *OrdersExecutionReport) GetLastQty() int64 {
	if m != nil {
		return m.LastQty
	}
	return 0
}

func (m *OrdersExecutionReport) GetLeavesQty() int64 {
	if m != nil {
		return m.LeavesQty
	}
	return 0
}

func (m *OrdersExecutionReport) GetCumQty() int64 {
	if m != nil {
		return m.CumQty
	}
	return 0
}

func (m *OrdersExecutionReport) GetAvgPx() int64 {
	if m != nil {
		return m.AvgPx
	}
	return 0
}

func (m *OrdersExecutionReport) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *OrdersExecutionReport) GetTransactTime() string {
	if m != nil {
		return m.TransactTime
	}
	return ""
}

func (m *OrdersExecutionReport) GetTrailer() *Trailer {
	if m != nil {
		return m.Trailer
	}
	return nil
}

func (m *OrdersExecutionReport) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*OrdersExecutionReport)(nil), "jim380.re.fix.OrdersExecutionReport")
}

func init() {
	proto.RegisterFile("re/fix/orders_execution_report.proto", fileDescriptor_612a8365f41aaf72)
}

var fileDescriptor_612a8365f41aaf72 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xc1, 0x6e, 0xd3, 0x4e,
	0x10, 0xc6, 0xe3, 0x7f, 0xd3, 0xa4, 0xd9, 0xb4, 0x7f, 0x60, 0x69, 0xaa, 0x51, 0x05, 0x56, 0x54,
	0x71, 0xc8, 0x05, 0xa7, 0xa2, 0x17, 0xc4, 0x11, 0x05, 0x84, 0x4f, 0x2d, 0x26, 0x27, 0x2e, 0xd1,
	0xc6, 0x9e, 0xa6, 0x8b, 0x6c, 0xaf, 0xb5, 0xbb, 0xa9, 0x9c, 0xb7, 0xe0, 0xb1, 0x38, 0xf6, 0xc8,
	0x11, 0x25, 0x57, 0x1e, 0x02, 0xed, 0xac, 0xdd, 0x52, 0x6e, 0xfb, 0x7d, 0xbf, 0xf1, 0xb7, 0x33,
	0x9e, 0x65, 0xaf, 0x34, 0x4e, 0xaf, 0x65, 0x3d, 0x55, 0x3a, 0x43, 0x6d, 0x16, 0x58, 0x63, 0xba,
	0xb6, 0x52, 0x95, 0x0b, 0x8d, 0x95, 0xd2, 0x36, 0xaa, 0xb4, 0xb2, 0x8a, 0x1f, 0x7d, 0x93, 0xc5,
	0xc5, 0xdb, 0xf3, 0x48, 0x63, 0x74, 0x2d, 0xeb, 0xd3, 0x97, 0xcd, 0x47, 0xc6, 0x8a, 0x32, 0x13,
	0x3a, 0x5b, 0x14, 0x68, 0x8c, 0x58, 0xa1, 0xaf, 0x3e, 0xfb, 0xdd, 0x65, 0xa3, 0x4b, 0xca, 0xfb,
	0xd0, 0xc6, 0x25, 0x94, 0xc6, 0x5f, 0xb0, 0x81, 0x41, 0x63, 0xa4, 0x2a, 0xe3, 0x19, 0x04, 0xe3,
	0x60, 0x32, 0x48, 0x1e, 0x0c, 0xfe, 0x9a, 0xf5, 0x6e, 0x50, 0x64, 0xa8, 0xe1, 0xbf, 0x71, 0x30,
	0x19, 0xbe, 0x19, 0x45, 0x8f, 0xae, 0x8d, 0x3e, 0x11, 0x4c, 0x9a, 0x22, 0x0e, 0xac, 0x9f, 0xe6,
	0x97, 0x3a, 0x8b, 0x67, 0xb0, 0x47, 0x51, 0xad, 0x74, 0x84, 0xe6, 0x89, 0x67, 0xd0, 0xf5, 0xa4,
	0x91, 0xfc, 0x84, 0xf5, 0xdc, 0x88, 0xf1, 0x0c, 0xf6, 0x09, 0x34, 0xca, 0x35, 0xa6, 0x74, 0xf6,
	0xc5, 0x0a, 0xbb, 0x36, 0xd0, 0xf3, 0x8d, 0xdd, 0x1b, 0xfc, 0x94, 0x1d, 0xb8, 0xba, 0xf9, 0xa6,
	0x42, 0xe8, 0x13, 0xbc, 0xd7, 0x2e, 0xd1, 0x6c, 0x8a, 0xa5, 0xca, 0xe1, 0xc0, 0x27, 0x7a, 0xc5,
	0x39, 0xeb, 0x1a, 0x99, 0x21, 0x0c, 0xc6, 0xc1, 0x64, 0x2f, 0xa1, 0xb3, 0xcb, 0xa1, 0x46, 0x3e,
	0xdb, 0x0d, 0x30, 0x9f, 0xd3, 0x6a, 0x7e, 0xcc, 0xf6, 0x2b, 0x2d, 0x53, 0x84, 0x21, 0x01, 0x2f,
	0xf8, 0x98, 0x0d, 0xad, 0x2c, 0x30, 0x2e, 0x3f, 0x2a, 0x9d, 0x22, 0x1c, 0x52, 0xd8, 0xdf, 0x96,
	0xbb, 0x3f, 0x17, 0xc6, 0x5e, 0xd5, 0x70, 0x44, 0xb0, 0x51, 0xee, 0x1f, 0xb8, 0x93, 0xbb, 0xea,
	0x7f, 0x02, 0xad, 0x74, 0xb3, 0xe6, 0x28, 0x6e, 0xd1, 0x38, 0xf6, 0x84, 0xd8, 0x83, 0xe1, 0xf2,
	0xd2, 0x75, 0xe1, 0xd0, 0x53, 0x9f, 0xe7, 0x95, 0xeb, 0x4f, 0xdc, 0xae, 0xae, 0x6a, 0x78, 0x46,
	0xb6, 0x17, 0x6e, 0x4a, 0x8b, 0xb5, 0x05, 0x4e, 0x4d, 0xd3, 0x99, 0x9f, 0xb1, 0x43, 0xab, 0x45,
	0x69, 0x44, 0x6a, 0xe7, 0xb2, 0x40, 0x78, 0x4e, 0xec, 0x91, 0xc7, 0xcf, 0x59, 0xdf, 0x6a, 0x21,
	0x73, 0xd4, 0x70, 0x4c, 0xbb, 0x3e, 0xf9, 0x67, 0xd7, 0x73, 0x4f, 0x93, 0xb6, 0x8c, 0xb6, 0xad,
	0x51, 0x58, 0xa5, 0x61, 0xd4, 0x6c, 0xdb, 0xcb, 0xf7, 0xef, 0x7e, 0x6c, 0xc3, 0xe0, 0x6e, 0x1b,
	0x06, 0xbf, 0xb6, 0x61, 0xf0, 0x7d, 0x17, 0x76, 0xee, 0x76, 0x61, 0xe7, 0xe7, 0x2e, 0xec, 0x7c,
	0x1d, 0xaf, 0xa4, 0xbd, 0x59, 0x2f, 0xa3, 0x54, 0x15, 0x53, 0x1f, 0x3f, 0x4d, 0x70, 0x5a, 0xd3,
	0xdb, 0xb5, 0x9b, 0x0a, 0xcd, 0xb2, 0x47, 0x2f, 0xf6, 0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x12, 0x4a, 0xd2, 0x8f, 0x07, 0x03, 0x00, 0x00,
}

func (m *OrdersExecutionReport) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrdersExecutionReport) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrdersExecutionReport) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xaa
	}
	if m.Trailer != nil {
		{
			size, err := m.Trailer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa2
	}
	if len(m.TransactTime) > 0 {
		i -= len(m.TransactTime)
		copy(dAtA[i:], m.TransactTime)
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(len(m.TransactTime)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x9a
	}
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(len(m.Text)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x92
	}
	if m.AvgPx != 0 {
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(m.AvgPx))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.CumQty != 0 {
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(m.CumQty))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.LeavesQty != 0 {
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(m.LeavesQty))
		i--
		dAtA[i] = 0x78
	}
	if m.LastQty != 0 {
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(m.LastQty))
		i--
		dAtA[i] = 0x70
	}
	if m.LastPx != 0 {
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(m.LastPx))
		i--
		dAtA[i] = 0x68
	}
	if m.TimeInForce != 0 {
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(m.TimeInForce))
		i--
		dAtA[i] = 0x60
	}
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.OrderQty) > 0 {
		i -= len(m.OrderQty)
		copy(dAtA[i:], m.OrderQty)
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(len(m.OrderQty)))
		i--
		dAtA[i] = 0x52
	}
	if m.Side != 0 {
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(m.Side))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.ExecType) > 0 {
		i -= len(m.ExecType)
		copy(dAtA[i:], m.ExecType)
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(len(m.ExecType)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.OrdStatus) > 0 {
		i -= len(m.OrdStatus)
		copy(dAtA[i:], m.OrdStatus)
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(len(m.OrdStatus)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ExecID) > 0 {
		i -= len(m.ExecID)
		copy(dAtA[i:], m.ExecID)
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(len(m.ExecID)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.OrderID) > 0 {
		i -= len(m.OrderID)
		copy(dAtA[i:], m.OrderID)
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(len(m.OrderID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ClOrdID) > 0 {
		i -= len(m.ClOrdID)
		copy(dAtA[i:], m.ClOrdID)
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(len(m.ClOrdID)))
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
			i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.SessionID) > 0 {
		i -= len(m.SessionID)
		copy(dAtA[i:], m.SessionID)
		i = encodeVarintOrdersExecutionReport(dAtA, i, uint64(len(m.SessionID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrdersExecutionReport(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrdersExecutionReport(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrdersExecutionReport) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SessionID)
	if l > 0 {
		n += 1 + l + sovOrdersExecutionReport(uint64(l))
	}
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovOrdersExecutionReport(uint64(l))
	}
	l = len(m.ClOrdID)
	if l > 0 {
		n += 1 + l + sovOrdersExecutionReport(uint64(l))
	}
	l = len(m.OrderID)
	if l > 0 {
		n += 1 + l + sovOrdersExecutionReport(uint64(l))
	}
	l = len(m.ExecID)
	if l > 0 {
		n += 1 + l + sovOrdersExecutionReport(uint64(l))
	}
	l = len(m.OrdStatus)
	if l > 0 {
		n += 1 + l + sovOrdersExecutionReport(uint64(l))
	}
	l = len(m.ExecType)
	if l > 0 {
		n += 1 + l + sovOrdersExecutionReport(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovOrdersExecutionReport(uint64(l))
	}
	if m.Side != 0 {
		n += 1 + sovOrdersExecutionReport(uint64(m.Side))
	}
	l = len(m.OrderQty)
	if l > 0 {
		n += 1 + l + sovOrdersExecutionReport(uint64(l))
	}
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovOrdersExecutionReport(uint64(l))
	}
	if m.TimeInForce != 0 {
		n += 1 + sovOrdersExecutionReport(uint64(m.TimeInForce))
	}
	if m.LastPx != 0 {
		n += 1 + sovOrdersExecutionReport(uint64(m.LastPx))
	}
	if m.LastQty != 0 {
		n += 1 + sovOrdersExecutionReport(uint64(m.LastQty))
	}
	if m.LeavesQty != 0 {
		n += 1 + sovOrdersExecutionReport(uint64(m.LeavesQty))
	}
	if m.CumQty != 0 {
		n += 2 + sovOrdersExecutionReport(uint64(m.CumQty))
	}
	if m.AvgPx != 0 {
		n += 2 + sovOrdersExecutionReport(uint64(m.AvgPx))
	}
	l = len(m.Text)
	if l > 0 {
		n += 2 + l + sovOrdersExecutionReport(uint64(l))
	}
	l = len(m.TransactTime)
	if l > 0 {
		n += 2 + l + sovOrdersExecutionReport(uint64(l))
	}
	if m.Trailer != nil {
		l = m.Trailer.Size()
		n += 2 + l + sovOrdersExecutionReport(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 2 + l + sovOrdersExecutionReport(uint64(l))
	}
	return n
}

func sovOrdersExecutionReport(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrdersExecutionReport(x uint64) (n int) {
	return sovOrdersExecutionReport(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrdersExecutionReport) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrdersExecutionReport
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
			return fmt.Errorf("proto: OrdersExecutionReport: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrdersExecutionReport: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
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
				return ErrInvalidLengthOrdersExecutionReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersExecutionReport
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
					return ErrIntOverflowOrdersExecutionReport
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
				return ErrInvalidLengthOrdersExecutionReport
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersExecutionReport
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
				return fmt.Errorf("proto: wrong wireType = %d for field ClOrdID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
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
				return ErrInvalidLengthOrdersExecutionReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersExecutionReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClOrdID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
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
				return ErrInvalidLengthOrdersExecutionReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersExecutionReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
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
				return ErrInvalidLengthOrdersExecutionReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersExecutionReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrdStatus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
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
				return ErrInvalidLengthOrdersExecutionReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersExecutionReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrdStatus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
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
				return ErrInvalidLengthOrdersExecutionReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersExecutionReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
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
				return ErrInvalidLengthOrdersExecutionReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersExecutionReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Side", wireType)
			}
			m.Side = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Side |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderQty", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
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
				return ErrInvalidLengthOrdersExecutionReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersExecutionReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderQty = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
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
				return ErrInvalidLengthOrdersExecutionReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersExecutionReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeInForce", wireType)
			}
			m.TimeInForce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeInForce |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastPx", wireType)
			}
			m.LastPx = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastPx |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastQty", wireType)
			}
			m.LastQty = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastQty |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeavesQty", wireType)
			}
			m.LeavesQty = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LeavesQty |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CumQty", wireType)
			}
			m.CumQty = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CumQty |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvgPx", wireType)
			}
			m.AvgPx = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AvgPx |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
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
				return ErrInvalidLengthOrdersExecutionReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersExecutionReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
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
				return ErrInvalidLengthOrdersExecutionReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersExecutionReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransactTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trailer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
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
				return ErrInvalidLengthOrdersExecutionReport
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersExecutionReport
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
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrdersExecutionReport
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
				return ErrInvalidLengthOrdersExecutionReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrdersExecutionReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrdersExecutionReport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrdersExecutionReport
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
func skipOrdersExecutionReport(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrdersExecutionReport
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
					return 0, ErrIntOverflowOrdersExecutionReport
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
					return 0, ErrIntOverflowOrdersExecutionReport
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
				return 0, ErrInvalidLengthOrdersExecutionReport
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrdersExecutionReport
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrdersExecutionReport
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrdersExecutionReport        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrdersExecutionReport          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrdersExecutionReport = fmt.Errorf("proto: unexpected end of group")
)
