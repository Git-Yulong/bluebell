/**
  @author: ZYL
  @date:
  @note
*/
package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"projects/bluebell/logic"
	"projects/bluebell/models"
	"strconv"
)

//
//  CreatePostHandler
//  @Description: 创建帖子Handler
//  @param context
//
func CreatePostHandler(context *gin.Context) {
	// 获取参数
	postPrams := new(models.Post)
	if err := context.ShouldBindJSON(postPrams); err != nil {
		zap.L().Error("获取前端参数错误，err=", zap.Error(err))
		ResponseErrorWithMsg(context, CodeInvalidParam, "参数有误！")
		return
	}
	err := logic.CreatePost(postPrams)
	if err != nil {
		zap.L().Error("创建帖子失败！err=", zap.Error(err))
		ResponseErrorWithMsg(context, CodeInvalidParam, "创建帖子失败！err="+err.Error())
		return
	}
	ResponseSuccess(context, "创建帖子成功!")
	// 调用service层
}

//
//  GetPostDetailHandler
//  @Description: 获取帖子详情Handler
//  @param context
//
func GetPostDetailHandler(context *gin.Context) {

	parma := context.Param("pid") // 帖子id
	pid, err := strconv.ParseInt(parma, 10, 64)
	if err != nil {
		zap.L().Error(fmt.Sprintf("strconv.ParseInt(%s) 错误， err=%s", parma, err.Error()))
		return
	}

	// 获取帖子详情， 获取作者信息，获取社区信息，获取评论信息
	data, err := logic.GetPostDetailByPid(pid)
	if err != nil {
		logMsg := fmt.Sprintf("controllers.GetPostDetailHandler CALL logic.GetPostDetailByPid(%d) err=%s", pid, err)
		zap.L().Error(logMsg)
		ResponseError(context, CodeServerBusy)
		return
	}

	ResponseSuccess(context, data)

}

//
//  GetPostListHandler
//  @Description: 分页获取帖子列表Handler
//  @param ctx
//
func GetPostListHandler(ctx *gin.Context) {
	pageStr := ctx.Query("page")
	pageSizeStr := ctx.Query("page_size")
	limitStr := ctx.Query("limit")

	page, err := strconv.ParseInt(pageStr, 10, 0)
	if err != nil {
		logMsg := fmt.Sprintf("GetPostListHandler CALL strconv.ParseInt(%d, %d, %d) failed, err = %s", page, 10, 64, err)
		zap.L().Error(logMsg)
		return
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 0)
	if err != nil {
		logMsg := fmt.Sprintf("GetPostListHandler CALL strconv.ParseInt(%d, %d, %d) failed, err = ", pageSize, 10, 64)
		zap.L().Error(logMsg)
		return
	}
	limit, err := strconv.ParseInt(limitStr, 10, 0)
	if err != nil {
		logMsg := fmt.Sprintf("GetPostListHandler CALL strconv.ParseInt(%d, %d, %d) failed, err = ", limit, 10, 64)
		zap.L().Error(logMsg)
		return
	}

	var (
		offset int64
	)
	if page == 0 {
		offset = 0
	} else {
		offset = (page - 1) * pageSize
	}

	data, err := logic.GetPostListByPage(int(offset), int(pageSize))
	if err != nil {
		logMsg := fmt.Sprintf("GetPostListHandler CALL logic.GetPostListByPage(%d, %d) failed, err = ", offset, pageSize)
		zap.L().Error(logMsg)
		return
	}
	ResponseSuccess(ctx, data)

}
