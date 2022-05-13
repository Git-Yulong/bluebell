/**
  @author: ZYL
  @date:
  @note
*/
package models

import "time"

type CommentDetail struct {
	CommentId  int64     `json:"comment_id" db:"comment_id"` // 评论id
	PostId     int64     `json:"post_id" db:"post_id"`       // 帖子id
	UserId     int64     `json:"user_id" db:"user_id"`       // 帖子评论用户的id
	UserName   string    `json:"username" db:"username"`     // 评论的用户名
	LikeCnt    int64     `json:"like_cnt" db:"like_cnt"`
	Content    string    `json:"content" db:"content"` // 评论内容
	Level      int8      `json:"level" db:"level"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}

type SecondLevelCommentDetail struct {
	PostId     int64     `json:"post_id" db:"post_id"`         // 帖子id
	Content    string    `json:"content" db:"content"`         // 评论内容
	UserId1    int64     `json:"user_id1" db:"user_id1"`       // 评论的用户id
	UserName1  string    `json:"username1" db:"username1"`     // 评论的用户名
	UserId2    int64     `json:"user_id2" db:"user_id2"`       // 被回复的用户id
	UserName2  string    `json:"username2" db:"username2"`     // 被回复的用户名
	Level      int8      `json:"level" db:"level"`             // 评论级别 分为1级评论和2级评论
	LikeCnt    int64     `json:"like_cnt" db:"like_cnt"`       // 点赞数量
	CreateTime time.Time `json:"create_time" db:"create_time"` // 评论时间
}
