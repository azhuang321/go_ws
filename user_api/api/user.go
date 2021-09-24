package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"user_api/errno"
	"user_api/global"
	"user_api/proto/gen/go/userpb"
	"user_api/utils"

	"user_api/forms"
	"user_api/forms/user"
)

func Register(ctx *gin.Context) {
	registerForm := user.RegisterFrom{}
	if ok := forms.BindJsonAndValidate(ctx, &registerForm); !ok {
		return
	}
	_,err := global.UserSrvClient.CreateUser(ctx, &userpb.UserInfo{
		Mobile:   registerForm.Mobile,
		Password: registerForm.Password,
	})
	if err != nil {
		zap.S().Errorf("用户服务出错:%s",err.Error())
		utils.ErrReturn(ctx,http.StatusInternalServerError,*errno.UserSrvErr)
		return
	}
	utils.OkReturn(ctx,[]string{},"注册成功")
}

func Login(ctx *gin.Context) {

}


