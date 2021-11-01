package models

import "github.com/go-playground/validator/v10"

func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(ParamSignUp)

	if su.Password != su.RePassword {
		// 输出错误提示信息，最后一个参数就是传递的param
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}

//定义请求参数结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password"  binding:"required,eqfield=Password"`
}

//定义登录请求参数结构体
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
