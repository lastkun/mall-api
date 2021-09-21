package main

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"mall-api/mall-user-web/initialize"

	"mall-api/mall-user-web/forms"
	"mall-api/mall-user-web/global"
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

	//绑定自定义validator
	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	v.RegisterValidation("mobile", forms.MobileValidator)
	//}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", forms.MobileValidator)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "非法的手机号码", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	zap.S().Debugf("[start] user-api port: %d", global.ServerConfig.Port)

	err := routers.Run(fmt.Sprintf(":%d", global.ServerConfig.Port))
	if err != nil {
		zap.S().Panic("[err] can not start user-api ", err.Error())
	}

}
