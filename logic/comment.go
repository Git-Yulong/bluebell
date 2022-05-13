/**
  @author: ZYL
  @date:
  @note
*/
package logic

import (
	"go.uber.org/zap"
	"projects/bluebell/dao/mysql"
	"projects/bluebell/dao/redis"
	"projects/bluebell/models"
)

//
//  GetFirstLevelCommentsByPid
//  @Description: 根据文章id获取一级评论
//  @param pid
//  @return []models.CommentDetail
//  @return error
//
func GetFirstLevelCommentsByPid(pid int64) ([]models.CommentDetail, error) {
	// 从mysql中根据帖子id查询对应的评论信息
	comments, err := mysql.GetFirstLevelCommentsByPid(pid)
	if err != nil {
		zap.L().Error("log.GetFirstLevelCommentsByPid(pid) CALL mysql.GetFirstLevelCommentsByPid(pid), ERR=", zap.Error(err))
		return nil, err
	}

	for i := range comments {
		// 从redis中获取点赞数  存到comments[i]里面
		count, err := redis.GetFirstLevelCommentLikeCount(comments[i].PostId, comments[i].CommentId)
		if err != nil {
			zap.L().Error("log.GetFirstLevelCommentsByPid(pid) CALL redis.GetFirstLevelCommentLikeCount, ERR=", zap.Error(err))
			continue
		}
		comments[i].LikeCnt = count
	}
	return comments, nil
}

//
//  GetSecondLevelCommentsByPostIdAndCommentId
//  @Description: 根据帖子id和一级评论id获取二级评论
//  @param postId
//  @param commentId
//  @return []models.SecondLevelCommentDetail
//  @return error
//
func GetSecondLevelCommentsByPostIdAndCommentId(postId, commentId int64) ([]models.SecondLevelCommentDetail, error) {
	// 从mysql中根据帖子id查询对应的评论信息
	return nil, nil
}
