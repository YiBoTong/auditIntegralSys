package db_log

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
)

func GetLogCount(where g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(config.LogTbName).Where("`delete`=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Count()
	return r, err
}

func GetLogs(offset int, limit int, where g.Map) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.LogTbName + " l").LeftJoin(config.UserTbName+" u", "l.user_id=u.user_id")
	sql.Fields("l.*,u.user_name")
	sql.Where("l.delete=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Limit(offset, limit).OrderBy("l.id desc").Select()
	return r.ToList(), err
}

func DelLog(logId int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	r, err := db.Table(config.LogTbName).Where("id=?", logId).Data(g.Map{"delete": 1}).Update()
	if err == nil {
		rows, err = r.RowsAffected()
	}
	return int(rows), err
}

func AddLog(log g.Map) (int, error) {
	db := g.DB()
	var lastId int64 = 0
	r, err := db.Table(config.LogTbName).Data(log).Insert()
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}
