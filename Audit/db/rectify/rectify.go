package db_rectify

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)


func Count(where g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(table.Rectify + " r")
	sql.LeftJoin(table.Draft+" d", "r.draft_id=d.id")
	sql.Where("r.delete=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Count()
	return r, err
}

func List(offset int, limit int, where g.Map) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.Rectify + " r")
	sql.LeftJoin(table.Draft+" d", "r.draft_id=d.id")
	sql.LeftJoin(table.Programme+" pt", "d.department_id=pt.id")
	sql.LeftJoin(table.Programme+" pq", "d.query_department_id=pq.id")
	sql.Fields("d.*,r.*,pt.title as department_name,pq.title as query_department_name")
	sql.Where("r.delete=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Limit(offset, limit).OrderBy("r.id desc").Select()
	return r.ToList(), err
}

func Get(id int) (entity.RectifyItem, error) {
	db := g.DB()
	rectify := entity.RectifyItem{}
	sql := db.Table(table.Rectify + " r")
	sql.LeftJoin(table.Draft+" d", "r.draft_id=d.id")
	sql.LeftJoin(table.Programme+" pt", "d.department_id=pt.id")
	sql.LeftJoin(table.Programme+" pq", "d.query_department_id=pq.id")
	sql.Fields("d.*,r.*,pt.title as department_name,pq.title as query_department_name")
	sql.Where("r.delete=?", 0)
	sql.And("r.id=?", id)
	r, err := sql.One()
	if err == nil {
		_ = r.ToStruct(&rectify)
	}
	return rectify, err
}

func Add(tx gdb.TX,confirmationId, draftId int) (int, error) {
	sql := tx.Table(table.Rectify).Data(g.Map{
		"confirmation_id": confirmationId,
		"draft_id":        draftId,
	})
	r, err := sql.Insert()
	id, _ := r.LastInsertId()
	return int(id), err
}