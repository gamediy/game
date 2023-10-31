// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: user/wallet/walletChangeLog.proto

package wallet

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListChangeLogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Page      int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size      int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	TransCode string `protobuf:"bytes,4,opt,name=transCode,proto3" json:"transCode,omitempty"`
}

func (x *ListChangeLogReq) Reset() {
	*x = ListChangeLogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_wallet_walletChangeLog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChangeLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChangeLogReq) ProtoMessage() {}

func (x *ListChangeLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_wallet_walletChangeLog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChangeLogReq.ProtoReflect.Descriptor instead.
func (*ListChangeLogReq) Descriptor() ([]byte, []int) {
	return file_user_wallet_walletChangeLog_proto_rawDescGZIP(), []int{0}
}

func (x *ListChangeLogReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ListChangeLogReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListChangeLogReq) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListChangeLogReq) GetTransCode() string {
	if x != nil {
		return x.TransCode
	}
	return ""
}

type ChangeLogItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TransId       string  `protobuf:"bytes,2,opt,name=transId,proto3" json:"transId,omitempty"`
	Uid           int64   `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	TransCode     int64   `protobuf:"varint,4,opt,name=transCode,proto3" json:"transCode,omitempty"`
	Title         string  `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	BalanceBefore float64 `protobuf:"fixed64,6,opt,name=balanceBefore,proto3" json:"balanceBefore,omitempty"`
	BalanceAfter  float64 `protobuf:"fixed64,7,opt,name=balanceAfter,proto3" json:"balanceAfter,omitempty"`
	Amount        float64 `protobuf:"fixed64,8,opt,name=amount,proto3" json:"amount,omitempty"`
	Note          string  `protobuf:"bytes,9,opt,name=note,proto3" json:"note,omitempty"`
	CreatedAt     string  `protobuf:"bytes,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *ChangeLogItem) Reset() {
	*x = ChangeLogItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_wallet_walletChangeLog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeLogItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeLogItem) ProtoMessage() {}

func (x *ChangeLogItem) ProtoReflect() protoreflect.Message {
	mi := &file_user_wallet_walletChangeLog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeLogItem.ProtoReflect.Descriptor instead.
func (*ChangeLogItem) Descriptor() ([]byte, []int) {
	return file_user_wallet_walletChangeLog_proto_rawDescGZIP(), []int{1}
}

func (x *ChangeLogItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChangeLogItem) GetTransId() string {
	if x != nil {
		return x.TransId
	}
	return ""
}

func (x *ChangeLogItem) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ChangeLogItem) GetTransCode() int64 {
	if x != nil {
		return x.TransCode
	}
	return 0
}

func (x *ChangeLogItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ChangeLogItem) GetBalanceBefore() float64 {
	if x != nil {
		return x.BalanceBefore
	}
	return 0
}

func (x *ChangeLogItem) GetBalanceAfter() float64 {
	if x != nil {
		return x.BalanceAfter
	}
	return 0
}

func (x *ChangeLogItem) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ChangeLogItem) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *ChangeLogItem) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type ListChangeLogRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64            `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*ChangeLogItem `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListChangeLogRes) Reset() {
	*x = ListChangeLogRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_wallet_walletChangeLog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChangeLogRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChangeLogRes) ProtoMessage() {}

func (x *ListChangeLogRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_wallet_walletChangeLog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChangeLogRes.ProtoReflect.Descriptor instead.
func (*ListChangeLogRes) Descriptor() ([]byte, []int) {
	return file_user_wallet_walletChangeLog_proto_rawDescGZIP(), []int{2}
}

func (x *ListChangeLogRes) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListChangeLogRes) GetList() []*ChangeLogItem {
	if x != nil {
		return x.List
	}
	return nil
}

var File_user_wallet_walletChangeLog_proto protoreflect.FileDescriptor

var file_user_wallet_walletChangeLog_proto_rawDesc = []byte{
	0x0a, 0x21, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x22, 0x6a, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x93, 0x02, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x4c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x66,
	0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x6f, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x53, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x29, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x32, 0x54, 0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x18,
	0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_wallet_walletChangeLog_proto_rawDescOnce sync.Once
	file_user_wallet_walletChangeLog_proto_rawDescData = file_user_wallet_walletChangeLog_proto_rawDesc
)

func file_user_wallet_walletChangeLog_proto_rawDescGZIP() []byte {
	file_user_wallet_walletChangeLog_proto_rawDescOnce.Do(func() {
		file_user_wallet_walletChangeLog_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_wallet_walletChangeLog_proto_rawDescData)
	})
	return file_user_wallet_walletChangeLog_proto_rawDescData
}

var file_user_wallet_walletChangeLog_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_user_wallet_walletChangeLog_proto_goTypes = []interface{}{
	(*ListChangeLogReq)(nil), // 0: wallet.ListChangeLogReq
	(*ChangeLogItem)(nil),    // 1: wallet.ChangeLogItem
	(*ListChangeLogRes)(nil), // 2: wallet.ListChangeLogRes
}
var file_user_wallet_walletChangeLog_proto_depIdxs = []int32{
	1, // 0: wallet.ListChangeLogRes.list:type_name -> wallet.ChangeLogItem
	0, // 1: wallet.WalletService.ListChangeLog:input_type -> wallet.ListChangeLogReq
	2, // 2: wallet.WalletService.ListChangeLog:output_type -> wallet.ListChangeLogRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_wallet_walletChangeLog_proto_init() }
func file_user_wallet_walletChangeLog_proto_init() {
	if File_user_wallet_walletChangeLog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_wallet_walletChangeLog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChangeLogReq); i {
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
		file_user_wallet_walletChangeLog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeLogItem); i {
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
		file_user_wallet_walletChangeLog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChangeLogRes); i {
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
			RawDescriptor: file_user_wallet_walletChangeLog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_wallet_walletChangeLog_proto_goTypes,
		DependencyIndexes: file_user_wallet_walletChangeLog_proto_depIdxs,
		MessageInfos:      file_user_wallet_walletChangeLog_proto_msgTypes,
	}.Build()
	File_user_wallet_walletChangeLog_proto = out.File
	file_user_wallet_walletChangeLog_proto_rawDesc = nil
	file_user_wallet_walletChangeLog_proto_goTypes = nil
	file_user_wallet_walletChangeLog_proto_depIdxs = nil
}
