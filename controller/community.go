package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
)

func GetCommunityHandler(c *gin.Context) {
	//获取分类列表
	data, err := logic.GetCommunityList()
	fmt.Println(data)
	if err != nil {
		zap.L().Error("GetCommunity() fail", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)

}

// 获取帖子详情
func CommunityDetailHandler(c *gin.Context) {
	//获取社区id
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//通过id查询详情
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, data)

}
