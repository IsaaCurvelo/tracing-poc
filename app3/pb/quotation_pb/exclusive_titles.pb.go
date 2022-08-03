// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: .protos/exclusive_titles.proto

package quotation_pb

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file___protos_exclusive_titles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file___protos_exclusive_titles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file___protos_exclusive_titles_proto_rawDescGZIP(), []int{0}
}

type ExclusiveTitlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VendorId string `protobuf:"bytes,1,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
}

func (x *ExclusiveTitlesRequest) Reset() {
	*x = ExclusiveTitlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file___protos_exclusive_titles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExclusiveTitlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExclusiveTitlesRequest) ProtoMessage() {}

func (x *ExclusiveTitlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file___protos_exclusive_titles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExclusiveTitlesRequest.ProtoReflect.Descriptor instead.
func (*ExclusiveTitlesRequest) Descriptor() ([]byte, []int) {
	return file___protos_exclusive_titles_proto_rawDescGZIP(), []int{1}
}

func (x *ExclusiveTitlesRequest) GetVendorId() string {
	if x != nil {
		return x.VendorId
	}
	return ""
}

type ExclusiveTitle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	VendorId string `protobuf:"bytes,3,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
}

func (x *ExclusiveTitle) Reset() {
	*x = ExclusiveTitle{}
	if protoimpl.UnsafeEnabled {
		mi := &file___protos_exclusive_titles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExclusiveTitle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExclusiveTitle) ProtoMessage() {}

func (x *ExclusiveTitle) ProtoReflect() protoreflect.Message {
	mi := &file___protos_exclusive_titles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExclusiveTitle.ProtoReflect.Descriptor instead.
func (*ExclusiveTitle) Descriptor() ([]byte, []int) {
	return file___protos_exclusive_titles_proto_rawDescGZIP(), []int{2}
}

func (x *ExclusiveTitle) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExclusiveTitle) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExclusiveTitle) GetVendorId() string {
	if x != nil {
		return x.VendorId
	}
	return ""
}

type ExclusiveTitlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExclusiveTitles []*ExclusiveTitle `protobuf:"bytes,1,rep,name=exclusive_titles,json=exclusiveTitles,proto3" json:"exclusive_titles,omitempty"`
}

func (x *ExclusiveTitlesResponse) Reset() {
	*x = ExclusiveTitlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file___protos_exclusive_titles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExclusiveTitlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExclusiveTitlesResponse) ProtoMessage() {}

func (x *ExclusiveTitlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file___protos_exclusive_titles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExclusiveTitlesResponse.ProtoReflect.Descriptor instead.
func (*ExclusiveTitlesResponse) Descriptor() ([]byte, []int) {
	return file___protos_exclusive_titles_proto_rawDescGZIP(), []int{3}
}

func (x *ExclusiveTitlesResponse) GetExclusiveTitles() []*ExclusiveTitle {
	if x != nil {
		return x.ExclusiveTitles
	}
	return nil
}

var File___protos_exclusive_titles_proto protoreflect.FileDescriptor

var file___protos_exclusive_titles_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73,
	0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x35, 0x0a,
	0x16, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x0e, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76,
	0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x17, 0x45, 0x78, 0x63, 0x6c, 0x75,
	0x73, 0x69, 0x76, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x10, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x5f,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x62, 0x2e, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x52, 0x0f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x73, 0x32, 0x62, 0x0a, 0x16, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x2e, 0x70,
	0x62, 0x2e, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x78,
	0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x2f, 0x70, 0x62, 0x2f, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file___protos_exclusive_titles_proto_rawDescOnce sync.Once
	file___protos_exclusive_titles_proto_rawDescData = file___protos_exclusive_titles_proto_rawDesc
)

func file___protos_exclusive_titles_proto_rawDescGZIP() []byte {
	file___protos_exclusive_titles_proto_rawDescOnce.Do(func() {
		file___protos_exclusive_titles_proto_rawDescData = protoimpl.X.CompressGZIP(file___protos_exclusive_titles_proto_rawDescData)
	})
	return file___protos_exclusive_titles_proto_rawDescData
}

var file___protos_exclusive_titles_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file___protos_exclusive_titles_proto_goTypes = []interface{}{
	(*Empty)(nil),                   // 0: pb.Empty
	(*ExclusiveTitlesRequest)(nil),  // 1: pb.ExclusiveTitlesRequest
	(*ExclusiveTitle)(nil),          // 2: pb.ExclusiveTitle
	(*ExclusiveTitlesResponse)(nil), // 3: pb.ExclusiveTitlesResponse
}
var file___protos_exclusive_titles_proto_depIdxs = []int32{
	2, // 0: pb.ExclusiveTitlesResponse.exclusive_titles:type_name -> pb.ExclusiveTitle
	1, // 1: pb.ExclusiveTitlesService.GetByVendorID:input_type -> pb.ExclusiveTitlesRequest
	3, // 2: pb.ExclusiveTitlesService.GetByVendorID:output_type -> pb.ExclusiveTitlesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file___protos_exclusive_titles_proto_init() }
func file___protos_exclusive_titles_proto_init() {
	if File___protos_exclusive_titles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file___protos_exclusive_titles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file___protos_exclusive_titles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExclusiveTitlesRequest); i {
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
		file___protos_exclusive_titles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExclusiveTitle); i {
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
		file___protos_exclusive_titles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExclusiveTitlesResponse); i {
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
			RawDescriptor: file___protos_exclusive_titles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file___protos_exclusive_titles_proto_goTypes,
		DependencyIndexes: file___protos_exclusive_titles_proto_depIdxs,
		MessageInfos:      file___protos_exclusive_titles_proto_msgTypes,
	}.Build()
	File___protos_exclusive_titles_proto = out.File
	file___protos_exclusive_titles_proto_rawDesc = nil
	file___protos_exclusive_titles_proto_goTypes = nil
	file___protos_exclusive_titles_proto_depIdxs = nil
}
