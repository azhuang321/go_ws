package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	middlewares "user_api/middlerwares"

	"user_api/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	zap.S().Infof("配置用户相关router")
	UserRouterNotAuthGroup := Router.Group("user")
	{
		UserRouterNotAuthGroup.POST("login", api.Login)
		UserRouterNotAuthGroup.POST("register", api.Register)
	}
	UserRouterAuthGroup := Router.Group("user").Use(middlewares.JWTAuth())
	{
		UserRouterAuthGroup.POST("info", api.GetUserInfo)
	}
}
