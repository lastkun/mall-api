package middlewares

import (
	"github.com/gin-gonic/gin"
	"mall-api/mall-user-web/utils"
	"net/http"
)

//管理员权限验证
func CheckAdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*utils.CustomClaims)

		if currentUser.Role != 0 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": "该用户无管理员权限",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
