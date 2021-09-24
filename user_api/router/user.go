package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"user_api/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	zap.S().Infof("配置用户相关router")
	UserRouterGroup := Router.Group("user")
	{
		UserRouterGroup.POST("login", api.Login)
		UserRouterGroup.POST("register", api.Register)
	}
}
