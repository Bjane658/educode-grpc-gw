// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package educode_submission_api

import (
	context "context"
	protobuf "educode/protobuf"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SubmissionServiceClient is the client API for SubmissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubmissionServiceClient interface {
	Get(ctx context.Context, in *SubmissionGetRequest, opts ...grpc.CallOption) (SubmissionService_GetClient, error)
	Submit(ctx context.Context, in *SubmissionRequest, opts ...grpc.CallOption) (*SubmissionResponse, error)
}

type submissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubmissionServiceClient(cc grpc.ClientConnInterface) SubmissionServiceClient {
	return &submissionServiceClient{cc}
}

func (c *submissionServiceClient) Get(ctx context.Context, in *SubmissionGetRequest, opts ...grpc.CallOption) (SubmissionService_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SubmissionService_serviceDesc.Streams[0], "/educode.submission.api.SubmissionService/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &submissionServiceGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SubmissionService_GetClient interface {
	Recv() (*protobuf.Submission, error)
	grpc.ClientStream
}

type submissionServiceGetClient struct {
	grpc.ClientStream
}

func (x *submissionServiceGetClient) Recv() (*protobuf.Submission, error) {
	m := new(protobuf.Submission)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *submissionServiceClient) Submit(ctx context.Context, in *SubmissionRequest, opts ...grpc.CallOption) (*SubmissionResponse, error) {
	out := new(SubmissionResponse)
	err := c.cc.Invoke(ctx, "/educode.submission.api.SubmissionService/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubmissionServiceServer is the server API for SubmissionService service.
// All implementations should embed UnimplementedSubmissionServiceServer
// for forward compatibility
type SubmissionServiceServer interface {
	Get(*SubmissionGetRequest, SubmissionService_GetServer) error
	Submit(context.Context, *SubmissionRequest) (*SubmissionResponse, error)
}

// UnimplementedSubmissionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSubmissionServiceServer struct {
}

func (UnimplementedSubmissionServiceServer) Get(*SubmissionGetRequest, SubmissionService_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSubmissionServiceServer) Submit(context.Context, *SubmissionRequest) (*SubmissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}

// UnsafeSubmissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubmissionServiceServer will
// result in compilation errors.
type UnsafeSubmissionServiceServer interface {
	mustEmbedUnimplementedSubmissionServiceServer()
}

func RegisterSubmissionServiceServer(s grpc.ServiceRegistrar, srv SubmissionServiceServer) {
	s.RegisterService(&_SubmissionService_serviceDesc, srv)
}

func _SubmissionService_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubmissionGetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SubmissionServiceServer).Get(m, &submissionServiceGetServer{stream})
}

type SubmissionService_GetServer interface {
	Send(*protobuf.Submission) error
	grpc.ServerStream
}

type submissionServiceGetServer struct {
	grpc.ServerStream
}

func (x *submissionServiceGetServer) Send(m *protobuf.Submission) error {
	return x.ServerStream.SendMsg(m)
}

func _SubmissionService_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServiceServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/educode.submission.api.SubmissionService/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServiceServer).Submit(ctx, req.(*SubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SubmissionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "educode.submission.api.SubmissionService",
	HandlerType: (*SubmissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _SubmissionService_Submit_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _SubmissionService_Get_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "educode/submission/api/v1.proto",
}
