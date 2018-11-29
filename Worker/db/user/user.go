package db_user

import (
	"auditIntegralSys/SystemSetup/entity"
	"auditIntegralSys/Worker/check"
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
)

func Login(userCode int, password string) (bool, int, error) {
	db := g.DB()
	var userLoginInfo entity.LoginInfo
	checkPd := false
	sql := db.Table(config.LoginTbName + " l")
	sql.InnerJoin(config.UserTbName+" u", "l.user_code=u.user_code")
	sql.Fields("l.*,u.user_id")
	sql.Where("l.delete=?", 0)
	sql.And("u.delete=?", 0)
	sql.And("l.user_code=?", userCode)
	sql.And("l.is_use=?", 1)
	res, err := sql.One()
	res.ToStruct(&userLoginInfo)
	if err == nil {
		checkPd = check.Password(userCode, password, userLoginInfo.Password)
	}
	return checkPd, userLoginInfo.UserId, err
}
