package initialize

import (
	"github.com/gin-gonic/gin"

	globalMiddlewares "user_api/middlerwares/global"
	"user_api/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	Router.Use(globalMiddlewares.Cors())

	ApiRouter := Router.Group("/v1")
	router.InitUserRouter(ApiRouter)
	router.InitCaptchaRouter(ApiRouter)

	return Router
}
