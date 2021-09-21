package api

import (
	"context"
	"fmt"
	"mall-api/mall-user-web/forms"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"mall-api/mall-user-web/global"
	"mall-api/mall-user-web/proto"
	"mall-api/mall-user-web/utils"
)

func GetUserList(ctx *gin.Context) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.Usc.Name, global.ServerConfig.Usc.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("连接用户grpc服务失败",
			"msg", err.Error())
	}

	client := proto.NewUserClient(conn)
	//获取参数中的页码页长
	pn, _ := strconv.Atoi(ctx.DefaultQuery("pn", "1"))
	pSize, _ := strconv.Atoi(ctx.DefaultQuery("psize", "10"))

	resp, err := client.GetUserList(context.Background(), &proto.PageInfo{
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
}
