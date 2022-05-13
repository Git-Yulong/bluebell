/**
  @author: ZYL
  @date:
  @note
*/
package mysql

import (
	"go.uber.org/zap"
	"projects/bluebell/models"
)

// GetFirstLevelCommentsByPid 根据帖子id获取一级评论
func GetFirstLevelCommentsByPid(pid int64) ([]models.CommentDetail, error) {
	sqlStr := `SELECT 
				 comment_id,
				 post_id, 
				 tb_user.user_id,
				 tb_user.username,
				 content, 
				 level, 
			   tb_comment.create_time
			FROM tb_comment 
			INNER JOIN tb_user ON tb_user.user_id = tb_comment.user_id
			WHERE post_id = ? AND level = 1`

	var comments []models.CommentDetail
	if err := db.Select(&comments, sqlStr, pid); err != nil {
		zap.L().Error("logic.GetFirstLevelCommentsByPid CALL db.Select(&comments, sqlStr, pid), err=", zap.Error(err))
		return nil, err
	}
	return comments, nil
}

// GetSecondLevelCommentsByPostIdAndCommentId 根据PostId和一级commentId获取二级评论
func GetSecondLevelCommentsByPostIdAndCommentId(postId int64, commentId int64) ([]models.SecondLevelCommentDetail, error) {
	sqlStr := `
			SELECT 
			tc1.post_id, 
			tc1.user_id reply_user_id, 
			tu1.username reply_username, 
			tc1.content, 
			tc2.user_id be_replied_user_id, 
			tu2.username be_replied_username
			FROM tb_comment tc1 
			INNER JOIN tb_comment tc2 ON tc1.trg_comment_id = tc2.comment_id 
			INNER JOIN tb_user tu1 ON tc1.user_id = tu1.user_id 
			INNER JOIN tb_user tu2 ON tc2.user_id = tu2.user_id 
			WHERE tc1.post_id=? AND tc1.trg_comment_id = ?;
	`

	result := make([]models.SecondLevelCommentDetail, 1)
	if err := db.Select(&result, sqlStr, postId, commentId); err != nil {
		zap.L().Error("mysql.GetSecondLevelCommentsByPostIdAndCommentId err = ", zap.Error(err))
		return nil, err
	}
	return result, nil
}
