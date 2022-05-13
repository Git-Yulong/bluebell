/**
  @author: ZYL
  @date:
  @note
*/
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projects/bluebell/controllers"
	"projects/bluebell/pkg/jwt"
	"strings"
)

//
//  JWTAuthMiddleware
//  @Description: 登录认证JWT 中间件
//  @return func(*gin.Context)
//
func JWTAuthMiddleware() func(*gin.Context) {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")

		if authHeader == "" {
			// 没有认证信息
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "请求头中没有认证信息",
			})
			ctx.Abort()
			return
		}

		// 能执行到这里，说明有认证信息
		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"code": 203,
				"msg":  "请求头中auth格式有误",
			})
			ctx.Abort()
			return
		}

		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 205,
				"msg":  "无效的token",
			})

			ctx.Abort()
			return
		}

		ctx.Set(controllers.CtxUserID, mc.UserID)
		ctx.Next()
	}
}
