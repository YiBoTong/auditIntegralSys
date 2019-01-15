package db_rectify

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/Audit/fun"
	entity2 "auditIntegralSys/Org/entity"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"strings"
	"time"
)

func getListSql(db gdb.DB, authorInfo entity2.User, where g.Map) *gdb.Model {
	sql := db.Table(table.Rectify + " r")
	sql.LeftJoin(table.Draft+" d", "r.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "d.programme_id=p.id")
	sql.LeftJoin(table.Department+" dt", "d.department_id=dt.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")
	sql.LeftJoin(table.RectifyReport+" rr", "rr.rectify_id=r.id AND rr.delete=0")
	sql.Where("r.delete=?", 0)
	sql.GroupBy("r.id")

	sql = fun.CheckIsMyData(*sql, authorInfo, where)

	if len(where) > 0 {
		sql.And(where)
	}
	return sql
}

func Count(authorInfo entity2.User, where g.Map) (int, error) {
	db := g.DB()
	r, err := getListSql(db, authorInfo, where).Count()
	return r, err
}

func List(authorInfo entity2.User, offset, limit int, where g.Map) (g.List, error) {
	db := g.DB()
	fields := []string{
		"d.*",
		"r.*",
		"dt.name as department_name",
		"dq.name as query_department_name",
		"p.title as programme_title",
		"p.start_time",
		"p.end_time",
		"p.plan_start_time",
		"p.plan_end_time",
		"rr.id as rectify_report_id",
		"rr.state as report_state",
	}
	sql := getListSql(db, authorInfo, where)
	sql.Fields(strings.Join(fields, ","))
	r, err := sql.Limit(offset, limit).OrderBy("r.id desc").Select()
	return r.ToList(), err
}

func Get(id int) (entity.RectifyItem, error) {
	db := g.DB()
	rectify := entity.RectifyItem{}
	fields := []string{
		"r.*",
		"u.user_name",
		"d.programme_id",
		"rr.id as rectify_report_id",
	}
	sql := db.Table(table.Rectify + " r")
	sql.LeftJoin(table.User+" u", "r.user_id=u.user_id")
	sql.LeftJoin(table.Draft+" d", "r.draft_id=d.id")
	sql.LeftJoin(table.RectifyReport+" rr", "r.id=rr.rectify_id")
	sql.Fields(strings.Join(fields, ","))
	sql.Where("r.delete=?", 0)
	sql.And("r.id=?", id)
	r, err := sql.One()
	if err == nil {
		_ = r.ToStruct(&rectify)
	}
	return rectify, err
}

func Add(tx gdb.TX, confirmationId, draftId int) (int, error) {
	rectify, _ := GetLastOne()
	year := time.Now().Year()
	number := fun.CreateNumber(rectify.Year, rectify.Number)
	sql := tx.Table(table.Rectify).Data(g.Map{
		"confirmation_id": confirmationId,
		"draft_id":        draftId,
		"year":            year,
		"number":          number,
	})
	r, err := sql.Insert()
	id, _ := r.LastInsertId()
	return int(id), err
}

func Update(id int, data g.Map, where ...g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(table.Rectify).Data(data)
	sql.Where("id=? AND `delete`=?", id, 0)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.Update()
	row, _ := r.RowsAffected()
	return int(row), err
}

func UpdateTX(id int, data g.Map, demand, suggest string, where ...g.Map) (int, error) {
	db := g.DB()
	row := 0
	tx, err := db.Begin()
	if err == nil {
		sql := tx.Table(table.Rectify).Data(data)
		sql.Where("id=? AND `delete`=?", id, 0)
		if len(where) > 0 {
			sql.And(where[0])
		}
		_, err = sql.Update()
		row += 1
	}
	if err == nil {
		_, _ = delDemand(*tx, id)
		_, err = addDemand(*tx, id, demand)
		row += 1
	}
	if err == nil {
		_, _ = delSuggest(*tx, id)
		_, err = addSuggest(*tx, id, suggest)
		row += 1
	}
	if err == nil {
		_ = tx.Commit()
	} else {
		_ = tx.Rollback()
	}
	return int(row), err
}

func GetLastOne() (entity.RectifyItem, error) {
	db := g.DB()
	rectify := entity.RectifyItem{}
	sql := db.Table(table.Rectify).Where("`delete`=?", 0)
	sql.OrderBy("id desc")
	r, err := sql.One()
	_ = r.ToStruct(&rectify)
	return rectify, err
}
