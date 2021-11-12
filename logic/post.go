package logic

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	//1.生成id
	p.ID = snowflake.GenID()
	//2.插入数据
	if err := mysql.CreatePost(p); err != nil {
		zap.L().Error("mysql.CreatePost(P) failed", zap.Error(err))

	}

	//返回
	return

}
