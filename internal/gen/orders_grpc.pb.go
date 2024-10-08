// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: orders.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ClientService_GetMenu_FullMethodName     = "/orders_service.ClientService/GetMenu"
	ClientService_CreateOrder_FullMethodName = "/orders_service.ClientService/CreateOrder"
	ClientService_CancelOrder_FullMethodName = "/orders_service.ClientService/CancelOrder"
)

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	GetMenu(ctx context.Context, in *GetMenuRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetMenuResponse], error)
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) GetMenu(ctx context.Context, in *GetMenuRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetMenuResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[0], ClientService_GetMenu_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetMenuRequest, GetMenuResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ClientService_GetMenuClient = grpc.ServerStreamingClient[GetMenuResponse]

func (c *clientServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, ClientService_CreateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, ClientService_CancelOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations must embed UnimplementedClientServiceServer
// for forward compatibility.
type ClientServiceServer interface {
	GetMenu(*GetMenuRequest, grpc.ServerStreamingServer[GetMenuResponse]) error
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
	mustEmbedUnimplementedClientServiceServer()
}

// UnimplementedClientServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedClientServiceServer struct{}

func (UnimplementedClientServiceServer) GetMenu(*GetMenuRequest, grpc.ServerStreamingServer[GetMenuResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetMenu not implemented")
}
func (UnimplementedClientServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedClientServiceServer) CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedClientServiceServer) mustEmbedUnimplementedClientServiceServer() {}
func (UnimplementedClientServiceServer) testEmbeddedByValue()                       {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	// If the following call pancis, it indicates UnimplementedClientServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_GetMenu_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetMenuRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).GetMenu(m, &grpc.GenericServerStream[GetMenuRequest, GetMenuResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ClientService_GetMenuServer = grpc.ServerStreamingServer[GetMenuResponse]

func _ClientService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_CancelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders_service.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _ClientService_CreateOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _ClientService_CancelOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMenu",
			Handler:       _ClientService_GetMenu_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "orders.proto",
}

const (
	RestaurantService_GetPendingOrders_FullMethodName = "/orders_service.RestaurantService/GetPendingOrders"
	RestaurantService_SetOrderToReady_FullMethodName  = "/orders_service.RestaurantService/SetOrderToReady"
)

// RestaurantServiceClient is the client API for RestaurantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RestaurantServiceClient interface {
	GetPendingOrders(ctx context.Context, in *GetPendingOrdersRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetPendingOrdersResponse], error)
	SetOrderToReady(ctx context.Context, in *SetOrderToReadyRequest, opts ...grpc.CallOption) (*SetOrderToReadyResponse, error)
}

type restaurantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRestaurantServiceClient(cc grpc.ClientConnInterface) RestaurantServiceClient {
	return &restaurantServiceClient{cc}
}

func (c *restaurantServiceClient) GetPendingOrders(ctx context.Context, in *GetPendingOrdersRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetPendingOrdersResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &RestaurantService_ServiceDesc.Streams[0], RestaurantService_GetPendingOrders_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetPendingOrdersRequest, GetPendingOrdersResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RestaurantService_GetPendingOrdersClient = grpc.ServerStreamingClient[GetPendingOrdersResponse]

func (c *restaurantServiceClient) SetOrderToReady(ctx context.Context, in *SetOrderToReadyRequest, opts ...grpc.CallOption) (*SetOrderToReadyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetOrderToReadyResponse)
	err := c.cc.Invoke(ctx, RestaurantService_SetOrderToReady_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestaurantServiceServer is the server API for RestaurantService service.
// All implementations must embed UnimplementedRestaurantServiceServer
// for forward compatibility.
type RestaurantServiceServer interface {
	GetPendingOrders(*GetPendingOrdersRequest, grpc.ServerStreamingServer[GetPendingOrdersResponse]) error
	SetOrderToReady(context.Context, *SetOrderToReadyRequest) (*SetOrderToReadyResponse, error)
	mustEmbedUnimplementedRestaurantServiceServer()
}

// UnimplementedRestaurantServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRestaurantServiceServer struct{}

func (UnimplementedRestaurantServiceServer) GetPendingOrders(*GetPendingOrdersRequest, grpc.ServerStreamingServer[GetPendingOrdersResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetPendingOrders not implemented")
}
func (UnimplementedRestaurantServiceServer) SetOrderToReady(context.Context, *SetOrderToReadyRequest) (*SetOrderToReadyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOrderToReady not implemented")
}
func (UnimplementedRestaurantServiceServer) mustEmbedUnimplementedRestaurantServiceServer() {}
func (UnimplementedRestaurantServiceServer) testEmbeddedByValue()                           {}

// UnsafeRestaurantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RestaurantServiceServer will
// result in compilation errors.
type UnsafeRestaurantServiceServer interface {
	mustEmbedUnimplementedRestaurantServiceServer()
}

func RegisterRestaurantServiceServer(s grpc.ServiceRegistrar, srv RestaurantServiceServer) {
	// If the following call pancis, it indicates UnimplementedRestaurantServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RestaurantService_ServiceDesc, srv)
}

func _RestaurantService_GetPendingOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetPendingOrdersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RestaurantServiceServer).GetPendingOrders(m, &grpc.GenericServerStream[GetPendingOrdersRequest, GetPendingOrdersResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RestaurantService_GetPendingOrdersServer = grpc.ServerStreamingServer[GetPendingOrdersResponse]

func _RestaurantService_SetOrderToReady_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetOrderToReadyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).SetOrderToReady(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestaurantService_SetOrderToReady_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).SetOrderToReady(ctx, req.(*SetOrderToReadyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RestaurantService_ServiceDesc is the grpc.ServiceDesc for RestaurantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RestaurantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders_service.RestaurantService",
	HandlerType: (*RestaurantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetOrderToReady",
			Handler:    _RestaurantService_SetOrderToReady_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPendingOrders",
			Handler:       _RestaurantService_GetPendingOrders_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "orders.proto",
}
