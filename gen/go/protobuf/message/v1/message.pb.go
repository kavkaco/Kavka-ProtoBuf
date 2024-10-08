// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: protobuf/message/v1/message.proto

package messagev1

import (
	v1 "github.com/kavkaco/Kavka-ProtoBuf/gen/go/protobuf/model/message/v1"
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

type FetchMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *FetchMessagesRequest) Reset() {
	*x = FetchMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_message_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchMessagesRequest) ProtoMessage() {}

func (x *FetchMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_message_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchMessagesRequest.ProtoReflect.Descriptor instead.
func (*FetchMessagesRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_message_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *FetchMessagesRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type FetchMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*v1.Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *FetchMessagesResponse) Reset() {
	*x = FetchMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_message_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchMessagesResponse) ProtoMessage() {}

func (x *FetchMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_message_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchMessagesResponse.ProtoReflect.Descriptor instead.
func (*FetchMessagesResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_message_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *FetchMessagesResponse) GetMessages() []*v1.Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type SendTextMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SendTextMessageRequest) Reset() {
	*x = SendTextMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_message_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTextMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTextMessageRequest) ProtoMessage() {}

func (x *SendTextMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_message_v1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTextMessageRequest.ProtoReflect.Descriptor instead.
func (*SendTextMessageRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_message_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *SendTextMessageRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *SendTextMessageRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type SendTextMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *v1.Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendTextMessageResponse) Reset() {
	*x = SendTextMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_message_v1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTextMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTextMessageResponse) ProtoMessage() {}

func (x *SendTextMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_message_v1_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTextMessageResponse.ProtoReflect.Descriptor instead.
func (*SendTextMessageResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_message_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *SendTextMessageResponse) GetMessage() *v1.Message {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_protobuf_message_v1_message_proto protoreflect.FileDescriptor

var file_protobuf_message_v1_message_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x15, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x16, 0x53, 0x65, 0x6e,
	0x64, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x22, 0x4e, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0xc6, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0f, 0x53,
	0x65, 0x6e, 0x64, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xaf, 0x01, 0x0a, 0x0e, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x76, 0x6b, 0x61, 0x63, 0x6f,
	0x2f, 0x4b, 0x61, 0x76, 0x6b, 0x61, 0x2d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protobuf_message_v1_message_proto_rawDescOnce sync.Once
	file_protobuf_message_v1_message_proto_rawDescData = file_protobuf_message_v1_message_proto_rawDesc
)

func file_protobuf_message_v1_message_proto_rawDescGZIP() []byte {
	file_protobuf_message_v1_message_proto_rawDescOnce.Do(func() {
		file_protobuf_message_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_message_v1_message_proto_rawDescData)
	})
	return file_protobuf_message_v1_message_proto_rawDescData
}

var file_protobuf_message_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protobuf_message_v1_message_proto_goTypes = []any{
	(*FetchMessagesRequest)(nil),    // 0: message.v1.FetchMessagesRequest
	(*FetchMessagesResponse)(nil),   // 1: message.v1.FetchMessagesResponse
	(*SendTextMessageRequest)(nil),  // 2: message.v1.SendTextMessageRequest
	(*SendTextMessageResponse)(nil), // 3: message.v1.SendTextMessageResponse
	(*v1.Message)(nil),              // 4: model.message.v1.Message
}
var file_protobuf_message_v1_message_proto_depIdxs = []int32{
	4, // 0: message.v1.FetchMessagesResponse.messages:type_name -> model.message.v1.Message
	4, // 1: message.v1.SendTextMessageResponse.message:type_name -> model.message.v1.Message
	0, // 2: message.v1.MessageService.FetchMessages:input_type -> message.v1.FetchMessagesRequest
	2, // 3: message.v1.MessageService.SendTextMessage:input_type -> message.v1.SendTextMessageRequest
	1, // 4: message.v1.MessageService.FetchMessages:output_type -> message.v1.FetchMessagesResponse
	3, // 5: message.v1.MessageService.SendTextMessage:output_type -> message.v1.SendTextMessageResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protobuf_message_v1_message_proto_init() }
func file_protobuf_message_v1_message_proto_init() {
	if File_protobuf_message_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_message_v1_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FetchMessagesRequest); i {
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
		file_protobuf_message_v1_message_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FetchMessagesResponse); i {
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
		file_protobuf_message_v1_message_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SendTextMessageRequest); i {
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
		file_protobuf_message_v1_message_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SendTextMessageResponse); i {
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
			RawDescriptor: file_protobuf_message_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_message_v1_message_proto_goTypes,
		DependencyIndexes: file_protobuf_message_v1_message_proto_depIdxs,
		MessageInfos:      file_protobuf_message_v1_message_proto_msgTypes,
	}.Build()
	File_protobuf_message_v1_message_proto = out.File
	file_protobuf_message_v1_message_proto_rawDesc = nil
	file_protobuf_message_v1_message_proto_goTypes = nil
	file_protobuf_message_v1_message_proto_depIdxs = nil
}
