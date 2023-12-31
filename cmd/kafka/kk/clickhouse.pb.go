// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: clickhouse/clickhouse.proto

package kk

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

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database  string             `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Table     string             `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	Timestamp string             `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Labels    map[string]string  `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Values    map[string]float64 `protobuf:"bytes,5,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clickhouse_clickhouse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_clickhouse_clickhouse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_clickhouse_clickhouse_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *Point) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *Point) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Point) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Point) GetValues() map[string]float64 {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_clickhouse_clickhouse_proto protoreflect.FileDescriptor

var file_clickhouse_clickhouse_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x63, 0x6c, 0x69,
	0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x02,
	0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a,
	0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x4c, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6d, 0x61,
	0x6c, 0x6c, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x42, 0x0f, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x48, 0x6f, 0x75, 0x73, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x0c, 0x2e, 0x2f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clickhouse_clickhouse_proto_rawDescOnce sync.Once
	file_clickhouse_clickhouse_proto_rawDescData = file_clickhouse_clickhouse_proto_rawDesc
)

func file_clickhouse_clickhouse_proto_rawDescGZIP() []byte {
	file_clickhouse_clickhouse_proto_rawDescOnce.Do(func() {
		file_clickhouse_clickhouse_proto_rawDescData = protoimpl.X.CompressGZIP(file_clickhouse_clickhouse_proto_rawDescData)
	})
	return file_clickhouse_clickhouse_proto_rawDescData
}

var file_clickhouse_clickhouse_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_clickhouse_clickhouse_proto_goTypes = []interface{}{
	(*Point)(nil), // 0: Point
	nil,           // 1: Point.LabelsEntry
	nil,           // 2: Point.ValuesEntry
}
var file_clickhouse_clickhouse_proto_depIdxs = []int32{
	1, // 0: Point.labels:type_name -> Point.LabelsEntry
	2, // 1: Point.values:type_name -> Point.ValuesEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_clickhouse_clickhouse_proto_init() }
func file_clickhouse_clickhouse_proto_init() {
	if File_clickhouse_clickhouse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clickhouse_clickhouse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
			RawDescriptor: file_clickhouse_clickhouse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_clickhouse_clickhouse_proto_goTypes,
		DependencyIndexes: file_clickhouse_clickhouse_proto_depIdxs,
		MessageInfos:      file_clickhouse_clickhouse_proto_msgTypes,
	}.Build()
	File_clickhouse_clickhouse_proto = out.File
	file_clickhouse_clickhouse_proto_rawDesc = nil
	file_clickhouse_clickhouse_proto_goTypes = nil
	file_clickhouse_clickhouse_proto_depIdxs = nil
}
