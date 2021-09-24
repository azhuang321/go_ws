package dao

import (
	"errors"
	"gorm.io/gorm"
	"user_srv/global"
)

type Users struct {
	Base
	Mobile string `json:"mobile"`
	Password string `json:"password"`
}

func (u *Users) CreateUser() (userId uint, err error) {
	err = global.DB.Create(u).Error
	if err != nil {
		return 0, err
	}
	return u.ID, nil
}

func (u *Users)GetUserInfo()(*Users,error){
	err := global.DB.Model(u).Where("mobile = ?",u.Mobile).First(u).Error
	if err != nil && !errors.Is(err,gorm.ErrRecordNotFound) {
		return nil, err
	}
	return u,nil
}


