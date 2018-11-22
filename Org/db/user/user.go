package db_user

import (
	"auditIntegralSys/_public/config"
	"auditIntegralSys/SystemSetup/entity"
	"gitee.com/johng/gf/g"
)

func HasUserCode(userCode int) (bool, entity.User, error) {
	db := g.DB()
	var user entity.User
	hasUserCode := false
	sql := db.Table(config.UserTbName).Where("`delete`=?", 0)
	sql.And("user_code=?", userCode)
	r, err := sql.One()
	r.ToStruct(&user)
	if err == nil && user.UserId > 0 {
		hasUserCode = true
	}
	return hasUserCode, user, err
}
