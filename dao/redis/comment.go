// Package redis
// @Description: 评论相关函数
package redis

import (
	"strconv"
)

//
//  GetFirstLevelCommentLikeCount
//  @Description: 获取一级评论的点赞数
//  @param postId 帖子id
//  @param commentId 评论id
//  @return int64 点赞数
//  @return error
//
func GetFirstLevelCommentLikeCount(postId, commentId int64) (int64, error) {
	postIdStr := strconv.FormatInt(postId, 10)
	commentIdStr := strconv.FormatInt(commentId, 10)
	key := KeyCommentFirstLevelZSet + ":" + postIdStr
	result, err := rdb.ZScore(key, commentIdStr).Result()
	if err != nil {
		return 0., err
	}
	return int64(result), nil
}
