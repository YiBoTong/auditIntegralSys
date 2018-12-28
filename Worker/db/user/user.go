package db_user

import (
	"auditIntegralSys/SystemSetup/entity"
	"auditIntegralSys/Worker/check"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
)

func Login(userCode int, password string) (bool, entity.LoginInfo, error) {
	db := g.DB()
	var userLoginInfo entity.LoginInfo
	checkPd := false
	sql := db.Table(table.Login + " l")
	sql.InnerJoin(table.User+" u", "l.user_code=u.user_code")
	sql.Fields("l.*,u.user_id")
	sql.Where("l.delete=?", 0)
	sql.And("u.delete=?", 0)
	sql.And("l.user_code=?", userCode)
	sql.And("l.is_use=?", 1)
	sql.OrderBy("l.login_id desc")
	res, err := sql.One()
	res.ToStruct(&userLoginInfo)
	if err == nil {
		checkPd = check.Password(userCode, password, userLoginInfo.Password)
	}
	return checkPd, userLoginInfo, err
}