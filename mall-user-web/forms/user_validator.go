package forms

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

//自定义验证器
//验证手机号码格式
func MobileValidator(v validator.FieldLevel) bool {
	mobile := v.Field().String()
	ok, _ := regexp.MatchString(`^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$`, mobile)
	return ok
}
