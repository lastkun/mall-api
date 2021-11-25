package handler

import (
	"github.com/gin-gonic/gin"
	"mall-api/mall-oss-web/utils"
)

func Token(c *gin.Context) {
	response := utils.Get_policy_token()
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Origin", "*")
	c.String(200, response)
}
