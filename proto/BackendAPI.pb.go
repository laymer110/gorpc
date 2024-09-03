// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: define/BackendAPI.proto

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

	ModelName   string     `protobuf:"bytes,1,opt,name=ModelName,proto3" json:"ModelName,omitempty"` // api名称或者
	Method      string     `protobuf:"bytes,2,opt,name=Method,proto3" json:"Method,omitempty"`       // 增删改查导入导出
	RequestData []byte     `protobuf:"bytes,3,opt,name=RequestData,proto3" json:"RequestData,omitempty"`
	User        *TokenInfo `protobuf:"bytes,4,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *BackendRequest) Reset() {
	*x = BackendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_BackendAPI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackendRequest) ProtoMessage() {}

func (x *BackendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_define_BackendAPI_proto_msgTypes[0]
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
	return file_define_BackendAPI_proto_rawDescGZIP(), []int{0}
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
		mi := &file_define_BackendAPI_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackendResponse) ProtoMessage() {}

func (x *BackendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_define_BackendAPI_proto_msgTypes[1]
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
	return file_define_BackendAPI_proto_rawDescGZIP(), []int{1}
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

type DBTableInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableName   string          `protobuf:"bytes,1,opt,name=TableName,proto3" json:"TableName,omitempty"`
	Description string          `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	Column      []*DBColumnInfo `protobuf:"bytes,3,rep,name=Column,proto3" json:"Column,omitempty"`
}

func (x *DBTableInfo) Reset() {
	*x = DBTableInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_BackendAPI_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBTableInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBTableInfo) ProtoMessage() {}

func (x *DBTableInfo) ProtoReflect() protoreflect.Message {
	mi := &file_define_BackendAPI_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBTableInfo.ProtoReflect.Descriptor instead.
func (*DBTableInfo) Descriptor() ([]byte, []int) {
	return file_define_BackendAPI_proto_rawDescGZIP(), []int{2}
}

func (x *DBTableInfo) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *DBTableInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DBTableInfo) GetColumn() []*DBColumnInfo {
	if x != nil {
		return x.Column
	}
	return nil
}

type DBColumnInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsPrimary    bool   `protobuf:"varint,1,opt,name=IsPrimary,proto3" json:"IsPrimary,omitempty"`
	FieldIndex   int64  `protobuf:"varint,2,opt,name=FieldIndex,proto3" json:"FieldIndex,omitempty"`
	DataType     string `protobuf:"bytes,3,opt,name=DataType,proto3" json:"DataType,omitempty"`
	DataLength   string `protobuf:"bytes,4,opt,name=DataLength,proto3" json:"DataLength,omitempty"` //兼容(8,2)
	ColumnName   string `protobuf:"bytes,5,opt,name=ColumnName,proto3" json:"ColumnName,omitempty"`
	DefaultValue string `protobuf:"bytes,6,opt,name=DefaultValue,proto3" json:"DefaultValue,omitempty"`
	Description  string `protobuf:"bytes,7,opt,name=Description,proto3" json:"Description,omitempty"` //描述
}

func (x *DBColumnInfo) Reset() {
	*x = DBColumnInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_BackendAPI_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBColumnInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBColumnInfo) ProtoMessage() {}

func (x *DBColumnInfo) ProtoReflect() protoreflect.Message {
	mi := &file_define_BackendAPI_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBColumnInfo.ProtoReflect.Descriptor instead.
func (*DBColumnInfo) Descriptor() ([]byte, []int) {
	return file_define_BackendAPI_proto_rawDescGZIP(), []int{3}
}

func (x *DBColumnInfo) GetIsPrimary() bool {
	if x != nil {
		return x.IsPrimary
	}
	return false
}

func (x *DBColumnInfo) GetFieldIndex() int64 {
	if x != nil {
		return x.FieldIndex
	}
	return 0
}

func (x *DBColumnInfo) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *DBColumnInfo) GetDataLength() string {
	if x != nil {
		return x.DataLength
	}
	return ""
}

func (x *DBColumnInfo) GetColumnName() string {
	if x != nil {
		return x.ColumnName
	}
	return ""
}

func (x *DBColumnInfo) GetDefaultValue() string {
	if x != nil {
		return x.DefaultValue
	}
	return ""
}

func (x *DBColumnInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type MigrationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool   `protobuf:"varint,1,opt,name=IsSuccess,proto3" json:"IsSuccess,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *MigrationResult) Reset() {
	*x = MigrationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_BackendAPI_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrationResult) ProtoMessage() {}

func (x *MigrationResult) ProtoReflect() protoreflect.Message {
	mi := &file_define_BackendAPI_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrationResult.ProtoReflect.Descriptor instead.
func (*MigrationResult) Descriptor() ([]byte, []int) {
	return file_define_BackendAPI_proto_rawDescGZIP(), []int{4}
}

func (x *MigrationResult) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *MigrationResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_define_BackendAPI_proto protoreflect.FileDescriptor

var file_define_BackendAPI_proto_rawDesc = []byte{
	0x0a, 0x17, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2f, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x41, 0x50, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x41, 0x50, 0x49, 0x1a, 0x18, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2f, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x94, 0x01, 0x0a, 0x0e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x22, 0x97, 0x01, 0x0a, 0x0f, 0x42, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x61,
	0x74, 0x61, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x44,
	0x61, 0x74, 0x61, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x50, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x7f, 0x0a, 0x0b, 0x44, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1c, 0x0a, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x30, 0x0a, 0x06, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41, 0x50, 0x49, 0x2e, 0x44, 0x42, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x43, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x22, 0xee, 0x01, 0x0a, 0x0c, 0x44, 0x42, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x44, 0x61, 0x74, 0x61, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a,
	0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x49, 0x0a, 0x0f, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa4, 0x01,
	0x0a, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41, 0x50, 0x49, 0x12, 0x43, 0x0a, 0x09,
	0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x42, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x41, 0x50, 0x49, 0x2e, 0x44, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x1b, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41, 0x50, 0x49, 0x2e,
	0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x51, 0x0a, 0x14, 0x43, 0x61, 0x6c, 0x6c, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x1a, 0x2e, 0x42, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x41, 0x50, 0x49, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41,
	0x50, 0x49, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x67, 0x6f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_define_BackendAPI_proto_rawDescOnce sync.Once
	file_define_BackendAPI_proto_rawDescData = file_define_BackendAPI_proto_rawDesc
)

func file_define_BackendAPI_proto_rawDescGZIP() []byte {
	file_define_BackendAPI_proto_rawDescOnce.Do(func() {
		file_define_BackendAPI_proto_rawDescData = protoimpl.X.CompressGZIP(file_define_BackendAPI_proto_rawDescData)
	})
	return file_define_BackendAPI_proto_rawDescData
}

var file_define_BackendAPI_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_define_BackendAPI_proto_goTypes = []any{
	(*BackendRequest)(nil),  // 0: BackendAPI.BackendRequest
	(*BackendResponse)(nil), // 1: BackendAPI.BackendResponse
	(*DBTableInfo)(nil),     // 2: BackendAPI.DBTableInfo
	(*DBColumnInfo)(nil),    // 3: BackendAPI.DBColumnInfo
	(*MigrationResult)(nil), // 4: BackendAPI.MigrationResult
	(*TokenInfo)(nil),       // 5: userManager.TokenInfo
}
var file_define_BackendAPI_proto_depIdxs = []int32{
	5, // 0: BackendAPI.BackendRequest.User:type_name -> userManager.TokenInfo
	3, // 1: BackendAPI.DBTableInfo.Column:type_name -> BackendAPI.DBColumnInfo
	2, // 2: BackendAPI.BackendAPI.Migration:input_type -> BackendAPI.DBTableInfo
	0, // 3: BackendAPI.BackendAPI.CallUnionDataBackend:input_type -> BackendAPI.BackendRequest
	4, // 4: BackendAPI.BackendAPI.Migration:output_type -> BackendAPI.MigrationResult
	1, // 5: BackendAPI.BackendAPI.CallUnionDataBackend:output_type -> BackendAPI.BackendResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_define_BackendAPI_proto_init() }
func file_define_BackendAPI_proto_init() {
	if File_define_BackendAPI_proto != nil {
		return
	}
	file_define_UserManager_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_define_BackendAPI_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_define_BackendAPI_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_define_BackendAPI_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DBTableInfo); i {
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
		file_define_BackendAPI_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DBColumnInfo); i {
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
		file_define_BackendAPI_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MigrationResult); i {
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
			RawDescriptor: file_define_BackendAPI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_define_BackendAPI_proto_goTypes,
		DependencyIndexes: file_define_BackendAPI_proto_depIdxs,
		MessageInfos:      file_define_BackendAPI_proto_msgTypes,
	}.Build()
	File_define_BackendAPI_proto = out.File
	file_define_BackendAPI_proto_rawDesc = nil
	file_define_BackendAPI_proto_goTypes = nil
	file_define_BackendAPI_proto_depIdxs = nil
}
