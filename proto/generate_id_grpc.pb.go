// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: generate_id.proto

package proto

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
	GenerateId_GetIds_FullMethodName = "/GenerateId/getIds"
)

// GenerateIdClient is the client API for GenerateId service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GenerateIdClient interface {
	GetIds(ctx context.Context, in *GetIdRequest, opts ...grpc.CallOption) (*GetIdResponse, error)
}

type generateIdClient struct {
	cc grpc.ClientConnInterface
}

func NewGenerateIdClient(cc grpc.ClientConnInterface) GenerateIdClient {
	return &generateIdClient{cc}
}

func (c *generateIdClient) GetIds(ctx context.Context, in *GetIdRequest, opts ...grpc.CallOption) (*GetIdResponse, error) {
	out := new(GetIdResponse)
	err := c.cc.Invoke(ctx, GenerateId_GetIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenerateIdServer is the server API for GenerateId service.
// All implementations must embed UnimplementedGenerateIdServer
// for forward compatibility
type GenerateIdServer interface {
	GetIds(context.Context, *GetIdRequest) (*GetIdResponse, error)
	mustEmbedUnimplementedGenerateIdServer()
}

// UnimplementedGenerateIdServer must be embedded to have forward compatible implementations.
type UnimplementedGenerateIdServer struct {
}

func (UnimplementedGenerateIdServer) GetIds(context.Context, *GetIdRequest) (*GetIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIds not implemented")
}
func (UnimplementedGenerateIdServer) mustEmbedUnimplementedGenerateIdServer() {}

// UnsafeGenerateIdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenerateIdServer will
// result in compilation errors.
type UnsafeGenerateIdServer interface {
	mustEmbedUnimplementedGenerateIdServer()
}

func RegisterGenerateIdServer(s grpc.ServiceRegistrar, srv GenerateIdServer) {
	s.RegisterService(&GenerateId_ServiceDesc, srv)
}

func _GenerateId_GetIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenerateIdServer).GetIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenerateId_GetIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenerateIdServer).GetIds(ctx, req.(*GetIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GenerateId_ServiceDesc is the grpc.ServiceDesc for GenerateId service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GenerateId_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GenerateId",
	HandlerType: (*GenerateIdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getIds",
			Handler:    _GenerateId_GetIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "generate_id.proto",
}
