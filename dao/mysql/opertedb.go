package mysql

import (
	"fmt"
	"web_app/pkg/snowflake"
)

type UserInfo struct {
	ID     uint
	UserId int64
	Name   string
	Gender string
	Hobby  string
}

func Opertedb() {
	userid := snowflake.GenID()
	fmt.Println(userid)
	u1 := UserInfo{5, userid, "七米", "男", "篮球"}

	gormdb.AutoMigrate(&UserInfo{})
	gormdb.Create(&u1)
	var u = new(UserInfo)
	gormdb.First(u)
	fmt.Println(u)
	gormdb.Model(&u).Update("name", "rolin")
	//gormdb.Delete(&u)

}
