package initialize

import (
	"github.com/gin-gonic/gin"

	"ws_srv/middlerwares"
	"ws_srv/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	Router.Use(middlewares.Cors())

	ApiRouter := Router.Group("/v1")
	router.InitWsRouter(ApiRouter)

	return Router
}
