package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
	"web_app/models"
)

func CreatePostHandler(c *gin.Context) {
	fmt.Println("开始处理createPost请求")
	//接收参数
	p := new(models.Post)
	zap.L().Info("controller.CreatePostHandler 接收参数")
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
	fmt.Println(p)
	//创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return

	}

	//返回
	ResponseSuccess(c, nil)

}
func GetPostDetailHandler(c *gin.Context) {
	//获取参数
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	fmt.Println(id)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
	}
	//获取数据
	data, err := logic.GetPostDetail(id)
	if err != nil {
		zap.L().Error("logic.GetPostDetail is fail", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	//返回结果
	ResponseSuccess(c, data)
}
