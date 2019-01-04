package db_auditReport

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func Count(where g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(table.AuditReport + " ar")
	sql.LeftJoin(table.Draft+" d", "ar.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "ar.programme_id=p.id")
	sql.LeftJoin(table.Department+" dd", "d.department_id=dd.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")
	sql.Where("ar.delete=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Count()
	return r, err
}

func List(offset int, limit int, where g.Map) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.AuditReport + " ar")
	sql.LeftJoin(table.Draft+" d", "ar.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "ar.programme_id=p.id")
	sql.LeftJoin(table.Department+" dd", "d.department_id=dd.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")
	sql.Fields("d.*,ar.*,p.start_time,d.time,p.end_time,dd.name as department_name,dq.name as query_department_name,p.title as programme_title")
	sql.Where("ar.delete=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Limit(offset, limit).OrderBy("ar.id desc").Select()
	return r.ToList(), err
}

func Add(tx gdb.TX, data g.Map) (int, error) {
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
