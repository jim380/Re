// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: re/mic/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryGetMarketIdentificationCodeRequest struct {
	MIC string `protobuf:"bytes,1,opt,name=MIC,proto3" json:"MIC,omitempty"`
}

func (m *QueryGetMarketIdentificationCodeRequest) Reset() {
	*m = QueryGetMarketIdentificationCodeRequest{}
}
func (m *QueryGetMarketIdentificationCodeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetMarketIdentificationCodeRequest) ProtoMessage()    {}
func (*QueryGetMarketIdentificationCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47859d0d82239e8d, []int{0}
}
func (m *QueryGetMarketIdentificationCodeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetMarketIdentificationCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetMarketIdentificationCodeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetMarketIdentificationCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetMarketIdentificationCodeRequest.Merge(m, src)
}
func (m *QueryGetMarketIdentificationCodeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetMarketIdentificationCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetMarketIdentificationCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetMarketIdentificationCodeRequest proto.InternalMessageInfo

func (m *QueryGetMarketIdentificationCodeRequest) GetMIC() string {
	if m != nil {
		return m.MIC
	}
	return ""
}

type QueryGetMarketIdentificationCodeResponse struct {
	MarketIdentificationCode MarketIdentificationCode `protobuf:"bytes,1,opt,name=MarketIdentificationCode,proto3" json:"MarketIdentificationCode"`
}

func (m *QueryGetMarketIdentificationCodeResponse) Reset() {
	*m = QueryGetMarketIdentificationCodeResponse{}
}
func (m *QueryGetMarketIdentificationCodeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetMarketIdentificationCodeResponse) ProtoMessage()    {}
func (*QueryGetMarketIdentificationCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47859d0d82239e8d, []int{1}
}
func (m *QueryGetMarketIdentificationCodeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetMarketIdentificationCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetMarketIdentificationCodeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetMarketIdentificationCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetMarketIdentificationCodeResponse.Merge(m, src)
}
func (m *QueryGetMarketIdentificationCodeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetMarketIdentificationCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetMarketIdentificationCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetMarketIdentificationCodeResponse proto.InternalMessageInfo

func (m *QueryGetMarketIdentificationCodeResponse) GetMarketIdentificationCode() MarketIdentificationCode {
	if m != nil {
		return m.MarketIdentificationCode
	}
	return MarketIdentificationCode{}
}

type QueryAllMarketIdentificationCodeRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllMarketIdentificationCodeRequest) Reset() {
	*m = QueryAllMarketIdentificationCodeRequest{}
}
func (m *QueryAllMarketIdentificationCodeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllMarketIdentificationCodeRequest) ProtoMessage()    {}
func (*QueryAllMarketIdentificationCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47859d0d82239e8d, []int{2}
}
func (m *QueryAllMarketIdentificationCodeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllMarketIdentificationCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllMarketIdentificationCodeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllMarketIdentificationCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllMarketIdentificationCodeRequest.Merge(m, src)
}
func (m *QueryAllMarketIdentificationCodeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllMarketIdentificationCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllMarketIdentificationCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllMarketIdentificationCodeRequest proto.InternalMessageInfo

func (m *QueryAllMarketIdentificationCodeRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllMarketIdentificationCodeResponse struct {
	MarketIdentificationCode []MarketIdentificationCode `protobuf:"bytes,1,rep,name=MarketIdentificationCode,proto3" json:"MarketIdentificationCode"`
	Pagination               *query.PageResponse        `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllMarketIdentificationCodeResponse) Reset() {
	*m = QueryAllMarketIdentificationCodeResponse{}
}
func (m *QueryAllMarketIdentificationCodeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllMarketIdentificationCodeResponse) ProtoMessage()    {}
func (*QueryAllMarketIdentificationCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47859d0d82239e8d, []int{3}
}
func (m *QueryAllMarketIdentificationCodeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllMarketIdentificationCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllMarketIdentificationCodeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllMarketIdentificationCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllMarketIdentificationCodeResponse.Merge(m, src)
}
func (m *QueryAllMarketIdentificationCodeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllMarketIdentificationCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllMarketIdentificationCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllMarketIdentificationCodeResponse proto.InternalMessageInfo

func (m *QueryAllMarketIdentificationCodeResponse) GetMarketIdentificationCode() []MarketIdentificationCode {
	if m != nil {
		return m.MarketIdentificationCode
	}
	return nil
}

func (m *QueryAllMarketIdentificationCodeResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetMarketIdentificationCodeRequest)(nil), "re.mic.QueryGetMarketIdentificationCodeRequest")
	proto.RegisterType((*QueryGetMarketIdentificationCodeResponse)(nil), "re.mic.QueryGetMarketIdentificationCodeResponse")
	proto.RegisterType((*QueryAllMarketIdentificationCodeRequest)(nil), "re.mic.QueryAllMarketIdentificationCodeRequest")
	proto.RegisterType((*QueryAllMarketIdentificationCodeResponse)(nil), "re.mic.QueryAllMarketIdentificationCodeResponse")
}

func init() { proto.RegisterFile("re/mic/v1beta1/query.proto", fileDescriptor_47859d0d82239e8d) }

var fileDescriptor_47859d0d82239e8d = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0xaa, 0xd3, 0x40,
	0x18, 0xcd, 0xf4, 0xea, 0x05, 0xc7, 0x8d, 0x0c, 0x2e, 0x2e, 0x51, 0x62, 0xc9, 0xc2, 0x7b, 0x29,
	0x32, 0xd3, 0x9f, 0x85, 0xa2, 0xab, 0xb6, 0x60, 0xe9, 0xa2, 0xa0, 0x59, 0xba, 0x29, 0x93, 0x74,
	0x8c, 0xa3, 0x49, 0x26, 0xcd, 0x4c, 0xc5, 0x22, 0x6e, 0x7c, 0x00, 0x11, 0x7c, 0x15, 0xf1, 0x19,
	0xba, 0x2c, 0xb8, 0xd0, 0x95, 0x48, 0xeb, 0x3b, 0xb8, 0x95, 0xcc, 0xa4, 0xda, 0x16, 0x62, 0xb2,
	0xb8, 0xbb, 0x21, 0x73, 0xce, 0x99, 0x73, 0xce, 0xf7, 0x05, 0xda, 0x19, 0x23, 0x31, 0x0f, 0xc8,
	0xeb, 0x8e, 0xcf, 0x14, 0xed, 0x90, 0xf9, 0x82, 0x65, 0x4b, 0x9c, 0x66, 0x42, 0x09, 0x74, 0x9a,
	0x31, 0x1c, 0xf3, 0xc0, 0xbe, 0x19, 0x8a, 0x50, 0xe8, 0x4f, 0x24, 0x3f, 0x99, 0x5b, 0xfb, 0x76,
	0x28, 0x44, 0x18, 0x31, 0x42, 0x53, 0x4e, 0x68, 0x92, 0x08, 0x45, 0x15, 0x17, 0x89, 0x2c, 0x6e,
	0x5b, 0x81, 0x90, 0xb1, 0x90, 0xc4, 0xa7, 0x92, 0x19, 0xd1, 0xbf, 0x4f, 0xa4, 0x34, 0xe4, 0x89,
	0x06, 0x17, 0x58, 0x72, 0xe4, 0x21, 0xa6, 0xd9, 0x2b, 0xa6, 0xa6, 0x7c, 0xc6, 0x12, 0xc5, 0x9f,
	0xf3, 0x40, 0x63, 0xa7, 0x81, 0x98, 0x31, 0x43, 0x70, 0x1f, 0xc1, 0xf3, 0xa7, 0xb9, 0xe4, 0x88,
	0xa9, 0x89, 0xc6, 0x8e, 0x0f, 0xa0, 0x43, 0x31, 0x63, 0x1e, 0x9b, 0x2f, 0x98, 0x54, 0xe8, 0x06,
	0x3c, 0x99, 0x8c, 0x87, 0x67, 0xa0, 0x09, 0x2e, 0xae, 0x79, 0xf9, 0xd1, 0xfd, 0x00, 0xe0, 0x45,
	0x35, 0x5b, 0xa6, 0x22, 0x91, 0x0c, 0xf9, 0xf0, 0xac, 0x0c, 0xa3, 0x35, 0xaf, 0x77, 0x9b, 0xd8,
	0xb4, 0x84, 0xcb, 0x70, 0x83, 0x2b, 0xab, 0x1f, 0x77, 0x2c, 0xaf, 0x54, 0xc7, 0x9d, 0x17, 0x69,
	0xfa, 0x51, 0x54, 0x95, 0xe6, 0x31, 0x84, 0xff, 0xda, 0x2b, 0x0c, 0xdc, 0xc5, 0xa6, 0x6a, 0x9c,
	0x57, 0x8d, 0xcd, 0xfc, 0x8a, 0x26, 0xf1, 0x13, 0x1a, 0xee, 0xb8, 0xde, 0x1e, 0xd3, 0xfd, 0xb6,
	0xeb, 0xe0, 0xbf, 0x6f, 0xd6, 0xea, 0xe0, 0xe4, 0x32, 0x3a, 0x40, 0xa3, 0x83, 0x60, 0x0d, 0x1d,
	0xec, 0xbc, 0x32, 0x98, 0x31, 0xb8, 0x9f, 0xac, 0xfb, 0xbb, 0x01, 0xaf, 0xea, 0x64, 0xe8, 0x0b,
	0x28, 0xf7, 0x8d, 0xc8, 0xce, 0x71, 0xcd, 0x3d, 0xb2, 0xdb, 0xf5, 0x09, 0xc6, 0x95, 0x7b, 0xff,
	0xfd, 0xd7, 0x5f, 0x9f, 0x1a, 0x1d, 0x44, 0xc8, 0x4b, 0x1e, 0xf7, 0x1e, 0xb4, 0x89, 0x67, 0xb6,
	0xbc, 0x7c, 0xbb, 0xc9, 0xdb, 0xc9, 0x78, 0xf8, 0x0e, 0x7d, 0x06, 0xf0, 0x56, 0x99, 0x7a, 0x3f,
	0x8a, 0x8e, 0xbc, 0x57, 0x6f, 0xcd, 0x91, 0xf7, 0x1a, 0x23, 0x77, 0xbb, 0xda, 0xfb, 0x3d, 0xd4,
	0xaa, 0xed, 0x5d, 0x0e, 0x1e, 0xae, 0x36, 0x0e, 0x58, 0x6f, 0x1c, 0xf0, 0x73, 0xe3, 0x80, 0x8f,
	0x5b, 0xc7, 0x5a, 0x6f, 0x1d, 0xeb, 0xfb, 0xd6, 0xb1, 0x9e, 0x35, 0x43, 0xae, 0x5e, 0x2c, 0x7c,
	0x1c, 0x88, 0x78, 0x4f, 0xef, 0x8d, 0x56, 0x54, 0xcb, 0x94, 0x49, 0xff, 0x54, 0xff, 0xd7, 0xbd,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x43, 0xef, 0x19, 0x8e, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a list of MarketIdentificationCode items.
	MarketIdentificationCode(ctx context.Context, in *QueryGetMarketIdentificationCodeRequest, opts ...grpc.CallOption) (*QueryGetMarketIdentificationCodeResponse, error)
	MarketIdentificationCodeAll(ctx context.Context, in *QueryAllMarketIdentificationCodeRequest, opts ...grpc.CallOption) (*QueryAllMarketIdentificationCodeResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) MarketIdentificationCode(ctx context.Context, in *QueryGetMarketIdentificationCodeRequest, opts ...grpc.CallOption) (*QueryGetMarketIdentificationCodeResponse, error) {
	out := new(QueryGetMarketIdentificationCodeResponse)
	err := c.cc.Invoke(ctx, "/re.mic.Query/MarketIdentificationCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MarketIdentificationCodeAll(ctx context.Context, in *QueryAllMarketIdentificationCodeRequest, opts ...grpc.CallOption) (*QueryAllMarketIdentificationCodeResponse, error) {
	out := new(QueryAllMarketIdentificationCodeResponse)
	err := c.cc.Invoke(ctx, "/re.mic.Query/MarketIdentificationCodeAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a list of MarketIdentificationCode items.
	MarketIdentificationCode(context.Context, *QueryGetMarketIdentificationCodeRequest) (*QueryGetMarketIdentificationCodeResponse, error)
	MarketIdentificationCodeAll(context.Context, *QueryAllMarketIdentificationCodeRequest) (*QueryAllMarketIdentificationCodeResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) MarketIdentificationCode(ctx context.Context, req *QueryGetMarketIdentificationCodeRequest) (*QueryGetMarketIdentificationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketIdentificationCode not implemented")
}
func (*UnimplementedQueryServer) MarketIdentificationCodeAll(ctx context.Context, req *QueryAllMarketIdentificationCodeRequest) (*QueryAllMarketIdentificationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketIdentificationCodeAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_MarketIdentificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetMarketIdentificationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MarketIdentificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/re.mic.Query/MarketIdentificationCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MarketIdentificationCode(ctx, req.(*QueryGetMarketIdentificationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_MarketIdentificationCodeAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllMarketIdentificationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MarketIdentificationCodeAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/re.mic.Query/MarketIdentificationCodeAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MarketIdentificationCodeAll(ctx, req.(*QueryAllMarketIdentificationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "re.mic.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MarketIdentificationCode",
			Handler:    _Query_MarketIdentificationCode_Handler,
		},
		{
			MethodName: "MarketIdentificationCodeAll",
			Handler:    _Query_MarketIdentificationCodeAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "re/mic/v1beta1/query.proto",
}

func (m *QueryGetMarketIdentificationCodeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetMarketIdentificationCodeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetMarketIdentificationCodeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MIC) > 0 {
		i -= len(m.MIC)
		copy(dAtA[i:], m.MIC)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.MIC)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetMarketIdentificationCodeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetMarketIdentificationCodeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetMarketIdentificationCodeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.MarketIdentificationCode.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllMarketIdentificationCodeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllMarketIdentificationCodeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllMarketIdentificationCodeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllMarketIdentificationCodeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllMarketIdentificationCodeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllMarketIdentificationCodeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.MarketIdentificationCode) > 0 {
		for iNdEx := len(m.MarketIdentificationCode) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MarketIdentificationCode[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetMarketIdentificationCodeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MIC)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetMarketIdentificationCodeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MarketIdentificationCode.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllMarketIdentificationCodeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllMarketIdentificationCodeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MarketIdentificationCode) > 0 {
		for _, e := range m.MarketIdentificationCode {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetMarketIdentificationCodeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetMarketIdentificationCodeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetMarketIdentificationCodeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MIC", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MIC = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetMarketIdentificationCodeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetMarketIdentificationCodeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetMarketIdentificationCodeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketIdentificationCode", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MarketIdentificationCode.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllMarketIdentificationCodeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllMarketIdentificationCodeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllMarketIdentificationCodeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllMarketIdentificationCodeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllMarketIdentificationCodeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllMarketIdentificationCodeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketIdentificationCode", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketIdentificationCode = append(m.MarketIdentificationCode, MarketIdentificationCode{})
			if err := m.MarketIdentificationCode[len(m.MarketIdentificationCode)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
