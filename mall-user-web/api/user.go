package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	middlewares "mall-api/mall-user-web/middleware"

	"mall-api/mall-user-web/forms"
	"mall-api/mall-user-web/global"
	"mall-api/mall-user-web/proto"
	"mall-api/mall-user-web/utils"
)

func GetUserList(ctx *gin.Context) {
	userClient := global.UserServiceClient
	//获取参数中的页码页长
	pn, _ := strconv.Atoi(ctx.DefaultQuery("pn", "1"))
	pSize, _ := strconv.Atoi(ctx.DefaultQuery("psize", "10"))

	resp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pn),
		PSize: uint32(pSize),
	})

	if err != nil {
		zap.S().Errorw("GetUserList 查询用户列表失败",
			"msg", err.Error())
		utils.StatusCodesHandler(err, ctx)
		return
	}

	result := make([]interface{}, 0)

	for _, value := range resp.Data {
		birthday := global.JsonTime(time.Unix(int64(value.Birthday), 0))
		user := global.UserResponse{
			Id:       value.Id,
			Nickname: value.Nickname,
			Birthday: birthday,
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}
		result = append(result, user)
	}

	ctx.JSON(http.StatusOK, result)
}

//通过账号密码登录
//请求底层grpc服务前需要先进行表单验证-- ctx.ShouldBindJSON
func LoginByPassword(ctx *gin.Context) {
	form := forms.LoginByPwdForm{}
	if err := ctx.ShouldBind(&form); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	//验证码校验
	if !store.Verify(form.CaptchaId, form.Captcha, true) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}

	client := global.UserServiceClient

	response, err := client.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: form.Mobile,
	})

	if err != nil {
		zap.S().Errorw("登录失败",
			"msg", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "登录失败",
		})
		return
	}

	checkResponse, err := client.CheckPassword(context.Background(), &proto.CheckPasswordRequest{
		Password:     form.PassWord,
		EncryptedPwd: response.Password,
	})

	if err != nil {
		zap.S().Errorw("检查密码失败",
			"msg", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "检查密码失败",
		})
		return
	}

	check := checkResponse.Success
	if !check {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "密码错误",
		})
	} else {
		//生成token
		j := middlewares.NewJWT()
		claims := utils.CustomClaims{
			ID:       response.Id,
			NickName: response.Nickname,
			Role:     response.Role,
			StandardClaims: jwt.StandardClaims{
				NotBefore: time.Now().Unix(),            //生效时间
				ExpiresAt: time.Now().Unix() + 60*60*24, //过期时间
				Issuer:    "mall",
			},
		}
		token, err := j.CreateToken(claims)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "生成Token失败",
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"id":           response.Id,
			"nickname":     response.Nickname,
			"token":        token,
			"expired_time": 1000 * (time.Now().Unix() + 60*60*24),
		})
	}

}

//注册
func RegisterByMobile(ctx *gin.Context) {
	form := forms.RegisterForm{}
	if err := ctx.ShouldBind(&form); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	//从redis中取出验证码并校验
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RC.Host, global.ServerConfig.RC.Port),
	})

	val, err := redisClient.Get(context.Background(), form.Mobile).Result()
	if err == redis.Nil {
		zap.S().Infof("注册-验证码已过期：不存在key = %s", form.Mobile)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码已过期，请重试",
		})
		return
	}

	//校验失败
	if val != form.Code {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码输入错误，请重试",
		})
		return
	}

	client := global.UserServiceClient

	resp, err := client.AddUser(context.Background(), &proto.AddUserRequest{
		Password: form.PassWord,
		Mobile:   form.Mobile,
		Nickname: form.Nickname,
	})

	if err != nil {
		zap.S().Errorf("注册失败--  AddUser : %s ", err.Error())
		utils.StatusCodesHandler(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("用户：%s : 您已注册成功，请登录", resp.Mobile),
	})
}
