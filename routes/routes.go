package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"web_app/controller"
	"web_app/logger"
	middleware "web_app/middleware/cors"
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
	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/version", jwt.JWTAuthMiddleware(), func(context *gin.Context) {
		//如果是登录用户，判断是否含有有效token
		//islogin := true

		context.JSON(200, gin.H{
			"message": fmt.Sprintf("%s", settings.Conf.Version),
			"date":    time.Now(),
			"m1":      fmt.Sprintf("%s", settings.Conf.Name),
		})

	})
	v1 := r.Group("/api/v1")
	//注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	//登录路由
	v1.POST("/login", controller.LoginHandler)

	v1.Use(jwt.JWTAuthMiddleware(), middleware.Cors())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		//post
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostsHandler)

		//投票功能
		v1.POST("/vote", controller.PostVoteHandler)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
