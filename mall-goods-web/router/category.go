package router

import (
	"github.com/gin-gonic/gin"
	"mall-api/mall-goods-web/api/category"
)

func InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("categorys")
	{
		CategoryRouter.GET("", category.List)
		CategoryRouter.DELETE("/:id", category.Delete)
		CategoryRouter.GET("/:id", category.Detail)
		CategoryRouter.POST("", category.New)
		CategoryRouter.PUT("/:id", category.Update)
	}
}