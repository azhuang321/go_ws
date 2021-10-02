package dao

import (
	"chat_srv/global"
	"errors"
	"gorm.io/gorm"
)

type UserFriend struct {
	Base
	UserID   uint32 `json:"user_id" gorm:"column:user_id"`
	FriendID uint32 `json:"friend_id" gorm:"column:friend_id"`
	GroupID  uint32 `json:"group_id" gorm:"column:group_id"`
	Status   uint8  `json:"status" gorm:"column:status"` // 好友关系(-1:拉黑;0删除;1拉黑)
}

func (cu *UserFriend) TableName() string {
	return "chat_user_friend"
}

func (cu *UserFriend) AddUserFriend() error {
	if err := global.DB.Create(cu).Error; err != nil {
		return err
	}
	return nil
}

func (cu *UserFriend) GetUserFriendList() ([]*Group,[]*UserFriend, error) {
	group := &Group{UserID:cu.UserID}
	userGroups,err := group.UserGroup()
	if err != nil {
		return []*Group{},[]*UserFriend{}, nil
	}
	groupIds := make([]uint32, 0)
	for _,val := range userGroups {
		groupIds = append(groupIds,val.ID)
	}

	var userFriendList []*UserFriend
	err = global.DB.Model(cu).Where("group_id in ?",groupIds).Where("user_id = ?",cu.UserID).Find(&userFriendList).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		return nil,nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return []*Group{},[]*UserFriend{}, nil
	}
	return userGroups,userFriendList, nil
}
