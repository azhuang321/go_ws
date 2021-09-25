// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: message.proto

package msgpb

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

type MsgType int32

const (
	MsgType_PING MsgType = 0
	MsgType_AUTH MsgType = 1
	MsgType_TEXT MsgType = 2
)

// Enum value maps for MsgType.
var (
	MsgType_name = map[int32]string{
		0: "PING",
		1: "AUTH",
		2: "TEXT",
	}
	MsgType_value = map[string]int32{
		"PING": 0,
		"AUTH": 1,
		"TEXT": 2,
	}
)

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}

func (x MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (MsgType) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgType.Descriptor instead.
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mine *Mine `protobuf:"bytes,1,opt,name=mine,proto3" json:"mine,omitempty"`
	To   *To   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Content) GetMine() *Mine {
	if x != nil {
		return x.Mine
	}
	return nil
}

func (x *Content) GetTo() *To {
	if x != nil {
		return x.To
	}
	return nil
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgType MsgType `protobuf:"varint,1,opt,name=msg_type,json=msgType,proto3,enum=Msg.MsgType" json:"msg_type,omitempty"`
	Path    string  `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Types that are assignable to Payload:
	//	*Msg_Token
	//	*Msg_Content
	Payload isMsg_Payload `protobuf_oneof:"payload"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *Msg) GetMsgType() MsgType {
	if x != nil {
		return x.MsgType
	}
	return MsgType_PING
}

func (x *Msg) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (m *Msg) GetPayload() isMsg_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Msg) GetToken() string {
	if x, ok := x.GetPayload().(*Msg_Token); ok {
		return x.Token
	}
	return ""
}

func (x *Msg) GetContent() *Content {
	if x, ok := x.GetPayload().(*Msg_Content); ok {
		return x.Content
	}
	return nil
}

type isMsg_Payload interface {
	isMsg_Payload()
}

type Msg_Token struct {
	Token string `protobuf:"bytes,3,opt,name=token,proto3,oneof"`
}

type Msg_Content struct {
	Content *Content `protobuf:"bytes,4,opt,name=content,proto3,oneof"`
}

func (*Msg_Token) isMsg_Payload() {}

func (*Msg_Content) isMsg_Payload() {}

type SendMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mine *Mine `protobuf:"bytes,1,opt,name=mine,proto3" json:"mine,omitempty"`
	To   *To   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *SendMsg) Reset() {
	*x = SendMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsg) ProtoMessage() {}

func (x *SendMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsg.ProtoReflect.Descriptor instead.
func (*SendMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *SendMsg) GetMine() *Mine {
	if x != nil {
		return x.Mine
	}
	return nil
}

func (x *SendMsg) GetTo() *To {
	if x != nil {
		return x.To
	}
	return nil
}

type HeartMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *HeartMsg) Reset() {
	*x = HeartMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartMsg) ProtoMessage() {}

func (x *HeartMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartMsg.ProtoReflect.Descriptor instead.
func (*HeartMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *HeartMsg) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

//自身消息
type Mine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Avatar   string `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Id       uint32 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Mine     bool   `protobuf:"varint,4,opt,name=mine,proto3" json:"mine,omitempty"`
	Content  string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Mine) Reset() {
	*x = Mine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mine) ProtoMessage() {}

func (x *Mine) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mine.ProtoReflect.Descriptor instead.
func (*Mine) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *Mine) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Mine) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Mine) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Mine) GetMine() bool {
	if x != nil {
		return x.Mine
	}
	return false
}

func (x *Mine) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type To struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Avatar      string `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Id          uint32 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Type        string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Cid         uint32 `protobuf:"varint,5,opt,name=cid,proto3" json:"cid,omitempty"`
	Content     string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp   uint32 `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Name        string `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	HistoryTime uint64 `protobuf:"varint,9,opt,name=history_time,json=historyTime,proto3" json:"history_time,omitempty"`
	Sign        string `protobuf:"bytes,10,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *To) Reset() {
	*x = To{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *To) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*To) ProtoMessage() {}

func (x *To) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use To.ProtoReflect.Descriptor instead.
func (*To) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *To) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *To) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *To) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *To) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *To) GetCid() uint32 {
	if x != nil {
		return x.Cid
	}
	return 0
}

func (x *To) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *To) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *To) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *To) GetHistoryTime() uint64 {
	if x != nil {
		return x.HistoryTime
	}
	return 0
}

func (x *To) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x4d, 0x73, 0x67, 0x22, 0x41, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x04, 0x6d, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x4d, 0x73, 0x67, 0x2e, 0x4d, 0x69, 0x6e, 0x65, 0x52, 0x04, 0x6d, 0x69, 0x6e, 0x65, 0x12, 0x17,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4d, 0x73, 0x67,
	0x2e, 0x54, 0x6f, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12,
	0x27, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x4d, 0x73, 0x67, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x73, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x09,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x41, 0x0a, 0x07, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x73, 0x67, 0x12, 0x1d, 0x0a, 0x04, 0x6d, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x73, 0x67, 0x2e, 0x4d, 0x69, 0x6e, 0x65, 0x52, 0x04, 0x6d,
	0x69, 0x6e, 0x65, 0x12, 0x17, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x4d, 0x73, 0x67, 0x2e, 0x54, 0x6f, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x08,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x78, 0x0a, 0x04,
	0x4d, 0x69, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x69, 0x6e, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6d, 0x69, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xf1, 0x01, 0x0a, 0x02, 0x54, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x2a, 0x27, 0x0a, 0x07, 0x4d, 0x73,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x41, 0x55, 0x54, 0x48, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45, 0x58,
	0x54, 0x10, 0x02, 0x42, 0x1a, 0x5a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x73, 0x67, 0x70, 0x62, 0x3b, 0x6d, 0x73, 0x67, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_message_proto_goTypes = []interface{}{
	(MsgType)(0),     // 0: Msg.MsgType
	(*Content)(nil),  // 1: Msg.Content
	(*Msg)(nil),      // 2: Msg.Msg
	(*SendMsg)(nil),  // 3: Msg.SendMsg
	(*HeartMsg)(nil), // 4: Msg.HeartMsg
	(*Mine)(nil),     // 5: Msg.Mine
	(*To)(nil),       // 6: Msg.To
}
var file_message_proto_depIdxs = []int32{
	5, // 0: Msg.Content.mine:type_name -> Msg.Mine
	6, // 1: Msg.Content.to:type_name -> Msg.To
	0, // 2: Msg.Msg.msg_type:type_name -> Msg.MsgType
	1, // 3: Msg.Msg.content:type_name -> Msg.Content
	5, // 4: Msg.SendMsg.mine:type_name -> Msg.Mine
	6, // 5: Msg.SendMsg.to:type_name -> Msg.To
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsg); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartMsg); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mine); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*To); i {
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
	file_message_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Msg_Token)(nil),
		(*Msg_Content)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
