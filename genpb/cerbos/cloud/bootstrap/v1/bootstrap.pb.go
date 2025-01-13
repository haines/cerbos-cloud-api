// Copyright 2021-2025 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: cerbos/cloud/bootstrap/v1/bootstrap.proto

package bootstrapv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	v1 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/bundle/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PDPConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Meta          *PDPConfig_Meta        `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	BundleInfo    *v1.BundleInfo         `protobuf:"bytes,2,opt,name=bundle_info,json=bundleInfo,proto3" json:"bundle_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PDPConfig) Reset() {
	*x = PDPConfig{}
	mi := &file_cerbos_cloud_bootstrap_v1_bootstrap_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PDPConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PDPConfig) ProtoMessage() {}

func (x *PDPConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_cloud_bootstrap_v1_bootstrap_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PDPConfig.ProtoReflect.Descriptor instead.
func (*PDPConfig) Descriptor() ([]byte, []int) {
	return file_cerbos_cloud_bootstrap_v1_bootstrap_proto_rawDescGZIP(), []int{0}
}

func (x *PDPConfig) GetMeta() *PDPConfig_Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *PDPConfig) GetBundleInfo() *v1.BundleInfo {
	if x != nil {
		return x.BundleInfo
	}
	return nil
}

type PDPConfig_Meta struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CommitHash    string                 `protobuf:"bytes,2,opt,name=commit_hash,json=commitHash,proto3" json:"commit_hash,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PDPConfig_Meta) Reset() {
	*x = PDPConfig_Meta{}
	mi := &file_cerbos_cloud_bootstrap_v1_bootstrap_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PDPConfig_Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PDPConfig_Meta) ProtoMessage() {}

func (x *PDPConfig_Meta) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_cloud_bootstrap_v1_bootstrap_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PDPConfig_Meta.ProtoReflect.Descriptor instead.
func (*PDPConfig_Meta) Descriptor() ([]byte, []int) {
	return file_cerbos_cloud_bootstrap_v1_bootstrap_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PDPConfig_Meta) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PDPConfig_Meta) GetCommitHash() string {
	if x != nil {
		return x.CommitHash
	}
	return ""
}

var File_cerbos_cloud_bootstrap_v1_bootstrap_proto protoreflect.FileDescriptor

var file_cerbos_cloud_bootstrap_v1_bootstrap_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x65, 0x72,
	0x62, 0x6f, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74,
	0x72, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x02, 0x0a, 0x09, 0x50, 0x44,
	0x50, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x45, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x44, 0x50, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x4b,
	0x0a, 0x0b, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x0a, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x73, 0x0a, 0x04, 0x4d,
	0x65, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x62, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x3b,
	0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cerbos_cloud_bootstrap_v1_bootstrap_proto_rawDescOnce sync.Once
	file_cerbos_cloud_bootstrap_v1_bootstrap_proto_rawDescData = file_cerbos_cloud_bootstrap_v1_bootstrap_proto_rawDesc
)

func file_cerbos_cloud_bootstrap_v1_bootstrap_proto_rawDescGZIP() []byte {
	file_cerbos_cloud_bootstrap_v1_bootstrap_proto_rawDescOnce.Do(func() {
		file_cerbos_cloud_bootstrap_v1_bootstrap_proto_rawDescData = protoimpl.X.CompressGZIP(file_cerbos_cloud_bootstrap_v1_bootstrap_proto_rawDescData)
	})
	return file_cerbos_cloud_bootstrap_v1_bootstrap_proto_rawDescData
}

var file_cerbos_cloud_bootstrap_v1_bootstrap_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cerbos_cloud_bootstrap_v1_bootstrap_proto_goTypes = []any{
	(*PDPConfig)(nil),             // 0: cerbos.cloud.bootstrap.v1.PDPConfig
	(*PDPConfig_Meta)(nil),        // 1: cerbos.cloud.bootstrap.v1.PDPConfig.Meta
	(*v1.BundleInfo)(nil),         // 2: cerbos.cloud.bundle.v1.BundleInfo
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_cerbos_cloud_bootstrap_v1_bootstrap_proto_depIdxs = []int32{
	1, // 0: cerbos.cloud.bootstrap.v1.PDPConfig.meta:type_name -> cerbos.cloud.bootstrap.v1.PDPConfig.Meta
	2, // 1: cerbos.cloud.bootstrap.v1.PDPConfig.bundle_info:type_name -> cerbos.cloud.bundle.v1.BundleInfo
	3, // 2: cerbos.cloud.bootstrap.v1.PDPConfig.Meta.created_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cerbos_cloud_bootstrap_v1_bootstrap_proto_init() }
func file_cerbos_cloud_bootstrap_v1_bootstrap_proto_init() {
	if File_cerbos_cloud_bootstrap_v1_bootstrap_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cerbos_cloud_bootstrap_v1_bootstrap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cerbos_cloud_bootstrap_v1_bootstrap_proto_goTypes,
		DependencyIndexes: file_cerbos_cloud_bootstrap_v1_bootstrap_proto_depIdxs,
		MessageInfos:      file_cerbos_cloud_bootstrap_v1_bootstrap_proto_msgTypes,
	}.Build()
	File_cerbos_cloud_bootstrap_v1_bootstrap_proto = out.File
	file_cerbos_cloud_bootstrap_v1_bootstrap_proto_rawDesc = nil
	file_cerbos_cloud_bootstrap_v1_bootstrap_proto_goTypes = nil
	file_cerbos_cloud_bootstrap_v1_bootstrap_proto_depIdxs = nil
}
