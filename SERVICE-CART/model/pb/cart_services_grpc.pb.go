// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: cart_services.proto

package pb

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

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartServiceClient interface {
	FindCartById(ctx context.Context, in *FindCartByCartIdRequest, opts ...grpc.CallOption) (*Cart, error)
	FindCartItems(ctx context.Context, in *FindCartItemsRequest, opts ...grpc.CallOption) (*FindCartItemsResponse, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) FindCartById(ctx context.Context, in *FindCartByCartIdRequest, opts ...grpc.CallOption) (*Cart, error) {
	out := new(Cart)
	err := c.cc.Invoke(ctx, "/model.CartService/FindCartById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) FindCartItems(ctx context.Context, in *FindCartItemsRequest, opts ...grpc.CallOption) (*FindCartItemsResponse, error) {
	out := new(FindCartItemsResponse)
	err := c.cc.Invoke(ctx, "/model.CartService/FindCartItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations must embed UnimplementedCartServiceServer
// for forward compatibility
type CartServiceServer interface {
	FindCartById(context.Context, *FindCartByCartIdRequest) (*Cart, error)
	FindCartItems(context.Context, *FindCartItemsRequest) (*FindCartItemsResponse, error)
	mustEmbedUnimplementedCartServiceServer()
}

// UnimplementedCartServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (UnimplementedCartServiceServer) FindCartById(context.Context, *FindCartByCartIdRequest) (*Cart, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCartById not implemented")
}
func (UnimplementedCartServiceServer) FindCartItems(context.Context, *FindCartItemsRequest) (*FindCartItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCartItems not implemented")
}
func (UnimplementedCartServiceServer) mustEmbedUnimplementedCartServiceServer() {}

// UnsafeCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServiceServer will
// result in compilation errors.
type UnsafeCartServiceServer interface {
	mustEmbedUnimplementedCartServiceServer()
}

func RegisterCartServiceServer(s grpc.ServiceRegistrar, srv CartServiceServer) {
	s.RegisterService(&CartService_ServiceDesc, srv)
}

func _CartService_FindCartById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCartByCartIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).FindCartById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.CartService/FindCartById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).FindCartById(ctx, req.(*FindCartByCartIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_FindCartItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCartItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).FindCartItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.CartService/FindCartItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).FindCartItems(ctx, req.(*FindCartItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CartService_ServiceDesc is the grpc.ServiceDesc for CartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindCartById",
			Handler:    _CartService_FindCartById_Handler,
		},
		{
			MethodName: "FindCartItems",
			Handler:    _CartService_FindCartItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart_services.proto",
}
