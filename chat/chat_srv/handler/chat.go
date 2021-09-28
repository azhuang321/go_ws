package handler

import (
	"chat_srv/model"
	"chat_srv/proto/gen/chat_pb"
	"chat_srv/utils"
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
	userGroup,userFriendList, err := model.GetUserFriendList(request.UserId)
	if err != nil {
		return nil, status.Error(codes.Unknown, "未知错误")
	}
	utils.PrettyPrint(userGroup)
	utils.PrettyPrint(userFriendList)

	userFriendListResp := make([]*chat_pb.GetFriendListResponse_UserFriend,0)
	for _, val := range userFriendList {
		friendInfo := &chat_pb.GetFriendListResponse_UserFriend{}
		friendInfo.FriendId = val.UserID
		friendInfo.GroupId = val.GroupID
		userFriendListResp = append(userFriendListResp,friendInfo)
	}

	userGroupResp := make([]*chat_pb.GetFriendListResponse_UserGroup,0)
	for _, val := range userGroup {
		groupInfo := &chat_pb.GetFriendListResponse_UserGroup{}
		groupInfo.Id = val.ID
		groupInfo.Name = val.Name
		userGroupResp = append(userGroupResp, groupInfo)
	}

	return &chat_pb.GetFriendListResponse{UserFriendLists: userFriendListResp,UserGroup: userGroupResp}, nil
}
