package router

import (
	"github.com/gin-gonic/gin"
	"mall-api/mall-goods-web/api/banner"
	middlewares "mall-api/mall-user-web/middleware"
)

func InitBannerRouter(Router *gin.RouterGroup) {
	BannerRouter := Router.Group("banners")
	{
		BannerRouter.GET("", banner.List)
		BannerRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.CheckAdminAuth(), banner.Delete)
		BannerRouter.POST("",  middlewares.JWTAuth(), middlewares.CheckAdminAuth(), banner.New)
		BannerRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.CheckAdminAuth(), banner.Update)
	}
}
