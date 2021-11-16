package logic

import (
	"fmt"
	"web_app/dao/mysql"
	"web_app/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	//mysql查询需要的数据
	fmt.Println("mysql.getCommunityList is")
	fmt.Println(mysql.GetCommunityList())
	return mysql.GetCommunityList()

}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityByID(id)

}
