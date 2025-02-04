// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: Records.proto

package health_sync

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
	MedicalRecordService_Create_FullMethodName = "/health_sync.MedicalRecordService/Create"
	MedicalRecordService_Update_FullMethodName = "/health_sync.MedicalRecordService/Update"
	MedicalRecordService_Delete_FullMethodName = "/health_sync.MedicalRecordService/Delete"
	MedicalRecordService_Get_FullMethodName    = "/health_sync.MedicalRecordService/Get"
	MedicalRecordService_List_FullMethodName   = "/health_sync.MedicalRecordService/List"
)

// MedicalRecordServiceClient is the client API for MedicalRecordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MedicalRecordServiceClient interface {
	Create(ctx context.Context, in *MedicalRecordCreate, opts ...grpc.CallOption) (*Void, error)
	Update(ctx context.Context, in *MedicalRecordUpdate, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*Void, error)
	Get(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*MedicalRecordRes, error)
	List(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*GetAllMedicalRecordsRes, error)
}

type medicalRecordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMedicalRecordServiceClient(cc grpc.ClientConnInterface) MedicalRecordServiceClient {
	return &medicalRecordServiceClient{cc}
}

func (c *medicalRecordServiceClient) Create(ctx context.Context, in *MedicalRecordCreate, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, MedicalRecordService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalRecordServiceClient) Update(ctx context.Context, in *MedicalRecordUpdate, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, MedicalRecordService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalRecordServiceClient) Delete(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, MedicalRecordService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalRecordServiceClient) Get(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*MedicalRecordRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MedicalRecordRes)
	err := c.cc.Invoke(ctx, MedicalRecordService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalRecordServiceClient) List(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*GetAllMedicalRecordsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllMedicalRecordsRes)
	err := c.cc.Invoke(ctx, MedicalRecordService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MedicalRecordServiceServer is the server API for MedicalRecordService service.
// All implementations must embed UnimplementedMedicalRecordServiceServer
// for forward compatibility.
type MedicalRecordServiceServer interface {
	Create(context.Context, *MedicalRecordCreate) (*Void, error)
	Update(context.Context, *MedicalRecordUpdate) (*Void, error)
	Delete(context.Context, *GetById) (*Void, error)
	Get(context.Context, *GetById) (*MedicalRecordRes, error)
	List(context.Context, *Filter) (*GetAllMedicalRecordsRes, error)
	mustEmbedUnimplementedMedicalRecordServiceServer()
}

// UnimplementedMedicalRecordServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMedicalRecordServiceServer struct{}

func (UnimplementedMedicalRecordServiceServer) Create(context.Context, *MedicalRecordCreate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMedicalRecordServiceServer) Update(context.Context, *MedicalRecordUpdate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMedicalRecordServiceServer) Delete(context.Context, *GetById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMedicalRecordServiceServer) Get(context.Context, *GetById) (*MedicalRecordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMedicalRecordServiceServer) List(context.Context, *Filter) (*GetAllMedicalRecordsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMedicalRecordServiceServer) mustEmbedUnimplementedMedicalRecordServiceServer() {}
func (UnimplementedMedicalRecordServiceServer) testEmbeddedByValue()                              {}

// UnsafeMedicalRecordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MedicalRecordServiceServer will
// result in compilation errors.
type UnsafeMedicalRecordServiceServer interface {
	mustEmbedUnimplementedMedicalRecordServiceServer()
}

func RegisterMedicalRecordServiceServer(s grpc.ServiceRegistrar, srv MedicalRecordServiceServer) {
	// If the following call pancis, it indicates UnimplementedMedicalRecordServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MedicalRecordService_ServiceDesc, srv)
}

func _MedicalRecordService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedicalRecordCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalRecordServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedicalRecordService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalRecordServiceServer).Create(ctx, req.(*MedicalRecordCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalRecordService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedicalRecordUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalRecordServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedicalRecordService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalRecordServiceServer).Update(ctx, req.(*MedicalRecordUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalRecordService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalRecordServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedicalRecordService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalRecordServiceServer).Delete(ctx, req.(*GetById))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalRecordService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalRecordServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedicalRecordService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalRecordServiceServer).Get(ctx, req.(*GetById))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalRecordService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalRecordServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedicalRecordService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalRecordServiceServer).List(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

// MedicalRecordService_ServiceDesc is the grpc.ServiceDesc for MedicalRecordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MedicalRecordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "health_sync.MedicalRecordService",
	HandlerType: (*MedicalRecordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _MedicalRecordService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _MedicalRecordService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _MedicalRecordService_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _MedicalRecordService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _MedicalRecordService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Records.proto",
}
