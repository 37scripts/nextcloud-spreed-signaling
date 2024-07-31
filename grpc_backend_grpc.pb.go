//*
// Standalone signaling server for the Nextcloud Spreed app.
// Copyright (C) 2022 struktur AG
//
// @author Joachim Bauch <bauch@struktur.de>
//
// @license GNU AGPL version 3 or any later version
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// source: grpc_backend.proto

package signaling

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RpcBackend_GetSessionCount_FullMethodName = "/signaling.RpcBackend/GetSessionCount"
)

// RpcBackendClient is the client API for RpcBackend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcBackendClient interface {
	GetSessionCount(ctx context.Context, in *GetSessionCountRequest, opts ...grpc.CallOption) (*GetSessionCountReply, error)
}

type rpcBackendClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcBackendClient(cc grpc.ClientConnInterface) RpcBackendClient {
	return &rpcBackendClient{cc}
}

func (c *rpcBackendClient) GetSessionCount(ctx context.Context, in *GetSessionCountRequest, opts ...grpc.CallOption) (*GetSessionCountReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSessionCountReply)
	err := c.cc.Invoke(ctx, RpcBackend_GetSessionCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcBackendServer is the server API for RpcBackend service.
// All implementations must embed UnimplementedRpcBackendServer
// for forward compatibility.
type RpcBackendServer interface {
	GetSessionCount(context.Context, *GetSessionCountRequest) (*GetSessionCountReply, error)
	mustEmbedUnimplementedRpcBackendServer()
}

// UnimplementedRpcBackendServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRpcBackendServer struct{}

func (UnimplementedRpcBackendServer) GetSessionCount(context.Context, *GetSessionCountRequest) (*GetSessionCountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessionCount not implemented")
}
func (UnimplementedRpcBackendServer) mustEmbedUnimplementedRpcBackendServer() {}
func (UnimplementedRpcBackendServer) testEmbeddedByValue()                    {}

// UnsafeRpcBackendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcBackendServer will
// result in compilation errors.
type UnsafeRpcBackendServer interface {
	mustEmbedUnimplementedRpcBackendServer()
}

func RegisterRpcBackendServer(s grpc.ServiceRegistrar, srv RpcBackendServer) {
	// If the following call pancis, it indicates UnimplementedRpcBackendServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RpcBackend_ServiceDesc, srv)
}

func _RpcBackend_GetSessionCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcBackendServer).GetSessionCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcBackend_GetSessionCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcBackendServer).GetSessionCount(ctx, req.(*GetSessionCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcBackend_ServiceDesc is the grpc.ServiceDesc for RpcBackend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcBackend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "signaling.RpcBackend",
	HandlerType: (*RpcBackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSessionCount",
			Handler:    _RpcBackend_GetSessionCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_backend.proto",
}
