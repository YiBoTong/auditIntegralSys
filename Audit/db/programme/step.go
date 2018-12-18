package db_programme

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addStep(ctx *gdb.TX, ProgrammeId int, data []g.Map) (int, error) {
	var rows int64 = 0
	len := len(data)
	for i := 0; i < len; i++ {
		data[i]["programme_id"] = ProgrammeId
	}
	res, err := ctx.BatchInsert(config.ProgrammeStepTbName, data, 5)
	if err == nil {
		rows, _ = res.RowsAffected()
	}
	return int(rows), err
}

func delStep(programmeId int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	r, err := db.Table(config.ProgrammeStepTbName).Where("programme_id=?", programmeId).Data(g.Map{"delete": 1}).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
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