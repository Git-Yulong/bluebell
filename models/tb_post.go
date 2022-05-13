package models

type TbPost struct {
	Id          int64  `json:"id"`
	PostId      int64  `json:"post_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	AuthorId    int64  `json:"author_id"`
	CommunityId int64  `json:"community_id"`
	Status      int64  `json:"status"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
	Share       int64  `json:"share"`
	PicPath     string `json:"pic_path"`
	VedioPath   string `json:"vedio_path"`
	LikeCnt     int64  `json:"like_cnt"`
}
