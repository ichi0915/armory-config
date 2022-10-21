// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.5
// source: proto/deploymentConfigurations/repository/Nexus.proto

package repository

import (
	permissions "github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/permissions"
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
type Nexus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled  bool           `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Searches []*NexusSearch `protobuf:"bytes,2,rep,name=searches,proto3" json:"searches,omitempty"`
}

func (x *Nexus) Reset() {
	*x = Nexus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_repository_Nexus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nexus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nexus) ProtoMessage() {}

func (x *Nexus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_repository_Nexus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nexus.ProtoReflect.Descriptor instead.
func (*Nexus) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_repository_Nexus_proto_rawDescGZIP(), []int{0}
}

func (x *Nexus) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Nexus) GetSearches() []*NexusSearch {
	if x != nil {
		return x.Searches
	}
	return nil
}

type NexusSearch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Permissions *permissions.Permissions `protobuf:"bytes,2,opt,name=permissions,proto3" json:"permissions,omitempty"`
	BaseUrl     string                   `protobuf:"bytes,3,opt,name=baseUrl,proto3" json:"baseUrl,omitempty"`
	Repo        string                   `protobuf:"bytes,4,opt,name=repo,proto3" json:"repo,omitempty"`
	NodeId      string                   `protobuf:"bytes,5,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Username    string                   `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	Password    string                   `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *NexusSearch) Reset() {
	*x = NexusSearch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_repository_Nexus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NexusSearch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NexusSearch) ProtoMessage() {}

func (x *NexusSearch) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_repository_Nexus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NexusSearch.ProtoReflect.Descriptor instead.
func (*NexusSearch) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_repository_Nexus_proto_rawDescGZIP(), []int{1}
}

func (x *NexusSearch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NexusSearch) GetPermissions() *permissions.Permissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *NexusSearch) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *NexusSearch) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *NexusSearch) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *NexusSearch) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *NexusSearch) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_proto_deploymentConfigurations_repository_Nexus_proto protoreflect.FileDescriptor

var file_proto_deploymentConfigurations_repository_Nexus_proto_rawDesc = []byte{
	0x0a, 0x35, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x4e, 0x65, 0x78, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x3c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x05, 0x4e, 0x65, 0x78, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x4e, 0x65, 0x78, 0x75, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x08, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x65, 0x73, 0x22, 0xe1, 0x01, 0x0a, 0x0b, 0x4e, 0x65, 0x78, 0x75, 0x73, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61,
	0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x2c, 0x5a, 0x2a, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_deploymentConfigurations_repository_Nexus_proto_rawDescOnce sync.Once
	file_proto_deploymentConfigurations_repository_Nexus_proto_rawDescData = file_proto_deploymentConfigurations_repository_Nexus_proto_rawDesc
)

func file_proto_deploymentConfigurations_repository_Nexus_proto_rawDescGZIP() []byte {
	file_proto_deploymentConfigurations_repository_Nexus_proto_rawDescOnce.Do(func() {
		file_proto_deploymentConfigurations_repository_Nexus_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_deploymentConfigurations_repository_Nexus_proto_rawDescData)
	})
	return file_proto_deploymentConfigurations_repository_Nexus_proto_rawDescData
}

var file_proto_deploymentConfigurations_repository_Nexus_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_deploymentConfigurations_repository_Nexus_proto_goTypes = []interface{}{
	(*Nexus)(nil),                   // 0: proto.repository.Nexus
	(*NexusSearch)(nil),             // 1: proto.repository.NexusSearch
	(*permissions.Permissions)(nil), // 2: proto.permissions.Permissions
}
var file_proto_deploymentConfigurations_repository_Nexus_proto_depIdxs = []int32{
	1, // 0: proto.repository.Nexus.searches:type_name -> proto.repository.NexusSearch
	2, // 1: proto.repository.NexusSearch.permissions:type_name -> proto.permissions.Permissions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_deploymentConfigurations_repository_Nexus_proto_init() }
func file_proto_deploymentConfigurations_repository_Nexus_proto_init() {
	if File_proto_deploymentConfigurations_repository_Nexus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_deploymentConfigurations_repository_Nexus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nexus); i {
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
		file_proto_deploymentConfigurations_repository_Nexus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NexusSearch); i {
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
			RawDescriptor: file_proto_deploymentConfigurations_repository_Nexus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_deploymentConfigurations_repository_Nexus_proto_goTypes,
		DependencyIndexes: file_proto_deploymentConfigurations_repository_Nexus_proto_depIdxs,
		MessageInfos:      file_proto_deploymentConfigurations_repository_Nexus_proto_msgTypes,
	}.Build()
	File_proto_deploymentConfigurations_repository_Nexus_proto = out.File
	file_proto_deploymentConfigurations_repository_Nexus_proto_rawDesc = nil
	file_proto_deploymentConfigurations_repository_Nexus_proto_goTypes = nil
	file_proto_deploymentConfigurations_repository_Nexus_proto_depIdxs = nil
}
