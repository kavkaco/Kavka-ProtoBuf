// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: protobuf/search/v1/search.proto

package searchv1

import (
	v1 "github.com/kavkaco/Kavka-ProtoBuf/gen/go/protobuf/model/chat/v1"
	v11 "github.com/kavkaco/Kavka-ProtoBuf/gen/go/protobuf/model/user/v1"
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

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_search_v1_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_search_v1_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_search_v1_search_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *SearchResponse_SearchResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_search_v1_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_search_v1_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_search_v1_search_proto_rawDescGZIP(), []int{1}
}

func (x *SearchResponse) GetResult() *SearchResponse_SearchResult {
	if x != nil {
		return x.Result
	}
	return nil
}

type SearchResponse_SearchResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chats []*v1.Chat  `protobuf:"bytes,1,rep,name=chats,proto3" json:"chats,omitempty"`
	Users []*v11.User `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *SearchResponse_SearchResult) Reset() {
	*x = SearchResponse_SearchResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_search_v1_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse_SearchResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse_SearchResult) ProtoMessage() {}

func (x *SearchResponse_SearchResult) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_search_v1_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse_SearchResult.ProtoReflect.Descriptor instead.
func (*SearchResponse_SearchResult) Descriptor() ([]byte, []int) {
	return file_protobuf_search_v1_search_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SearchResponse_SearchResult) GetChats() []*v1.Chat {
	if x != nil {
		return x.Chats
	}
	return nil
}

func (x *SearchResponse_SearchResult) GetUsers() []*v11.User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_protobuf_search_v1_search_proto protoreflect.FileDescriptor

var file_protobuf_search_v1_search_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x68, 0x61,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0xb6, 0x01, 0x0a, 0x0e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x64, 0x0a, 0x0c,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x29, 0x0a, 0x05,
	0x63, 0x68, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x05, 0x63, 0x68, 0x61, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x32, 0x50, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0xa7, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x61, 0x76, 0x6b, 0x61, 0x63, 0x6f, 0x2f, 0x4b, 0x61, 0x76, 0x6b, 0x61, 0x2d,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f,
	0x76, 0x31, 0x3b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58,
	0x58, 0xaa, 0x02, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_search_v1_search_proto_rawDescOnce sync.Once
	file_protobuf_search_v1_search_proto_rawDescData = file_protobuf_search_v1_search_proto_rawDesc
)

func file_protobuf_search_v1_search_proto_rawDescGZIP() []byte {
	file_protobuf_search_v1_search_proto_rawDescOnce.Do(func() {
		file_protobuf_search_v1_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_search_v1_search_proto_rawDescData)
	})
	return file_protobuf_search_v1_search_proto_rawDescData
}

var file_protobuf_search_v1_search_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protobuf_search_v1_search_proto_goTypes = []any{
	(*SearchRequest)(nil),               // 0: search.v1.SearchRequest
	(*SearchResponse)(nil),              // 1: search.v1.SearchResponse
	(*SearchResponse_SearchResult)(nil), // 2: search.v1.SearchResponse.SearchResult
	(*v1.Chat)(nil),                     // 3: model.chat.v1.Chat
	(*v11.User)(nil),                    // 4: model.user.v1.User
}
var file_protobuf_search_v1_search_proto_depIdxs = []int32{
	2, // 0: search.v1.SearchResponse.result:type_name -> search.v1.SearchResponse.SearchResult
	3, // 1: search.v1.SearchResponse.SearchResult.chats:type_name -> model.chat.v1.Chat
	4, // 2: search.v1.SearchResponse.SearchResult.users:type_name -> model.user.v1.User
	0, // 3: search.v1.SearchService.Search:input_type -> search.v1.SearchRequest
	1, // 4: search.v1.SearchService.Search:output_type -> search.v1.SearchResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protobuf_search_v1_search_proto_init() }
func file_protobuf_search_v1_search_proto_init() {
	if File_protobuf_search_v1_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_search_v1_search_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SearchRequest); i {
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
		file_protobuf_search_v1_search_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SearchResponse); i {
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
		file_protobuf_search_v1_search_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SearchResponse_SearchResult); i {
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
			RawDescriptor: file_protobuf_search_v1_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_search_v1_search_proto_goTypes,
		DependencyIndexes: file_protobuf_search_v1_search_proto_depIdxs,
		MessageInfos:      file_protobuf_search_v1_search_proto_msgTypes,
	}.Build()
	File_protobuf_search_v1_search_proto = out.File
	file_protobuf_search_v1_search_proto_rawDesc = nil
	file_protobuf_search_v1_search_proto_goTypes = nil
	file_protobuf_search_v1_search_proto_depIdxs = nil
}
