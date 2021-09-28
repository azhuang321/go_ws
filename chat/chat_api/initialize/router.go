package initialize

import (
	"github.com/gin-gonic/gin"

	globalMiddlewares "chat_api/middlewares/global"
	"chat_api/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	Router.Use(globalMiddlewares.Cors())

	ApiRouter := Router.Group("/v1")
	router.InitUserRouter(ApiRouter)

	return Router
}
