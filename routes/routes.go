package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"web_app/controller"
	"web_app/logger"
	"web_app/middleware/jwt"
	"web_app/settings"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	fmt.Println("打印viper值")
	r.GET("/version", jwt.JWTAuthMiddleware(), func(context *gin.Context) {
		//如果是登录用户，判断是否含有有效token
		//islogin := true

		context.JSON(200, gin.H{
			"message": fmt.Sprintf("%s", settings.Conf.Version),
			"date":    time.Now(),
			"m1":      fmt.Sprintf("%s", settings.Conf.Name),
		})

	})
	//注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	//登录路由
	r.POST("/login", controller.LoginHandler)
	return r
}
