package check

import (
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g/util/gconv"
)

func Password(userCode int, userPd string, sqlPd string) bool {
	return userPd != "" && util.GetPasswordStr(userPd, gconv.String(userCode)) == sqlPd
}