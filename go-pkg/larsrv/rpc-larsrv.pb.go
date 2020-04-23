// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc-larsrv.proto

package larsrv

import (
	context "context"
	fmt "fmt"
	laruche "github.com/benka-me/laruche/go-pkg/laruche"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("rpc-larsrv.proto", fileDescriptor_40fa15abad12b977) }

var fileDescriptor_40fa15abad12b977 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x2a, 0x48, 0xd6,
	0xcd, 0x49, 0x2c, 0x2a, 0x2e, 0x2a, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0,
	0xa4, 0x78, 0x90, 0x45, 0xa5, 0x8c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73,
	0xf5, 0x93, 0x52, 0xf3, 0xb2, 0x13, 0x75, 0x73, 0x53, 0xf5, 0x73, 0x12, 0x8b, 0x4a, 0x93, 0x33,
	0x52, 0xf5, 0xc1, 0x4a, 0x92, 0x4a, 0xd3, 0x60, 0x02, 0x10, 0x3d, 0x46, 0x0b, 0x98, 0xb9, 0xd8,
	0x7c, 0xc0, 0x86, 0x08, 0x19, 0x73, 0x71, 0x05, 0x94, 0x26, 0xe5, 0x64, 0x16, 0x67, 0x38, 0xa5,
	0xa6, 0x0a, 0x09, 0xe8, 0xc1, 0x14, 0x06, 0x94, 0x82, 0x45, 0xa4, 0x84, 0xd1, 0x45, 0x82, 0x52,
	0x8b, 0x85, 0xcc, 0xb8, 0xb8, 0xa1, 0x9a, 0x3c, 0x32, 0xcb, 0x52, 0x85, 0x04, 0x51, 0xd4, 0x80,
	0x84, 0xa4, 0x44, 0x30, 0x84, 0x40, 0xfa, 0x4c, 0xb9, 0x78, 0x02, 0x8a, 0x32, 0xcb, 0x12, 0x4b,
	0x32, 0xab, 0x52, 0x49, 0xb0, 0xce, 0x82, 0x8b, 0x17, 0xae, 0x8d, 0x34, 0x0b, 0xf5, 0xb9, 0xd8,
	0x82, 0x53, 0x4b, 0x48, 0xb0, 0xca, 0x88, 0x8b, 0x3d, 0x38, 0xb5, 0x84, 0x34, 0x4b, 0xd4, 0xb9,
	0xd8, 0xdc, 0x21, 0x96, 0xf0, 0xc3, 0xe5, 0xc1, 0xc6, 0x15, 0x4a, 0xf1, 0x20, 0x0b, 0x08, 0x69,
	0x71, 0xb1, 0x43, 0x14, 0x16, 0x23, 0x39, 0x07, 0xc4, 0x05, 0x29, 0xe5, 0x45, 0x11, 0x71, 0x4a,
	0xbb, 0xf0, 0x50, 0x8e, 0xe1, 0xc6, 0x43, 0x39, 0x86, 0x0f, 0x0f, 0xe5, 0x18, 0x1b, 0x1e, 0xc9,
	0x31, 0xae, 0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e,
	0xc9, 0x31, 0xbe, 0x78, 0x24, 0xc7, 0xf0, 0xe1, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17,
	0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x80, 0x27, 0x31, 0xe8, 0x16, 0xa7, 0x16,
	0x95, 0xa5, 0x16, 0xe9, 0xa7, 0xe7, 0xeb, 0x16, 0x64, 0xa7, 0xeb, 0x43, 0x12, 0x51, 0x12, 0x1b,
	0x38, 0x45, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xc4, 0x3f, 0x78, 0x6f, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LarsrvClient is the client API for Larsrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LarsrvClient interface {
	PublishBee(ctx context.Context, in *laruche.PushBee, opts ...grpc.CallOption) (*laruche.PushBeeRes, error)
	PublishHive(ctx context.Context, in *laruche.PushHive, opts ...grpc.CallOption) (*laruche.PushHiveRes, error)
	PrivatizeBee(ctx context.Context, in *laruche.PushBee, opts ...grpc.CallOption) (*laruche.PushBeeRes, error)
	PrivatizeHive(ctx context.Context, in *laruche.PushHive, opts ...grpc.CallOption) (*laruche.PushHiveRes, error)
	SetBee(ctx context.Context, in *laruche.PushBee, opts ...grpc.CallOption) (*laruche.PushBeeRes, error)
	SetHive(ctx context.Context, in *laruche.PushHive, opts ...grpc.CallOption) (*laruche.PushHiveRes, error)
	GetBee(ctx context.Context, in *laruche.BeeReq, opts ...grpc.CallOption) (*laruche.Bee, error)
	GetBees(ctx context.Context, in *laruche.BeesReq, opts ...grpc.CallOption) (*laruche.Bees, error)
}

type larsrvClient struct {
	cc *grpc.ClientConn
}

func NewLarsrvClient(cc *grpc.ClientConn) LarsrvClient {
	return &larsrvClient{cc}
}

func (c *larsrvClient) PublishBee(ctx context.Context, in *laruche.PushBee, opts ...grpc.CallOption) (*laruche.PushBeeRes, error) {
	out := new(laruche.PushBeeRes)
	err := c.cc.Invoke(ctx, "/larsrv.Larsrv/PublishBee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *larsrvClient) PublishHive(ctx context.Context, in *laruche.PushHive, opts ...grpc.CallOption) (*laruche.PushHiveRes, error) {
	out := new(laruche.PushHiveRes)
	err := c.cc.Invoke(ctx, "/larsrv.Larsrv/PublishHive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *larsrvClient) PrivatizeBee(ctx context.Context, in *laruche.PushBee, opts ...grpc.CallOption) (*laruche.PushBeeRes, error) {
	out := new(laruche.PushBeeRes)
	err := c.cc.Invoke(ctx, "/larsrv.Larsrv/PrivatizeBee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *larsrvClient) PrivatizeHive(ctx context.Context, in *laruche.PushHive, opts ...grpc.CallOption) (*laruche.PushHiveRes, error) {
	out := new(laruche.PushHiveRes)
	err := c.cc.Invoke(ctx, "/larsrv.Larsrv/PrivatizeHive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *larsrvClient) SetBee(ctx context.Context, in *laruche.PushBee, opts ...grpc.CallOption) (*laruche.PushBeeRes, error) {
	out := new(laruche.PushBeeRes)
	err := c.cc.Invoke(ctx, "/larsrv.Larsrv/SetBee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *larsrvClient) SetHive(ctx context.Context, in *laruche.PushHive, opts ...grpc.CallOption) (*laruche.PushHiveRes, error) {
	out := new(laruche.PushHiveRes)
	err := c.cc.Invoke(ctx, "/larsrv.Larsrv/SetHive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *larsrvClient) GetBee(ctx context.Context, in *laruche.BeeReq, opts ...grpc.CallOption) (*laruche.Bee, error) {
	out := new(laruche.Bee)
	err := c.cc.Invoke(ctx, "/larsrv.Larsrv/GetBee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *larsrvClient) GetBees(ctx context.Context, in *laruche.BeesReq, opts ...grpc.CallOption) (*laruche.Bees, error) {
	out := new(laruche.Bees)
	err := c.cc.Invoke(ctx, "/larsrv.Larsrv/GetBees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LarsrvServer is the server API for Larsrv service.
type LarsrvServer interface {
	PublishBee(context.Context, *laruche.PushBee) (*laruche.PushBeeRes, error)
	PublishHive(context.Context, *laruche.PushHive) (*laruche.PushHiveRes, error)
	PrivatizeBee(context.Context, *laruche.PushBee) (*laruche.PushBeeRes, error)
	PrivatizeHive(context.Context, *laruche.PushHive) (*laruche.PushHiveRes, error)
	SetBee(context.Context, *laruche.PushBee) (*laruche.PushBeeRes, error)
	SetHive(context.Context, *laruche.PushHive) (*laruche.PushHiveRes, error)
	GetBee(context.Context, *laruche.BeeReq) (*laruche.Bee, error)
	GetBees(context.Context, *laruche.BeesReq) (*laruche.Bees, error)
}

// UnimplementedLarsrvServer can be embedded to have forward compatible implementations.
type UnimplementedLarsrvServer struct {
}

func (*UnimplementedLarsrvServer) PublishBee(ctx context.Context, req *laruche.PushBee) (*laruche.PushBeeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishBee not implemented")
}
func (*UnimplementedLarsrvServer) PublishHive(ctx context.Context, req *laruche.PushHive) (*laruche.PushHiveRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishHive not implemented")
}
func (*UnimplementedLarsrvServer) PrivatizeBee(ctx context.Context, req *laruche.PushBee) (*laruche.PushBeeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrivatizeBee not implemented")
}
func (*UnimplementedLarsrvServer) PrivatizeHive(ctx context.Context, req *laruche.PushHive) (*laruche.PushHiveRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrivatizeHive not implemented")
}
func (*UnimplementedLarsrvServer) SetBee(ctx context.Context, req *laruche.PushBee) (*laruche.PushBeeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBee not implemented")
}
func (*UnimplementedLarsrvServer) SetHive(ctx context.Context, req *laruche.PushHive) (*laruche.PushHiveRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHive not implemented")
}
func (*UnimplementedLarsrvServer) GetBee(ctx context.Context, req *laruche.BeeReq) (*laruche.Bee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBee not implemented")
}
func (*UnimplementedLarsrvServer) GetBees(ctx context.Context, req *laruche.BeesReq) (*laruche.Bees, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBees not implemented")
}

func RegisterLarsrvServer(s *grpc.Server, srv LarsrvServer) {
	s.RegisterService(&_Larsrv_serviceDesc, srv)
}

func _Larsrv_PublishBee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(laruche.PushBee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LarsrvServer).PublishBee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larsrv.Larsrv/PublishBee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LarsrvServer).PublishBee(ctx, req.(*laruche.PushBee))
	}
	return interceptor(ctx, in, info, handler)
}

func _Larsrv_PublishHive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(laruche.PushHive)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LarsrvServer).PublishHive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larsrv.Larsrv/PublishHive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LarsrvServer).PublishHive(ctx, req.(*laruche.PushHive))
	}
	return interceptor(ctx, in, info, handler)
}

func _Larsrv_PrivatizeBee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(laruche.PushBee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LarsrvServer).PrivatizeBee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larsrv.Larsrv/PrivatizeBee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LarsrvServer).PrivatizeBee(ctx, req.(*laruche.PushBee))
	}
	return interceptor(ctx, in, info, handler)
}

func _Larsrv_PrivatizeHive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(laruche.PushHive)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LarsrvServer).PrivatizeHive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larsrv.Larsrv/PrivatizeHive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LarsrvServer).PrivatizeHive(ctx, req.(*laruche.PushHive))
	}
	return interceptor(ctx, in, info, handler)
}

func _Larsrv_SetBee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(laruche.PushBee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LarsrvServer).SetBee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larsrv.Larsrv/SetBee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LarsrvServer).SetBee(ctx, req.(*laruche.PushBee))
	}
	return interceptor(ctx, in, info, handler)
}

func _Larsrv_SetHive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(laruche.PushHive)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LarsrvServer).SetHive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larsrv.Larsrv/SetHive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LarsrvServer).SetHive(ctx, req.(*laruche.PushHive))
	}
	return interceptor(ctx, in, info, handler)
}

func _Larsrv_GetBee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(laruche.BeeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LarsrvServer).GetBee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larsrv.Larsrv/GetBee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LarsrvServer).GetBee(ctx, req.(*laruche.BeeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Larsrv_GetBees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(laruche.BeesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LarsrvServer).GetBees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larsrv.Larsrv/GetBees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LarsrvServer).GetBees(ctx, req.(*laruche.BeesReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Larsrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "larsrv.Larsrv",
	HandlerType: (*LarsrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishBee",
			Handler:    _Larsrv_PublishBee_Handler,
		},
		{
			MethodName: "PublishHive",
			Handler:    _Larsrv_PublishHive_Handler,
		},
		{
			MethodName: "PrivatizeBee",
			Handler:    _Larsrv_PrivatizeBee_Handler,
		},
		{
			MethodName: "PrivatizeHive",
			Handler:    _Larsrv_PrivatizeHive_Handler,
		},
		{
			MethodName: "SetBee",
			Handler:    _Larsrv_SetBee_Handler,
		},
		{
			MethodName: "SetHive",
			Handler:    _Larsrv_SetHive_Handler,
		},
		{
			MethodName: "GetBee",
			Handler:    _Larsrv_GetBee_Handler,
		},
		{
			MethodName: "GetBees",
			Handler:    _Larsrv_GetBees_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc-larsrv.proto",
}