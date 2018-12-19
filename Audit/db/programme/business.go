package db_programme

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addBusiness(ctx *gdb.TX, ProgrammeId int, data []g.Map) (int, error) {
	l := len(data)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < len(data); i++ {
		data[i]["programme_id"] = ProgrammeId
	}
	res, err := ctx.BatchInsert(config.ProgrammeBusinessTbName, data, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func updateBusiness(ctx *gdb.TX, ProgrammeId int, data []g.Map) (int, error) {
	l := len(data)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < len(data); i++ {
		data[i]["programme_id"] = ProgrammeId
		data[i]["delete"] = 0
	}
	res, err := ctx.BatchSave(config.ProgrammeBusinessTbName, data, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func delBusiness(tx *gdb.TX, programmeId int) (int, error) {
	r, err := tx.Table(config.ProgrammeBusinessTbName).Where("programme_id=?", programmeId).Data(g.Map{"delete": 1}).Update()
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func editBusiness(tx *gdb.TX, programmeId int, update []g.Map) (int, error) {
	l := len(update)
	_, _ = delBusiness(tx, programmeId)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		update[i]["programme_id"] = programmeId
		update[i]["delete"] = 0
	}
	r, err := tx.BatchSave(config.ProgrammeBusinessTbName, update, 5)
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func GetBusiness(programmeId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.ProgrammeBusinessTbName).Where("programme_id=?", programmeId)
	sql.And("`delete`=?", 0)
	sql.OrderBy("`order` asc")
	res, err := sql.All()
	return res.ToList(), err
}
