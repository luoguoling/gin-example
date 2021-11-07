package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

//获取当前登陆用户id
var ErrorUserNotLogin = errors.New("错误用户未登录")

const CtxUserIDKey = "userID"

func getCurrentUserId(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return 0, err
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return 0, err
	}
	return

}
