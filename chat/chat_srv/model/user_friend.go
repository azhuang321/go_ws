package model

import (
	"chat_srv/global"
	"chat_srv/model/dao"
	"go.uber.org/zap"
)

func AddFriend(userId, friendId, groupId uint32) bool {
	tx := global.DB.Begin()
	userFriendModel1 := &dao.UserFriend{
		UserID:   userId,
		FriendID: friendId,
		GroupID:  groupId,
	}
	if err := userFriendModel1.AddUserFriend(); err != nil {
		zap.S().Errorf("添加好友失败:%s", err.Error())
		return false
	}
	userFriendModel2 := &dao.UserFriend{
		UserID:   friendId,
		FriendID: userId,
	}
	if err := userFriendModel2.AddUserFriend(); err != nil {
		tx.Rollback()
		zap.S().Errorf("添加好友失败:%s", err.Error())
		return false
	}
	tx.Commit()
	return true
}

func GetUserFriendList(userId uint32) ([]*dao.Group,[]*dao.UserFriend,error) {
	userFiend := dao.UserFriend{UserID: userId}
	userGroup,userFriendList, err := userFiend.GetUserFriendList()
	if err != nil {
		zap.S().Errorf("查询好友失败:%s", err.Error())
		return nil,nil, err
	}
	return userGroup,userFriendList, nil
}
