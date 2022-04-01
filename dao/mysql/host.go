package mysql

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"web_app/models"
)

func GetHost(pageNum, pageSize int) (hosts []*models.Host, err error) {
	sqlStr := `select region, hostname, publicip, privateip, os, status,remark,create_time
	from host
	limit ?,?
	`
	err = db.Select(&hosts, sqlStr, pageNum, pageSize)
	if err == sql.ErrNoRows {
		zap.L().Error("mysql.GetPosts failed", zap.Error(err))
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query posts failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	return

}
func GetHostDetail(id int64) (host *models.Host, err error) {
	host = new(models.Host)
	sqlStr := `select region, hostname, publicip, privateip, os, status,remark,create_time
	from host
	where id = ?
	`
	err = db.Get(host, sqlStr, id)
	fmt.Println("打印host详细信息:")
	fmt.Println(host)
	if err == sql.ErrNoRows {
		zap.L().Error("mysql.GetHostDetail failed", zap.Error(err))
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query host failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
		return
	}

	return host, nil
}

// CreateHost 数据库入库操作
func CreateHost(host *models.Host) (err error) {
	sqlStr := `insert into host(
		region,hostname,publicip,privateip,os,status,remark)
		values (?,?,?,?,?,?,?)
		`
	fmt.Println("打印的值", host.Remark, host.Privateip)
	_, err = db.Exec(sqlStr, host.Region, host.Hostname, host.Publicip, host.Privateip, host.Os, host.Status, host.Remark)
	return

}

//UpdateHost  修改数据
func UpdateHost(host *models.Host) (err error) {
	err = gormdb.Save(host).Error
	return
}

// DeleteHost 删除主机
func DeleteHost(id int64) (err error) {
	sqlStr := `delete  from host where id = ?`
	_, err = db.Exec(sqlStr, id)
	return
}
