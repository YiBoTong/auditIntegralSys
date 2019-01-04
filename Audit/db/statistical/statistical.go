package db_statistical

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
)

func Count(where g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(table.AuditReport + " ar")
	sql.LeftJoin(table.Draft+" d", "ar.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "ar.programme_id=p.id")
	sql.LeftJoin(table.Department+" dd", "d.department_id=dd.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")
	sql.Where("ar.delete=? AND ar.state=?", 0, state.Publish)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Count()
	return r, err
}

// 审计报告上报了的才能出现统计
func List(offset int, limit int, where g.Map) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.AuditReport + " ar")
	sql.LeftJoin(table.Draft+" d", "ar.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "ar.programme_id=p.id")
	sql.LeftJoin(table.Department+" dd", "d.department_id=dd.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")
	sql.Fields("d.*,ar.*,p.start_time,p.end_time,dd.name as department_name,dq.name as query_department_name,p.title as programme_title")
	sql.Where("ar.delete=? AND ar.state=?", 0, state.Publish)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Limit(offset, limit).OrderBy("ar.id desc").Select()
	return r.ToList(), err
}

func Get(id int, where ...g.Map) (entity.StatisticalListItem, error) {
	db := g.DB()
	sql := db.Table(table.AuditReport + " ar")
	sql.LeftJoin(table.Draft+" d", "ar.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "ar.programme_id=p.id")
	sql.LeftJoin(table.Department+" dd", "d.department_id=dd.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")
	sql.Fields("d.*,ar.*,p.start_time,p.end_time,dd.name as department_name,dq.name as query_department_name,p.title as programme_title")
	sql.Where("ar.id=? AND ar.delete=? AND ar.state=?", id, 0, state.Publish)
	if len(where) > 0 {
		sql.And(where[0])
	}
	item := entity.StatisticalListItem{}
	r, err := sql.One()
	_ = r.ToStruct(&item)
	return item, err
}
