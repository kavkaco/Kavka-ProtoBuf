// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: protobuf/chat/v1/chat.proto

package chatv1

import (
	v1 "github.com/kavkaco/Kavka-ProtoBuf/gen/go/protobuf/model/chat/v1"
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

type GetDirectChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipientUserId string `protobuf:"bytes,1,opt,name=recipient_user_id,json=recipientUserId,proto3" json:"recipient_user_id,omitempty"`
}

func (x *GetDirectChatRequest) Reset() {
	*x = GetDirectChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_v1_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDirectChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDirectChatRequest) ProtoMessage() {}

func (x *GetDirectChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_v1_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDirectChatRequest.ProtoReflect.Descriptor instead.
func (*GetDirectChatRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_v1_chat_proto_rawDescGZIP(), []int{0}
}

func (x *GetDirectChatRequest) GetRecipientUserId() string {
	if x != nil {
		return x.RecipientUserId
	}
	return ""
}

type GetDirectChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat *v1.Chat `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *GetDirectChatResponse) Reset() {
	*x = GetDirectChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_v1_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDirectChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDirectChatResponse) ProtoMessage() {}

func (x *GetDirectChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_v1_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDirectChatResponse.ProtoReflect.Descriptor instead.
func (*GetDirectChatResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_v1_chat_proto_rawDescGZIP(), []int{1}
}

func (x *GetDirectChatResponse) GetChat() *v1.Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type GetChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *GetChatRequest) Reset() {
	*x = GetChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_v1_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatRequest) ProtoMessage() {}

func (x *GetChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_v1_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatRequest.ProtoReflect.Descriptor instead.
func (*GetChatRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_v1_chat_proto_rawDescGZIP(), []int{2}
}

func (x *GetChatRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type GetChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat *v1.Chat `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *GetChatResponse) Reset() {
	*x = GetChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_v1_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatResponse) ProtoMessage() {}

func (x *GetChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_v1_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatResponse.ProtoReflect.Descriptor instead.
func (*GetChatResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_v1_chat_proto_rawDescGZIP(), []int{3}
}

func (x *GetChatResponse) GetChat() *v1.Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type GetUserChatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserChatsRequest) Reset() {
	*x = GetUserChatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_v1_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserChatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserChatsRequest) ProtoMessage() {}

func (x *GetUserChatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_v1_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserChatsRequest.ProtoReflect.Descriptor instead.
func (*GetUserChatsRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_v1_chat_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserChatsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserChatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chats []*v1.Chat `protobuf:"bytes,1,rep,name=chats,proto3" json:"chats,omitempty"`
}

func (x *GetUserChatsResponse) Reset() {
	*x = GetUserChatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_v1_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserChatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserChatsResponse) ProtoMessage() {}

func (x *GetUserChatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_v1_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserChatsResponse.ProtoReflect.Descriptor instead.
func (*GetUserChatsResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_v1_chat_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserChatsResponse) GetChats() []*v1.Chat {
	if x != nil {
		return x.Chats
	}
	return nil
}

type CreateDirectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipientUserId string `protobuf:"bytes,1,opt,name=recipient_user_id,json=recipientUserId,proto3" json:"recipient_user_id,omitempty"`
}

func (x *CreateDirectRequest) Reset() {
	*x = CreateDirectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_v1_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDirectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDirectRequest) ProtoMessage() {}

func (x *CreateDirectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_v1_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDirectRequest.ProtoReflect.Descriptor instead.
func (*CreateDirectRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_v1_chat_proto_rawDescGZIP(), []int{6}
}

func (x *CreateDirectRequest) GetRecipientUserId() string {
	if x != nil {
		return x.RecipientUserId
	}
	return ""
}

type CreateDirectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat *v1.Chat `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *CreateDirectResponse) Reset() {
	*x = CreateDirectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_v1_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDirectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDirectResponse) ProtoMessage() {}

func (x *CreateDirectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_v1_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDirectResponse.ProtoReflect.Descriptor instead.
func (*CreateDirectResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_v1_chat_proto_rawDescGZIP(), []int{7}
}

func (x *CreateDirectResponse) GetChat() *v1.Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type CreateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_v1_chat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_v1_chat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_v1_chat_proto_rawDescGZIP(), []int{8}
}

func (x *CreateGroupRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateGroupRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateGroupRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat *v1.Chat `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *CreateGroupResponse) Reset() {
	*x = CreateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_v1_chat_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupResponse) ProtoMessage() {}

func (x *CreateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_v1_chat_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateGroupResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_v1_chat_proto_rawDescGZIP(), []int{9}
}

func (x *CreateGroupResponse) GetChat() *v1.Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type CreateChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateChannelRequest) Reset() {
	*x = CreateChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_v1_chat_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChannelRequest) ProtoMessage() {}

func (x *CreateChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_v1_chat_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChannelRequest.ProtoReflect.Descriptor instead.
func (*CreateChannelRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_v1_chat_proto_rawDescGZIP(), []int{10}
}

func (x *CreateChannelRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateChannelRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateChannelRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat *v1.Chat `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *CreateChannelResponse) Reset() {
	*x = CreateChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_v1_chat_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChannelResponse) ProtoMessage() {}

func (x *CreateChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_v1_chat_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChannelResponse.ProtoReflect.Descriptor instead.
func (*CreateChannelResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_v1_chat_proto_rawDescGZIP(), []int{11}
}

func (x *CreateChannelResponse) GetChat() *v1.Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type JoinChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *JoinChatRequest) Reset() {
	*x = JoinChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_v1_chat_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinChatRequest) ProtoMessage() {}

func (x *JoinChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_v1_chat_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinChatRequest.ProtoReflect.Descriptor instead.
func (*JoinChatRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_v1_chat_proto_rawDescGZIP(), []int{12}
}

func (x *JoinChatRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type JoinChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat *v1.Chat `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *JoinChatResponse) Reset() {
	*x = JoinChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_v1_chat_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinChatResponse) ProtoMessage() {}

func (x *JoinChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_v1_chat_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinChatResponse.ProtoReflect.Descriptor instead.
func (*JoinChatResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_v1_chat_proto_rawDescGZIP(), []int{13}
}

func (x *JoinChatResponse) GetChat() *v1.Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

var File_protobuf_chat_v1_chat_proto protoreflect.FileDescriptor

var file_protobuf_chat_v1_chat_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x40, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22,
	0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x2e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x05, 0x63, 0x68, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x05, 0x63, 0x68, 0x61, 0x74, 0x73, 0x22, 0x41, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x68, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3e, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27,
	0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x6a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04,
	0x63, 0x68, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x2a, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49,
	0x64, 0x22, 0x3b, 0x0a, 0x10, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x32, 0x9e,
	0x04, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x43, 0x68, 0x61, 0x74, 0x12,
	0x1d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x73,
	0x12, 0x1c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4d, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12,
	0x1c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1b, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1d, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x08,
	0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69,
	0x6e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x97, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x42,
	0x09, 0x43, 0x68, 0x61, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x76, 0x6b, 0x61, 0x63, 0x6f,
	0x2f, 0x4b, 0x61, 0x76, 0x6b, 0x61, 0x2d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x68, 0x61, 0x74, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x07, 0x43, 0x68, 0x61, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x43, 0x68, 0x61, 0x74, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x08, 0x43, 0x68, 0x61, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protobuf_chat_v1_chat_proto_rawDescOnce sync.Once
	file_protobuf_chat_v1_chat_proto_rawDescData = file_protobuf_chat_v1_chat_proto_rawDesc
)

func file_protobuf_chat_v1_chat_proto_rawDescGZIP() []byte {
	file_protobuf_chat_v1_chat_proto_rawDescOnce.Do(func() {
		file_protobuf_chat_v1_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_chat_v1_chat_proto_rawDescData)
	})
	return file_protobuf_chat_v1_chat_proto_rawDescData
}

var file_protobuf_chat_v1_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_protobuf_chat_v1_chat_proto_goTypes = []any{
	(*GetDirectChatRequest)(nil),  // 0: chat.v1.GetDirectChatRequest
	(*GetDirectChatResponse)(nil), // 1: chat.v1.GetDirectChatResponse
	(*GetChatRequest)(nil),        // 2: chat.v1.GetChatRequest
	(*GetChatResponse)(nil),       // 3: chat.v1.GetChatResponse
	(*GetUserChatsRequest)(nil),   // 4: chat.v1.GetUserChatsRequest
	(*GetUserChatsResponse)(nil),  // 5: chat.v1.GetUserChatsResponse
	(*CreateDirectRequest)(nil),   // 6: chat.v1.CreateDirectRequest
	(*CreateDirectResponse)(nil),  // 7: chat.v1.CreateDirectResponse
	(*CreateGroupRequest)(nil),    // 8: chat.v1.CreateGroupRequest
	(*CreateGroupResponse)(nil),   // 9: chat.v1.CreateGroupResponse
	(*CreateChannelRequest)(nil),  // 10: chat.v1.CreateChannelRequest
	(*CreateChannelResponse)(nil), // 11: chat.v1.CreateChannelResponse
	(*JoinChatRequest)(nil),       // 12: chat.v1.JoinChatRequest
	(*JoinChatResponse)(nil),      // 13: chat.v1.JoinChatResponse
	(*v1.Chat)(nil),               // 14: model.chat.v1.Chat
}
var file_protobuf_chat_v1_chat_proto_depIdxs = []int32{
	14, // 0: chat.v1.GetDirectChatResponse.chat:type_name -> model.chat.v1.Chat
	14, // 1: chat.v1.GetChatResponse.chat:type_name -> model.chat.v1.Chat
	14, // 2: chat.v1.GetUserChatsResponse.chats:type_name -> model.chat.v1.Chat
	14, // 3: chat.v1.CreateDirectResponse.chat:type_name -> model.chat.v1.Chat
	14, // 4: chat.v1.CreateGroupResponse.chat:type_name -> model.chat.v1.Chat
	14, // 5: chat.v1.CreateChannelResponse.chat:type_name -> model.chat.v1.Chat
	14, // 6: chat.v1.JoinChatResponse.chat:type_name -> model.chat.v1.Chat
	2,  // 7: chat.v1.ChatService.GetChat:input_type -> chat.v1.GetChatRequest
	0,  // 8: chat.v1.ChatService.GetDirectChat:input_type -> chat.v1.GetDirectChatRequest
	4,  // 9: chat.v1.ChatService.GetUserChats:input_type -> chat.v1.GetUserChatsRequest
	6,  // 10: chat.v1.ChatService.CreateDirect:input_type -> chat.v1.CreateDirectRequest
	8,  // 11: chat.v1.ChatService.CreateGroup:input_type -> chat.v1.CreateGroupRequest
	10, // 12: chat.v1.ChatService.CreateChannel:input_type -> chat.v1.CreateChannelRequest
	12, // 13: chat.v1.ChatService.JoinChat:input_type -> chat.v1.JoinChatRequest
	3,  // 14: chat.v1.ChatService.GetChat:output_type -> chat.v1.GetChatResponse
	1,  // 15: chat.v1.ChatService.GetDirectChat:output_type -> chat.v1.GetDirectChatResponse
	5,  // 16: chat.v1.ChatService.GetUserChats:output_type -> chat.v1.GetUserChatsResponse
	7,  // 17: chat.v1.ChatService.CreateDirect:output_type -> chat.v1.CreateDirectResponse
	9,  // 18: chat.v1.ChatService.CreateGroup:output_type -> chat.v1.CreateGroupResponse
	11, // 19: chat.v1.ChatService.CreateChannel:output_type -> chat.v1.CreateChannelResponse
	13, // 20: chat.v1.ChatService.JoinChat:output_type -> chat.v1.JoinChatResponse
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_protobuf_chat_v1_chat_proto_init() }
func file_protobuf_chat_v1_chat_proto_init() {
	if File_protobuf_chat_v1_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_chat_v1_chat_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetDirectChatRequest); i {
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
		file_protobuf_chat_v1_chat_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetDirectChatResponse); i {
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
		file_protobuf_chat_v1_chat_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetChatRequest); i {
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
		file_protobuf_chat_v1_chat_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetChatResponse); i {
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
		file_protobuf_chat_v1_chat_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserChatsRequest); i {
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
		file_protobuf_chat_v1_chat_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserChatsResponse); i {
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
		file_protobuf_chat_v1_chat_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CreateDirectRequest); i {
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
		file_protobuf_chat_v1_chat_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CreateDirectResponse); i {
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
		file_protobuf_chat_v1_chat_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*CreateGroupRequest); i {
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
		file_protobuf_chat_v1_chat_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*CreateGroupResponse); i {
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
		file_protobuf_chat_v1_chat_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*CreateChannelRequest); i {
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
		file_protobuf_chat_v1_chat_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*CreateChannelResponse); i {
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
		file_protobuf_chat_v1_chat_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*JoinChatRequest); i {
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
		file_protobuf_chat_v1_chat_proto_msgTypes[13].Exporter = func(v any, i int) any {
			switch v := v.(*JoinChatResponse); i {
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
			RawDescriptor: file_protobuf_chat_v1_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_chat_v1_chat_proto_goTypes,
		DependencyIndexes: file_protobuf_chat_v1_chat_proto_depIdxs,
		MessageInfos:      file_protobuf_chat_v1_chat_proto_msgTypes,
	}.Build()
	File_protobuf_chat_v1_chat_proto = out.File
	file_protobuf_chat_v1_chat_proto_rawDesc = nil
	file_protobuf_chat_v1_chat_proto_goTypes = nil
	file_protobuf_chat_v1_chat_proto_depIdxs = nil
}
