// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remote_command.proto

/*
Package telegram is a generated protocol buffer package.

It is generated from these files:
	remote_command.proto

It has these top-level messages:
	RemoteCommandRequest
	RemoteRequest
*/
package remoteTelegramCommands

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RemoteCommandRequest struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Group       int64  `protobuf:"varint,3,opt,name=group" json:"group,omitempty"`
}

func (m *RemoteCommandRequest) Reset()                    { *m = RemoteCommandRequest{} }
func (m *RemoteCommandRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoteCommandRequest) ProtoMessage()               {}
func (*RemoteCommandRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RemoteCommandRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RemoteCommandRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RemoteCommandRequest) GetGroup() int64 {
	if m != nil {
		return m.Group
	}
	return 0
}

type RemoteRequest struct {
	From    string `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *RemoteRequest) Reset()                    { *m = RemoteRequest{} }
func (m *RemoteRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoteRequest) ProtoMessage()               {}
func (*RemoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RemoteRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *RemoteRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RemoteCommandRequest)(nil), "telegram.RemoteCommandRequest")
	proto.RegisterType((*RemoteRequest)(nil), "telegram.RemoteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RemoteCommand service

type RemoteCommandClient interface {
	RegisterCommand(ctx context.Context, in *RemoteCommandRequest, opts ...grpc.CallOption) (RemoteCommand_RegisterCommandClient, error)
}

type remoteCommandClient struct {
	cc *grpc.ClientConn
}

func NewRemoteCommandClient(cc *grpc.ClientConn) RemoteCommandClient {
	return &remoteCommandClient{cc}
}

func (c *remoteCommandClient) RegisterCommand(ctx context.Context, in *RemoteCommandRequest, opts ...grpc.CallOption) (RemoteCommand_RegisterCommandClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RemoteCommand_serviceDesc.Streams[0], c.cc, "/telegram.RemoteCommand/RegisterCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &remoteCommandRegisterCommandClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RemoteCommand_RegisterCommandClient interface {
	Recv() (*RemoteRequest, error)
	grpc.ClientStream
}

type remoteCommandRegisterCommandClient struct {
	grpc.ClientStream
}

func (x *remoteCommandRegisterCommandClient) Recv() (*RemoteRequest, error) {
	m := new(RemoteRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RemoteCommand service

type RemoteCommandServer interface {
	RegisterCommand(*RemoteCommandRequest, RemoteCommand_RegisterCommandServer) error
}

func RegisterRemoteCommandServer(s *grpc.Server, srv RemoteCommandServer) {
	s.RegisterService(&_RemoteCommand_serviceDesc, srv)
}

func _RemoteCommand_RegisterCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RemoteCommandRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RemoteCommandServer).RegisterCommand(m, &remoteCommandRegisterCommandServer{stream})
}

type RemoteCommand_RegisterCommandServer interface {
	Send(*RemoteRequest) error
	grpc.ServerStream
}

type remoteCommandRegisterCommandServer struct {
	grpc.ServerStream
}

func (x *remoteCommandRegisterCommandServer) Send(m *RemoteRequest) error {
	return x.ServerStream.SendMsg(m)
}

var _RemoteCommand_serviceDesc = grpc.ServiceDesc{
	ServiceName: "telegram.RemoteCommand",
	HandlerType: (*RemoteCommandServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterCommand",
			Handler:       _RemoteCommand_RegisterCommand_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "remote_command.proto",
}

func init() { proto.RegisterFile("remote_command.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xbb, 0xca, 0xc2, 0x30,
	0x18, 0xfd, 0xf3, 0xd7, 0xeb, 0x27, 0x22, 0x84, 0x82, 0xc1, 0x41, 0x4a, 0xa7, 0x4e, 0x45, 0x74,
	0x76, 0x72, 0x77, 0xc8, 0x0b, 0x94, 0xb4, 0xfd, 0x0c, 0x05, 0xd3, 0xd4, 0x24, 0x7d, 0x7f, 0x31,
	0x6d, 0xa0, 0x88, 0xdb, 0xb9, 0xe5, 0x24, 0x27, 0x10, 0x1b, 0x54, 0xda, 0x61, 0x51, 0x69, 0xa5,
	0x44, 0x5b, 0xe7, 0x9d, 0xd1, 0x4e, 0xd3, 0x95, 0xc3, 0x27, 0x4a, 0x23, 0x54, 0x5a, 0x42, 0xcc,
	0x7d, 0xe2, 0x36, 0x04, 0x38, 0xbe, 0x7a, 0xb4, 0x8e, 0x52, 0x98, 0xb5, 0x42, 0x21, 0x23, 0x09,
	0xc9, 0xd6, 0xdc, 0x63, 0x9a, 0xc0, 0xa6, 0x46, 0x5b, 0x99, 0xa6, 0x73, 0x8d, 0x6e, 0xd9, 0xbf,
	0xb7, 0xa6, 0x12, 0x8d, 0x61, 0x2e, 0x8d, 0xee, 0x3b, 0x16, 0x25, 0x24, 0x8b, 0xf8, 0x40, 0xd2,
	0x2b, 0x6c, 0x87, 0x3b, 0x26, 0xe5, 0x0f, 0xa3, 0x55, 0x28, 0xff, 0x60, 0xca, 0x60, 0xa9, 0xd0,
	0x5a, 0x21, 0x71, 0x2c, 0x0e, 0xf4, 0x5c, 0x84, 0xe3, 0xe3, 0x13, 0xe9, 0x1d, 0x76, 0x1c, 0x65,
	0x63, 0x1d, 0x9a, 0x20, 0x1d, 0xf3, 0xb0, 0x28, 0xff, 0x35, 0xe7, 0xb0, 0xff, 0xf6, 0x47, 0x23,
	0xfd, 0x3b, 0x91, 0x72, 0xe1, 0x3f, 0xe5, 0xf2, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x45, 0x45, 0xbf,
	0xb8, 0x2c, 0x01, 0x00, 0x00,
}
