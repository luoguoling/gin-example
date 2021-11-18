package logic

import (
	"fmt"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	//1.生成id
	p.ID = snowflake.GenID()
	//2.插入数据
	if err1 := mysql.CreatePost(p); err1 != nil {
		zap.L().Error("mysql.CreatePost(P) failed", zap.Error(err1))
		return

	}
	//返回
	return

}

func GetPostDetail(id int64) (data *models.ApiPostDetail, err error) {
	data, err = mysql.GetPostDetail(id)
	if err != nil {
		zap.L().Error("mysql.GetPostDetail(id) failed", zap.Error(err))
		return nil, err
	}
	//通过用户id获取用户
	user, err := mysql.GetUserByID(data.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID(data.AuthorID) failed", zap.Error(err))
	}
	data.AuthorName = user.Username
	community, err := mysql.GetCommunityByID(data.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed", zap.String("community_id", fmt.Sprint(data.CommunityID)), zap.Error(err))
		return
	}
	data.CommunityName = community.Name
	fmt.Println("logic.data is ", data)
	return data, nil
}

//获取所有帖子的数据
func GetPosts(pageNum, pageSize int) (posts []*models.ApiPostDetail, err error) {
	posts, err = mysql.GetPosts(pageNum, pageSize)
	if err != nil {
		zap.L().Error("mysql.GetPosts failed", zap.Error(err))
	}
	postSlice := make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		postnew := post
		//通过用户id获取用户
		user, err := mysql.GetUserByID(postnew.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(data.AuthorID) failed", zap.Error(err))
		}
		postnew.AuthorName = user.Username
		community, err := mysql.GetCommunityByID(postnew.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityByID() failed", zap.String("community_id", fmt.Sprint(postnew.CommunityID)), zap.Error(err))

		}
		postnew.CommunityName = community.Name
		postSlice = append(postSlice, postnew)

	}
	return postSlice, nil

}
