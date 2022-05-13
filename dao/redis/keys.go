// Package redis
// @Description: 保存redis中的key
package redis

var (
	KeyPostTimeZSet          = "bluebell:post:time"     // ZSet 根据帖子创建时间排序 bluebell:post:time -> {member1, score1| member2, score2...}
	KeyPostScoreZSet         = "bluebell:post:score"    // ZSet 根据投票分数排序 bluebell:post:score -> {member1, score1| member2, score2...}
	KeyPostUserSet           = "bluebell:post:voted"    // Set， 前缀， 每一篇文章对应一个set，里面对应着为这这篇文章投过票的用户
	KeyCommentFirstLevelZSet = "bluebell:comment:first" // 加上postid 表示个帖子一级评论点赞数的zset
)
