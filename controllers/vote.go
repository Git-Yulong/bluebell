/**
  @author: ZYL
  @date:
  @note
*/
package controllers

import (
	"github.com/gin-gonic/gin"
	"projects/bluebell/logic"
	"projects/bluebell/models"
)

// VotePostHandler POST请求， 给帖子投票
func VotePostHandler(context *gin.Context) {
	var params models.ParmaVotePost
	if err := context.ShouldBindJSON(&params); err != nil {
		ResponseError(context, CodeInvalidParam)
		return
	}

	if err := logic.VotePost(&params); err != nil {
		ResponseError(context, CodeServerBusy)
		return
	}

}
