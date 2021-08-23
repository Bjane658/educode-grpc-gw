// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package educode_catalogue_api

import (
	context "context"
	protobuf "educode/protobuf"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CatalogueServiceClient is the client API for CatalogueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogueServiceClient interface {
	FindById(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*protobuf.Challenge, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (CatalogueService_ListClient, error)
	Put(ctx context.Context, in *protobuf.Challenge, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Whitelist(ctx context.Context, in *WhitelistEntry, opts ...grpc.CallOption) (*emptypb.Empty, error)
	IsWhitelisted(ctx context.Context, in *WhitelistEntry, opts ...grpc.CallOption) (*WhitelistEntry, error)
}

type catalogueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogueServiceClient(cc grpc.ClientConnInterface) CatalogueServiceClient {
	return &catalogueServiceClient{cc}
}

func (c *catalogueServiceClient) FindById(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*protobuf.Challenge, error) {
	out := new(protobuf.Challenge)
	err := c.cc.Invoke(ctx, "/educode.catalogue.api.CatalogueService/FindById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogueServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (CatalogueService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CatalogueService_serviceDesc.Streams[0], "/educode.catalogue.api.CatalogueService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &catalogueServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CatalogueService_ListClient interface {
	Recv() (*protobuf.Challenge, error)
	grpc.ClientStream
}

type catalogueServiceListClient struct {
	grpc.ClientStream
}

func (x *catalogueServiceListClient) Recv() (*protobuf.Challenge, error) {
	m := new(protobuf.Challenge)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *catalogueServiceClient) Put(ctx context.Context, in *protobuf.Challenge, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/educode.catalogue.api.CatalogueService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogueServiceClient) Whitelist(ctx context.Context, in *WhitelistEntry, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/educode.catalogue.api.CatalogueService/Whitelist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogueServiceClient) IsWhitelisted(ctx context.Context, in *WhitelistEntry, opts ...grpc.CallOption) (*WhitelistEntry, error) {
	out := new(WhitelistEntry)
	err := c.cc.Invoke(ctx, "/educode.catalogue.api.CatalogueService/IsWhitelisted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogueServiceServer is the server API for CatalogueService service.
// All implementations should embed UnimplementedCatalogueServiceServer
// for forward compatibility
type CatalogueServiceServer interface {
	FindById(context.Context, *FindRequest) (*protobuf.Challenge, error)
	List(*ListRequest, CatalogueService_ListServer) error
	Put(context.Context, *protobuf.Challenge) (*emptypb.Empty, error)
	Whitelist(context.Context, *WhitelistEntry) (*emptypb.Empty, error)
	IsWhitelisted(context.Context, *WhitelistEntry) (*WhitelistEntry, error)
}

// UnimplementedCatalogueServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCatalogueServiceServer struct {
}

func (UnimplementedCatalogueServiceServer) FindById(context.Context, *FindRequest) (*protobuf.Challenge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedCatalogueServiceServer) List(*ListRequest, CatalogueService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCatalogueServiceServer) Put(context.Context, *protobuf.Challenge) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedCatalogueServiceServer) Whitelist(context.Context, *WhitelistEntry) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Whitelist not implemented")
}
func (UnimplementedCatalogueServiceServer) IsWhitelisted(context.Context, *WhitelistEntry) (*WhitelistEntry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsWhitelisted not implemented")
}

// UnsafeCatalogueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogueServiceServer will
// result in compilation errors.
type UnsafeCatalogueServiceServer interface {
	mustEmbedUnimplementedCatalogueServiceServer()
}

func RegisterCatalogueServiceServer(s grpc.ServiceRegistrar, srv CatalogueServiceServer) {
	s.RegisterService(&_CatalogueService_serviceDesc, srv)
}

func _CatalogueService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogueServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/educode.catalogue.api.CatalogueService/FindById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogueServiceServer).FindById(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogueService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CatalogueServiceServer).List(m, &catalogueServiceListServer{stream})
}

type CatalogueService_ListServer interface {
	Send(*protobuf.Challenge) error
	grpc.ServerStream
}

type catalogueServiceListServer struct {
	grpc.ServerStream
}

func (x *catalogueServiceListServer) Send(m *protobuf.Challenge) error {
	return x.ServerStream.SendMsg(m)
}

func _CatalogueService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.Challenge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogueServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/educode.catalogue.api.CatalogueService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogueServiceServer).Put(ctx, req.(*protobuf.Challenge))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogueService_Whitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhitelistEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogueServiceServer).Whitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/educode.catalogue.api.CatalogueService/Whitelist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogueServiceServer).Whitelist(ctx, req.(*WhitelistEntry))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogueService_IsWhitelisted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhitelistEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogueServiceServer).IsWhitelisted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/educode.catalogue.api.CatalogueService/IsWhitelisted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogueServiceServer).IsWhitelisted(ctx, req.(*WhitelistEntry))
	}
	return interceptor(ctx, in, info, handler)
}

var _CatalogueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "educode.catalogue.api.CatalogueService",
	HandlerType: (*CatalogueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindById",
			Handler:    _CatalogueService_FindById_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _CatalogueService_Put_Handler,
		},
		{
			MethodName: "Whitelist",
			Handler:    _CatalogueService_Whitelist_Handler,
		},
		{
			MethodName: "IsWhitelisted",
			Handler:    _CatalogueService_IsWhitelisted_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _CatalogueService_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "educode/catalogue/api/v1.proto",
}
