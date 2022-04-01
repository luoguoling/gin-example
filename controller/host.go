package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"
	"web_app/pkg/util"
	"web_app/settings"
)

func GetHostHandler(c *gin.Context) {
	//获取用户数据
	hosts, err := logic.GetHost(util.GetPage(c), settings.Conf.PageSize)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回数据
	ResponseSuccess(c, hosts)
}

func GetHostDetailHandler(c *gin.Context) {
	//获取参数
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	fmt.Println("传入参数的值:")
	fmt.Println(id)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
	}
	//查询数据
	host, err := logic.GetHostDetail(id)
	//返回数据
	ResponseSuccess(c, host)

}

// CreateHostHandler 添加Host
func CreateHostHandler(c *gin.Context) {
	//接手校验参数
	host := new(models.Host)
	if err := c.ShouldBind(host); err != nil {
		zap.L().Error("crontroller CreateHostHandler  c.shouldbind err", zap.Any("err", err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//处理数据添加logic
	if err := logic.CreateHost(host); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	//返回结果
	ResponseSuccess(c, nil)
}

func UpdateHostHandler(c *gin.Context) {
	//获取参数
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	fmt.Println("传入参数的值:")
	fmt.Println(id)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
	}
	host, err := mysql.GetHostDetail(id)
	if err != nil {
		zap.L().Error("logic updatehost", zap.Error(err))
	}
	c.BindJSON(&host)
	//处理数据
	if err := logic.UpdateHost(host); err != nil {
		zap.L().Error("logic UpdateHostHandler ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, nil)
}

// DeleteHostHandler  删除主机
func DeleteHostHandler(c *gin.Context) {
	//获取参数
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	fmt.Println("开始执行删除操作")
	fmt.Println("delete参数id", id)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
	}
	if err := logic.DeleteHost(id); err != nil {
		zap.L().Error("controller DeleteHostHandler", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, nil)
}
