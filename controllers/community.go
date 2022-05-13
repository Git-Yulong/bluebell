/**
  @author: ZYL
  @date:
  @note: 社区相关的handler
*/
package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"projects/bluebell/logic"
	"strconv"
)

//
//  GetCommunityListHandler
//  @Description: 获取社区列表的Handler
//  @param context *gin.Context
//
func GetCommunityListHandler(context *gin.Context) {
	communityList, err := logic.GetCommunityList()
	if err != nil {
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, communityList)
}

//
//  GetCommunityDetailListHandler
//  @Description: 获取某个社区详情
//  @param context
//
func GetCommunityDetailListHandler(context *gin.Context) {

	idStr := context.Param("id")               // 获取社区id
	id, err := strconv.ParseInt(idStr, 10, 64) //  转换成整数
	if id <= 0 {
		zap.L().Error("controller.GetCommunityDetailListHandler(context *gin.Context), err =  ", zap.Error(err))
		ResponseError(context, CodeInvalidParam)
		return
	}
	community, err := logic.GetCommunityById(id) //  调用logic获取社区详情
	if err != nil {
		zap.L().Error("controller.GetCommunityDetailListHandler CALL logic.GetCommunityById(id), err=", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, community)
}
