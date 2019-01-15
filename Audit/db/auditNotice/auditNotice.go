package db_auditNotice

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/Audit/fun"
	entity2 "auditIntegralSys/Org/entity"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"strings"
)

func getListSql(db gdb.DB, authorInfo entity2.User, where g.Map) *gdb.Model {
	sql := db.Table(table.AuditNotice + " an")
	sql.LeftJoin(table.Draft+" d", "an.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "d.programme_id=p.id")
	sql.LeftJoin(table.Department+" dd", "d.department_id=dd.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")
	sql.Where("an.delete=?", 0)
	sql.GroupBy("an.id")

	sql = fun.CheckIsMyData(*sql, authorInfo, where)

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
		"an.*",
		"p.start_time",
		"p.end_time",
		"d.query_start_time,d.query_end_time",
		"dd.name as department_name",
		"dq.name as query_department_name",
		"p.title as programme_title",
	}
	sql.Fields(strings.Join(fields, ","))
	r, err := sql.Limit(offset, limit).OrderBy("an.id desc").Select()
	return r.ToList(), err
}

func Add(tx *gdb.TX, data g.Map) (int, error) {
	r, err := tx.Table(table.AuditNotice).Data(data).Insert()
	id, _ := r.LastInsertId()
	return int(id), err
}

func Get(id int, where ...g.Map) (entity.AuditNoticeItem, error) {
	db := g.DB()
	auditNoticeItem := entity.AuditNoticeItem{}
	sql := db.Table(table.AuditNotice)
	sql.Where("id=? AND `delete`=?", id, 0)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.One()
	if err == nil {
		_ = r.ToStruct(&auditNoticeItem)
	}
	return auditNoticeItem, err
}

func GetByDraft(draftId int, where ...g.Map) (entity.AuditNoticeItem, error) {
	db := g.DB()
	auditNoticeItem := entity.AuditNoticeItem{}
	sql := db.Table(table.AuditNotice)
	sql.Where("draft_id=? AND `delete`=?", draftId, 0)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.One()
	if err == nil {
		_ = r.ToStruct(&auditNoticeItem)
	}
	return auditNoticeItem, err
}

func GetLastOne(where ...g.Map) (entity.AuditNoticeItem, error) {
	db := g.DB()
	auditNoticeItem := entity.AuditNoticeItem{}
	sql := db.Table(table.AuditNotice)
	sql.Where("`delete`=?", 0)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.OrderBy("id desc").One()
	if err == nil {
		_ = r.ToStruct(&auditNoticeItem)
	}
	return auditNoticeItem, err
}
