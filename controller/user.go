package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"web_app/logic"
	"web_app/models"
)

//注册
func SignUpHandler(c *gin.Context) {
	// 1.参数校验
	p := new(models.ParamSignUp)

	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数错误,返回响应
		zap.L().Error("SignUp with valid param", zap.Error(err))
		//判断err是不是validator.ValidationErrors错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数错误",
			"err": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	fmt.Println(p.Username, p.RePassword)
	// 2.业务处理
	err := logic.SignUp(p)
	// 3.返回响应
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": p,
		})
	}

}

//登录
func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("Login with valid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录请求参数错误",
			"err": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	if err := logic.Login(p); err != nil {
		zap.L().Error("用户登录失败!!!", zap.String("username", p.Username), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功!!!",
		})
	}

}
