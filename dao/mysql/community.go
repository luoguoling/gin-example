package mysql

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"web_app/models"
)

//获取列表数据
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id,community_name from  community"
	err = db.Select(&communityList, sqlStr)
	if err != sql.ErrNoRows {
		zap.L().Warn("communityList is null")
		err = nil
	}

	return
}

//获取详情
func GetCommunityByID(id int64) (community *models.CommunityDetail, err error) {
	fmt.Println("数据库接收到的id", id)
	community = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time
	from community
	where community_id = ?`
	err = db.Get(community, sqlStr, id)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
		return
	}

	fmt.Println("mysql", community)
	return community, err
}
