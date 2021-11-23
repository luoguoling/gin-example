package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
)

// CommunityHandler 升级版帖子列表接口
// @Summary 帖子列表接口
// @Description 查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Param name query string false "查询参数"
// @Success 200 {object} _Community
// @Router /community [get]
func CommunityHandler(c *gin.Context) {
	//获取分类列表
	communityList, err := logic.GetCommunityList()
	fmt.Println("社区数据是....")
	fmt.Println(communityList)
	if err != nil {
		zap.L().Error("GetCommunity() fail", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, communityList)

}

// CommunityDetailHandler 升级版帖子列表详情接口
// @Summary 帖子列表接口
// @Description 查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Param id query string true "查询参数"
// @Success 200 {object} _CommunityDetail
// @Router /community/{id} [get]
// 获取帖子详情
func CommunityDetailHandler(c *gin.Context) {
	//获取社区id
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	fmt.Sprintf("分类id:%d", id)
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
