// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// TunnelClient is the client API for Tunnel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TunnelClient interface {
	InitTunnel(ctx context.Context, opts ...grpc.CallOption) (Tunnel_InitTunnelClient, error)
	SendData(ctx context.Context, in *SendDataRequest, opts ...grpc.CallOption) (*SendDataResponse, error)
}

type tunnelClient struct {
	cc grpc.ClientConnInterface
}

func NewTunnelClient(cc grpc.ClientConnInterface) TunnelClient {
	return &tunnelClient{cc}
}

func (c *tunnelClient) InitTunnel(ctx context.Context, opts ...grpc.CallOption) (Tunnel_InitTunnelClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tunnel_ServiceDesc.Streams[0], "/tunnel.Tunnel/InitTunnel", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelInitTunnelClient{stream}
	return x, nil
}

type Tunnel_InitTunnelClient interface {
	Send(*InitTunnelRequest) error
	Recv() (*StreamData, error)
	grpc.ClientStream
}

type tunnelInitTunnelClient struct {
	grpc.ClientStream
}

func (x *tunnelInitTunnelClient) Send(m *InitTunnelRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelInitTunnelClient) Recv() (*StreamData, error) {
	m := new(StreamData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tunnelClient) SendData(ctx context.Context, in *SendDataRequest, opts ...grpc.CallOption) (*SendDataResponse, error) {
	out := new(SendDataResponse)
	err := c.cc.Invoke(ctx, "/tunnel.Tunnel/SendData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TunnelServer is the server API for Tunnel service.
// All implementations must embed UnimplementedTunnelServer
// for forward compatibility
type TunnelServer interface {
	InitTunnel(Tunnel_InitTunnelServer) error
	SendData(context.Context, *SendDataRequest) (*SendDataResponse, error)
	mustEmbedUnimplementedTunnelServer()
}

// UnimplementedTunnelServer must be embedded to have forward compatible implementations.
type UnimplementedTunnelServer struct {
}

func (UnimplementedTunnelServer) InitTunnel(Tunnel_InitTunnelServer) error {
	return status.Errorf(codes.Unimplemented, "method InitTunnel not implemented")
}
func (UnimplementedTunnelServer) SendData(context.Context, *SendDataRequest) (*SendDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendData not implemented")
}
func (UnimplementedTunnelServer) mustEmbedUnimplementedTunnelServer() {}

// UnsafeTunnelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TunnelServer will
// result in compilation errors.
type UnsafeTunnelServer interface {
	mustEmbedUnimplementedTunnelServer()
}

func RegisterTunnelServer(s grpc.ServiceRegistrar, srv TunnelServer) {
	s.RegisterService(&Tunnel_ServiceDesc, srv)
}

func _Tunnel_InitTunnel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelServer).InitTunnel(&tunnelInitTunnelServer{stream})
}

type Tunnel_InitTunnelServer interface {
	Send(*StreamData) error
	Recv() (*InitTunnelRequest, error)
	grpc.ServerStream
}

type tunnelInitTunnelServer struct {
	grpc.ServerStream
}

func (x *tunnelInitTunnelServer) Send(m *StreamData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelInitTunnelServer) Recv() (*InitTunnelRequest, error) {
	m := new(InitTunnelRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Tunnel_SendData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServer).SendData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tunnel.Tunnel/SendData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServer).SendData(ctx, req.(*SendDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tunnel_ServiceDesc is the grpc.ServiceDesc for Tunnel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tunnel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tunnel.Tunnel",
	HandlerType: (*TunnelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendData",
			Handler:    _Tunnel_SendData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InitTunnel",
			Handler:       _Tunnel_InitTunnel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/proto/tunnel.proto",
}
