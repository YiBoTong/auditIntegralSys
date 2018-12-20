package db_programme

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
)

func AddDepExamines(programmeId int, data g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(config.ProgrammeDepExamineTbName)
	sql.Data(data)
	r, err := sql.Insert()
	id, _ := r.LastInsertId()
	return int(id), err
}

func GetDepExamines(programmeId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.ProgrammeDepExamineTbName + " d")
	sql.LeftJoin(config.UserTbName+" u", "d.user_id=u.user_id")
	sql.Fields("d.*,u.user_name")
	sql.Where("d.programme_id=?", programmeId)
	sql.OrderBy("d.id asc")
	r, err := sql.All()
	return r.ToList(), err
}
