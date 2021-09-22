package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"ws_srv/handler"
)

func InitWsRouter(Router *gin.RouterGroup) {
	zap.S().Infof("配置用户相关router")
	UserRouterGroup := Router.Group("ws")
	{
		UserRouterGroup.GET("test", handler.Test)
	}
}
