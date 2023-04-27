package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64b0f1bc979aed9f, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type GetLink struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLink) Reset()         { *m = GetLink{} }
func (m *GetLink) String() string { return proto.CompactTextString(m) }
func (*GetLink) ProtoMessage()    {}
func (*GetLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_64b0f1bc979aed9f, []int{1}
}

func (m *GetLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLink.Unmarshal(m, b)
}
func (m *GetLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLink.Marshal(b, m, deterministic)
}
func (m *GetLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLink.Merge(m, src)
}
func (m *GetLink) XXX_Size() int {
	return xxx_messageInfo_GetLink.Size(m)
}
func (m *GetLink) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLink.DiscardUnknown(m)
}

var xxx_messageInfo_GetLink proto.InternalMessageInfo

func (m *GetLink) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Response struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Shorturl             string   `protobuf:"bytes,2,opt,name=shorturl,proto3" json:"shorturl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_64b0f1bc979aed9f, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Response) GetShorturl() string {
	if m != nil {
		return m.Shorturl
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*GetLink)(nil), "GetLink")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("Message.proto", fileDescriptor_64b0f1bc979aed9f) }

var fileDescriptor_64b0f1bc979aed9f = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xf5, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe4, 0xe2, 0x75, 0x2e, 0x4a,
	0x4d, 0x2c, 0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe0, 0x62, 0x2e, 0x2d,
	0xca, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x95, 0xa4, 0xb9, 0xd8, 0xdd, 0x53,
	0x4b, 0x7c, 0x32, 0xf3, 0xb2, 0xb1, 0x48, 0x5a, 0x70, 0x71, 0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7,
	0x15, 0xa7, 0x62, 0xca, 0x0a, 0x49, 0x71, 0x71, 0x14, 0x67, 0xe4, 0x17, 0x95, 0x80, 0x84, 0x99,
	0xc0, 0xc2, 0x70, 0xbe, 0x51, 0x00, 0x17, 0x67, 0x30, 0x88, 0x0d, 0x36, 0x58, 0x95, 0x8b, 0x0d,
	0xe2, 0x0c, 0x21, 0x3e, 0x3d, 0x14, 0xf7, 0x48, 0x71, 0xea, 0xc1, 0xcc, 0x57, 0x62, 0x10, 0x92,
	0xe1, 0x62, 0x76, 0x4f, 0x2d, 0x11, 0xe2, 0xd0, 0x83, 0x3a, 0x08, 0x45, 0xd6, 0x89, 0x2b, 0x0a,
	0x6e, 0x7a, 0x12, 0x1b, 0xd8, 0x7b, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x61, 0xd8,
	0x6e, 0xef, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShortLinkClient is the client API for ShortLink service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShortLinkClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Response, error)
	Get(ctx context.Context, in *GetLink, opts ...grpc.CallOption) (*Response, error)
}

type shortLinkClient struct {
	cc *grpc.ClientConn
}

func NewShortLinkClient(cc *grpc.ClientConn) ShortLinkClient {
	return &shortLinkClient{cc}
}

func (c *shortLinkClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ShortLink/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortLinkClient) Get(ctx context.Context, in *GetLink, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ShortLink/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortLinkServer is the server API for ShortLink service.
type ShortLinkServer interface {
	Create(context.Context, *CreateRequest) (*Response, error)
	Get(context.Context, *GetLink) (*Response, error)
}

// UnimplementedShortLinkServer can be embedded to have forward compatible implementations.
type UnimplementedShortLinkServer struct {
}

func (*UnimplementedShortLinkServer) Create(ctx context.Context, req *CreateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedShortLinkServer) Get(ctx context.Context, req *GetLink) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterShortLinkServer(s *grpc.Server, srv ShortLinkServer) {
	s.RegisterService(&_ShortLink_serviceDesc, srv)
}

func _ShortLink_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortLinkServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShortLink/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortLinkServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortLink_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLink)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortLinkServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShortLink/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortLinkServer).Get(ctx, req.(*GetLink))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShortLink_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShortLink",
	HandlerType: (*ShortLinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ShortLink_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ShortLink_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Message.proto",
}
