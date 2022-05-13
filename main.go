package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"projects/bluebell/controllers"
	"projects/bluebell/dao/mysql"
	"projects/bluebell/dao/redis"
	"projects/bluebell/logger"
	"projects/bluebell/pkg/dwf"
	"projects/bluebell/pkg/snowflake"
	"projects/bluebell/routers"
	"projects/bluebell/settings"
	"syscall"
	"time"
)

// 通用 go web 开发脚手架

// 1 加载配置
// 2 初始化日志
// 3 初始化 MYSQL连接
// 4 初始化Redis
// 5 注册路由
// 6 启动服务
func main() {
	// 1 加载配置viper
	if err := settings.Init(); err != nil {
		log.Printf("settings init failed, err: %v\n", err)
		return
	}

	// 2 初始化日志
	if err := logger.Init(); err != nil {
		log.Printf("logger init failed, err: %v\n", err)
		return
	}
	defer zap.L().Sync() // 同步到缓冲区

	zap.L().Debug("log init success...")
	// 3 初始化 MYSQL连接
	if err := mysql.Init(); err != nil {
		zap.L().Error("mysql init failed, err:", zap.Error(err))
		return
	}
	defer mysql.Close()

	// 4 初始化 Redis 连接
	if err := redis.Init(); err != nil {
		zap.L().Error("redis init failed, err:", zap.Error(err))
		return
	}
	defer redis.Close()

	if err := snowflake.Init(viper.GetString("app.start_time"), viper.GetInt64("app.machine_id")); err != nil {

		//log.Fatal(err)
		zap.L().Error("snowflake init failed, err:", zap.Error(err))
		return
	}

	if err := dwf.Init(); err != nil {
		zap.L().Error("dirtyWordFilter init failed, err:", zap.Error(err))
		return
	}

	// 初始化校验器 validator
	if err := controllers.InitTrans("zh"); err != nil {
		zap.L().Error("validator init failed ", zap.Error(err))
		return
	}

	// 5 注册路由
	r := routers.Setup()
	r.Run(":" + viper.GetString("app.port"))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
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
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
	return
}
