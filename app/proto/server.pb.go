// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetDataRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDataRequest) Reset()         { *m = GetDataRequest{} }
func (m *GetDataRequest) String() string { return proto.CompactTextString(m) }
func (*GetDataRequest) ProtoMessage()    {}
func (*GetDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *GetDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataRequest.Unmarshal(m, b)
}
func (m *GetDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataRequest.Marshal(b, m, deterministic)
}
func (m *GetDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataRequest.Merge(m, src)
}
func (m *GetDataRequest) XXX_Size() int {
	return xxx_messageInfo_GetDataRequest.Size(m)
}
func (m *GetDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataRequest proto.InternalMessageInfo

func (m *GetDataRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetDataResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDataResponse) Reset()         { *m = GetDataResponse{} }
func (m *GetDataResponse) String() string { return proto.CompactTextString(m) }
func (*GetDataResponse) ProtoMessage()    {}
func (*GetDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *GetDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataResponse.Unmarshal(m, b)
}
func (m *GetDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataResponse.Marshal(b, m, deterministic)
}
func (m *GetDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataResponse.Merge(m, src)
}
func (m *GetDataResponse) XXX_Size() int {
	return xxx_messageInfo_GetDataResponse.Size(m)
}
func (m *GetDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataResponse proto.InternalMessageInfo

func (m *GetDataResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetDataRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDataRequest) Reset()         { *m = SetDataRequest{} }
func (m *SetDataRequest) String() string { return proto.CompactTextString(m) }
func (*SetDataRequest) ProtoMessage()    {}
func (*SetDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{2}
}

func (m *SetDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDataRequest.Unmarshal(m, b)
}
func (m *SetDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDataRequest.Marshal(b, m, deterministic)
}
func (m *SetDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDataRequest.Merge(m, src)
}
func (m *SetDataRequest) XXX_Size() int {
	return xxx_messageInfo_SetDataRequest.Size(m)
}
func (m *SetDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDataRequest proto.InternalMessageInfo

func (m *SetDataRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetDataRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ListDataResponse struct {
	Data                 []*Data  `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDataResponse) Reset()         { *m = ListDataResponse{} }
func (m *ListDataResponse) String() string { return proto.CompactTextString(m) }
func (*ListDataResponse) ProtoMessage()    {}
func (*ListDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{3}
}

func (m *ListDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDataResponse.Unmarshal(m, b)
}
func (m *ListDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDataResponse.Marshal(b, m, deterministic)
}
func (m *ListDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDataResponse.Merge(m, src)
}
func (m *ListDataResponse) XXX_Size() int {
	return xxx_messageInfo_ListDataResponse.Size(m)
}
func (m *ListDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDataResponse proto.InternalMessageInfo

func (m *ListDataResponse) GetData() []*Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type Data struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{4}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Data) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type DeleteDataRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDataRequest) Reset()         { *m = DeleteDataRequest{} }
func (m *DeleteDataRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDataRequest) ProtoMessage()    {}
func (*DeleteDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{5}
}

func (m *DeleteDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDataRequest.Unmarshal(m, b)
}
func (m *DeleteDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDataRequest.Marshal(b, m, deterministic)
}
func (m *DeleteDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDataRequest.Merge(m, src)
}
func (m *DeleteDataRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDataRequest.Size(m)
}
func (m *DeleteDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDataRequest proto.InternalMessageInfo

func (m *DeleteDataRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDataRequest)(nil), "proto.GetDataRequest")
	proto.RegisterType((*GetDataResponse)(nil), "proto.GetDataResponse")
	proto.RegisterType((*SetDataRequest)(nil), "proto.SetDataRequest")
	proto.RegisterType((*ListDataResponse)(nil), "proto.ListDataResponse")
	proto.RegisterType((*Data)(nil), "proto.Data")
	proto.RegisterType((*DeleteDataRequest)(nil), "proto.DeleteDataRequest")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0xd2, 0xe9, 0xf9,
	0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x60, 0x5e, 0x52, 0x69, 0x9a, 0x7e, 0x6a, 0x6e, 0x41, 0x49, 0x25,
	0x44, 0x8d, 0x92, 0x12, 0x17, 0x9f, 0x7b, 0x6a, 0x89, 0x4b, 0x62, 0x49, 0x62, 0x50, 0x6a, 0x61,
	0x69, 0x6a, 0x71, 0x89, 0x90, 0x00, 0x17, 0x73, 0x76, 0x6a, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x67, 0x10, 0x88, 0xa9, 0xa4, 0xce, 0xc5, 0x0f, 0x57, 0x53, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a,
	0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x0a, 0x55, 0x06, 0xe1, 0x28, 0x59, 0x70, 0xf1,
	0x05, 0x13, 0x30, 0x0c, 0xa1, 0x93, 0x09, 0x59, 0xa7, 0x31, 0x97, 0x80, 0x4f, 0x66, 0x31, 0xaa,
	0x1d, 0xf2, 0x5c, 0x2c, 0x29, 0x89, 0x25, 0x89, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xdc,
	0x10, 0x07, 0xeb, 0x81, 0x95, 0x80, 0x25, 0x94, 0xf4, 0xb8, 0x58, 0x40, 0x3c, 0xa2, 0x2d, 0x51,
	0xe5, 0x12, 0x74, 0x49, 0xcd, 0x49, 0x2d, 0x49, 0xc5, 0xeb, 0x42, 0xa3, 0x46, 0x26, 0x2e, 0xb6,
	0x60, 0x70, 0x38, 0x0a, 0x59, 0x71, 0xb1, 0x43, 0x7d, 0x2e, 0x24, 0x0a, 0xb5, 0x1f, 0x35, 0xb4,
	0xa4, 0xc4, 0xd0, 0x85, 0x21, 0x8e, 0x57, 0x62, 0x00, 0xe9, 0x0d, 0x46, 0xd3, 0x1b, 0x8c, 0xae,
	0x17, 0x12, 0x33, 0x7a, 0xb0, 0x98, 0xd1, 0x73, 0x05, 0xc5, 0x8c, 0x12, 0x83, 0x90, 0x2d, 0x17,
	0x07, 0x2c, 0x38, 0x84, 0x70, 0xa8, 0x92, 0x12, 0x87, 0x1a, 0x8a, 0x1e, 0x6e, 0x4a, 0x0c, 0x42,
	0x0e, 0x5c, 0x5c, 0x08, 0x8f, 0x0a, 0x49, 0xc0, 0x42, 0x0e, 0xdd, 0xef, 0xb8, 0x1d, 0x90, 0xc4,
	0x06, 0x16, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xab, 0x11, 0x73, 0x06, 0x51, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerClient interface {
	GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error)
	SetData(ctx context.Context, in *SetDataRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ListData(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListDataResponse, error)
	DeleteData(ctx context.Context, in *DeleteDataRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type serverClient struct {
	cc *grpc.ClientConn
}

func NewServerClient(cc *grpc.ClientConn) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error) {
	out := new(GetDataResponse)
	err := c.cc.Invoke(ctx, "/proto.Server/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) SetData(ctx context.Context, in *SetDataRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Server/SetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) ListData(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListDataResponse, error) {
	out := new(ListDataResponse)
	err := c.cc.Invoke(ctx, "/proto.Server/ListData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) DeleteData(ctx context.Context, in *DeleteDataRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Server/DeleteData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
type ServerServer interface {
	GetData(context.Context, *GetDataRequest) (*GetDataResponse, error)
	SetData(context.Context, *SetDataRequest) (*empty.Empty, error)
	ListData(context.Context, *empty.Empty) (*ListDataResponse, error)
	DeleteData(context.Context, *DeleteDataRequest) (*empty.Empty, error)
}

// UnimplementedServerServer can be embedded to have forward compatible implementations.
type UnimplementedServerServer struct {
}

func (*UnimplementedServerServer) GetData(ctx context.Context, req *GetDataRequest) (*GetDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (*UnimplementedServerServer) SetData(ctx context.Context, req *SetDataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetData not implemented")
}
func (*UnimplementedServerServer) ListData(ctx context.Context, req *empty.Empty) (*ListDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListData not implemented")
}
func (*UnimplementedServerServer) DeleteData(ctx context.Context, req *DeleteDataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteData not implemented")
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Server/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).GetData(ctx, req.(*GetDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_SetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).SetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Server/SetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).SetData(ctx, req.(*SetDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_ListData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).ListData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Server/ListData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).ListData(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_DeleteData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).DeleteData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Server/DeleteData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).DeleteData(ctx, req.(*DeleteDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _Server_GetData_Handler,
		},
		{
			MethodName: "SetData",
			Handler:    _Server_SetData_Handler,
		},
		{
			MethodName: "ListData",
			Handler:    _Server_ListData_Handler,
		},
		{
			MethodName: "DeleteData",
			Handler:    _Server_DeleteData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
