/**
  @author: ZYL
  @date:
  @note
*/
package logic

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"path"
	"projects/bluebell/dao/mysql"
	"projects/bluebell/dao/redis"
	"projects/bluebell/models"
	"projects/bluebell/pkg/jwt"
	"projects/bluebell/pkg/snowflake"
	"strconv"
)

// FindUserByTel 根据手机号查找用户
// tel 手机号
func CheckUserExistByTel(tel string) (error, bool) {
	return mysql.CheckUserExistByTel(tel)
}

func SaveTelCheckCode(tel string, code string) error {
	key := "user:tel:" + tel
	err := redis.StringSet(key, code)
	return err

}

// FindUserByUserId 根据用户名id查找用户
func FindUserByUserId(userId string) (error, *models.User) {
	// 先去redis中找对应的uid
	// 如果不存在，则去mysql中查找
	uid, _ := strconv.ParseInt(userId, 10, 64)
	user, err := mysql.GetUserByUid(uid)
	if err != nil {
		return err, nil
	} else {
		return nil, user
	}
}

// 处理注册请求
func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户是否已经存在

	exist, err := mysql.CheckUserExist(p.Username)
	if err != nil {
		// 查询数据出错
		return err
	}
	if exist {
		// 用户已存在
		return errors.New(fmt.Sprintf("username：%s 已存在，不能重复注册!", p.Username))
	}

	// 根据雪花算法生产useId
	uid := snowflake.GenID()
	user := models.User{
		UserId:    uid,
		Username:  p.Username,
		Password:  p.Password,
		Gender:    sql.NullInt16{Int16: p.Gender, Valid: true},
		Telephone: sql.NullString{String: p.Telephone, Valid: true},
	}

	user.HeadPic = path.Join(viper.GetString("file.base"), viper.GetString("file.head_pic"), "default.jpg")
	if len(p.Email) > 0 {
		user.Email = sql.NullString{String: p.Email, Valid: true}
	} else {
		user.Email = sql.NullString{Valid: true}
	}

	// 写入数据库
	err = mysql.SaveUser(&user)
	if err != nil {
		zap.L().Error("mysql.SaveUser, 保存user出错: ", zap.Error(err))
		return err
	}
	return nil
}

// Login 登录操作
func Login(p *models.ParamLogin) (*models.User, error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	dbUser, err := mysql.Login(user)

	if err != nil {
		zap.L().Error("logic.Login failed ", zap.Error(err))
		return nil, err
	}

	token, err := jwt.GenToken(dbUser.UserId, dbUser.Username)
	if err != nil {
		zap.L().Error("生成token失败")
		return nil, err

	}
	dbUser.Token = token
	// 生成JWT token
	return dbUser, nil
}

func UpdateUserHeadPic(userId string, path string) error {
	// 检查参数
	uid, err := strconv.Atoi(userId)

	if err != nil {
		zap.L().Error("用户id不合法，err=", zap.Error(err))
		return err
	}

	// 查看用户是否存在
	flg, err := mysql.CheckUserExistByUserId(int64(uid))
	if flg == false {
		zap.L().Error("用户id不存在，err=", zap.Error(err))
		return err
	}

	// 更新数据
	err = mysql.UpdateUserHeadPic(int64(uid), path)
	if err != nil {
		return err
	}
	return nil

}
