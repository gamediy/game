// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: user/withdraw/withdraw.proto

package withdraw

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

// pay password status
type PayPassStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *PayPassStatusReq) Reset() {
	*x = PayPassStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_withdraw_withdraw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayPassStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayPassStatusReq) ProtoMessage() {}

func (x *PayPassStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_withdraw_withdraw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayPassStatusReq.ProtoReflect.Descriptor instead.
func (*PayPassStatusReq) Descriptor() ([]byte, []int) {
	return file_user_withdraw_withdraw_proto_rawDescGZIP(), []int{0}
}

func (x *PayPassStatusReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type PayPassStatusRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty" dc:"0 not set , 1 already set"` // 0 not set , 1 already set
}

func (x *PayPassStatusRes) Reset() {
	*x = PayPassStatusRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_withdraw_withdraw_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayPassStatusRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayPassStatusRes) ProtoMessage() {}

func (x *PayPassStatusRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_withdraw_withdraw_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayPassStatusRes.ProtoReflect.Descriptor instead.
func (*PayPassStatusRes) Descriptor() ([]byte, []int) {
	return file_user_withdraw_withdraw_proto_rawDescGZIP(), []int{1}
}

func (x *PayPassStatusRes) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// set pay password
type SetPayPassReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Pass string `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
}

func (x *SetPayPassReq) Reset() {
	*x = SetPayPassReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_withdraw_withdraw_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPayPassReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPayPassReq) ProtoMessage() {}

func (x *SetPayPassReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_withdraw_withdraw_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPayPassReq.ProtoReflect.Descriptor instead.
func (*SetPayPassReq) Descriptor() ([]byte, []int) {
	return file_user_withdraw_withdraw_proto_rawDescGZIP(), []int{2}
}

func (x *SetPayPassReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SetPayPassReq) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

type SetPayPassRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetPayPassRes) Reset() {
	*x = SetPayPassRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_withdraw_withdraw_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPayPassRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPayPassRes) ProtoMessage() {}

func (x *SetPayPassRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_withdraw_withdraw_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPayPassRes.ProtoReflect.Descriptor instead.
func (*SetPayPassRes) Descriptor() ([]byte, []int) {
	return file_user_withdraw_withdraw_proto_rawDescGZIP(), []int{3}
}

// bind withdraw account
type BindWithdrawAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BankId  int64  `protobuf:"varint,1,opt,name=bankId,proto3" json:"bankId,omitempty" dc:"bank id"`      // bank id
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty" dc:"card number"` // card number
	Title   string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty" dc:"cardholder"`      // cardholder
	Pass    string `protobuf:"bytes,4,opt,name=pass,proto3" json:"pass,omitempty" dc:"pay pass"`          // pay pass
	Uid     int64  `protobuf:"varint,5,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *BindWithdrawAccountReq) Reset() {
	*x = BindWithdrawAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_withdraw_withdraw_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindWithdrawAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindWithdrawAccountReq) ProtoMessage() {}

func (x *BindWithdrawAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_withdraw_withdraw_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindWithdrawAccountReq.ProtoReflect.Descriptor instead.
func (*BindWithdrawAccountReq) Descriptor() ([]byte, []int) {
	return file_user_withdraw_withdraw_proto_rawDescGZIP(), []int{4}
}

func (x *BindWithdrawAccountReq) GetBankId() int64 {
	if x != nil {
		return x.BankId
	}
	return 0
}

func (x *BindWithdrawAccountReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BindWithdrawAccountReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BindWithdrawAccountReq) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

func (x *BindWithdrawAccountReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type BindWithdrawAccountRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BindWithdrawAccountRes) Reset() {
	*x = BindWithdrawAccountRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_withdraw_withdraw_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindWithdrawAccountRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindWithdrawAccountRes) ProtoMessage() {}

func (x *BindWithdrawAccountRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_withdraw_withdraw_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindWithdrawAccountRes.ProtoReflect.Descriptor instead.
func (*BindWithdrawAccountRes) Descriptor() ([]byte, []int) {
	return file_user_withdraw_withdraw_proto_rawDescGZIP(), []int{5}
}

// list withdraw account
type ListWithdrawAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Uid  int64 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ListWithdrawAccountReq) Reset() {
	*x = ListWithdrawAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_withdraw_withdraw_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWithdrawAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWithdrawAccountReq) ProtoMessage() {}

func (x *ListWithdrawAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_withdraw_withdraw_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWithdrawAccountReq.ProtoReflect.Descriptor instead.
func (*ListWithdrawAccountReq) Descriptor() ([]byte, []int) {
	return file_user_withdraw_withdraw_proto_rawDescGZIP(), []int{6}
}

func (x *ListWithdrawAccountReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListWithdrawAccountReq) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListWithdrawAccountReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type WithdrawAccountItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Net      string `protobuf:"bytes,2,opt,name=net,proto3" json:"net,omitempty"`
	Protocol string `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Uid      int64  `protobuf:"varint,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Account  string `protobuf:"bytes,5,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *WithdrawAccountItem) Reset() {
	*x = WithdrawAccountItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_withdraw_withdraw_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawAccountItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawAccountItem) ProtoMessage() {}

func (x *WithdrawAccountItem) ProtoReflect() protoreflect.Message {
	mi := &file_user_withdraw_withdraw_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawAccountItem.ProtoReflect.Descriptor instead.
func (*WithdrawAccountItem) Descriptor() ([]byte, []int) {
	return file_user_withdraw_withdraw_proto_rawDescGZIP(), []int{7}
}

func (x *WithdrawAccountItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WithdrawAccountItem) GetNet() string {
	if x != nil {
		return x.Net
	}
	return ""
}

func (x *WithdrawAccountItem) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *WithdrawAccountItem) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *WithdrawAccountItem) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

// create withdraw
type CreateWithdrawReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AmountItemId      int64   `protobuf:"varint,1,opt,name=amountItemId,proto3" json:"amountItemId,omitempty"`
	Amount            float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	WithdrawAccountId int64   `protobuf:"varint,3,opt,name=withdrawAccountId,proto3" json:"withdrawAccountId,omitempty"`
}

func (x *CreateWithdrawReq) Reset() {
	*x = CreateWithdrawReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_withdraw_withdraw_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWithdrawReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWithdrawReq) ProtoMessage() {}

func (x *CreateWithdrawReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_withdraw_withdraw_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWithdrawReq.ProtoReflect.Descriptor instead.
func (*CreateWithdrawReq) Descriptor() ([]byte, []int) {
	return file_user_withdraw_withdraw_proto_rawDescGZIP(), []int{8}
}

func (x *CreateWithdrawReq) GetAmountItemId() int64 {
	if x != nil {
		return x.AmountItemId
	}
	return 0
}

func (x *CreateWithdrawReq) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateWithdrawReq) GetWithdrawAccountId() int64 {
	if x != nil {
		return x.WithdrawAccountId
	}
	return 0
}

type CreateWithdrawRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateWithdrawRes) Reset() {
	*x = CreateWithdrawRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_withdraw_withdraw_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWithdrawRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWithdrawRes) ProtoMessage() {}

func (x *CreateWithdrawRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_withdraw_withdraw_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWithdrawRes.ProtoReflect.Descriptor instead.
func (*CreateWithdrawRes) Descriptor() ([]byte, []int) {
	return file_user_withdraw_withdraw_proto_rawDescGZIP(), []int{9}
}

var File_user_withdraw_withdraw_proto protoreflect.FileDescriptor

var file_user_withdraw_withdraw_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2f,
	0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x22, 0x24, 0x0a, 0x10, 0x50, 0x61, 0x79, 0x50,
	0x61, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x2a,
	0x0a, 0x10, 0x50, 0x61, 0x79, 0x50, 0x61, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x35, 0x0a, 0x0d, 0x53, 0x65,
	0x74, 0x50, 0x61, 0x79, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x73,
	0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x50, 0x61, 0x79, 0x50, 0x61, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x16, 0x42, 0x69, 0x6e, 0x64, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x62,
	0x61, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x42,
	0x69, 0x6e, 0x64, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x22, 0x52, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x7f, 0x0a, 0x13, 0x77, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6e, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e,
	0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7d, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x12,
	0x22, 0x0a, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x77,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x32, 0xc1,
	0x02, 0x0a, 0x0f, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x47, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x50, 0x61, 0x73, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1a, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x50,
	0x61, 0x79, 0x50, 0x61, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a,
	0x1a, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x50, 0x61, 0x79, 0x50, 0x61,
	0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x0a, 0x53,
	0x65, 0x74, 0x50, 0x61, 0x79, 0x50, 0x61, 0x73, 0x73, 0x12, 0x17, 0x2e, 0x77, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x79, 0x50, 0x61, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x17, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x53, 0x65,
	0x74, 0x50, 0x61, 0x79, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x12, 0x59, 0x0a, 0x13, 0x42,
	0x69, 0x6e, 0x64, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x20, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x42, 0x69,
	0x6e, 0x64, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e,
	0x42, 0x69, 0x6e, 0x64, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x1b, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52,
	0x65, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_withdraw_withdraw_proto_rawDescOnce sync.Once
	file_user_withdraw_withdraw_proto_rawDescData = file_user_withdraw_withdraw_proto_rawDesc
)

func file_user_withdraw_withdraw_proto_rawDescGZIP() []byte {
	file_user_withdraw_withdraw_proto_rawDescOnce.Do(func() {
		file_user_withdraw_withdraw_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_withdraw_withdraw_proto_rawDescData)
	})
	return file_user_withdraw_withdraw_proto_rawDescData
}

var file_user_withdraw_withdraw_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_user_withdraw_withdraw_proto_goTypes = []interface{}{
	(*PayPassStatusReq)(nil),       // 0: withdraw.PayPassStatusReq
	(*PayPassStatusRes)(nil),       // 1: withdraw.PayPassStatusRes
	(*SetPayPassReq)(nil),          // 2: withdraw.SetPayPassReq
	(*SetPayPassRes)(nil),          // 3: withdraw.SetPayPassRes
	(*BindWithdrawAccountReq)(nil), // 4: withdraw.BindWithdrawAccountReq
	(*BindWithdrawAccountRes)(nil), // 5: withdraw.BindWithdrawAccountRes
	(*ListWithdrawAccountReq)(nil), // 6: withdraw.ListWithdrawAccountReq
	(*WithdrawAccountItem)(nil),    // 7: withdraw.withdrawAccountItem
	(*CreateWithdrawReq)(nil),      // 8: withdraw.CreateWithdrawReq
	(*CreateWithdrawRes)(nil),      // 9: withdraw.CreateWithdrawRes
}
var file_user_withdraw_withdraw_proto_depIdxs = []int32{
	0, // 0: withdraw.WithdrawService.PayPassStatus:input_type -> withdraw.PayPassStatusReq
	2, // 1: withdraw.WithdrawService.SetPayPass:input_type -> withdraw.SetPayPassReq
	4, // 2: withdraw.WithdrawService.BindWithdrawAccount:input_type -> withdraw.BindWithdrawAccountReq
	8, // 3: withdraw.WithdrawService.CreateWithdraw:input_type -> withdraw.CreateWithdrawReq
	1, // 4: withdraw.WithdrawService.PayPassStatus:output_type -> withdraw.PayPassStatusRes
	3, // 5: withdraw.WithdrawService.SetPayPass:output_type -> withdraw.SetPayPassRes
	5, // 6: withdraw.WithdrawService.BindWithdrawAccount:output_type -> withdraw.BindWithdrawAccountRes
	9, // 7: withdraw.WithdrawService.CreateWithdraw:output_type -> withdraw.CreateWithdrawRes
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_withdraw_withdraw_proto_init() }
func file_user_withdraw_withdraw_proto_init() {
	if File_user_withdraw_withdraw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_withdraw_withdraw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayPassStatusReq); i {
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
		file_user_withdraw_withdraw_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayPassStatusRes); i {
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
		file_user_withdraw_withdraw_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPayPassReq); i {
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
		file_user_withdraw_withdraw_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPayPassRes); i {
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
		file_user_withdraw_withdraw_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindWithdrawAccountReq); i {
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
		file_user_withdraw_withdraw_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindWithdrawAccountRes); i {
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
		file_user_withdraw_withdraw_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWithdrawAccountReq); i {
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
		file_user_withdraw_withdraw_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawAccountItem); i {
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
		file_user_withdraw_withdraw_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWithdrawReq); i {
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
		file_user_withdraw_withdraw_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWithdrawRes); i {
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
			RawDescriptor: file_user_withdraw_withdraw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_withdraw_withdraw_proto_goTypes,
		DependencyIndexes: file_user_withdraw_withdraw_proto_depIdxs,
		MessageInfos:      file_user_withdraw_withdraw_proto_msgTypes,
	}.Build()
	File_user_withdraw_withdraw_proto = out.File
	file_user_withdraw_withdraw_proto_rawDesc = nil
	file_user_withdraw_withdraw_proto_goTypes = nil
	file_user_withdraw_withdraw_proto_depIdxs = nil
}
