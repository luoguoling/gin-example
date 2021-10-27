package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
	"web_app/settings"
)

var gormdb *gorm.DB

func InitGorm(cfg *settings.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname,
	)
	gormdb, err = gorm.Open("mysql", dsn)
	if err != nil {
		zap.L().Error("mysql gorm数据库连接失败!!!")
	}
	return err

}
func GormClose() {
	_ = gormdb.Close()

}
