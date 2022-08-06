// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: statsservice/stats.proto

package statsservice

import (
	_ "github.com/satjinder/med8r/gprotos"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Drilldowns string `protobuf:"bytes,1,opt,name=drilldowns,proto3" json:"drilldowns,omitempty"`
}

func (x *GetStatsRequest) Reset() {
	*x = GetStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statsservice_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsRequest) ProtoMessage() {}

func (x *GetStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_statsservice_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsRequest.ProtoReflect.Descriptor instead.
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return file_statsservice_stats_proto_rawDescGZIP(), []int{0}
}

func (x *GetStatsRequest) GetDrilldowns() string {
	if x != nil {
		return x.Drilldowns
	}
	return ""
}

type GetStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Stats `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetStatsResponse) Reset() {
	*x = GetStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statsservice_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsResponse) ProtoMessage() {}

func (x *GetStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_statsservice_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsResponse.ProtoReflect.Descriptor instead.
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return file_statsservice_stats_proto_rawDescGZIP(), []int{1}
}

func (x *GetStatsResponse) GetData() []*Stats {
	if x != nil {
		return x.Data
	}
	return nil
}

type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year string `protobuf:"bytes,1,opt,name=Year,proto3" json:"Year,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statsservice_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_statsservice_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_statsservice_stats_proto_rawDescGZIP(), []int{2}
}

func (x *Stats) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

var File_statsservice_stats_proto protoreflect.FileDescriptor

var file_statsservice_stats_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x74, 0x61, 0x74, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x64, 0x72, 0x69, 0x6c, 0x6c, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x72, 0x69, 0x6c, 0x6c, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x22, 0x34, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x1b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x59,
	0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x32,
	0x65, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x73, 0x41, 0x50, 0x49, 0x12, 0x59, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0xa2, 0x82, 0x19, 0x18, 0x0a, 0x11,
	0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x72, 0x65, 0x61,
	0x64, 0x12, 0x03, 0x4a, 0x57, 0x54, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x74, 0x6a, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x6d,
	0x65, 0x64, 0x38, 0x72, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x3b, 0x73, 0x74, 0x61, 0x74, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_statsservice_stats_proto_rawDescOnce sync.Once
	file_statsservice_stats_proto_rawDescData = file_statsservice_stats_proto_rawDesc
)

func file_statsservice_stats_proto_rawDescGZIP() []byte {
	file_statsservice_stats_proto_rawDescOnce.Do(func() {
		file_statsservice_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_statsservice_stats_proto_rawDescData)
	})
	return file_statsservice_stats_proto_rawDescData
}

var file_statsservice_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_statsservice_stats_proto_goTypes = []interface{}{
	(*GetStatsRequest)(nil),  // 0: stats.GetStatsRequest
	(*GetStatsResponse)(nil), // 1: stats.GetStatsResponse
	(*Stats)(nil),            // 2: stats.stats
}
var file_statsservice_stats_proto_depIdxs = []int32{
	2, // 0: stats.GetStatsResponse.data:type_name -> stats.stats
	0, // 1: stats.StatsAPI.GetStats:input_type -> stats.GetStatsRequest
	1, // 2: stats.StatsAPI.GetStats:output_type -> stats.GetStatsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_statsservice_stats_proto_init() }
func file_statsservice_stats_proto_init() {
	if File_statsservice_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_statsservice_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatsRequest); i {
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
		file_statsservice_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatsResponse); i {
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
		file_statsservice_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats); i {
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
			RawDescriptor: file_statsservice_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_statsservice_stats_proto_goTypes,
		DependencyIndexes: file_statsservice_stats_proto_depIdxs,
		MessageInfos:      file_statsservice_stats_proto_msgTypes,
	}.Build()
	File_statsservice_stats_proto = out.File
	file_statsservice_stats_proto_rawDesc = nil
	file_statsservice_stats_proto_goTypes = nil
	file_statsservice_stats_proto_depIdxs = nil
}
