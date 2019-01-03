package db_integral

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addChangeScore(tx gdb.TX, data g.Map) (int, error) {
	r, err := tx.Table(table.IntegralEdit).Data(data).Insert()
	id, _ := r.LastInsertId()
	return int(id), err
}

func updateChangeScore(tx gdb.TX, changeScoreId int, data g.Map) (int, error) {
	r, err := tx.Table(table.IntegralEdit).Data(data).Where("id=? AND `delete`=0", changeScoreId).Update()
	row, _ := r.RowsAffected()
	return int(row), err
}

func delChangeScore(tx gdb.TX, integralId int) (int, error) {
	r, err := tx.Table(table.IntegralEdit).Data(g.Map{"delete": 1}).Where("integral_id=?", integralId).Update()
	row, _ := r.RowsAffected()
	return int(row), err
}

func GetChangeScore(integralId int) (entity.IntegralChangeScore, error) {
	db := g.DB()
	integralChangeScore := entity.IntegralChangeScore{}
	r, err := db.Table(table.IntegralEdit).Where("integral_id=? AND `delete`=0", integralId).One()
	_ = r.ToStruct(&integralChangeScore)
	return integralChangeScore, err
}
