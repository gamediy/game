// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: user/deposit/deposit.proto

package deposit

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

type DepositAmountItemsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *DepositAmountItemsReq) Reset() {
	*x = DepositAmountItemsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_deposit_deposit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositAmountItemsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositAmountItemsReq) ProtoMessage() {}

func (x *DepositAmountItemsReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_deposit_deposit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositAmountItemsReq.ProtoReflect.Descriptor instead.
func (*DepositAmountItemsReq) Descriptor() ([]byte, []int) {
	return file_user_deposit_deposit_proto_rawDescGZIP(), []int{0}
}

func (x *DepositAmountItemsReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type AmountItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title            string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Status           int64   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Detail           string  `protobuf:"bytes,4,opt,name=detail,proto3" json:"detail,omitempty"`
	AmountCategoryId int64   `protobuf:"varint,5,opt,name=amount_category_id,json=amountCategoryId,proto3" json:"amount_category_id,omitempty"`
	Net              string  `protobuf:"bytes,6,opt,name=net,proto3" json:"net,omitempty"`
	Min              float64 `protobuf:"fixed64,7,opt,name=min,proto3" json:"min,omitempty"`
	Max              float64 `protobuf:"fixed64,8,opt,name=max,proto3" json:"max,omitempty"`
	Fee              float64 `protobuf:"fixed64,9,opt,name=fee,proto3" json:"fee,omitempty"`
	Type             int64   `protobuf:"varint,10,opt,name=type,proto3" json:"type,omitempty"`
	Logo             string  `protobuf:"bytes,11,opt,name=logo,proto3" json:"logo,omitempty"`
	Sort             int64   `protobuf:"varint,12,opt,name=sort,proto3" json:"sort,omitempty"`
	Country          string  `protobuf:"bytes,13,opt,name=country,proto3" json:"country,omitempty"`
	Currency         string  `protobuf:"bytes,14,opt,name=currency,proto3" json:"currency,omitempty"`
	Protocol         string  `protobuf:"bytes,15,opt,name=protocol,proto3" json:"protocol,omitempty"`
	CreatedAt        string  `protobuf:"bytes,16,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        string  `protobuf:"bytes,17,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Address          string  `protobuf:"bytes,18,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *AmountItem) Reset() {
	*x = AmountItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_deposit_deposit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmountItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmountItem) ProtoMessage() {}

func (x *AmountItem) ProtoReflect() protoreflect.Message {
	mi := &file_user_deposit_deposit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmountItem.ProtoReflect.Descriptor instead.
func (*AmountItem) Descriptor() ([]byte, []int) {
	return file_user_deposit_deposit_proto_rawDescGZIP(), []int{1}
}

func (x *AmountItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AmountItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AmountItem) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AmountItem) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *AmountItem) GetAmountCategoryId() int64 {
	if x != nil {
		return x.AmountCategoryId
	}
	return 0
}

func (x *AmountItem) GetNet() string {
	if x != nil {
		return x.Net
	}
	return ""
}

func (x *AmountItem) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *AmountItem) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *AmountItem) GetFee() float64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *AmountItem) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *AmountItem) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *AmountItem) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *AmountItem) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *AmountItem) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *AmountItem) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *AmountItem) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AmountItem) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *AmountItem) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type DepositBox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string        `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	CategoryId int64         `protobuf:"varint,2,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	Item       []*AmountItem `protobuf:"bytes,3,rep,name=item,proto3" json:"item,omitempty"`
}

func (x *DepositBox) Reset() {
	*x = DepositBox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_deposit_deposit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositBox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositBox) ProtoMessage() {}

func (x *DepositBox) ProtoReflect() protoreflect.Message {
	mi := &file_user_deposit_deposit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositBox.ProtoReflect.Descriptor instead.
func (*DepositBox) Descriptor() ([]byte, []int) {
	return file_user_deposit_deposit_proto_rawDescGZIP(), []int{2}
}

func (x *DepositBox) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DepositBox) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *DepositBox) GetItem() []*AmountItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type DepositAmountItemsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tips string        `protobuf:"bytes,1,opt,name=tips,proto3" json:"tips,omitempty"`
	List []*DepositBox `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *DepositAmountItemsRes) Reset() {
	*x = DepositAmountItemsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_deposit_deposit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositAmountItemsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositAmountItemsRes) ProtoMessage() {}

func (x *DepositAmountItemsRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_deposit_deposit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositAmountItemsRes.ProtoReflect.Descriptor instead.
func (*DepositAmountItemsRes) Descriptor() ([]byte, []int) {
	return file_user_deposit_deposit_proto_rawDescGZIP(), []int{3}
}

func (x *DepositAmountItemsRes) GetTips() string {
	if x != nil {
		return x.Tips
	}
	return ""
}

func (x *DepositAmountItemsRes) GetList() []*DepositBox {
	if x != nil {
		return x.List
	}
	return nil
}

// create deposit
type CreateDepositReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayId           int64   `protobuf:"varint,1,opt,name=payId,proto3" json:"payId,omitempty" dc:"充值类型ID"` // 充值类型ID
	Amount          float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	TransferOrderNo string  `protobuf:"bytes,3,opt,name=transferOrderNo,proto3" json:"transferOrderNo,omitempty" dc:"交易成功的订单号"` // 交易成功的订单号
	TransferImg     string  `protobuf:"bytes,4,opt,name=transferImg,proto3" json:"transferImg,omitempty" dc:"凭证"`               // 凭证
}

func (x *CreateDepositReq) Reset() {
	*x = CreateDepositReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_deposit_deposit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDepositReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDepositReq) ProtoMessage() {}

func (x *CreateDepositReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_deposit_deposit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDepositReq.ProtoReflect.Descriptor instead.
func (*CreateDepositReq) Descriptor() ([]byte, []int) {
	return file_user_deposit_deposit_proto_rawDescGZIP(), []int{4}
}

func (x *CreateDepositReq) GetPayId() int64 {
	if x != nil {
		return x.PayId
	}
	return 0
}

func (x *CreateDepositReq) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateDepositReq) GetTransferOrderNo() string {
	if x != nil {
		return x.TransferOrderNo
	}
	return ""
}

func (x *CreateDepositReq) GetTransferImg() string {
	if x != nil {
		return x.TransferImg
	}
	return ""
}

type CreateDepositRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderNo string `protobuf:"bytes,1,opt,name=orderNo,proto3" json:"orderNo,omitempty" dc:"交易ID"` // 交易ID
}

func (x *CreateDepositRes) Reset() {
	*x = CreateDepositRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_deposit_deposit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDepositRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDepositRes) ProtoMessage() {}

func (x *CreateDepositRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_deposit_deposit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDepositRes.ProtoReflect.Descriptor instead.
func (*CreateDepositRes) Descriptor() ([]byte, []int) {
	return file_user_deposit_deposit_proto_rawDescGZIP(), []int{5}
}

func (x *CreateDepositRes) GetOrderNo() string {
	if x != nil {
		return x.OrderNo
	}
	return ""
}

var File_user_deposit_deposit_proto protoreflect.FileDescriptor

var file_user_deposit_deposit_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2f, 0x64,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x22, 0x29, 0x0a, 0x15, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x22, 0xbe, 0x03, 0x0a, 0x0a, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6e, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c,
	0x6f, 0x67, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x6b, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x42, 0x6f, 0x78, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x54,
	0x0a, 0x15, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x70, 0x73, 0x12, 0x27, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x42, 0x6f, 0x78, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x8c, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x79,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x79, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e,
	0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x6d, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x49, 0x6d, 0x67, 0x22, 0x2c, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e,
	0x6f, 0x32, 0xb1, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1e,
	0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1e,
	0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x12, 0x45,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x12,
	0x19, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x64, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x64, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_deposit_deposit_proto_rawDescOnce sync.Once
	file_user_deposit_deposit_proto_rawDescData = file_user_deposit_deposit_proto_rawDesc
)

func file_user_deposit_deposit_proto_rawDescGZIP() []byte {
	file_user_deposit_deposit_proto_rawDescOnce.Do(func() {
		file_user_deposit_deposit_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_deposit_deposit_proto_rawDescData)
	})
	return file_user_deposit_deposit_proto_rawDescData
}

var file_user_deposit_deposit_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_deposit_deposit_proto_goTypes = []interface{}{
	(*DepositAmountItemsReq)(nil), // 0: deposit.DepositAmountItemsReq
	(*AmountItem)(nil),            // 1: deposit.AmountItem
	(*DepositBox)(nil),            // 2: deposit.depositBox
	(*DepositAmountItemsRes)(nil), // 3: deposit.DepositAmountItemsRes
	(*CreateDepositReq)(nil),      // 4: deposit.CreateDepositReq
	(*CreateDepositRes)(nil),      // 5: deposit.CreateDepositRes
}
var file_user_deposit_deposit_proto_depIdxs = []int32{
	1, // 0: deposit.depositBox.item:type_name -> deposit.AmountItem
	2, // 1: deposit.DepositAmountItemsRes.list:type_name -> deposit.depositBox
	0, // 2: deposit.DepositService.ListDepositAmountItems:input_type -> deposit.DepositAmountItemsReq
	4, // 3: deposit.DepositService.CreateDeposit:input_type -> deposit.CreateDepositReq
	3, // 4: deposit.DepositService.ListDepositAmountItems:output_type -> deposit.DepositAmountItemsRes
	5, // 5: deposit.DepositService.CreateDeposit:output_type -> deposit.CreateDepositRes
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_deposit_deposit_proto_init() }
func file_user_deposit_deposit_proto_init() {
	if File_user_deposit_deposit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_deposit_deposit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositAmountItemsReq); i {
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
		file_user_deposit_deposit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmountItem); i {
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
		file_user_deposit_deposit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositBox); i {
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
		file_user_deposit_deposit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositAmountItemsRes); i {
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
		file_user_deposit_deposit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDepositReq); i {
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
		file_user_deposit_deposit_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDepositRes); i {
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
			RawDescriptor: file_user_deposit_deposit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_deposit_deposit_proto_goTypes,
		DependencyIndexes: file_user_deposit_deposit_proto_depIdxs,
		MessageInfos:      file_user_deposit_deposit_proto_msgTypes,
	}.Build()
	File_user_deposit_deposit_proto = out.File
	file_user_deposit_deposit_proto_rawDesc = nil
	file_user_deposit_deposit_proto_goTypes = nil
	file_user_deposit_deposit_proto_depIdxs = nil
}
