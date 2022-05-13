/**
  @author: ZYL
  @date:
  @note
*/
package logic

import (
	"projects/bluebell/models"
)

func VotePost(params *models.ParmaVotePost) error {
	//1 判断帖子投票时间是否过期
	//2 判断该用户是否已经为该帖子投过票
	//3 计算得分，写入redis
	//4 将该用户id加入set中代表已经为该帖子投过票
	return nil
}
