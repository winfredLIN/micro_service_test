// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: greeting.proto

package greeting

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

// GreetingServiceClient is the client API for GreetingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetingServiceClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
}

type greetingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetingServiceClient(cc grpc.ClientConnInterface) GreetingServiceClient {
	return &greetingServiceClient{cc}
}

func (c *greetingServiceClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, "/api.protobuf.greeting.GreetingService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetingServiceServer is the server API for GreetingService service.
// All implementations must embed UnimplementedGreetingServiceServer
// for forward compatibility
type GreetingServiceServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
	mustEmbedUnimplementedGreetingServiceServer()
}

// UnimplementedGreetingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetingServiceServer struct {
}

func (UnimplementedGreetingServiceServer) SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreetingServiceServer) mustEmbedUnimplementedGreetingServiceServer() {}

// UnsafeGreetingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetingServiceServer will
// result in compilation errors.
type UnsafeGreetingServiceServer interface {
	mustEmbedUnimplementedGreetingServiceServer()
}

func RegisterGreetingServiceServer(s grpc.ServiceRegistrar, srv GreetingServiceServer) {
	s.RegisterService(&GreetingService_ServiceDesc, srv)
}

func _GreetingService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.protobuf.greeting.GreetingService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingServiceServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GreetingService_ServiceDesc is the grpc.ServiceDesc for GreetingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.protobuf.greeting.GreetingService",
	HandlerType: (*GreetingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GreetingService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greeting.proto",
}

// RegistrationServiceClient is the client API for RegistrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistrationServiceClient interface {
	Register(ctx context.Context, in *InformationRequest, opts ...grpc.CallOption) (*AnswerResponse, error)
}

type registrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistrationServiceClient(cc grpc.ClientConnInterface) RegistrationServiceClient {
	return &registrationServiceClient{cc}
}

func (c *registrationServiceClient) Register(ctx context.Context, in *InformationRequest, opts ...grpc.CallOption) (*AnswerResponse, error) {
	out := new(AnswerResponse)
	err := c.cc.Invoke(ctx, "/api.protobuf.greeting.RegistrationService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationServiceServer is the server API for RegistrationService service.
// All implementations must embed UnimplementedRegistrationServiceServer
// for forward compatibility
type RegistrationServiceServer interface {
	Register(context.Context, *InformationRequest) (*AnswerResponse, error)
	mustEmbedUnimplementedRegistrationServiceServer()
}

// UnimplementedRegistrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRegistrationServiceServer struct {
}

func (UnimplementedRegistrationServiceServer) Register(context.Context, *InformationRequest) (*AnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRegistrationServiceServer) mustEmbedUnimplementedRegistrationServiceServer() {}

// UnsafeRegistrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistrationServiceServer will
// result in compilation errors.
type UnsafeRegistrationServiceServer interface {
	mustEmbedUnimplementedRegistrationServiceServer()
}

func RegisterRegistrationServiceServer(s grpc.ServiceRegistrar, srv RegistrationServiceServer) {
	s.RegisterService(&RegistrationService_ServiceDesc, srv)
}

func _RegistrationService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.protobuf.greeting.RegistrationService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServiceServer).Register(ctx, req.(*InformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegistrationService_ServiceDesc is the grpc.ServiceDesc for RegistrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegistrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.protobuf.greeting.RegistrationService",
	HandlerType: (*RegistrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _RegistrationService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greeting.proto",
}
