// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.9.1
// source: customer.proto

package main

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type Customers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*Customer `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *Customers) Reset() {
	*x = Customers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customers) ProtoMessage() {}

func (x *Customers) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customers.ProtoReflect.Descriptor instead.
func (*Customers) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{1}
}

func (x *Customers) GetRecords() []*Customer {
	if x != nil {
		return x.Records
	}
	return nil
}

var File_customer_proto protoreflect.FileDescriptor

var file_customer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x35, 0x0a,
	0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x07, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x32, 0x6f, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x0e,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x0e,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x00,
	0x12, 0x33, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customer_proto_rawDescOnce sync.Once
	file_customer_proto_rawDescData = file_customer_proto_rawDesc
)

func file_customer_proto_rawDescGZIP() []byte {
	file_customer_proto_rawDescOnce.Do(func() {
		file_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_customer_proto_rawDescData)
	})
	return file_customer_proto_rawDescData
}

var file_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customer_proto_goTypes = []interface{}{
	(*Customer)(nil),    // 0: data.customer
	(*Customers)(nil),   // 1: data.customers
	(*empty.Empty)(nil), // 2: google.protobuf.Empty
}
var file_customer_proto_depIdxs = []int32{
	0, // 0: data.customers.records:type_name -> data.customer
	0, // 1: data.CustomerData.Insert:input_type -> data.customer
	2, // 2: data.CustomerData.GetAll:input_type -> google.protobuf.Empty
	0, // 3: data.CustomerData.Insert:output_type -> data.customer
	1, // 4: data.CustomerData.GetAll:output_type -> data.customers
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_customer_proto_init() }
func file_customer_proto_init() {
	if File_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customers); i {
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
			RawDescriptor: file_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customer_proto_goTypes,
		DependencyIndexes: file_customer_proto_depIdxs,
		MessageInfos:      file_customer_proto_msgTypes,
	}.Build()
	File_customer_proto = out.File
	file_customer_proto_rawDesc = nil
	file_customer_proto_goTypes = nil
	file_customer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerDataClient is the client API for CustomerData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerDataClient interface {
	Insert(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error)
	GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Customers, error)
}

type customerDataClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerDataClient(cc grpc.ClientConnInterface) CustomerDataClient {
	return &customerDataClient{cc}
}

func (c *customerDataClient) Insert(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/data.CustomerData/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerDataClient) GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Customers, error) {
	out := new(Customers)
	err := c.cc.Invoke(ctx, "/data.CustomerData/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerDataServer is the server API for CustomerData service.
type CustomerDataServer interface {
	Insert(context.Context, *Customer) (*Customer, error)
	GetAll(context.Context, *empty.Empty) (*Customers, error)
}

// UnimplementedCustomerDataServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerDataServer struct {
}

func (*UnimplementedCustomerDataServer) Insert(context.Context, *Customer) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (*UnimplementedCustomerDataServer) GetAll(context.Context, *empty.Empty) (*Customers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

func RegisterCustomerDataServer(s *grpc.Server, srv CustomerDataServer) {
	s.RegisterService(&_CustomerData_serviceDesc, srv)
}

func _CustomerData_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerDataServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.CustomerData/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerDataServer).Insert(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerData_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerDataServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.CustomerData/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerDataServer).GetAll(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerData_serviceDesc = grpc.ServiceDesc{
	ServiceName: "data.CustomerData",
	HandlerType: (*CustomerDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _CustomerData_Insert_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _CustomerData_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer.proto",
}
