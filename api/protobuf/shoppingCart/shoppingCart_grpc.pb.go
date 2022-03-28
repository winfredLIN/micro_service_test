// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package shoppingCart

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

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartServiceClient interface {
	CartAdd(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	CartLook(ctx context.Context, in *LookRequest, opts ...grpc.CallOption) (*LookResponse, error)
	CommodityShow(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) CartAdd(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/api.protobuf.shoppingCart.CartService/CartAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) CartLook(ctx context.Context, in *LookRequest, opts ...grpc.CallOption) (*LookResponse, error) {
	out := new(LookResponse)
	err := c.cc.Invoke(ctx, "/api.protobuf.shoppingCart.CartService/CartLook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) CommodityShow(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error) {
	out := new(ShowResponse)
	err := c.cc.Invoke(ctx, "/api.protobuf.shoppingCart.CartService/CommodityShow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations must embed UnimplementedCartServiceServer
// for forward compatibility
type CartServiceServer interface {
	CartAdd(context.Context, *AddRequest) (*AddResponse, error)
	CartLook(context.Context, *LookRequest) (*LookResponse, error)
	CommodityShow(context.Context, *ShowRequest) (*ShowResponse, error)
	mustEmbedUnimplementedCartServiceServer()
}

// UnimplementedCartServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (UnimplementedCartServiceServer) CartAdd(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartAdd not implemented")
}
func (UnimplementedCartServiceServer) CartLook(context.Context, *LookRequest) (*LookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartLook not implemented")
}
func (UnimplementedCartServiceServer) CommodityShow(context.Context, *ShowRequest) (*ShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommodityShow not implemented")
}
func (UnimplementedCartServiceServer) mustEmbedUnimplementedCartServiceServer() {}

// UnsafeCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServiceServer will
// result in compilation errors.
type UnsafeCartServiceServer interface {
	mustEmbedUnimplementedCartServiceServer()
}

func RegisterCartServiceServer(s grpc.ServiceRegistrar, srv CartServiceServer) {
	s.RegisterService(&CartService_ServiceDesc, srv)
}

func _CartService_CartAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).CartAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.protobuf.shoppingCart.CartService/CartAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).CartAdd(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_CartLook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).CartLook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.protobuf.shoppingCart.CartService/CartLook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).CartLook(ctx, req.(*LookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_CommodityShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).CommodityShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.protobuf.shoppingCart.CartService/CommodityShow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).CommodityShow(ctx, req.(*ShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CartService_ServiceDesc is the grpc.ServiceDesc for CartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.protobuf.shoppingCart.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CartAdd",
			Handler:    _CartService_CartAdd_Handler,
		},
		{
			MethodName: "CartLook",
			Handler:    _CartService_CartLook_Handler,
		},
		{
			MethodName: "CommodityShow",
			Handler:    _CartService_CommodityShow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/protobuf/shoppingCart/shoppingCart.proto",
}
