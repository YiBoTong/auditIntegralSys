package db_punishNotice

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
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
	sql := db.Table(table.PunishNotice + " pn")
	sql.LeftJoin(table.Draft+" d", "pn.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "d.programme_id=p.id")
	sql.LeftJoin(table.Department+" dt", "d.department_id=dt.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")
	sql.Fields("d.*,pn.*,dt.name as department_name,dq.name as query_department_name,p.title as programme_title")
	sql.Where("pn.delete=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Limit(offset, limit).OrderBy("pn.id desc").Select()
	return r.ToList(), err
}

func Get(id int) (entity.PunishNoticeItem, error) {
	db := g.DB()
	confirmation := entity.PunishNoticeItem{}
	sql := db.Table(table.PunishNotice + " p")
	sql.LeftJoin(table.Draft+" d", "p.draft_id=d.id")
	sql.LeftJoin(table.Programme+" pt", "d.department_id=pt.id")
	sql.LeftJoin(table.Programme+" pq", "d.query_department_id=pq.id")
	sql.Fields("d.*,p.*,pt.title as department_name,pq.title as query_department_name")
	sql.Where("p.delete=?", 0)
	sql.And("p.id=?", id)
	r, err := sql.One()
	if err == nil {
		_ = r.ToStruct(&confirmation)
	}
	return confirmation, err
}

func Add(tx gdb.TX,confirmationId, draftId int) (int, error) {
	sql := tx.Table(table.PunishNotice).Data(g.Map{
		"confirmation_id": confirmationId,
		"draft_id":        draftId,
	})
	r, err := sql.Insert()
	id, _ := r.LastInsertId()
	return int(id), err
}

func Update(id int, data g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(table.PunishNotice).Data(data)
	sql.Where("`delete`=? AND state=?", 0, state.Draft)
	sql.And("id=?", id)
	r, err := sql.Update()
	row, _ := r.RowsAffected()
	return int(row), err
}
