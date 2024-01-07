// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: re/oracle/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgCosmoshubTxs struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	OracleID string `protobuf:"bytes,2,opt,name=oracleID,proto3" json:"oracleID,omitempty"`
	Address  string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgCosmoshubTxs) Reset()         { *m = MsgCosmoshubTxs{} }
func (m *MsgCosmoshubTxs) String() string { return proto.CompactTextString(m) }
func (*MsgCosmoshubTxs) ProtoMessage()    {}
func (*MsgCosmoshubTxs) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e5b0255adc1fd5, []int{0}
}
func (m *MsgCosmoshubTxs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCosmoshubTxs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCosmoshubTxs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCosmoshubTxs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCosmoshubTxs.Merge(m, src)
}
func (m *MsgCosmoshubTxs) XXX_Size() int {
	return m.Size()
}
func (m *MsgCosmoshubTxs) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCosmoshubTxs.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCosmoshubTxs proto.InternalMessageInfo

func (m *MsgCosmoshubTxs) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCosmoshubTxs) GetOracleID() string {
	if m != nil {
		return m.OracleID
	}
	return ""
}

func (m *MsgCosmoshubTxs) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MsgCosmoshubTxsResponse struct {
}

func (m *MsgCosmoshubTxsResponse) Reset()         { *m = MsgCosmoshubTxsResponse{} }
func (m *MsgCosmoshubTxsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCosmoshubTxsResponse) ProtoMessage()    {}
func (*MsgCosmoshubTxsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e5b0255adc1fd5, []int{1}
}
func (m *MsgCosmoshubTxsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCosmoshubTxsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCosmoshubTxsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCosmoshubTxsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCosmoshubTxsResponse.Merge(m, src)
}
func (m *MsgCosmoshubTxsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCosmoshubTxsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCosmoshubTxsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCosmoshubTxsResponse proto.InternalMessageInfo

type MsgBitcoinTxs struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	OracleID string `protobuf:"bytes,2,opt,name=oracleID,proto3" json:"oracleID,omitempty"`
	Address  string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgBitcoinTxs) Reset()         { *m = MsgBitcoinTxs{} }
func (m *MsgBitcoinTxs) String() string { return proto.CompactTextString(m) }
func (*MsgBitcoinTxs) ProtoMessage()    {}
func (*MsgBitcoinTxs) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e5b0255adc1fd5, []int{2}
}
func (m *MsgBitcoinTxs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBitcoinTxs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBitcoinTxs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBitcoinTxs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBitcoinTxs.Merge(m, src)
}
func (m *MsgBitcoinTxs) XXX_Size() int {
	return m.Size()
}
func (m *MsgBitcoinTxs) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBitcoinTxs.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBitcoinTxs proto.InternalMessageInfo

func (m *MsgBitcoinTxs) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgBitcoinTxs) GetOracleID() string {
	if m != nil {
		return m.OracleID
	}
	return ""
}

func (m *MsgBitcoinTxs) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MsgBitcoinTxsResponse struct {
}

func (m *MsgBitcoinTxsResponse) Reset()         { *m = MsgBitcoinTxsResponse{} }
func (m *MsgBitcoinTxsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgBitcoinTxsResponse) ProtoMessage()    {}
func (*MsgBitcoinTxsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e5b0255adc1fd5, []int{3}
}
func (m *MsgBitcoinTxsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBitcoinTxsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBitcoinTxsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBitcoinTxsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBitcoinTxsResponse.Merge(m, src)
}
func (m *MsgBitcoinTxsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgBitcoinTxsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBitcoinTxsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBitcoinTxsResponse proto.InternalMessageInfo

type MsgEthereumTxs struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	OracleID string `protobuf:"bytes,2,opt,name=oracleID,proto3" json:"oracleID,omitempty"`
	Address  string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgEthereumTxs) Reset()         { *m = MsgEthereumTxs{} }
func (m *MsgEthereumTxs) String() string { return proto.CompactTextString(m) }
func (*MsgEthereumTxs) ProtoMessage()    {}
func (*MsgEthereumTxs) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e5b0255adc1fd5, []int{4}
}
func (m *MsgEthereumTxs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEthereumTxs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEthereumTxs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEthereumTxs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEthereumTxs.Merge(m, src)
}
func (m *MsgEthereumTxs) XXX_Size() int {
	return m.Size()
}
func (m *MsgEthereumTxs) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEthereumTxs.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEthereumTxs proto.InternalMessageInfo

func (m *MsgEthereumTxs) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgEthereumTxs) GetOracleID() string {
	if m != nil {
		return m.OracleID
	}
	return ""
}

func (m *MsgEthereumTxs) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MsgEthereumTxsResponse struct {
}

func (m *MsgEthereumTxsResponse) Reset()         { *m = MsgEthereumTxsResponse{} }
func (m *MsgEthereumTxsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgEthereumTxsResponse) ProtoMessage()    {}
func (*MsgEthereumTxsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e5b0255adc1fd5, []int{5}
}
func (m *MsgEthereumTxsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEthereumTxsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEthereumTxsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEthereumTxsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEthereumTxsResponse.Merge(m, src)
}
func (m *MsgEthereumTxsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgEthereumTxsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEthereumTxsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEthereumTxsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCosmoshubTxs)(nil), "re.oracle.MsgCosmoshubTxs")
	proto.RegisterType((*MsgCosmoshubTxsResponse)(nil), "re.oracle.MsgCosmoshubTxsResponse")
	proto.RegisterType((*MsgBitcoinTxs)(nil), "re.oracle.MsgBitcoinTxs")
	proto.RegisterType((*MsgBitcoinTxsResponse)(nil), "re.oracle.MsgBitcoinTxsResponse")
	proto.RegisterType((*MsgEthereumTxs)(nil), "re.oracle.MsgEthereumTxs")
	proto.RegisterType((*MsgEthereumTxsResponse)(nil), "re.oracle.MsgEthereumTxsResponse")
}

func init() { proto.RegisterFile("re/oracle/v1beta1/tx.proto", fileDescriptor_61e5b0255adc1fd5) }

var fileDescriptor_61e5b0255adc1fd5 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0x7b, 0x16, 0xd4, 0xbe, 0xfe, 0x83, 0x80, 0x36, 0xbd, 0xe1, 0xa8, 0x71, 0x71, 0xca,
	0x59, 0xbb, 0xb8, 0xb8, 0x54, 0x05, 0x45, 0xe2, 0x50, 0x9c, 0x5c, 0x34, 0x49, 0x5f, 0xd2, 0x88,
	0xe9, 0x95, 0x7b, 0xaf, 0x52, 0xbf, 0x85, 0x1f, 0xcb, 0xb1, 0xa3, 0xa3, 0xb4, 0xab, 0x1f, 0x42,
	0x6c, 0x4d, 0x4c, 0x0a, 0x19, 0x3b, 0xbe, 0xfc, 0x78, 0x9e, 0x1f, 0x3c, 0x77, 0xc0, 0x35, 0x4a,
	0xa5, 0xfd, 0xf0, 0x05, 0xe5, 0x6b, 0x2b, 0x40, 0xe3, 0xb7, 0xa4, 0x19, 0xbb, 0x43, 0xad, 0x8c,
	0xb2, 0x6a, 0x1a, 0xdd, 0x05, 0x73, 0x7c, 0xd8, 0xf3, 0x28, 0xba, 0x50, 0x94, 0x28, 0xea, 0x8f,
	0x82, 0xfb, 0x31, 0x59, 0x36, 0x6c, 0x84, 0x1a, 0x7d, 0xa3, 0xb4, 0xcd, 0x9a, 0xec, 0xb8, 0xd6,
	0x4d, 0x4f, 0x8b, 0xc3, 0xe6, 0x22, 0x76, 0x73, 0x69, 0xaf, 0xcd, 0x51, 0x76, 0xff, 0xa6, 0xfc,
	0x5e, 0x4f, 0x23, 0x91, 0x5d, 0x5d, 0xa4, 0xfe, 0x4e, 0xa7, 0x01, 0xf5, 0x25, 0x45, 0x17, 0x69,
	0xa8, 0x06, 0x84, 0xce, 0x23, 0xec, 0x78, 0x14, 0x75, 0x62, 0x13, 0xaa, 0x78, 0xb0, 0x0a, 0x77,
	0x1d, 0xf6, 0x0b, 0x82, 0xcc, 0xfc, 0x04, 0xbb, 0x1e, 0x45, 0x57, 0xa6, 0x8f, 0x1a, 0x47, 0xc9,
	0x2a, 0xd4, 0x36, 0x1c, 0x14, 0x0d, 0xa9, 0xfb, 0xf4, 0x9b, 0x41, 0xd5, 0xa3, 0xc8, 0xba, 0x83,
	0xed, 0xc2, 0xf0, 0xdc, 0xcd, 0xde, 0xc5, 0x5d, 0x5a, 0x8c, 0x3b, 0xe5, 0x2c, 0xed, 0xb5, 0xae,
	0x01, 0xf2, 0x53, 0x16, 0x13, 0xff, 0x84, 0x37, 0xcb, 0x48, 0xd6, 0x74, 0x0b, 0x5b, 0xf9, 0x69,
	0x1a, 0xc5, 0x40, 0x0e, 0xf1, 0xc3, 0x52, 0x94, 0x96, 0x75, 0xce, 0x3f, 0xa6, 0x82, 0x4d, 0xa6,
	0x82, 0x7d, 0x4d, 0x05, 0x7b, 0x9f, 0x89, 0xca, 0x64, 0x26, 0x2a, 0x9f, 0x33, 0x51, 0x79, 0x38,
	0x8a, 0x62, 0xd3, 0x1f, 0x05, 0x6e, 0xa8, 0x12, 0xf9, 0x1c, 0x27, 0xed, 0xb3, 0x13, 0xd9, 0x45,
	0x39, 0x4e, 0xff, 0xad, 0x79, 0x1b, 0x22, 0x05, 0xeb, 0xf3, 0x3f, 0xdb, 0xfe, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x85, 0xb2, 0x65, 0x5d, 0xd1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CosmoshubTxs(ctx context.Context, in *MsgCosmoshubTxs, opts ...grpc.CallOption) (*MsgCosmoshubTxsResponse, error)
	BitcoinTxs(ctx context.Context, in *MsgBitcoinTxs, opts ...grpc.CallOption) (*MsgBitcoinTxsResponse, error)
	EthereumTxs(ctx context.Context, in *MsgEthereumTxs, opts ...grpc.CallOption) (*MsgEthereumTxsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CosmoshubTxs(ctx context.Context, in *MsgCosmoshubTxs, opts ...grpc.CallOption) (*MsgCosmoshubTxsResponse, error) {
	out := new(MsgCosmoshubTxsResponse)
	err := c.cc.Invoke(ctx, "/re.oracle.Msg/CosmoshubTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BitcoinTxs(ctx context.Context, in *MsgBitcoinTxs, opts ...grpc.CallOption) (*MsgBitcoinTxsResponse, error) {
	out := new(MsgBitcoinTxsResponse)
	err := c.cc.Invoke(ctx, "/re.oracle.Msg/BitcoinTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) EthereumTxs(ctx context.Context, in *MsgEthereumTxs, opts ...grpc.CallOption) (*MsgEthereumTxsResponse, error) {
	out := new(MsgEthereumTxsResponse)
	err := c.cc.Invoke(ctx, "/re.oracle.Msg/EthereumTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CosmoshubTxs(context.Context, *MsgCosmoshubTxs) (*MsgCosmoshubTxsResponse, error)
	BitcoinTxs(context.Context, *MsgBitcoinTxs) (*MsgBitcoinTxsResponse, error)
	EthereumTxs(context.Context, *MsgEthereumTxs) (*MsgEthereumTxsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CosmoshubTxs(ctx context.Context, req *MsgCosmoshubTxs) (*MsgCosmoshubTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CosmoshubTxs not implemented")
}
func (*UnimplementedMsgServer) BitcoinTxs(ctx context.Context, req *MsgBitcoinTxs) (*MsgBitcoinTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BitcoinTxs not implemented")
}
func (*UnimplementedMsgServer) EthereumTxs(ctx context.Context, req *MsgEthereumTxs) (*MsgEthereumTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EthereumTxs not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CosmoshubTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCosmoshubTxs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CosmoshubTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/re.oracle.Msg/CosmoshubTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CosmoshubTxs(ctx, req.(*MsgCosmoshubTxs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BitcoinTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBitcoinTxs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BitcoinTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/re.oracle.Msg/BitcoinTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BitcoinTxs(ctx, req.(*MsgBitcoinTxs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_EthereumTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEthereumTxs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EthereumTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/re.oracle.Msg/EthereumTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EthereumTxs(ctx, req.(*MsgEthereumTxs))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "re.oracle.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CosmoshubTxs",
			Handler:    _Msg_CosmoshubTxs_Handler,
		},
		{
			MethodName: "BitcoinTxs",
			Handler:    _Msg_BitcoinTxs_Handler,
		},
		{
			MethodName: "EthereumTxs",
			Handler:    _Msg_EthereumTxs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "re/oracle/v1beta1/tx.proto",
}

func (m *MsgCosmoshubTxs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCosmoshubTxs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCosmoshubTxs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OracleID) > 0 {
		i -= len(m.OracleID)
		copy(dAtA[i:], m.OracleID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OracleID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCosmoshubTxsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCosmoshubTxsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCosmoshubTxsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgBitcoinTxs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBitcoinTxs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBitcoinTxs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OracleID) > 0 {
		i -= len(m.OracleID)
		copy(dAtA[i:], m.OracleID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OracleID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBitcoinTxsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBitcoinTxsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBitcoinTxsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgEthereumTxs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEthereumTxs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEthereumTxs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OracleID) > 0 {
		i -= len(m.OracleID)
		copy(dAtA[i:], m.OracleID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OracleID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgEthereumTxsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEthereumTxsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEthereumTxsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCosmoshubTxs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.OracleID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCosmoshubTxsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgBitcoinTxs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.OracleID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgBitcoinTxsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgEthereumTxs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.OracleID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgEthereumTxsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCosmoshubTxs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCosmoshubTxs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCosmoshubTxs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCosmoshubTxsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCosmoshubTxsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCosmoshubTxsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgBitcoinTxs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBitcoinTxs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBitcoinTxs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgBitcoinTxsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBitcoinTxsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBitcoinTxsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgEthereumTxs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgEthereumTxs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEthereumTxs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgEthereumTxsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgEthereumTxsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEthereumTxsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
