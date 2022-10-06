// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: photo/v1/photo.proto

package photov1

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

// PhotoServiceClient is the client API for PhotoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PhotoServiceClient interface {
	// GetPhotos returns all the photos
	GetPhotos(ctx context.Context, in *GetPhotosRequest, opts ...grpc.CallOption) (*GetPhotosResponse, error)
	// GetByHash returns a photo by its hash
	GetByHash(ctx context.Context, in *GetByHashRequest, opts ...grpc.CallOption) (*GetByHashResponse, error)
	// ContentByHash returns the photo content by its hash
	ContentByHash(ctx context.Context, in *ContentByHashRequest, opts ...grpc.CallOption) (PhotoService_ContentByHashClient, error)
}

type photoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPhotoServiceClient(cc grpc.ClientConnInterface) PhotoServiceClient {
	return &photoServiceClient{cc}
}

func (c *photoServiceClient) GetPhotos(ctx context.Context, in *GetPhotosRequest, opts ...grpc.CallOption) (*GetPhotosResponse, error) {
	out := new(GetPhotosResponse)
	err := c.cc.Invoke(ctx, "/photo.v1.PhotoService/GetPhotos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoServiceClient) GetByHash(ctx context.Context, in *GetByHashRequest, opts ...grpc.CallOption) (*GetByHashResponse, error) {
	out := new(GetByHashResponse)
	err := c.cc.Invoke(ctx, "/photo.v1.PhotoService/GetByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoServiceClient) ContentByHash(ctx context.Context, in *ContentByHashRequest, opts ...grpc.CallOption) (PhotoService_ContentByHashClient, error) {
	stream, err := c.cc.NewStream(ctx, &PhotoService_ServiceDesc.Streams[0], "/photo.v1.PhotoService/ContentByHash", opts...)
	if err != nil {
		return nil, err
	}
	x := &photoServiceContentByHashClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PhotoService_ContentByHashClient interface {
	Recv() (*PhotoServiceContentByHashResponse, error)
	grpc.ClientStream
}

type photoServiceContentByHashClient struct {
	grpc.ClientStream
}

func (x *photoServiceContentByHashClient) Recv() (*PhotoServiceContentByHashResponse, error) {
	m := new(PhotoServiceContentByHashResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PhotoServiceServer is the server API for PhotoService service.
// All implementations should embed UnimplementedPhotoServiceServer
// for forward compatibility
type PhotoServiceServer interface {
	// GetPhotos returns all the photos
	GetPhotos(context.Context, *GetPhotosRequest) (*GetPhotosResponse, error)
	// GetByHash returns a photo by its hash
	GetByHash(context.Context, *GetByHashRequest) (*GetByHashResponse, error)
	// ContentByHash returns the photo content by its hash
	ContentByHash(*ContentByHashRequest, PhotoService_ContentByHashServer) error
}

// UnimplementedPhotoServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPhotoServiceServer struct {
}

func (UnimplementedPhotoServiceServer) GetPhotos(context.Context, *GetPhotosRequest) (*GetPhotosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPhotos not implemented")
}
func (UnimplementedPhotoServiceServer) GetByHash(context.Context, *GetByHashRequest) (*GetByHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByHash not implemented")
}
func (UnimplementedPhotoServiceServer) ContentByHash(*ContentByHashRequest, PhotoService_ContentByHashServer) error {
	return status.Errorf(codes.Unimplemented, "method ContentByHash not implemented")
}

// UnsafePhotoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PhotoServiceServer will
// result in compilation errors.
type UnsafePhotoServiceServer interface {
	mustEmbedUnimplementedPhotoServiceServer()
}

func RegisterPhotoServiceServer(s grpc.ServiceRegistrar, srv PhotoServiceServer) {
	s.RegisterService(&PhotoService_ServiceDesc, srv)
}

func _PhotoService_GetPhotos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPhotosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServiceServer).GetPhotos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/photo.v1.PhotoService/GetPhotos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServiceServer).GetPhotos(ctx, req.(*GetPhotosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoService_GetByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServiceServer).GetByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/photo.v1.PhotoService/GetByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServiceServer).GetByHash(ctx, req.(*GetByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoService_ContentByHash_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ContentByHashRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PhotoServiceServer).ContentByHash(m, &photoServiceContentByHashServer{stream})
}

type PhotoService_ContentByHashServer interface {
	Send(*PhotoServiceContentByHashResponse) error
	grpc.ServerStream
}

type photoServiceContentByHashServer struct {
	grpc.ServerStream
}

func (x *photoServiceContentByHashServer) Send(m *PhotoServiceContentByHashResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PhotoService_ServiceDesc is the grpc.ServiceDesc for PhotoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PhotoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "photo.v1.PhotoService",
	HandlerType: (*PhotoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPhotos",
			Handler:    _PhotoService_GetPhotos_Handler,
		},
		{
			MethodName: "GetByHash",
			Handler:    _PhotoService_GetByHash_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ContentByHash",
			Handler:       _PhotoService_ContentByHash_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "photo/v1/photo.proto",
}
