package initialize

import (
	"github.com/gin-gonic/gin"

	globalMiddlewares "ws_srv/middlerwares/global"
	"ws_srv/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	Router.Use(globalMiddlewares.Cors())

	ApiRouter := Router.Group("/v1")
	router.InitWsRouter(ApiRouter)

	return Router
}
