package util

import (
	"gitee.com/johng/gf/g/crypto/gmd5"
)

func GetPasswordStr(pd string, userCode string) string {
	// 密码前增加前缀
	return gmd5.Encrypt("YBT_" + userCode + pd)
}
