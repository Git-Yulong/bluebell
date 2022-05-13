/**
  @author: ZYL
  @date:
  @note
*/
package logic

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"projects/bluebell/dao/mysql"
	"projects/bluebell/models"
	"projects/bluebell/pkg/dwf"
	"projects/bluebell/pkg/snowflake"
	"time"
)

//
//  CreatePost
//  @Description: 创建帖子
//  @param post
//  @return error
//
func CreatePost(post *models.Post) error {
	// 检查参数是否合法
	if len(post.Title) <= 1 {
		return errors.New("标题不能太短!")
	}

	// 这一步最好检查社区id是否存在
	err, flg := mysql.CheckCommunityExist(post.CommunityID)

	if err != nil || flg == false {
		return errors.New("所选社区不存在错误!")
	}

	err, newTitle := dwf.FilterDirtyWord(post.Title, "*")
	if err != nil {
		zap.L().Error("过滤标题敏感词发生错误, err", zap.Error(err))
		return err
	}
	post.Title = newTitle

	err, newContent := dwf.FilterDirtyWord(post.Content, "*")
	if err != nil {
		zap.L().Error("过滤内容敏感词发生错误, err", zap.Error(err))
		return err
	}
	post.Content = newContent

	fmt.Println(post)

	// 根据雪花算法生成帖子id
	post.PostId = snowflake.GenID()
	post.Status = 0 // 表示待审核
	post.CreateTime = time.Now().Local()
	post.LikeCnt = 0
	err = mysql.CreatePost(post)
	if err != nil {
		return err
	}
	return nil
}

//
//  GetPostDetailByPid
//  @Description: 获取帖子详情
//  @param pid 帖子id
//  @return *models.ApiPostDetail
//  @return error
//
func GetPostDetailByPid(pid int64) (*models.ApiPostDetail, error) {
	// 获取帖子信息
	post, err := mysql.GetPostById(pid)
	if err != nil {
		errMsg := fmt.Sprintf("logic.GetPostDetailByPostId(%d) CALL mysql.GetPostById(%d) , err=%s", pid, pid, err)
		zap.L().Error(errMsg)
		return nil, err
	}

	// 获取社区信息
	cid := post.CommunityID
	community, err := mysql.GetCommunityDetailByCid(cid)

	// 获取作者信息
	authorId := post.AuthorID
	user, err := mysql.GetUserByUid(authorId)
	// 获取一级评论信息
	comments, err := mysql.GetFirstLevelCommentsByPid(post.PostId)
	var res = models.ApiPostDetail{
		Post:      post,
		User:      user,
		Comments:  comments,
		Community: community,
	}
	return &res, nil
}

//
//  GetPostListByPage
//  @Description: 分页形式获取帖子列表
//  @param pageNum
//  @param pageSize
//  @return []models.ApiPostDetail2
//  @return error
//
func GetPostListByPage(pageNum, pageSize int) ([]models.ApiPostDetail2, error) {
	data, err := mysql.SelectPostListByPage(pageNum, pageSize)
	if err != nil {
		logMsg := fmt.Sprintf("mysql.GetPostListByPage(%d, %d) failed, err = %s.", pageNum, pageSize, err)
		zap.L().Error(logMsg)
		return nil, err
	}
	if len(data) == 0 {
		// 查询结果为空
		err = errors.New("len(data) == 0")
		logMsg := fmt.Sprintf("mysql.GetPostListByPage(%d, %d) failed, err = %s.", pageNum, pageSize, err)
		zap.L().Error(logMsg)
		return nil, err
	}

	pids := make([]int64, 0, len(data))
	for i := 0; i < len(data); i++ {
		pids = append(pids, data[i].PostId)
	}

	for i := 0; i < len(pids); i++ {
		fmt.Println("pids=", pids)
	}

	detail, err := mysql.GetPostDetail(pids)
	if err != nil {
		logMsg := fmt.Sprintf("mysql.mysql.GetPostDetail(pids) failed, err = %s.", err)
		zap.L().Error(logMsg)
		return nil, err
	}
	return detail, nil
}
