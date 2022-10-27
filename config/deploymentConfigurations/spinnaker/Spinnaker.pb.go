// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.14.0
// source: proto/deploymentConfigurations/spinnaker/Spinnaker.proto

package spinnaker

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

// import "proto/deploymentConfigurations/persistentStorage/S3.proto";
// import "proto/deploymentConfigurations/persistentStorage/Azs.proto";
// import "proto/deploymentConfigurations/persistentStorage/Oracle.proto";
// import "proto/deploymentConfigurations/persistentStorage/Gcs.proto";
// Configuration for the clouddriver microservice.
type Spinnaker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extensibility *Extensibility `protobuf:"bytes,1,opt,name=extensibility,proto3" json:"extensibility,omitempty"`
}

func (x *Spinnaker) Reset() {
	*x = Spinnaker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spinnaker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spinnaker) ProtoMessage() {}

func (x *Spinnaker) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spinnaker.ProtoReflect.Descriptor instead.
func (*Spinnaker) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDescGZIP(), []int{0}
}

func (x *Spinnaker) GetExtensibility() *Extensibility {
	if x != nil {
		return x.Extensibility
	}
	return nil
}

type Extensibility struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// map<string, Plugins> plugins = 1;
	Plugins      *Test         `protobuf:"bytes,1,opt,name=plugins,proto3" json:"plugins,omitempty"`
	Repositories *Repositories `protobuf:"bytes,2,opt,name=repositories,proto3" json:"repositories,omitempty"`
}

func (x *Extensibility) Reset() {
	*x = Extensibility{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extensibility) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extensibility) ProtoMessage() {}

func (x *Extensibility) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extensibility.ProtoReflect.Descriptor instead.
func (*Extensibility) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDescGZIP(), []int{1}
}

func (x *Extensibility) GetPlugins() *Test {
	if x != nil {
		return x.Plugins
	}
	return nil
}

func (x *Extensibility) GetRepositories() *Repositories {
	if x != nil {
		return x.Repositories
	}
	return nil
}

type Plugins struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string name = 1;
	Id         string        `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Enabled    bool          `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Version    string        `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Config     *Config       `protobuf:"bytes,5,opt,name=config,proto3" json:"config,omitempty"`
	Extensions []*Extensions `protobuf:"bytes,6,rep,name=extensions,proto3" json:"extensions,omitempty"`
}

func (x *Plugins) Reset() {
	*x = Plugins{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugins) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugins) ProtoMessage() {}

func (x *Plugins) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugins.ProtoReflect.Descriptor instead.
func (*Plugins) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDescGZIP(), []int{2}
}

func (x *Plugins) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Plugins) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Plugins) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Plugins) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Plugins) GetExtensions() []*Extensions {
	if x != nil {
		return x.Extensions
	}
	return nil
}

type Extensions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id      string  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Enabled bool    `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Config  *Config `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Extensions) Reset() {
	*x = Extensions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extensions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extensions) ProtoMessage() {}

func (x *Extensions) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extensions.ProtoReflect.Descriptor instead.
func (*Extensions) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDescGZIP(), []int{3}
}

func (x *Extensions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Extensions) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Extensions) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Extensions) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDescGZIP(), []int{4}
}

type Repositories struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Url  string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Repositories) Reset() {
	*x = Repositories{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repositories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repositories) ProtoMessage() {}

func (x *Repositories) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repositories.ProtoReflect.Descriptor instead.
func (*Repositories) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDescGZIP(), []int{5}
}

func (x *Repositories) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Repositories) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Repositories) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugins map[string]*Plugins `protobuf:"bytes,1,rep,name=plugins,proto3" json:"plugins,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDescGZIP(), []int{6}
}

func (x *Test) GetPlugins() map[string]*Plugins {
	if x != nil {
		return x.Plugins
	}
	return nil
}

var File_proto_deploymentConfigurations_spinnaker_Spinnaker_proto protoreflect.FileDescriptor

var file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDesc = []byte{
	0x0a, 0x38, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x53, 0x70, 0x69, 0x6e, 0x6e,
	0x61, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x22, 0x51, 0x0a, 0x09, 0x53,
	0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0d, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65,
	0x72, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x0d, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x83,
	0x01, 0x0a, 0x0d, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x2f, 0x0a, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61,
	0x6b, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x12, 0x41, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x07, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x70, 0x69,
	0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x7b, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2f,
	0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x08, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x44, 0x0a, 0x0c, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22,
	0x9a, 0x01, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x07, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x1a, 0x54, 0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2b, 0x5a, 0x29,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDescOnce sync.Once
	file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDescData = file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDesc
)

func file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDescGZIP() []byte {
	file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDescOnce.Do(func() {
		file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDescData)
	})
	return file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDescData
}

var file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_goTypes = []interface{}{
	(*Spinnaker)(nil),     // 0: proto.spinnaker.Spinnaker
	(*Extensibility)(nil), // 1: proto.spinnaker.Extensibility
	(*Plugins)(nil),       // 2: proto.spinnaker.Plugins
	(*Extensions)(nil),    // 3: proto.spinnaker.Extensions
	(*Config)(nil),        // 4: proto.spinnaker.Config
	(*Repositories)(nil),  // 5: proto.spinnaker.Repositories
	(*Test)(nil),          // 6: proto.spinnaker.Test
	nil,                   // 7: proto.spinnaker.Test.PluginsEntry
}
var file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_depIdxs = []int32{
	1, // 0: proto.spinnaker.Spinnaker.extensibility:type_name -> proto.spinnaker.Extensibility
	6, // 1: proto.spinnaker.Extensibility.plugins:type_name -> proto.spinnaker.Test
	5, // 2: proto.spinnaker.Extensibility.repositories:type_name -> proto.spinnaker.Repositories
	4, // 3: proto.spinnaker.Plugins.config:type_name -> proto.spinnaker.Config
	3, // 4: proto.spinnaker.Plugins.extensions:type_name -> proto.spinnaker.Extensions
	4, // 5: proto.spinnaker.Extensions.config:type_name -> proto.spinnaker.Config
	7, // 6: proto.spinnaker.Test.plugins:type_name -> proto.spinnaker.Test.PluginsEntry
	2, // 7: proto.spinnaker.Test.PluginsEntry.value:type_name -> proto.spinnaker.Plugins
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_init() }
func file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_init() {
	if File_proto_deploymentConfigurations_spinnaker_Spinnaker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spinnaker); i {
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
		file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extensibility); i {
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
		file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugins); i {
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
		file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extensions); i {
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
		file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repositories); i {
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
		file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
			RawDescriptor: file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_goTypes,
		DependencyIndexes: file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_depIdxs,
		MessageInfos:      file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_msgTypes,
	}.Build()
	File_proto_deploymentConfigurations_spinnaker_Spinnaker_proto = out.File
	file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_rawDesc = nil
	file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_goTypes = nil
	file_proto_deploymentConfigurations_spinnaker_Spinnaker_proto_depIdxs = nil
}
