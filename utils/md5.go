/**
  @author: ZYL
  @date:
  @note
*/
package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 加密盐
const salt string = "zylzxy"

// MD5Encrypt MD5加密
func MD5Encrypt(password string) string {
	hash := md5.New()
	hash.Write([]byte(salt))
	return hex.EncodeToString(hash.Sum([]byte(password)))
}
