package initialize

import (
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver" // 集成grpc-consul解析器
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"mall-api/mall-goods-web/global"
	"mall-api/mall-goods-web/proto"
)

//服务发现
//配置了grpc-consul解析器 并使用内置负载均衡策略
func InitService() {
	consulInfo := global.ServerConfig.CC
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.Usc.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("用户grpc服务连接失败")
	}

	GoodsServiceClient := proto.NewGoodsClient(conn)
	global.GoodsServiceClient = GoodsServiceClient
}
