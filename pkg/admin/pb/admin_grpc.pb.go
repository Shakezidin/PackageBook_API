// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: admin.proto

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
	Admin_AdminLoginRequest_FullMethodName     = "/pb.Admin/AdminLoginRequest"
	Admin_AdminAddCategory_FullMethodName      = "/pb.Admin/AdminAddCategory"
	Admin_AdminViewPackages_FullMethodName     = "/pb.Admin/AdminViewPackages"
	Admin_AdminViewpackage_FullMethodName      = "/pb.Admin/AdminViewpackage"
	Admin_AdminPacakgeStatus_FullMethodName    = "/pb.Admin/AdminPacakgeStatus"
	Admin_AdminViewCategories_FullMethodName   = "/pb.Admin/AdminViewCategories"
	Admin_AdminViewDestination_FullMethodName  = "/pb.Admin/AdminViewDestination"
	Admin_AdminViewActivity_FullMethodName     = "/pb.Admin/AdminViewActivity"
	Admin_AdminViewCoordinators_FullMethodName = "/pb.Admin/AdminViewCoordinators"
	Admin_AdminViewBookings_FullMethodName     = "/pb.Admin/AdminViewBookings"
	Admin_AdminViewBooking_FullMethodName      = "/pb.Admin/AdminViewBooking"
	Admin_AdminViewDashboard_FullMethodName    = "/pb.Admin/AdminViewDashboard"
	Admin_AdminSearchBooking_FullMethodName    = "/pb.Admin/AdminSearchBooking"
	Admin_AdminViewUsers_FullMethodName        = "/pb.Admin/AdminViewUsers"
	Admin_AdminViewUser_FullMethodName         = "/pb.Admin/AdminViewUser"
)

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	AdminLoginRequest(ctx context.Context, in *AdminLogin, opts ...grpc.CallOption) (*AdminResponse, error)
	AdminAddCategory(ctx context.Context, in *AdminCategory, opts ...grpc.CallOption) (*AdminResponse, error)
	AdminViewPackages(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminPackages, error)
	AdminViewpackage(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminPackage, error)
	AdminPacakgeStatus(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminResponse, error)
	AdminViewCategories(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminCategories, error)
	AdminViewDestination(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminDestination, error)
	AdminViewActivity(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminActivity, error)
	AdminViewCoordinators(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminUsers, error)
	AdminViewBookings(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminHistories, error)
	AdminViewBooking(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminHistory, error)
	AdminViewDashboard(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminDashboard, error)
	AdminSearchBooking(ctx context.Context, in *AdminBookingSearchCriteria, opts ...grpc.CallOption) (*AdminHistories, error)
	AdminViewUsers(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminUsers, error)
	AdminViewUser(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminUser, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) AdminLoginRequest(ctx context.Context, in *AdminLogin, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, Admin_AdminLoginRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AdminAddCategory(ctx context.Context, in *AdminCategory, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, Admin_AdminAddCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AdminViewPackages(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminPackages, error) {
	out := new(AdminPackages)
	err := c.cc.Invoke(ctx, Admin_AdminViewPackages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AdminViewpackage(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminPackage, error) {
	out := new(AdminPackage)
	err := c.cc.Invoke(ctx, Admin_AdminViewpackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AdminPacakgeStatus(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, Admin_AdminPacakgeStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AdminViewCategories(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminCategories, error) {
	out := new(AdminCategories)
	err := c.cc.Invoke(ctx, Admin_AdminViewCategories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AdminViewDestination(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminDestination, error) {
	out := new(AdminDestination)
	err := c.cc.Invoke(ctx, Admin_AdminViewDestination_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AdminViewActivity(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminActivity, error) {
	out := new(AdminActivity)
	err := c.cc.Invoke(ctx, Admin_AdminViewActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AdminViewCoordinators(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminUsers, error) {
	out := new(AdminUsers)
	err := c.cc.Invoke(ctx, Admin_AdminViewCoordinators_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AdminViewBookings(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminHistories, error) {
	out := new(AdminHistories)
	err := c.cc.Invoke(ctx, Admin_AdminViewBookings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AdminViewBooking(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminHistory, error) {
	out := new(AdminHistory)
	err := c.cc.Invoke(ctx, Admin_AdminViewBooking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AdminViewDashboard(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminDashboard, error) {
	out := new(AdminDashboard)
	err := c.cc.Invoke(ctx, Admin_AdminViewDashboard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AdminSearchBooking(ctx context.Context, in *AdminBookingSearchCriteria, opts ...grpc.CallOption) (*AdminHistories, error) {
	out := new(AdminHistories)
	err := c.cc.Invoke(ctx, Admin_AdminSearchBooking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AdminViewUsers(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminUsers, error) {
	out := new(AdminUsers)
	err := c.cc.Invoke(ctx, Admin_AdminViewUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AdminViewUser(ctx context.Context, in *AdminView, opts ...grpc.CallOption) (*AdminUser, error) {
	out := new(AdminUser)
	err := c.cc.Invoke(ctx, Admin_AdminViewUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	AdminLoginRequest(context.Context, *AdminLogin) (*AdminResponse, error)
	AdminAddCategory(context.Context, *AdminCategory) (*AdminResponse, error)
	AdminViewPackages(context.Context, *AdminView) (*AdminPackages, error)
	AdminViewpackage(context.Context, *AdminView) (*AdminPackage, error)
	AdminPacakgeStatus(context.Context, *AdminView) (*AdminResponse, error)
	AdminViewCategories(context.Context, *AdminView) (*AdminCategories, error)
	AdminViewDestination(context.Context, *AdminView) (*AdminDestination, error)
	AdminViewActivity(context.Context, *AdminView) (*AdminActivity, error)
	AdminViewCoordinators(context.Context, *AdminView) (*AdminUsers, error)
	AdminViewBookings(context.Context, *AdminView) (*AdminHistories, error)
	AdminViewBooking(context.Context, *AdminView) (*AdminHistory, error)
	AdminViewDashboard(context.Context, *AdminView) (*AdminDashboard, error)
	AdminSearchBooking(context.Context, *AdminBookingSearchCriteria) (*AdminHistories, error)
	AdminViewUsers(context.Context, *AdminView) (*AdminUsers, error)
	AdminViewUser(context.Context, *AdminView) (*AdminUser, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) AdminLoginRequest(context.Context, *AdminLogin) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLoginRequest not implemented")
}
func (UnimplementedAdminServer) AdminAddCategory(context.Context, *AdminCategory) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminAddCategory not implemented")
}
func (UnimplementedAdminServer) AdminViewPackages(context.Context, *AdminView) (*AdminPackages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminViewPackages not implemented")
}
func (UnimplementedAdminServer) AdminViewpackage(context.Context, *AdminView) (*AdminPackage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminViewpackage not implemented")
}
func (UnimplementedAdminServer) AdminPacakgeStatus(context.Context, *AdminView) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminPacakgeStatus not implemented")
}
func (UnimplementedAdminServer) AdminViewCategories(context.Context, *AdminView) (*AdminCategories, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminViewCategories not implemented")
}
func (UnimplementedAdminServer) AdminViewDestination(context.Context, *AdminView) (*AdminDestination, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminViewDestination not implemented")
}
func (UnimplementedAdminServer) AdminViewActivity(context.Context, *AdminView) (*AdminActivity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminViewActivity not implemented")
}
func (UnimplementedAdminServer) AdminViewCoordinators(context.Context, *AdminView) (*AdminUsers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminViewCoordinators not implemented")
}
func (UnimplementedAdminServer) AdminViewBookings(context.Context, *AdminView) (*AdminHistories, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminViewBookings not implemented")
}
func (UnimplementedAdminServer) AdminViewBooking(context.Context, *AdminView) (*AdminHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminViewBooking not implemented")
}
func (UnimplementedAdminServer) AdminViewDashboard(context.Context, *AdminView) (*AdminDashboard, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminViewDashboard not implemented")
}
func (UnimplementedAdminServer) AdminSearchBooking(context.Context, *AdminBookingSearchCriteria) (*AdminHistories, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminSearchBooking not implemented")
}
func (UnimplementedAdminServer) AdminViewUsers(context.Context, *AdminView) (*AdminUsers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminViewUsers not implemented")
}
func (UnimplementedAdminServer) AdminViewUser(context.Context, *AdminView) (*AdminUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminViewUser not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_AdminLoginRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminLoginRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminLoginRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminLoginRequest(ctx, req.(*AdminLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AdminAddCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminAddCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminAddCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminAddCategory(ctx, req.(*AdminCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AdminViewPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminViewPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminViewPackages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminViewPackages(ctx, req.(*AdminView))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AdminViewpackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminViewpackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminViewpackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminViewpackage(ctx, req.(*AdminView))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AdminPacakgeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminPacakgeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminPacakgeStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminPacakgeStatus(ctx, req.(*AdminView))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AdminViewCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminViewCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminViewCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminViewCategories(ctx, req.(*AdminView))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AdminViewDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminViewDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminViewDestination_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminViewDestination(ctx, req.(*AdminView))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AdminViewActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminViewActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminViewActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminViewActivity(ctx, req.(*AdminView))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AdminViewCoordinators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminViewCoordinators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminViewCoordinators_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminViewCoordinators(ctx, req.(*AdminView))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AdminViewBookings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminViewBookings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminViewBookings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminViewBookings(ctx, req.(*AdminView))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AdminViewBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminViewBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminViewBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminViewBooking(ctx, req.(*AdminView))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AdminViewDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminViewDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminViewDashboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminViewDashboard(ctx, req.(*AdminView))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AdminSearchBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminBookingSearchCriteria)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminSearchBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminSearchBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminSearchBooking(ctx, req.(*AdminBookingSearchCriteria))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AdminViewUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminViewUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminViewUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminViewUsers(ctx, req.(*AdminView))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AdminViewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminViewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminViewUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminViewUser(ctx, req.(*AdminView))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminLoginRequest",
			Handler:    _Admin_AdminLoginRequest_Handler,
		},
		{
			MethodName: "AdminAddCategory",
			Handler:    _Admin_AdminAddCategory_Handler,
		},
		{
			MethodName: "AdminViewPackages",
			Handler:    _Admin_AdminViewPackages_Handler,
		},
		{
			MethodName: "AdminViewpackage",
			Handler:    _Admin_AdminViewpackage_Handler,
		},
		{
			MethodName: "AdminPacakgeStatus",
			Handler:    _Admin_AdminPacakgeStatus_Handler,
		},
		{
			MethodName: "AdminViewCategories",
			Handler:    _Admin_AdminViewCategories_Handler,
		},
		{
			MethodName: "AdminViewDestination",
			Handler:    _Admin_AdminViewDestination_Handler,
		},
		{
			MethodName: "AdminViewActivity",
			Handler:    _Admin_AdminViewActivity_Handler,
		},
		{
			MethodName: "AdminViewCoordinators",
			Handler:    _Admin_AdminViewCoordinators_Handler,
		},
		{
			MethodName: "AdminViewBookings",
			Handler:    _Admin_AdminViewBookings_Handler,
		},
		{
			MethodName: "AdminViewBooking",
			Handler:    _Admin_AdminViewBooking_Handler,
		},
		{
			MethodName: "AdminViewDashboard",
			Handler:    _Admin_AdminViewDashboard_Handler,
		},
		{
			MethodName: "AdminSearchBooking",
			Handler:    _Admin_AdminSearchBooking_Handler,
		},
		{
			MethodName: "AdminViewUsers",
			Handler:    _Admin_AdminViewUsers_Handler,
		},
		{
			MethodName: "AdminViewUser",
			Handler:    _Admin_AdminViewUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}
