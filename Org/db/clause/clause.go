package db_clause

import (
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
)

func GetClauseCount(where g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(config.ClauseTbName).Where("`delete`=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Count()
	return r, err
}

func GetClauses(offset int, limit int, where g.Map) (g.List, error) {
	db := g.DB()
	sql := db.Table(config.ClauseTbName + " c")
	sql.LeftJoin(config.UserTbName+" u", "c.author_id=u.user_id")
	sql.Fields("c.*,u.user_name as author_name")
	sql.Where("c.delete=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Limit(offset, limit).OrderBy("id desc").Select()
	return r.ToList(), err
}

func GetClauseTitle(offset int, limit int, departmentId int, title string) (g.List, error) {
	db := g.DB()
	sql := db.Table(config.ClauseTbName)
	if departmentId != 0 {
		sql.Where("`delete`=? AND (department_id=? OR department_id=-1)", 0, departmentId)
	} else {
		sql.Where("`delete`=? AND department_id=-1", 0)
	}
	sql.And("title like ?", "%"+title+"%")
	r, err := sql.Limit(offset, limit).OrderBy("id desc").Select()
	return r.ToList(), err
}

func GetClause(id int) (entity.Clause, error) {
	var Clause entity.Clause
	db := g.DB()
	sql := db.Table(config.ClauseTbName + " c")
	sql.LeftJoin(config.UserTbName+" u", "c.author_id=u.user_id")
	sql.Fields("c.*,u.user_name as author_name")
	sql.Where("c.delete=?", 0)
	sql.And("c.id=?", id)
	r, err := sql.One()
	_ = r.ToStruct(&Clause)
	return Clause, err
}

func AddClause(Clause g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	r, err := db.Insert(config.ClauseTbName, Clause)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func UpdateClause(id int, Clause g.Map) (int, error) {
	var rows int64 = 0
	db := g.DB()
	r, err := db.Table(config.ClauseTbName).Where("id=?", id).Data(Clause).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}

func DelClause(id int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	r, err := db.Table(config.ClauseTbName).Where("id=?", id).Data(g.Map{"delete": 1}).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}
