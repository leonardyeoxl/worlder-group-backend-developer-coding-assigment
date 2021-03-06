// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.12.2
// source: data/data.proto

package data

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

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensorValue int64  `protobuf:"varint,1,opt,name=SensorValue,proto3" json:"SensorValue,omitempty"`
	ID1         int64  `protobuf:"varint,2,opt,name=ID1,proto3" json:"ID1,omitempty"`
	ID2         string `protobuf:"bytes,3,opt,name=ID2,proto3" json:"ID2,omitempty"`
	Timestamp   int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_data_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetSensorValue() int64 {
	if x != nil {
		return x.SensorValue
	}
	return 0
}

func (x *Data) GetID1() int64 {
	if x != nil {
		return x.ID1
	}
	return 0
}

func (x *Data) GetID2() string {
	if x != nil {
		return x.ID2
	}
	return ""
}

func (x *Data) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Collection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Data `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Collection) Reset() {
	*x = Collection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Collection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collection) ProtoMessage() {}

func (x *Collection) ProtoReflect() protoreflect.Message {
	mi := &file_data_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collection.ProtoReflect.Descriptor instead.
func (*Collection) Descriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{1}
}

func (x *Collection) GetData() []*Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID1            int64  `protobuf:"varint,1,opt,name=ID1,proto3" json:"ID1,omitempty"`
	ID2            string `protobuf:"bytes,2,opt,name=ID2,proto3" json:"ID2,omitempty"`
	StartTimestamp int64  `protobuf:"varint,3,opt,name=StartTimestamp,proto3" json:"StartTimestamp,omitempty"`
	EndTimestamp   int64  `protobuf:"varint,4,opt,name=EndTimestamp,proto3" json:"EndTimestamp,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_data_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetID1() int64 {
	if x != nil {
		return x.ID1
	}
	return 0
}

func (x *Request) GetID2() string {
	if x != nil {
		return x.ID2
	}
	return ""
}

func (x *Request) GetStartTimestamp() int64 {
	if x != nil {
		return x.StartTimestamp
	}
	return 0
}

func (x *Request) GetEndTimestamp() int64 {
	if x != nil {
		return x.EndTimestamp
	}
	return 0
}

var File_data_data_proto protoreflect.FileDescriptor

var file_data_data_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6a, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x20, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x44, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x49, 0x44, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x44, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x49, 0x44, 0x32, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x2c, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x79, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x49, 0x44, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x49, 0x44, 0x31, 0x12, 0x10,
	0x0a, 0x03, 0x49, 0x44, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x49, 0x44, 0x32,
	0x12, 0x26, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x48, 0x0a, 0x13,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x0c, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x10, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_data_proto_rawDescOnce sync.Once
	file_data_data_proto_rawDescData = file_data_data_proto_rawDesc
)

func file_data_data_proto_rawDescGZIP() []byte {
	file_data_data_proto_rawDescOnce.Do(func() {
		file_data_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_data_proto_rawDescData)
	})
	return file_data_data_proto_rawDescData
}

var file_data_data_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_data_data_proto_goTypes = []interface{}{
	(*Data)(nil),       // 0: data.Data
	(*Collection)(nil), // 1: data.Collection
	(*Request)(nil),    // 2: data.Request
}
var file_data_data_proto_depIdxs = []int32{
	0, // 0: data.Collection.data:type_name -> data.Data
	2, // 1: data.RetrieveDataService.RetrieveData:input_type -> data.Request
	1, // 2: data.RetrieveDataService.RetrieveData:output_type -> data.Collection
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_data_data_proto_init() }
func file_data_data_proto_init() {
	if File_data_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_data_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Collection); i {
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
		file_data_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
			RawDescriptor: file_data_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_data_data_proto_goTypes,
		DependencyIndexes: file_data_data_proto_depIdxs,
		MessageInfos:      file_data_data_proto_msgTypes,
	}.Build()
	File_data_data_proto = out.File
	file_data_data_proto_rawDesc = nil
	file_data_data_proto_goTypes = nil
	file_data_data_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RetrieveDataServiceClient is the client API for RetrieveDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RetrieveDataServiceClient interface {
	RetrieveData(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Collection, error)
}

type retrieveDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRetrieveDataServiceClient(cc grpc.ClientConnInterface) RetrieveDataServiceClient {
	return &retrieveDataServiceClient{cc}
}

func (c *retrieveDataServiceClient) RetrieveData(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Collection, error) {
	out := new(Collection)
	err := c.cc.Invoke(ctx, "/data.RetrieveDataService/RetrieveData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RetrieveDataServiceServer is the server API for RetrieveDataService service.
type RetrieveDataServiceServer interface {
	RetrieveData(context.Context, *Request) (*Collection, error)
}

// UnimplementedRetrieveDataServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRetrieveDataServiceServer struct {
}

func (*UnimplementedRetrieveDataServiceServer) RetrieveData(context.Context, *Request) (*Collection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveData not implemented")
}

func RegisterRetrieveDataServiceServer(s *grpc.Server, srv RetrieveDataServiceServer) {
	s.RegisterService(&_RetrieveDataService_serviceDesc, srv)
}

func _RetrieveDataService_RetrieveData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrieveDataServiceServer).RetrieveData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.RetrieveDataService/RetrieveData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrieveDataServiceServer).RetrieveData(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _RetrieveDataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "data.RetrieveDataService",
	HandlerType: (*RetrieveDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrieveData",
			Handler:    _RetrieveDataService_RetrieveData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data/data.proto",
}
