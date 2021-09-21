package model

import (
	"user_srv/model/dao"
	"user_srv/utils"
)

func CreateUser(username, password string) (userId uint, err error) {
	var user dao.Users
	user.Username = username
	user.Password = utils.MD5(password)
	return user.CreateUser()
}
