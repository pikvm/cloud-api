// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.11
// source: proxy_for_agent/proxy_for_agent.proto

package proxy_for_agent

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

// ProxyForAgentClient is the client API for ProxyForAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProxyForAgentClient interface {
	EventsChannel(ctx context.Context, opts ...grpc.CallOption) (ProxyForAgent_EventsChannelClient, error)
	ConnectionChannel(ctx context.Context, opts ...grpc.CallOption) (ProxyForAgent_ConnectionChannelClient, error)
}

type proxyForAgentClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyForAgentClient(cc grpc.ClientConnInterface) ProxyForAgentClient {
	return &proxyForAgentClient{cc}
}

func (c *proxyForAgentClient) EventsChannel(ctx context.Context, opts ...grpc.CallOption) (ProxyForAgent_EventsChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProxyForAgent_ServiceDesc.Streams[0], "/proxy_for_agent.ProxyForAgent/EventsChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyForAgentEventsChannelClient{stream}
	return x, nil
}

type ProxyForAgent_EventsChannelClient interface {
	Send(*AgentEvent) error
	Recv() (*ProxyEvent, error)
	grpc.ClientStream
}

type proxyForAgentEventsChannelClient struct {
	grpc.ClientStream
}

func (x *proxyForAgentEventsChannelClient) Send(m *AgentEvent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *proxyForAgentEventsChannelClient) Recv() (*ProxyEvent, error) {
	m := new(ProxyEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *proxyForAgentClient) ConnectionChannel(ctx context.Context, opts ...grpc.CallOption) (ProxyForAgent_ConnectionChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProxyForAgent_ServiceDesc.Streams[1], "/proxy_for_agent.ProxyForAgent/ConnectionChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyForAgentConnectionChannelClient{stream}
	return x, nil
}

type ProxyForAgent_ConnectionChannelClient interface {
	Send(*ConnectionMessage) error
	Recv() (*ConnectionMessage, error)
	grpc.ClientStream
}

type proxyForAgentConnectionChannelClient struct {
	grpc.ClientStream
}

func (x *proxyForAgentConnectionChannelClient) Send(m *ConnectionMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *proxyForAgentConnectionChannelClient) Recv() (*ConnectionMessage, error) {
	m := new(ConnectionMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProxyForAgentServer is the server API for ProxyForAgent service.
// All implementations must embed UnimplementedProxyForAgentServer
// for forward compatibility
type ProxyForAgentServer interface {
	EventsChannel(ProxyForAgent_EventsChannelServer) error
	ConnectionChannel(ProxyForAgent_ConnectionChannelServer) error
	mustEmbedUnimplementedProxyForAgentServer()
}

// UnimplementedProxyForAgentServer must be embedded to have forward compatible implementations.
type UnimplementedProxyForAgentServer struct {
}

func (UnimplementedProxyForAgentServer) EventsChannel(ProxyForAgent_EventsChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method EventsChannel not implemented")
}
func (UnimplementedProxyForAgentServer) ConnectionChannel(ProxyForAgent_ConnectionChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectionChannel not implemented")
}
func (UnimplementedProxyForAgentServer) mustEmbedUnimplementedProxyForAgentServer() {}

// UnsafeProxyForAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProxyForAgentServer will
// result in compilation errors.
type UnsafeProxyForAgentServer interface {
	mustEmbedUnimplementedProxyForAgentServer()
}

func RegisterProxyForAgentServer(s grpc.ServiceRegistrar, srv ProxyForAgentServer) {
	s.RegisterService(&ProxyForAgent_ServiceDesc, srv)
}

func _ProxyForAgent_EventsChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProxyForAgentServer).EventsChannel(&proxyForAgentEventsChannelServer{stream})
}

type ProxyForAgent_EventsChannelServer interface {
	Send(*ProxyEvent) error
	Recv() (*AgentEvent, error)
	grpc.ServerStream
}

type proxyForAgentEventsChannelServer struct {
	grpc.ServerStream
}

func (x *proxyForAgentEventsChannelServer) Send(m *ProxyEvent) error {
	return x.ServerStream.SendMsg(m)
}

func (x *proxyForAgentEventsChannelServer) Recv() (*AgentEvent, error) {
	m := new(AgentEvent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProxyForAgent_ConnectionChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProxyForAgentServer).ConnectionChannel(&proxyForAgentConnectionChannelServer{stream})
}

type ProxyForAgent_ConnectionChannelServer interface {
	Send(*ConnectionMessage) error
	Recv() (*ConnectionMessage, error)
	grpc.ServerStream
}

type proxyForAgentConnectionChannelServer struct {
	grpc.ServerStream
}

func (x *proxyForAgentConnectionChannelServer) Send(m *ConnectionMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *proxyForAgentConnectionChannelServer) Recv() (*ConnectionMessage, error) {
	m := new(ConnectionMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProxyForAgent_ServiceDesc is the grpc.ServiceDesc for ProxyForAgent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProxyForAgent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proxy_for_agent.ProxyForAgent",
	HandlerType: (*ProxyForAgentServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EventsChannel",
			Handler:       _ProxyForAgent_EventsChannel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ConnectionChannel",
			Handler:       _ProxyForAgent_ConnectionChannel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proxy_for_agent/proxy_for_agent.proto",
}