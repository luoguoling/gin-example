package logic

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
)

func GetCommunity() (data []models.Community, err error) {
	//mysql查询需要的数据
	data, err = mysql.GetCommunity()
	if err != nil {
		zap.L().Error("GetCommunity执行查询语句出错", zap.Error(err))
		return

	}
	return
}
