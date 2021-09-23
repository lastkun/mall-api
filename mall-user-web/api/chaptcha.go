package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"mall-api/mall-user-web/forms"
	"mall-api/mall-user-web/global"
	"mall-api/mall-user-web/utils"
	"net/http"
	"time"
)

var store = base64Captcha.DefaultMemStore

//获取图形验证码
func GetCaptcha(ctx *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("生成验证码错误,: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成验证码错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath":   b64s,
	})
}

//注册验证码发送 --模拟
func SendVerCode(ctx *gin.Context) {
	tempCode := 123456 //临时验证码 阿里云短信服务无法通过

	form := forms.ValCodeForm{}
	if err := ctx.ShouldBind(&form); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	//生成验证码同时存放到redis中
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RC.Host, global.ServerConfig.RC.Port),
	})

	err := redisClient.Set(context.Background(), form.Mobile, tempCode, time.Duration(global.ServerConfig.RC.Expire)*time.Second).Err()
	if err != nil {
		zap.S().Errorf("验证码发送 SendVerCode :", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "redis服务出现故障，请检查",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "验证码发送成功",
	})
}
