package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
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

func CheckUserExist(username string) error {
	//对密码加密
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExit
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
	sqlStr := "select user_id, username, password from user where username = ?"
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库失败
		return err
	}

	//判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorPasswordWrong
	}
	return

}

//根据用户id返回用户信息
func GetUserByID(userid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select 
	user_id,username
	from user
	where user_id = ?
	`
	err = db.Get(user, sqlStr, userid)
	if err == sql.ErrNoRows {
		err = ErrorUserNotExist
	}
	if err != nil {
		return
	}
	return
}
