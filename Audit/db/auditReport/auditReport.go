package db_auditReport

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/Audit/fun"
	entity2 "auditIntegralSys/Org/entity"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
	"time"
)

func getListSql(db gdb.DB, authorInfo entity2.User, where g.Map) *gdb.Model {
	sql := db.Table(table.AuditReport + " ar")
	sql.LeftJoin(table.Draft+" d", "ar.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "ar.programme_id=p.id")
	sql.LeftJoin(table.Department+" dd", "d.department_id=dd.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")
	sql.Where("ar.delete=?", 0)
	sql.GroupBy("ar.id")
	// 项目名称模糊查询
	if where["project_name"] != nil && where["project_name"] != "" {
		sql.And("d.project_name like ?", strings.Replace("%?%", "?", gconv.String(where["project_name"]), 1))
		delete(where, "project_name")
	}
	// 查询自己和别人已发布并且公布自己的数据
	//sql.And("(c.author_id=? OR (c.author_id!=? AND c.state=?))", authorId, authorId, state.Publish)
	if len(where) > 0 {
		sql.And(where)
	}
	return sql
}

func Count(userInfo entity2.User, where g.Map) (int, error) {
	db := g.DB()
	r, err := getListSql(db, userInfo, where).Count()
	return r, err
}

func List(userInfo entity2.User, offset int, limit int, where g.Map) (g.List, error) {
	db := g.DB()
	sql := getListSql(db, userInfo, where)
	fields := []string{
		"d.*",
		"ar.*",
		"p.start_time",
		"p.end_time",
		"d.query_start_time,d.query_end_time",
		"dd.name as department_name",
		"dq.name as query_department_name",
		"p.title as programme_title",
	}
	sql.Fields(strings.Join(fields, ","))
	r, err := sql.Limit(offset, limit).OrderBy("ar.id desc").Select()
	return r.ToList(), err
}

func Add(tx gdb.TX, data g.Map) (int, error) {
	auditReport, _ := GetLastOne()
	year := time.Now().Year()
	number := fun.CreateNumber(auditReport.Year, auditReport.Number)
	data["year"] = year
	data["number"] = number
	r, err := tx.Table(table.AuditReport).Data(data).Insert()
	id, _ := r.LastInsertId()
	return int(id), err
}

func Get(id int, where ...g.Map) (entity.AuditReportItem, error) {
	db := g.DB()
	auditReportItem := entity.AuditReportItem{}
	sql := db.Table(table.AuditReport)
	sql.Where("id=? AND `delete`=?", id, 0)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.One()
	if err == nil {
		_ = r.ToStruct(&auditReportItem)
	}
	return auditReportItem, err
}

func update(tx gdb.TX, id int, data g.Map) (int, error) {
	sql := tx.Table(table.AuditReport).Data(data).Where("id=? AND state=? AND `delete`=?", id, state.Draft, 0)
	r, err := sql.Update()
	row, _ := r.RowsAffected()
	return int(row), err
}

func Update(id int, data g.Map, where ...g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(table.AuditReport).Data(data).Where("id=? AND `delete`=?", id, 0)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.Update()
	row, _ := r.RowsAffected()
	return int(row), err
}

func Edit(id int, basicInfo, reason, plan string, auditReportData g.Map) (int, error) {
	db := g.DB()
	rows := 0
	tx, err := db.Begin()
	if err == nil {
		rows, err = update(*tx, id, auditReportData)
	}
	if err == nil && rows != 0 {
		_, _ = delBasicInfo(*tx, id)
		_, err = addBasicInfo(*tx, id, basicInfo)
	}
	if err == nil && rows != 0 {
		_, _ = delReason(*tx, id)
		_, err = addReason(*tx, id, reason)
	}
	if err == nil && rows != 0 {
		_, _ = delPlan(*tx, id)
		_, err = addPlan(*tx, id, plan)
	}
	if err == nil {
		_ = tx.Commit()
	} else {
		rows = 0
		_ = tx.Rollback()
	}
	return rows, err
}

func GetLastOne() (entity.AuditReportItem, error) {
	db := g.DB()
	confirmation := entity.AuditReportItem{}
	sql := db.Table(table.AuditReport).Where("`delete`=?", 0)
	sql.OrderBy("id desc")
	r, err := sql.One()
	_ = r.ToStruct(&confirmation)
	return confirmation, err
}
