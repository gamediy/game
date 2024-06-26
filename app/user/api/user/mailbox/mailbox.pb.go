// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: user/mailbox/mailbox.proto

package mailbox

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

type ListMailBoxReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size     int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Receiver string `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Read     string `protobuf:"bytes,4,opt,name=read,proto3" json:"read,omitempty"`
	Type     string `protobuf:"bytes,5,opt,name=implement,proto3" json:"implement,omitempty" dc:"0 普通 1 赠送"` // 0 普通 1 赠送
}

func (x *ListMailBoxReq) Reset() {
	*x = ListMailBoxReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_mailbox_mailbox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMailBoxReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMailBoxReq) ProtoMessage() {}

func (x *ListMailBoxReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_mailbox_mailbox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMailBoxReq.ProtoReflect.Descriptor instead.
func (*ListMailBoxReq) Descriptor() ([]byte, []int) {
	return file_user_mailbox_mailbox_proto_rawDescGZIP(), []int{0}
}

func (x *ListMailBoxReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListMailBoxReq) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListMailBoxReq) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *ListMailBoxReq) GetRead() string {
	if x != nil {
		return x.Read
	}
	return ""
}

func (x *ListMailBoxReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ListMailBoxRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64      `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*MailBox `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListMailBoxRes) Reset() {
	*x = ListMailBoxRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_mailbox_mailbox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMailBoxRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMailBoxRes) ProtoMessage() {}

func (x *ListMailBoxRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_mailbox_mailbox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMailBoxRes.ProtoReflect.Descriptor instead.
func (*ListMailBoxRes) Descriptor() ([]byte, []int) {
	return file_user_mailbox_mailbox_proto_rawDescGZIP(), []int{1}
}

func (x *ListMailBoxRes) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListMailBoxRes) GetList() []*MailBox {
	if x != nil {
		return x.List
	}
	return nil
}

type MailBox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Receiver  int64  `protobuf:"varint,4,opt,name=receiver,proto3" json:"receiver,omitempty"`
	ReadTime  string `protobuf:"bytes,5,opt,name=readTime,proto3" json:"readTime,omitempty"`
	ReadStart string `protobuf:"bytes,6,opt,name=readStart,proto3" json:"readStart,omitempty"`
	Read      int64  `protobuf:"varint,7,opt,name=read,proto3" json:"read,omitempty"`
	Type      int64  `protobuf:"varint,8,opt,name=implement,proto3" json:"implement,omitempty" dc:"0 普通 1赠送"` // 0 普通 1赠送
}

func (x *MailBox) Reset() {
	*x = MailBox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_mailbox_mailbox_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailBox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailBox) ProtoMessage() {}

func (x *MailBox) ProtoReflect() protoreflect.Message {
	mi := &file_user_mailbox_mailbox_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailBox.ProtoReflect.Descriptor instead.
func (*MailBox) Descriptor() ([]byte, []int) {
	return file_user_mailbox_mailbox_proto_rawDescGZIP(), []int{2}
}

func (x *MailBox) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MailBox) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MailBox) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MailBox) GetReceiver() int64 {
	if x != nil {
		return x.Receiver
	}
	return 0
}

func (x *MailBox) GetReadTime() string {
	if x != nil {
		return x.ReadTime
	}
	return ""
}

func (x *MailBox) GetReadStart() string {
	if x != nil {
		return x.ReadStart
	}
	return ""
}

func (x *MailBox) GetRead() int64 {
	if x != nil {
		return x.Read
	}
	return 0
}

func (x *MailBox) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

type MailBoxTotalUnReadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *MailBoxTotalUnReadReq) Reset() {
	*x = MailBoxTotalUnReadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_mailbox_mailbox_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailBoxTotalUnReadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailBoxTotalUnReadReq) ProtoMessage() {}

func (x *MailBoxTotalUnReadReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_mailbox_mailbox_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailBoxTotalUnReadReq.ProtoReflect.Descriptor instead.
func (*MailBoxTotalUnReadReq) Descriptor() ([]byte, []int) {
	return file_user_mailbox_mailbox_proto_rawDescGZIP(), []int{3}
}

func (x *MailBoxTotalUnReadReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type MailBoxTotalUnReadRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *MailBoxTotalUnReadRes) Reset() {
	*x = MailBoxTotalUnReadRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_mailbox_mailbox_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailBoxTotalUnReadRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailBoxTotalUnReadRes) ProtoMessage() {}

func (x *MailBoxTotalUnReadRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_mailbox_mailbox_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailBoxTotalUnReadRes.ProtoReflect.Descriptor instead.
func (*MailBoxTotalUnReadRes) Descriptor() ([]byte, []int) {
	return file_user_mailbox_mailbox_proto_rawDescGZIP(), []int{4}
}

func (x *MailBoxTotalUnReadRes) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_user_mailbox_mailbox_proto protoreflect.FileDescriptor

var file_user_mailbox_mailbox_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x2f, 0x6d,
	0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x61,
	0x69, 0x6c, 0x62, 0x6f, 0x78, 0x22, 0x7c, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x69,
	0x6c, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x65, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x61, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x4c, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x42,
	0x6f, 0x78, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6c,
	0x62, 0x6f, 0x78, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0xc7, 0x01, 0x0a, 0x07, 0x4d, 0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x61, 0x64, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x72, 0x65, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x29, 0x0a, 0x15, 0x4d,
	0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x6e, 0x52, 0x65, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x15, 0x4d, 0x61, 0x69, 0x6c, 0x42, 0x6f,
	0x78, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75,
	0x6d, 0x32, 0xac, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x69, 0x6c,
	0x42, 0x6f, 0x78, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x6d,
	0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x42,
	0x6f, 0x78, 0x52, 0x65, 0x73, 0x12, 0x59, 0x0a, 0x17, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x61,
	0x69, 0x6c, 0x42, 0x6f, 0x78, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x6e, 0x52, 0x65, 0x61, 0x64,
	0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x42,
	0x6f, 0x78, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x1e, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x42,
	0x6f, 0x78, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_mailbox_mailbox_proto_rawDescOnce sync.Once
	file_user_mailbox_mailbox_proto_rawDescData = file_user_mailbox_mailbox_proto_rawDesc
)

func file_user_mailbox_mailbox_proto_rawDescGZIP() []byte {
	file_user_mailbox_mailbox_proto_rawDescOnce.Do(func() {
		file_user_mailbox_mailbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_mailbox_mailbox_proto_rawDescData)
	})
	return file_user_mailbox_mailbox_proto_rawDescData
}

var file_user_mailbox_mailbox_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_mailbox_mailbox_proto_goTypes = []interface{}{
	(*ListMailBoxReq)(nil),        // 0: mailbox.ListMailBoxReq
	(*ListMailBoxRes)(nil),        // 1: mailbox.ListMailBoxRes
	(*MailBox)(nil),               // 2: mailbox.MailBox
	(*MailBoxTotalUnReadReq)(nil), // 3: mailbox.MailBoxTotalUnReadReq
	(*MailBoxTotalUnReadRes)(nil), // 4: mailbox.MailBoxTotalUnReadRes
}
var file_user_mailbox_mailbox_proto_depIdxs = []int32{
	2, // 0: mailbox.ListMailBoxRes.list:type_name -> mailbox.MailBox
	0, // 1: mailbox.MailBoxService.ListMailBox:input_type -> mailbox.ListMailBoxReq
	3, // 2: mailbox.MailBoxService.CountMailBoxTotalUnRead:input_type -> mailbox.MailBoxTotalUnReadReq
	1, // 3: mailbox.MailBoxService.ListMailBox:output_type -> mailbox.ListMailBoxRes
	4, // 4: mailbox.MailBoxService.CountMailBoxTotalUnRead:output_type -> mailbox.MailBoxTotalUnReadRes
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_mailbox_mailbox_proto_init() }
func file_user_mailbox_mailbox_proto_init() {
	if File_user_mailbox_mailbox_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_mailbox_mailbox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMailBoxReq); i {
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
		file_user_mailbox_mailbox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMailBoxRes); i {
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
		file_user_mailbox_mailbox_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailBox); i {
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
		file_user_mailbox_mailbox_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailBoxTotalUnReadReq); i {
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
		file_user_mailbox_mailbox_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailBoxTotalUnReadRes); i {
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
			RawDescriptor: file_user_mailbox_mailbox_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_mailbox_mailbox_proto_goTypes,
		DependencyIndexes: file_user_mailbox_mailbox_proto_depIdxs,
		MessageInfos:      file_user_mailbox_mailbox_proto_msgTypes,
	}.Build()
	File_user_mailbox_mailbox_proto = out.File
	file_user_mailbox_mailbox_proto_rawDesc = nil
	file_user_mailbox_mailbox_proto_goTypes = nil
	file_user_mailbox_mailbox_proto_depIdxs = nil
}
