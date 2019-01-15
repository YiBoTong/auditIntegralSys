package db_rectify

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addDemand(tx gdb.TX, rectifyId int, content string) (int, error) {
	r, err := tx.Table(table.RectifyDemand).Data(g.Map{
		"rectify_id": rectifyId,
		"content":     content,
	}).Insert()
	id, _ := r.LastInsertId()
	return int(id), err
}

func delDemand(tx gdb.TX, rectifyId int) (int, error) {
	r, err := tx.Table(table.RectifyDemand).Data(g.Map{"delete": 1}).Where("rectify_id=?", rectifyId).Update()
	row, _ := r.RowsAffected()
	return int(row), err
}

func GetDemand(rectifyId int) (entity.RectifyContent, error) {
	db := g.DB()
	content := entity.RectifyContent{}
	sql := db.Table(table.RectifyDemand)
	sql.Where("`delete`=? AND rectify_id=?", 0, rectifyId)
	r, err := sql.One()
	_ = r.ToStruct(&content)
	return content, err
}
