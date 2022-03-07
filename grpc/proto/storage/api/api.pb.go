// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: api.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StorageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StorageRequest) Reset() {
	*x = StorageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageRequest) ProtoMessage() {}

func (x *StorageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageRequest.ProtoReflect.Descriptor instead.
func (*StorageRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *StorageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StorageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StorageReply) Reset() {
	*x = StorageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageReply) ProtoMessage() {}

func (x *StorageReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageReply.ProtoReflect.Descriptor instead.
func (*StorageReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *StorageReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x22, 0x24, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0xef, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x12, 0x38, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0c, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x35, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x0d, 0x5a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_goTypes = []interface{}{
	(*StorageRequest)(nil), // 0: api.StorageRequest
	(*StorageReply)(nil),   // 1: api.StorageReply
}
var file_api_proto_depIdxs = []int32{
	0, // 0: api.LocalStorage.CreateVolume:input_type -> api.StorageRequest
	0, // 1: api.LocalStorage.RemoveVolume:input_type -> api.StorageRequest
	0, // 2: api.LocalStorage.ResizeVolume:input_type -> api.StorageRequest
	0, // 3: api.LocalStorage.ListVolumes:input_type -> api.StorageRequest
	1, // 4: api.LocalStorage.CreateVolume:output_type -> api.StorageReply
	1, // 5: api.LocalStorage.RemoveVolume:output_type -> api.StorageReply
	1, // 6: api.LocalStorage.ResizeVolume:output_type -> api.StorageReply
	1, // 7: api.LocalStorage.ListVolumes:output_type -> api.StorageReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LocalStorageClient is the client API for LocalStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LocalStorageClient interface {
	CreateVolume(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*StorageReply, error)
	RemoveVolume(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*StorageReply, error)
	ResizeVolume(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*StorageReply, error)
	ListVolumes(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*StorageReply, error)
}

type localStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewLocalStorageClient(cc grpc.ClientConnInterface) LocalStorageClient {
	return &localStorageClient{cc}
}

func (c *localStorageClient) CreateVolume(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*StorageReply, error) {
	out := new(StorageReply)
	err := c.cc.Invoke(ctx, "/api.LocalStorage/CreateVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localStorageClient) RemoveVolume(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*StorageReply, error) {
	out := new(StorageReply)
	err := c.cc.Invoke(ctx, "/api.LocalStorage/RemoveVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localStorageClient) ResizeVolume(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*StorageReply, error) {
	out := new(StorageReply)
	err := c.cc.Invoke(ctx, "/api.LocalStorage/ResizeVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localStorageClient) ListVolumes(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*StorageReply, error) {
	out := new(StorageReply)
	err := c.cc.Invoke(ctx, "/api.LocalStorage/ListVolumes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocalStorageServer is the server API for LocalStorage service.
type LocalStorageServer interface {
	CreateVolume(context.Context, *StorageRequest) (*StorageReply, error)
	RemoveVolume(context.Context, *StorageRequest) (*StorageReply, error)
	ResizeVolume(context.Context, *StorageRequest) (*StorageReply, error)
	ListVolumes(context.Context, *StorageRequest) (*StorageReply, error)
}

// UnimplementedLocalStorageServer can be embedded to have forward compatible implementations.
type UnimplementedLocalStorageServer struct {
}

func (*UnimplementedLocalStorageServer) CreateVolume(context.Context, *StorageRequest) (*StorageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVolume not implemented")
}
func (*UnimplementedLocalStorageServer) RemoveVolume(context.Context, *StorageRequest) (*StorageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveVolume not implemented")
}
func (*UnimplementedLocalStorageServer) ResizeVolume(context.Context, *StorageRequest) (*StorageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResizeVolume not implemented")
}
func (*UnimplementedLocalStorageServer) ListVolumes(context.Context, *StorageRequest) (*StorageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVolumes not implemented")
}

func RegisterLocalStorageServer(s *grpc.Server, srv LocalStorageServer) {
	s.RegisterService(&_LocalStorage_serviceDesc, srv)
}

func _LocalStorage_CreateVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalStorageServer).CreateVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.LocalStorage/CreateVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalStorageServer).CreateVolume(ctx, req.(*StorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocalStorage_RemoveVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalStorageServer).RemoveVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.LocalStorage/RemoveVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalStorageServer).RemoveVolume(ctx, req.(*StorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocalStorage_ResizeVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalStorageServer).ResizeVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.LocalStorage/ResizeVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalStorageServer).ResizeVolume(ctx, req.(*StorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocalStorage_ListVolumes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalStorageServer).ListVolumes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.LocalStorage/ListVolumes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalStorageServer).ListVolumes(ctx, req.(*StorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LocalStorage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.LocalStorage",
	HandlerType: (*LocalStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVolume",
			Handler:    _LocalStorage_CreateVolume_Handler,
		},
		{
			MethodName: "RemoveVolume",
			Handler:    _LocalStorage_RemoveVolume_Handler,
		},
		{
			MethodName: "ResizeVolume",
			Handler:    _LocalStorage_ResizeVolume_Handler,
		},
		{
			MethodName: "ListVolumes",
			Handler:    _LocalStorage_ListVolumes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
