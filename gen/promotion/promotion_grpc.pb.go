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
	Promotion_AddPromoProduct_FullMethodName    = "/promotion.Promotion/AddPromoProduct"
	Promotion_GetPromoProducts_FullMethodName   = "/promotion.Promotion/GetPromoProducts"
	Promotion_DeletePromoProduct_FullMethodName = "/promotion.Promotion/DeletePromoProduct"
)

// PromotionClient is the client API for Promotion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PromotionClient interface {
	AddPromoProduct(ctx context.Context, in *AddPromoProductRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetPromoProducts(ctx context.Context, in *GetPromoProductsRequest, opts ...grpc.CallOption) (*GetPromoProductsResponse, error)
	DeletePromoProduct(ctx context.Context, in *DeletePromoProductRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type promotionClient struct {
	cc grpc.ClientConnInterface
}

func NewPromotionClient(cc grpc.ClientConnInterface) PromotionClient {
	return &promotionClient{cc}
}

func (c *promotionClient) AddPromoProduct(ctx context.Context, in *AddPromoProductRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Promotion_AddPromoProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionClient) GetPromoProducts(ctx context.Context, in *GetPromoProductsRequest, opts ...grpc.CallOption) (*GetPromoProductsResponse, error) {
	out := new(GetPromoProductsResponse)
	err := c.cc.Invoke(ctx, Promotion_GetPromoProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionClient) DeletePromoProduct(ctx context.Context, in *DeletePromoProductRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Promotion_DeletePromoProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromotionServer is the server API for Promotion service.
// All implementations must embed UnimplementedPromotionServer
// for forward compatibility
type PromotionServer interface {
	AddPromoProduct(context.Context, *AddPromoProductRequest) (*empty.Empty, error)
	GetPromoProducts(context.Context, *GetPromoProductsRequest) (*GetPromoProductsResponse, error)
	DeletePromoProduct(context.Context, *DeletePromoProductRequest) (*empty.Empty, error)
	mustEmbedUnimplementedPromotionServer()
}

// UnimplementedPromotionServer must be embedded to have forward compatible implementations.
type UnimplementedPromotionServer struct {
}

func (UnimplementedPromotionServer) AddPromoProduct(context.Context, *AddPromoProductRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPromoProduct not implemented")
}
func (UnimplementedPromotionServer) GetPromoProducts(context.Context, *GetPromoProductsRequest) (*GetPromoProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPromoProducts not implemented")
}
func (UnimplementedPromotionServer) DeletePromoProduct(context.Context, *DeletePromoProductRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePromoProduct not implemented")
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

func _Promotion_AddPromoProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPromoProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServer).AddPromoProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Promotion_AddPromoProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServer).AddPromoProduct(ctx, req.(*AddPromoProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Promotion_GetPromoProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPromoProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServer).GetPromoProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Promotion_GetPromoProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServer).GetPromoProducts(ctx, req.(*GetPromoProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Promotion_DeletePromoProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePromoProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServer).DeletePromoProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Promotion_DeletePromoProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServer).DeletePromoProduct(ctx, req.(*DeletePromoProductRequest))
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
			MethodName: "AddPromoProduct",
			Handler:    _Promotion_AddPromoProduct_Handler,
		},
		{
			MethodName: "GetPromoProducts",
			Handler:    _Promotion_GetPromoProducts_Handler,
		},
		{
			MethodName: "DeletePromoProduct",
			Handler:    _Promotion_DeletePromoProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "promotion.proto",
}
