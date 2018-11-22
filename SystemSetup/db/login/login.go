package db_login

import (
	"auditIntegralSys/_public/config"
	"auditIntegralSys/SystemSetup/entity"
	"gitee.com/johng/gf/g"
)

func GetUserCount(where g.Map) (int, error) {
	db := g.DB()
	// SELECT COUNT(1) FROM login l INNER JOIN users u ON (l.user_code=u.user_code) WHERE l.delete=0 AND u.delete=0
	sql := db.Table(config.LoginTbName + " l").InnerJoin(config.UserTbName+" u", "l.user_code=u.user_code")
	sql.Fields("COUNT(1)")
	sql.Where("l.delete=?", 0)
	sql.And("u.delete=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Count()
	return r, err
}

func GetLoginList(offset int, limit int, where g.Map) ([]map[string]interface{}, error) {
	db := g.DB()
	// SELECT l.is_use,l.change_pd_time,l.login_time,l.author_id,u.*,a.user_name as author_name FROM login l INNER JOIN users u ON (l.user_code=u.user_code) LEFT JOIN users a ON (l.author_id=u.user_id) WHERE l.delete=0 AND u.delete=0 ORDER BY u.user_id desc LIMIT 0, 20
	sql := db.Table(config.LoginTbName + " l").InnerJoin(config.UserTbName+" u", "l.user_code=u.user_code")
	sql.LeftJoin(config.UserTbName+" a", "l.author_id=a.user_id")
	sql.Fields("l.is_use,l.change_pd_time,l.login_time,l.author_id,l.login_num,u.*,a.user_name as author_name")
	sql.Where("l.delete=?", 0)
	sql.And("u.delete=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Limit(offset, limit).OrderBy("u.user_id desc").Select()
	return r.ToList(), err
}

func AddLogin(login g.Map) (int, error) {
	db := g.DB()
	var lastId int64 = 0
	r, err := db.Table(config.LoginTbName).Data(login).Replace()
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func UpdateLogin(user g.Map, userCode int, deleted int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	sql := db.Table(config.LoginTbName).Data(user).Where("user_code=?", userCode)
	sql.And("`delete`=?", deleted)
	r, err := sql.Update()
	if err == nil {
		rows, err = r.RowsAffected()
	}
	return int(rows), err
}

func HasUserCode(userCode int, checkAll bool) (bool, entity.LoginInfo, error) {
	db := g.DB()
	hasUserCode := false
	var info entity.LoginInfo
	sql := db.Table(config.LoginTbName).Where("user_code=?", userCode)
	if !checkAll {
		sql.And("`delete`=?", 0)
	}
	r, err := sql.One()
	err = r.ToStruct(&info)
	if err == nil && info.LoginId > 0 {
		hasUserCode = true
	}
	return hasUserCode, info, err
}

func DelLogin(userCode int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	r, err := db.Table(config.LoginTbName).Where("user_code=?", userCode).Data(g.Map{"delete": 1}).Update()
	if err == nil {
		rows, err = r.RowsAffected()
	}
	return int(rows), err
}
