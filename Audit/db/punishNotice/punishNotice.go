package db_punishNotice

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/table"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"strings"
)

func Count(where g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(table.PunishNotice + " p")
	sql.LeftJoin(table.Draft+" d", "p.draft_id=d.id")
	sql.Where("p.delete=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Count()
	return r, err
}

func List(offset int, limit int, where g.Map) (g.List, error) {
	db := g.DB()
	fields := []string{
		"d.*",
		"pn.*",
		"u.user_name",
		"dt.name as department_name",
		"dq.name as query_department_name",
		"p.title as programme_title",
		"p.start_time",
		"p.end_time",
		"p.plan_start_time",
		"p.plan_end_time",
	}
	sql := db.Table(table.PunishNotice + " pn")
	sql.LeftJoin(table.Draft+" d", "pn.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "d.programme_id=p.id")
	sql.LeftJoin(table.Department+" dt", "d.department_id=dt.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")
	sql.LeftJoin(table.User+" u", "pn.user_id=u.user_id")
	sql.Fields(strings.Join(fields, ","))
	sql.Where("pn.delete=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Limit(offset, limit).OrderBy("pn.id desc").Select()
	return r.ToList(), err
}

func Get(id int, where ...g.Map) (entity.PunishNoticeItem, error) {
	db := g.DB()
	fields := []string{
		"d.*",
		"pn.*",
		"u.user_name",
		"dt.name as department_name",
		"dq.name as query_department_name",
		"p.title as programme_title",
		"p.start_time",
		"p.end_time",
		"p.plan_start_time",
		"p.plan_end_time",
	}
	confirmation := entity.PunishNoticeItem{}
	sql := db.Table(table.PunishNotice + " pn")
	sql.LeftJoin(table.Draft+" d", "pn.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "d.programme_id=p.id")
	sql.LeftJoin(table.Department+" dt", "d.department_id=dt.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")
	sql.LeftJoin(table.User+" u", "pn.user_id=u.user_id")
	sql.Fields(strings.Join(fields, ","))
	sql.Where("pn.delete=?", 0)
	sql.And("pn.id=?", id)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.One()
	if err == nil {
		_ = r.ToStruct(&confirmation)
	}
	return confirmation, err
}

func Add(tx gdb.TX, confirmationId, draftId int, userIds []int) (int, error) {
	data := g.List{}
	nowTime := util.GetLocalNowTimeStr()
	for _, v := range userIds {
		data = append(data, g.Map{
			"confirmation_id": confirmationId,
			"draft_id":        draftId,
			"user_id":         v,
			"time":            nowTime,
		})
	}
	if len(data) < 1 {
		return 0, nil
	}
	r, err := tx.BatchInsert(table.PunishNotice, data, 3)
	id, _ := r.LastInsertId()
	return int(id), err
}

func Update(id int, data g.Map, where ...g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(table.PunishNotice).Data(data)
	sql.Where("`delete`=?", 0)
	sql.And("id=?", id)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.Update()
	row, _ := r.RowsAffected()
	return int(row), err
}
