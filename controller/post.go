package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/models"
)

func CreatePostHandler(c *gin.Context) {
	//接收参数
	if err := c.ShouldBindJSON(&models.Post{}); err != nil {
		zap.L().Error("参数获取错误", zap.Error(err))
	}
	//写入处理

	//返回

}
