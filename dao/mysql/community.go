package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"web_app/models"
)

//获取列表数据
func GetCommunity() (data []models.Community, err error) {
	sqlStr := "select community_id,community_name from  community"
	rows, err := db.Query(sqlStr)
	if err != nil {
		zap.L().Error("GetCommunity 执行query报错", zap.Error(err))
		fmt.Println("查询数据库报错!!!!")
	}
	var community models.Community
	for rows.Next() {
		err := rows.Scan(&community)
		if err != nil {
			fmt.Println("赋值报错", err)
		}

	}
	if err != nil {
		zap.L().Error("GetCommunity查询数据报错", zap.Error(err))
		return
	}

	return
}
