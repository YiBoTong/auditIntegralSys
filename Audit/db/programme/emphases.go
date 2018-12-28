package db_programme

import (
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addEmphases(ctx *gdb.TX, ProgrammeId int, data []g.Map) (int, error) {
	l := len(data)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		data[i]["programme_id"] = ProgrammeId
	}
	res, err := ctx.BatchInsert(table.ProgrammeEmphases, data, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func updateEmphases(ctx *gdb.TX, ProgrammeId int, data []g.Map) (int, error) {
	l := len(data)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		data[i]["programme_id"] = ProgrammeId
		data[i]["delete"] = 0
	}
	res, err := ctx.BatchSave(table.ProgrammeEmphases, data, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func editEmphases(tx *gdb.TX, programmeId int, update []g.Map) (int, error) {
	l := len(update)
	_, _ = delEmphases(tx, programmeId)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		update[i]["programme_id"] = programmeId
		update[i]["delete"] = 0
	}
	r, err := tx.BatchSave(table.ProgrammeEmphases, update, 5)
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func delEmphases(tx *gdb.TX, programmeId int) (int, error) {
	r, err := tx.Table(table.ProgrammeEmphases).Where("programme_id=?", programmeId).Data(g.Map{"delete": 1}).Update()
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func GetEmphases(programmeId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(table.ProgrammeEmphases).Where("programme_id=?", programmeId)
	sql.And("`delete`=?", 0)
	sql.OrderBy("`order` asc")
	res, err := sql.All()
	return res.ToList(), err
}
