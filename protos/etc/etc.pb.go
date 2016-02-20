// Code generated by protoc-gen-go.
// source: etc.proto
// DO NOT EDIT!

/*
Package etc is a generated protocol buffer package.

It is generated from these files:
	etc.proto

It has these top-level messages:
	GetReq
	GetReply
*/
package etc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type GetReq struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetReq) Reset()                    { *m = GetReq{} }
func (m *GetReq) String() string            { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()               {}
func (*GetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetReply struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *GetReply) Reset()                    { *m = GetReply{} }
func (m *GetReply) String() string            { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()               {}
func (*GetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*GetReq)(nil), "etc.GetReq")
	proto.RegisterType((*GetReply)(nil), "etc.GetReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Etc service

type EtcClient interface {
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetReply, error)
}

type etcClient struct {
	cc *grpc.ClientConn
}

func NewEtcClient(cc *grpc.ClientConn) EtcClient {
	return &etcClient{cc}
}

func (c *etcClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := grpc.Invoke(ctx, "/etc.Etc/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Etc service

type EtcServer interface {
	Get(context.Context, *GetReq) (*GetReply, error)
}

func RegisterEtcServer(s *grpc.Server, srv EtcServer) {
	s.RegisterService(&_Etc_serviceDesc, srv)
}

func _Etc_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(EtcServer).Get(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Etc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "etc.Etc",
	HandlerType: (*EtcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Etc_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2d, 0x49, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x06, 0x32, 0x95, 0xa4, 0xb8, 0xd8, 0xdc, 0x53, 0x4b,
	0x82, 0x52, 0x0b, 0x85, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0x40, 0x4c, 0x25, 0x05, 0x2e, 0x0e, 0xb0, 0x5c, 0x41, 0x4e, 0xa5, 0x90, 0x08, 0x17, 0x6b,
	0x59, 0x62, 0x4e, 0x69, 0x2a, 0x54, 0x1e, 0xc2, 0x31, 0xd2, 0xe2, 0x62, 0x76, 0x2d, 0x49, 0x16,
	0x52, 0xe6, 0x62, 0x06, 0x2a, 0x14, 0xe2, 0xd6, 0x03, 0x19, 0x0e, 0x31, 0x4e, 0x8a, 0x17, 0xc1,
	0x01, 0xea, 0x57, 0x62, 0x48, 0x62, 0x03, 0xdb, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x7a,
	0xf9, 0x49, 0x26, 0x82, 0x00, 0x00, 0x00,
}