/**
  @author: ZYL
  @date:
  @note
*/
package models

import "time"

// 内存对齐
type Post struct {
	PostId      int64     `json:"post_id,string" db:"post_id"`                       // 帖子id
	AuthorID    int64     `json:"author_id" db:"author_id" binding:"required"`       // 作者id
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"` // 社区id
	Status      int32     `json:"status" db:"status"`                                // 帖子状态
	Title       string    `json:"title" db:"title" binding:"required"`               // 帖子标题
	Content     string    `json:"content" db:"content" binding:"required"`           // 帖子内容
	CreateTime  time.Time `json:"create_time" db:"create_time"`                      // 帖子创建时间
	UpdateTime  time.Time `json:"update_time" db:"update_time"`                      // 帖子创建时间
	Share       bool      `json:"share" db:"share"`
	PicPath     string    `json:"pic_path" db:"pic_path"`
	VideoPath   string    `json:"video_path" db:"vedio_path"`
	LikeCnt     int64     `json:"like_cnt" db:"like_cnt"`
}

type ApiPostDetail struct {
	Post      *Post           `json:"post"`
	User      *User           `json:"author"`
	Comments  []CommentDetail `json:"comments"`
	Community *Community      `json:"community"`
}

type ApiPostDetail2 struct {
	PostId        int64  `json:"post_id,string" db:"post_id"`                       // 帖子id
	AuthorID      int64  `json:"author_id" db:"author_id" binding:"required"`       // 作者id
	CommunityID   int64  `json:"community_id" db:"community_id" binding:"required"` // 社区id
	LikeCnt       int64  `json:"like_cnt" db:"like_cnt"`
	Title         string `json:"title" db:"title" binding:"required"`     // 帖子标题
	Content       string `json:"content" db:"content" binding:"required"` // 帖子内容
	HeadPic       string `json:"head_pic" db:"head_pic"`
	UserName      string `json:"username" db:"username"`
	CommunityName string `json:"community_name" db:"community_name"`
}
