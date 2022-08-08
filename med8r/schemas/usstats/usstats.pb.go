// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: usstats/usstats.proto

package usstats

import (
	_ "github.com/satjinder/med8r/schemas/gprotos"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	anypb "google.golang.org/protobuf/types/known/anypb"
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
	Measures   string `protobuf:"bytes,2,opt,name=measures,proto3" json:"measures,omitempty"`
}

func (x *GetStatsRequest) Reset() {
	*x = GetStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usstats_usstats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsRequest) ProtoMessage() {}

func (x *GetStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usstats_usstats_proto_msgTypes[0]
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
	return file_usstats_usstats_proto_rawDescGZIP(), []int{0}
}

func (x *GetStatsRequest) GetDrilldowns() string {
	if x != nil {
		return x.Drilldowns
	}
	return ""
}

func (x *GetStatsRequest) GetMeasures() string {
	if x != nil {
		return x.Measures
	}
	return ""
}

type GetStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   []*Data   `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Source []*Source `protobuf:"bytes,2,rep,name=source,proto3" json:"source,omitempty"`
}

func (x *GetStatsResponse) Reset() {
	*x = GetStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usstats_usstats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsResponse) ProtoMessage() {}

func (x *GetStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usstats_usstats_proto_msgTypes[1]
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
	return file_usstats_usstats_proto_rawDescGZIP(), []int{1}
}

func (x *GetStatsResponse) GetData() []*Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetStatsResponse) GetSource() []*Source {
	if x != nil {
		return x.Source
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDNation   string `protobuf:"bytes,1,opt,name=ID_nation,json=ID Nation,proto3" json:"ID_nation,omitempty"`
	Nation     string `protobuf:"bytes,2,opt,name=nation,json=Nation,proto3" json:"nation,omitempty"`
	IDYear     uint32 `protobuf:"varint,3,opt,name=ID_year,json=ID Year,proto3" json:"ID_year,omitempty"`
	Year       string `protobuf:"bytes,4,opt,name=year,json=Year,proto3" json:"year,omitempty"`
	Population uint32 `protobuf:"varint,5,opt,name=population,json=Population,proto3" json:"population,omitempty"`
	SlugNation string `protobuf:"bytes,6,opt,name=Slug_nation,json=Slug Nation,proto3" json:"Slug_nation,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usstats_usstats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_usstats_usstats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_usstats_usstats_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetIDNation() string {
	if x != nil {
		return x.IDNation
	}
	return ""
}

func (x *Data) GetNation() string {
	if x != nil {
		return x.Nation
	}
	return ""
}

func (x *Data) GetIDYear() uint32 {
	if x != nil {
		return x.IDYear
	}
	return 0
}

func (x *Data) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Data) GetPopulation() uint32 {
	if x != nil {
		return x.Population
	}
	return 0
}

func (x *Data) GetSlugNation() string {
	if x != nil {
		return x.SlugNation
	}
	return ""
}

type Annotations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceName        string `protobuf:"bytes,1,opt,name=source_name,json=sourceName,proto3" json:"source_name,omitempty"`
	SourceDescription string `protobuf:"bytes,2,opt,name=source_description,json=sourceDescription,proto3" json:"source_description,omitempty"`
	DatasetName       string `protobuf:"bytes,3,opt,name=dataset_name,json=datasetName,proto3" json:"dataset_name,omitempty"`
	DatasetLink       string `protobuf:"bytes,4,opt,name=dataset_link,json=datasetLink,proto3" json:"dataset_link,omitempty"`
	TableId           string `protobuf:"bytes,5,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	Topic             string `protobuf:"bytes,6,opt,name=topic,proto3" json:"topic,omitempty"`
	Subtopic          string `protobuf:"bytes,7,opt,name=subtopic,proto3" json:"subtopic,omitempty"`
}

func (x *Annotations) Reset() {
	*x = Annotations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usstats_usstats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Annotations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Annotations) ProtoMessage() {}

func (x *Annotations) ProtoReflect() protoreflect.Message {
	mi := &file_usstats_usstats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Annotations.ProtoReflect.Descriptor instead.
func (*Annotations) Descriptor() ([]byte, []int) {
	return file_usstats_usstats_proto_rawDescGZIP(), []int{3}
}

func (x *Annotations) GetSourceName() string {
	if x != nil {
		return x.SourceName
	}
	return ""
}

func (x *Annotations) GetSourceDescription() string {
	if x != nil {
		return x.SourceDescription
	}
	return ""
}

func (x *Annotations) GetDatasetName() string {
	if x != nil {
		return x.DatasetName
	}
	return ""
}

func (x *Annotations) GetDatasetLink() string {
	if x != nil {
		return x.DatasetLink
	}
	return ""
}

func (x *Annotations) GetTableId() string {
	if x != nil {
		return x.TableId
	}
	return ""
}

func (x *Annotations) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Annotations) GetSubtopic() string {
	if x != nil {
		return x.Subtopic
	}
	return ""
}

type Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Measures      []string     `protobuf:"bytes,1,rep,name=measures,proto3" json:"measures,omitempty"`
	Annotations   *Annotations `protobuf:"bytes,2,opt,name=annotations,proto3" json:"annotations,omitempty"`
	Name          string       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Substitutions []*anypb.Any `protobuf:"bytes,4,rep,name=substitutions,proto3" json:"substitutions,omitempty"`
}

func (x *Source) Reset() {
	*x = Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usstats_usstats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_usstats_usstats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_usstats_usstats_proto_rawDescGZIP(), []int{4}
}

func (x *Source) GetMeasures() []string {
	if x != nil {
		return x.Measures
	}
	return nil
}

func (x *Source) GetAnnotations() *Annotations {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Source) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Source) GetSubstitutions() []*anypb.Any {
	if x != nil {
		return x.Substitutions
	}
	return nil
}

var File_usstats_usstats_proto protoreflect.FileDescriptor

var file_usstats_usstats_proto_rawDesc = []byte{
	0x0a, 0x15, 0x75, 0x73, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x75, 0x73, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x75, 0x73, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x72, 0x69, 0x6c, 0x6c, 0x64,
	0x6f, 0x77, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x72, 0x69, 0x6c,
	0x6c, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x73, 0x22, 0x5e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x49,
	0x44, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x49, 0x44, 0x20, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x44, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x49, 0x44, 0x20, 0x59, 0x65, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x53, 0x6c, 0x75, 0x67, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x6c, 0x75, 0x67, 0x20, 0x4e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xf0, 0x01, 0x0a, 0x0b, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x22, 0xac, 0x01, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0b, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x75, 0x73, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x74,
	0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x32, 0x88, 0x02, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x73, 0x41, 0x50, 0x49,
	0x12, 0xfb, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x18, 0x2e,
	0x75, 0x73, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0xb9, 0x01, 0xa2, 0x82, 0x19, 0xb4, 0x01, 0x0a, 0x11, 0x61, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x1a, 0x26, 0x0a,
	0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a,
	0x01, 0x31, 0x12, 0x11, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x3a, 0x72, 0x65, 0x61, 0x64, 0x1a, 0x77, 0x0a, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x2d, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x12, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x03, 0x47, 0x45, 0x54, 0x12, 0x12, 0x0a, 0x0b, 0x75,
	0x72, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x03, 0x47, 0x45, 0x54, 0x12,
	0x1b, 0x0a, 0x0f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6b,
	0x65, 0x79, 0x12, 0x08, 0x55, 0x53, 0x2d, 0x53, 0x54, 0x41, 0x54, 0x53, 0x12, 0x10, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x08, 0x55, 0x53, 0x2d, 0x53, 0x54, 0x41, 0x54, 0x53, 0x42, 0x34,
	0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x74,
	0x6a, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x64, 0x38, 0x72, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x73, 0x2f, 0x75, 0x73, 0x73, 0x74, 0x61, 0x74, 0x73, 0x3b, 0x75, 0x73, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_usstats_usstats_proto_rawDescOnce sync.Once
	file_usstats_usstats_proto_rawDescData = file_usstats_usstats_proto_rawDesc
)

func file_usstats_usstats_proto_rawDescGZIP() []byte {
	file_usstats_usstats_proto_rawDescOnce.Do(func() {
		file_usstats_usstats_proto_rawDescData = protoimpl.X.CompressGZIP(file_usstats_usstats_proto_rawDescData)
	})
	return file_usstats_usstats_proto_rawDescData
}

var file_usstats_usstats_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_usstats_usstats_proto_goTypes = []interface{}{
	(*GetStatsRequest)(nil),  // 0: usstats.GetStatsRequest
	(*GetStatsResponse)(nil), // 1: usstats.GetStatsResponse
	(*Data)(nil),             // 2: usstats.Data
	(*Annotations)(nil),      // 3: usstats.Annotations
	(*Source)(nil),           // 4: usstats.Source
	(*anypb.Any)(nil),        // 5: google.protobuf.Any
}
var file_usstats_usstats_proto_depIdxs = []int32{
	2, // 0: usstats.GetStatsResponse.data:type_name -> usstats.Data
	4, // 1: usstats.GetStatsResponse.source:type_name -> usstats.Source
	3, // 2: usstats.Source.annotations:type_name -> usstats.Annotations
	5, // 3: usstats.Source.substitutions:type_name -> google.protobuf.Any
	0, // 4: usstats.StatsAPI.GetStats:input_type -> usstats.GetStatsRequest
	1, // 5: usstats.StatsAPI.GetStats:output_type -> usstats.GetStatsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_usstats_usstats_proto_init() }
func file_usstats_usstats_proto_init() {
	if File_usstats_usstats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_usstats_usstats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_usstats_usstats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_usstats_usstats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_usstats_usstats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Annotations); i {
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
		file_usstats_usstats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Source); i {
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
			RawDescriptor: file_usstats_usstats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_usstats_usstats_proto_goTypes,
		DependencyIndexes: file_usstats_usstats_proto_depIdxs,
		MessageInfos:      file_usstats_usstats_proto_msgTypes,
	}.Build()
	File_usstats_usstats_proto = out.File
	file_usstats_usstats_proto_rawDesc = nil
	file_usstats_usstats_proto_goTypes = nil
	file_usstats_usstats_proto_depIdxs = nil
}
