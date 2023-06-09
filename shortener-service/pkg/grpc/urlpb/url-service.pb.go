// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: proto/url-service.proto

package urlpb

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

type NewShortURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LongURL string `protobuf:"bytes,1,opt,name=longURL,proto3" json:"longURL,omitempty"`
}

func (x *NewShortURLRequest) Reset() {
	*x = NewShortURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewShortURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewShortURLRequest) ProtoMessage() {}

func (x *NewShortURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewShortURLRequest.ProtoReflect.Descriptor instead.
func (*NewShortURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_url_service_proto_rawDescGZIP(), []int{0}
}

func (x *NewShortURLRequest) GetLongURL() string {
	if x != nil {
		return x.LongURL
	}
	return ""
}

type NewShortURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortURL string `protobuf:"bytes,1,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
}

func (x *NewShortURLResponse) Reset() {
	*x = NewShortURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewShortURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewShortURLResponse) ProtoMessage() {}

func (x *NewShortURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewShortURLResponse.ProtoReflect.Descriptor instead.
func (*NewShortURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_url_service_proto_rawDescGZIP(), []int{1}
}

func (x *NewShortURLResponse) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

type GetLongURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortURL string `protobuf:"bytes,1,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
}

func (x *GetLongURLRequest) Reset() {
	*x = GetLongURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLongURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLongURLRequest) ProtoMessage() {}

func (x *GetLongURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLongURLRequest.ProtoReflect.Descriptor instead.
func (*GetLongURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_url_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetLongURLRequest) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

type GetLongURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LongURL string `protobuf:"bytes,1,opt,name=longURL,proto3" json:"longURL,omitempty"`
}

func (x *GetLongURLResponse) Reset() {
	*x = GetLongURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLongURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLongURLResponse) ProtoMessage() {}

func (x *GetLongURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLongURLResponse.ProtoReflect.Descriptor instead.
func (*GetLongURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_url_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetLongURLResponse) GetLongURL() string {
	if x != nil {
		return x.LongURL
	}
	return ""
}

var File_proto_url_service_proto protoreflect.FileDescriptor

var file_proto_url_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x64, 0x62, 0x22, 0x2e, 0x0a,
	0x12, 0x4e, 0x65, 0x77, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x6e, 0x67, 0x55, 0x52, 0x4c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x6e, 0x67, 0x55, 0x52, 0x4c, 0x22, 0x31, 0x0a,
	0x13, 0x4e, 0x65, 0x77, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c,
	0x22, 0x2f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52,
	0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52,
	0x4c, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x6e, 0x67, 0x55,
	0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x6e, 0x67, 0x55, 0x52,
	0x4c, 0x32, 0x8e, 0x01, 0x0a, 0x0a, 0x55, 0x52, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x43, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x52, 0x4c, 0x12, 0x16, 0x2e, 0x64, 0x62, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x64,
	0x62, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x6e, 0x67,
	0x55, 0x52, 0x4c, 0x12, 0x15, 0x2e, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x6e, 0x67,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x64, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x75, 0x72, 0x6c, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_url_service_proto_rawDescOnce sync.Once
	file_proto_url_service_proto_rawDescData = file_proto_url_service_proto_rawDesc
)

func file_proto_url_service_proto_rawDescGZIP() []byte {
	file_proto_url_service_proto_rawDescOnce.Do(func() {
		file_proto_url_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_url_service_proto_rawDescData)
	})
	return file_proto_url_service_proto_rawDescData
}

var file_proto_url_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_url_service_proto_goTypes = []interface{}{
	(*NewShortURLRequest)(nil),  // 0: db.NewShortURLRequest
	(*NewShortURLResponse)(nil), // 1: db.NewShortURLResponse
	(*GetLongURLRequest)(nil),   // 2: db.GetLongURLRequest
	(*GetLongURLResponse)(nil),  // 3: db.GetLongURLResponse
}
var file_proto_url_service_proto_depIdxs = []int32{
	0, // 0: db.URLService.GenerateShortURL:input_type -> db.NewShortURLRequest
	2, // 1: db.URLService.GetLongURL:input_type -> db.GetLongURLRequest
	1, // 2: db.URLService.GenerateShortURL:output_type -> db.NewShortURLResponse
	3, // 3: db.URLService.GetLongURL:output_type -> db.GetLongURLResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_url_service_proto_init() }
func file_proto_url_service_proto_init() {
	if File_proto_url_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_url_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewShortURLRequest); i {
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
		file_proto_url_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewShortURLResponse); i {
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
		file_proto_url_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLongURLRequest); i {
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
		file_proto_url_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLongURLResponse); i {
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
			RawDescriptor: file_proto_url_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_url_service_proto_goTypes,
		DependencyIndexes: file_proto_url_service_proto_depIdxs,
		MessageInfos:      file_proto_url_service_proto_msgTypes,
	}.Build()
	File_proto_url_service_proto = out.File
	file_proto_url_service_proto_rawDesc = nil
	file_proto_url_service_proto_goTypes = nil
	file_proto_url_service_proto_depIdxs = nil
}
