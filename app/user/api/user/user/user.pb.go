// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: user/user/user.proto

package user

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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account     string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	ClientAgent string `protobuf:"bytes,3,opt,name=clientAgent,proto3" json:"clientAgent,omitempty"`
	DeviceId    string `protobuf:"bytes,4,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	Xid         string `protobuf:"bytes,5,opt,name=xid,proto3" json:"xid,omitempty"`
	Ip          string `protobuf:"bytes,6,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_user_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginRequest) GetClientAgent() string {
	if x != nil {
		return x.ClientAgent
	}
	return ""
}

func (x *LoginRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *LoginRequest) GetXid() string {
	if x != nil {
		return x.Xid
	}
	return ""
}

func (x *LoginRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type RegRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account     string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Ip          string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Xid         string `protobuf:"bytes,4,opt,name=xid,proto3" json:"xid,omitempty"`
	Phone       string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Email       string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	ClientAgent string `protobuf:"bytes,7,opt,name=clientAgent,proto3" json:"clientAgent,omitempty"`
	DeviceId    string `protobuf:"bytes,8,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
}

func (x *RegRequest) Reset() {
	*x = RegRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegRequest) ProtoMessage() {}

func (x *RegRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegRequest.ProtoReflect.Descriptor instead.
func (*RegRequest) Descriptor() ([]byte, []int) {
	return file_user_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *RegRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *RegRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *RegRequest) GetXid() string {
	if x != nil {
		return x.Xid
	}
	return ""
}

func (x *RegRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RegRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegRequest) GetClientAgent() string {
	if x != nil {
		return x.ClientAgent
	}
	return ""
}

func (x *RegRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type RegReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RegReply) Reset() {
	*x = RegReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegReply) ProtoMessage() {}

func (x *RegReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegReply.ProtoReflect.Descriptor instead.
func (*RegReply) Descriptor() ([]byte, []int) {
	return file_user_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *RegReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_user_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *LoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UserInfoRequest) Reset() {
	*x = UserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRequest) ProtoMessage() {}

func (x *UserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRequest.ProtoReflect.Descriptor instead.
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return file_user_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserInfoRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid        int32  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	UidInt64   int64  `protobuf:"varint,2,opt,name=uidInt64,proto3" json:"uidInt64,omitempty"`
	Account    string `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	ClientIP   string `protobuf:"bytes,4,opt,name=clientIP,proto3" json:"clientIP,omitempty"`
	Lang       string `protobuf:"bytes,5,opt,name=lang,proto3" json:"lang,omitempty"`
	Pid        int32  `protobuf:"varint,6,opt,name=pid,proto3" json:"pid,omitempty"`
	ParentPath string `protobuf:"bytes,7,opt,name=parentPath,proto3" json:"parentPath,omitempty"`
	Phone      string `protobuf:"bytes,8,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *UserInfoReply) Reset() {
	*x = UserInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoReply) ProtoMessage() {}

func (x *UserInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoReply.ProtoReflect.Descriptor instead.
func (*UserInfoReply) Descriptor() ([]byte, []int) {
	return file_user_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserInfoReply) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *UserInfoReply) GetUidInt64() int64 {
	if x != nil {
		return x.UidInt64
	}
	return 0
}

func (x *UserInfoReply) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserInfoReply) GetClientIP() string {
	if x != nil {
		return x.ClientIP
	}
	return ""
}

func (x *UserInfoReply) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *UserInfoReply) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *UserInfoReply) GetParentPath() string {
	if x != nil {
		return x.ParentPath
	}
	return ""
}

func (x *UserInfoReply) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type WalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *WalletRequest) Reset() {
	*x = WalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletRequest) ProtoMessage() {}

func (x *WalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletRequest.ProtoReflect.Descriptor instead.
func (*WalletRequest) Descriptor() ([]byte, []int) {
	return file_user_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *WalletRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type WalletReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance     float64 `protobuf:"fixed64,1,opt,name=balance,proto3" json:"balance,omitempty"`
	Freeze      float64 `protobuf:"fixed64,2,opt,name=freeze,proto3" json:"freeze,omitempty"`
	TotalProfit float64 `protobuf:"fixed64,3,opt,name=totalProfit,proto3" json:"totalProfit,omitempty"`
}

func (x *WalletReply) Reset() {
	*x = WalletReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletReply) ProtoMessage() {}

func (x *WalletReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletReply.ProtoReflect.Descriptor instead.
func (*WalletReply) Descriptor() ([]byte, []int) {
	return file_user_user_user_proto_rawDescGZIP(), []int{7}
}

func (x *WalletReply) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *WalletReply) GetFreeze() float64 {
	if x != nil {
		return x.Freeze
	}
	return 0
}

func (x *WalletReply) GetTotalProfit() float64 {
	if x != nil {
		return x.TotalProfit
	}
	return 0
}

type UpdateLoginPassReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldPass string `protobuf:"bytes,1,opt,name=oldPass,proto3" json:"oldPass,omitempty"`
	NewPass string `protobuf:"bytes,2,opt,name=newPass,proto3" json:"newPass,omitempty"`
	Uid     int64  `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *UpdateLoginPassReq) Reset() {
	*x = UpdateLoginPassReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLoginPassReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLoginPassReq) ProtoMessage() {}

func (x *UpdateLoginPassReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLoginPassReq.ProtoReflect.Descriptor instead.
func (*UpdateLoginPassReq) Descriptor() ([]byte, []int) {
	return file_user_user_user_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateLoginPassReq) GetOldPass() string {
	if x != nil {
		return x.OldPass
	}
	return ""
}

func (x *UpdateLoginPassReq) GetNewPass() string {
	if x != nil {
		return x.NewPass
	}
	return ""
}

func (x *UpdateLoginPassReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type UpdateLoginPassRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateLoginPassRes) Reset() {
	*x = UpdateLoginPassRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLoginPassRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLoginPassRes) ProtoMessage() {}

func (x *UpdateLoginPassRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLoginPassRes.ProtoReflect.Descriptor instead.
func (*UpdateLoginPassRes) Descriptor() ([]byte, []int) {
	return file_user_user_user_proto_rawDescGZIP(), []int{9}
}

type SendMsgCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *SendMsgCodeReq) Reset() {
	*x = SendMsgCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgCodeReq) ProtoMessage() {}

func (x *SendMsgCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgCodeReq.ProtoReflect.Descriptor instead.
func (*SendMsgCodeReq) Descriptor() ([]byte, []int) {
	return file_user_user_user_proto_rawDescGZIP(), []int{10}
}

func (x *SendMsgCodeReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type SendMsgCodeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendMsgCodeRes) Reset() {
	*x = SendMsgCodeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_user_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgCodeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgCodeRes) ProtoMessage() {}

func (x *SendMsgCodeRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_user_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgCodeRes.ProtoReflect.Descriptor instead.
func (*SendMsgCodeRes) Descriptor() ([]byte, []int) {
	return file_user_user_user_proto_rawDescGZIP(), []int{11}
}

type BindPhoneReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Uid   int64  `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *BindPhoneReq) Reset() {
	*x = BindPhoneReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_user_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindPhoneReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindPhoneReq) ProtoMessage() {}

func (x *BindPhoneReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_user_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindPhoneReq.ProtoReflect.Descriptor instead.
func (*BindPhoneReq) Descriptor() ([]byte, []int) {
	return file_user_user_user_proto_rawDescGZIP(), []int{12}
}

func (x *BindPhoneReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *BindPhoneReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *BindPhoneReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type BindPhoneRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *BindPhoneRes) Reset() {
	*x = BindPhoneRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_user_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindPhoneRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindPhoneRes) ProtoMessage() {}

func (x *BindPhoneRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_user_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindPhoneRes.ProtoReflect.Descriptor instead.
func (*BindPhoneRes) Descriptor() ([]byte, []int) {
	return file_user_user_user_proto_rawDescGZIP(), []int{13}
}

func (x *BindPhoneRes) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_user_user_user_proto protoreflect.FileDescriptor

var file_user_user_user_proto_rawDesc = []byte{
	0x0a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xa4, 0x01, 0x0a,
	0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x78, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x78, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x70, 0x22, 0xce, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x78, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x78, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x22, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x27, 0x0a, 0x0f, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0xcf, 0x01, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x69, 0x64, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x69, 0x64, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x50, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x21, 0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x61, 0x0a, 0x0b, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x22, 0x5a, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e,
	0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65,
	0x77, 0x50, 0x61, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x22, 0x26, 0x0a,
	0x0e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x22, 0x4a, 0x0a, 0x0c, 0x42, 0x69, 0x6e, 0x64, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x0c, 0x42, 0x69, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x32, 0x86, 0x03, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x03, 0x52, 0x65, 0x67,
	0x12, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x36, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x06, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x45, 0x0a, 0x0f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x12, 0x18,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x12, 0x39, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x12, 0x33, 0x0a,
	0x09, 0x42, 0x69, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_user_user_proto_rawDescOnce sync.Once
	file_user_user_user_proto_rawDescData = file_user_user_user_proto_rawDesc
)

func file_user_user_user_proto_rawDescGZIP() []byte {
	file_user_user_user_proto_rawDescOnce.Do(func() {
		file_user_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_user_user_proto_rawDescData)
	})
	return file_user_user_user_proto_rawDescData
}

var file_user_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_user_user_user_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),       // 0: user.LoginRequest
	(*RegRequest)(nil),         // 1: user.RegRequest
	(*RegReply)(nil),           // 2: user.RegReply
	(*LoginReply)(nil),         // 3: user.LoginReply
	(*UserInfoRequest)(nil),    // 4: user.UserInfoRequest
	(*UserInfoReply)(nil),      // 5: user.UserInfoReply
	(*WalletRequest)(nil),      // 6: user.WalletRequest
	(*WalletReply)(nil),        // 7: user.WalletReply
	(*UpdateLoginPassReq)(nil), // 8: user.UpdateLoginPassReq
	(*UpdateLoginPassRes)(nil), // 9: user.UpdateLoginPassRes
	(*SendMsgCodeReq)(nil),     // 10: user.SendMsgCodeReq
	(*SendMsgCodeRes)(nil),     // 11: user.SendMsgCodeRes
	(*BindPhoneReq)(nil),       // 12: user.BindPhoneReq
	(*BindPhoneRes)(nil),       // 13: user.BindPhoneRes
}
var file_user_user_user_proto_depIdxs = []int32{
	1,  // 0: user.UserService.Reg:input_type -> user.RegRequest
	0,  // 1: user.UserService.Login:input_type -> user.LoginRequest
	4,  // 2: user.UserService.UserInfo:input_type -> user.UserInfoRequest
	6,  // 3: user.UserService.Wallet:input_type -> user.WalletRequest
	8,  // 4: user.UserService.UpdateLoginPass:input_type -> user.UpdateLoginPassReq
	10, // 5: user.UserService.SendMsgCode:input_type -> user.SendMsgCodeReq
	12, // 6: user.UserService.BindPhone:input_type -> user.BindPhoneReq
	2,  // 7: user.UserService.Reg:output_type -> user.RegReply
	3,  // 8: user.UserService.Login:output_type -> user.LoginReply
	5,  // 9: user.UserService.UserInfo:output_type -> user.UserInfoReply
	7,  // 10: user.UserService.Wallet:output_type -> user.WalletReply
	9,  // 11: user.UserService.UpdateLoginPass:output_type -> user.UpdateLoginPassRes
	11, // 12: user.UserService.SendMsgCode:output_type -> user.SendMsgCodeRes
	13, // 13: user.UserService.BindPhone:output_type -> user.BindPhoneRes
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_user_user_user_proto_init() }
func file_user_user_user_proto_init() {
	if File_user_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_user_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegRequest); i {
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
		file_user_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegReply); i {
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
		file_user_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_user_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoRequest); i {
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
		file_user_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoReply); i {
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
		file_user_user_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletRequest); i {
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
		file_user_user_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletReply); i {
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
		file_user_user_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLoginPassReq); i {
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
		file_user_user_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLoginPassRes); i {
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
		file_user_user_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgCodeReq); i {
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
		file_user_user_user_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgCodeRes); i {
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
		file_user_user_user_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindPhoneReq); i {
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
		file_user_user_user_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindPhoneRes); i {
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
			RawDescriptor: file_user_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_user_user_proto_goTypes,
		DependencyIndexes: file_user_user_user_proto_depIdxs,
		MessageInfos:      file_user_user_user_proto_msgTypes,
	}.Build()
	File_user_user_user_proto = out.File
	file_user_user_user_proto_rawDesc = nil
	file_user_user_user_proto_goTypes = nil
	file_user_user_user_proto_depIdxs = nil
}
