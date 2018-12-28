package db_programme

import (
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addBasis(tx *gdb.TX, programmeId int, data []g.Map) (int, error) {
	l := len(data)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		data[i]["programme_id"] = programmeId
	}
	res, err := tx.BatchInsert(table.ProgrammeBasis, data, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func updateBasis(tx *gdb.TX, programmeId int, data []g.Map) (int, error) {
	l := len(data)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		data[i]["programme_id"] = programmeId
		data[i]["delete"] = 0
	}
	res, err := tx.BatchSave(table.ProgrammeBasis, data, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func delBasis(tx *gdb.TX, programmeId int) (int, error) {
	r, err := tx.Table(table.ProgrammeBasis).Where("programme_id=?", programmeId).Data(g.Map{"delete": 1}).Update()
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func editBasis(tx *gdb.TX, programmeId int, update []g.Map) (int, error) {
	l := len(update)
	_, _ = delBasis(tx, programmeId)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		update[i]["programme_id"] = programmeId
		update[i]["delete"] = 0
	}
	r, err := tx.BatchSave(table.ProgrammeBasis, update, 5)
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func GetBasis(programmeId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(table.ProgrammeBasis).Where("programme_id=?", programmeId)
	sql.And("`delete`=?", 0)
	sql.OrderBy("`order` asc")
	res, err := sql.All()
	return res.ToList(), err
}
