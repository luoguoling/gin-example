package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//返回数据结构体

type Response struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code ResCode) {
	rd := &Response{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)

}

//自己指定错误
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	rd := &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)

}

func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &Response{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)

}
