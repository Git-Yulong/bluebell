package models

type TbUser struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"user_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Telephone  string `json:"telephone"`
	Gender     int64  `json:"gender"`
	HeadPic    string `json:"head_pic"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
