package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/satori/go.uuid"
	"go.uber.org/zap"

	"mall-api/mall-goods-web/global"
	"mall-api/mall-goods-web/initialize"
	"mall-api/mall-goods-web/utils/consul"
)

func main() {
	//初始化logger
	initialize.InitGlobalLogger()
	//初始化配置
	initialize.InitConfig()
	//初始化翻译器--用于表单验证
	initialize.InitTrans("zh")
	//初始化routers
	routers := initialize.Routers()
	//初始化grpc服务client
	initialize.InitService()

	//服务注册--consul
	registerClient := consul.NewRegisterClient(global.ServerConfig.CC.Host, global.ServerConfig.CC.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := registerClient.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("[err] goods-api 服务注册失败", err.Error())
	}

	zap.S().Infof("[start] goods-api port: %d", global.ServerConfig.Port)

	go func() {
		err = routers.Run(fmt.Sprintf(":%d", global.ServerConfig.Port))
		if err != nil {
			zap.S().Panic("[err] can not start goods-api ", err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	err = registerClient.DeRegister(serviceId)

	if err != nil {
		zap.S().Panic("[err] goods-api 注销服务失败")
	} else {
		zap.S().Info("goods-api 注销成功")
	}

}
