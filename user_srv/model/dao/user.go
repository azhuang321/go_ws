package dao

import "user_srv/global"

type Users struct {
	Base
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *Users) CreateUser() (userId uint, err error) {
	err = global.DB.Create(u).Error
	if err != nil {
		return 0, err
	}
	return u.ID, nil
}
