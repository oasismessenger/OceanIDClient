// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.0--rc2
// source: ocean_id.proto

package idService

import (
	types "github.com/oasismessenger/OceanIDClient/schemes/types"
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

type IDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DC        types.EnumOceanDC `protobuf:"varint,1,opt,name=DC,proto3,enum=Types.EnumOceanDC" json:"DC,omitempty"`
	WorkerId  uint64            `protobuf:"varint,2,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	RequestId uint64            `protobuf:"varint,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *IDRequest) Reset() {
	*x = IDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocean_id_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDRequest) ProtoMessage() {}

func (x *IDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ocean_id_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDRequest.ProtoReflect.Descriptor instead.
func (*IDRequest) Descriptor() ([]byte, []int) {
	return file_ocean_id_proto_rawDescGZIP(), []int{0}
}

func (x *IDRequest) GetDC() types.EnumOceanDC {
	if x != nil {
		return x.DC
	}
	return types.EnumOceanDC(0)
}

func (x *IDRequest) GetWorkerId() uint64 {
	if x != nil {
		return x.WorkerId
	}
	return 0
}

func (x *IDRequest) GetRequestId() uint64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

type IDBulkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DC        types.EnumOceanDC `protobuf:"varint,1,opt,name=DC,proto3,enum=Types.EnumOceanDC" json:"DC,omitempty"`
	WorkerId  uint32            `protobuf:"varint,2,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	RequestId uint64            `protobuf:"varint,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	BulkSize  uint32            `protobuf:"varint,4,opt,name=bulk_size,json=bulkSize,proto3" json:"bulk_size,omitempty"`
}

func (x *IDBulkRequest) Reset() {
	*x = IDBulkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocean_id_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDBulkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDBulkRequest) ProtoMessage() {}

func (x *IDBulkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ocean_id_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDBulkRequest.ProtoReflect.Descriptor instead.
func (*IDBulkRequest) Descriptor() ([]byte, []int) {
	return file_ocean_id_proto_rawDescGZIP(), []int{1}
}

func (x *IDBulkRequest) GetDC() types.EnumOceanDC {
	if x != nil {
		return x.DC
	}
	return types.EnumOceanDC(0)
}

func (x *IDBulkRequest) GetWorkerId() uint32 {
	if x != nil {
		return x.WorkerId
	}
	return 0
}

func (x *IDBulkRequest) GetRequestId() uint64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *IDBulkRequest) GetBulkSize() uint32 {
	if x != nil {
		return x.BulkSize
	}
	return 0
}

type IDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp uint64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ReplyId   uint64 `protobuf:"varint,3,opt,name=reply_id,json=replyId,proto3" json:"reply_id,omitempty"`
}

func (x *IDReply) Reset() {
	*x = IDReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocean_id_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDReply) ProtoMessage() {}

func (x *IDReply) ProtoReflect() protoreflect.Message {
	mi := &file_ocean_id_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDReply.ProtoReflect.Descriptor instead.
func (*IDReply) Descriptor() ([]byte, []int) {
	return file_ocean_id_proto_rawDescGZIP(), []int{2}
}

func (x *IDReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IDReply) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *IDReply) GetReplyId() uint64 {
	if x != nil {
		return x.ReplyId
	}
	return 0
}

type IDBulkReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids       []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Timestamp uint64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ReplyId   uint64  `protobuf:"varint,3,opt,name=reply_id,json=replyId,proto3" json:"reply_id,omitempty"`
	Size      uint32  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *IDBulkReply) Reset() {
	*x = IDBulkReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocean_id_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDBulkReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDBulkReply) ProtoMessage() {}

func (x *IDBulkReply) ProtoReflect() protoreflect.Message {
	mi := &file_ocean_id_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDBulkReply.ProtoReflect.Descriptor instead.
func (*IDBulkReply) Descriptor() ([]byte, []int) {
	return file_ocean_id_proto_rawDescGZIP(), []int{3}
}

func (x *IDBulkReply) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *IDBulkReply) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *IDBulkReply) GetReplyId() uint64 {
	if x != nil {
		return x.ReplyId
	}
	return 0
}

func (x *IDBulkReply) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_ocean_id_proto protoreflect.FileDescriptor

var file_ocean_id_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x69, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x20, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x5f, 0x69,
	0x64, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a,
	0x09, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x02, 0x44, 0x43,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x4f, 0x63, 0x65, 0x61, 0x6e, 0x44, 0x43, 0x52, 0x02, 0x44, 0x43, 0x12, 0x1b,
	0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x0d, 0x49,
	0x44, 0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x02,
	0x44, 0x43, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x63, 0x65, 0x61, 0x6e, 0x44, 0x43, 0x52, 0x02, 0x44, 0x43,
	0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x75, 0x6c, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x62, 0x75, 0x6c, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x52, 0x0a, 0x07, 0x49, 0x44, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x22, 0x6c, 0x0a,
	0x0b, 0x49, 0x44, 0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08,
	0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x72, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x32, 0x89, 0x01, 0x0a, 0x07,
	0x4f, 0x63, 0x65, 0x61, 0x6e, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x0a, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x49, 0x44, 0x12, 0x14, 0x2e, 0x69, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x69, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x44, 0x0a, 0x0e, 0x42, 0x75, 0x6c, 0x6b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x49, 0x44, 0x12, 0x18, 0x2e, 0x69, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x49, 0x44, 0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x69, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x44, 0x42, 0x75, 0x6c, 0x6b,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x73, 0x2f, 0x69, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x3b, 0x69,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ocean_id_proto_rawDescOnce sync.Once
	file_ocean_id_proto_rawDescData = file_ocean_id_proto_rawDesc
)

func file_ocean_id_proto_rawDescGZIP() []byte {
	file_ocean_id_proto_rawDescOnce.Do(func() {
		file_ocean_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_ocean_id_proto_rawDescData)
	})
	return file_ocean_id_proto_rawDescData
}

var file_ocean_id_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ocean_id_proto_goTypes = []interface{}{
	(*IDRequest)(nil),      // 0: idService.IDRequest
	(*IDBulkRequest)(nil),  // 1: idService.IDBulkRequest
	(*IDReply)(nil),        // 2: idService.IDReply
	(*IDBulkReply)(nil),    // 3: idService.IDBulkReply
	(types.EnumOceanDC)(0), // 4: Types.EnumOceanDC
}
var file_ocean_id_proto_depIdxs = []int32{
	4, // 0: idService.IDRequest.DC:type_name -> Types.EnumOceanDC
	4, // 1: idService.IDBulkRequest.DC:type_name -> Types.EnumOceanDC
	0, // 2: idService.OceanID.GenerateID:input_type -> idService.IDRequest
	1, // 3: idService.OceanID.BulkGenerateID:input_type -> idService.IDBulkRequest
	2, // 4: idService.OceanID.GenerateID:output_type -> idService.IDReply
	3, // 5: idService.OceanID.BulkGenerateID:output_type -> idService.IDBulkReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ocean_id_proto_init() }
func file_ocean_id_proto_init() {
	if File_ocean_id_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ocean_id_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDRequest); i {
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
		file_ocean_id_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDBulkRequest); i {
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
		file_ocean_id_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDReply); i {
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
		file_ocean_id_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDBulkReply); i {
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
			RawDescriptor: file_ocean_id_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ocean_id_proto_goTypes,
		DependencyIndexes: file_ocean_id_proto_depIdxs,
		MessageInfos:      file_ocean_id_proto_msgTypes,
	}.Build()
	File_ocean_id_proto = out.File
	file_ocean_id_proto_rawDesc = nil
	file_ocean_id_proto_goTypes = nil
	file_ocean_id_proto_depIdxs = nil
}
