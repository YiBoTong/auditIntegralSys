package check

import (
	"auditIntegralSys/_public/util"
)

func Password(userCode string, userPd string, sqlPd string) bool {
	return userPd != "" && util.GetPasswordStr(userPd, userCode) == sqlPd
}

func PasswordLen(password string) bool {
	len := len(password)
	return len >= 6 && len <= 18
}
