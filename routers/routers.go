// Package routers
// @Description: 设置路由
package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"projects/bluebell/controllers"
	"projects/bluebell/logger"
	"projects/bluebell/middleware"
)

func Setup() *gin.Engine {
	r := gin.New()
	//r.StaticFS("./upload", http.Dir("./upload"))
	mode := viper.GetString("app.mode")
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(logger.GinLogger(), logger.GinRecovery(true)) // 整合zap日志

	// /api/v1
	v1 := r.Group("/api/v1")
	v1.POST("/signup", controllers.SignUpHandler)
	v1.POST("/login", controllers.LoginHandler)
	v1.GET("/ping", middleware.JWTAuthMiddleware(), controllers.PingHandler)
	v1.GET("/telcode/:tel", controllers.TelCheckCodeHandler)
	v1.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})

	v1.Use(middleware.JWTAuthMiddleware()) // 使用认证中间件
	{
		v1.GET("/community", controllers.GetCommunityListHandler)
		v1.GET("/community/:id", controllers.GetCommunityDetailListHandler)
		v1.POST("/post", controllers.CreatePostHandler)
		v1.POST("/upload", controllers.UploadHeadPicHandler)
		v1.GET("/first_level_comments/:pid", controllers.GetFirstLevelCommentsByPostId)
		v1.GET("/post_detail/:pid", controllers.GetPostDetailHandler)
		v1.GET("/post/", controllers.GetPostListHandler)
	}

	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
