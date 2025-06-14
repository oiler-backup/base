// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: proto/metrics.proto

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BackupMetricsService_ReportSuccessfulBackup_FullMethodName = "/backupmetrics.BackupMetricsService/ReportSuccessfulBackup"
	BackupMetricsService_ReportRestoreStatus_FullMethodName    = "/backupmetrics.BackupMetricsService/ReportRestoreStatus"
)

// BackupMetricsServiceClient is the client API for BackupMetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackupMetricsServiceClient interface {
	ReportSuccessfulBackup(ctx context.Context, in *BackupMetrics, opts ...grpc.CallOption) (*empty.Empty, error)
	ReportRestoreStatus(ctx context.Context, in *RestoreMetrics, opts ...grpc.CallOption) (*empty.Empty, error)
}

type backupMetricsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBackupMetricsServiceClient(cc grpc.ClientConnInterface) BackupMetricsServiceClient {
	return &backupMetricsServiceClient{cc}
}

func (c *backupMetricsServiceClient) ReportSuccessfulBackup(ctx context.Context, in *BackupMetrics, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, BackupMetricsService_ReportSuccessfulBackup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupMetricsServiceClient) ReportRestoreStatus(ctx context.Context, in *RestoreMetrics, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, BackupMetricsService_ReportRestoreStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupMetricsServiceServer is the server API for BackupMetricsService service.
// All implementations must embed UnimplementedBackupMetricsServiceServer
// for forward compatibility.
type BackupMetricsServiceServer interface {
	ReportSuccessfulBackup(context.Context, *BackupMetrics) (*empty.Empty, error)
	ReportRestoreStatus(context.Context, *RestoreMetrics) (*empty.Empty, error)
	mustEmbedUnimplementedBackupMetricsServiceServer()
}

// UnimplementedBackupMetricsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBackupMetricsServiceServer struct{}

func (UnimplementedBackupMetricsServiceServer) ReportSuccessfulBackup(context.Context, *BackupMetrics) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportSuccessfulBackup not implemented")
}
func (UnimplementedBackupMetricsServiceServer) ReportRestoreStatus(context.Context, *RestoreMetrics) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportRestoreStatus not implemented")
}
func (UnimplementedBackupMetricsServiceServer) mustEmbedUnimplementedBackupMetricsServiceServer() {}
func (UnimplementedBackupMetricsServiceServer) testEmbeddedByValue()                              {}

// UnsafeBackupMetricsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackupMetricsServiceServer will
// result in compilation errors.
type UnsafeBackupMetricsServiceServer interface {
	mustEmbedUnimplementedBackupMetricsServiceServer()
}

func RegisterBackupMetricsServiceServer(s grpc.ServiceRegistrar, srv BackupMetricsServiceServer) {
	// If the following call pancis, it indicates UnimplementedBackupMetricsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BackupMetricsService_ServiceDesc, srv)
}

func _BackupMetricsService_ReportSuccessfulBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupMetrics)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupMetricsServiceServer).ReportSuccessfulBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupMetricsService_ReportSuccessfulBackup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupMetricsServiceServer).ReportSuccessfulBackup(ctx, req.(*BackupMetrics))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupMetricsService_ReportRestoreStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreMetrics)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupMetricsServiceServer).ReportRestoreStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupMetricsService_ReportRestoreStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupMetricsServiceServer).ReportRestoreStatus(ctx, req.(*RestoreMetrics))
	}
	return interceptor(ctx, in, info, handler)
}

// BackupMetricsService_ServiceDesc is the grpc.ServiceDesc for BackupMetricsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BackupMetricsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "backupmetrics.BackupMetricsService",
	HandlerType: (*BackupMetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportSuccessfulBackup",
			Handler:    _BackupMetricsService_ReportSuccessfulBackup_Handler,
		},
		{
			MethodName: "ReportRestoreStatus",
			Handler:    _BackupMetricsService_ReportRestoreStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/metrics.proto",
}
