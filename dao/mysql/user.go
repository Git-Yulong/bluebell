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
	"projects/bluebell/utils"
)

// CheckUserExist 根据用户名进行查找
func CheckUserExist(username string) (bool, error) {
	sqlStr := "select count(*) from tb_user where username=?"
	cnt := 0
	err := db.Get(&cnt, sqlStr, username)
	if err != nil {
		zap.L().Error("CheckUserExist:根据用户名查询用户错误 ", zap.Error(err))
		return false, err
	}

	return cnt > 0, nil
}

// CheckUserExist 根据用户名进行查找
func CheckUserExistByUserId(userId int64) (bool, error) {
	sqlStr := "select count(*) from tb_user where user_id=?"
	cnt := 0
	err := db.Get(&cnt, sqlStr, userId)
	if err != nil {
		zap.L().Error("CheckUserExist:根据用户名查询用户错误 ", zap.Error(err))
		return false, err
	}

	return cnt > 0, nil
}

// SaveUser 保存用户到数据库
func SaveUser(user *models.User) error {
	user.Password = utils.MD5Encrypt(user.Password)

	sqlStr := `insert into tb_user(user_id, username, password, email, telephone, gender, head_pic) values(?,?,?,?,?,?,?)`
	// 这里返回的是 Result{last_id, row_effected} 可以不用
	_, err := db.Exec(sqlStr, user.UserId, user.Username, user.Password, user.Email.String, user.Telephone.String, user.Gender.Int16, user.HeadPic)
	if err != nil {
		zap.L().Error("mysql.SaveUser failed ", zap.Error(err))
		return err
	}
	return nil
}

// Login 用户登录
// 登录成功，err为nil, 否则不为nil
func Login(user *models.User) (*models.User, error) {
	oldPWD := utils.MD5Encrypt(user.Password)                                    // 加密后再跟数据库中的密码对比
	sqlStr := "select user_id, username, password from tb_user where username=?" // 创建了索引idx_username_pwd(username, password) 不要 select *
	var dbUser models.User
	err := db.Get(&dbUser, sqlStr, user.Username)
	if err != nil {
		// 不要告诉客户端用户不存在，为了安全起见
		return nil, err
	}
	// err == nil，但是密码不正确
	if oldPWD != dbUser.Password {
		// 密码不正确
		zap.L().Error("mysql.Login failed ", zap.Error(err))
		return nil, errors.New("用户名或者密码错误密码错误！")
	}
	user.Password = ""
	return &dbUser, nil
}

// GetUserByUid 根据用户id 查找用户
// uid 用户id
func GetUserByUid(uid int64) (*models.User, error) {
	sqlStr := `SELECT * FROM tb_user WHERE user_id=?`
	var dbUser models.User
	err := db.Get(&dbUser, sqlStr, uid)
	if err != nil {
		logMsg := fmt.Sprintf("mysql.GetUserByUid(%d), err=%s", uid, err)
		zap.L().Error(logMsg)
		return nil, err
	}
	return &dbUser, nil
}

// CheckUserExistByTel 判断手机号是否存在
func CheckUserExistByTel(tel string) (error, bool) {
	sqlStr := `SELECT * FROM tb_user WHERE telephone=?`
	var dbUser models.User

	err := db.Get(&dbUser, sqlStr, tel)
	if err != nil {
		zap.L().Error("用户不存在,id=" + tel)
		return err, false
	}
	return nil, true
}

// 更新用户头像
func UpdateUserHeadPic(userId int64, path string) error {
	sqlStr := `update tb_user set head_pic =? where user_id=?`
	_, err := db.Exec(sqlStr, path, userId)
	if err != nil {
		zap.L().Error("更新用户头像失败", zap.Error(err))
		return err
	}
	return nil
}
