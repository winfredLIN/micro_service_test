// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stock

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

// StockServiceClient is the client API for StockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockServiceClient interface {
	// 商品库存变动
	StockChange(ctx context.Context, in *ChangeRequest, opts ...grpc.CallOption) (*ChangeResponse, error)
	StockShow(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error)
}

type stockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStockServiceClient(cc grpc.ClientConnInterface) StockServiceClient {
	return &stockServiceClient{cc}
}

func (c *stockServiceClient) StockChange(ctx context.Context, in *ChangeRequest, opts ...grpc.CallOption) (*ChangeResponse, error) {
	out := new(ChangeResponse)
	err := c.cc.Invoke(ctx, "/api.protobuf.stock.StockService/StockChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) StockShow(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error) {
	out := new(ShowResponse)
	err := c.cc.Invoke(ctx, "/api.protobuf.stock.StockService/StockShow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServiceServer is the server API for StockService service.
// All implementations must embed UnimplementedStockServiceServer
// for forward compatibility
type StockServiceServer interface {
	// 商品库存变动
	StockChange(context.Context, *ChangeRequest) (*ChangeResponse, error)
	StockShow(context.Context, *ShowRequest) (*ShowResponse, error)
	mustEmbedUnimplementedStockServiceServer()
}

// UnimplementedStockServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStockServiceServer struct {
}

func (UnimplementedStockServiceServer) StockChange(context.Context, *ChangeRequest) (*ChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StockChange not implemented")
}
func (UnimplementedStockServiceServer) StockShow(context.Context, *ShowRequest) (*ShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StockShow not implemented")
}
func (UnimplementedStockServiceServer) mustEmbedUnimplementedStockServiceServer() {}

// UnsafeStockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockServiceServer will
// result in compilation errors.
type UnsafeStockServiceServer interface {
	mustEmbedUnimplementedStockServiceServer()
}

func RegisterStockServiceServer(s grpc.ServiceRegistrar, srv StockServiceServer) {
	s.RegisterService(&StockService_ServiceDesc, srv)
}

func _StockService_StockChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).StockChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.protobuf.stock.StockService/StockChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).StockChange(ctx, req.(*ChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_StockShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).StockShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.protobuf.stock.StockService/StockShow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).StockShow(ctx, req.(*ShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StockService_ServiceDesc is the grpc.ServiceDesc for StockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.protobuf.stock.StockService",
	HandlerType: (*StockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StockChange",
			Handler:    _StockService_StockChange_Handler,
		},
		{
			MethodName: "StockShow",
			Handler:    _StockService_StockShow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/protobuf/stock/stock.proto",
}
