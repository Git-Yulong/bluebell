package models

type TbComment struct {
	Id         int64  `json:"id"`
	CommentId  int64  `json:"comment_id"`
	PostId     int64  `json:"post_id"`
	UserId     int64  `json:"user_id"`
	Content    string `json:"content"`
	Level      int64  `json:"level"`
	ParentId   int64  `json:"parent_id"`
	CreateTime string `json:"create_time"`
	Share      int64  `json:"share"`
	LikeCnt    int64  `json:"like_cnt"`
}
