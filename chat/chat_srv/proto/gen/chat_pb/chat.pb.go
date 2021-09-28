// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: chat.proto

package chat_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FriendId uint32 `protobuf:"varint,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	GroupId  uint32 `protobuf:"varint,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *AddFriendRequest) Reset() {
	*x = AddFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendRequest) ProtoMessage() {}

func (x *AddFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFriendRequest.ProtoReflect.Descriptor instead.
func (*AddFriendRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *AddFriendRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddFriendRequest) GetFriendId() uint32 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

func (x *AddFriendRequest) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type GetFriendListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetFriendListRequest) Reset() {
	*x = GetFriendListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFriendListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFriendListRequest) ProtoMessage() {}

func (x *GetFriendListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFriendListRequest.ProtoReflect.Descriptor instead.
func (*GetFriendListRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *GetFriendListRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetFriendListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserFriendLists []*GetFriendListResponse_UserFriend `protobuf:"bytes,1,rep,name=user_friend_lists,json=userFriendLists,proto3" json:"user_friend_lists,omitempty"`
}

func (x *GetFriendListResponse) Reset() {
	*x = GetFriendListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFriendListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFriendListResponse) ProtoMessage() {}

func (x *GetFriendListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFriendListResponse.ProtoReflect.Descriptor instead.
func (*GetFriendListResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *GetFriendListResponse) GetUserFriendLists() []*GetFriendListResponse_UserFriend {
	if x != nil {
		return x.UserFriendLists
	}
	return nil
}

type GetFriendListResponse_UserFriend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FriendId  uint32 `protobuf:"varint,1,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	GroupId   uint32 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	GroupName string `protobuf:"bytes,3,opt,name=group_name,json=groupName,proto3" json:"group_name"` // @gotags: json:"group_name"
}

func (x *GetFriendListResponse_UserFriend) Reset() {
	*x = GetFriendListResponse_UserFriend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFriendListResponse_UserFriend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFriendListResponse_UserFriend) ProtoMessage() {}

func (x *GetFriendListResponse_UserFriend) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFriendListResponse_UserFriend.ProtoReflect.Descriptor instead.
func (*GetFriendListResponse_UserFriend) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GetFriendListResponse_UserFriend) GetFriendId() uint32 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

func (x *GetFriendListResponse_UserFriend) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GetFriendListResponse_UserFriend) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x10, 0x41, 0x64, 0x64,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x2f,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0xcb, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x11, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x1a, 0x63, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x7e, 0x0a,
	0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x12, 0x11, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3e, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x5a,
	0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x70, 0x62, 0x3b, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chat_proto_goTypes = []interface{}{
	(*AddFriendRequest)(nil),                 // 0: AddFriendRequest
	(*GetFriendListRequest)(nil),             // 1: GetFriendListRequest
	(*GetFriendListResponse)(nil),            // 2: GetFriendListResponse
	(*GetFriendListResponse_UserFriend)(nil), // 3: GetFriendListResponse.UserFriend
	(*emptypb.Empty)(nil),                    // 4: google.protobuf.Empty
}
var file_chat_proto_depIdxs = []int32{
	3, // 0: GetFriendListResponse.user_friend_lists:type_name -> GetFriendListResponse.UserFriend
	0, // 1: Chat.AddFriend:input_type -> AddFriendRequest
	1, // 2: Chat.GetFriendList:input_type -> GetFriendListRequest
	4, // 3: Chat.AddFriend:output_type -> google.protobuf.Empty
	2, // 4: Chat.GetFriendList:output_type -> GetFriendListResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFriendRequest); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFriendListRequest); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFriendListResponse); i {
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
		file_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFriendListResponse_UserFriend); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetFriendList(ctx context.Context, in *GetFriendListRequest, opts ...grpc.CallOption) (*GetFriendListResponse, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Chat/AddFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetFriendList(ctx context.Context, in *GetFriendListRequest, opts ...grpc.CallOption) (*GetFriendListResponse, error) {
	out := new(GetFriendListResponse)
	err := c.cc.Invoke(ctx, "/Chat/GetFriendList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	AddFriend(context.Context, *AddFriendRequest) (*emptypb.Empty, error)
	GetFriendList(context.Context, *GetFriendListRequest) (*GetFriendListResponse, error)
}

// UnimplementedChatServer can be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (*UnimplementedChatServer) AddFriend(context.Context, *AddFriendRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (*UnimplementedChatServer) GetFriendList(context.Context, *GetFriendListRequest) (*GetFriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendList not implemented")
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chat/AddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).AddFriend(ctx, req.(*AddFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chat/GetFriendList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetFriendList(ctx, req.(*GetFriendListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFriend",
			Handler:    _Chat_AddFriend_Handler,
		},
		{
			MethodName: "GetFriendList",
			Handler:    _Chat_GetFriendList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}