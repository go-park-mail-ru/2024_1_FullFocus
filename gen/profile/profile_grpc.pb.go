// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: profile.proto

package gen

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Profile_CreateProfile_FullMethodName               = "/profile.Profile/CreateProfile"
	Profile_GetProfileByID_FullMethodName              = "/profile.Profile/GetProfileByID"
	Profile_GetProfileNamesByIDs_FullMethodName        = "/profile.Profile/GetProfileNamesByIDs"
	Profile_GetProfileMetaInfo_FullMethodName          = "/profile.Profile/GetProfileMetaInfo"
	Profile_GetAvatarByID_FullMethodName               = "/profile.Profile/GetAvatarByID"
	Profile_UpdateAvatarByProfileID_FullMethodName     = "/profile.Profile/UpdateAvatarByProfileID"
	Profile_UpdateProfile_FullMethodName               = "/profile.Profile/UpdateProfile"
	Profile_DeleteAvatarByProfileID_FullMethodName     = "/profile.Profile/DeleteAvatarByProfileID"
	Profile_GetProfileNamesAvatarsByIDs_FullMethodName = "/profile.Profile/GetProfileNamesAvatarsByIDs"
)

// ProfileClient is the client API for Profile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileClient interface {
	CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetProfileByID(ctx context.Context, in *GetProfileByIDRequest, opts ...grpc.CallOption) (*GetProfileByIDResponse, error)
	GetProfileNamesByIDs(ctx context.Context, in *GetProfileNamesByIDsRequest, opts ...grpc.CallOption) (*GetProfileNamesByIDsResponse, error)
	GetProfileMetaInfo(ctx context.Context, in *GetProfileMetaInfoRequest, opts ...grpc.CallOption) (*GetProfileMetaInfoResponse, error)
	GetAvatarByID(ctx context.Context, in *GetAvatarByIDRequest, opts ...grpc.CallOption) (*GetAvatarByIDResponse, error)
	UpdateAvatarByProfileID(ctx context.Context, in *UpdateAvatarByProfileIDRequest, opts ...grpc.CallOption) (*UpdateAvatarByProfileIDResponse, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteAvatarByProfileID(ctx context.Context, in *DeleteAvatarByProfileIDRequest, opts ...grpc.CallOption) (*DeleteAvatarByProfileIDResponse, error)
	GetProfileNamesAvatarsByIDs(ctx context.Context, in *GetProfileNamesAvatarsRequest, opts ...grpc.CallOption) (*GetProfileNamesAvatarsResponse, error)
}

type profileClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileClient(cc grpc.ClientConnInterface) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Profile_CreateProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetProfileByID(ctx context.Context, in *GetProfileByIDRequest, opts ...grpc.CallOption) (*GetProfileByIDResponse, error) {
	out := new(GetProfileByIDResponse)
	err := c.cc.Invoke(ctx, Profile_GetProfileByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetProfileNamesByIDs(ctx context.Context, in *GetProfileNamesByIDsRequest, opts ...grpc.CallOption) (*GetProfileNamesByIDsResponse, error) {
	out := new(GetProfileNamesByIDsResponse)
	err := c.cc.Invoke(ctx, Profile_GetProfileNamesByIDs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetProfileMetaInfo(ctx context.Context, in *GetProfileMetaInfoRequest, opts ...grpc.CallOption) (*GetProfileMetaInfoResponse, error) {
	out := new(GetProfileMetaInfoResponse)
	err := c.cc.Invoke(ctx, Profile_GetProfileMetaInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetAvatarByID(ctx context.Context, in *GetAvatarByIDRequest, opts ...grpc.CallOption) (*GetAvatarByIDResponse, error) {
	out := new(GetAvatarByIDResponse)
	err := c.cc.Invoke(ctx, Profile_GetAvatarByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) UpdateAvatarByProfileID(ctx context.Context, in *UpdateAvatarByProfileIDRequest, opts ...grpc.CallOption) (*UpdateAvatarByProfileIDResponse, error) {
	out := new(UpdateAvatarByProfileIDResponse)
	err := c.cc.Invoke(ctx, Profile_UpdateAvatarByProfileID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Profile_UpdateProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) DeleteAvatarByProfileID(ctx context.Context, in *DeleteAvatarByProfileIDRequest, opts ...grpc.CallOption) (*DeleteAvatarByProfileIDResponse, error) {
	out := new(DeleteAvatarByProfileIDResponse)
	err := c.cc.Invoke(ctx, Profile_DeleteAvatarByProfileID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetProfileNamesAvatarsByIDs(ctx context.Context, in *GetProfileNamesAvatarsRequest, opts ...grpc.CallOption) (*GetProfileNamesAvatarsResponse, error) {
	out := new(GetProfileNamesAvatarsResponse)
	err := c.cc.Invoke(ctx, Profile_GetProfileNamesAvatarsByIDs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServer is the server API for Profile service.
// All implementations must embed UnimplementedProfileServer
// for forward compatibility
type ProfileServer interface {
	CreateProfile(context.Context, *CreateProfileRequest) (*empty.Empty, error)
	GetProfileByID(context.Context, *GetProfileByIDRequest) (*GetProfileByIDResponse, error)
	GetProfileNamesByIDs(context.Context, *GetProfileNamesByIDsRequest) (*GetProfileNamesByIDsResponse, error)
	GetProfileMetaInfo(context.Context, *GetProfileMetaInfoRequest) (*GetProfileMetaInfoResponse, error)
	GetAvatarByID(context.Context, *GetAvatarByIDRequest) (*GetAvatarByIDResponse, error)
	UpdateAvatarByProfileID(context.Context, *UpdateAvatarByProfileIDRequest) (*UpdateAvatarByProfileIDResponse, error)
	UpdateProfile(context.Context, *UpdateProfileRequest) (*empty.Empty, error)
	DeleteAvatarByProfileID(context.Context, *DeleteAvatarByProfileIDRequest) (*DeleteAvatarByProfileIDResponse, error)
	GetProfileNamesAvatarsByIDs(context.Context, *GetProfileNamesAvatarsRequest) (*GetProfileNamesAvatarsResponse, error)
	mustEmbedUnimplementedProfileServer()
}

// UnimplementedProfileServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServer struct {
}

func (UnimplementedProfileServer) CreateProfile(context.Context, *CreateProfileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfile not implemented")
}
func (UnimplementedProfileServer) GetProfileByID(context.Context, *GetProfileByIDRequest) (*GetProfileByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileByID not implemented")
}
func (UnimplementedProfileServer) GetProfileNamesByIDs(context.Context, *GetProfileNamesByIDsRequest) (*GetProfileNamesByIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileNamesByIDs not implemented")
}
func (UnimplementedProfileServer) GetProfileMetaInfo(context.Context, *GetProfileMetaInfoRequest) (*GetProfileMetaInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileMetaInfo not implemented")
}
func (UnimplementedProfileServer) GetAvatarByID(context.Context, *GetAvatarByIDRequest) (*GetAvatarByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvatarByID not implemented")
}
func (UnimplementedProfileServer) UpdateAvatarByProfileID(context.Context, *UpdateAvatarByProfileIDRequest) (*UpdateAvatarByProfileIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAvatarByProfileID not implemented")
}
func (UnimplementedProfileServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedProfileServer) DeleteAvatarByProfileID(context.Context, *DeleteAvatarByProfileIDRequest) (*DeleteAvatarByProfileIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAvatarByProfileID not implemented")
}
func (UnimplementedProfileServer) GetProfileNamesAvatarsByIDs(context.Context, *GetProfileNamesAvatarsRequest) (*GetProfileNamesAvatarsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileNamesAvatarsByIDs not implemented")
}
func (UnimplementedProfileServer) mustEmbedUnimplementedProfileServer() {}

// UnsafeProfileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServer will
// result in compilation errors.
type UnsafeProfileServer interface {
	mustEmbedUnimplementedProfileServer()
}

func RegisterProfileServer(s grpc.ServiceRegistrar, srv ProfileServer) {
	s.RegisterService(&Profile_ServiceDesc, srv)
}

func _Profile_CreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).CreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_CreateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).CreateProfile(ctx, req.(*CreateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetProfileByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetProfileByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetProfileByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetProfileByID(ctx, req.(*GetProfileByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetProfileNamesByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileNamesByIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetProfileNamesByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetProfileNamesByIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetProfileNamesByIDs(ctx, req.(*GetProfileNamesByIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetProfileMetaInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileMetaInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetProfileMetaInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetProfileMetaInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetProfileMetaInfo(ctx, req.(*GetProfileMetaInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetAvatarByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvatarByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetAvatarByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetAvatarByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetAvatarByID(ctx, req.(*GetAvatarByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_UpdateAvatarByProfileID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAvatarByProfileIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).UpdateAvatarByProfileID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_UpdateAvatarByProfileID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).UpdateAvatarByProfileID(ctx, req.(*UpdateAvatarByProfileIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_UpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_DeleteAvatarByProfileID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAvatarByProfileIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).DeleteAvatarByProfileID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_DeleteAvatarByProfileID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).DeleteAvatarByProfileID(ctx, req.(*DeleteAvatarByProfileIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetProfileNamesAvatarsByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileNamesAvatarsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetProfileNamesAvatarsByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetProfileNamesAvatarsByIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetProfileNamesAvatarsByIDs(ctx, req.(*GetProfileNamesAvatarsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Profile_ServiceDesc is the grpc.ServiceDesc for Profile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Profile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profile.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProfile",
			Handler:    _Profile_CreateProfile_Handler,
		},
		{
			MethodName: "GetProfileByID",
			Handler:    _Profile_GetProfileByID_Handler,
		},
		{
			MethodName: "GetProfileNamesByIDs",
			Handler:    _Profile_GetProfileNamesByIDs_Handler,
		},
		{
			MethodName: "GetProfileMetaInfo",
			Handler:    _Profile_GetProfileMetaInfo_Handler,
		},
		{
			MethodName: "GetAvatarByID",
			Handler:    _Profile_GetAvatarByID_Handler,
		},
		{
			MethodName: "UpdateAvatarByProfileID",
			Handler:    _Profile_UpdateAvatarByProfileID_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _Profile_UpdateProfile_Handler,
		},
		{
			MethodName: "DeleteAvatarByProfileID",
			Handler:    _Profile_DeleteAvatarByProfileID_Handler,
		},
		{
			MethodName: "GetProfileNamesAvatarsByIDs",
			Handler:    _Profile_GetProfileNamesAvatarsByIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}
