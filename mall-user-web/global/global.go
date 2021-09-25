package global

import (
	ut "github.com/go-playground/universal-translator"
	"mall-api/mall-user-web/config"
	"mall-api/mall-user-web/proto"
)

//全局变量 多处复用
var (
	ServerConfig      = &config.ServerConfig{}
	NacosConfig       = &config.NacosConfig{}
	Trans             ut.Translator
	UserServiceClient proto.UserClient
)
