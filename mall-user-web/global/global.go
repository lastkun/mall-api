package global

import (
	ut "github.com/go-playground/universal-translator"
	"mall-api/mall-user-web/config"
)

//全局变量 多处复用
var (
	ServerConfig = &config.ServerConfig{}
	Trans        ut.Translator
)