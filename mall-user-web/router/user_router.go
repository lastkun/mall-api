package router

import (
	"github.com/gin-gonic/gin"
	middlewares "mall-api/mall-user-web/middleware"

	"mall-api/mall-user-web/api"
)

func UserRouter(router *gin.RouterGroup) {
	//对某一组url加上权限验证 router.Group("user").Use()
	userRouter := router.Group("user")
	{
		userRouter.GET("list", middlewares.JWTAuth(), middlewares.CheckAdminAuth(), api.GetUserList)
		userRouter.POST("loginByPassword", api.LoginByPassword)
	}

}
