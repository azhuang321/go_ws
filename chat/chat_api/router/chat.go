package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"chat_api/api"
	"chat_api/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	zap.S().Infof("配置用户相关router")
	UserRouterGroup := Router.Group("chat").Use(middlewares.JWTAuth())
	{
		UserRouterGroup.POST("getUserFriendList", api.GetUserFriendList)
	}
}
