// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// GoodsClient is the client API for Goods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsClient interface {
	CreateGoods(ctx context.Context, in *CreateGoodsRequest, opts ...grpc.CallOption) (*CreateGoodsReply, error)
	UpdateGoods(ctx context.Context, in *UpdateGoodsRequest, opts ...grpc.CallOption) (*UpdateGoodsReply, error)
	DeleteGoods(ctx context.Context, in *DeleteGoodsRequest, opts ...grpc.CallOption) (*DeleteGoodsReply, error)
	GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsReply, error)
	ListGoods(ctx context.Context, in *ListGoodsRequest, opts ...grpc.CallOption) (*ListGoodsReply, error)
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersReply, error)
	CreateOrders(ctx context.Context, in *CreateOrdersRequest, opts ...grpc.CallOption) (*CreateOrdersReply, error)
	// 分布式事务-tcc-try
	CreateOrdersTccTry(ctx context.Context, in *CreateOrdersRequest, opts ...grpc.CallOption) (*CreateOrdersReply, error)
	// 分布式事务-tcc-confirm
	CreateOrdersConfirm(ctx context.Context, in *CreateOrdersRequest, opts ...grpc.CallOption) (*CreateOrdersReply, error)
	// 分布式事务-tcc-cancel
	CreateOrdersTccCancel(ctx context.Context, in *CreateOrdersRequest, opts ...grpc.CallOption) (*CreateOrdersReply, error)
	UpdateOrders(ctx context.Context, in *UpdateOrdersRequest, opts ...grpc.CallOption) (*UpdateOrdersReply, error)
	DeleteOrders(ctx context.Context, in *DeleteOrdersRequest, opts ...grpc.CallOption) (*DeleteOrdersReply, error)
	ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersReply, error)
}

type goodsClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsClient(cc grpc.ClientConnInterface) GoodsClient {
	return &goodsClient{cc}
}

func (c *goodsClient) CreateGoods(ctx context.Context, in *CreateGoodsRequest, opts ...grpc.CallOption) (*CreateGoodsReply, error) {
	out := new(CreateGoodsReply)
	err := c.cc.Invoke(ctx, "/api.goods.service.v1.Goods/CreateGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateGoods(ctx context.Context, in *UpdateGoodsRequest, opts ...grpc.CallOption) (*UpdateGoodsReply, error) {
	out := new(UpdateGoodsReply)
	err := c.cc.Invoke(ctx, "/api.goods.service.v1.Goods/UpdateGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DeleteGoods(ctx context.Context, in *DeleteGoodsRequest, opts ...grpc.CallOption) (*DeleteGoodsReply, error) {
	out := new(DeleteGoodsReply)
	err := c.cc.Invoke(ctx, "/api.goods.service.v1.Goods/DeleteGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsReply, error) {
	out := new(GetGoodsReply)
	err := c.cc.Invoke(ctx, "/api.goods.service.v1.Goods/GetGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) ListGoods(ctx context.Context, in *ListGoodsRequest, opts ...grpc.CallOption) (*ListGoodsReply, error) {
	out := new(ListGoodsReply)
	err := c.cc.Invoke(ctx, "/api.goods.service.v1.Goods/ListGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersReply, error) {
	out := new(GetOrdersReply)
	err := c.cc.Invoke(ctx, "/api.goods.service.v1.Goods/GetOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateOrders(ctx context.Context, in *CreateOrdersRequest, opts ...grpc.CallOption) (*CreateOrdersReply, error) {
	out := new(CreateOrdersReply)
	err := c.cc.Invoke(ctx, "/api.goods.service.v1.Goods/CreateOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateOrdersTccTry(ctx context.Context, in *CreateOrdersRequest, opts ...grpc.CallOption) (*CreateOrdersReply, error) {
	out := new(CreateOrdersReply)
	err := c.cc.Invoke(ctx, "/api.goods.service.v1.Goods/CreateOrdersTccTry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateOrdersConfirm(ctx context.Context, in *CreateOrdersRequest, opts ...grpc.CallOption) (*CreateOrdersReply, error) {
	out := new(CreateOrdersReply)
	err := c.cc.Invoke(ctx, "/api.goods.service.v1.Goods/CreateOrdersConfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateOrdersTccCancel(ctx context.Context, in *CreateOrdersRequest, opts ...grpc.CallOption) (*CreateOrdersReply, error) {
	out := new(CreateOrdersReply)
	err := c.cc.Invoke(ctx, "/api.goods.service.v1.Goods/CreateOrdersTccCancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateOrders(ctx context.Context, in *UpdateOrdersRequest, opts ...grpc.CallOption) (*UpdateOrdersReply, error) {
	out := new(UpdateOrdersReply)
	err := c.cc.Invoke(ctx, "/api.goods.service.v1.Goods/UpdateOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DeleteOrders(ctx context.Context, in *DeleteOrdersRequest, opts ...grpc.CallOption) (*DeleteOrdersReply, error) {
	out := new(DeleteOrdersReply)
	err := c.cc.Invoke(ctx, "/api.goods.service.v1.Goods/DeleteOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersReply, error) {
	out := new(ListOrdersReply)
	err := c.cc.Invoke(ctx, "/api.goods.service.v1.Goods/ListOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServer is the server API for Goods service.
// All implementations must embed UnimplementedGoodsServer
// for forward compatibility
type GoodsServer interface {
	CreateGoods(context.Context, *CreateGoodsRequest) (*CreateGoodsReply, error)
	UpdateGoods(context.Context, *UpdateGoodsRequest) (*UpdateGoodsReply, error)
	DeleteGoods(context.Context, *DeleteGoodsRequest) (*DeleteGoodsReply, error)
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsReply, error)
	ListGoods(context.Context, *ListGoodsRequest) (*ListGoodsReply, error)
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersReply, error)
	CreateOrders(context.Context, *CreateOrdersRequest) (*CreateOrdersReply, error)
	// 分布式事务-tcc-try
	CreateOrdersTccTry(context.Context, *CreateOrdersRequest) (*CreateOrdersReply, error)
	// 分布式事务-tcc-confirm
	CreateOrdersConfirm(context.Context, *CreateOrdersRequest) (*CreateOrdersReply, error)
	// 分布式事务-tcc-cancel
	CreateOrdersTccCancel(context.Context, *CreateOrdersRequest) (*CreateOrdersReply, error)
	UpdateOrders(context.Context, *UpdateOrdersRequest) (*UpdateOrdersReply, error)
	DeleteOrders(context.Context, *DeleteOrdersRequest) (*DeleteOrdersReply, error)
	ListOrders(context.Context, *ListOrdersRequest) (*ListOrdersReply, error)
	mustEmbedUnimplementedGoodsServer()
}

// UnimplementedGoodsServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsServer struct {
}

func (UnimplementedGoodsServer) CreateGoods(context.Context, *CreateGoodsRequest) (*CreateGoodsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGoods not implemented")
}
func (UnimplementedGoodsServer) UpdateGoods(context.Context, *UpdateGoodsRequest) (*UpdateGoodsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGoods not implemented")
}
func (UnimplementedGoodsServer) DeleteGoods(context.Context, *DeleteGoodsRequest) (*DeleteGoodsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGoods not implemented")
}
func (UnimplementedGoodsServer) GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoods not implemented")
}
func (UnimplementedGoodsServer) ListGoods(context.Context, *ListGoodsRequest) (*ListGoodsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGoods not implemented")
}
func (UnimplementedGoodsServer) GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedGoodsServer) CreateOrders(context.Context, *CreateOrdersRequest) (*CreateOrdersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrders not implemented")
}
func (UnimplementedGoodsServer) CreateOrdersTccTry(context.Context, *CreateOrdersRequest) (*CreateOrdersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrdersTccTry not implemented")
}
func (UnimplementedGoodsServer) CreateOrdersConfirm(context.Context, *CreateOrdersRequest) (*CreateOrdersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrdersConfirm not implemented")
}
func (UnimplementedGoodsServer) CreateOrdersTccCancel(context.Context, *CreateOrdersRequest) (*CreateOrdersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrdersTccCancel not implemented")
}
func (UnimplementedGoodsServer) UpdateOrders(context.Context, *UpdateOrdersRequest) (*UpdateOrdersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrders not implemented")
}
func (UnimplementedGoodsServer) DeleteOrders(context.Context, *DeleteOrdersRequest) (*DeleteOrdersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrders not implemented")
}
func (UnimplementedGoodsServer) ListOrders(context.Context, *ListOrdersRequest) (*ListOrdersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (UnimplementedGoodsServer) mustEmbedUnimplementedGoodsServer() {}

// UnsafeGoodsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServer will
// result in compilation errors.
type UnsafeGoodsServer interface {
	mustEmbedUnimplementedGoodsServer()
}

func RegisterGoodsServer(s grpc.ServiceRegistrar, srv GoodsServer) {
	s.RegisterService(&Goods_ServiceDesc, srv)
}

func _Goods_CreateGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.goods.service.v1.Goods/CreateGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateGoods(ctx, req.(*CreateGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.goods.service.v1.Goods/UpdateGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateGoods(ctx, req.(*UpdateGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DeleteGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DeleteGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.goods.service.v1.Goods/DeleteGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DeleteGoods(ctx, req.(*DeleteGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.goods.service.v1.Goods/GetGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGoods(ctx, req.(*GetGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_ListGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).ListGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.goods.service.v1.Goods/ListGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).ListGoods(ctx, req.(*ListGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.goods.service.v1.Goods/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.goods.service.v1.Goods/CreateOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateOrders(ctx, req.(*CreateOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateOrdersTccTry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateOrdersTccTry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.goods.service.v1.Goods/CreateOrdersTccTry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateOrdersTccTry(ctx, req.(*CreateOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateOrdersConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateOrdersConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.goods.service.v1.Goods/CreateOrdersConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateOrdersConfirm(ctx, req.(*CreateOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateOrdersTccCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateOrdersTccCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.goods.service.v1.Goods/CreateOrdersTccCancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateOrdersTccCancel(ctx, req.(*CreateOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.goods.service.v1.Goods/UpdateOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateOrders(ctx, req.(*UpdateOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DeleteOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DeleteOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.goods.service.v1.Goods/DeleteOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DeleteOrders(ctx, req.(*DeleteOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.goods.service.v1.Goods/ListOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).ListOrders(ctx, req.(*ListOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Goods_ServiceDesc is the grpc.ServiceDesc for Goods service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Goods_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.goods.service.v1.Goods",
	HandlerType: (*GoodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGoods",
			Handler:    _Goods_CreateGoods_Handler,
		},
		{
			MethodName: "UpdateGoods",
			Handler:    _Goods_UpdateGoods_Handler,
		},
		{
			MethodName: "DeleteGoods",
			Handler:    _Goods_DeleteGoods_Handler,
		},
		{
			MethodName: "GetGoods",
			Handler:    _Goods_GetGoods_Handler,
		},
		{
			MethodName: "ListGoods",
			Handler:    _Goods_ListGoods_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _Goods_GetOrders_Handler,
		},
		{
			MethodName: "CreateOrders",
			Handler:    _Goods_CreateOrders_Handler,
		},
		{
			MethodName: "CreateOrdersTccTry",
			Handler:    _Goods_CreateOrdersTccTry_Handler,
		},
		{
			MethodName: "CreateOrdersConfirm",
			Handler:    _Goods_CreateOrdersConfirm_Handler,
		},
		{
			MethodName: "CreateOrdersTccCancel",
			Handler:    _Goods_CreateOrdersTccCancel_Handler,
		},
		{
			MethodName: "UpdateOrders",
			Handler:    _Goods_UpdateOrders_Handler,
		},
		{
			MethodName: "DeleteOrders",
			Handler:    _Goods_DeleteOrders_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _Goods_ListOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/goods.proto",
}
