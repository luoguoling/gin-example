package logic

import (
	"fmt"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	//判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 1.生成uid
	userId := snowflake.GenID()
	//用户密码加密
	//构造user实例
	user := &models.User{
		UserID:   userId,
		Username: p.Username,
		Password: p.Password,
	}

	//保存到数据库

	err = mysql.InsertUser(user)
	return err
}

//登录
func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		fmt.Println("用户检查失败")
		return "", err
	}
	//生成jwt的token
	fmt.Println("key is ....")
	fmt.Println(jwt.GenerateToken(user.UserID, user.Username))
	return jwt.GenerateToken(user.UserID, user.Username)

}
