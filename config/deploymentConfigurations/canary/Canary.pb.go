// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.14.0
// source: proto/deploymentConfigurations/canary/Canary.proto

package canary

import (
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

// Configuration for the clouddriver microservice.
type Canary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled               bool                   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	ServiceIntegrations   []*ServiceIntegrations `protobuf:"bytes,2,rep,name=serviceIntegrations,proto3" json:"serviceIntegrations,omitempty"`
	ReduxLogger           bool                   `protobuf:"varint,3,opt,name=reduxLogger,proto3" json:"reduxLogger,omitempty"`
	DefaultJudge          string                 `protobuf:"bytes,4,opt,name=defaultJudge,proto3" json:"defaultJudge,omitempty"`
	StagesEnabled         bool                   `protobuf:"varint,5,opt,name=stagesEnabled,proto3" json:"stagesEnabled,omitempty"`
	TemplatesEnabled      bool                   `protobuf:"varint,6,opt,name=templatesEnabled,proto3" json:"templatesEnabled,omitempty"`
	ShowAllConfigsEnabled bool                   `protobuf:"varint,7,opt,name=showAllConfigsEnabled,proto3" json:"showAllConfigsEnabled,omitempty"`
}

func (x *Canary) Reset() {
	*x = Canary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_canary_Canary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Canary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Canary) ProtoMessage() {}

func (x *Canary) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_canary_Canary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Canary.ProtoReflect.Descriptor instead.
func (*Canary) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_canary_Canary_proto_rawDescGZIP(), []int{0}
}

func (x *Canary) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Canary) GetServiceIntegrations() []*ServiceIntegrations {
	if x != nil {
		return x.ServiceIntegrations
	}
	return nil
}

func (x *Canary) GetReduxLogger() bool {
	if x != nil {
		return x.ReduxLogger
	}
	return false
}

func (x *Canary) GetDefaultJudge() string {
	if x != nil {
		return x.DefaultJudge
	}
	return ""
}

func (x *Canary) GetStagesEnabled() bool {
	if x != nil {
		return x.StagesEnabled
	}
	return false
}

func (x *Canary) GetTemplatesEnabled() bool {
	if x != nil {
		return x.TemplatesEnabled
	}
	return false
}

func (x *Canary) GetShowAllConfigsEnabled() bool {
	if x != nil {
		return x.ShowAllConfigsEnabled
	}
	return false
}

type ServiceIntegrations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Enabled            bool              `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Accounts           []*CanaryAccounts `protobuf:"bytes,3,rep,name=accounts,proto3" json:"accounts,omitempty"`
	GcsEnabled         bool              `protobuf:"varint,4,opt,name=gcsEnabled,proto3" json:"gcsEnabled,omitempty"`
	StackdriverEnabled bool              `protobuf:"varint,5,opt,name=stackdriverEnabled,proto3" json:"stackdriverEnabled,omitempty"`
	S3Enabled          bool              `protobuf:"varint,6,opt,name=s3Enabled,proto3" json:"s3Enabled,omitempty"`
}

func (x *ServiceIntegrations) Reset() {
	*x = ServiceIntegrations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_canary_Canary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceIntegrations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceIntegrations) ProtoMessage() {}

func (x *ServiceIntegrations) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_canary_Canary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceIntegrations.ProtoReflect.Descriptor instead.
func (*ServiceIntegrations) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_canary_Canary_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceIntegrations) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceIntegrations) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *ServiceIntegrations) GetAccounts() []*CanaryAccounts {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *ServiceIntegrations) GetGcsEnabled() bool {
	if x != nil {
		return x.GcsEnabled
	}
	return false
}

func (x *ServiceIntegrations) GetStackdriverEnabled() bool {
	if x != nil {
		return x.StackdriverEnabled
	}
	return false
}

func (x *ServiceIntegrations) GetS3Enabled() bool {
	if x != nil {
		return x.S3Enabled
	}
	return false
}

type CanaryAccounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Project              string    `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	JsonPath             string    `protobuf:"bytes,3,opt,name=jsonPath,proto3" json:"jsonPath,omitempty"`
	Bucket               string    `protobuf:"bytes,4,opt,name=bucket,proto3" json:"bucket,omitempty"`
	BucketLocation       string    `protobuf:"bytes,5,opt,name=bucketLocation,proto3" json:"bucketLocation,omitempty"`
	RootFolder           string    `protobuf:"bytes,6,opt,name=rootFolder,proto3" json:"rootFolder,omitempty"`
	SupportedTypes       []string  `protobuf:"bytes,7,rep,name=supportedTypes,proto3" json:"supportedTypes,omitempty"`
	Endpoint             *Endpoint `protobuf:"bytes,8,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Username             string    `protobuf:"bytes,9,opt,name=username,proto3" json:"username,omitempty"`
	Password             string    `protobuf:"bytes,10,opt,name=password,proto3" json:"password,omitempty"`
	UsernamePasswordFile string    `protobuf:"bytes,11,opt,name=usernamePasswordFile,proto3" json:"usernamePasswordFile,omitempty"`
	ApiKey               string    `protobuf:"bytes,12,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	ApplicationKey       string    `protobuf:"bytes,13,opt,name=applicationKey,proto3" json:"applicationKey,omitempty"`
	AccessToken          string    `protobuf:"bytes,14,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	DefaultScopeKey      string    `protobuf:"bytes,15,opt,name=defaultScopeKey,proto3" json:"defaultScopeKey,omitempty"`
	DefaultLocationKey   string    `protobuf:"bytes,16,opt,name=defaultLocationKey,proto3" json:"defaultLocationKey,omitempty"`
	Region               string    `protobuf:"bytes,17,opt,name=region,proto3" json:"region,omitempty"`
	ProfileName          string    `protobuf:"bytes,18,opt,name=profileName,proto3" json:"profileName,omitempty"`
	// string Awsendpoint = 19;
	AccessKeyId     string `protobuf:"bytes,20,opt,name=accessKeyId,proto3" json:"accessKeyId,omitempty"`
	SecretAccessKey string `protobuf:"bytes,21,opt,name=secretAccessKey,proto3" json:"secretAccessKey,omitempty"`
}

func (x *CanaryAccounts) Reset() {
	*x = CanaryAccounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_canary_Canary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanaryAccounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanaryAccounts) ProtoMessage() {}

func (x *CanaryAccounts) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_canary_Canary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanaryAccounts.ProtoReflect.Descriptor instead.
func (*CanaryAccounts) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_canary_Canary_proto_rawDescGZIP(), []int{2}
}

func (x *CanaryAccounts) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CanaryAccounts) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *CanaryAccounts) GetJsonPath() string {
	if x != nil {
		return x.JsonPath
	}
	return ""
}

func (x *CanaryAccounts) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *CanaryAccounts) GetBucketLocation() string {
	if x != nil {
		return x.BucketLocation
	}
	return ""
}

func (x *CanaryAccounts) GetRootFolder() string {
	if x != nil {
		return x.RootFolder
	}
	return ""
}

func (x *CanaryAccounts) GetSupportedTypes() []string {
	if x != nil {
		return x.SupportedTypes
	}
	return nil
}

func (x *CanaryAccounts) GetEndpoint() *Endpoint {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

func (x *CanaryAccounts) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CanaryAccounts) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CanaryAccounts) GetUsernamePasswordFile() string {
	if x != nil {
		return x.UsernamePasswordFile
	}
	return ""
}

func (x *CanaryAccounts) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *CanaryAccounts) GetApplicationKey() string {
	if x != nil {
		return x.ApplicationKey
	}
	return ""
}

func (x *CanaryAccounts) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *CanaryAccounts) GetDefaultScopeKey() string {
	if x != nil {
		return x.DefaultScopeKey
	}
	return ""
}

func (x *CanaryAccounts) GetDefaultLocationKey() string {
	if x != nil {
		return x.DefaultLocationKey
	}
	return ""
}

func (x *CanaryAccounts) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *CanaryAccounts) GetProfileName() string {
	if x != nil {
		return x.ProfileName
	}
	return ""
}

func (x *CanaryAccounts) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *CanaryAccounts) GetSecretAccessKey() string {
	if x != nil {
		return x.SecretAccessKey
	}
	return ""
}

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseUrl string `protobuf:"bytes,1,opt,name=baseUrl,proto3" json:"baseUrl,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_canary_Canary_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_canary_Canary_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_canary_Canary_proto_rawDescGZIP(), []int{3}
}

func (x *Endpoint) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

var File_proto_deploymentConfigurations_canary_Canary_proto protoreflect.FileDescriptor

var file_proto_deploymentConfigurations_canary_Canary_proto_rawDesc = []byte{
	0x0a, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x43, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61,
	0x72, 0x79, 0x22, 0xc5, 0x02, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x53, 0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e,
	0x61, 0x72, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x72, 0x65, 0x64, 0x75, 0x78, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x75, 0x78, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x22,
	0x0a, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4a, 0x75, 0x64,
	0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x10, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x15, 0x73, 0x68, 0x6f, 0x77, 0x41, 0x6c, 0x6c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x15, 0x73, 0x68, 0x6f, 0x77, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0xeb, 0x01, 0x0a, 0x13, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x38, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72,
	0x79, 0x2e, 0x43, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x63,
	0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x67, 0x63, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x33,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73,
	0x33, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0xc4, 0x05, 0x0a, 0x0e, 0x43, 0x61, 0x6e,
	0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x73, 0x6f,
	0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f,
	0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x26, 0x0a,
	0x0e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x74, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x74, 0x46,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x32, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x32, 0x0a, 0x14, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x28, 0x0a, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x4b,
	0x65, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x12, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x49, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x22,
	0x24, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61,
	0x73, 0x65, 0x55, 0x72, 0x6c, 0x42, 0x28, 0x5a, 0x26, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_deploymentConfigurations_canary_Canary_proto_rawDescOnce sync.Once
	file_proto_deploymentConfigurations_canary_Canary_proto_rawDescData = file_proto_deploymentConfigurations_canary_Canary_proto_rawDesc
)

func file_proto_deploymentConfigurations_canary_Canary_proto_rawDescGZIP() []byte {
	file_proto_deploymentConfigurations_canary_Canary_proto_rawDescOnce.Do(func() {
		file_proto_deploymentConfigurations_canary_Canary_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_deploymentConfigurations_canary_Canary_proto_rawDescData)
	})
	return file_proto_deploymentConfigurations_canary_Canary_proto_rawDescData
}

var file_proto_deploymentConfigurations_canary_Canary_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_deploymentConfigurations_canary_Canary_proto_goTypes = []interface{}{
	(*Canary)(nil),              // 0: proto.canary.Canary
	(*ServiceIntegrations)(nil), // 1: proto.canary.serviceIntegrations
	(*CanaryAccounts)(nil),      // 2: proto.canary.CanaryAccounts
	(*Endpoint)(nil),            // 3: proto.canary.Endpoint
}
var file_proto_deploymentConfigurations_canary_Canary_proto_depIdxs = []int32{
	1, // 0: proto.canary.Canary.serviceIntegrations:type_name -> proto.canary.serviceIntegrations
	2, // 1: proto.canary.serviceIntegrations.accounts:type_name -> proto.canary.CanaryAccounts
	3, // 2: proto.canary.CanaryAccounts.endpoint:type_name -> proto.canary.Endpoint
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_deploymentConfigurations_canary_Canary_proto_init() }
func file_proto_deploymentConfigurations_canary_Canary_proto_init() {
	if File_proto_deploymentConfigurations_canary_Canary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_deploymentConfigurations_canary_Canary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Canary); i {
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
		file_proto_deploymentConfigurations_canary_Canary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceIntegrations); i {
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
		file_proto_deploymentConfigurations_canary_Canary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CanaryAccounts); i {
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
		file_proto_deploymentConfigurations_canary_Canary_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
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
			RawDescriptor: file_proto_deploymentConfigurations_canary_Canary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_deploymentConfigurations_canary_Canary_proto_goTypes,
		DependencyIndexes: file_proto_deploymentConfigurations_canary_Canary_proto_depIdxs,
		MessageInfos:      file_proto_deploymentConfigurations_canary_Canary_proto_msgTypes,
	}.Build()
	File_proto_deploymentConfigurations_canary_Canary_proto = out.File
	file_proto_deploymentConfigurations_canary_Canary_proto_rawDesc = nil
	file_proto_deploymentConfigurations_canary_Canary_proto_goTypes = nil
	file_proto_deploymentConfigurations_canary_Canary_proto_depIdxs = nil
}
