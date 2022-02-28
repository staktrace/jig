// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package httpgreet

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HttpGreeterClient is the client API for HttpGreeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HttpGreeterClient interface {
	// Hello greets.
	GetHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	PostHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	PostHelloURL(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type httpGreeterClient struct {
	cc grpc.ClientConnInterface
}

func NewHttpGreeterClient(cc grpc.ClientConnInterface) HttpGreeterClient {
	return &httpGreeterClient{cc}
}

func (c *httpGreeterClient) GetHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/httpgreet.HttpGreeter/GetHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpGreeterClient) PostHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/httpgreet.HttpGreeter/PostHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpGreeterClient) PostHelloURL(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/httpgreet.HttpGreeter/PostHelloURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HttpGreeterServer is the server API for HttpGreeter service.
// All implementations must embed UnimplementedHttpGreeterServer
// for forward compatibility
type HttpGreeterServer interface {
	// Hello greets.
	GetHello(context.Context, *HelloRequest) (*HelloResponse, error)
	PostHello(context.Context, *HelloRequest) (*HelloResponse, error)
	PostHelloURL(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedHttpGreeterServer()
}

// UnimplementedHttpGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedHttpGreeterServer struct {
}

func (UnimplementedHttpGreeterServer) GetHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHello not implemented")
}
func (UnimplementedHttpGreeterServer) PostHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostHello not implemented")
}
func (UnimplementedHttpGreeterServer) PostHelloURL(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostHelloURL not implemented")
}
func (UnimplementedHttpGreeterServer) mustEmbedUnimplementedHttpGreeterServer() {}

// UnsafeHttpGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HttpGreeterServer will
// result in compilation errors.
type UnsafeHttpGreeterServer interface {
	mustEmbedUnimplementedHttpGreeterServer()
}

func RegisterHttpGreeterServer(s grpc.ServiceRegistrar, srv HttpGreeterServer) {
	s.RegisterService(&HttpGreeter_ServiceDesc, srv)
}

func _HttpGreeter_GetHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpGreeterServer).GetHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/httpgreet.HttpGreeter/GetHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpGreeterServer).GetHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpGreeter_PostHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpGreeterServer).PostHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/httpgreet.HttpGreeter/PostHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpGreeterServer).PostHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpGreeter_PostHelloURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpGreeterServer).PostHelloURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/httpgreet.HttpGreeter/PostHelloURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpGreeterServer).PostHelloURL(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HttpGreeter_ServiceDesc is the grpc.ServiceDesc for HttpGreeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HttpGreeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "httpgreet.HttpGreeter",
	HandlerType: (*HttpGreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHello",
			Handler:    _HttpGreeter_GetHello_Handler,
		},
		{
			MethodName: "PostHello",
			Handler:    _HttpGreeter_PostHello_Handler,
		},
		{
			MethodName: "PostHelloURL",
			Handler:    _HttpGreeter_PostHelloURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "httpgreet/httpgreet.proto",
}