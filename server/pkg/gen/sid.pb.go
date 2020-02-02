// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sid.proto

package sid

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type HealthStatus_Status int32

const (
	HealthStatus_INACTIVE HealthStatus_Status = 0
	HealthStatus_READY    HealthStatus_Status = 1
	HealthStatus_WORKING  HealthStatus_Status = 2
	HealthStatus_LEAVING  HealthStatus_Status = 3
)

var HealthStatus_Status_name = map[int32]string{
	0: "INACTIVE",
	1: "READY",
	2: "WORKING",
	3: "LEAVING",
}

var HealthStatus_Status_value = map[string]int32{
	"INACTIVE": 0,
	"READY":    1,
	"WORKING":  2,
	"LEAVING":  3,
}

func (x HealthStatus_Status) String() string {
	return proto.EnumName(HealthStatus_Status_name, int32(x))
}

func (HealthStatus_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_923cdea5d0e24e4a, []int{2, 0}
}

type Job_JobStatus int32

const (
	Job_QUEUED    Job_JobStatus = 0
	Job_BUILDING  Job_JobStatus = 1
	Job_ABANDONED Job_JobStatus = 2
	Job_COMPLETED Job_JobStatus = 3
	Job_FAILED    Job_JobStatus = 5
)

var Job_JobStatus_name = map[int32]string{
	0: "QUEUED",
	1: "BUILDING",
	2: "ABANDONED",
	3: "COMPLETED",
	5: "FAILED",
}

var Job_JobStatus_value = map[string]int32{
	"QUEUED":    0,
	"BUILDING":  1,
	"ABANDONED": 2,
	"COMPLETED": 3,
	"FAILED":    5,
}

func (x Job_JobStatus) String() string {
	return proto.EnumName(Job_JobStatus_name, int32(x))
}

func (Job_JobStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_923cdea5d0e24e4a, []int{3, 0}
}

type JobRunEvent_EventType int32

const (
	JobRunEvent_RUN_LOG JobRunEvent_EventType = 0
	JobRunEvent_ERROR   JobRunEvent_EventType = 1
)

var JobRunEvent_EventType_name = map[int32]string{
	0: "RUN_LOG",
	1: "ERROR",
}

var JobRunEvent_EventType_value = map[string]int32{
	"RUN_LOG": 0,
	"ERROR":   1,
}

func (x JobRunEvent_EventType) String() string {
	return proto.EnumName(JobRunEvent_EventType_name, int32(x))
}

func (JobRunEvent_EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_923cdea5d0e24e4a, []int{4, 0}
}

type Token struct {
	Token                string               `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ExpiresAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_923cdea5d0e24e4a, []int{0}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetExpiresAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

// Login Request
type LoginRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	NewPassword          string   `protobuf:"bytes,3,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_923cdea5d0e24e4a, []int{1}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

// HealthStatus
type HealthStatus struct {
	Status               HealthStatus_Status  `protobuf:"varint,1,opt,name=status,proto3,enum=sid.HealthStatus_Status" json:"status,omitempty"`
	StatusAt             *timestamp.Timestamp `protobuf:"bytes,2,opt,name=status_at,json=statusAt,proto3" json:"status_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HealthStatus) Reset()         { *m = HealthStatus{} }
func (m *HealthStatus) String() string { return proto.CompactTextString(m) }
func (*HealthStatus) ProtoMessage()    {}
func (*HealthStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_923cdea5d0e24e4a, []int{2}
}

func (m *HealthStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthStatus.Unmarshal(m, b)
}
func (m *HealthStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthStatus.Marshal(b, m, deterministic)
}
func (m *HealthStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthStatus.Merge(m, src)
}
func (m *HealthStatus) XXX_Size() int {
	return xxx_messageInfo_HealthStatus.Size(m)
}
func (m *HealthStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HealthStatus proto.InternalMessageInfo

func (m *HealthStatus) GetStatus() HealthStatus_Status {
	if m != nil {
		return m.Status
	}
	return HealthStatus_INACTIVE
}

func (m *HealthStatus) GetStatusAt() *timestamp.Timestamp {
	if m != nil {
		return m.StatusAt
	}
	return nil
}

type Job struct {
	RepoName             string               `protobuf:"bytes,1,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	RepoSshUrl           string               `protobuf:"bytes,2,opt,name=repo_ssh_url,json=repoSshUrl,proto3" json:"repo_ssh_url,omitempty"`
	CommitHexsha         string               `protobuf:"bytes,3,opt,name=commit_hexsha,json=commitHexsha,proto3" json:"commit_hexsha,omitempty"`
	JobStatus            Job_JobStatus        `protobuf:"varint,4,opt,name=job_status,json=jobStatus,proto3,enum=sid.Job_JobStatus" json:"job_status,omitempty"`
	StatusAt             *timestamp.Timestamp `protobuf:"bytes,5,opt,name=status_at,json=statusAt,proto3" json:"status_at,omitempty"`
	JobUuid              string               `protobuf:"bytes,6,opt,name=job_uuid,json=jobUuid,proto3" json:"job_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_923cdea5d0e24e4a, []int{3}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

func (m *Job) GetRepoSshUrl() string {
	if m != nil {
		return m.RepoSshUrl
	}
	return ""
}

func (m *Job) GetCommitHexsha() string {
	if m != nil {
		return m.CommitHexsha
	}
	return ""
}

func (m *Job) GetJobStatus() Job_JobStatus {
	if m != nil {
		return m.JobStatus
	}
	return Job_QUEUED
}

func (m *Job) GetStatusAt() *timestamp.Timestamp {
	if m != nil {
		return m.StatusAt
	}
	return nil
}

func (m *Job) GetJobUuid() string {
	if m != nil {
		return m.JobUuid
	}
	return ""
}

type JobRunEvent struct {
	Type                 JobRunEvent_EventType `protobuf:"varint,1,opt,name=type,proto3,enum=sid.JobRunEvent_EventType" json:"type,omitempty"`
	Content              string                `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	EventAt              *timestamp.Timestamp  `protobuf:"bytes,3,opt,name=event_at,json=eventAt,proto3" json:"event_at,omitempty"`
	Job                  *Job                  `protobuf:"bytes,4,opt,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *JobRunEvent) Reset()         { *m = JobRunEvent{} }
func (m *JobRunEvent) String() string { return proto.CompactTextString(m) }
func (*JobRunEvent) ProtoMessage()    {}
func (*JobRunEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_923cdea5d0e24e4a, []int{4}
}

func (m *JobRunEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobRunEvent.Unmarshal(m, b)
}
func (m *JobRunEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobRunEvent.Marshal(b, m, deterministic)
}
func (m *JobRunEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobRunEvent.Merge(m, src)
}
func (m *JobRunEvent) XXX_Size() int {
	return xxx_messageInfo_JobRunEvent.Size(m)
}
func (m *JobRunEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_JobRunEvent.DiscardUnknown(m)
}

var xxx_messageInfo_JobRunEvent proto.InternalMessageInfo

func (m *JobRunEvent) GetType() JobRunEvent_EventType {
	if m != nil {
		return m.Type
	}
	return JobRunEvent_RUN_LOG
}

func (m *JobRunEvent) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *JobRunEvent) GetEventAt() *timestamp.Timestamp {
	if m != nil {
		return m.EventAt
	}
	return nil
}

func (m *JobRunEvent) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type CheckInResponse struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckInResponse) Reset()         { *m = CheckInResponse{} }
func (m *CheckInResponse) String() string { return proto.CompactTextString(m) }
func (*CheckInResponse) ProtoMessage()    {}
func (*CheckInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_923cdea5d0e24e4a, []int{5}
}

func (m *CheckInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckInResponse.Unmarshal(m, b)
}
func (m *CheckInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckInResponse.Marshal(b, m, deterministic)
}
func (m *CheckInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckInResponse.Merge(m, src)
}
func (m *CheckInResponse) XXX_Size() int {
	return xxx_messageInfo_CheckInResponse.Size(m)
}
func (m *CheckInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckInResponse proto.InternalMessageInfo

func (m *CheckInResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type Repo struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SshUrl               string               `protobuf:"bytes,2,opt,name=ssh_url,json=sshUrl,proto3" json:"ssh_url,omitempty"`
	Enabled              bool                 `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	AddedBy              int32                `protobuf:"varint,4,opt,name=added_by,json=addedBy,proto3" json:"added_by,omitempty"`
	AddedAt              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=added_at,json=addedAt,proto3" json:"added_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Repo) Reset()         { *m = Repo{} }
func (m *Repo) String() string { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()    {}
func (*Repo) Descriptor() ([]byte, []int) {
	return fileDescriptor_923cdea5d0e24e4a, []int{6}
}

func (m *Repo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repo.Unmarshal(m, b)
}
func (m *Repo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repo.Marshal(b, m, deterministic)
}
func (m *Repo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repo.Merge(m, src)
}
func (m *Repo) XXX_Size() int {
	return xxx_messageInfo_Repo.Size(m)
}
func (m *Repo) XXX_DiscardUnknown() {
	xxx_messageInfo_Repo.DiscardUnknown(m)
}

var xxx_messageInfo_Repo proto.InternalMessageInfo

func (m *Repo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Repo) GetSshUrl() string {
	if m != nil {
		return m.SshUrl
	}
	return ""
}

func (m *Repo) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Repo) GetAddedBy() int32 {
	if m != nil {
		return m.AddedBy
	}
	return 0
}

func (m *Repo) GetAddedAt() *timestamp.Timestamp {
	if m != nil {
		return m.AddedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("sid.HealthStatus_Status", HealthStatus_Status_name, HealthStatus_Status_value)
	proto.RegisterEnum("sid.Job_JobStatus", Job_JobStatus_name, Job_JobStatus_value)
	proto.RegisterEnum("sid.JobRunEvent_EventType", JobRunEvent_EventType_name, JobRunEvent_EventType_value)
	proto.RegisterType((*Token)(nil), "sid.Token")
	proto.RegisterType((*LoginRequest)(nil), "sid.LoginRequest")
	proto.RegisterType((*HealthStatus)(nil), "sid.HealthStatus")
	proto.RegisterType((*Job)(nil), "sid.Job")
	proto.RegisterType((*JobRunEvent)(nil), "sid.JobRunEvent")
	proto.RegisterType((*CheckInResponse)(nil), "sid.CheckInResponse")
	proto.RegisterType((*Repo)(nil), "sid.Repo")
}

func init() { proto.RegisterFile("sid.proto", fileDescriptor_923cdea5d0e24e4a) }

var fileDescriptor_923cdea5d0e24e4a = []byte{
	// 776 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x51, 0x73, 0xdb, 0x44,
	0x10, 0xb6, 0xe2, 0x58, 0x96, 0xd6, 0x2e, 0x88, 0xa5, 0x33, 0x18, 0x31, 0x53, 0x82, 0xca, 0x0c,
	0x7e, 0xa9, 0x5a, 0xc2, 0x30, 0x0c, 0xc3, 0x93, 0x12, 0x8b, 0xd4, 0xc1, 0xd8, 0xe5, 0x62, 0x17,
	0x78, 0xd2, 0x48, 0xd1, 0x36, 0x56, 0x6a, 0xeb, 0x84, 0xee, 0x44, 0xea, 0x5f, 0xc3, 0xbf, 0xe0,
	0x9d, 0x5f, 0xc0, 0x5f, 0x62, 0xee, 0x24, 0xbb, 0x4e, 0xe9, 0x4c, 0xf2, 0x60, 0x6b, 0xf7, 0xdb,
	0x4f, 0xb7, 0xda, 0xef, 0xbe, 0x3b, 0xb0, 0x45, 0x96, 0xfa, 0x45, 0xc9, 0x25, 0xc7, 0xb6, 0xc8,
	0x52, 0xf7, 0xf3, 0x2b, 0xce, 0xaf, 0x56, 0xf4, 0x54, 0x43, 0x49, 0xf5, 0xea, 0xa9, 0xcc, 0xd6,
	0x24, 0x64, 0xbc, 0x2e, 0x6a, 0x96, 0xf7, 0x1b, 0x74, 0xe6, 0xfc, 0x35, 0xe5, 0xf8, 0x10, 0x3a,
	0x52, 0x05, 0x03, 0xe3, 0xc8, 0x18, 0xda, 0xac, 0x4e, 0xf0, 0x7b, 0x00, 0x7a, 0x53, 0x64, 0x25,
	0x89, 0x28, 0x96, 0x83, 0x83, 0x23, 0x63, 0xd8, 0x3b, 0x76, 0xfd, 0x7a, 0x51, 0x7f, 0xbb, 0xa8,
	0x3f, 0xdf, 0x2e, 0xca, 0xec, 0x86, 0x1d, 0x48, 0x6f, 0x0d, 0xfd, 0x09, 0xbf, 0xca, 0x72, 0x46,
	0x7f, 0x54, 0x24, 0x24, 0x3e, 0x02, 0xc8, 0x52, 0xca, 0x65, 0xf6, 0x2a, 0xa3, 0xb2, 0xe9, 0xb2,
	0x87, 0xa0, 0x0b, 0x56, 0x11, 0x0b, 0x71, 0xc3, 0xcb, 0x54, 0x37, 0xb2, 0xd9, 0x2e, 0xc7, 0x2f,
	0xa0, 0x9f, 0xd3, 0x4d, 0xb4, 0xab, 0xb7, 0x75, 0xbd, 0x97, 0xd3, 0xcd, 0x8b, 0x06, 0xf2, 0xfe,
	0x36, 0xa0, 0xff, 0x9c, 0xe2, 0x95, 0x5c, 0x5e, 0xc8, 0x58, 0x56, 0x02, 0x9f, 0x81, 0x29, 0x74,
	0xa4, 0x7b, 0x7d, 0x70, 0x3c, 0xf0, 0x95, 0x36, 0xfb, 0x14, 0xbf, 0x7e, 0xb0, 0x86, 0x87, 0xdf,
	0x81, 0x5d, 0x47, 0xf7, 0x9b, 0xd5, 0xaa, 0xc9, 0x81, 0xf4, 0x7e, 0x00, 0xb3, 0x69, 0xda, 0x07,
	0x6b, 0x3c, 0x0d, 0x4e, 0xe7, 0xe3, 0x97, 0xa1, 0xd3, 0x42, 0x1b, 0x3a, 0x2c, 0x0c, 0x46, 0xbf,
	0x3b, 0x06, 0xf6, 0xa0, 0xfb, 0xeb, 0x8c, 0xfd, 0x34, 0x9e, 0x9e, 0x39, 0x07, 0x2a, 0x99, 0x84,
	0xc1, 0x4b, 0x95, 0xb4, 0xbd, 0x7f, 0x0e, 0xa0, 0x7d, 0xce, 0x13, 0xfc, 0x0c, 0xec, 0x92, 0x0a,
	0x1e, 0xe5, 0xf1, 0x9a, 0x1a, 0x79, 0x2c, 0x05, 0x4c, 0xe3, 0x35, 0xe1, 0x11, 0xf4, 0x75, 0x51,
	0x88, 0x65, 0x54, 0x95, 0xab, 0x46, 0x20, 0x50, 0xd8, 0x85, 0x58, 0x2e, 0xca, 0x15, 0x3e, 0x86,
	0x07, 0x97, 0x7c, 0xbd, 0xce, 0x64, 0xb4, 0xa4, 0x37, 0x62, 0x19, 0x37, 0x1a, 0xf5, 0x6b, 0xf0,
	0xb9, 0xc6, 0xf0, 0x6b, 0x80, 0x6b, 0x9e, 0x44, 0x8d, 0x2e, 0x87, 0x5a, 0x17, 0xd4, 0xba, 0x9c,
	0xf3, 0x44, 0xfd, 0x1a, 0x45, 0xec, 0xeb, 0x6d, 0x78, 0x5b, 0x94, 0xce, 0xfd, 0x45, 0xc1, 0x4f,
	0xc1, 0x52, 0xbd, 0xaa, 0x2a, 0x4b, 0x07, 0xa6, 0xfe, 0x96, 0xee, 0x35, 0x4f, 0x16, 0x55, 0x96,
	0x7a, 0x33, 0xb0, 0x77, 0xbd, 0x10, 0xc0, 0xfc, 0x65, 0x11, 0x2e, 0xc2, 0x91, 0xd3, 0x52, 0xf2,
	0x9d, 0x2c, 0xc6, 0x93, 0x91, 0x52, 0xc6, 0xc0, 0x07, 0x60, 0x07, 0x27, 0xc1, 0x74, 0x34, 0x9b,
	0x86, 0x23, 0xe7, 0x40, 0xa5, 0xa7, 0xb3, 0x9f, 0x5f, 0x4c, 0xc2, 0x79, 0x38, 0x72, 0xda, 0xea,
	0xbd, 0x1f, 0x83, 0xf1, 0x24, 0x1c, 0x39, 0x1d, 0xef, 0x5f, 0x03, 0x7a, 0xe7, 0x3c, 0x61, 0x55,
	0x1e, 0xfe, 0x49, 0xb9, 0x44, 0x1f, 0x0e, 0xe5, 0xa6, 0xa0, 0x66, 0xe7, 0xdd, 0xed, 0x84, 0xdb,
	0xba, 0xaf, 0xff, 0xe7, 0x9b, 0x82, 0x98, 0xe6, 0xe1, 0x00, 0xba, 0x97, 0x3c, 0x97, 0x94, 0xcb,
	0x46, 0xd9, 0x6d, 0x8a, 0xdf, 0x82, 0x45, 0x8a, 0xac, 0xa6, 0x6f, 0xdf, 0x39, 0x7d, 0x57, 0x73,
	0x03, 0x89, 0x2e, 0xb4, 0xaf, 0x79, 0xa2, 0x15, 0xee, 0x1d, 0x5b, 0xbb, 0xfe, 0x0a, 0xf4, 0x1e,
	0x83, 0xbd, 0xeb, 0xaf, 0xac, 0xc0, 0x16, 0xd3, 0x68, 0x32, 0x3b, 0xab, 0xfd, 0x12, 0x32, 0x36,
	0x63, 0x8e, 0xe1, 0x3d, 0x81, 0x0f, 0x4f, 0x97, 0x74, 0xf9, 0x7a, 0x9c, 0x33, 0x12, 0x05, 0xcf,
	0x05, 0xa9, 0x03, 0x52, 0x36, 0xf1, 0x5b, 0x7f, 0xd4, 0xb9, 0xf7, 0x97, 0x01, 0x87, 0x8c, 0x0a,
	0x8e, 0x08, 0x87, 0x7b, 0x06, 0xd2, 0x31, 0x7e, 0x02, 0xdd, 0xdb, 0xbe, 0x31, 0x45, 0xed, 0x99,
	0x01, 0x74, 0x29, 0x8f, 0x93, 0x15, 0xd5, 0x27, 0xca, 0x62, 0xdb, 0x54, 0x6d, 0x5e, 0x9c, 0xa6,
	0x94, 0x46, 0xc9, 0x46, 0x0f, 0xd1, 0x61, 0x5d, 0x9d, 0x9f, 0x6c, 0x94, 0x22, 0x75, 0xe9, 0x5e,
	0x7e, 0xa8, 0x5f, 0x0b, 0xe4, 0xb1, 0xb2, 0xf9, 0x45, 0x96, 0xe2, 0x57, 0x60, 0x9e, 0x91, 0x54,
	0x86, 0xff, 0xe8, 0x7f, 0x07, 0xd2, 0xdd, 0x29, 0xe5, 0xb5, 0xf0, 0x11, 0x98, 0x41, 0x9a, 0x2a,
	0xe2, 0x0e, 0xbd, 0x55, 0x1f, 0x42, 0x47, 0xdf, 0x2f, 0xcd, 0x3a, 0xfb, 0x77, 0x8d, 0x0b, 0x1a,
	0xd2, 0x17, 0x9b, 0xd7, 0xc2, 0x27, 0x00, 0xa7, 0xcb, 0x38, 0xbf, 0x22, 0x75, 0x59, 0xdc, 0x4d,
	0xff, 0x12, 0xac, 0x33, 0x92, 0x4a, 0x4d, 0x81, 0xb6, 0xae, 0xa8, 0xd8, 0x7d, 0x1b, 0x7a, 0xad,
	0x67, 0x06, 0x9e, 0xc0, 0xc7, 0xfb, 0x9f, 0xde, 0x6c, 0xd6, 0xfb, 0x86, 0x7a, 0xa8, 0xa1, 0x77,
	0x76, 0xd3, 0x6b, 0x0d, 0x0d, 0xf4, 0xa1, 0xcf, 0xe8, 0x92, 0x97, 0x69, 0xed, 0x4d, 0x74, 0xde,
	0x35, 0xea, 0xfe, 0xc0, 0x43, 0x23, 0x31, 0xb5, 0xc0, 0xdf, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff,
	0xaa, 0x71, 0x65, 0xef, 0xe6, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SidClient is the client API for Sid service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SidClient interface {
	// Obtains a job from the queue
	GetJob(ctx context.Context, in *HealthStatus, opts ...grpc.CallOption) (*Job, error)
	// Adds a job to the queue
	AddJob(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Job, error)
	// Log in
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Token, error)
	ChangePass(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Token, error)
	// server to client stream of repos matching repo filter
	GetRepos(ctx context.Context, in *Repo, opts ...grpc.CallOption) (Sid_GetReposClient, error)
	// A client-to-server streaming RPC.
	//
	// Streams health status of the client as it changes.
	HealthStatusCheckIn(ctx context.Context, opts ...grpc.CallOption) (Sid_HealthStatusCheckInClient, error)
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of JobRunEvents on a job being run, returning a
	// job when done.
	RecordJobRun(ctx context.Context, opts ...grpc.CallOption) (Sid_RecordJobRunClient, error)
}

type sidClient struct {
	cc *grpc.ClientConn
}

func NewSidClient(cc *grpc.ClientConn) SidClient {
	return &sidClient{cc}
}

func (c *sidClient) GetJob(ctx context.Context, in *HealthStatus, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, "/sid.Sid/GetJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sidClient) AddJob(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, "/sid.Sid/AddJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sidClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/sid.Sid/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sidClient) ChangePass(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/sid.Sid/ChangePass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sidClient) GetRepos(ctx context.Context, in *Repo, opts ...grpc.CallOption) (Sid_GetReposClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sid_serviceDesc.Streams[0], "/sid.Sid/GetRepos", opts...)
	if err != nil {
		return nil, err
	}
	x := &sidGetReposClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Sid_GetReposClient interface {
	Recv() (*Repo, error)
	grpc.ClientStream
}

type sidGetReposClient struct {
	grpc.ClientStream
}

func (x *sidGetReposClient) Recv() (*Repo, error) {
	m := new(Repo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sidClient) HealthStatusCheckIn(ctx context.Context, opts ...grpc.CallOption) (Sid_HealthStatusCheckInClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sid_serviceDesc.Streams[1], "/sid.Sid/HealthStatusCheckIn", opts...)
	if err != nil {
		return nil, err
	}
	x := &sidHealthStatusCheckInClient{stream}
	return x, nil
}

type Sid_HealthStatusCheckInClient interface {
	Send(*HealthStatus) error
	CloseAndRecv() (*CheckInResponse, error)
	grpc.ClientStream
}

type sidHealthStatusCheckInClient struct {
	grpc.ClientStream
}

func (x *sidHealthStatusCheckInClient) Send(m *HealthStatus) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sidHealthStatusCheckInClient) CloseAndRecv() (*CheckInResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CheckInResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sidClient) RecordJobRun(ctx context.Context, opts ...grpc.CallOption) (Sid_RecordJobRunClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sid_serviceDesc.Streams[2], "/sid.Sid/RecordJobRun", opts...)
	if err != nil {
		return nil, err
	}
	x := &sidRecordJobRunClient{stream}
	return x, nil
}

type Sid_RecordJobRunClient interface {
	Send(*JobRunEvent) error
	CloseAndRecv() (*Job, error)
	grpc.ClientStream
}

type sidRecordJobRunClient struct {
	grpc.ClientStream
}

func (x *sidRecordJobRunClient) Send(m *JobRunEvent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sidRecordJobRunClient) CloseAndRecv() (*Job, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Job)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SidServer is the server API for Sid service.
type SidServer interface {
	// Obtains a job from the queue
	GetJob(context.Context, *HealthStatus) (*Job, error)
	// Adds a job to the queue
	AddJob(context.Context, *Job) (*Job, error)
	// Log in
	Login(context.Context, *LoginRequest) (*Token, error)
	ChangePass(context.Context, *LoginRequest) (*Token, error)
	// server to client stream of repos matching repo filter
	GetRepos(*Repo, Sid_GetReposServer) error
	// A client-to-server streaming RPC.
	//
	// Streams health status of the client as it changes.
	HealthStatusCheckIn(Sid_HealthStatusCheckInServer) error
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of JobRunEvents on a job being run, returning a
	// job when done.
	RecordJobRun(Sid_RecordJobRunServer) error
}

// UnimplementedSidServer can be embedded to have forward compatible implementations.
type UnimplementedSidServer struct {
}

func (*UnimplementedSidServer) GetJob(ctx context.Context, req *HealthStatus) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}
func (*UnimplementedSidServer) AddJob(ctx context.Context, req *Job) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddJob not implemented")
}
func (*UnimplementedSidServer) Login(ctx context.Context, req *LoginRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedSidServer) ChangePass(ctx context.Context, req *LoginRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePass not implemented")
}
func (*UnimplementedSidServer) GetRepos(req *Repo, srv Sid_GetReposServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRepos not implemented")
}
func (*UnimplementedSidServer) HealthStatusCheckIn(srv Sid_HealthStatusCheckInServer) error {
	return status.Errorf(codes.Unimplemented, "method HealthStatusCheckIn not implemented")
}
func (*UnimplementedSidServer) RecordJobRun(srv Sid_RecordJobRunServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordJobRun not implemented")
}

func RegisterSidServer(s *grpc.Server, srv SidServer) {
	s.RegisterService(&_Sid_serviceDesc, srv)
}

func _Sid_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SidServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sid.Sid/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SidServer).GetJob(ctx, req.(*HealthStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sid_AddJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SidServer).AddJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sid.Sid/AddJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SidServer).AddJob(ctx, req.(*Job))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sid_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SidServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sid.Sid/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SidServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sid_ChangePass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SidServer).ChangePass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sid.Sid/ChangePass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SidServer).ChangePass(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sid_GetRepos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Repo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SidServer).GetRepos(m, &sidGetReposServer{stream})
}

type Sid_GetReposServer interface {
	Send(*Repo) error
	grpc.ServerStream
}

type sidGetReposServer struct {
	grpc.ServerStream
}

func (x *sidGetReposServer) Send(m *Repo) error {
	return x.ServerStream.SendMsg(m)
}

func _Sid_HealthStatusCheckIn_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SidServer).HealthStatusCheckIn(&sidHealthStatusCheckInServer{stream})
}

type Sid_HealthStatusCheckInServer interface {
	SendAndClose(*CheckInResponse) error
	Recv() (*HealthStatus, error)
	grpc.ServerStream
}

type sidHealthStatusCheckInServer struct {
	grpc.ServerStream
}

func (x *sidHealthStatusCheckInServer) SendAndClose(m *CheckInResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sidHealthStatusCheckInServer) Recv() (*HealthStatus, error) {
	m := new(HealthStatus)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Sid_RecordJobRun_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SidServer).RecordJobRun(&sidRecordJobRunServer{stream})
}

type Sid_RecordJobRunServer interface {
	SendAndClose(*Job) error
	Recv() (*JobRunEvent, error)
	grpc.ServerStream
}

type sidRecordJobRunServer struct {
	grpc.ServerStream
}

func (x *sidRecordJobRunServer) SendAndClose(m *Job) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sidRecordJobRunServer) Recv() (*JobRunEvent, error) {
	m := new(JobRunEvent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Sid_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sid.Sid",
	HandlerType: (*SidServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJob",
			Handler:    _Sid_GetJob_Handler,
		},
		{
			MethodName: "AddJob",
			Handler:    _Sid_AddJob_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Sid_Login_Handler,
		},
		{
			MethodName: "ChangePass",
			Handler:    _Sid_ChangePass_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRepos",
			Handler:       _Sid_GetRepos_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HealthStatusCheckIn",
			Handler:       _Sid_HealthStatusCheckIn_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RecordJobRun",
			Handler:       _Sid_RecordJobRun_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "sid.proto",
}
