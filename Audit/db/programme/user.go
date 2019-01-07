package db_programme

import (
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addUser(ctx *gdb.TX, ProgrammeId int, data []g.Map) (int, error) {
	l := len(data)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		data[i]["programme_id"] = ProgrammeId
	}
	res, err := ctx.BatchInsert(table.ProgrammeUser, data, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func updateUser(ctx *gdb.TX, ProgrammeId int, data []g.Map) (int, error) {
	l := len(data)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		data[i]["programme_id"] = ProgrammeId
		data[i]["delete"] = 0
	}
	res, err := ctx.BatchSave(table.ProgrammeUser, data, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func editUser(tx *gdb.TX, programmeId int, update []g.Map) (int, error) {
	l := len(update)
	_, _ = delUser(tx, programmeId)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		update[i]["programme_id"] = programmeId
		update[i]["delete"] = 0
	}
	r, err := tx.BatchSave(table.ProgrammeUser, update, 5)
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func delUser(tx *gdb.TX, programmeId int) (int, error) {
	r, err := tx.Table(table.ProgrammeUser).Where("programme_id=?", programmeId).Data(g.Map{"delete": 1}).Update()
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func GetUser(programmeId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(table.ProgrammeUser + " p")
	sql.LeftJoin(table.User+" u", "p.user_id=u.user_id")
	sql.Where("p.programme_id=?", programmeId)
	sql.And("p.delete=?", 0)
	sql.Fields("p.*,u.user_name")
	sql.OrderBy("p.order asc")
	res, err := sql.All()
	return res.ToList(), err
}
