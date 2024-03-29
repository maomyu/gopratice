// Code generated by protoc-gen-go. DO NOT EDIT.
// source: p.proto

package p

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CallRequest struct {
	Greeting             string            `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	Infos                map[string]string `protobuf:"bytes,2,rep,name=infos,proto3" json:"infos,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CallRequest) Reset()         { *m = CallRequest{} }
func (m *CallRequest) String() string { return proto.CompactTextString(m) }
func (*CallRequest) ProtoMessage()    {}
func (*CallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe44bf90c1c2369, []int{0}
}

func (m *CallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallRequest.Unmarshal(m, b)
}
func (m *CallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallRequest.Marshal(b, m, deterministic)
}
func (m *CallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallRequest.Merge(m, src)
}
func (m *CallRequest) XXX_Size() int {
	return xxx_messageInfo_CallRequest.Size(m)
}
func (m *CallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallRequest proto.InternalMessageInfo

func (m *CallRequest) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func (m *CallRequest) GetInfos() map[string]string {
	if m != nil {
		return m.Infos
	}
	return nil
}

type CallResponse struct {
	Reply                string     `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	Details              []*any.Any `protobuf:"bytes,2,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CallResponse) Reset()         { *m = CallResponse{} }
func (m *CallResponse) String() string { return proto.CompactTextString(m) }
func (*CallResponse) ProtoMessage()    {}
func (*CallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe44bf90c1c2369, []int{1}
}

func (m *CallResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallResponse.Unmarshal(m, b)
}
func (m *CallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallResponse.Marshal(b, m, deterministic)
}
func (m *CallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallResponse.Merge(m, src)
}
func (m *CallResponse) XXX_Size() int {
	return xxx_messageInfo_CallResponse.Size(m)
}
func (m *CallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CallResponse proto.InternalMessageInfo

func (m *CallResponse) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func (m *CallResponse) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

type Res struct {
	Reply                string   `protobuf:"bytes,4,opt,name=reply,proto3" json:"reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Res) Reset()         { *m = Res{} }
func (m *Res) String() string { return proto.CompactTextString(m) }
func (*Res) ProtoMessage()    {}
func (*Res) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe44bf90c1c2369, []int{2}
}

func (m *Res) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Res.Unmarshal(m, b)
}
func (m *Res) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Res.Marshal(b, m, deterministic)
}
func (m *Res) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Res.Merge(m, src)
}
func (m *Res) XXX_Size() int {
	return xxx_messageInfo_Res.Size(m)
}
func (m *Res) XXX_DiscardUnknown() {
	xxx_messageInfo_Res.DiscardUnknown(m)
}

var xxx_messageInfo_Res proto.InternalMessageInfo

func (m *Res) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func init() {
	proto.RegisterType((*CallRequest)(nil), "CallRequest")
	proto.RegisterMapType((map[string]string)(nil), "CallRequest.InfosEntry")
	proto.RegisterType((*CallResponse)(nil), "CallResponse")
	proto.RegisterType((*Res)(nil), "Res")
}

func init() { proto.RegisterFile("p.proto", fileDescriptor_abe44bf90c1c2369) }

var fileDescriptor_abe44bf90c1c2369 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x49, 0x42, 0x09, 0xbd, 0x16, 0x09, 0x59, 0x95, 0x88, 0xc2, 0x52, 0x65, 0xca, 0x82,
	0x87, 0x32, 0x50, 0xb1, 0x21, 0xc4, 0xc0, 0xea, 0xf2, 0x02, 0x2e, 0x5c, 0xa3, 0x08, 0xcb, 0x36,
	0xb6, 0x53, 0xc9, 0xaf, 0xc1, 0x13, 0x23, 0xd7, 0xa4, 0xf1, 0x76, 0xff, 0xef, 0xff, 0x3e, 0xdf,
	0x1d, 0x94, 0x9a, 0x6a, 0xa3, 0x9c, 0xaa, 0xe7, 0x5c, 0xfa, 0x58, 0x36, 0xbf, 0x19, 0x2c, 0x5e,
	0xb9, 0x10, 0x0c, 0x7f, 0x06, 0xb4, 0x8e, 0xd4, 0x70, 0xdd, 0x19, 0x44, 0xd7, 0xcb, 0xae, 0xca,
	0xd6, 0x59, 0x3b, 0x67, 0x67, 0x4d, 0x1e, 0x60, 0xd6, 0xcb, 0x83, 0xb2, 0x55, 0xbe, 0x2e, 0xda,
	0xc5, 0xe6, 0x8e, 0x26, 0x8d, 0xf4, 0x3d, 0xbc, 0xbc, 0x49, 0x67, 0x3c, 0x8b, 0xa9, 0x7a, 0x0b,
	0x30, 0x99, 0xe4, 0x16, 0x8a, 0x6f, 0xf4, 0xff, 0xcc, 0x50, 0x92, 0x15, 0xcc, 0x8e, 0x5c, 0x0c,
	0x58, 0xe5, 0x27, 0x2f, 0x8a, 0xe7, 0x7c, 0x9b, 0x35, 0x1f, 0xb0, 0x8c, 0x68, 0xab, 0x95, 0xb4,
	0x18, 0x92, 0x06, 0xb5, 0x18, 0xbb, 0xa3, 0x20, 0x14, 0xca, 0x2f, 0x74, 0xbc, 0x17, 0xe3, 0x40,
	0x2b, 0xda, 0x29, 0xd5, 0x09, 0x8c, 0xab, 0xed, 0x87, 0x03, 0x7d, 0x91, 0x9e, 0x8d, 0xa1, 0xe6,
	0x1e, 0x0a, 0x86, 0x76, 0x82, 0x5d, 0x26, 0xb0, 0xcd, 0x53, 0x3c, 0xc3, 0x0e, 0xcd, 0xb1, 0xff,
	0x44, 0xd2, 0x42, 0xb9, 0xe3, 0x3e, 0x38, 0x64, 0x99, 0xae, 0x59, 0xdf, 0xd0, 0x74, 0xb2, 0xe6,
	0x62, 0x7f, 0x75, 0xfa, 0xec, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x80, 0xc4, 0x30, 0xb6, 0x5d,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CallServiceClient is the client API for CallService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CallServiceClient interface {
	SayCall(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error)
}

type callServiceClient struct {
	cc *grpc.ClientConn
}

func NewCallServiceClient(cc *grpc.ClientConn) CallServiceClient {
	return &callServiceClient{cc}
}

func (c *callServiceClient) SayCall(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := c.cc.Invoke(ctx, "/CallService/SayCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CallServiceServer is the server API for CallService service.
type CallServiceServer interface {
	SayCall(context.Context, *CallRequest) (*CallResponse, error)
}

// UnimplementedCallServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCallServiceServer struct {
}

func (*UnimplementedCallServiceServer) SayCall(ctx context.Context, req *CallRequest) (*CallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayCall not implemented")
}

func RegisterCallServiceServer(s *grpc.Server, srv CallServiceServer) {
	s.RegisterService(&_CallService_serviceDesc, srv)
}

func _CallService_SayCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).SayCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CallService/SayCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).SayCall(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CallService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CallService",
	HandlerType: (*CallServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayCall",
			Handler:    _CallService_SayCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "p.proto",
}
