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

func (u UserService) CreateUser(ctx context.Context, request *userpb.UserInfo) (*emptypb.Empty, error) {
	_, err := model.CreateUser(request.Username, request.Password)
	if err != nil {
		return nil, status.Error(codes.Unknown, "数据库错误")
	}
	return &emptypb.Empty{}, nil
}
