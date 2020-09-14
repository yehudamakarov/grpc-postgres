// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package users

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*User, error)
	AddUsers(ctx context.Context, opts ...grpc.CallOption) (UserService_AddUsersClient, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*User, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (UserService_ListUsersClient, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

var userServiceAddUserStreamDesc = &grpc.StreamDesc{
	StreamName: "AddUser",
}

func (c *userServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/users.UserService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var userServiceAddUsersStreamDesc = &grpc.StreamDesc{
	StreamName:    "AddUsers",
	ClientStreams: true,
}

func (c *userServiceClient) AddUsers(ctx context.Context, opts ...grpc.CallOption) (UserService_AddUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, userServiceAddUsersStreamDesc, "/users.UserService/AddUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceAddUsersClient{stream}
	return x, nil
}

type UserService_AddUsersClient interface {
	Send(*AddUserRequest) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type userServiceAddUsersClient struct {
	grpc.ClientStream
}

func (x *userServiceAddUsersClient) Send(m *AddUserRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceAddUsersClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var userServiceGetUserStreamDesc = &grpc.StreamDesc{
	StreamName: "GetUser",
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/users.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var userServiceDeleteUserStreamDesc = &grpc.StreamDesc{
	StreamName: "DeleteUser",
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/users.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var userServiceListUsersStreamDesc = &grpc.StreamDesc{
	StreamName:    "ListUsers",
	ServerStreams: true,
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (UserService_ListUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, userServiceListUsersStreamDesc, "/users.UserService/ListUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceListUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_ListUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type userServiceListUsersClient struct {
	grpc.ClientStream
}

func (x *userServiceListUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceService is the service API for UserService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterUserServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type UserServiceService struct {
	AddUser    func(context.Context, *AddUserRequest) (*User, error)
	AddUsers   func(UserService_AddUsersServer) error
	GetUser    func(context.Context, *GetUserRequest) (*User, error)
	DeleteUser func(context.Context, *DeleteUserRequest) (*User, error)
	ListUsers  func(*ListUsersRequest, UserService_ListUsersServer) error
}

func (s *UserServiceService) addUser(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/users.UserService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *UserServiceService) addUsers(_ interface{}, stream grpc.ServerStream) error {
	return s.AddUsers(&userServiceAddUsersServer{stream})
}
func (s *UserServiceService) getUser(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/users.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *UserServiceService) deleteUser(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/users.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *UserServiceService) listUsers(_ interface{}, stream grpc.ServerStream) error {
	m := new(ListUsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return s.ListUsers(m, &userServiceListUsersServer{stream})
}

type UserService_AddUsersServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*AddUserRequest, error)
	grpc.ServerStream
}

type userServiceAddUsersServer struct {
	grpc.ServerStream
}

func (x *userServiceAddUsersServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceAddUsersServer) Recv() (*AddUserRequest, error) {
	m := new(AddUserRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type UserService_ListUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type userServiceListUsersServer struct {
	grpc.ServerStream
}

func (x *userServiceListUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

// RegisterUserServiceService registers a service implementation with a gRPC server.
func RegisterUserServiceService(s grpc.ServiceRegistrar, srv *UserServiceService) {
	srvCopy := *srv
	if srvCopy.AddUser == nil {
		srvCopy.AddUser = func(context.Context, *AddUserRequest) (*User, error) {
			return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
		}
	}
	if srvCopy.AddUsers == nil {
		srvCopy.AddUsers = func(UserService_AddUsersServer) error {
			return status.Errorf(codes.Unimplemented, "method AddUsers not implemented")
		}
	}
	if srvCopy.GetUser == nil {
		srvCopy.GetUser = func(context.Context, *GetUserRequest) (*User, error) {
			return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
		}
	}
	if srvCopy.DeleteUser == nil {
		srvCopy.DeleteUser = func(context.Context, *DeleteUserRequest) (*User, error) {
			return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
		}
	}
	if srvCopy.ListUsers == nil {
		srvCopy.ListUsers = func(*ListUsersRequest, UserService_ListUsersServer) error {
			return status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "users.UserService",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "AddUser",
				Handler:    srvCopy.addUser,
			},
			{
				MethodName: "GetUser",
				Handler:    srvCopy.getUser,
			},
			{
				MethodName: "DeleteUser",
				Handler:    srvCopy.deleteUser,
			},
		},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "AddUsers",
				Handler:       srvCopy.addUsers,
				ClientStreams: true,
			},
			{
				StreamName:    "ListUsers",
				Handler:       srvCopy.listUsers,
				ServerStreams: true,
			},
		},
		Metadata: "users.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewUserServiceService creates a new UserServiceService containing the
// implemented methods of the UserService service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewUserServiceService(s interface{}) *UserServiceService {
	ns := &UserServiceService{}
	if h, ok := s.(interface {
		AddUser(context.Context, *AddUserRequest) (*User, error)
	}); ok {
		ns.AddUser = h.AddUser
	}
	if h, ok := s.(interface {
		AddUsers(UserService_AddUsersServer) error
	}); ok {
		ns.AddUsers = h.AddUsers
	}
	if h, ok := s.(interface {
		GetUser(context.Context, *GetUserRequest) (*User, error)
	}); ok {
		ns.GetUser = h.GetUser
	}
	if h, ok := s.(interface {
		DeleteUser(context.Context, *DeleteUserRequest) (*User, error)
	}); ok {
		ns.DeleteUser = h.DeleteUser
	}
	if h, ok := s.(interface {
		ListUsers(*ListUsersRequest, UserService_ListUsersServer) error
	}); ok {
		ns.ListUsers = h.ListUsers
	}
	return ns
}

// UnstableUserServiceService is the service API for UserService service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableUserServiceService interface {
	AddUser(context.Context, *AddUserRequest) (*User, error)
	AddUsers(UserService_AddUsersServer) error
	GetUser(context.Context, *GetUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*User, error)
	ListUsers(*ListUsersRequest, UserService_ListUsersServer) error
}
