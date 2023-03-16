// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: common.proto

package v1

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

type IRequestProtocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: msgpack:"action"
	Action string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty" msgpack:"action"`
	// @inject_tag: msgpack:"method"
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty" msgpack:"method"`
	// @inject_tag: msgpack:"callback"
	Callback string `protobuf:"bytes,3,opt,name=callback,proto3" json:"callback,omitempty" msgpack:"callback"`
	// @inject_tag: msgpack:"isCompress"
	IsCompress bool `protobuf:"varint,4,opt,name=isCompress,proto3" json:"isCompress,omitempty" msgpack:"isCompress"` // true 表示 data 为 proto 编码的字节数组
	// @inject_tag: msgpack:"channelId"
	ChannelId int32 `protobuf:"varint,5,opt,name=channelId,proto3" json:"channelId,omitempty" msgpack:"channelId"`
	// @inject_tag: msgpack:"data"
	Data []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty" msgpack:"data"` // data 是 proto 编码的字节
}

func (x *IRequestProtocol) Reset() {
	*x = IRequestProtocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IRequestProtocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IRequestProtocol) ProtoMessage() {}

func (x *IRequestProtocol) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IRequestProtocol.ProtoReflect.Descriptor instead.
func (*IRequestProtocol) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *IRequestProtocol) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *IRequestProtocol) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *IRequestProtocol) GetCallback() string {
	if x != nil {
		return x.Callback
	}
	return ""
}

func (x *IRequestProtocol) GetIsCompress() bool {
	if x != nil {
		return x.IsCompress
	}
	return false
}

func (x *IRequestProtocol) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *IRequestProtocol) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type IResponseProtocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: msgpack:"code" json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" msgpack:"code"`
	// @inject_tag: msgpack:"isCompress" json:"isCompress"
	IsCompress bool `protobuf:"varint,2,opt,name=isCompress,proto3" json:"isCompress" msgpack:"isCompress"` // true 表示 data 为 proto 编码的字节数组
	// @inject_tag: msgpack:"data" json:"data"
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data" msgpack:"data"` // data 是 proto 编码的字节
	// @inject_tag: msgpack:"callback" json:"callback"
	Callback string `protobuf:"bytes,4,opt,name=callback,proto3" json:"callback" msgpack:"callback"` //  string msg = 5; // 异常时的详细信息
}

func (x *IResponseProtocol) Reset() {
	*x = IResponseProtocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IResponseProtocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IResponseProtocol) ProtoMessage() {}

func (x *IResponseProtocol) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IResponseProtocol.ProtoReflect.Descriptor instead.
func (*IResponseProtocol) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *IResponseProtocol) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *IResponseProtocol) GetIsCompress() bool {
	if x != nil {
		return x.IsCompress
	}
	return false
}

func (x *IResponseProtocol) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *IResponseProtocol) GetCallback() string {
	if x != nil {
		return x.Callback
	}
	return ""
}

// 该结构序列化后，放在 IResponseProtocol.data 中
type RespData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: msgpack:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty" msgpack:"code"`
	// @inject_tag: msgpack:"msg"
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty" msgpack:"msg"`
	// @inject_tag: msgpack:"data"
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty" msgpack:"data"` // google.protobuf.Struct data = 3;
}

func (x *RespData) Reset() {
	*x = RespData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespData) ProtoMessage() {}

func (x *RespData) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespData.ProtoReflect.Descriptor instead.
func (*RespData) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *RespData) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RespData) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RespData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ErrInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 备注信息 | @inject_tag: msgpack:"msg" json:"msg"
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg" msgpack:"msg"`
}

func (x *ErrInfo) Reset() {
	*x = ErrInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrInfo) ProtoMessage() {}

func (x *ErrInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrInfo.ProtoReflect.Descriptor instead.
func (*ErrInfo) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *ErrInfo) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// 客户端建立连接后，请求一次登录，参数是用户的 token
type FishUserLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户的 token，一般从平台获取 | @inject_tag: msgpack:"token" json:"token"
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token" msgpack:"token"`
}

func (x *FishUserLogin) Reset() {
	*x = FishUserLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FishUserLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FishUserLogin) ProtoMessage() {}

func (x *FishUserLogin) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FishUserLogin.ProtoReflect.Descriptor instead.
func (*FishUserLogin) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *FishUserLogin) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AccountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user id | @inject_tag: msgpack:"uid" json:"uid"
	Uid int32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid" msgpack:"uid"`
	// | @inject_tag: msgpack:"nickname" json:"nickname"
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname" msgpack:"nickname"`
	// | @inject_tag: msgpack:"avatar" json:"avatar"
	Avatar string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar" msgpack:"avatar"`
	// | @inject_tag: msgpack:"gender" json:"gender"
	Gender int32 `protobuf:"varint,4,opt,name=gender,proto3" json:"gender" msgpack:"gender"`
}

func (x *AccountInfo) Reset() {
	*x = AccountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountInfo) ProtoMessage() {}

func (x *AccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountInfo.ProtoReflect.Descriptor instead.
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (x *AccountInfo) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *AccountInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *AccountInfo) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *AccountInfo) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

type NormalInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 备注信息 | @inject_tag: msgpack:"msg" json:"msg"
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg" msgpack:"msg"`
}

func (x *NormalInfo) Reset() {
	*x = NormalInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NormalInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormalInfo) ProtoMessage() {}

func (x *NormalInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormalInfo.ProtoReflect.Descriptor instead.
func (*NormalInfo) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{6}
}

func (x *NormalInfo) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// 加入游戏时对应的游戏 id
// 一场游戏的 id；全局唯一；不能有两场游戏 id 一样。
type DataOfJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 平台提供的房间 id | @inject_tag: msgpack:"roomId" json:"roomId"
	RoomId string `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId" msgpack:"roomId"`
	// 对局 id | @inject_tag: msgpack:"gameId" json:"gameId"
	GameId string `protobuf:"bytes,2,opt,name=gameId,proto3" json:"gameId" msgpack:"gameId"`
	// 当前玩家的 nft id | @inject_tag: msgpack:"nft" json:"nft"
	Nft int32 `protobuf:"varint,3,opt,name=nft,proto3" json:"nft" msgpack:"nft"`
	// 是否与机器人对战 | @inject_tag: msgpack:"needBot" json:"needBot"
	NeedBot int32 `protobuf:"varint,4,opt,name=needBot,proto3" json:"needBot" msgpack:"needBot"`
	// 机器人的 uid | @inject_tag: msgpack:"botId" json:"botId"
	BotId int32 `protobuf:"varint,5,opt,name=botId,proto3" json:"botId" msgpack:"botId"`
	// 机器人的 nft | @inject_tag: msgpack:"botNft" json:"botNft"
	BotNft int32 `protobuf:"varint,6,opt,name=botNft,proto3" json:"botNft" msgpack:"botNft"`
	// 门票 | @inject_tag: msgpack:"ticket" json:"ticket"
	Ticket float32 `protobuf:"fixed32,7,opt,name=ticket,proto3" json:"ticket" msgpack:"ticket"`
	// 对战类型，区分 1v1，淘汰赛等 | @inject_tag: msgpack:"pvpType" json:"pvpType"
	PvpType int32 `protobuf:"varint,8,opt,name=pvpType,proto3" json:"pvpType" msgpack:"pvpType"`
	// 机器人的 ai 策略类型 | @inject_tag: msgpack:"aiPolicyType" json:"aiPolicyType"
	AiPolicyType int32 `protobuf:"varint,9,opt,name=aiPolicyType,proto3" json:"aiPolicyType" msgpack:"aiPolicyType"`
	// 人数 | @inject_tag: msgpack:"num" json:"num"
	Num int32 `protobuf:"varint,10,opt,name=num,proto3" json:"num" msgpack:"num"`
	// 游戏挡位 | @inject_tag: msgpack:"level" json:"level"
	Level int32 `protobuf:"varint,11,opt,name=level,proto3" json:"level" msgpack:"level"`
	// 适配的最大坐标 | @inject_tag: msgpack:"posY" json:"posY"
	PosY int32 `protobuf:"varint,12,opt,name=posY,proto3" json:"posY" msgpack:"posY"`
}

func (x *DataOfJoin) Reset() {
	*x = DataOfJoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataOfJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataOfJoin) ProtoMessage() {}

func (x *DataOfJoin) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataOfJoin.ProtoReflect.Descriptor instead.
func (*DataOfJoin) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{7}
}

func (x *DataOfJoin) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *DataOfJoin) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *DataOfJoin) GetNft() int32 {
	if x != nil {
		return x.Nft
	}
	return 0
}

func (x *DataOfJoin) GetNeedBot() int32 {
	if x != nil {
		return x.NeedBot
	}
	return 0
}

func (x *DataOfJoin) GetBotId() int32 {
	if x != nil {
		return x.BotId
	}
	return 0
}

func (x *DataOfJoin) GetBotNft() int32 {
	if x != nil {
		return x.BotNft
	}
	return 0
}

func (x *DataOfJoin) GetTicket() float32 {
	if x != nil {
		return x.Ticket
	}
	return 0
}

func (x *DataOfJoin) GetPvpType() int32 {
	if x != nil {
		return x.PvpType
	}
	return 0
}

func (x *DataOfJoin) GetAiPolicyType() int32 {
	if x != nil {
		return x.AiPolicyType
	}
	return 0
}

func (x *DataOfJoin) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *DataOfJoin) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *DataOfJoin) GetPosY() int32 {
	if x != nil {
		return x.PosY
	}
	return 0
}

type DataInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: msgpack:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty" msgpack:"code"`
	// @inject_tag: msgpack:"msg"
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty" msgpack:"msg"`
}

func (x *DataInfoResp) Reset() {
	*x = DataInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataInfoResp) ProtoMessage() {}

func (x *DataInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataInfoResp.ProtoReflect.Descriptor instead.
func (*DataInfoResp) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{8}
}

func (x *DataInfoResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DataInfoResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x62, 0x6c, 0x6f, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x22, 0xb0, 0x01, 0x0a, 0x10, 0x49, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x43,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69,
	0x73, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x77, 0x0a, 0x11, 0x49,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x22, 0x44, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1b, 0x0a, 0x07, 0x45, 0x72,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x25, 0x0a, 0x0d, 0x46, 0x69, 0x73, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6b,
	0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x1e, 0x0a, 0x0a, 0x4e,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xa8, 0x02, 0x0a, 0x0a,
	0x44, 0x61, 0x74, 0x61, 0x4f, 0x66, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x66,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x66, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6e, 0x65, 0x65, 0x64, 0x42, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6e,
	0x65, 0x65, 0x64, 0x42, 0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x6f, 0x74, 0x4e, 0x66, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6f,
	0x74, 0x4e, 0x66, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x76, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70,
	0x76, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x69, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61, 0x69,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75,
	0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x59, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x6f, 0x73, 0x59, 0x22, 0x34, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x78, 0x79, 0x2f, 0x62,
	0x6c, 0x6f, 0x6f, 0x64, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x62, 0x6c, 0x6f, 0x6f, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_common_proto_goTypes = []interface{}{
	(*IRequestProtocol)(nil),  // 0: blood.v1.IRequestProtocol
	(*IResponseProtocol)(nil), // 1: blood.v1.IResponseProtocol
	(*RespData)(nil),          // 2: blood.v1.RespData
	(*ErrInfo)(nil),           // 3: blood.v1.ErrInfo
	(*FishUserLogin)(nil),     // 4: blood.v1.FishUserLogin
	(*AccountInfo)(nil),       // 5: blood.v1.AccountInfo
	(*NormalInfo)(nil),        // 6: blood.v1.NormalInfo
	(*DataOfJoin)(nil),        // 7: blood.v1.DataOfJoin
	(*DataInfoResp)(nil),      // 8: blood.v1.DataInfoResp
}
var file_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IRequestProtocol); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IResponseProtocol); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespData); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrInfo); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FishUserLogin); i {
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
		file_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountInfo); i {
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
		file_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NormalInfo); i {
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
		file_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataOfJoin); i {
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
		file_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataInfoResp); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
