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

func (cu *UserFriend) GetUserFriendList() ([]map[string]interface{}, error) {
	var userFriendList []map[string]interface{}
	err := global.DB.Table("chat_user_friend AS a").Where("a.user_id = ?", cu.UserID).Select(
		"a.friend_id", "a.group_id", "b.name group_name").Joins(
		"left join chat_group b on a.group_id = b.id",
	).Scan(&userFriendList).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return []map[string]interface{}{}, nil
	}
	return userFriendList, nil
}
