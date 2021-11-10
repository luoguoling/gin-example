package mysql

import (
	"database/sql"
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
	community = new(models.CommunityDetail)
	sqlStr := "select community_id,community_name,introduction,create_time from  community where community_id = ?"
	if err := db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return
}
