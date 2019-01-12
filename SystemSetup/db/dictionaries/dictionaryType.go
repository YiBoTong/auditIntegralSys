package db_dictionaries

import (
	"auditIntegralSys/SystemSetup/entity"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func getListSql(db gdb.DB, where g.Map) *gdb.Model {
	sql := db.Table(table.DictionaryType + " d")
	sql.LeftJoin(table.User+" u", "d.user_id=u.user_id")
	sql.Where("d.delete=?", 0)
	// 项目名称模糊查询
	if where["title"] != nil && where["title"] != "" {
		sql.And("d.title like ?", strings.Replace("%?%", "?", gconv.String(where["title"]), 1))
		delete(where, "title")
	}
	if len(where) > 0 {
		sql.And(where)
	}
	return sql
}

func GetDictionaryTypeCount(where g.Map) (int, error) {
	db := g.DB()
	sql := getListSql(db, where)
	r, err := sql.Count()
	return r, err
}

func GetDictionaryTypes(offset int, limit int, where g.Map) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := getListSql(db, where)
	sql.Fields("d.*,u.user_name")
	r, err := sql.Limit(offset, limit).OrderBy("d.id desc").Select()
	return r.ToList(), err
}

func GetDictionaryType(id int) (entity.DictionaryType, error) {
	var dictionaryType entity.DictionaryType
	db := g.DB()
	r, err := db.Table(table.DictionaryType + " d").LeftJoin(table.User+" u", "d.user_id=u.user_id").Fields("d.*,u.user_name").Where("d.id=?", id).And("d.delete=?", 0).One()
	_ = r.ToStruct(&dictionaryType)
	return dictionaryType, err
}

func AddDictionaryType(dictionaryType g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	r, err := db.Insert(table.DictionaryType, dictionaryType)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func UpdateDictionaryType(id int, dictionaryType g.Map) (int, error) {
	var rows int64 = 0
	db := g.DB()
	r, err := db.Table(table.DictionaryType).Where("id=?", id).Data(dictionaryType).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}

func DelDictionaryType(typeId int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	r, err := db.Table(table.DictionaryType).Where("id=?", typeId).Data(g.Map{"delete": 1}).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}
