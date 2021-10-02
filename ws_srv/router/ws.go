package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	middlewares "ws_srv/middlerwares"

	"ws_srv/api"
)

func InitWsRouter(Router *gin.RouterGroup) {
	zap.S().Infof("配置websocket router")
	Router.GET("websocket", middlewares.JWTAuth(), api.WebSocket)
}
