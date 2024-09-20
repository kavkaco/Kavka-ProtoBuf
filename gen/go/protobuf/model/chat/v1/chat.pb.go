// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: protobuf/model/chat/v1/chat.proto

package chatv1

import (
	v1 "github.com/kavkaco/Kavka-ProtoBuf/gen/go/protobuf/model/user/v1"
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

type ChatType int32

const (
	ChatType_CHAT_TYPE_UNSPECIFIED ChatType = 0
	ChatType_CHAT_TYPE_CHANNEL     ChatType = 1
	ChatType_CHAT_TYPE_GROUP       ChatType = 2
	ChatType_CHAT_TYPE_DIRECT      ChatType = 3
)

// Enum value maps for ChatType.
var (
	ChatType_name = map[int32]string{
		0: "CHAT_TYPE_UNSPECIFIED",
		1: "CHAT_TYPE_CHANNEL",
		2: "CHAT_TYPE_GROUP",
		3: "CHAT_TYPE_DIRECT",
	}
	ChatType_value = map[string]int32{
		"CHAT_TYPE_UNSPECIFIED": 0,
		"CHAT_TYPE_CHANNEL":     1,
		"CHAT_TYPE_GROUP":       2,
		"CHAT_TYPE_DIRECT":      3,
	}
)

func (x ChatType) Enum() *ChatType {
	p := new(ChatType)
	*p = x
	return p
}

func (x ChatType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatType) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_model_chat_v1_chat_proto_enumTypes[0].Descriptor()
}

func (ChatType) Type() protoreflect.EnumType {
	return &file_protobuf_model_chat_v1_chat_proto_enumTypes[0]
}

func (x ChatType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatType.Descriptor instead.
func (ChatType) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_model_chat_v1_chat_proto_rawDescGZIP(), []int{0}
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LastName string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_model_chat_v1_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_model_chat_v1_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_protobuf_model_chat_v1_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Member) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Member) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Member) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type LastMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageType    string `protobuf:"bytes,1,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	MessageCaption string `protobuf:"bytes,2,opt,name=message_caption,json=messageCaption,proto3" json:"message_caption,omitempty"`
}

func (x *LastMessage) Reset() {
	*x = LastMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_model_chat_v1_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LastMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastMessage) ProtoMessage() {}

func (x *LastMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_model_chat_v1_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LastMessage.ProtoReflect.Descriptor instead.
func (*LastMessage) Descriptor() ([]byte, []int) {
	return file_protobuf_model_chat_v1_chat_proto_rawDescGZIP(), []int{1}
}

func (x *LastMessage) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *LastMessage) GetMessageCaption() string {
	if x != nil {
		return x.MessageCaption
	}
	return ""
}

type ChatDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ChatDetailType:
	//
	//	*ChatDetail_ChannelDetail
	//	*ChatDetail_GroupDetail
	//	*ChatDetail_DirectDetail
	ChatDetailType isChatDetail_ChatDetailType `protobuf_oneof:"chat_detail_type"`
}

func (x *ChatDetail) Reset() {
	*x = ChatDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_model_chat_v1_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatDetail) ProtoMessage() {}

func (x *ChatDetail) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_model_chat_v1_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatDetail.ProtoReflect.Descriptor instead.
func (*ChatDetail) Descriptor() ([]byte, []int) {
	return file_protobuf_model_chat_v1_chat_proto_rawDescGZIP(), []int{2}
}

func (m *ChatDetail) GetChatDetailType() isChatDetail_ChatDetailType {
	if m != nil {
		return m.ChatDetailType
	}
	return nil
}

func (x *ChatDetail) GetChannelDetail() *ChannelChatDetail {
	if x, ok := x.GetChatDetailType().(*ChatDetail_ChannelDetail); ok {
		return x.ChannelDetail
	}
	return nil
}

func (x *ChatDetail) GetGroupDetail() *GroupChatDetail {
	if x, ok := x.GetChatDetailType().(*ChatDetail_GroupDetail); ok {
		return x.GroupDetail
	}
	return nil
}

func (x *ChatDetail) GetDirectDetail() *DirectChatDetail {
	if x, ok := x.GetChatDetailType().(*ChatDetail_DirectDetail); ok {
		return x.DirectDetail
	}
	return nil
}

type isChatDetail_ChatDetailType interface {
	isChatDetail_ChatDetailType()
}

type ChatDetail_ChannelDetail struct {
	ChannelDetail *ChannelChatDetail `protobuf:"bytes,1,opt,name=channel_detail,json=channelDetail,proto3,oneof"`
}

type ChatDetail_GroupDetail struct {
	GroupDetail *GroupChatDetail `protobuf:"bytes,2,opt,name=group_detail,json=groupDetail,proto3,oneof"`
}

type ChatDetail_DirectDetail struct {
	DirectDetail *DirectChatDetail `protobuf:"bytes,3,opt,name=direct_detail,json=directDetail,proto3,oneof"`
}

func (*ChatDetail_ChannelDetail) isChatDetail_ChatDetailType() {}

func (*ChatDetail_GroupDetail) isChatDetail_ChatDetailType() {}

func (*ChatDetail_DirectDetail) isChatDetail_ChatDetailType() {}

type ChannelChatDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Members      []string `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	Admins       []string `protobuf:"bytes,3,rep,name=admins,proto3" json:"admins,omitempty"`
	Owner        string   `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	RemovedUsers []string `protobuf:"bytes,5,rep,name=removed_users,json=removedUsers,proto3" json:"removed_users,omitempty"`
	Username     string   `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	Description  string   `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ChannelChatDetail) Reset() {
	*x = ChannelChatDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_model_chat_v1_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelChatDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelChatDetail) ProtoMessage() {}

func (x *ChannelChatDetail) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_model_chat_v1_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelChatDetail.ProtoReflect.Descriptor instead.
func (*ChannelChatDetail) Descriptor() ([]byte, []int) {
	return file_protobuf_model_chat_v1_chat_proto_rawDescGZIP(), []int{3}
}

func (x *ChannelChatDetail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ChannelChatDetail) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *ChannelChatDetail) GetAdmins() []string {
	if x != nil {
		return x.Admins
	}
	return nil
}

func (x *ChannelChatDetail) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ChannelChatDetail) GetRemovedUsers() []string {
	if x != nil {
		return x.RemovedUsers
	}
	return nil
}

func (x *ChannelChatDetail) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ChannelChatDetail) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GroupChatDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Members      []string `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	Admins       []string `protobuf:"bytes,3,rep,name=admins,proto3" json:"admins,omitempty"`
	Owner        string   `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	RemovedUsers []string `protobuf:"bytes,5,rep,name=removed_users,json=removedUsers,proto3" json:"removed_users,omitempty"`
	Username     string   `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	Description  string   `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GroupChatDetail) Reset() {
	*x = GroupChatDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_model_chat_v1_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupChatDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupChatDetail) ProtoMessage() {}

func (x *GroupChatDetail) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_model_chat_v1_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupChatDetail.ProtoReflect.Descriptor instead.
func (*GroupChatDetail) Descriptor() ([]byte, []int) {
	return file_protobuf_model_chat_v1_chat_proto_rawDescGZIP(), []int{4}
}

func (x *GroupChatDetail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GroupChatDetail) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *GroupChatDetail) GetAdmins() []string {
	if x != nil {
		return x.Admins
	}
	return nil
}

func (x *GroupChatDetail) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GroupChatDetail) GetRemovedUsers() []string {
	if x != nil {
		return x.RemovedUsers
	}
	return nil
}

func (x *GroupChatDetail) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GroupChatDetail) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type DirectChatDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipient *v1.User `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
}

func (x *DirectChatDetail) Reset() {
	*x = DirectChatDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_model_chat_v1_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectChatDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectChatDetail) ProtoMessage() {}

func (x *DirectChatDetail) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_model_chat_v1_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectChatDetail.ProtoReflect.Descriptor instead.
func (*DirectChatDetail) Descriptor() ([]byte, []int) {
	return file_protobuf_model_chat_v1_chat_proto_rawDescGZIP(), []int{5}
}

func (x *DirectChatDetail) GetRecipient() *v1.User {
	if x != nil {
		return x.Recipient
	}
	return nil
}

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId      string       `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	ChatType    ChatType     `protobuf:"varint,2,opt,name=chat_type,json=chatType,proto3,enum=model.chat.v1.ChatType" json:"chat_type,omitempty"`
	ChatDetail  *ChatDetail  `protobuf:"bytes,3,opt,name=chat_detail,json=chatDetail,proto3" json:"chat_detail,omitempty"`
	LastMessage *LastMessage `protobuf:"bytes,4,opt,name=last_message,json=lastMessage,proto3" json:"last_message,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_model_chat_v1_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_model_chat_v1_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_protobuf_model_chat_v1_chat_proto_rawDescGZIP(), []int{6}
}

func (x *Chat) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *Chat) GetChatType() ChatType {
	if x != nil {
		return x.ChatType
	}
	return ChatType_CHAT_TYPE_UNSPECIFIED
}

func (x *Chat) GetChatDetail() *ChatDetail {
	if x != nil {
		return x.ChatDetail
	}
	return nil
}

func (x *Chat) GetLastMessage() *LastMessage {
	if x != nil {
		return x.LastMessage
	}
	return nil
}

var File_protobuf_model_chat_v1_chat_proto protoreflect.FileDescriptor

var file_protobuf_model_chat_v1_chat_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x59, 0x0a, 0x0b, 0x4c, 0x61, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x61, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf8, 0x01, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x49, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x48, 0x00, 0x52,
	0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x43,
	0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x68, 0x61, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x48, 0x00, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x46, 0x0a, 0x0d, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x43, 0x68, 0x61, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x48, 0x00, 0x52, 0x0c, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x12, 0x0a, 0x10, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22,
	0xd4, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd2, 0x01, 0x0a, 0x0f, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x43, 0x68, 0x61, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x45, 0x0a, 0x10, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x43, 0x68, 0x61, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x31, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x22, 0xd0, 0x01, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68,
	0x61, 0x74, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x67, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11,
	0x43, 0x48, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45,
	0x4c, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x48, 0x41, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x10, 0x03, 0x42, 0xbc,
	0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x43, 0x68, 0x61, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61,
	0x76, 0x6b, 0x61, 0x63, 0x6f, 0x2f, 0x4b, 0x61, 0x76, 0x6b, 0x61, 0x2d, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x42, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f,
	0x76, 0x31, 0x3b, 0x63, 0x68, 0x61, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x43, 0x58, 0xaa,
	0x02, 0x0d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x0d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x43, 0x68, 0x61, 0x74, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x19, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x43, 0x68, 0x61, 0x74, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x3a, 0x3a, 0x43, 0x68, 0x61, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_model_chat_v1_chat_proto_rawDescOnce sync.Once
	file_protobuf_model_chat_v1_chat_proto_rawDescData = file_protobuf_model_chat_v1_chat_proto_rawDesc
)

func file_protobuf_model_chat_v1_chat_proto_rawDescGZIP() []byte {
	file_protobuf_model_chat_v1_chat_proto_rawDescOnce.Do(func() {
		file_protobuf_model_chat_v1_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_model_chat_v1_chat_proto_rawDescData)
	})
	return file_protobuf_model_chat_v1_chat_proto_rawDescData
}

var file_protobuf_model_chat_v1_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_model_chat_v1_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protobuf_model_chat_v1_chat_proto_goTypes = []any{
	(ChatType)(0),             // 0: model.chat.v1.ChatType
	(*Member)(nil),            // 1: model.chat.v1.Member
	(*LastMessage)(nil),       // 2: model.chat.v1.LastMessage
	(*ChatDetail)(nil),        // 3: model.chat.v1.ChatDetail
	(*ChannelChatDetail)(nil), // 4: model.chat.v1.ChannelChatDetail
	(*GroupChatDetail)(nil),   // 5: model.chat.v1.GroupChatDetail
	(*DirectChatDetail)(nil),  // 6: model.chat.v1.DirectChatDetail
	(*Chat)(nil),              // 7: model.chat.v1.Chat
	(*v1.User)(nil),           // 8: model.user.v1.User
}
var file_protobuf_model_chat_v1_chat_proto_depIdxs = []int32{
	4, // 0: model.chat.v1.ChatDetail.channel_detail:type_name -> model.chat.v1.ChannelChatDetail
	5, // 1: model.chat.v1.ChatDetail.group_detail:type_name -> model.chat.v1.GroupChatDetail
	6, // 2: model.chat.v1.ChatDetail.direct_detail:type_name -> model.chat.v1.DirectChatDetail
	8, // 3: model.chat.v1.DirectChatDetail.recipient:type_name -> model.user.v1.User
	0, // 4: model.chat.v1.Chat.chat_type:type_name -> model.chat.v1.ChatType
	3, // 5: model.chat.v1.Chat.chat_detail:type_name -> model.chat.v1.ChatDetail
	2, // 6: model.chat.v1.Chat.last_message:type_name -> model.chat.v1.LastMessage
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_protobuf_model_chat_v1_chat_proto_init() }
func file_protobuf_model_chat_v1_chat_proto_init() {
	if File_protobuf_model_chat_v1_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_model_chat_v1_chat_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Member); i {
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
		file_protobuf_model_chat_v1_chat_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LastMessage); i {
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
		file_protobuf_model_chat_v1_chat_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ChatDetail); i {
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
		file_protobuf_model_chat_v1_chat_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ChannelChatDetail); i {
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
		file_protobuf_model_chat_v1_chat_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GroupChatDetail); i {
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
		file_protobuf_model_chat_v1_chat_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DirectChatDetail); i {
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
		file_protobuf_model_chat_v1_chat_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Chat); i {
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
	file_protobuf_model_chat_v1_chat_proto_msgTypes[2].OneofWrappers = []any{
		(*ChatDetail_ChannelDetail)(nil),
		(*ChatDetail_GroupDetail)(nil),
		(*ChatDetail_DirectDetail)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_model_chat_v1_chat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_model_chat_v1_chat_proto_goTypes,
		DependencyIndexes: file_protobuf_model_chat_v1_chat_proto_depIdxs,
		EnumInfos:         file_protobuf_model_chat_v1_chat_proto_enumTypes,
		MessageInfos:      file_protobuf_model_chat_v1_chat_proto_msgTypes,
	}.Build()
	File_protobuf_model_chat_v1_chat_proto = out.File
	file_protobuf_model_chat_v1_chat_proto_rawDesc = nil
	file_protobuf_model_chat_v1_chat_proto_goTypes = nil
	file_protobuf_model_chat_v1_chat_proto_depIdxs = nil
}
