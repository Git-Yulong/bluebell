/**
  @author: ZYL
  @date:
  @note
*/
package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"path"
	"projects/bluebell/dao/redis"
	"projects/bluebell/logic"
	"projects/bluebell/models"
	"projects/bluebell/utils"
)

// TelCheckCodeHandler 向手机发送验证码
func TelCheckCodeHandler(c *gin.Context) {
	// 1 从前端获取用户手机号， 验证手机号的合法性
	// 2 验证用户是否存在（验证手机号对应的用户id是否存在）
	// 3 生成验证码, 并将验证码存入redis
	// 4 向对应手机发送验证码

	// 1 从前端获取用户手机号， 验证手机号的合法性
	mobile := c.Param("tel") // 获取手机
	flg := utils.IsMobile(mobile)
	if !flg {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "手机号不存在或者非法",
		})
		return
	}

	//// 2 验证用户是否存在（验证手机号对应的用户id是否存在）
	//err, flg := logic.CheckUserExistByTel(mobile)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"msg": "手机号不存在或者非法",
	//	})
	//	return
	//}

	// 3 生成随机验证码
	code := utils.TelCheckCode()
	err, s := utils.SendMsg(mobile, code)

	if s == "success" {
		err := logic.SaveTelCheckCode(mobile, code)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "验证法发送失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":            fmt.Sprintf("验证码已经发送至%s, %s", mobile, "3分钟后失效"),
			"tel_check_code": code,
		})
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "验证法发送失败",
			"err": err.Error(),
		})
		return
	}

}

func SignUpHandler(context *gin.Context) {
	// 1 获取前端的注册参数
	// 2 检查手机号码是否合法
	// 3 从redis中获取手机号码对应的验证码 并设置其过期
	// 4 校验其他信息
	// 5 存入数据库

	// 1 获取前端的注册参数
	p := new(models.ParamSignUp)
	if err := context.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid params ", zap.Error(err))
		ResponseErrorWithMsg(context, CodeInvalidParam, err.Error())
		return
	}
	if flg := utils.IsMobile(p.Telephone); !flg {
		ResponseErrorWithMsg(context, CodeInvalidParam, "手机号码不合法")
		return
	}

	// 获取手机验证码
	key := "user:tel:" + p.Telephone
	telCheckCode, err := redis.GetString(key)
	if telCheckCode != p.TelCheckCode {
		ResponseErrorWithMsg(context, CodeInvalidParam, "验证码错误")
		return
	}

	// 业务处理
	err = logic.SignUp(p)
	if err != nil {
		zap.L().Error("logic.SignUp, 保存用户失败 ", zap.Error(err))
		ResponseErrorWithMsg(context, CodeServerBusy, "保存用户失败")
		return
	}
	// 设置 key 过期
	redis.ExpireKey(key, 0)
	// 返回响应
	ResponseSuccess(context, nil)

}

type User struct {
	Name     string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginHandler 用户登录处理Handler
// 参数校验
func LoginHandler(context *gin.Context) {
	p := new(models.ParamLogin)
	// 参数绑定
	if err := context.BindJSON(p); err != nil {
		zap.L().Error("Login with invalid params ", zap.Error(err))
		ResponseErrorWithMsg(context, CodeInvalidParam, err)
		return
	}
	user, err := logic.Login(p)
	if err != nil {
		ResponseError(context, CodeInvalidParam)
		return
	}

	fmt.Println(user)

	ResponseSuccess(context, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserId), // id值大于1<<53-1  int64类型的最大值是1<<63-1
		"user_name": user.Username,
		"token":     user.Token,
	})
}

// 更新头像
func UploadHeadPicHandler(context *gin.Context) {
	userId := context.PostForm("user_id")
	file, err := context.FormFile("file")
	if err != nil {
		zap.L().Error("获取文件失败，err=", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	// 文件保存路径
	dir := path.Join(viper.GetString("file.base_path"), viper.GetString("file.head_pic"), userId)
	fmt.Println(dir)
	if _, err := os.Stat(dir); err != nil {
		// 文件夹不存在
		err = os.Mkdir(dir, 0711)
		if err != nil {
			zap.L().Error("创建用户头像目录失败，err=", zap.Error(err))
			return
		}
	}
	filePath := path.Join(dir, file.Filename)
	err = logic.UpdateUserHeadPic(userId, filePath)
	if err != nil {
		zap.L().Error("更新用户头像失败, err=", zap.Error(err))
		return
	}

	// 保存文件
	err = context.SaveUploadedFile(file, filePath)
	if err != nil {
		zap.L().Error("保存文件失败，err=", zap.Error(err))
	}
	ResponseSuccess(context, ResponseData{
		Code: CodeSuccess,
		Msg:  "success",
		Data: gin.H{
			"url": "http://" + context.Request.Host + filePath,
		},
	})
}

func PingHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
