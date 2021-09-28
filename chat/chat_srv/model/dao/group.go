package dao

import (
	"chat_srv/global"
	"errors"
	"gorm.io/gorm"
)

type Group struct {
	Base
	UserID uint32 `json:"user_id" gorm:"column:user_id"`
	Name   string `json:"name" gorm:"column:name"`
}

func (g *Group) TableName() string {
	return "chat_group"
}

func (g *Group) UserGroup() ([]*Group,error) {
	userGroups := make([]*Group,0)
	if err := global.DB.Where("user_id = ?",g.UserID).Find(&userGroups).Error; err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound) {
			return userGroups, nil
		}
		return nil, err
	}
	return userGroups, nil
}