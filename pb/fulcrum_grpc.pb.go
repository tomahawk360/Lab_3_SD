// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: fulcrum.proto

package pb

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

// BrokerServiceClient is the client API for BrokerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrokerServiceClient interface {
	GetSoldiers(ctx context.Context, in *GetSoldiersServiceReq, opts ...grpc.CallOption) (*GetSoldiersServiceRes, error)
	GetServer(ctx context.Context, in *GetServerServiceReq, opts ...grpc.CallOption) (*GetServerServiceRes, error)
}

type brokerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerServiceClient(cc grpc.ClientConnInterface) BrokerServiceClient {
	return &brokerServiceClient{cc}
}

func (c *brokerServiceClient) GetSoldiers(ctx context.Context, in *GetSoldiersServiceReq, opts ...grpc.CallOption) (*GetSoldiersServiceRes, error) {
	out := new(GetSoldiersServiceRes)
	err := c.cc.Invoke(ctx, "/BrokerService/GetSoldiers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerServiceClient) GetServer(ctx context.Context, in *GetServerServiceReq, opts ...grpc.CallOption) (*GetServerServiceRes, error) {
	out := new(GetServerServiceRes)
	err := c.cc.Invoke(ctx, "/BrokerService/GetServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrokerServiceServer is the server API for BrokerService service.
// All implementations must embed UnimplementedBrokerServiceServer
// for forward compatibility
type BrokerServiceServer interface {
	GetSoldiers(context.Context, *GetSoldiersServiceReq) (*GetSoldiersServiceRes, error)
	GetServer(context.Context, *GetServerServiceReq) (*GetServerServiceRes, error)
	mustEmbedUnimplementedBrokerServiceServer()
}

// UnimplementedBrokerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBrokerServiceServer struct {
}

func (UnimplementedBrokerServiceServer) GetSoldiers(context.Context, *GetSoldiersServiceReq) (*GetSoldiersServiceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSoldiers not implemented")
}
func (UnimplementedBrokerServiceServer) GetServer(context.Context, *GetServerServiceReq) (*GetServerServiceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServer not implemented")
}
func (UnimplementedBrokerServiceServer) mustEmbedUnimplementedBrokerServiceServer() {}

// UnsafeBrokerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrokerServiceServer will
// result in compilation errors.
type UnsafeBrokerServiceServer interface {
	mustEmbedUnimplementedBrokerServiceServer()
}

func RegisterBrokerServiceServer(s grpc.ServiceRegistrar, srv BrokerServiceServer) {
	s.RegisterService(&BrokerService_ServiceDesc, srv)
}

func _BrokerService_GetSoldiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSoldiersServiceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServiceServer).GetSoldiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BrokerService/GetSoldiers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServiceServer).GetSoldiers(ctx, req.(*GetSoldiersServiceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrokerService_GetServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerServiceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServiceServer).GetServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BrokerService/GetServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServiceServer).GetServer(ctx, req.(*GetServerServiceReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BrokerService_ServiceDesc is the grpc.ServiceDesc for BrokerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BrokerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BrokerService",
	HandlerType: (*BrokerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSoldiers",
			Handler:    _BrokerService_GetSoldiers_Handler,
		},
		{
			MethodName: "GetServer",
			Handler:    _BrokerService_GetServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fulcrum.proto",
}

// InformerServiceClient is the client API for InformerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InformerServiceClient interface {
	AddBase(ctx context.Context, in *AddBaseServiceReq, opts ...grpc.CallOption) (*ConnectServiceRes, error)
	RenameBase(ctx context.Context, in *RenameBaseServiceReq, opts ...grpc.CallOption) (*ConnectServiceRes, error)
	UpdateValue(ctx context.Context, in *UpdateValueServiceReq, opts ...grpc.CallOption) (*ConnectServiceRes, error)
	DeleteBase(ctx context.Context, in *DeleteBaseServiceReq, opts ...grpc.CallOption) (*ConnectServiceRes, error)
}

type informerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInformerServiceClient(cc grpc.ClientConnInterface) InformerServiceClient {
	return &informerServiceClient{cc}
}

func (c *informerServiceClient) AddBase(ctx context.Context, in *AddBaseServiceReq, opts ...grpc.CallOption) (*ConnectServiceRes, error) {
	out := new(ConnectServiceRes)
	err := c.cc.Invoke(ctx, "/InformerService/AddBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informerServiceClient) RenameBase(ctx context.Context, in *RenameBaseServiceReq, opts ...grpc.CallOption) (*ConnectServiceRes, error) {
	out := new(ConnectServiceRes)
	err := c.cc.Invoke(ctx, "/InformerService/RenameBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informerServiceClient) UpdateValue(ctx context.Context, in *UpdateValueServiceReq, opts ...grpc.CallOption) (*ConnectServiceRes, error) {
	out := new(ConnectServiceRes)
	err := c.cc.Invoke(ctx, "/InformerService/UpdateValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informerServiceClient) DeleteBase(ctx context.Context, in *DeleteBaseServiceReq, opts ...grpc.CallOption) (*ConnectServiceRes, error) {
	out := new(ConnectServiceRes)
	err := c.cc.Invoke(ctx, "/InformerService/DeleteBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InformerServiceServer is the server API for InformerService service.
// All implementations must embed UnimplementedInformerServiceServer
// for forward compatibility
type InformerServiceServer interface {
	AddBase(context.Context, *AddBaseServiceReq) (*ConnectServiceRes, error)
	RenameBase(context.Context, *RenameBaseServiceReq) (*ConnectServiceRes, error)
	UpdateValue(context.Context, *UpdateValueServiceReq) (*ConnectServiceRes, error)
	DeleteBase(context.Context, *DeleteBaseServiceReq) (*ConnectServiceRes, error)
	mustEmbedUnimplementedInformerServiceServer()
}

// UnimplementedInformerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInformerServiceServer struct {
}

func (UnimplementedInformerServiceServer) AddBase(context.Context, *AddBaseServiceReq) (*ConnectServiceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBase not implemented")
}
func (UnimplementedInformerServiceServer) RenameBase(context.Context, *RenameBaseServiceReq) (*ConnectServiceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameBase not implemented")
}
func (UnimplementedInformerServiceServer) UpdateValue(context.Context, *UpdateValueServiceReq) (*ConnectServiceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateValue not implemented")
}
func (UnimplementedInformerServiceServer) DeleteBase(context.Context, *DeleteBaseServiceReq) (*ConnectServiceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBase not implemented")
}
func (UnimplementedInformerServiceServer) mustEmbedUnimplementedInformerServiceServer() {}

// UnsafeInformerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InformerServiceServer will
// result in compilation errors.
type UnsafeInformerServiceServer interface {
	mustEmbedUnimplementedInformerServiceServer()
}

func RegisterInformerServiceServer(s grpc.ServiceRegistrar, srv InformerServiceServer) {
	s.RegisterService(&InformerService_ServiceDesc, srv)
}

func _InformerService_AddBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBaseServiceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformerServiceServer).AddBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InformerService/AddBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformerServiceServer).AddBase(ctx, req.(*AddBaseServiceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _InformerService_RenameBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameBaseServiceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformerServiceServer).RenameBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InformerService/RenameBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformerServiceServer).RenameBase(ctx, req.(*RenameBaseServiceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _InformerService_UpdateValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateValueServiceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformerServiceServer).UpdateValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InformerService/UpdateValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformerServiceServer).UpdateValue(ctx, req.(*UpdateValueServiceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _InformerService_DeleteBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBaseServiceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformerServiceServer).DeleteBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InformerService/DeleteBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformerServiceServer).DeleteBase(ctx, req.(*DeleteBaseServiceReq))
	}
	return interceptor(ctx, in, info, handler)
}

// InformerService_ServiceDesc is the grpc.ServiceDesc for InformerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InformerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "InformerService",
	HandlerType: (*InformerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBase",
			Handler:    _InformerService_AddBase_Handler,
		},
		{
			MethodName: "RenameBase",
			Handler:    _InformerService_RenameBase_Handler,
		},
		{
			MethodName: "UpdateValue",
			Handler:    _InformerService_UpdateValue_Handler,
		},
		{
			MethodName: "DeleteBase",
			Handler:    _InformerService_DeleteBase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fulcrum.proto",
}

// ServidorServiceClient is the client API for ServidorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServidorServiceClient interface {
	AskServer(ctx context.Context, in *AskServerServiceReq, opts ...grpc.CallOption) (*AskServerServiceRes, error)
}

type servidorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServidorServiceClient(cc grpc.ClientConnInterface) ServidorServiceClient {
	return &servidorServiceClient{cc}
}

func (c *servidorServiceClient) AskServer(ctx context.Context, in *AskServerServiceReq, opts ...grpc.CallOption) (*AskServerServiceRes, error) {
	out := new(AskServerServiceRes)
	err := c.cc.Invoke(ctx, "/ServidorService/AskServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServidorServiceServer is the server API for ServidorService service.
// All implementations must embed UnimplementedServidorServiceServer
// for forward compatibility
type ServidorServiceServer interface {
	AskServer(context.Context, *AskServerServiceReq) (*AskServerServiceRes, error)
	mustEmbedUnimplementedServidorServiceServer()
}

// UnimplementedServidorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServidorServiceServer struct {
}

func (UnimplementedServidorServiceServer) AskServer(context.Context, *AskServerServiceReq) (*AskServerServiceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskServer not implemented")
}
func (UnimplementedServidorServiceServer) mustEmbedUnimplementedServidorServiceServer() {}

// UnsafeServidorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServidorServiceServer will
// result in compilation errors.
type UnsafeServidorServiceServer interface {
	mustEmbedUnimplementedServidorServiceServer()
}

func RegisterServidorServiceServer(s grpc.ServiceRegistrar, srv ServidorServiceServer) {
	s.RegisterService(&ServidorService_ServiceDesc, srv)
}

func _ServidorService_AskServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AskServerServiceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServidorServiceServer).AskServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServidorService/AskServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServidorServiceServer).AskServer(ctx, req.(*AskServerServiceReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ServidorService_ServiceDesc is the grpc.ServiceDesc for ServidorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServidorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ServidorService",
	HandlerType: (*ServidorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AskServer",
			Handler:    _ServidorService_AskServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fulcrum.proto",
}

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogServiceClient interface {
	SendLogs(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
}

type logServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogServiceClient(cc grpc.ClientConnInterface) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) SendLogs(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/LogService/SendLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServiceServer is the server API for LogService service.
// All implementations must embed UnimplementedLogServiceServer
// for forward compatibility
type LogServiceServer interface {
	SendLogs(context.Context, *LogRequest) (*LogResponse, error)
	mustEmbedUnimplementedLogServiceServer()
}

// UnimplementedLogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogServiceServer struct {
}

func (UnimplementedLogServiceServer) SendLogs(context.Context, *LogRequest) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLogs not implemented")
}
func (UnimplementedLogServiceServer) mustEmbedUnimplementedLogServiceServer() {}

// UnsafeLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServiceServer will
// result in compilation errors.
type UnsafeLogServiceServer interface {
	mustEmbedUnimplementedLogServiceServer()
}

func RegisterLogServiceServer(s grpc.ServiceRegistrar, srv LogServiceServer) {
	s.RegisterService(&LogService_ServiceDesc, srv)
}

func _LogService_SendLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).SendLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LogService/SendLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).SendLogs(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogService_ServiceDesc is the grpc.ServiceDesc for LogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendLogs",
			Handler:    _LogService_SendLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fulcrum.proto",
}
