// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package greet

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

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// Hello greets.
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	// HelloClientStream greets everyone at once.
	HelloClientStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_HelloClientStreamClient, error)
	// HelloServerStream greets repeatedly.
	HelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_HelloServerStreamClient, error)
	// HelloBidiStream greets everyone individually.
	HelloBidiStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_HelloBidiStreamClient, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/greet.Greeter/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) HelloClientStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_HelloClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[0], "/greet.Greeter/HelloClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterHelloClientStreamClient{stream}
	return x, nil
}

type Greeter_HelloClientStreamClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloResponse, error)
	grpc.ClientStream
}

type greeterHelloClientStreamClient struct {
	grpc.ClientStream
}

func (x *greeterHelloClientStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterHelloClientStreamClient) CloseAndRecv() (*HelloResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) HelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_HelloServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[1], "/greet.Greeter/HelloServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterHelloServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_HelloServerStreamClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greeterHelloServerStreamClient struct {
	grpc.ClientStream
}

func (x *greeterHelloServerStreamClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) HelloBidiStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_HelloBidiStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[2], "/greet.Greeter/HelloBidiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterHelloBidiStreamClient{stream}
	return x, nil
}

type Greeter_HelloBidiStreamClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greeterHelloBidiStreamClient struct {
	grpc.ClientStream
}

func (x *greeterHelloBidiStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterHelloBidiStreamClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// Hello greets.
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	// HelloClientStream greets everyone at once.
	HelloClientStream(Greeter_HelloClientStreamServer) error
	// HelloServerStream greets repeatedly.
	HelloServerStream(*HelloRequest, Greeter_HelloServerStreamServer) error
	// HelloBidiStream greets everyone individually.
	HelloBidiStream(Greeter_HelloBidiStreamServer) error
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedGreeterServer) HelloClientStream(Greeter_HelloClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloClientStream not implemented")
}
func (UnimplementedGreeterServer) HelloServerStream(*HelloRequest, Greeter_HelloServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloServerStream not implemented")
}
func (UnimplementedGreeterServer) HelloBidiStream(Greeter_HelloBidiStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloBidiStream not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.Greeter/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_HelloClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).HelloClientStream(&greeterHelloClientStreamServer{stream})
}

type Greeter_HelloClientStreamServer interface {
	SendAndClose(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterHelloClientStreamServer struct {
	grpc.ServerStream
}

func (x *greeterHelloClientStreamServer) SendAndClose(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterHelloClientStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_HelloServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).HelloServerStream(m, &greeterHelloServerStreamServer{stream})
}

type Greeter_HelloServerStreamServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type greeterHelloServerStreamServer struct {
	grpc.ServerStream
}

func (x *greeterHelloServerStreamServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_HelloBidiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).HelloBidiStream(&greeterHelloBidiStreamServer{stream})
}

type Greeter_HelloBidiStreamServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterHelloBidiStreamServer struct {
	grpc.ServerStream
}

func (x *greeterHelloBidiStreamServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterHelloBidiStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Greeter_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloClientStream",
			Handler:       _Greeter_HelloClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "HelloServerStream",
			Handler:       _Greeter_HelloServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HelloBidiStream",
			Handler:       _Greeter_HelloBidiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greet/greeter.proto",
}