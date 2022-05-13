/**
  @author: ZYL
  @date:
  @note
*/
package mysql

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"projects/bluebell/models"
)

// 获取全部社区信息
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `SELECT community_id, community_name FROM tb_community`
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityById(cid int64) (*models.Community, error) {
	sqlStr := `SELECT community_id, community_name, introduction FROM tb_community WHERE community_id = ?`
	var dbModel models.Community
	if err := db.Get(&dbModel, sqlStr, cid); err != nil {
		zap.L().Error("msql.GetCommunityById, 获取社区详情失败， err=", zap.Error(err))
		return nil, err
	}
	return &dbModel, nil
}

func CheckCommunityExist(cid int64) (error, bool) {
	sqlStr := `select count(*) from tb_community where community_id = ?`
	var cnt int
	err := db.Get(&cnt, sqlStr, cid)
	if err != nil {
		return err, false
	}
	if cnt > 0 {
		return nil, true
	} else {
		return nil, false
	}
}

// GetCommunityDetailByCid 获取社区信息
func GetCommunityDetailByCid(cid int64) (*models.Community, error) {
	sqlStr := `select community_id, community_name, introduction, create_time from tb_community where community_id=?`
	var community = new(models.Community)
	err := db.Get(community, sqlStr, cid)
	if err != nil {
		logMsg := fmt.Sprintf("mysql.GetCommunityDetailByCid(%d), err=%s", cid, err)
		zap.L().Error(logMsg)
		return nil, err
	}

	return community, nil

}
