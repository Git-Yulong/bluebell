package models

type TbCommunity struct {
	Id            int64  `json:"id"`
	CommunityId   int64  `json:"community_id"`
	CommunityName string `json:"community_name"`
	Introduction  string `json:"introduction"`
	CreateTime    string `json:"create_time"`
	UpdateTime    string `json:"update_time"`
}
