// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: ledger/queryresult/kv_query_result.proto

package queryresult

import (
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

// KV -- QueryResult for range/execute query. Holds a key and corresponding value.
type KV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value     []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KV) Reset() {
	*x = KV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_queryresult_kv_query_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KV) ProtoMessage() {}

func (x *KV) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_queryresult_kv_query_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KV.ProtoReflect.Descriptor instead.
func (*KV) Descriptor() ([]byte, []int) {
	return file_ledger_queryresult_kv_query_result_proto_rawDescGZIP(), []int{0}
}

func (x *KV) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *KV) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KV) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// KeyModification -- QueryResult for history query. Holds a transaction ID, value,
// timestamp, and delete marker which resulted from a history query.
type KeyModification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId      string                 `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	Value     []byte                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	IsDelete  bool                   `protobuf:"varint,4,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
}

func (x *KeyModification) Reset() {
	*x = KeyModification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_queryresult_kv_query_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyModification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyModification) ProtoMessage() {}

func (x *KeyModification) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_queryresult_kv_query_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyModification.ProtoReflect.Descriptor instead.
func (*KeyModification) Descriptor() ([]byte, []int) {
	return file_ledger_queryresult_kv_query_result_proto_rawDescGZIP(), []int{1}
}

func (x *KeyModification) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *KeyModification) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *KeyModification) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *KeyModification) GetIsDelete() bool {
	if x != nil {
		return x.IsDelete
	}
	return false
}

var File_ledger_queryresult_kv_query_result_proto protoreflect.FileDescriptor

var file_ledger_queryresult_kv_query_result_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2f, 0x6b, 0x76, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x02, 0x4b, 0x56, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x0f, 0x4b, 0x65, 0x79, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x78, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0xd4, 0x01, 0x0a, 0x30, 0x6f,
	0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x66,
	0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42,
	0x12, 0x4b, 0x56, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66, 0x61,
	0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0xa2, 0x02, 0x03, 0x51, 0x58, 0x58, 0xaa, 0x02, 0x0b,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0xca, 0x02, 0x0b, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0xe2, 0x02, 0x17, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ledger_queryresult_kv_query_result_proto_rawDescOnce sync.Once
	file_ledger_queryresult_kv_query_result_proto_rawDescData = file_ledger_queryresult_kv_query_result_proto_rawDesc
)

func file_ledger_queryresult_kv_query_result_proto_rawDescGZIP() []byte {
	file_ledger_queryresult_kv_query_result_proto_rawDescOnce.Do(func() {
		file_ledger_queryresult_kv_query_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_ledger_queryresult_kv_query_result_proto_rawDescData)
	})
	return file_ledger_queryresult_kv_query_result_proto_rawDescData
}

var file_ledger_queryresult_kv_query_result_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ledger_queryresult_kv_query_result_proto_goTypes = []interface{}{
	(*KV)(nil),                    // 0: queryresult.KV
	(*KeyModification)(nil),       // 1: queryresult.KeyModification
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_ledger_queryresult_kv_query_result_proto_depIdxs = []int32{
	2, // 0: queryresult.KeyModification.timestamp:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ledger_queryresult_kv_query_result_proto_init() }
func file_ledger_queryresult_kv_query_result_proto_init() {
	if File_ledger_queryresult_kv_query_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ledger_queryresult_kv_query_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KV); i {
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
		file_ledger_queryresult_kv_query_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyModification); i {
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
			RawDescriptor: file_ledger_queryresult_kv_query_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ledger_queryresult_kv_query_result_proto_goTypes,
		DependencyIndexes: file_ledger_queryresult_kv_query_result_proto_depIdxs,
		MessageInfos:      file_ledger_queryresult_kv_query_result_proto_msgTypes,
	}.Build()
	File_ledger_queryresult_kv_query_result_proto = out.File
	file_ledger_queryresult_kv_query_result_proto_rawDesc = nil
	file_ledger_queryresult_kv_query_result_proto_goTypes = nil
	file_ledger_queryresult_kv_query_result_proto_depIdxs = nil
}
