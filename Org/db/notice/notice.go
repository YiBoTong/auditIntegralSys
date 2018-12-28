package db_notice

import (
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
)

func GetNoticeCount(where g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(table.Notice).Where("`delete`=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Count()
	return r, err
}

func GetNotices(offset int, limit int, where g.Map) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(table.Notice).Where("`delete`=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Limit(offset, limit).OrderBy("id desc").Select()
	return r.ToList(), err
}

func GetNotice(id int) (entity.Notice, error) {
	var Notice entity.Notice
	db := g.DB()
	r, err := db.Table(table.Notice).Where("id=?", id).And("`delete`=?", 0).One()
	_ = r.ToStruct(&Notice)
	return Notice, err
}

func AddNotice(Notice g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	r, err := db.Insert(table.Notice, Notice)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func UpdateNotice(id int, Notice g.Map) (int, error) {
	var rows int64 = 0
	db := g.DB()
	r, err := db.Table(table.Notice).Where("id=?", id).Data(Notice).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}

func DelNotice(id int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	r, err := db.Table(table.Notice).Where("id=?", id).Data(g.Map{"delete": 1}).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}
