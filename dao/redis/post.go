/**
  @author: ZYL
  @date:
  @note
*/
package redis

// CreatePost 创建帖子
func CreatePost(postId int64, title string, content string, authorId int, communityId int64) {
	// 生成文章id
	rdb.Do("incr")

}
