package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"user_srv/model"
	"user_srv/proto/gen/userpb"
)

type UserService struct{}

func (u UserService) CreateUser(ctx context.Context, request *userpb.CreateUserRequest) (*emptypb.Empty, error) {
	_, err := model.CreateUser(request.Mobile, request.Password)
	if err != nil {
		return nil, status.Error(codes.Unknown, "数据库错误")
	}
	return &emptypb.Empty{}, nil
}

func (u UserService) IsExistUser(ctx context.Context, request *userpb.UserRequest) (*userpb.IsExistUserResponse, error) {
	isExist, err := model.IsExistUser(request.Mobile)
	if err != nil {
		return nil, status.Error(codes.Unknown, "数据库错误")
	}
	return &userpb.IsExistUserResponse{IsExist: isExist}, nil
}

func (u UserService) GetUserInfo(ctx context.Context, request *userpb.UserRequest) (*userpb.GetUserInfoResponse, error) {
	userInfo, err := model.GetUserInfo(request.Mobile)
	if err != nil {
		return nil, status.Error(codes.Unknown, "数据库错误")
	}
	return &userpb.GetUserInfoResponse{
		Id:       uint32(userInfo.ID),
		Mobile:   userInfo.Mobile,
		Password: userInfo.Password,
		Avatar:   userInfo.Avatar,
		Sign:     userInfo.Sign,
	}, nil
}

func (u UserService) CheckPwd(ctx context.Context, request *userpb.CreateUserRequest) (*userpb.CheckPwdResponse, error) {
	isRight, userInfo, err := model.CheckPwd(request.Mobile, request.Password)
	if err != nil {
		return nil, status.Error(codes.Unknown, "数据库错误")
	}
	return &userpb.CheckPwdResponse{IsRight: isRight, UserInfo: &userpb.GetUserInfoResponse{
		Id:     uint32(userInfo.ID),
		Mobile: userInfo.Password,
	}}, nil
}

func (u UserService) GetUsersInfo(ctx context.Context, req *userpb.GetUsersInfoReq) (*userpb.GetUsersInfoResp, error) {
	usersInfo, err := model.GetUsersInfo(req.Ids)
	if err != nil {
		return nil, status.Error(codes.Unknown, "数据库错误")
	}
	var returnUsersInfo []*userpb.GetUserInfoResponse
	for _, val := range usersInfo {
		userInfo := &userpb.GetUserInfoResponse{}
		userInfo.Id = val.ID
		userInfo.Mobile = val.Mobile
		userInfo.Avatar = val.Avatar
		userInfo.Sign = val.Sign
		returnUsersInfo = append(returnUsersInfo, userInfo)
	}

	return &userpb.GetUsersInfoResp{UsersInfo: returnUsersInfo}, nil
}
