package db_programme

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addBasis(ctx gdb.TX, programmeId int, data []g.Map) (int, error) {
	var rows int64 = 0
	len := len(data)
	for i := 0; i < len; i++ {
		data[i]["programme_id"] = programmeId
	}
	res, err := ctx.BatchInsert(config.ProgrammeBasisTbName, data, 5)
	if err == nil {
		rows, _ = res.RowsAffected()
	}
	return int(rows), err
}

func delBasis(ctx *gdb.TX, programmeId int) (int64, error) {
	var rows int64 = 0
	r, err := ctx.Table(config.ProgrammeBasisTbName).Where("programme_id=?", programmeId).Data(g.Map{"delete": 1}).Update()
	if err == nil {
		rows, err = r.RowsAffected()
	}
	return rows, err
}

//func updateBasis(update []g.Map) (int,error) {
//	var rows int64 = 0
//
//}
//
//func editBasis(programmeId int, add, update []g.Map) (int64, error) {
//	var rows int64 = 0
//	row, err := delBasis(programmeId)
//	rows += int64(row)
//	if len(add) >0 {
//
//	}
//	return int64(rows), err
//}

func GetBasis(programmeId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.ProgrammeBasisTbName).Where("programme_id=?", programmeId)
	sql.And("`delete`=?", 0)
	sql.OrderBy("`order` asc")
	res, err := sql.All()
	return res.ToList(), err
}
