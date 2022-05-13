/**
  @author: ZYL
  @date:
  @note
*/
package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID         int64  `db:"id"`
	UserId     int64  `db:"user_id"`
	Username   string `db:"username"`
	Password   string `db:"password"`
	HeadPic    string `db:"head_pic"`
	Token      string
	Email      sql.NullString `db:"email"`
	Telephone  sql.NullString `db:"telephone"`
	Gender     sql.NullInt16  `db:"gender"`
	CreateTime *time.Time     `db:"create_time"`
	UpdateTime *time.Time     `db:"update_time"`
}
