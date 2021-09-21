package initialize

import (
	"github.com/gin-gonic/gin"
	userRouter "mall-api/mall-user-web/router"
)

//初始化routers
func Routers() *gin.Engine {
	router := gin.Default()
	mainGroup := router.Group("/api")
	userRouter.UserRouter(mainGroup)
	return router
}
