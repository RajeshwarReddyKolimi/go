// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: protos/fileservice.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UploadFileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filename      string                 `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Data          []byte                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	mi := &file_protos_fileservice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fileservice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_protos_fileservice_proto_rawDescGZIP(), []int{0}
}

func (x *UploadFileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *UploadFileRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type UploadFileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadFileResponse) Reset() {
	*x = UploadFileResponse{}
	mi := &file_protos_fileservice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResponse) ProtoMessage() {}

func (x *UploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fileservice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResponse.ProtoReflect.Descriptor instead.
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return file_protos_fileservice_proto_rawDescGZIP(), []int{1}
}

func (x *UploadFileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DownloadFileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filename      string                 `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Data          []byte                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadFileRequest) Reset() {
	*x = DownloadFileRequest{}
	mi := &file_protos_fileservice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileRequest) ProtoMessage() {}

func (x *DownloadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fileservice_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileRequest.ProtoReflect.Descriptor instead.
func (*DownloadFileRequest) Descriptor() ([]byte, []int) {
	return file_protos_fileservice_proto_rawDescGZIP(), []int{2}
}

func (x *DownloadFileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *DownloadFileRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type DownloadFileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadFileResponse) Reset() {
	*x = DownloadFileResponse{}
	mi := &file_protos_fileservice_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileResponse) ProtoMessage() {}

func (x *DownloadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fileservice_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileResponse.ProtoReflect.Descriptor instead.
func (*DownloadFileResponse) Descriptor() ([]byte, []int) {
	return file_protos_fileservice_proto_rawDescGZIP(), []int{3}
}

func (x *DownloadFileResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type MetaDataRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filename      string                 `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetaDataRequest) Reset() {
	*x = MetaDataRequest{}
	mi := &file_protos_fileservice_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetaDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaDataRequest) ProtoMessage() {}

func (x *MetaDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fileservice_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaDataRequest.ProtoReflect.Descriptor instead.
func (*MetaDataRequest) Descriptor() ([]byte, []int) {
	return file_protos_fileservice_proto_rawDescGZIP(), []int{4}
}

func (x *MetaDataRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type MetaDataResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filename      string                 `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Size          int64                  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	MimeType      string                 `protobuf:"bytes,3,opt,name=mimeType,proto3" json:"mimeType,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetaDataResponse) Reset() {
	*x = MetaDataResponse{}
	mi := &file_protos_fileservice_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetaDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaDataResponse) ProtoMessage() {}

func (x *MetaDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_fileservice_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaDataResponse.ProtoReflect.Descriptor instead.
func (*MetaDataResponse) Descriptor() ([]byte, []int) {
	return file_protos_fileservice_proto_rawDescGZIP(), []int{5}
}

func (x *MetaDataResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *MetaDataResponse) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *MetaDataResponse) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *MetaDataResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_protos_fileservice_proto protoreflect.FileDescriptor

const file_protos_fileservice_proto_rawDesc = "" +
	"\n" +
	"\x18protos/fileservice.proto\x12\vfileservice\"C\n" +
	"\x11UploadFileRequest\x12\x1a\n" +
	"\bfilename\x18\x01 \x01(\tR\bfilename\x12\x12\n" +
	"\x04data\x18\x02 \x01(\fR\x04data\".\n" +
	"\x12UploadFileResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"E\n" +
	"\x13DownloadFileRequest\x12\x1a\n" +
	"\bfilename\x18\x01 \x01(\tR\bfilename\x12\x12\n" +
	"\x04data\x18\x02 \x01(\fR\x04data\"*\n" +
	"\x14DownloadFileResponse\x12\x12\n" +
	"\x04data\x18\x01 \x01(\fR\x04data\"-\n" +
	"\x0fMetaDataRequest\x12\x1a\n" +
	"\bfilename\x18\x01 \x01(\tR\bfilename\"|\n" +
	"\x10MetaDataResponse\x12\x1a\n" +
	"\bfilename\x18\x01 \x01(\tR\bfilename\x12\x12\n" +
	"\x04size\x18\x02 \x01(\x03R\x04size\x12\x1a\n" +
	"\bmimeType\x18\x03 \x01(\tR\bmimeType\x12\x1c\n" +
	"\tcreatedAt\x18\x04 \x01(\tR\tcreatedAt2\x81\x02\n" +
	"\vFileService\x12O\n" +
	"\n" +
	"UploadFile\x12\x1e.fileservice.UploadFileRequest\x1a\x1f.fileservice.UploadFileResponse(\x01\x12U\n" +
	"\fDownloadFile\x12 .fileservice.DownloadFileRequest\x1a!.fileservice.DownloadFileResponse0\x01\x12J\n" +
	"\vGetMetaData\x12\x1c.fileservice.MetaDataRequest\x1a\x1d.fileservice.MetaDataResponseB\tZ\a/protosb\x06proto3"

var (
	file_protos_fileservice_proto_rawDescOnce sync.Once
	file_protos_fileservice_proto_rawDescData []byte
)

func file_protos_fileservice_proto_rawDescGZIP() []byte {
	file_protos_fileservice_proto_rawDescOnce.Do(func() {
		file_protos_fileservice_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protos_fileservice_proto_rawDesc), len(file_protos_fileservice_proto_rawDesc)))
	})
	return file_protos_fileservice_proto_rawDescData
}

var file_protos_fileservice_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_fileservice_proto_goTypes = []any{
	(*UploadFileRequest)(nil),    // 0: fileservice.UploadFileRequest
	(*UploadFileResponse)(nil),   // 1: fileservice.UploadFileResponse
	(*DownloadFileRequest)(nil),  // 2: fileservice.DownloadFileRequest
	(*DownloadFileResponse)(nil), // 3: fileservice.DownloadFileResponse
	(*MetaDataRequest)(nil),      // 4: fileservice.MetaDataRequest
	(*MetaDataResponse)(nil),     // 5: fileservice.MetaDataResponse
}
var file_protos_fileservice_proto_depIdxs = []int32{
	0, // 0: fileservice.FileService.UploadFile:input_type -> fileservice.UploadFileRequest
	2, // 1: fileservice.FileService.DownloadFile:input_type -> fileservice.DownloadFileRequest
	4, // 2: fileservice.FileService.GetMetaData:input_type -> fileservice.MetaDataRequest
	1, // 3: fileservice.FileService.UploadFile:output_type -> fileservice.UploadFileResponse
	3, // 4: fileservice.FileService.DownloadFile:output_type -> fileservice.DownloadFileResponse
	5, // 5: fileservice.FileService.GetMetaData:output_type -> fileservice.MetaDataResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_fileservice_proto_init() }
func file_protos_fileservice_proto_init() {
	if File_protos_fileservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protos_fileservice_proto_rawDesc), len(file_protos_fileservice_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_fileservice_proto_goTypes,
		DependencyIndexes: file_protos_fileservice_proto_depIdxs,
		MessageInfos:      file_protos_fileservice_proto_msgTypes,
	}.Build()
	File_protos_fileservice_proto = out.File
	file_protos_fileservice_proto_goTypes = nil
	file_protos_fileservice_proto_depIdxs = nil
}
