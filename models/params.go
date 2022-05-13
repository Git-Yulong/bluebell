/**
  @author: ZYL
  @date:
  @note
*/
package models

type ParamSignUp struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	RePassword   string `json:"re_password" binding:"required,eqfield=Password"`
	Email        string `json:"email"`
	Gender       int16  `json:"gender"`
	Telephone    string `json:"tel" binding:"required"`
	TelCheckCode string `json:"tel_check_code" binding:"required"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParmaVotePost struct {
	UserId    int64 `json:"user_id,string" binding:"required"`
	PostId    int64 `json:"post_id,string" binding:"required"`
	Direction int8  `json:"direction" binding:"required,oneof=0 1 -1"`
}
