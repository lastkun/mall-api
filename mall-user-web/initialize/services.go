package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mall-api/mall-user-web/global"
	"mall-api/mall-user-web/proto"
)

func InitService() {
	serviceHost := ""
	servicePort := 0

	//从注册中心拉取grpc服务信息
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.CC.Host,
		global.ServerConfig.CC.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	services, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.Usc.Name))
	if err != nil {
		panic(err)
	}

	for _, v := range services {
		serviceHost = v.Address
		servicePort = v.Port
		break
	}

	if serviceHost == "" {
		zap.S().Fatal("用户grpc服务连接失败")
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", serviceHost, servicePort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("连接用户grpc服务失败",
			"msg", err.Error())
	}

	userClient := proto.NewUserClient(conn)

	global.UserServiceClient = userClient
}
