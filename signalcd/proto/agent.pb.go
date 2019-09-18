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

type DeploymentStatus_Phase int32

const (
	DeploymentStatus_UNKNOWN  DeploymentStatus_Phase = 0
	DeploymentStatus_SUCCESS  DeploymentStatus_Phase = 1
	DeploymentStatus_FAILURE  DeploymentStatus_Phase = 2
	DeploymentStatus_PROGRESS DeploymentStatus_Phase = 3
	DeploymentStatus_PENDING  DeploymentStatus_Phase = 4
	DeploymentStatus_KILLED   DeploymentStatus_Phase = 5
)

var DeploymentStatus_Phase_name = map[int32]string{
	0: "UNKNOWN",
	1: "SUCCESS",
	2: "FAILURE",
	3: "PROGRESS",
	4: "PENDING",
	5: "KILLED",
}

var DeploymentStatus_Phase_value = map[string]int32{
	"UNKNOWN":  0,
	"SUCCESS":  1,
	"FAILURE":  2,
	"PROGRESS": 3,
	"PENDING":  4,
	"KILLED":   5,
}

func (x DeploymentStatus_Phase) String() string {
	return proto.EnumName(DeploymentStatus_Phase_name, int32(x))
}

func (DeploymentStatus_Phase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e7aa30ce6e358323, []int{1, 0}
}

type Deployment struct {
	Number               int64             `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Created              int64             `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Started              int64             `protobuf:"varint,3,opt,name=started,proto3" json:"started,omitempty"`
	Finished             int64             `protobuf:"varint,4,opt,name=finished,proto3" json:"finished,omitempty"`
	Status               *DeploymentStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Pipeline             *Pipeline         `protobuf:"bytes,6,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Deployment) Reset()         { *m = Deployment{} }
func (m *Deployment) String() string { return proto.CompactTextString(m) }
func (*Deployment) ProtoMessage()    {}
func (*Deployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7aa30ce6e358323, []int{0}
}

func (m *Deployment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deployment.Unmarshal(m, b)
}
func (m *Deployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deployment.Marshal(b, m, deterministic)
}
func (m *Deployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deployment.Merge(m, src)
}
func (m *Deployment) XXX_Size() int {
	return xxx_messageInfo_Deployment.Size(m)
}
func (m *Deployment) XXX_DiscardUnknown() {
	xxx_messageInfo_Deployment.DiscardUnknown(m)
}

var xxx_messageInfo_Deployment proto.InternalMessageInfo

func (m *Deployment) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Deployment) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Deployment) GetStarted() int64 {
	if m != nil {
		return m.Started
	}
	return 0
}

func (m *Deployment) GetFinished() int64 {
	if m != nil {
		return m.Finished
	}
	return 0
}

func (m *Deployment) GetStatus() *DeploymentStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Deployment) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

type DeploymentStatus struct {
	Phase                DeploymentStatus_Phase `protobuf:"varint,1,opt,name=phase,proto3,enum=signalcd.DeploymentStatus_Phase" json:"phase,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DeploymentStatus) Reset()         { *m = DeploymentStatus{} }
func (m *DeploymentStatus) String() string { return proto.CompactTextString(m) }
func (*DeploymentStatus) ProtoMessage()    {}
func (*DeploymentStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7aa30ce6e358323, []int{1}
}

func (m *DeploymentStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeploymentStatus.Unmarshal(m, b)
}
func (m *DeploymentStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeploymentStatus.Marshal(b, m, deterministic)
}
func (m *DeploymentStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeploymentStatus.Merge(m, src)
}
func (m *DeploymentStatus) XXX_Size() int {
	return xxx_messageInfo_DeploymentStatus.Size(m)
}
func (m *DeploymentStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_DeploymentStatus.DiscardUnknown(m)
}

var xxx_messageInfo_DeploymentStatus proto.InternalMessageInfo

func (m *DeploymentStatus) GetPhase() DeploymentStatus_Phase {
	if m != nil {
		return m.Phase
	}
	return DeploymentStatus_UNKNOWN
}

type CurrentDeploymentRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrentDeploymentRequest) Reset()         { *m = CurrentDeploymentRequest{} }
func (m *CurrentDeploymentRequest) String() string { return proto.CompactTextString(m) }
func (*CurrentDeploymentRequest) ProtoMessage()    {}
func (*CurrentDeploymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7aa30ce6e358323, []int{2}
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
	return fileDescriptor_e7aa30ce6e358323, []int{3}
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
	return fileDescriptor_e7aa30ce6e358323, []int{4}
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
	return fileDescriptor_e7aa30ce6e358323, []int{5}
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

type Pipeline struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Steps                []*Step  `protobuf:"bytes,3,rep,name=steps,proto3" json:"steps,omitempty"`
	Checks               []*Check `protobuf:"bytes,4,rep,name=checks,proto3" json:"checks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pipeline) Reset()         { *m = Pipeline{} }
func (m *Pipeline) String() string { return proto.CompactTextString(m) }
func (*Pipeline) ProtoMessage()    {}
func (*Pipeline) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7aa30ce6e358323, []int{6}
}

func (m *Pipeline) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pipeline.Unmarshal(m, b)
}
func (m *Pipeline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pipeline.Marshal(b, m, deterministic)
}
func (m *Pipeline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pipeline.Merge(m, src)
}
func (m *Pipeline) XXX_Size() int {
	return xxx_messageInfo_Pipeline.Size(m)
}
func (m *Pipeline) XXX_DiscardUnknown() {
	xxx_messageInfo_Pipeline.DiscardUnknown(m)
}

var xxx_messageInfo_Pipeline proto.InternalMessageInfo

func (m *Pipeline) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Pipeline) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pipeline) GetSteps() []*Step {
	if m != nil {
		return m.Steps
	}
	return nil
}

func (m *Pipeline) GetChecks() []*Check {
	if m != nil {
		return m.Checks
	}
	return nil
}

type Step struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image                string   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	ImagePullSecrets     []string `protobuf:"bytes,3,rep,name=ImagePullSecrets,proto3" json:"ImagePullSecrets,omitempty"`
	Commands             []string `protobuf:"bytes,4,rep,name=commands,proto3" json:"commands,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Step) Reset()         { *m = Step{} }
func (m *Step) String() string { return proto.CompactTextString(m) }
func (*Step) ProtoMessage()    {}
func (*Step) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7aa30ce6e358323, []int{7}
}

func (m *Step) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Step.Unmarshal(m, b)
}
func (m *Step) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Step.Marshal(b, m, deterministic)
}
func (m *Step) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Step.Merge(m, src)
}
func (m *Step) XXX_Size() int {
	return xxx_messageInfo_Step.Size(m)
}
func (m *Step) XXX_DiscardUnknown() {
	xxx_messageInfo_Step.DiscardUnknown(m)
}

var xxx_messageInfo_Step proto.InternalMessageInfo

func (m *Step) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Step) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Step) GetImagePullSecrets() []string {
	if m != nil {
		return m.ImagePullSecrets
	}
	return nil
}

func (m *Step) GetCommands() []string {
	if m != nil {
		return m.Commands
	}
	return nil
}

type Check struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image                string   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	ImagePullSecrets     []string `protobuf:"bytes,3,rep,name=ImagePullSecrets,proto3" json:"ImagePullSecrets,omitempty"`
	Duration             int64    `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Check) Reset()         { *m = Check{} }
func (m *Check) String() string { return proto.CompactTextString(m) }
func (*Check) ProtoMessage()    {}
func (*Check) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7aa30ce6e358323, []int{8}
}

func (m *Check) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Check.Unmarshal(m, b)
}
func (m *Check) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Check.Marshal(b, m, deterministic)
}
func (m *Check) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Check.Merge(m, src)
}
func (m *Check) XXX_Size() int {
	return xxx_messageInfo_Check.Size(m)
}
func (m *Check) XXX_DiscardUnknown() {
	xxx_messageInfo_Check.DiscardUnknown(m)
}

var xxx_messageInfo_Check proto.InternalMessageInfo

func (m *Check) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Check) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Check) GetImagePullSecrets() []string {
	if m != nil {
		return m.ImagePullSecrets
	}
	return nil
}

func (m *Check) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type AgentStatusRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentStatusRequest) Reset()         { *m = AgentStatusRequest{} }
func (m *AgentStatusRequest) String() string { return proto.CompactTextString(m) }
func (*AgentStatusRequest) ProtoMessage()    {}
func (*AgentStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7aa30ce6e358323, []int{9}
}

func (m *AgentStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentStatusRequest.Unmarshal(m, b)
}
func (m *AgentStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentStatusRequest.Marshal(b, m, deterministic)
}
func (m *AgentStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentStatusRequest.Merge(m, src)
}
func (m *AgentStatusRequest) XXX_Size() int {
	return xxx_messageInfo_AgentStatusRequest.Size(m)
}
func (m *AgentStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AgentStatusRequest proto.InternalMessageInfo

func (m *AgentStatusRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AgentStatusResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentStatusResponse) Reset()         { *m = AgentStatusResponse{} }
func (m *AgentStatusResponse) String() string { return proto.CompactTextString(m) }
func (*AgentStatusResponse) ProtoMessage()    {}
func (*AgentStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7aa30ce6e358323, []int{10}
}

func (m *AgentStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentStatusResponse.Unmarshal(m, b)
}
func (m *AgentStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentStatusResponse.Marshal(b, m, deterministic)
}
func (m *AgentStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentStatusResponse.Merge(m, src)
}
func (m *AgentStatusResponse) XXX_Size() int {
	return xxx_messageInfo_AgentStatusResponse.Size(m)
}
func (m *AgentStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AgentStatusResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("signalcd.DeploymentStatus_Phase", DeploymentStatus_Phase_name, DeploymentStatus_Phase_value)
	proto.RegisterType((*Deployment)(nil), "signalcd.Deployment")
	proto.RegisterType((*DeploymentStatus)(nil), "signalcd.DeploymentStatus")
	proto.RegisterType((*CurrentDeploymentRequest)(nil), "signalcd.CurrentDeploymentRequest")
	proto.RegisterType((*CurrentDeploymentResponse)(nil), "signalcd.CurrentDeploymentResponse")
	proto.RegisterType((*SetDeploymentStatusRequest)(nil), "signalcd.SetDeploymentStatusRequest")
	proto.RegisterType((*SetDeploymentStatusResponse)(nil), "signalcd.SetDeploymentStatusResponse")
	proto.RegisterType((*Pipeline)(nil), "signalcd.Pipeline")
	proto.RegisterType((*Step)(nil), "signalcd.Step")
	proto.RegisterType((*Check)(nil), "signalcd.Check")
	proto.RegisterType((*AgentStatusRequest)(nil), "signalcd.AgentStatusRequest")
	proto.RegisterType((*AgentStatusResponse)(nil), "signalcd.AgentStatusResponse")
}

func init() { proto.RegisterFile("signalcd/proto/agent.proto", fileDescriptor_e7aa30ce6e358323) }

var fileDescriptor_e7aa30ce6e358323 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x49, 0x9c, 0x26, 0x93, 0xd2, 0xba, 0xdb, 0x82, 0x8c, 0xa1, 0x52, 0xb4, 0x14, 0x11,
	0x71, 0x48, 0x25, 0x23, 0x71, 0x6f, 0xd3, 0x50, 0x85, 0x56, 0x69, 0xb4, 0x56, 0x84, 0x84, 0x90,
	0x90, 0x63, 0x0f, 0x89, 0x45, 0xfc, 0x81, 0x77, 0x8d, 0xe8, 0x81, 0x7f, 0xc2, 0x85, 0x7f, 0xc5,
	0xcf, 0x41, 0x5e, 0xdb, 0x71, 0xa8, 0x13, 0xe0, 0xc2, 0xc9, 0xfb, 0xf6, 0xbd, 0xd9, 0x37, 0xb3,
	0x33, 0x6b, 0x30, 0xb8, 0x37, 0x0f, 0xec, 0xa5, 0xe3, 0x9e, 0x46, 0x71, 0x28, 0xc2, 0x53, 0x7b,
	0x8e, 0x81, 0xe8, 0xcb, 0x35, 0x69, 0x15, 0x1c, 0xfd, 0xa9, 0x00, 0x5c, 0x60, 0xb4, 0x0c, 0x6f,
	0x7d, 0x0c, 0x04, 0x79, 0x08, 0xcd, 0x20, 0xf1, 0x67, 0x18, 0xeb, 0x4a, 0x57, 0xe9, 0xd5, 0x59,
	0x8e, 0x88, 0x0e, 0x3b, 0x4e, 0x8c, 0xb6, 0x40, 0x57, 0xaf, 0x49, 0xa2, 0x80, 0x29, 0xc3, 0x85,
	0x1d, 0xa7, 0x4c, 0x3d, 0x63, 0x72, 0x48, 0x0c, 0x68, 0x7d, 0xf4, 0x02, 0x8f, 0x2f, 0xd0, 0xd5,
	0x1b, 0x92, 0x5a, 0x61, 0x62, 0x42, 0x93, 0x0b, 0x5b, 0x24, 0x5c, 0x57, 0xbb, 0x4a, 0xaf, 0x63,
	0x1a, 0xfd, 0x22, 0xa3, 0x7e, 0x99, 0x8d, 0x25, 0x15, 0x2c, 0x57, 0x92, 0x3e, 0xb4, 0x22, 0x2f,
	0xc2, 0xa5, 0x17, 0xa0, 0xde, 0x94, 0x51, 0xa4, 0x8c, 0x9a, 0xe4, 0x0c, 0x5b, 0x69, 0xe8, 0x0f,
	0x05, 0xb4, 0xbb, 0x87, 0x91, 0x57, 0xa0, 0x46, 0x0b, 0x9b, 0xa3, 0xac, 0x6f, 0xcf, 0xec, 0x6e,
	0xf7, 0xed, 0x4f, 0x52, 0x1d, 0xcb, 0xe4, 0x74, 0x0a, 0xaa, 0xc4, 0xa4, 0x03, 0x3b, 0xd3, 0xf1,
	0xd5, 0xf8, 0xe6, 0xed, 0x58, 0xbb, 0x97, 0x02, 0x6b, 0x3a, 0x18, 0x0c, 0x2d, 0x4b, 0x53, 0x52,
	0xf0, 0xfa, 0x6c, 0x74, 0x3d, 0x65, 0x43, 0xad, 0x46, 0x76, 0xa1, 0x35, 0x61, 0x37, 0x97, 0x2c,
	0xa5, 0xea, 0x29, 0x35, 0x19, 0x8e, 0x2f, 0x46, 0xe3, 0x4b, 0xad, 0x41, 0x00, 0x9a, 0x57, 0xa3,
	0xeb, 0xeb, 0xe1, 0x85, 0xa6, 0x52, 0x03, 0xf4, 0x41, 0x12, 0xc7, 0x18, 0x88, 0xd2, 0x9e, 0xe1,
	0xe7, 0x04, 0xb9, 0xa0, 0x1f, 0xe0, 0xd1, 0x06, 0x8e, 0x47, 0x61, 0xc0, 0x91, 0x9c, 0xc3, 0x81,
	0x73, 0x97, 0x94, 0x35, 0x75, 0xcc, 0xa3, 0x4d, 0x35, 0xb1, 0xaa, 0x9c, 0x2e, 0xc0, 0xb0, 0x50,
	0x54, 0xee, 0x3b, 0xb3, 0xdf, 0x3a, 0x0a, 0x65, 0xeb, 0x6a, 0xff, 0xda, 0x3a, 0x7a, 0x0c, 0x8f,
	0x37, 0x3a, 0x65, 0xc5, 0xd0, 0x6f, 0xd0, 0x2a, 0xfa, 0x47, 0xf6, 0xa0, 0xe6, 0xb9, 0xd2, 0xb2,
	0xcd, 0x6a, 0x9e, 0x4b, 0x08, 0x34, 0x02, 0xdb, 0x47, 0x69, 0xd6, 0x66, 0x72, 0x4d, 0x4e, 0x40,
	0xe5, 0x02, 0x23, 0xae, 0xd7, 0xbb, 0xf5, 0x5e, 0xc7, 0xdc, 0x2b, 0x33, 0xb0, 0x04, 0x46, 0x2c,
	0x23, 0xc9, 0x73, 0x68, 0x3a, 0x0b, 0x74, 0x3e, 0x71, 0xbd, 0x21, 0x65, 0xfb, 0xa5, 0x6c, 0x90,
	0xee, 0xb3, 0x9c, 0xa6, 0x5f, 0xa1, 0x91, 0xc6, 0xad, 0xac, 0x94, 0x35, 0xab, 0x23, 0x50, 0x3d,
	0xdf, 0x9e, 0x17, 0xfe, 0x19, 0x20, 0x2f, 0x40, 0x1b, 0xa5, 0x8b, 0x49, 0xb2, 0x5c, 0x5a, 0xe8,
	0xc4, 0x28, 0xb2, 0x5c, 0xda, 0xac, 0xb2, 0x9f, 0x3e, 0x03, 0x27, 0xf4, 0x7d, 0x3b, 0x70, 0xb3,
	0x44, 0xda, 0x6c, 0x85, 0xe9, 0x2d, 0xa8, 0x32, 0x95, 0xff, 0x67, 0xed, 0x26, 0xb1, 0x2d, 0xbc,
	0x30, 0x28, 0x5e, 0x60, 0x81, 0x69, 0x0f, 0xc8, 0xd9, 0xbc, 0xd2, 0xf4, 0x0d, 0x79, 0xd0, 0x07,
	0x70, 0xf8, 0x9b, 0x32, 0x6b, 0x9a, 0xf9, 0xbd, 0x06, 0xbb, 0xd9, 0x3e, 0xc6, 0x5f, 0x3c, 0x07,
	0xc9, 0x7b, 0x38, 0xa8, 0xcc, 0x2b, 0xa1, 0x6b, 0x97, 0xbe, 0x65, 0xd0, 0x8d, 0xa7, 0x7f, 0xd4,
	0xe4, 0x03, 0x3f, 0x83, 0xc3, 0x0d, 0x23, 0x44, 0x4e, 0xd6, 0x7a, 0xbf, 0x75, 0x96, 0x8d, 0x67,
	0x7f, 0x51, 0xe5, 0x1e, 0x6f, 0xa0, 0xb3, 0x56, 0x29, 0x79, 0x52, 0x46, 0x55, 0xaf, 0xca, 0x38,
	0xde, 0xc2, 0x66, 0x67, 0x9d, 0xef, 0xbf, 0xbb, 0x5f, 0xf0, 0xf2, 0x9f, 0x3b, 0x6b, 0xca, 0xcf,
	0xcb, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xff, 0x1f, 0x5d, 0xa2, 0x98, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentServiceClient interface {
	CurrentDeployment(ctx context.Context, in *CurrentDeploymentRequest, opts ...grpc.CallOption) (*CurrentDeploymentResponse, error)
	SetDeploymentStatus(ctx context.Context, in *SetDeploymentStatusRequest, opts ...grpc.CallOption) (*SetDeploymentStatusResponse, error)
	AgentStatus(ctx context.Context, in *AgentStatusRequest, opts ...grpc.CallOption) (*AgentStatusResponse, error)
}

type agentServiceClient struct {
	cc *grpc.ClientConn
}

func NewAgentServiceClient(cc *grpc.ClientConn) AgentServiceClient {
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

func (c *agentServiceClient) AgentStatus(ctx context.Context, in *AgentStatusRequest, opts ...grpc.CallOption) (*AgentStatusResponse, error) {
	out := new(AgentStatusResponse)
	err := c.cc.Invoke(ctx, "/signalcd.AgentService/AgentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
type AgentServiceServer interface {
	CurrentDeployment(context.Context, *CurrentDeploymentRequest) (*CurrentDeploymentResponse, error)
	SetDeploymentStatus(context.Context, *SetDeploymentStatusRequest) (*SetDeploymentStatusResponse, error)
	AgentStatus(context.Context, *AgentStatusRequest) (*AgentStatusResponse, error)
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
func (*UnimplementedAgentServiceServer) AgentStatus(ctx context.Context, req *AgentStatusRequest) (*AgentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AgentStatus not implemented")
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

func _AgentService_AgentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).AgentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signalcd.AgentService/AgentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).AgentStatus(ctx, req.(*AgentStatusRequest))
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
			MethodName: "AgentStatus",
			Handler:    _AgentService_AgentStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "signalcd/proto/agent.proto",
}
