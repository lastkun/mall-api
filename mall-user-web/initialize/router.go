package initialize

import (
	"github.com/gin-gonic/gin"
	middlewares "mall-api/mall-user-web/middleware"
	userRouter "mall-api/mall-user-web/router"
)

//初始化routers
func Routers() *gin.Engine {
	router := gin.Default()
	//解决跨域问题
	router.Use(middlewares.Cors())
	mainGroup := router.Group("/api")
	userRouter.UserRouter(mainGroup)
	userRouter.BaseRouter(mainGroup)
	return router
}
