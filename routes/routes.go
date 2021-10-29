package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"web_app/controller"
	"web_app/logger"
	"web_app/settings"
)

func Setup() *gin.Engine {

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	fmt.Println("打印viper值")
	r.GET("/version", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": fmt.Sprintf("%s", settings.Conf.Version),
			"date":    time.Now(),
			"m1":      fmt.Sprintf("%s", settings.Conf.Name),
		})

	})
	//注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	return r
}
