package router

import (
	"github.com/gin-gonic/gin"
	"mall-api/mall-goods-web/api/goods"
	middlewares "mall-api/mall-user-web/middleware"
)

func InitGoodsRouter(r *gin.RouterGroup)  {
	routerGroup := r.Group("goods")
	{
		routerGroup.GET("", goods.List)
		routerGroup.POST("", middlewares.JWTAuth(), middlewares.CheckAdminAuth(), goods.New)
		routerGroup.GET("/:id", goods.Detail)
		routerGroup.DELETE("/:id",middlewares.JWTAuth(), middlewares.CheckAdminAuth(), goods.Delete)
		routerGroup.GET("/:id/stocks", goods.Stocks) //获取库存
		routerGroup.PUT("/:id",middlewares.JWTAuth(), middlewares.CheckAdminAuth(), goods.Update)
		routerGroup.PATCH("/:id",middlewares.JWTAuth(), middlewares.CheckAdminAuth(), goods.UpdateStatus) //上架、热门、新品标签 状态
	}
}