package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/logic"
)

func GetCommunityHandler(c *gin.Context) {
	//获取分类列表
	data, err := logic.GetCommunity()
	if err != nil {
		zap.L().Error("GetCommunity() fail", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)

}
