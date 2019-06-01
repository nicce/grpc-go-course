// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Calculator struct {
	ValueOne             int32    `protobuf:"varint,1,opt,name=valueOne,proto3" json:"valueOne,omitempty"`
	ValueTwo             int32    `protobuf:"varint,2,opt,name=valueTwo,proto3" json:"valueTwo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Calculator) Reset()         { *m = Calculator{} }
func (m *Calculator) String() string { return proto.CompactTextString(m) }
func (*Calculator) ProtoMessage()    {}
func (*Calculator) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
}

func (m *Calculator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Calculator.Unmarshal(m, b)
}
func (m *Calculator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Calculator.Marshal(b, m, deterministic)
}
func (m *Calculator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Calculator.Merge(m, src)
}
func (m *Calculator) XXX_Size() int {
	return xxx_messageInfo_Calculator.Size(m)
}
func (m *Calculator) XXX_DiscardUnknown() {
	xxx_messageInfo_Calculator.DiscardUnknown(m)
}

var xxx_messageInfo_Calculator proto.InternalMessageInfo

func (m *Calculator) GetValueOne() int32 {
	if m != nil {
		return m.ValueOne
	}
	return 0
}

func (m *Calculator) GetValueTwo() int32 {
	if m != nil {
		return m.ValueTwo
	}
	return 0
}

type CalculatorRequest struct {
	Calculator           *Calculator `protobuf:"bytes,1,opt,name=calculator,proto3" json:"calculator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CalculatorRequest) Reset()         { *m = CalculatorRequest{} }
func (m *CalculatorRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatorRequest) ProtoMessage()    {}
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
}

func (m *CalculatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorRequest.Unmarshal(m, b)
}
func (m *CalculatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorRequest.Marshal(b, m, deterministic)
}
func (m *CalculatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorRequest.Merge(m, src)
}
func (m *CalculatorRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatorRequest.Size(m)
}
func (m *CalculatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorRequest proto.InternalMessageInfo

func (m *CalculatorRequest) GetCalculator() *Calculator {
	if m != nil {
		return m.Calculator
	}
	return nil
}

type CalculatorResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorResponse) Reset()         { *m = CalculatorResponse{} }
func (m *CalculatorResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorResponse) ProtoMessage()    {}
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{2}
}

func (m *CalculatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorResponse.Unmarshal(m, b)
}
func (m *CalculatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorResponse.Marshal(b, m, deterministic)
}
func (m *CalculatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorResponse.Merge(m, src)
}
func (m *CalculatorResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorResponse.Size(m)
}
func (m *CalculatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorResponse proto.InternalMessageInfo

func (m *CalculatorResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type PrimeNumber struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumber) Reset()         { *m = PrimeNumber{} }
func (m *PrimeNumber) String() string { return proto.CompactTextString(m) }
func (*PrimeNumber) ProtoMessage()    {}
func (*PrimeNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{3}
}

func (m *PrimeNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumber.Unmarshal(m, b)
}
func (m *PrimeNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumber.Marshal(b, m, deterministic)
}
func (m *PrimeNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumber.Merge(m, src)
}
func (m *PrimeNumber) XXX_Size() int {
	return xxx_messageInfo_PrimeNumber.Size(m)
}
func (m *PrimeNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumber.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumber proto.InternalMessageInfo

func (m *PrimeNumber) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PrimeNumberRequest struct {
	Number               *PrimeNumber `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PrimeNumberRequest) Reset()         { *m = PrimeNumberRequest{} }
func (m *PrimeNumberRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberRequest) ProtoMessage()    {}
func (*PrimeNumberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{4}
}

func (m *PrimeNumberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberRequest.Unmarshal(m, b)
}
func (m *PrimeNumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberRequest.Marshal(b, m, deterministic)
}
func (m *PrimeNumberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberRequest.Merge(m, src)
}
func (m *PrimeNumberRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberRequest.Size(m)
}
func (m *PrimeNumberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberRequest proto.InternalMessageInfo

func (m *PrimeNumberRequest) GetNumber() *PrimeNumber {
	if m != nil {
		return m.Number
	}
	return nil
}

type PrimeNumberResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberResponse) Reset()         { *m = PrimeNumberResponse{} }
func (m *PrimeNumberResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberResponse) ProtoMessage()    {}
func (*PrimeNumberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{5}
}

func (m *PrimeNumberResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberResponse.Unmarshal(m, b)
}
func (m *PrimeNumberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberResponse.Marshal(b, m, deterministic)
}
func (m *PrimeNumberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberResponse.Merge(m, src)
}
func (m *PrimeNumberResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberResponse.Size(m)
}
func (m *PrimeNumberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberResponse proto.InternalMessageInfo

func (m *PrimeNumberResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Calculator)(nil), "calculator.Calculator")
	proto.RegisterType((*CalculatorRequest)(nil), "calculator.CalculatorRequest")
	proto.RegisterType((*CalculatorResponse)(nil), "calculator.CalculatorResponse")
	proto.RegisterType((*PrimeNumber)(nil), "calculator.PrimeNumber")
	proto.RegisterType((*PrimeNumberRequest)(nil), "calculator.PrimeNumberRequest")
	proto.RegisterType((*PrimeNumberResponse)(nil), "calculator.PrimeNumberResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x4a, 0xc3, 0x50,
	0x10, 0x85, 0x8d, 0x62, 0x91, 0xa9, 0x08, 0x8e, 0x50, 0x43, 0xc1, 0x2a, 0x17, 0x84, 0x2e, 0xb4,
	0x95, 0x08, 0x3e, 0x80, 0x56, 0x37, 0x82, 0x4a, 0xdb, 0x95, 0x1b, 0x49, 0xc2, 0x2c, 0x02, 0x49,
	0x26, 0xde, 0x9f, 0xf6, 0xe5, 0x7c, 0x38, 0x31, 0xbf, 0xd3, 0x45, 0xdc, 0xdd, 0x93, 0x73, 0x72,
	0xee, 0x77, 0x87, 0x81, 0x69, 0x1c, 0xa6, 0xb1, 0x4b, 0x43, 0xcb, 0x7a, 0xde, 0x1d, 0x8b, 0x48,
	0x88, 0x59, 0xa1, 0xd9, 0x32, 0x42, 0xf7, 0x45, 0x2d, 0x00, 0x9e, 0x5a, 0x85, 0x63, 0x38, 0xda,
	0x84, 0xa9, 0xa3, 0xf7, 0x9c, 0x7c, 0xef, 0xca, 0x9b, 0x1e, 0x2e, 0x5b, 0xdd, 0x7a, 0xeb, 0x2d,
	0xfb, 0xfb, 0xc2, 0x5b, 0x6f, 0x59, 0xbd, 0xc2, 0x69, 0xd7, 0xb2, 0xa4, 0x6f, 0x47, 0xc6, 0xe2,
	0x03, 0x88, 0x8b, 0xca, 0xba, 0x61, 0x30, 0x9a, 0x09, 0x1a, 0xf1, 0x8b, 0x44, 0xba, 0x01, 0x94,
	0x65, 0xa6, 0xe0, 0xdc, 0x10, 0x8e, 0x60, 0xa0, 0xc9, 0xb8, 0xd4, 0xd6, 0x60, 0xb5, 0x52, 0xd7,
	0x30, 0xfc, 0xd0, 0x49, 0x46, 0x6f, 0x2e, 0x8b, 0x48, 0xff, 0xc5, 0xf2, 0xf2, 0xd4, 0xc4, 0x2a,
	0xa5, 0x9e, 0x01, 0x45, 0xac, 0x41, 0x9c, 0xef, 0xa4, 0x87, 0xc1, 0xb9, 0xc4, 0x93, 0xf9, 0xa6,
	0xe6, 0x16, 0xce, 0x76, 0x6a, 0xfe, 0x87, 0x0b, 0x7e, 0x3c, 0x39, 0x98, 0x15, 0xe9, 0x4d, 0x12,
	0x13, 0xbe, 0xc0, 0xc1, 0xca, 0x65, 0x78, 0xd1, 0x33, 0x8b, 0x8a, 0x6d, 0x3c, 0xe9, 0xb3, 0xab,
	0x3b, 0xd5, 0x1e, 0x7e, 0x81, 0x2f, 0x60, 0x16, 0x14, 0x73, 0x56, 0xb0, 0x49, 0x6c, 0xc2, 0x39,
	0x4e, 0xfa, 0x5e, 0x52, 0xb7, 0x5f, 0xf6, 0xfa, 0x4d, 0xfd, 0x9d, 0xf7, 0x78, 0xf2, 0x79, 0x2c,
	0x37, 0x29, 0x1a, 0x94, 0xfb, 0x73, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x87, 0x95, 0x63, 0x7d,
	0x6b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	// Unary
	Sum(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	// Server streaming
	PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberDecompositionClient interface {
	Recv() (*PrimeNumberResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberDecompositionClient) Recv() (*PrimeNumberResponse, error) {
	m := new(PrimeNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	// Unary
	Sum(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
	// Server streaming
	PrimeNumberDecomposition(*PrimeNumberRequest, CalculatorService_PrimeNumberDecompositionServer) error
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumberDecomposition(m, &calculatorServicePrimeNumberDecompositionServer{stream})
}

type CalculatorService_PrimeNumberDecompositionServer interface {
	Send(*PrimeNumberResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberDecompositionServer) Send(m *PrimeNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _CalculatorService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
