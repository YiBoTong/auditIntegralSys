package db_log

import (
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func getListSql(db gdb.DB, authorId int, where g.Map) *gdb.Model {
	sql := db.Table(table.Log + " l")
	sql.LeftJoin(table.User+" u", "l.user_id=u.user_id")
	sql.Where("l.user_id=? AND l.delete=?", authorId, 0)
	// 项目名称模糊查询
	if where["msg"] != nil && where["msg"] != "" {
		sql.And("l.msg like ?", strings.Replace("%?%", "?", gconv.String(where["msg"]), 1))
		delete(where, "msg")
	}
	if len(where) > 0 {
		sql.And(where)
	}
	return sql
}

func GetLogCount(authorId int, where g.Map) (int, error) {
	db := g.DB()
	r, err := getListSql(db, authorId, where).Count()
	return r, err
}

func GetLogs(authorId, offset, limit int, where g.Map) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := getListSql(db, authorId, where)
	sql.Fields("l.*,u.user_name")
	r, err := sql.Limit(offset, limit).OrderBy("l.id desc").Select()
	return r.ToList(), err
}

func DelLog(logId int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	r, err := db.Table(table.Log).Where("id=?", logId).Data(g.Map{"delete": 1}).Update()
	if err == nil {
		rows, err = r.RowsAffected()
	}
	return int(rows), err
}

func AddLog(log g.Map) (int, error) {
	db := g.DB()
	var lastId int64 = 0
	r, err := db.Table(table.Log).Data(log).Insert()
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}
