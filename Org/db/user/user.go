package db_user

import (
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
)

func HasUserCode(userCode string) (bool, entity.User, error) {
	db := g.DB()
	var user entity.User
	hasUserCode := false
	sql := db.Table(table.User).Where("`delete`=?", 0)
	sql.And("user_code=?", userCode)
	sql.Limit(0, 1)
	r, err := sql.One()
	_ = r.ToStruct(&user)
	if err == nil && user.UserId != 0 {
		hasUserCode = true
	}
	return hasUserCode, user, err
}

func GetUserCount(where g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(table.User).Where("`delete`=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Count()
	return r, err
}

func GetUsers(offset int, limit int, where g.Map) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(table.User + " u")
	sql.LeftJoin(table.Department+" d", "u.department_id=d.id")
	sql.Fields("u.*,d.name as department_name")
	sql.Where("u.delete=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Limit(offset, limit).OrderBy("u.user_id desc").Select()
	return r.ToList(), err
}

func AddUser(user []g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	// 批次5条数据写入
	r, err := db.BatchInsert(table.User, user, 5)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func GetUser(userId int) (entity.User, error) {
	var user entity.User
	db := g.DB()
	sql := db.Table(table.User + " u")
	sql.LeftJoin(table.Department+" d", "u.department_id=d.id")
	sql.Fields("u.*,d.name as department_name")
	sql.Where("u.user_id=?", userId)
	sql.And("u.delete=?", 0)
	r, err := sql.One()
	_ = r.ToStruct(&user)
	return user, err
}

func UpdateUser(userId int, user g.Map) (int, error) {
	var rows int64 = 0
	db := g.DB()
	r, err := db.Table(table.User).Where("user_id=?", userId).Data(user).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}

func DelUser(userId int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	r, err := db.Table(table.User).Where("user_id=?", userId).Data(g.Map{"delete": 1}).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}
