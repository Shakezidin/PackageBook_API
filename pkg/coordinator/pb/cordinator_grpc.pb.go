// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: cordinator.proto

package __

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
	Coordinator_CoordinatorLoginRequest_FullMethodName        = "/pb.Coordinator/CoordinatorLoginRequest"
	Coordinator_CoordinatorSignupRequest_FullMethodName       = "/pb.Coordinator/CoordinatorSignupRequest"
	Coordinator_CoordinatorSignupVerifyRequest_FullMethodName = "/pb.Coordinator/CoordinatorSignupVerifyRequest"
	Coordinator_CoordinatorAddPackage_FullMethodName          = "/pb.Coordinator/CoordinatorAddPackage"
	Coordinator_CoordinatorAddDestination_FullMethodName      = "/pb.Coordinator/CoordinatorAddDestination"
	Coordinator_CoordinatorAddActivity_FullMethodName         = "/pb.Coordinator/CoordinatorAddActivity"
)

// CoordinatorClient is the client API for Coordinator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoordinatorClient interface {
	CoordinatorLoginRequest(ctx context.Context, in *CoorinatorLogin, opts ...grpc.CallOption) (*CordinatorLoginResponce, error)
	CoordinatorSignupRequest(ctx context.Context, in *Signup, opts ...grpc.CallOption) (*SignupResponce, error)
	CoordinatorSignupVerifyRequest(ctx context.Context, in *Verify, opts ...grpc.CallOption) (*VerifyResponce, error)
	CoordinatorAddPackage(ctx context.Context, in *AddPackage, opts ...grpc.CallOption) (*AddPackageResponce, error)
	CoordinatorAddDestination(ctx context.Context, in *AddDestination, opts ...grpc.CallOption) (*AddDestinationResponce, error)
	CoordinatorAddActivity(ctx context.Context, in *AddActivity, opts ...grpc.CallOption) (*AddActivityResponce, error)
}

type coordinatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCoordinatorClient(cc grpc.ClientConnInterface) CoordinatorClient {
	return &coordinatorClient{cc}
}

func (c *coordinatorClient) CoordinatorLoginRequest(ctx context.Context, in *CoorinatorLogin, opts ...grpc.CallOption) (*CordinatorLoginResponce, error) {
	out := new(CordinatorLoginResponce)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorLoginRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorSignupRequest(ctx context.Context, in *Signup, opts ...grpc.CallOption) (*SignupResponce, error) {
	out := new(SignupResponce)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorSignupRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorSignupVerifyRequest(ctx context.Context, in *Verify, opts ...grpc.CallOption) (*VerifyResponce, error) {
	out := new(VerifyResponce)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorSignupVerifyRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorAddPackage(ctx context.Context, in *AddPackage, opts ...grpc.CallOption) (*AddPackageResponce, error) {
	out := new(AddPackageResponce)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorAddPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorAddDestination(ctx context.Context, in *AddDestination, opts ...grpc.CallOption) (*AddDestinationResponce, error) {
	out := new(AddDestinationResponce)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorAddDestination_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorAddActivity(ctx context.Context, in *AddActivity, opts ...grpc.CallOption) (*AddActivityResponce, error) {
	out := new(AddActivityResponce)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorAddActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinatorServer is the server API for Coordinator service.
// All implementations must embed UnimplementedCoordinatorServer
// for forward compatibility
type CoordinatorServer interface {
	CoordinatorLoginRequest(context.Context, *CoorinatorLogin) (*CordinatorLoginResponce, error)
	CoordinatorSignupRequest(context.Context, *Signup) (*SignupResponce, error)
	CoordinatorSignupVerifyRequest(context.Context, *Verify) (*VerifyResponce, error)
	CoordinatorAddPackage(context.Context, *AddPackage) (*AddPackageResponce, error)
	CoordinatorAddDestination(context.Context, *AddDestination) (*AddDestinationResponce, error)
	CoordinatorAddActivity(context.Context, *AddActivity) (*AddActivityResponce, error)
	mustEmbedUnimplementedCoordinatorServer()
}

// UnimplementedCoordinatorServer must be embedded to have forward compatible implementations.
type UnimplementedCoordinatorServer struct {
}

func (UnimplementedCoordinatorServer) CoordinatorLoginRequest(context.Context, *CoorinatorLogin) (*CordinatorLoginResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorLoginRequest not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorSignupRequest(context.Context, *Signup) (*SignupResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorSignupRequest not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorSignupVerifyRequest(context.Context, *Verify) (*VerifyResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorSignupVerifyRequest not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorAddPackage(context.Context, *AddPackage) (*AddPackageResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorAddPackage not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorAddDestination(context.Context, *AddDestination) (*AddDestinationResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorAddDestination not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorAddActivity(context.Context, *AddActivity) (*AddActivityResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorAddActivity not implemented")
}
func (UnimplementedCoordinatorServer) mustEmbedUnimplementedCoordinatorServer() {}

// UnsafeCoordinatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoordinatorServer will
// result in compilation errors.
type UnsafeCoordinatorServer interface {
	mustEmbedUnimplementedCoordinatorServer()
}

func RegisterCoordinatorServer(s grpc.ServiceRegistrar, srv CoordinatorServer) {
	s.RegisterService(&Coordinator_ServiceDesc, srv)
}

func _Coordinator_CoordinatorLoginRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoorinatorLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorLoginRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorLoginRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorLoginRequest(ctx, req.(*CoorinatorLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorSignupRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Signup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorSignupRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorSignupRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorSignupRequest(ctx, req.(*Signup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorSignupVerifyRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Verify)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorSignupVerifyRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorSignupVerifyRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorSignupVerifyRequest(ctx, req.(*Verify))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorAddPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPackage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorAddPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorAddPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorAddPackage(ctx, req.(*AddPackage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorAddDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDestination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorAddDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorAddDestination_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorAddDestination(ctx, req.(*AddDestination))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorAddActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddActivity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorAddActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorAddActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorAddActivity(ctx, req.(*AddActivity))
	}
	return interceptor(ctx, in, info, handler)
}

// Coordinator_ServiceDesc is the grpc.ServiceDesc for Coordinator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Coordinator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Coordinator",
	HandlerType: (*CoordinatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CoordinatorLoginRequest",
			Handler:    _Coordinator_CoordinatorLoginRequest_Handler,
		},
		{
			MethodName: "CoordinatorSignupRequest",
			Handler:    _Coordinator_CoordinatorSignupRequest_Handler,
		},
		{
			MethodName: "CoordinatorSignupVerifyRequest",
			Handler:    _Coordinator_CoordinatorSignupVerifyRequest_Handler,
		},
		{
			MethodName: "CoordinatorAddPackage",
			Handler:    _Coordinator_CoordinatorAddPackage_Handler,
		},
		{
			MethodName: "CoordinatorAddDestination",
			Handler:    _Coordinator_CoordinatorAddDestination_Handler,
		},
		{
			MethodName: "CoordinatorAddActivity",
			Handler:    _Coordinator_CoordinatorAddActivity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cordinator.proto",
}
