// Code generated by ICEBERG protoc-gen-go. DO NOT EDIT EXCEPET SERVER VERSION.
// source: hello.proto

/*
Package hello is a generated protocol buffer package.

It is generated from these files:
	hello.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package hello

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	"context"
	"github.com/kwins/iceberg/frame"
	"github.com/kwins/iceberg/frame/config"
	"github.com/kwins/iceberg/frame/protocol"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// HelloRequest 请求结构
type HelloRequest struct {
	Name string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty" xml:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// HelloResponse 响应结构
type HelloResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty" xml:"message,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "hello.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "hello.HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.

// Client API for Hello service
// iceberg server version,relation to server uri.
var hello_version = frame.SrvVersionName[frame.SV1]

// SayHello 定义SayHello方法
func SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	task, err := frame.ReadyTask(ctx, "SayHello", "hello", in)
	if err != nil {
		return nil, err
	}
	if span := frame.SpanWithTask(ctx, task); span != nil {
		defer span.Finish()
	}
	back, err := frame.DeliverTo(task)
	if err != nil {
		return nil, err
	}

	var out HelloResponse
	if err := protocol.Unpack(task.GetFormat(), back.GetBody(), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// HelloServer Server API for Hello service
type HelloServer interface {
	SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error)
}

// RegisterHelloServer register HelloServer with etcd info
func RegisterHelloServer(srv HelloServer, cfg *config.BaseCfg) {
	frame.RegisterAndServe(&helloServerDesc, srv, cfg)
}

// hello server SayHello handler
func helloSayHelloHandler(srv interface{}, ctx context.Context, format protocol.RestfulFormat, data []byte) ([]byte, error) {
	var in HelloRequest
	if err := protocol.Unpack(format, data, &in); err != nil {
		return nil, err
	}

	helloResp, err := srv.(HelloServer).SayHello(ctx, &in)
	if err != nil {
		return nil, err
	}
	b, err := protocol.Pack(format, helloResp)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// hello server describe
var helloServerDesc = frame.ServiceDesc{
	Version:     hello_version,
	ServiceName: "Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []frame.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    helloSayHelloHandler,
		},
	},
	ServiceURI: []string{
		"/services/" + hello_version + "/hello",
	},
	Metadata: "hello.Hello",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x94, 0xb8, 0x78, 0x3c,
	0x40, 0x8c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc,
	0x54, 0x09, 0x16, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x93, 0x8b, 0x17, 0xaa, 0xa6,
	0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d,
	0x55, 0x82, 0x11, 0xac, 0x0e, 0xc6, 0x35, 0x72, 0xe0, 0x62, 0x05, 0x2b, 0x15, 0x32, 0xe7, 0xe2,
	0x08, 0x4e, 0xac, 0x84, 0xb0, 0x85, 0xf5, 0x20, 0x16, 0x23, 0x5b, 0x24, 0x25, 0x82, 0x2a, 0x08,
	0x31, 0x59, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x3c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74,
	0x1d, 0x84, 0x03, 0xad, 0x00, 0x00, 0x00,
}