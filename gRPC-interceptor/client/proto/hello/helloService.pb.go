// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: helloService.proto

package hellopb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ManRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ManRequest) Reset() {
	*x = ManRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManRequest) ProtoMessage() {}

func (x *ManRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManRequest.ProtoReflect.Descriptor instead.
func (*ManRequest) Descriptor() ([]byte, []int) {
	return file_helloService_proto_rawDescGZIP(), []int{0}
}

type ManResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ManResponse) Reset() {
	*x = ManResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManResponse) ProtoMessage() {}

func (x *ManResponse) ProtoReflect() protoreflect.Message {
	mi := &file_helloService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManResponse.ProtoReflect.Descriptor instead.
func (*ManResponse) Descriptor() ([]byte, []int) {
	return file_helloService_proto_rawDescGZIP(), []int{1}
}

func (x *ManResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CatRequest) Reset() {
	*x = CatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatRequest) ProtoMessage() {}

func (x *CatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatRequest.ProtoReflect.Descriptor instead.
func (*CatRequest) Descriptor() ([]byte, []int) {
	return file_helloService_proto_rawDescGZIP(), []int{2}
}

type CatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CatResponse) Reset() {
	*x = CatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatResponse) ProtoMessage() {}

func (x *CatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_helloService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatResponse.ProtoReflect.Descriptor instead.
func (*CatResponse) Descriptor() ([]byte, []int) {
	return file_helloService_proto_rawDescGZIP(), []int{3}
}

func (x *CatResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DogRequest) Reset() {
	*x = DogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DogRequest) ProtoMessage() {}

func (x *DogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DogRequest.ProtoReflect.Descriptor instead.
func (*DogRequest) Descriptor() ([]byte, []int) {
	return file_helloService_proto_rawDescGZIP(), []int{4}
}

type DogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DogResponse) Reset() {
	*x = DogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DogResponse) ProtoMessage() {}

func (x *DogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_helloService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DogResponse.ProtoReflect.Descriptor instead.
func (*DogResponse) Descriptor() ([]byte, []int) {
	return file_helloService_proto_rawDescGZIP(), []int{5}
}

func (x *DogResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_helloService_proto protoreflect.FileDescriptor

var file_helloService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x6c, 0x6c, 0x61, 0x72, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x4d, 0x61,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x27, 0x0a, 0x0b, 0x4d, 0x61, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x27, 0x0a, 0x0b, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x44, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x27, 0x0a, 0x0b, 0x44, 0x6f, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0xd2, 0x02, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x6a, 0x0a, 0x03, 0x4d, 0x61, 0x6e, 0x12, 0x30, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x6c, 0x6c, 0x61, 0x72,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x2e, 0x4d,
	0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x6c, 0x6c,
	0x61, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74,
	0x2e, 0x4d, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x03,
	0x43, 0x61, 0x74, 0x12, 0x30, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x6c, 0x6c, 0x61, 0x72, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x6c, 0x6c, 0x61, 0x72, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x2e, 0x43, 0x61, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x03, 0x44, 0x6f, 0x67, 0x12,
	0x30, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x6c, 0x6c, 0x61, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x2e, 0x44, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x31, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73,
	0x65, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x6c, 0x6c, 0x61, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x2e, 0x44, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloService_proto_rawDescOnce sync.Once
	file_helloService_proto_rawDescData = file_helloService_proto_rawDesc
)

func file_helloService_proto_rawDescGZIP() []byte {
	file_helloService_proto_rawDescOnce.Do(func() {
		file_helloService_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloService_proto_rawDescData)
	})
	return file_helloService_proto_rawDescData
}

var file_helloService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_helloService_proto_goTypes = []interface{}{
	(*ManRequest)(nil),  // 0: github.com.sevendollar.grpcintercept.ManRequest
	(*ManResponse)(nil), // 1: github.com.sevendollar.grpcintercept.ManResponse
	(*CatRequest)(nil),  // 2: github.com.sevendollar.grpcintercept.CatRequest
	(*CatResponse)(nil), // 3: github.com.sevendollar.grpcintercept.CatResponse
	(*DogRequest)(nil),  // 4: github.com.sevendollar.grpcintercept.DogRequest
	(*DogResponse)(nil), // 5: github.com.sevendollar.grpcintercept.DogResponse
}
var file_helloService_proto_depIdxs = []int32{
	0, // 0: github.com.sevendollar.grpcintercept.HelloService.Man:input_type -> github.com.sevendollar.grpcintercept.ManRequest
	2, // 1: github.com.sevendollar.grpcintercept.HelloService.Cat:input_type -> github.com.sevendollar.grpcintercept.CatRequest
	4, // 2: github.com.sevendollar.grpcintercept.HelloService.Dog:input_type -> github.com.sevendollar.grpcintercept.DogRequest
	1, // 3: github.com.sevendollar.grpcintercept.HelloService.Man:output_type -> github.com.sevendollar.grpcintercept.ManResponse
	3, // 4: github.com.sevendollar.grpcintercept.HelloService.Cat:output_type -> github.com.sevendollar.grpcintercept.CatResponse
	5, // 5: github.com.sevendollar.grpcintercept.HelloService.Dog:output_type -> github.com.sevendollar.grpcintercept.DogResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_helloService_proto_init() }
func file_helloService_proto_init() {
	if File_helloService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManRequest); i {
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
		file_helloService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManResponse); i {
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
		file_helloService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CatRequest); i {
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
		file_helloService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CatResponse); i {
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
		file_helloService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DogRequest); i {
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
		file_helloService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DogResponse); i {
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
			RawDescriptor: file_helloService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloService_proto_goTypes,
		DependencyIndexes: file_helloService_proto_depIdxs,
		MessageInfos:      file_helloService_proto_msgTypes,
	}.Build()
	File_helloService_proto = out.File
	file_helloService_proto_rawDesc = nil
	file_helloService_proto_goTypes = nil
	file_helloService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	Man(ctx context.Context, in *ManRequest, opts ...grpc.CallOption) (*ManResponse, error)
	Cat(ctx context.Context, in *CatRequest, opts ...grpc.CallOption) (*CatResponse, error)
	Dog(ctx context.Context, in *DogRequest, opts ...grpc.CallOption) (*DogResponse, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Man(ctx context.Context, in *ManRequest, opts ...grpc.CallOption) (*ManResponse, error) {
	out := new(ManResponse)
	err := c.cc.Invoke(ctx, "/github.com.sevendollar.grpcintercept.HelloService/Man", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) Cat(ctx context.Context, in *CatRequest, opts ...grpc.CallOption) (*CatResponse, error) {
	out := new(CatResponse)
	err := c.cc.Invoke(ctx, "/github.com.sevendollar.grpcintercept.HelloService/Cat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) Dog(ctx context.Context, in *DogRequest, opts ...grpc.CallOption) (*DogResponse, error) {
	out := new(DogResponse)
	err := c.cc.Invoke(ctx, "/github.com.sevendollar.grpcintercept.HelloService/Dog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	Man(context.Context, *ManRequest) (*ManResponse, error)
	Cat(context.Context, *CatRequest) (*CatResponse, error)
	Dog(context.Context, *DogRequest) (*DogResponse, error)
}

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (*UnimplementedHelloServiceServer) Man(context.Context, *ManRequest) (*ManResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Man not implemented")
}
func (*UnimplementedHelloServiceServer) Cat(context.Context, *CatRequest) (*CatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cat not implemented")
}
func (*UnimplementedHelloServiceServer) Dog(context.Context, *DogRequest) (*DogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dog not implemented")
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_Man_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Man(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.sevendollar.grpcintercept.HelloService/Man",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Man(ctx, req.(*ManRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_Cat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Cat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.sevendollar.grpcintercept.HelloService/Cat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Cat(ctx, req.(*CatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_Dog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Dog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.sevendollar.grpcintercept.HelloService/Dog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Dog(ctx, req.(*DogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.sevendollar.grpcintercept.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Man",
			Handler:    _HelloService_Man_Handler,
		},
		{
			MethodName: "Cat",
			Handler:    _HelloService_Cat_Handler,
		},
		{
			MethodName: "Dog",
			Handler:    _HelloService_Dog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloService.proto",
}
