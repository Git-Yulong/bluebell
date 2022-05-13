/**
  @author: ZYL
  @date:
  @note
*/
package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const CtxUserID = "userid"

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentIUserId 获取当前登录用户的userid
// 如果登录了，则返回当前用户的userid, nil
// 否则返回0, err
func getCurrentIUserId(ctx *gin.Context) (int64, error) {
	uid, ok := ctx.Get(CtxUserID)
	if !ok {

		return 0, ErrorUserNotLogin
	}
	return uid.(int64), nil
}
