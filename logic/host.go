package logic

import (
	"fmt"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
)

func GetHost(pageNum, pageSize int) (hosts []*models.Host, err error) {
	fmt.Println(pageNum, pageSize)
	hosts, err = mysql.GetHost(pageNum, pageSize)
	fmt.Println(err)
	if err != nil {
		zap.L().Error("logic gethost报错!!!", zap.Error(err))
	}
	return

}
func GetHostDetail(id int64) (host *models.Host, err error) {
	//mysql查询出数据
	host, err = mysql.GetHostDetail(id)
	if err != nil {
		zap.L().Error("logic gethostdetail报错!!!", zap.Error(err))
	}
	return host, nil
}

// CreateHost 添加host
func CreateHost(host *models.Host) (err error) {
	//数据入库操作
	if err1 := mysql.CreateHost(host); err1 != nil {
		zap.L().Error("logic createhost error", zap.Error(err1))
		return
	}
	return
}

// UpdateHost 更新Host
func UpdateHost(host *models.Host) (err error) {
	if err1 := mysql.UpdateHost(host); err1 != nil {
		zap.L().Error("logic updatehost", zap.Error(err1))
		return
	}
	return

}

// DeleteHost  删除主机
func DeleteHost(id int64) (err error) {
	if err1 := mysql.DeleteHost(id); err1 != nil {
		zap.L().Error("logic DeleteHost", zap.Error(err1))
		return
	}
	return
}
