// Code generated by protoc-gen-go. DO NOT EDIT.
// source: signalcd/proto/agent.proto

package signalcdproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CurrentDeploymentRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrentDeploymentRequest) Reset()         { *m = CurrentDeploymentRequest{} }
func (m *CurrentDeploymentRequest) String() string { return proto.CompactTextString(m) }
func (*CurrentDeploymentRequest) ProtoMessage()    {}
func (*CurrentDeploymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7aa30ce6e358323, []int{0}
}

func (m *CurrentDeploymentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrentDeploymentRequest.Unmarshal(m, b)
}
func (m *CurrentDeploymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrentDeploymentRequest.Marshal(b, m, deterministic)
}
func (m *CurrentDeploymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentDeploymentRequest.Merge(m, src)
}
func (m *CurrentDeploymentRequest) XXX_Size() int {
	return xxx_messageInfo_CurrentDeploymentRequest.Size(m)
}
func (m *CurrentDeploymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentDeploymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentDeploymentRequest proto.InternalMessageInfo

type CurrentDeploymentResponse struct {
	CurrentDeployment    *Deployment `protobuf:"bytes,1,opt,name=currentDeployment,proto3" json:"currentDeployment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CurrentDeploymentResponse) Reset()         { *m = CurrentDeploymentResponse{} }
func (m *CurrentDeploymentResponse) String() string { return proto.CompactTextString(m) }
func (*CurrentDeploymentResponse) ProtoMessage()    {}
func (*CurrentDeploymentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7aa30ce6e358323, []int{1}
}

func (m *CurrentDeploymentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrentDeploymentResponse.Unmarshal(m, b)
}
func (m *CurrentDeploymentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrentDeploymentResponse.Marshal(b, m, deterministic)
}
func (m *CurrentDeploymentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentDeploymentResponse.Merge(m, src)
}
func (m *CurrentDeploymentResponse) XXX_Size() int {
	return xxx_messageInfo_CurrentDeploymentResponse.Size(m)
}
func (m *CurrentDeploymentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentDeploymentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentDeploymentResponse proto.InternalMessageInfo

func (m *CurrentDeploymentResponse) GetCurrentDeployment() *Deployment {
	if m != nil {
		return m.CurrentDeployment
	}
	return nil
}

type SetDeploymentStatusRequest struct {
	Number               int64             `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Status               *DeploymentStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SetDeploymentStatusRequest) Reset()         { *m = SetDeploymentStatusRequest{} }
func (m *SetDeploymentStatusRequest) String() string { return proto.CompactTextString(m) }
func (*SetDeploymentStatusRequest) ProtoMessage()    {}
func (*SetDeploymentStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7aa30ce6e358323, []int{2}
}

func (m *SetDeploymentStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDeploymentStatusRequest.Unmarshal(m, b)
}
func (m *SetDeploymentStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDeploymentStatusRequest.Marshal(b, m, deterministic)
}
func (m *SetDeploymentStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDeploymentStatusRequest.Merge(m, src)
}
func (m *SetDeploymentStatusRequest) XXX_Size() int {
	return xxx_messageInfo_SetDeploymentStatusRequest.Size(m)
}
func (m *SetDeploymentStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDeploymentStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDeploymentStatusRequest proto.InternalMessageInfo

func (m *SetDeploymentStatusRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *SetDeploymentStatusRequest) GetStatus() *DeploymentStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type SetDeploymentStatusResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDeploymentStatusResponse) Reset()         { *m = SetDeploymentStatusResponse{} }
func (m *SetDeploymentStatusResponse) String() string { return proto.CompactTextString(m) }
func (*SetDeploymentStatusResponse) ProtoMessage()    {}
func (*SetDeploymentStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7aa30ce6e358323, []int{3}
}

func (m *SetDeploymentStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDeploymentStatusResponse.Unmarshal(m, b)
}
func (m *SetDeploymentStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDeploymentStatusResponse.Marshal(b, m, deterministic)
}
func (m *SetDeploymentStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDeploymentStatusResponse.Merge(m, src)
}
func (m *SetDeploymentStatusResponse) XXX_Size() int {
	return xxx_messageInfo_SetDeploymentStatusResponse.Size(m)
}
func (m *SetDeploymentStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDeploymentStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetDeploymentStatusResponse proto.InternalMessageInfo

type StepLogsRequest struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Step                 int64    `protobuf:"varint,2,opt,name=step,proto3" json:"step,omitempty"`
	Logs                 []byte   `protobuf:"bytes,3,opt,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StepLogsRequest) Reset()         { *m = StepLogsRequest{} }
func (m *StepLogsRequest) String() string { return proto.CompactTextString(m) }
func (*StepLogsRequest) ProtoMessage()    {}
func (*StepLogsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7aa30ce6e358323, []int{4}
}

func (m *StepLogsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepLogsRequest.Unmarshal(m, b)
}
func (m *StepLogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepLogsRequest.Marshal(b, m, deterministic)
}
func (m *StepLogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepLogsRequest.Merge(m, src)
}
func (m *StepLogsRequest) XXX_Size() int {
	return xxx_messageInfo_StepLogsRequest.Size(m)
}
func (m *StepLogsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StepLogsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StepLogsRequest proto.InternalMessageInfo

func (m *StepLogsRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *StepLogsRequest) GetStep() int64 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *StepLogsRequest) GetLogs() []byte {
	if m != nil {
		return m.Logs
	}
	return nil
}

type StepLogsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StepLogsResponse) Reset()         { *m = StepLogsResponse{} }
func (m *StepLogsResponse) String() string { return proto.CompactTextString(m) }
func (*StepLogsResponse) ProtoMessage()    {}
func (*StepLogsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7aa30ce6e358323, []int{5}
}

func (m *StepLogsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepLogsResponse.Unmarshal(m, b)
}
func (m *StepLogsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepLogsResponse.Marshal(b, m, deterministic)
}
func (m *StepLogsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepLogsResponse.Merge(m, src)
}
func (m *StepLogsResponse) XXX_Size() int {
	return xxx_messageInfo_StepLogsResponse.Size(m)
}
func (m *StepLogsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StepLogsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StepLogsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CurrentDeploymentRequest)(nil), "signalcd.CurrentDeploymentRequest")
	proto.RegisterType((*CurrentDeploymentResponse)(nil), "signalcd.CurrentDeploymentResponse")
	proto.RegisterType((*SetDeploymentStatusRequest)(nil), "signalcd.SetDeploymentStatusRequest")
	proto.RegisterType((*SetDeploymentStatusResponse)(nil), "signalcd.SetDeploymentStatusResponse")
	proto.RegisterType((*StepLogsRequest)(nil), "signalcd.StepLogsRequest")
	proto.RegisterType((*StepLogsResponse)(nil), "signalcd.StepLogsResponse")
}

func init() { proto.RegisterFile("signalcd/proto/agent.proto", fileDescriptor_e7aa30ce6e358323) }

var fileDescriptor_e7aa30ce6e358323 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x51, 0x4b, 0xc3, 0x30,
	0x10, 0xc7, 0xd9, 0x2a, 0x65, 0x9c, 0x93, 0xb9, 0x28, 0xd2, 0x45, 0x84, 0x11, 0x15, 0xf6, 0xd4,
	0x41, 0xfd, 0x04, 0x9b, 0x3e, 0xfa, 0x62, 0xfb, 0x26, 0x82, 0xb4, 0xf5, 0xa8, 0x83, 0x2e, 0x89,
	0x4d, 0x2a, 0xec, 0x73, 0xf8, 0x85, 0x65, 0x69, 0x4a, 0xeb, 0xd6, 0xb9, 0xa7, 0x5e, 0xef, 0xff,
	0xbf, 0xfb, 0x71, 0x77, 0x01, 0xaa, 0x56, 0x19, 0x8f, 0xf3, 0xf4, 0x63, 0x2e, 0x0b, 0xa1, 0xc5,
	0x3c, 0xce, 0x90, 0x6b, 0xdf, 0xc4, 0x64, 0x50, 0x6b, 0x74, 0xd7, 0xa5, 0x37, 0x12, 0x55, 0xe5,
	0x62, 0x14, 0xbc, 0xc7, 0xb2, 0x28, 0x90, 0xeb, 0x27, 0x94, 0xb9, 0xd8, 0xac, 0x91, 0xeb, 0x10,
	0xbf, 0x4a, 0x54, 0x9a, 0xbd, 0xc3, 0xa4, 0x43, 0x53, 0x52, 0x70, 0x85, 0x64, 0x09, 0xe3, 0x74,
	0x57, 0xf4, 0x7a, 0xd3, 0xde, 0xec, 0x34, 0xb8, 0xf4, 0x6b, 0xa0, 0xdf, 0x2a, 0xdc, 0xb7, 0xb3,
	0x4f, 0xa0, 0x11, 0xb6, 0x12, 0x91, 0x8e, 0x75, 0xa9, 0x2c, 0x9e, 0x5c, 0x81, 0xcb, 0xcb, 0x75,
	0x82, 0x85, 0x69, 0xeb, 0x84, 0xf6, 0x8f, 0x04, 0xe0, 0x2a, 0x63, 0xf4, 0xfa, 0x06, 0x47, 0xbb,
	0x70, 0xb6, 0x95, 0x75, 0xb2, 0x1b, 0xb8, 0xee, 0x24, 0x55, 0xc3, 0xb0, 0x17, 0x18, 0x45, 0x1a,
	0xe5, 0xb3, 0xc8, 0x8e, 0xd2, 0x09, 0x9c, 0x28, 0x8d, 0xd2, 0xb0, 0x9d, 0xd0, 0xc4, 0xdb, 0x5c,
	0x2e, 0x32, 0xe5, 0x39, 0xd3, 0xde, 0x6c, 0x18, 0x9a, 0x98, 0x11, 0x38, 0x6f, 0x5a, 0x56, 0x98,
	0xe0, 0xa7, 0x0f, 0xc3, 0xc5, 0xf6, 0x44, 0x11, 0x16, 0xdf, 0xab, 0x14, 0xc9, 0x1b, 0x8c, 0xf7,
	0x36, 0x4c, 0x58, 0x33, 0xcf, 0xa1, 0xd3, 0xd0, 0xdb, 0x7f, 0x3d, 0xf6, 0x44, 0x09, 0x5c, 0x74,
	0x0c, 0x4d, 0xee, 0x9a, 0xda, 0xc3, 0xdb, 0xa7, 0xf7, 0x47, 0x5c, 0x96, 0xb1, 0x80, 0x41, 0x3d,
	0x26, 0x99, 0xb4, 0x4a, 0xfe, 0x6e, 0x93, 0xd2, 0x2e, 0xa9, 0x6a, 0xb1, 0x1c, 0xbd, 0x9e, 0xd5,
	0xa2, 0x79, 0x93, 0x89, 0x6b, 0x3e, 0x0f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x61, 0xf9, 0x32,
	0x40, 0xde, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentServiceClient interface {
	CurrentDeployment(ctx context.Context, in *CurrentDeploymentRequest, opts ...grpc.CallOption) (*CurrentDeploymentResponse, error)
	SetDeploymentStatus(ctx context.Context, in *SetDeploymentStatusRequest, opts ...grpc.CallOption) (*SetDeploymentStatusResponse, error)
	StepLogs(ctx context.Context, in *StepLogsRequest, opts ...grpc.CallOption) (*StepLogsResponse, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) CurrentDeployment(ctx context.Context, in *CurrentDeploymentRequest, opts ...grpc.CallOption) (*CurrentDeploymentResponse, error) {
	out := new(CurrentDeploymentResponse)
	err := c.cc.Invoke(ctx, "/signalcd.AgentService/CurrentDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) SetDeploymentStatus(ctx context.Context, in *SetDeploymentStatusRequest, opts ...grpc.CallOption) (*SetDeploymentStatusResponse, error) {
	out := new(SetDeploymentStatusResponse)
	err := c.cc.Invoke(ctx, "/signalcd.AgentService/SetDeploymentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) StepLogs(ctx context.Context, in *StepLogsRequest, opts ...grpc.CallOption) (*StepLogsResponse, error) {
	out := new(StepLogsResponse)
	err := c.cc.Invoke(ctx, "/signalcd.AgentService/StepLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
type AgentServiceServer interface {
	CurrentDeployment(context.Context, *CurrentDeploymentRequest) (*CurrentDeploymentResponse, error)
	SetDeploymentStatus(context.Context, *SetDeploymentStatusRequest) (*SetDeploymentStatusResponse, error)
	StepLogs(context.Context, *StepLogsRequest) (*StepLogsResponse, error)
}

// UnimplementedAgentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (*UnimplementedAgentServiceServer) CurrentDeployment(ctx context.Context, req *CurrentDeploymentRequest) (*CurrentDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentDeployment not implemented")
}
func (*UnimplementedAgentServiceServer) SetDeploymentStatus(ctx context.Context, req *SetDeploymentStatusRequest) (*SetDeploymentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeploymentStatus not implemented")
}
func (*UnimplementedAgentServiceServer) StepLogs(ctx context.Context, req *StepLogsRequest) (*StepLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StepLogs not implemented")
}

func RegisterAgentServiceServer(s *grpc.Server, srv AgentServiceServer) {
	s.RegisterService(&_AgentService_serviceDesc, srv)
}

func _AgentService_CurrentDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).CurrentDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signalcd.AgentService/CurrentDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).CurrentDeployment(ctx, req.(*CurrentDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_SetDeploymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeploymentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).SetDeploymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signalcd.AgentService/SetDeploymentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).SetDeploymentStatus(ctx, req.(*SetDeploymentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_StepLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StepLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).StepLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signalcd.AgentService/StepLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).StepLogs(ctx, req.(*StepLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "signalcd.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CurrentDeployment",
			Handler:    _AgentService_CurrentDeployment_Handler,
		},
		{
			MethodName: "SetDeploymentStatus",
			Handler:    _AgentService_SetDeploymentStatus_Handler,
		},
		{
			MethodName: "StepLogs",
			Handler:    _AgentService_StepLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "signalcd/proto/agent.proto",
}
