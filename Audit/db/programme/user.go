package db_programme

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addUser(ctx *gdb.Tx, ProgrammeId int, data []g.Map) (int, error) {
	var rows int64 = 0
	len := len(data)
	for i := 0; i < len; i++ {
		data[i]["programme_id"] = ProgrammeId
	}
	res, err := ctx.BatchInsert(config.ProgrammeUserTbName, data, 5)
	if err == nil {
		rows, _ = res.RowsAffected()
	}
	return int(rows), err
}