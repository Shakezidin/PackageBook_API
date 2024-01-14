// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: user.proto

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
	User_UserLoginRequest_FullMethodName        = "/pb.User/UserLoginRequest"
	User_UserSignupRequest_FullMethodName       = "/pb.User/UserSignupRequest"
	User_UserSignupVerifyRequest_FullMethodName = "/pb.User/UserSignupVerifyRequest"
	User_UserViewPackage_FullMethodName         = "/pb.User/UserViewPackage"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	UserLoginRequest(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*UserLoginResponce, error)
	UserSignupRequest(ctx context.Context, in *Signup, opts ...grpc.CallOption) (*SignupResponce, error)
	UserSignupVerifyRequest(ctx context.Context, in *Verify, opts ...grpc.CallOption) (*VerifyResponce, error)
	UserViewPackage(ctx context.Context, in *ViewPackage, opts ...grpc.CallOption) (*ViewPacakgeResponce, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) UserLoginRequest(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*UserLoginResponce, error) {
	out := new(UserLoginResponce)
	err := c.cc.Invoke(ctx, User_UserLoginRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserSignupRequest(ctx context.Context, in *Signup, opts ...grpc.CallOption) (*SignupResponce, error) {
	out := new(SignupResponce)
	err := c.cc.Invoke(ctx, User_UserSignupRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserSignupVerifyRequest(ctx context.Context, in *Verify, opts ...grpc.CallOption) (*VerifyResponce, error) {
	out := new(VerifyResponce)
	err := c.cc.Invoke(ctx, User_UserSignupVerifyRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserViewPackage(ctx context.Context, in *ViewPackage, opts ...grpc.CallOption) (*ViewPacakgeResponce, error) {
	out := new(ViewPacakgeResponce)
	err := c.cc.Invoke(ctx, User_UserViewPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	UserLoginRequest(context.Context, *UserLogin) (*UserLoginResponce, error)
	UserSignupRequest(context.Context, *Signup) (*SignupResponce, error)
	UserSignupVerifyRequest(context.Context, *Verify) (*VerifyResponce, error)
	UserViewPackage(context.Context, *ViewPackage) (*ViewPacakgeResponce, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) UserLoginRequest(context.Context, *UserLogin) (*UserLoginResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLoginRequest not implemented")
}
func (UnimplementedUserServer) UserSignupRequest(context.Context, *Signup) (*SignupResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignupRequest not implemented")
}
func (UnimplementedUserServer) UserSignupVerifyRequest(context.Context, *Verify) (*VerifyResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignupVerifyRequest not implemented")
}
func (UnimplementedUserServer) UserViewPackage(context.Context, *ViewPackage) (*ViewPacakgeResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserViewPackage not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_UserLoginRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserLoginRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserLoginRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLoginRequest(ctx, req.(*UserLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserSignupRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Signup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserSignupRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserSignupRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserSignupRequest(ctx, req.(*Signup))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserSignupVerifyRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Verify)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserSignupVerifyRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserSignupVerifyRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserSignupVerifyRequest(ctx, req.(*Verify))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserViewPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewPackage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserViewPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserViewPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserViewPackage(ctx, req.(*ViewPackage))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLoginRequest",
			Handler:    _User_UserLoginRequest_Handler,
		},
		{
			MethodName: "UserSignupRequest",
			Handler:    _User_UserSignupRequest_Handler,
		},
		{
			MethodName: "UserSignupVerifyRequest",
			Handler:    _User_UserSignupVerifyRequest_Handler,
		},
		{
			MethodName: "UserViewPackage",
			Handler:    _User_UserViewPackage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
