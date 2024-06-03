// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: BackendAPI/BackendAPI.proto

package gorpc

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

type BackendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelName   string          `protobuf:"bytes,1,opt,name=ModelName,proto3" json:"ModelName,omitempty"` // api名称或者
	Method      string          `protobuf:"bytes,2,opt,name=Method,proto3" json:"Method,omitempty"`       // 增删改查导入导出
	RequestData []byte          `protobuf:"bytes,3,opt,name=RequestData,proto3" json:"RequestData,omitempty"`
	User        *TokenInfo      `protobuf:"bytes,4,opt,name=User,proto3" json:"User,omitempty"`
	Prameter    *SearchPrameter `protobuf:"bytes,5,opt,name=Prameter,proto3" json:"Prameter,omitempty"`
}

func (x *BackendRequest) Reset() {
	*x = BackendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BackendAPI_BackendAPI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackendRequest) ProtoMessage() {}

func (x *BackendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_BackendAPI_BackendAPI_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackendRequest.ProtoReflect.Descriptor instead.
func (*BackendRequest) Descriptor() ([]byte, []int) {
	return file_BackendAPI_BackendAPI_proto_rawDescGZIP(), []int{0}
}

func (x *BackendRequest) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *BackendRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *BackendRequest) GetRequestData() []byte {
	if x != nil {
		return x.RequestData
	}
	return nil
}

func (x *BackendRequest) GetUser() *TokenInfo {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *BackendRequest) GetPrameter() *SearchPrameter {
	if x != nil {
		return x.Prameter
	}
	return nil
}

type BackendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataCount int64  `protobuf:"varint,1,opt,name=DataCount,proto3" json:"DataCount,omitempty"`
	PageCount int64  `protobuf:"varint,2,opt,name=PageCount,proto3" json:"PageCount,omitempty"`
	Code      int64  `protobuf:"varint,3,opt,name=Code,proto3" json:"Code,omitempty"`
	Message   string `protobuf:"bytes,4,opt,name=Message,proto3" json:"Message,omitempty"`
	RespData  []byte `protobuf:"bytes,5,opt,name=RespData,proto3" json:"RespData,omitempty"`
}

func (x *BackendResponse) Reset() {
	*x = BackendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BackendAPI_BackendAPI_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackendResponse) ProtoMessage() {}

func (x *BackendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_BackendAPI_BackendAPI_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackendResponse.ProtoReflect.Descriptor instead.
func (*BackendResponse) Descriptor() ([]byte, []int) {
	return file_BackendAPI_BackendAPI_proto_rawDescGZIP(), []int{1}
}

func (x *BackendResponse) GetDataCount() int64 {
	if x != nil {
		return x.DataCount
	}
	return 0
}

func (x *BackendResponse) GetPageCount() int64 {
	if x != nil {
		return x.PageCount
	}
	return 0
}

func (x *BackendResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BackendResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BackendResponse) GetRespData() []byte {
	if x != nil {
		return x.RespData
	}
	return nil
}

type SearchPrameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize   int64 `protobuf:"varint,1,opt,name=PageSize,proto3" json:"PageSize,omitempty"`     //每页数量
	PageNumber int64 `protobuf:"varint,2,opt,name=PageNumber,proto3" json:"PageNumber,omitempty"` //第几页
	PageCount  int64 `protobuf:"varint,3,opt,name=PageCount,proto3" json:"PageCount,omitempty"`   //共有多少页
}

func (x *SearchPrameter) Reset() {
	*x = SearchPrameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BackendAPI_BackendAPI_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPrameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPrameter) ProtoMessage() {}

func (x *SearchPrameter) ProtoReflect() protoreflect.Message {
	mi := &file_BackendAPI_BackendAPI_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPrameter.ProtoReflect.Descriptor instead.
func (*SearchPrameter) Descriptor() ([]byte, []int) {
	return file_BackendAPI_BackendAPI_proto_rawDescGZIP(), []int{2}
}

func (x *SearchPrameter) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchPrameter) GetPageNumber() int64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *SearchPrameter) GetPageCount() int64 {
	if x != nil {
		return x.PageCount
	}
	return 0
}

var File_BackendAPI_BackendAPI_proto protoreflect.FileDescriptor

var file_BackendAPI_BackendAPI_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41, 0x50, 0x49, 0x2f, 0x42, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x41, 0x50, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x42,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41, 0x50, 0x49, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x0e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x50, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x42, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41, 0x50, 0x49, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x08, 0x50, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x22, 0x97, 0x01, 0x0a, 0x0f, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x50, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x22, 0x6a, 0x0a, 0x0e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x67,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x50,
	0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x50, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x5f, 0x0a, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x41, 0x50, 0x49, 0x12, 0x51, 0x0a, 0x14, 0x43, 0x61, 0x6c, 0x6c, 0x55, 0x6e, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x1a, 0x2e,
	0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41, 0x50, 0x49, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x42, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x41, 0x50, 0x49, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_BackendAPI_BackendAPI_proto_rawDescOnce sync.Once
	file_BackendAPI_BackendAPI_proto_rawDescData = file_BackendAPI_BackendAPI_proto_rawDesc
)

func file_BackendAPI_BackendAPI_proto_rawDescGZIP() []byte {
	file_BackendAPI_BackendAPI_proto_rawDescOnce.Do(func() {
		file_BackendAPI_BackendAPI_proto_rawDescData = protoimpl.X.CompressGZIP(file_BackendAPI_BackendAPI_proto_rawDescData)
	})
	return file_BackendAPI_BackendAPI_proto_rawDescData
}

var file_BackendAPI_BackendAPI_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_BackendAPI_BackendAPI_proto_goTypes = []interface{}{
	(*BackendRequest)(nil),  // 0: BackendAPI.BackendRequest
	(*BackendResponse)(nil), // 1: BackendAPI.BackendResponse
	(*SearchPrameter)(nil),  // 2: BackendAPI.SearchPrameter
	(*TokenInfo)(nil),       // 3: userManager.TokenInfo
}
var file_BackendAPI_BackendAPI_proto_depIdxs = []int32{
	3, // 0: BackendAPI.BackendRequest.User:type_name -> userManager.TokenInfo
	2, // 1: BackendAPI.BackendRequest.Prameter:type_name -> BackendAPI.SearchPrameter
	0, // 2: BackendAPI.BackendAPI.CallUnionDataBackend:input_type -> BackendAPI.BackendRequest
	1, // 3: BackendAPI.BackendAPI.CallUnionDataBackend:output_type -> BackendAPI.BackendResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_BackendAPI_BackendAPI_proto_init() }
func file_BackendAPI_BackendAPI_proto_init() {
	if File_BackendAPI_BackendAPI_proto != nil {
		return
	}
	file_proto_UserManager_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BackendAPI_BackendAPI_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackendRequest); i {
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
		file_BackendAPI_BackendAPI_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackendResponse); i {
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
		file_BackendAPI_BackendAPI_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPrameter); i {
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
			RawDescriptor: file_BackendAPI_BackendAPI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_BackendAPI_BackendAPI_proto_goTypes,
		DependencyIndexes: file_BackendAPI_BackendAPI_proto_depIdxs,
		MessageInfos:      file_BackendAPI_BackendAPI_proto_msgTypes,
	}.Build()
	File_BackendAPI_BackendAPI_proto = out.File
	file_BackendAPI_BackendAPI_proto_rawDesc = nil
	file_BackendAPI_BackendAPI_proto_goTypes = nil
	file_BackendAPI_BackendAPI_proto_depIdxs = nil
}
