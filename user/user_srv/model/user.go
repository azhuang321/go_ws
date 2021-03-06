package model

import (
	"go.uber.org/zap"
	"user_srv/model/dao"
	"user_srv/utils"
)

func CreateUser(mobile, password string) (userId uint32, err error) {
	var user dao.Users
	user.Mobile = mobile
	user.Password = utils.MD5(password)
	return user.CreateUser()
}

func IsExistUser(mobile string) (isExist bool, err error) {
	var user dao.Users
	user.Mobile = mobile
	userInfo, err := user.GetUserInfo()
	if err != nil {
		return false, err
	}
	if userInfo.ID > 0 {
		return true, nil
	}
	return false, nil
}

func GetUserInfo(mobile string) (userInfo *dao.Users, err error) {
	userInfo = &dao.Users{}
	userInfo.Mobile = mobile
	userInfo, err = userInfo.GetUserInfo()
	if err != nil {
		return nil, err
	}

	if userInfo.ID > 0 {
		return userInfo, nil
	}
	return &dao.Users{}, nil
}

func CheckPwd(mobile, password string) (isRight bool, userInfo *dao.Users, err error) {
	userInfo = &dao.Users{}
	userInfo.Mobile = mobile
	userInfo, err = userInfo.GetUserInfo()
	if err != nil {
		return false, nil, err
	}
	if userInfo.ID <= 0 || userInfo.Password != utils.MD5(password) {
		return false, userInfo, nil
	}
	return true, userInfo, nil
}

func GetUsersInfo(ids []uint32) ([]*dao.Users, error) {
	var users []*dao.Users
	users, err := dao.GetUsersInfo(ids)
	if err != nil {
		zap.S().Errorf("查询多个用户信息出错:%s", err.Error())
		return nil, err
	}
	return users, nil
}
