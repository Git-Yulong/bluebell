/**
  @author: ZYL
  @date:
  @note: 评论相关的Handler
*/
package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"projects/bluebell/logic"
	"strconv"
)

//
//  GetFirstLevelCommentsByPostId
//  @Description: 根据帖子id，获取一级评论记录
//  @param context
//
func GetFirstLevelCommentsByPostId(context *gin.Context) {
	// 获取帖子pid
	parma := context.Param("pid")
	// 前端传过来是string类型，解析成int64位的
	pid, err := strconv.ParseInt(parma, 10, 64)

	if err != nil {
		zap.Error(err)
		ResponseError(context, CodeInvalidParam) // 请求参数有误
		return
	}

	// 到这里说明已经成功获取post_id
	// 调用logic层获取评论信息
	comments, err := logic.GetFirstLevelCommentsByPid(pid)
	if err != nil {
		zap.L().Error("controller.GetFirstLevelCommentsByPostId CALL logic.GetFirstLevelCommentsByPid(pid) err = ", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, comments)
}

//
//  GetSecondLevelCommentByPostId
//  @Description: 根据帖子id，获取二级评论记录
//  @param context
//
func GetSecondLevelCommentByPostId(context *gin.Context) {

}
