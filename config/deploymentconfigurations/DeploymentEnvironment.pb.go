// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.5
// source: proto/deploymentConfigurations/deploymentEnv/DeploymentEnvironment.proto

package deploymentConfigurations

import (
	deploymentEnv "config/deploymentConfigurations/deploymentEnv"
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
type DeploymentEnvironment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Container      []*deploymentEnv.Containers     `protobuf:"bytes,1,rep,name=container,proto3" json:"container,omitempty"`
	InitContainers []*deploymentEnv.InitContainers `protobuf:"bytes,2,rep,name=initContainers,proto3" json:"initContainers,omitempty"`
	Volumes        []*deploymentEnv.Volumes        `protobuf:"bytes,3,rep,name=volumes,proto3" json:"volumes,omitempty"`
}

func (x *DeploymentEnvironment) Reset() {
	*x = DeploymentEnvironment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentEnvironment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentEnvironment) ProtoMessage() {}

func (x *DeploymentEnvironment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentEnvironment.ProtoReflect.Descriptor instead.
func (*DeploymentEnvironment) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescGZIP(), []int{0}
}

func (x *DeploymentEnvironment) GetContainer() []*deploymentEnv.Containers {
	if x != nil {
		return x.Container
	}
	return nil
}

func (x *DeploymentEnvironment) GetInitContainers() []*deploymentEnv.InitContainers {
	if x != nil {
		return x.InitContainers
	}
	return nil
}

func (x *DeploymentEnvironment) GetVolumes() []*deploymentEnv.Volumes {
	if x != nil {
		return x.Volumes
	}
	return nil
}

var File_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto protoreflect.FileDescriptor

var file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDesc = []byte{
	0x0a, 0x48, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x2f, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x3d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x45, 0x6e, 0x76, 0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x45, 0x6e, 0x76, 0x2f, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x09,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x3e, 0x0a, 0x07, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x52,
	0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x42, 0x21, 0x5a, 0x1f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescOnce sync.Once
	file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescData = file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDesc
)

func file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescGZIP() []byte {
	file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescOnce.Do(func() {
		file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescData)
	})
	return file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescData
}

var file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_goTypes = []interface{}{
	(*DeploymentEnvironment)(nil),        // 0: proto.deploymentEnvironment.DeploymentEnvironment
	(*deploymentEnv.Containers)(nil),     // 1: proto.deploymentEnvironment.Containers
	(*deploymentEnv.InitContainers)(nil), // 2: proto.deploymentEnvironment.InitContainers
	(*deploymentEnv.Volumes)(nil),        // 3: proto.deploymentEnvironment.Volumes
}
var file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_depIdxs = []int32{
	1, // 0: proto.deploymentEnvironment.DeploymentEnvironment.container:type_name -> proto.deploymentEnvironment.Containers
	2, // 1: proto.deploymentEnvironment.DeploymentEnvironment.initContainers:type_name -> proto.deploymentEnvironment.InitContainers
	3, // 2: proto.deploymentEnvironment.DeploymentEnvironment.volumes:type_name -> proto.deploymentEnvironment.Volumes
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_init() }
func file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_init() {
	if File_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentEnvironment); i {
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
			RawDescriptor: file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_goTypes,
		DependencyIndexes: file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_depIdxs,
		MessageInfos:      file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes,
	}.Build()
	File_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto = out.File
	file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDesc = nil
	file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_goTypes = nil
	file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_depIdxs = nil
}
