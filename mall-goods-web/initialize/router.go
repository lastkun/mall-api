package initialize

import (
	"github.com/gin-gonic/gin"
	middlewares "mall-api/mall-goods-web/middleware"
	GoodsRouter "mall-api/mall-goods-web/router"
)

//初始化routers
func Routers() *gin.Engine {
	router := gin.Default()
	//解决跨域问题
	router.Use(middlewares.Cors())
	mainGroup := router.Group("/api")
	GoodsRouter.GoodsRouter(mainGroup)
	GoodsRouter.BaseRouter(mainGroup)
	return router
}
