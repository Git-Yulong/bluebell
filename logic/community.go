/**
  @author: ZYL
  @date:
  @note
*/
package logic

import (
	"projects/bluebell/dao/mysql"
	"projects/bluebell/models"
)

//
//  GetCommunityList
//  @Description: 获取所有的社区
//  @return []*models.Community
//  @return error
//
func GetCommunityList() ([]*models.Community, error) {
	// 从数据库中找
	return mysql.GetCommunityList()

}

//
//  GetCommunityById
//  @Description: 根据社区id获取社区详情
//  @param cid
//  @return *models.Community
//  @return error
//
func GetCommunityById(cid int64) (*models.Community, error) {
	return mysql.GetCommunityById(cid)
}
