// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: promotion.proto

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
	Promotion_AddPromoProductInfo_FullMethodName    = "/promotion.Promotion/AddPromoProductInfo"
	Promotion_GetPromoProductsInfo_FullMethodName   = "/promotion.Promotion/GetPromoProductsInfo"
	Promotion_DeletePromoProductInfo_FullMethodName = "/promotion.Promotion/DeletePromoProductInfo"
)

// PromotionClient is the client API for Promotion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PromotionClient interface {
	AddPromoProductInfo(ctx context.Context, in *AddPromoProductRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetPromoProductsInfo(ctx context.Context, in *GetPromoProductsRequest, opts ...grpc.CallOption) (*GetPromoProductsResponse, error)
	DeletePromoProductInfo(ctx context.Context, in *DeletePromoProductRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type promotionClient struct {
	cc grpc.ClientConnInterface
}

func NewPromotionClient(cc grpc.ClientConnInterface) PromotionClient {
	return &promotionClient{cc}
}

func (c *promotionClient) AddPromoProductInfo(ctx context.Context, in *AddPromoProductRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Promotion_AddPromoProductInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionClient) GetPromoProductsInfo(ctx context.Context, in *GetPromoProductsRequest, opts ...grpc.CallOption) (*GetPromoProductsResponse, error) {
	out := new(GetPromoProductsResponse)
	err := c.cc.Invoke(ctx, Promotion_GetPromoProductsInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionClient) DeletePromoProductInfo(ctx context.Context, in *DeletePromoProductRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Promotion_DeletePromoProductInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromotionServer is the server API for Promotion service.
// All implementations must embed UnimplementedPromotionServer
// for forward compatibility
type PromotionServer interface {
	AddPromoProductInfo(context.Context, *AddPromoProductRequest) (*empty.Empty, error)
	GetPromoProductsInfo(context.Context, *GetPromoProductsRequest) (*GetPromoProductsResponse, error)
	DeletePromoProductInfo(context.Context, *DeletePromoProductRequest) (*empty.Empty, error)
	mustEmbedUnimplementedPromotionServer()
}

// UnimplementedPromotionServer must be embedded to have forward compatible implementations.
type UnimplementedPromotionServer struct {
}

func (UnimplementedPromotionServer) AddPromoProductInfo(context.Context, *AddPromoProductRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPromoProductInfo not implemented")
}
func (UnimplementedPromotionServer) GetPromoProductsInfo(context.Context, *GetPromoProductsRequest) (*GetPromoProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPromoProductsInfo not implemented")
}
func (UnimplementedPromotionServer) DeletePromoProductInfo(context.Context, *DeletePromoProductRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePromoProductInfo not implemented")
}
func (UnimplementedPromotionServer) mustEmbedUnimplementedPromotionServer() {}

// UnsafePromotionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PromotionServer will
// result in compilation errors.
type UnsafePromotionServer interface {
	mustEmbedUnimplementedPromotionServer()
}

func RegisterPromotionServer(s grpc.ServiceRegistrar, srv PromotionServer) {
	s.RegisterService(&Promotion_ServiceDesc, srv)
}

func _Promotion_AddPromoProductInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPromoProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServer).AddPromoProductInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Promotion_AddPromoProductInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServer).AddPromoProductInfo(ctx, req.(*AddPromoProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Promotion_GetPromoProductsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPromoProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServer).GetPromoProductsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Promotion_GetPromoProductsInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServer).GetPromoProductsInfo(ctx, req.(*GetPromoProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Promotion_DeletePromoProductInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePromoProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServer).DeletePromoProductInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Promotion_DeletePromoProductInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServer).DeletePromoProductInfo(ctx, req.(*DeletePromoProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Promotion_ServiceDesc is the grpc.ServiceDesc for Promotion service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Promotion_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "promotion.Promotion",
	HandlerType: (*PromotionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPromoProductInfo",
			Handler:    _Promotion_AddPromoProductInfo_Handler,
		},
		{
			MethodName: "GetPromoProductsInfo",
			Handler:    _Promotion_GetPromoProductsInfo_Handler,
		},
		{
			MethodName: "DeletePromoProductInfo",
			Handler:    _Promotion_DeletePromoProductInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "promotion.proto",
}
