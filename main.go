package main

import (
	"context"
	"flag"
	"fmt"
	"web_app/controller"
	"web_app/pkg/snowflake"
	//"web_app/pkg/util"

	//"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/routes"
	"web_app/settings"
)

//go web脚手架模板1
func main() {
	//if len(os.Args) < 2 {
	//	fmt.Println("请输入配置文件config")
	//	return
	//}
	var filepath string
	flag.StringVar(&filepath, "c", "filepath", "配置文件")
	flag.Parse()
	//1.加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Printf("配置文件初始化失败,err:%v\n", err)
		return
	} else {
		fmt.Println("配置文件加载成功!!!!")
	}
	//2.初始化配置
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("日志始化失败,err:%v\n", err)
		return
	}
	zap.L().Sync() //缓存区日志追加到日志中
	//3.初始化mysql连接
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		fmt.Printf("mysql数据库始化失败,err:%v\n", err)
		return
	}
	defer mysql.Close()

	//3.1 初始化gorm连接
	if err := mysql.InitGorm(settings.Conf.MysqlConfig); err != nil {
		fmt.Printf("gorm数据库始化失败,err:%v\n", err)
		return
	}
	defer mysql.GormClose()

	//4.初始化redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("redis始化失败,err:%v\n", err)
		return
	} else {
		fmt.Println("redis连接成功!!!")
	}
	defer redis.Close()
	redis.Operatedb()
	fmt.Println(settings.Conf.MachineID)
	//雪花id生成
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Println("打印的值是:")
		fmt.Println(settings.Conf.StartTime, settings.Conf.MachineID)
		fmt.Printf("init snowflake failed,err:%v\n", err)
		return
	}
	mysql.Opertedb()
	fmt.Println(snowflake.GenID())
	//初始化gin校验使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Println("翻译初始化错误!!!")
		return
	}
	//getlocation
	// util.GetLocation()
	//5.注册路由
	r := routes.Setup(settings.Conf.Mode)
	//6.启动服务(优雅关机)
	srv := &http.Server{
		//Addr:    fmt.Sprintf("%d", viper.GetInt("app.port")),
		Addr:    fmt.Sprintf(":%s", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Info("shutdown Server.....", zap.Error(err))
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
	zap.L().Info("Server exiting")
}
