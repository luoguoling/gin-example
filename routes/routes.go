package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"web_app/controller"
	_ "web_app/docs" // 千万不要忘了导入把你上一步生成的docs
	"web_app/logger"
	middleware "web_app/middleware/cors"
	"web_app/middleware/jwt"
	"web_app/middleware/ratelimit"
	"web_app/settings"

	"github.com/gin-contrib/pprof"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1")
	//注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	//登录路由
	v1.POST("/login", controller.LoginHandler)
	v1.GET("/communityList", controller.CommunityHandler)
	v1.GET("/community/:id", controller.CommunityDetailHandler)
	//v1.Use(jwt.JWTAuthMiddleware(), middleware.Cors(), ratelimit.RateLimitMiddleware(time.Second*2, 1))
	v1.Use(middleware.Cors(), ratelimit.RateLimitMiddleware(time.Second*2, 10000))
	{
		//post
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostsHandler)

		//投票功能
		v1.POST("/vote", controller.PostVoteHandler)
	}
	pprof.Register(r) //注册pprof
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	v1.GET("/hosts", controller.GetHostHandler)
	v1.GET("/host/:id", controller.GetHostDetailHandler)
	v1.POST("/host", controller.CreateHostHandler)
	v1.PUT("/host/:id", controller.UpdateHostHandler)
	v1.DELETE("/host/:id", controller.DeleteHostHandler)
	return r
}
