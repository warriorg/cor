// Code generated by protoc-gen-go.
// source: task.proto
// DO NOT EDIT!

/*
Package task is a generated protocol buffer package.

It is generated from these files:
	task.proto

It has these top-level messages:
	TaskRequest
	TaskResponse
*/
package task

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type TaskRequest struct {
	Time     *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=time" json:"time,omitempty"`
	EndTime  *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=endTime" json:"endTime,omitempty"`
	MaxCount int32                      `protobuf:"varint,3,opt,name=maxCount" json:"maxCount,omitempty"`
	Every    int32                      `protobuf:"varint,4,opt,name=every" json:"every,omitempty"`
	Unit     string                     `protobuf:"bytes,5,opt,name=unit" json:"unit,omitempty"`
	Url      string                     `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	Method   string                     `protobuf:"bytes,7,opt,name=method" json:"method,omitempty"`
	Body     string                     `protobuf:"bytes,8,opt,name=body" json:"body,omitempty"`
	Header   map[string]string          `protobuf:"bytes,9,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *TaskRequest) Reset()                    { *m = TaskRequest{} }
func (m *TaskRequest) String() string            { return proto.CompactTextString(m) }
func (*TaskRequest) ProtoMessage()               {}
func (*TaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TaskRequest) GetTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *TaskRequest) GetEndTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *TaskRequest) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

type TaskResponse struct {
	Code    int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *TaskResponse) Reset()                    { *m = TaskResponse{} }
func (m *TaskResponse) String() string            { return proto.CompactTextString(m) }
func (*TaskResponse) ProtoMessage()               {}
func (*TaskResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*TaskRequest)(nil), "task.TaskRequest")
	proto.RegisterType((*TaskResponse)(nil), "task.TaskResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for TaskService service

type TaskServiceClient interface {
}

type taskServiceClient struct {
	cc *grpc.ClientConn
}

func NewTaskServiceClient(cc *grpc.ClientConn) TaskServiceClient {
	return &taskServiceClient{cc}
}

// Server API for TaskService service

type TaskServiceServer interface {
}

func RegisterTaskServiceServer(s *grpc.Server, srv TaskServiceServer) {
	s.RegisterService(&_TaskService_serviceDesc, srv)
}

var _TaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "task.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    fileDescriptor0,
}

func init() { proto.RegisterFile("task.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x49, 0x9b, 0xa4, 0xcd, 0xcd, 0xf7, 0x81, 0x0c, 0x22, 0x43, 0x40, 0x0c, 0x5d, 0x65,
	0x35, 0x85, 0xaa, 0xa0, 0xe2, 0x4e, 0x04, 0xd7, 0x63, 0x5f, 0x60, 0xda, 0x5c, 0xdb, 0xd2, 0x26,
	0x53, 0xe7, 0x4f, 0x30, 0x8f, 0xec, 0x5b, 0xc8, 0xcc, 0x34, 0xd2, 0x9d, 0xbb, 0x73, 0xce, 0x9c,
	0x5f, 0xb8, 0x9c, 0x00, 0x18, 0xa1, 0xf7, 0xec, 0xa8, 0xa4, 0x91, 0x24, 0x76, 0xba, 0xb8, 0xd9,
	0x48, 0xb9, 0x39, 0xe0, 0xdc, 0x67, 0x2b, 0xfb, 0x31, 0x37, 0xbb, 0x06, 0xb5, 0x11, 0xcd, 0x31,
	0xd4, 0x66, 0xdf, 0x23, 0xc8, 0x97, 0x42, 0xef, 0x39, 0x7e, 0x5a, 0xd4, 0x86, 0x30, 0x88, 0x5d,
	0x85, 0x46, 0x65, 0x54, 0xe5, 0x8b, 0x82, 0x05, 0x9e, 0x0d, 0x3c, 0x5b, 0x0e, 0x3c, 0xf7, 0x3d,
	0x72, 0x07, 0x13, 0x6c, 0x6b, 0x97, 0xd2, 0xd1, 0x9f, 0xc8, 0x50, 0x25, 0x05, 0x4c, 0x1b, 0xf1,
	0xf5, 0x22, 0x6d, 0x6b, 0xe8, 0xb8, 0x8c, 0xaa, 0x84, 0xff, 0x7a, 0x72, 0x09, 0x09, 0x76, 0xa8,
	0x7a, 0x1a, 0xfb, 0x87, 0x60, 0x08, 0x81, 0xd8, 0xb6, 0x3b, 0x43, 0x93, 0x32, 0xaa, 0x32, 0xee,
	0x35, 0xb9, 0x80, 0xb1, 0x55, 0x07, 0x9a, 0xfa, 0xc8, 0x49, 0x72, 0x05, 0x69, 0x83, 0x66, 0x2b,
	0x6b, 0x3a, 0xf1, 0xe1, 0xc9, 0x39, 0x7a, 0x25, 0xeb, 0x9e, 0x4e, 0x03, 0xed, 0x34, 0xb9, 0x87,
	0x74, 0x8b, 0xa2, 0x46, 0x45, 0xb3, 0x72, 0x5c, 0xe5, 0x8b, 0x6b, 0xe6, 0xd7, 0x3b, 0x1b, 0x83,
	0xbd, 0xf9, 0xf7, 0xd7, 0xd6, 0xa8, 0x9e, 0x9f, 0xca, 0xc5, 0x23, 0xe4, 0x67, 0xb1, 0xbb, 0x61,
	0x8f, 0xbd, 0x9f, 0x2b, 0xe3, 0x4e, 0xba, 0xfb, 0x3b, 0x71, 0xb0, 0x61, 0x8f, 0x8c, 0x07, 0xf3,
	0x34, 0x7a, 0x88, 0x66, 0xcf, 0xf0, 0x2f, 0x7c, 0x5d, 0x1f, 0x65, 0xab, 0xd1, 0x5d, 0xb5, 0x96,
	0x75, 0xd8, 0x3a, 0xe1, 0x5e, 0x13, 0x0a, 0x93, 0x06, 0xb5, 0x16, 0x9b, 0x81, 0x1f, 0xec, 0xe2,
	0x7f, 0xf8, 0x51, 0xef, 0xa8, 0xba, 0xdd, 0x1a, 0x57, 0xa9, 0xdf, 0xf7, 0xf6, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x0d, 0xa3, 0x3d, 0x02, 0xf4, 0x01, 0x00, 0x00,
}
