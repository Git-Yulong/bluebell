/**
  @author: ZYL
  @date:
  @note
*/
package mysql

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"projects/bluebell/models"
	"strconv"
)

func CreatePost(post *models.Post) error {
	sqlStr := `insert into tb_post(post_id, title, content, author_id, community_id, status, create_time, update_time, like_cnt) values(?,?,?,?,?,?,?,?,?)`
	_, err := db.Exec(sqlStr, post.PostId, post.Title, post.Content, post.AuthorID, post.CommunityID, post.Status, post.CreateTime, post.CreateTime, post.LikeCnt)

	if err != nil {
		zap.L().Error("保存post到数据库中失败, err=", zap.Error(err))
		return err
	}

	return nil
}

func GetPostById(pid int64) (*models.Post, error) {
	var post models.Post
	sqlStr := `select post_id, title, content, author_id, community_id, status, create_time, update_time, share, like_cnt 
			from tb_post where post_id = ?`
	err := db.Get(&post, sqlStr, pid)
	if err != nil {
		errMsg := fmt.Sprintf("mysql.GetPostById(%d), err= %s", pid, err)
		zap.L().Error(errMsg)
		return nil, err
	}
	return &post, nil

}

// SelectPostListByPage 分页获取帖子
func SelectPostListByPage(pageNum, pageSize int) ([]models.Post, error) {
	sqlStr := `select post_id, title, content, author_id, community_id, status, create_time, update_time, share, like_cnt 
			from tb_post where status = 1 order by create_time limit ?, ?`
	var posts []models.Post
	err := db.Select(&posts, sqlStr, pageNum, pageSize)
	if err != nil {
		logMsg := fmt.Sprintf("mysql.SelectPostListByPage CALL db.Select(&posts, %s, %d, %d)", sqlStr, pageNum, pageSize)
		zap.L().Error(logMsg)
		return nil, err
	}

	return posts, nil
}

// GetPostDetail 获取post详情
func GetPostDetail(pids []int64) ([]models.ApiPostDetail2, error) {
	if len(pids) <= 0 {
		return nil, errors.New("pids为空")
	}

	sqlStr := `
		SELECT 
		post_id, 
		title, 
		content, 
		like_cnt, 
		author_id,
		tb_user.username,
		tb_user.head_pic,
		tb_community.community_id, 
		community_name
		from tb_post 
		INNER JOIN tb_user 
			ON tb_user.user_id = author_id
		INNER JOIN tb_community 
			ON tb_community.community_id = tb_post.community_id
		WHERE tb_post.post_id IN 
	`
	sqlStr += `(`
	for i := 0; i < len(pids)-1; i++ {
		sqlStr += strconv.FormatInt(pids[i], 10) + ","
	}
	sqlStr += strconv.FormatInt(pids[len(pids)-1], 10)
	sqlStr += `);`
	var apiPostDetail2 []models.ApiPostDetail2
	err := db.Select(&apiPostDetail2, sqlStr)
	if err != nil {
		logMsg := fmt.Sprintf("mysql.GetPostDetail(post *models.Post) CALL db.Get(postDetail,sqlStr) failed, err=%s", err)
		zap.L().Error(logMsg)
		return nil, err
	}
	return apiPostDetail2, nil
}
