// Code generated by protoc-gen-go. DO NOT EDIT.
// source: urls.proto

/*
Package v2beta is a generated protocol buffer package.

It is generated from these files:
	urls.proto

It has these top-level messages:
	Url
	GetUrlsRequest
	GetUrlsResponse
	GetUrlRequest
	GetUrlResponse
	CreateUrlRequest
	CreateUrlResponse
*/
package v2beta

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Url struct {
	Key      string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Link     string `protobuf:"bytes,2,opt,name=link" json:"link,omitempty"`
	Creation string `protobuf:"bytes,3,opt,name=creation" json:"creation,omitempty"`
}

func (m *Url) Reset()                    { *m = Url{} }
func (m *Url) String() string            { return proto.CompactTextString(m) }
func (*Url) ProtoMessage()               {}
func (*Url) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Url) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Url) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Url) GetCreation() string {
	if m != nil {
		return m.Creation
	}
	return ""
}

type GetUrlsRequest struct {
}

func (m *GetUrlsRequest) Reset()                    { *m = GetUrlsRequest{} }
func (m *GetUrlsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUrlsRequest) ProtoMessage()               {}
func (*GetUrlsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetUrlsResponse struct {
	Keys []string `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
}

func (m *GetUrlsResponse) Reset()                    { *m = GetUrlsResponse{} }
func (m *GetUrlsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUrlsResponse) ProtoMessage()               {}
func (*GetUrlsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetUrlsResponse) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type GetUrlRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetUrlRequest) Reset()                    { *m = GetUrlRequest{} }
func (m *GetUrlRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUrlRequest) ProtoMessage()               {}
func (*GetUrlRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetUrlRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetUrlResponse struct {
	Url *Url `protobuf:"bytes,1,opt,name=Url" json:"Url,omitempty"`
}

func (m *GetUrlResponse) Reset()                    { *m = GetUrlResponse{} }
func (m *GetUrlResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUrlResponse) ProtoMessage()               {}
func (*GetUrlResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetUrlResponse) GetUrl() *Url {
	if m != nil {
		return m.Url
	}
	return nil
}

type CreateUrlRequest struct {
	Link string `protobuf:"bytes,1,opt,name=link" json:"link,omitempty"`
}

func (m *CreateUrlRequest) Reset()                    { *m = CreateUrlRequest{} }
func (m *CreateUrlRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateUrlRequest) ProtoMessage()               {}
func (*CreateUrlRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CreateUrlRequest) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

type CreateUrlResponse struct {
	Url *Url `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
}

func (m *CreateUrlResponse) Reset()                    { *m = CreateUrlResponse{} }
func (m *CreateUrlResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateUrlResponse) ProtoMessage()               {}
func (*CreateUrlResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CreateUrlResponse) GetUrl() *Url {
	if m != nil {
		return m.Url
	}
	return nil
}

func init() {
	proto.RegisterType((*Url)(nil), "v2beta.Url")
	proto.RegisterType((*GetUrlsRequest)(nil), "v2beta.GetUrlsRequest")
	proto.RegisterType((*GetUrlsResponse)(nil), "v2beta.GetUrlsResponse")
	proto.RegisterType((*GetUrlRequest)(nil), "v2beta.GetUrlRequest")
	proto.RegisterType((*GetUrlResponse)(nil), "v2beta.GetUrlResponse")
	proto.RegisterType((*CreateUrlRequest)(nil), "v2beta.CreateUrlRequest")
	proto.RegisterType((*CreateUrlResponse)(nil), "v2beta.CreateUrlResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UrlService service

type UrlServiceClient interface {
	// List returns all available URLs keys
	List(ctx context.Context, in *GetUrlsRequest, opts ...grpc.CallOption) (*GetUrlsResponse, error)
	// Create creates a new Url
	Create(ctx context.Context, in *CreateUrlRequest, opts ...grpc.CallOption) (*CreateUrlResponse, error)
	// Get return a Url
	Get(ctx context.Context, in *GetUrlRequest, opts ...grpc.CallOption) (*GetUrlResponse, error)
}

type urlServiceClient struct {
	cc *grpc.ClientConn
}

func NewUrlServiceClient(cc *grpc.ClientConn) UrlServiceClient {
	return &urlServiceClient{cc}
}

func (c *urlServiceClient) List(ctx context.Context, in *GetUrlsRequest, opts ...grpc.CallOption) (*GetUrlsResponse, error) {
	out := new(GetUrlsResponse)
	err := grpc.Invoke(ctx, "/v2beta.UrlService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlServiceClient) Create(ctx context.Context, in *CreateUrlRequest, opts ...grpc.CallOption) (*CreateUrlResponse, error) {
	out := new(CreateUrlResponse)
	err := grpc.Invoke(ctx, "/v2beta.UrlService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlServiceClient) Get(ctx context.Context, in *GetUrlRequest, opts ...grpc.CallOption) (*GetUrlResponse, error) {
	out := new(GetUrlResponse)
	err := grpc.Invoke(ctx, "/v2beta.UrlService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UrlService service

type UrlServiceServer interface {
	// List returns all available URLs keys
	List(context.Context, *GetUrlsRequest) (*GetUrlsResponse, error)
	// Create creates a new Url
	Create(context.Context, *CreateUrlRequest) (*CreateUrlResponse, error)
	// Get return a Url
	Get(context.Context, *GetUrlRequest) (*GetUrlResponse, error)
}

func RegisterUrlServiceServer(s *grpc.Server, srv UrlServiceServer) {
	s.RegisterService(&_UrlService_serviceDesc, srv)
}

func _UrlService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUrlsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2beta.UrlService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServiceServer).List(ctx, req.(*GetUrlsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2beta.UrlService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServiceServer).Create(ctx, req.(*CreateUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2beta.UrlService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServiceServer).Get(ctx, req.(*GetUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UrlService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v2beta.UrlService",
	HandlerType: (*UrlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _UrlService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UrlService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UrlService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "urls.proto",
}

func init() { proto.RegisterFile("urls.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x49, 0x53, 0xca, 0xff, 0xdf, 0x6a, 0x6d, 0x2f, 0xb5, 0x8d, 0x41, 0xa1, 0x0e, 0x28,
	0xe2, 0xa2, 0x81, 0xb8, 0x73, 0xeb, 0xa2, 0x1b, 0x05, 0xa9, 0xd4, 0x7d, 0x5a, 0x2e, 0x25, 0x64,
	0xc8, 0xd4, 0x99, 0x49, 0xa1, 0x88, 0x1b, 0xc1, 0x27, 0xf0, 0xd1, 0x7c, 0x05, 0x1f, 0x44, 0x66,
	0x92, 0xb4, 0x4d, 0xab, 0xbb, 0x9b, 0x93, 0x73, 0xbf, 0x9c, 0x93, 0x19, 0x80, 0x4c, 0x72, 0x35,
	0x5c, 0x48, 0xa1, 0x05, 0x36, 0x96, 0xe1, 0x94, 0x74, 0xe4, 0x9f, 0xce, 0x85, 0x98, 0x73, 0x0a,
	0xa2, 0x45, 0x1c, 0x44, 0x69, 0x2a, 0x74, 0xa4, 0x63, 0x91, 0x16, 0x2e, 0x36, 0x02, 0x77, 0x22,
	0x39, 0xb6, 0xc1, 0x4d, 0x68, 0xe5, 0x39, 0x03, 0xe7, 0xea, 0xff, 0xd8, 0x8c, 0x88, 0x50, 0xe7,
	0x71, 0x9a, 0x78, 0x35, 0x2b, 0xd9, 0x19, 0x7d, 0xf8, 0x37, 0x93, 0x64, 0xf7, 0x3d, 0xd7, 0xea,
	0xeb, 0x67, 0xd6, 0x86, 0xd6, 0x88, 0xf4, 0x44, 0x72, 0x35, 0xa6, 0x97, 0x8c, 0x94, 0x66, 0x17,
	0x70, 0xb4, 0x56, 0xd4, 0x42, 0xa4, 0x8a, 0x0c, 0x34, 0xa1, 0x95, 0xf2, 0x9c, 0x81, 0x6b, 0xa0,
	0x66, 0x66, 0xe7, 0x70, 0x98, 0xdb, 0x8a, 0xbd, 0xfd, 0x2c, 0x2c, 0x28, 0xd9, 0x6b, 0xd0, 0x99,
	0x8d, 0x6d, 0x3d, 0xcd, 0xb0, 0x39, 0xcc, 0xab, 0x0e, 0x8d, 0xc3, 0xe8, 0xec, 0x12, 0xda, 0x77,
	0x26, 0x18, 0x6d, 0x61, 0xcb, 0x42, 0xce, 0xa6, 0x10, 0x0b, 0xa1, 0xb3, 0xe5, 0xdb, 0xb0, 0xb3,
	0x3f, 0xd8, 0x99, 0xe4, 0xe1, 0x47, 0x0d, 0x60, 0x22, 0xf9, 0x13, 0xc9, 0x65, 0x3c, 0x23, 0x7c,
	0x80, 0xfa, 0x7d, 0xac, 0x34, 0xf6, 0x4a, 0x63, 0xf5, 0x2f, 0xf8, 0xfd, 0x3d, 0x3d, 0xff, 0x0c,
	0xeb, 0xbe, 0x7f, 0x7d, 0x7f, 0xd6, 0x5a, 0x78, 0x10, 0xe4, 0x86, 0xc0, 0x9c, 0x1d, 0x3e, 0x43,
	0x23, 0x4f, 0x84, 0x5e, 0xb9, 0xb8, 0xdb, 0xc4, 0x3f, 0xf9, 0xe5, 0x4d, 0x01, 0xed, 0x5b, 0x68,
	0x87, 0x55, 0xa0, 0xb7, 0xce, 0x35, 0x3e, 0x82, 0x3b, 0x22, 0x8d, 0xc7, 0xd5, 0x34, 0x25, 0xb1,
	0xb7, 0x2b, 0x17, 0x38, 0xdf, 0xe2, 0xba, 0x88, 0xdb, 0xb8, 0xe0, 0x35, 0xa1, 0xd5, 0xdb, 0xb4,
	0x61, 0x2f, 0xd0, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0xdd, 0x48, 0x44, 0x74, 0x02,
	0x00, 0x00,
}
