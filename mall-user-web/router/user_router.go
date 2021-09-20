package router

import (
	"github.com/gin-gonic/gin"
	"mall-api/mall-user-web/api"
)

func UserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")

	userRouter.GET("list", api.GetUserList)
}
