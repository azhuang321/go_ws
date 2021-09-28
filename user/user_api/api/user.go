package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
	"user_api/errno"
	"user_api/global"
	"user_api/proto/gen/go/userpb"
	"user_api/utils"

	"user_api/forms"
)

func Register(ctx *gin.Context) {
	userForm := forms.UserForm{}
	if ok := forms.BindJsonAndValidate(ctx, &userForm, ""); !ok {
		return
	}
	isExist, err := global.UserSrvClient.IsExistUser(ctx, &userpb.UserRequest{
		Mobile: userForm.Mobile,
	})
	if err != nil {
		zap.S().Errorf("用户服务出错:%s", err.Error())
		utils.ErrReturn(ctx, http.StatusInternalServerError, errno.UserSrvErr)
		return
	}
	if isExist.IsExist {
		utils.ErrReturn(ctx, http.StatusAlreadyReported, errno.ErrUserIsExist)
		return
	}
	_, err = global.UserSrvClient.CreateUser(ctx, &userpb.CreateUserRequest{
		Mobile:   userForm.Mobile,
		Password: userForm.Password,
	})
	if err != nil {
		zap.S().Errorf("用户服务出错:%s", err.Error())
		utils.ErrReturn(ctx, http.StatusInternalServerError, errno.UserSrvErr)
		return
	}
	utils.OkReturn(ctx, []string{}, "注册成功")
}

func Login(ctx *gin.Context) {
	userForm := forms.UserForm{}
	if ok := forms.BindJsonAndValidate(ctx, &userForm, "login"); !ok {
		return
	}
	resp, err := global.UserSrvClient.CheckPwd(ctx, &userpb.CreateUserRequest{Mobile: userForm.Mobile, Password: userForm.Password})
	if err != nil {
		zap.S().Errorf("用户服务出错:%s", err.Error())
		utils.ErrReturn(ctx, http.StatusInternalServerError, errno.UserSrvErr)
		return
	}
	if !resp.IsRight {
		utils.ErrReturn(ctx, http.StatusUnauthorized, errno.ErrUserPassword)
		return
	}
	j := utils.NewJWT()
	claims := utils.CustomClaims{
		ID:     resp.UserInfo.Id,
		Mobile: userForm.Mobile,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*24*int64(global.Config.Jwt.ExpireTimeDay),
			Issuer:    "mychat",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		zap.S().Errorf("生成token失败:%s", err.Error())
		rErr := errno.ErrRuntime.ReplaceMsg("生成token失败")
		utils.ErrReturn(ctx, http.StatusInternalServerError, &rErr)
		return
	}
	utils.OkReturn(ctx, gin.H{
		"id":         resp.UserInfo.Id,
		"mobile":     userForm.Mobile,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	}, "")
}

func GetUserInfo(ctx *gin.Context) {
	userForm := &forms.UserForm{}
	if err := forms.BindJsonAndValidate(ctx, userForm, "getUserInfo"); !err {
		return
	}
	resp, err := global.UserSrvClient.GetUserInfo(ctx, &userpb.UserRequest{Mobile: userForm.Mobile})
	if err != nil {
		zap.S().Errorf("用户服务出错:%s", err.Error())
		utils.ErrReturn(ctx, http.StatusInternalServerError, errno.UserSrvErr)
	}
	resp.Password = ""
	utils.OkReturn(ctx, resp, "")
}
