package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"web_app/models"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
		post_id,title,content,author_id,community_id)
		values (?,?,?,?,?)
		`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return

}

func GetPostDetail(id int64) (post *models.ApiPostDetail, err error) {
	post = new(models.ApiPostDetail)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id = ?
	`
	err = db.Get(post, sqlStr, id)
	if err == sql.ErrNoRows {
		zap.L().Error("mysql.GetPostDetail failed", zap.Error(err))
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query post failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	return
}

//获取所有帖子
func GetPosts(pageNum, pageSize int) (posts []*models.ApiPostDetail, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	limit ?,?
	`
	err = db.Select(&posts, sqlStr, pageNum, pageSize)
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
