package initialize

import (
	"github.com/gin-gonic/gin"

	"user_api/middlerwares"
	"user_api/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	Router.Use(middlewares.Cors())

	ApiRouter := Router.Group("/v1")
	router.InitUserRouter(ApiRouter)

	return Router
}
