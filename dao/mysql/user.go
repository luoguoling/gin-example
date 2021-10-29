package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"web_app/models"
)

type User struct {
	Id         int8   `json:"id"`
	UserId     int64  `json:"user_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func TableName() {
	return

}

func CheckUserExist(username string) error {
	//对密码加密
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已经存在！！！")
	}
	return nil
}

const secret = "aaaa"

//func InsertUser(p *models.ParamSignUp) (err error) {
//	//数据入库
//	gormdb.AutoMigrate(&p)
//	err = gormdb.Create(&p).Error
//	return err
//}
// InsertUser 插入用户数据
func InsertUser(user *models.User) (err error) {
	//对密码加密
	user.Password = encryptPassword(user.Password)
	//执行sql语句
	sqlStr := "insert into user(user_id,username,password) values (?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}
func encryptPassword(opassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(opassword)))
}
