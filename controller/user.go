package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"
)

//注册
func SignUpHandler(c *gin.Context) {
	// 1.参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数错误,返回响应
		zap.L().Error("SignUp with valid param1", zap.Error(err))
		//判断err是不是validator.ValidationErrors错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	fmt.Println(p.Username, p.RePassword)
	// 2.业务处理
	err := logic.SignUp(p)
	// 3.返回响应
	if errors.Is(err, mysql.ErrorUserExit) {
		zap.L().Error("用户已经存在!!!,不要重复注册!!!")
		ResponseError(c, CodeUserExist)
		return
	}
	if err != nil {
		zap.L().Error("mysql.Register() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)

}

//登录
func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	fmt.Println("登录开始了.....")
	fmt.Println(c.RemoteIP())
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("Login with valid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//业务逻辑处理
	fmt.Println("用户名是:")
	fmt.Println(p.Username)
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID), // id值大于1<<53-1  int64类型的最大值是1<<63-1
		"user_name": user.Username,
		"token":     user.Token,
	})

}
