package router

import (
	"github.com/gin-gonic/gin"
	"mall-api/mall-user-web/api"
)

func BaseRouter(router *gin.RouterGroup) {
	baseRouter := router.Group("base")
	{
		baseRouter.GET("getCaptcha", api.GetCaptcha)
	}
}
