package db_programme

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addStep(ctx *gdb.TX, ProgrammeId int, data []g.Map) (int, error) {
	l := len(data)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		data[i]["programme_id"] = ProgrammeId
	}
	res, err := ctx.BatchInsert(config.ProgrammeStepTbName, data, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func updateStep(ctx *gdb.TX, ProgrammeId int, data []g.Map) (int, error) {
	l := len(data)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		data[i]["programme_id"] = ProgrammeId
		data[i]["delete"] = 0
	}
	res, err := ctx.BatchSave(config.ProgrammeStepTbName, data, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func editStep(tx *gdb.TX, programmeId int, update []g.Map) (int, error) {
	l := len(update)
	_, _ = delStep(tx, programmeId)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		update[i]["programme_id"] = programmeId
		update[i]["delete"] = 0
	}
	r, err := tx.BatchSave(config.ProgrammeStepTbName, update, 5)
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func delStep(tx *gdb.TX, programmeId int) (int, error) {
	r, err := tx.Table(config.ProgrammeStepTbName).Where("programme_id=?", programmeId).Data(g.Map{"delete": 1}).Update()
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func GetStep(programmeId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.ProgrammeStepTbName).Where("programme_id=?", programmeId)
	sql.And("`delete`=?", 0)
	sql.OrderBy("`order` asc")
	res, err := sql.All()
	return res.ToList(), err
}
