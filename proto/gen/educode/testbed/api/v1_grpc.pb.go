// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package educode_testbed_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TestbedServiceClient is the client API for TestbedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestbedServiceClient interface {
	Execute(ctx context.Context, in *ExecutionRequest, opts ...grpc.CallOption) (*ExecutionResponse, error)
}

type testbedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestbedServiceClient(cc grpc.ClientConnInterface) TestbedServiceClient {
	return &testbedServiceClient{cc}
}

func (c *testbedServiceClient) Execute(ctx context.Context, in *ExecutionRequest, opts ...grpc.CallOption) (*ExecutionResponse, error) {
	out := new(ExecutionResponse)
	err := c.cc.Invoke(ctx, "/educode.testbed.api.TestbedService/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestbedServiceServer is the server API for TestbedService service.
// All implementations should embed UnimplementedTestbedServiceServer
// for forward compatibility
type TestbedServiceServer interface {
	Execute(context.Context, *ExecutionRequest) (*ExecutionResponse, error)
}

// UnimplementedTestbedServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTestbedServiceServer struct {
}

func (UnimplementedTestbedServiceServer) Execute(context.Context, *ExecutionRequest) (*ExecutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

// UnsafeTestbedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestbedServiceServer will
// result in compilation errors.
type UnsafeTestbedServiceServer interface {
	mustEmbedUnimplementedTestbedServiceServer()
}

func RegisterTestbedServiceServer(s grpc.ServiceRegistrar, srv TestbedServiceServer) {
	s.RegisterService(&_TestbedService_serviceDesc, srv)
}

func _TestbedService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestbedServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/educode.testbed.api.TestbedService/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestbedServiceServer).Execute(ctx, req.(*ExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestbedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "educode.testbed.api.TestbedService",
	HandlerType: (*TestbedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _TestbedService_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "educode/testbed/api/v1.proto",
}
