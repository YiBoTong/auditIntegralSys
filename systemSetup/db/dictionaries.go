package db

import (
	"auditIntegralSys/_public/config"
	"auditIntegralSys/systemSetup/entity"
	"gitee.com/johng/gf/g"
)

func GetDictionaryTypeCount(where g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(config.DictionaryTypeTbName).Where("'delete'=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Count()
	return r, err
}

func GetDictionaryTypes(offset int, limit int, where g.Map) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.DictionaryTypeTbName + " d").LeftJoin(config.UserTbName+" u", "d.user_id=u.id").Fields("d.*,u.user_name").Where("'d.delete'=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Limit(offset, limit).OrderBy("id desc").Select()
	return r.ToList(), err
}

func GetDictionaryType(id int) (entity.DictionaryType, error) {
	var dictionaryType entity.DictionaryType
	db := g.DB()
	r, err := db.Table(config.DictionaryTypeTbName + " d").LeftJoin(config.UserTbName+" u", "d.user_id=u.id").Fields("d.*,u.user_name").Where("d.id=?", id).And("d.delete=?", 0).One()
	r.ToStruct(&dictionaryType)
	return dictionaryType, err
}

func AddDictionaryType(dictionaryType g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	r, err := db.Insert(config.DictionaryTypeTbName, dictionaryType)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func UpdateDictionaryType(id int, dictionaryType g.Map) (int, error) {
	var rows int64 = 0
	db := g.DB()
	r, err := db.Table(config.DictionaryTypeTbName).Where("id=?", id).Data(dictionaryType).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}

func DelDictionaryType(typeId int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	r, err := db.Table(config.DictionaryTypeTbName).Where("id=?", typeId).Data(g.Map{"delete": 1}).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}

func AddDictionaries(dictionaries []g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	// 批次5条数据写入
	r, err := db.BatchInsert(config.DictionaryTbName, dictionaries, 5)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), nil
}

func GetDictionaries(typeId int) ([]map[string]interface{}, error) {
	db := g.DB()
	r, err := db.Table(config.DictionaryTbName).Where("type_id=?", typeId).And("'delete'=?", 0).OrderBy("'order' desc").All()
	return r.ToList(), err
}

func DelDictionary() {
}
