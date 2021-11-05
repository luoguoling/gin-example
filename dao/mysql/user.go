package mysql

import (
	"crypto/md5"
	"database/sql"
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

var (
	ErrorUserExist       = errors.New(("用户已经存在"))
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户密码错误")
)

func CheckUserExist(username string) error {
	//对密码加密
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
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

//加密方法
func encryptPassword(opassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(opassword)))
}

//检查登录用户是否正确
func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := "select username,password from user where username = ?"
	err = db.Get(user, sqlStr, user.Username)
	if err != nil && err != sql.ErrNoRows {
		// 查询数据库出错
		return
	}
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}

	//判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return

}
