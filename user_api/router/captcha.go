package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"user_api/api"
)

func InitCaptchaRouter(Router *gin.RouterGroup) {
	zap.S().Infof("配置用户相关router")
	UserRouterGroup := Router.Group("/")
	{
		UserRouterGroup.GET("captcha", api.GetCaptcha)
	}
}
