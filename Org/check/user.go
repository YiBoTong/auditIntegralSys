package check

import (
	"auditIntegralSys/Org/db/user"
	"auditIntegralSys/_public/config"
)

// 检测员工好是否存在（传userId将排除此userId后检测）
func HasUserCode(userCode string, userId int) (bool, string, error) {
	msg := ""
	hasCode, userInfo, err := db_user.HasUserCode(userCode)
	if userId != 0 && userInfo.UserId == userId {
		hasCode = false
	} else if err == nil {
		msg = config.UserCode
		if hasCode {
			msg += config.Had
		} else {
			msg += config.NoHad
		}
	}
	return hasCode, msg, err
}