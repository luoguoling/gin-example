package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/logic"
	"web_app/models"
)

func CreatePostHandler(c *gin.Context) {
	//接收参数
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJson(p) error", zap.Any("err", err))
		zap.L().Error("参数获取错误", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//从c获取道当前用户得id
	userID, err := getCurrentUserId(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	//创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return

	}

	//返回
	ResponseSuccess(c, nil)

}
