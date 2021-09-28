package handler

import (
	"chat_srv/model"
	"chat_srv/proto/gen/chat_pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatService struct{}

func (c ChatService) AddFriend(ctx context.Context, request *chat_pb.AddFriendRequest) (*emptypb.Empty, error) {
	if !model.AddFriend(request.UserId, request.FriendId, request.GroupId) {
		return nil, status.Error(codes.Unknown, "未知错误")
	}
	return &emptypb.Empty{}, nil
}

func (c ChatService) GetFriendList(ctx context.Context, request *chat_pb.GetFriendListRequest) (*chat_pb.GetFriendListResponse, error) {
	userFriendList, err := model.GetUserFriendList(request.UserId)
	if err != nil {
		return nil, status.Error(codes.Unknown, "未知错误")
	}

	var userFriendListResp []*chat_pb.GetFriendListResponse_UserFriend
	for _, val := range userFriendList {
		userFriend := &chat_pb.GetFriendListResponse_UserFriend{}
		userFriend.FriendId = val["friend_id"].(uint32)
		userFriend.GroupId = val["group_id"].(uint32)
		if val["group_name"] != nil {
			userFriend.GroupName = val["group_name"].(string)
		} else {
			userFriend.GroupName = ""
		}
		userFriendListResp = append(userFriendListResp, userFriend)
	}
	return &chat_pb.GetFriendListResponse{UserFriendLists: userFriendListResp}, nil
}
