package logic

import (
	"context"
	"fmt"
	"web_app/dao/mysql"
	"web_app/dao/redis"
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
func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		fmt.Println("logic.user 用户检查失败")
		return nil, err
	}
	//生成jwt的token
	fmt.Println("key is ....")
	var ctx = context.Background()
	token, _, _ := jwt.GenerateToken(user.UserID, user.Username)
	err = redis.Rdb.Set(ctx, p.Username, token, 0).Err()
	user.Token = token
	if err != nil {
		fmt.Println("插入数据失败!!!")

	}
	return

}
