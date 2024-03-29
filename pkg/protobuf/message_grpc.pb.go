// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: pkg/protobuf/message.proto

package protobuf

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

const (
	GreetService_Hello_FullMethodName       = "/proto.GreetService/Hello"
	GreetService_HelloStream_FullMethodName = "/proto.GreetService/HelloStream"
)

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	// Unary RPC
	// 客户端发起了一个RPC请求到服务端，服务端进行业务处理并返回响应给客户端，
	// 这是gRPC最基本的一种工作方式（Unary RPC）
	Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// Streaming RPC
	// 服务端流，客户端发出一个RPC请求，服务端客户端与之间建立一个单向的流。
	// 服务端可以向流中写入多个响应消息，最后主动关闭流；而客户端需要监听这个流，不断获取响应直到流关闭。
	// 应用场景举例：客户端向服务端发送一个股票代码，服务端就把该股票的实时数据源源不断的返回给客户端。
	//
	// 客户端流，客户端传入多个请求对象，服务端返回一个响应结果。
	// 典型的应用场景举例：物联网终端向服务器上报数据、大数据流式计算等。
	//
	// 双向流，双向流式RPC即客户端和服务端均为流式的RPC，能发送多个请求对象也能接收到多个响应对象。
	// 典型应用示例：聊天应用等。
	HelloStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_HelloStreamClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, GreetService_Hello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) HelloStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_HelloStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], GreetService_HelloStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceHelloStreamClient{stream}
	return x, nil
}

type GreetService_HelloStreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type greetServiceHelloStreamClient struct {
	grpc.ClientStream
}

func (x *greetServiceHelloStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceHelloStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	// Unary RPC
	// 客户端发起了一个RPC请求到服务端，服务端进行业务处理并返回响应给客户端，
	// 这是gRPC最基本的一种工作方式（Unary RPC）
	Hello(context.Context, *Request) (*Response, error)
	// Streaming RPC
	// 服务端流，客户端发出一个RPC请求，服务端客户端与之间建立一个单向的流。
	// 服务端可以向流中写入多个响应消息，最后主动关闭流；而客户端需要监听这个流，不断获取响应直到流关闭。
	// 应用场景举例：客户端向服务端发送一个股票代码，服务端就把该股票的实时数据源源不断的返回给客户端。
	//
	// 客户端流，客户端传入多个请求对象，服务端返回一个响应结果。
	// 典型的应用场景举例：物联网终端向服务器上报数据、大数据流式计算等。
	//
	// 双向流，双向流式RPC即客户端和服务端均为流式的RPC，能发送多个请求对象也能接收到多个响应对象。
	// 典型应用示例：聊天应用等。
	HelloStream(GreetService_HelloStreamServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) Hello(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedGreetServiceServer) HelloStream(GreetService_HelloStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloStream not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreetService_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Hello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_HelloStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).HelloStream(&greetServiceHelloStreamServer{stream})
}

type GreetService_HelloStreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type greetServiceHelloStreamServer struct {
	grpc.ServerStream
}

func (x *greetServiceHelloStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceHelloStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _GreetService_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloStream",
			Handler:       _GreetService_HelloStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/protobuf/message.proto",
}
