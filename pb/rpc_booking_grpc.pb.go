// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: rpc_booking.proto

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

// RPCBookingClient is the client API for RPCBooking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCBookingClient interface {
	FindById(ctx context.Context, in *BookingParamId, opts ...grpc.CallOption) (*Booking, error)
	CreateBooking(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*Booking, error)
	UpdateBooking(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*Booking, error)
	SearchBooking(ctx context.Context, in *SearchBookingRequest, opts ...grpc.CallOption) (*SearchBookingResponse, error)
}

type rPCBookingClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCBookingClient(cc grpc.ClientConnInterface) RPCBookingClient {
	return &rPCBookingClient{cc}
}

func (c *rPCBookingClient) FindById(ctx context.Context, in *BookingParamId, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/tuns_go_flight.RPCBooking/FindById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCBookingClient) CreateBooking(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/tuns_go_flight.RPCBooking/CreateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCBookingClient) UpdateBooking(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/tuns_go_flight.RPCBooking/UpdateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCBookingClient) SearchBooking(ctx context.Context, in *SearchBookingRequest, opts ...grpc.CallOption) (*SearchBookingResponse, error) {
	out := new(SearchBookingResponse)
	err := c.cc.Invoke(ctx, "/tuns_go_flight.RPCBooking/SearchBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCBookingServer is the server API for RPCBooking service.
// All implementations must embed UnimplementedRPCBookingServer
// for forward compatibility
type RPCBookingServer interface {
	FindById(context.Context, *BookingParamId) (*Booking, error)
	CreateBooking(context.Context, *Booking) (*Booking, error)
	UpdateBooking(context.Context, *Booking) (*Booking, error)
	SearchBooking(context.Context, *SearchBookingRequest) (*SearchBookingResponse, error)
	mustEmbedUnimplementedRPCBookingServer()
}

// UnimplementedRPCBookingServer must be embedded to have forward compatible implementations.
type UnimplementedRPCBookingServer struct {
}

func (UnimplementedRPCBookingServer) FindById(context.Context, *BookingParamId) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedRPCBookingServer) CreateBooking(context.Context, *Booking) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedRPCBookingServer) UpdateBooking(context.Context, *Booking) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBooking not implemented")
}
func (UnimplementedRPCBookingServer) SearchBooking(context.Context, *SearchBookingRequest) (*SearchBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBooking not implemented")
}
func (UnimplementedRPCBookingServer) mustEmbedUnimplementedRPCBookingServer() {}

// UnsafeRPCBookingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCBookingServer will
// result in compilation errors.
type UnsafeRPCBookingServer interface {
	mustEmbedUnimplementedRPCBookingServer()
}

func RegisterRPCBookingServer(s grpc.ServiceRegistrar, srv RPCBookingServer) {
	s.RegisterService(&RPCBooking_ServiceDesc, srv)
}

func _RPCBooking_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingParamId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCBookingServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tuns_go_flight.RPCBooking/FindById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCBookingServer).FindById(ctx, req.(*BookingParamId))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCBooking_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Booking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCBookingServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tuns_go_flight.RPCBooking/CreateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCBookingServer).CreateBooking(ctx, req.(*Booking))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCBooking_UpdateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Booking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCBookingServer).UpdateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tuns_go_flight.RPCBooking/UpdateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCBookingServer).UpdateBooking(ctx, req.(*Booking))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCBooking_SearchBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCBookingServer).SearchBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tuns_go_flight.RPCBooking/SearchBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCBookingServer).SearchBooking(ctx, req.(*SearchBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCBooking_ServiceDesc is the grpc.ServiceDesc for RPCBooking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCBooking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tuns_go_flight.RPCBooking",
	HandlerType: (*RPCBookingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindById",
			Handler:    _RPCBooking_FindById_Handler,
		},
		{
			MethodName: "CreateBooking",
			Handler:    _RPCBooking_CreateBooking_Handler,
		},
		{
			MethodName: "UpdateBooking",
			Handler:    _RPCBooking_UpdateBooking_Handler,
		},
		{
			MethodName: "SearchBooking",
			Handler:    _RPCBooking_SearchBooking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_booking.proto",
}
