// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: aspen/transport/grpc/v1/kv.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	FeedbackService_Exec_FullMethodName = "/aspen.v1.FeedbackService/Exec"
)

// FeedbackServiceClient is the client API for FeedbackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedbackServiceClient interface {
	Exec(ctx context.Context, in *FeedbackMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type feedbackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedbackServiceClient(cc grpc.ClientConnInterface) FeedbackServiceClient {
	return &feedbackServiceClient{cc}
}

func (c *feedbackServiceClient) Exec(ctx context.Context, in *FeedbackMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FeedbackService_Exec_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedbackServiceServer is the server API for FeedbackService service.
// All implementations should embed UnimplementedFeedbackServiceServer
// for forward compatibility
type FeedbackServiceServer interface {
	Exec(context.Context, *FeedbackMessage) (*emptypb.Empty, error)
}

// UnimplementedFeedbackServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFeedbackServiceServer struct {
}

func (UnimplementedFeedbackServiceServer) Exec(context.Context, *FeedbackMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}

// UnsafeFeedbackServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedbackServiceServer will
// result in compilation errors.
type UnsafeFeedbackServiceServer interface {
	mustEmbedUnimplementedFeedbackServiceServer()
}

func RegisterFeedbackServiceServer(s grpc.ServiceRegistrar, srv FeedbackServiceServer) {
	s.RegisterService(&FeedbackService_ServiceDesc, srv)
}

func _FeedbackService_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedbackMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServiceServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedbackService_Exec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServiceServer).Exec(ctx, req.(*FeedbackMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedbackService_ServiceDesc is the grpc.ServiceDesc for FeedbackService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedbackService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aspen.v1.FeedbackService",
	HandlerType: (*FeedbackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exec",
			Handler:    _FeedbackService_Exec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aspen/transport/grpc/v1/kv.proto",
}

const (
	TxService_Exec_FullMethodName = "/aspen.v1.TxService/Exec"
)

// TxServiceClient is the client API for TxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TxServiceClient interface {
	Exec(ctx context.Context, in *TxRequest, opts ...grpc.CallOption) (*TxRequest, error)
}

type txServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTxServiceClient(cc grpc.ClientConnInterface) TxServiceClient {
	return &txServiceClient{cc}
}

func (c *txServiceClient) Exec(ctx context.Context, in *TxRequest, opts ...grpc.CallOption) (*TxRequest, error) {
	out := new(TxRequest)
	err := c.cc.Invoke(ctx, TxService_Exec_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TxServiceServer is the server API for TxService service.
// All implementations should embed UnimplementedTxServiceServer
// for forward compatibility
type TxServiceServer interface {
	Exec(context.Context, *TxRequest) (*TxRequest, error)
}

// UnimplementedTxServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTxServiceServer struct {
}

func (UnimplementedTxServiceServer) Exec(context.Context, *TxRequest) (*TxRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}

// UnsafeTxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TxServiceServer will
// result in compilation errors.
type UnsafeTxServiceServer interface {
	mustEmbedUnimplementedTxServiceServer()
}

func RegisterTxServiceServer(s grpc.ServiceRegistrar, srv TxServiceServer) {
	s.RegisterService(&TxService_ServiceDesc, srv)
}

func _TxService_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServiceServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TxService_Exec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServiceServer).Exec(ctx, req.(*TxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TxService_ServiceDesc is the grpc.ServiceDesc for TxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aspen.v1.TxService",
	HandlerType: (*TxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exec",
			Handler:    _TxService_Exec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aspen/transport/grpc/v1/kv.proto",
}

const (
	LeaseService_Exec_FullMethodName = "/aspen.v1.LeaseService/Exec"
)

// LeaseServiceClient is the client API for LeaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LeaseServiceClient interface {
	Exec(ctx context.Context, in *TxRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type leaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLeaseServiceClient(cc grpc.ClientConnInterface) LeaseServiceClient {
	return &leaseServiceClient{cc}
}

func (c *leaseServiceClient) Exec(ctx context.Context, in *TxRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LeaseService_Exec_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeaseServiceServer is the server API for LeaseService service.
// All implementations should embed UnimplementedLeaseServiceServer
// for forward compatibility
type LeaseServiceServer interface {
	Exec(context.Context, *TxRequest) (*emptypb.Empty, error)
}

// UnimplementedLeaseServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLeaseServiceServer struct {
}

func (UnimplementedLeaseServiceServer) Exec(context.Context, *TxRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}

// UnsafeLeaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LeaseServiceServer will
// result in compilation errors.
type UnsafeLeaseServiceServer interface {
	mustEmbedUnimplementedLeaseServiceServer()
}

func RegisterLeaseServiceServer(s grpc.ServiceRegistrar, srv LeaseServiceServer) {
	s.RegisterService(&LeaseService_ServiceDesc, srv)
}

func _LeaseService_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaseServiceServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LeaseService_Exec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaseServiceServer).Exec(ctx, req.(*TxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LeaseService_ServiceDesc is the grpc.ServiceDesc for LeaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LeaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aspen.v1.LeaseService",
	HandlerType: (*LeaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exec",
			Handler:    _LeaseService_Exec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aspen/transport/grpc/v1/kv.proto",
}

const (
	RecoveryService_Exec_FullMethodName = "/aspen.v1.RecoveryService/Exec"
)

// RecoveryServiceClient is the client API for RecoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecoveryServiceClient interface {
	Exec(ctx context.Context, opts ...grpc.CallOption) (RecoveryService_ExecClient, error)
}

type recoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecoveryServiceClient(cc grpc.ClientConnInterface) RecoveryServiceClient {
	return &recoveryServiceClient{cc}
}

func (c *recoveryServiceClient) Exec(ctx context.Context, opts ...grpc.CallOption) (RecoveryService_ExecClient, error) {
	stream, err := c.cc.NewStream(ctx, &RecoveryService_ServiceDesc.Streams[0], RecoveryService_Exec_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &recoveryServiceExecClient{stream}
	return x, nil
}

type RecoveryService_ExecClient interface {
	Send(*RecoveryRequest) error
	Recv() (*RecoveryResponse, error)
	grpc.ClientStream
}

type recoveryServiceExecClient struct {
	grpc.ClientStream
}

func (x *recoveryServiceExecClient) Send(m *RecoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *recoveryServiceExecClient) Recv() (*RecoveryResponse, error) {
	m := new(RecoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RecoveryServiceServer is the server API for RecoveryService service.
// All implementations should embed UnimplementedRecoveryServiceServer
// for forward compatibility
type RecoveryServiceServer interface {
	Exec(RecoveryService_ExecServer) error
}

// UnimplementedRecoveryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRecoveryServiceServer struct {
}

func (UnimplementedRecoveryServiceServer) Exec(RecoveryService_ExecServer) error {
	return status.Errorf(codes.Unimplemented, "method Exec not implemented")
}

// UnsafeRecoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecoveryServiceServer will
// result in compilation errors.
type UnsafeRecoveryServiceServer interface {
	mustEmbedUnimplementedRecoveryServiceServer()
}

func RegisterRecoveryServiceServer(s grpc.ServiceRegistrar, srv RecoveryServiceServer) {
	s.RegisterService(&RecoveryService_ServiceDesc, srv)
}

func _RecoveryService_Exec_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RecoveryServiceServer).Exec(&recoveryServiceExecServer{stream})
}

type RecoveryService_ExecServer interface {
	Send(*RecoveryResponse) error
	Recv() (*RecoveryRequest, error)
	grpc.ServerStream
}

type recoveryServiceExecServer struct {
	grpc.ServerStream
}

func (x *recoveryServiceExecServer) Send(m *RecoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *recoveryServiceExecServer) Recv() (*RecoveryRequest, error) {
	m := new(RecoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RecoveryService_ServiceDesc is the grpc.ServiceDesc for RecoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aspen.v1.RecoveryService",
	HandlerType: (*RecoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Exec",
			Handler:       _RecoveryService_Exec_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "aspen/transport/grpc/v1/kv.proto",
}
