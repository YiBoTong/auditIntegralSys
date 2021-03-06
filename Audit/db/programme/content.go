package db_programme

import (
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addContent(ctx *gdb.TX, ProgrammeId int, data []g.Map) (int, error) {
	l := len(data)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		data[i]["programme_id"] = ProgrammeId
	}
	res, err := ctx.BatchInsert(table.ProgrammeContent, data, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func updateContent(ctx *gdb.TX, ProgrammeId int, data []g.Map) (int, error) {
	l := len(data)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		data[i]["programme_id"] = ProgrammeId
		data[i]["delete"] = 0
	}
	res, err := ctx.BatchSave(table.ProgrammeContent, data, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func editContent(tx *gdb.TX, programmeId int, update []g.Map) (int, error) {
	l := len(update)
	_, _ = delContent(tx, programmeId)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		update[i]["programme_id"] = programmeId
		update[i]["delete"] = 0
	}
	r, err := tx.BatchSave(table.ProgrammeContent, update, 5)
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func delContent(tx *gdb.TX, programmeId int) (int, error) {
	r, err := tx.Table(table.ProgrammeContent).Where("programme_id=?", programmeId).Data(g.Map{"delete": 1}).Update()
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func GetContent(programmeId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(table.ProgrammeContent).Where("programme_id=?", programmeId)
	sql.And("`delete`=?", 0)
	sql.OrderBy("`order` asc")
	res, err := sql.All()
	return res.ToList(), err
}
