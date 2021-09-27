package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	middlewares "ws_srv/middlerwares"

	"ws_srv/api"
)

func InitWsRouter(Router *gin.RouterGroup) {
	zap.S().Infof("配置用户相关router")
	UserRouterGroup := Router.Group("ws")
	{
		UserRouterGroup.GET("test", middlewares.JWTAuth(), api.Test)
	}
}
