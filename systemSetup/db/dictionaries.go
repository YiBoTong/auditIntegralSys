package db

import (
	"auditIntegralSys/_public/config"
	"auditIntegralSys/systemSetup/entity"
	"database/sql/driver"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func GetDictionaryTypeCount(where g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(config.DictionaryTypeTbName).Where("`delete`=?", 0)
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
	return int(lastId), err
}

func GetDictionaries(typeId int) ([]map[string]interface{}, error) {
	db := g.DB()
	r, err := db.Table(config.DictionaryTbName).Where("type_id=?", typeId).And("`delete`=?", 0).OrderBy("`order` asc").All()
	return r.ToList(), err
}

func UpdateDictionaries(typeId int, add []g.Map, update []g.Map, updateIds []int) (bool, error) {
	db := g.DB()
	ctx, err := db.Begin()
	if err == nil {
		_, err = delDictionaries(ctx, typeId, updateIds)
	}
	if err == nil {
		_, err = addDictionaries(ctx, add)
	}
	if err == nil {
		_, err = updateDictionaries(ctx, update)
	}
	if err == nil {
		ctx.Commit()
	} else {
		ctx.Rollback()
	}
	return err == nil, err
}

func addDictionaries(ctx *gdb.Tx, add []g.Map) (int, error) {
	var rows int64 = 0
	// 批次5条数据写入
	r, err := ctx.BatchInsert(config.DictionaryTbName, add, 5)
	if err == nil {
		rows, err = r.RowsAffected()
	}
	return int(rows), err
}

func updateDictionaries(ctx *gdb.Tx, update []g.Map) (int, error) {
	var rows int64 = 0
	// 批次5条数据写入
	r, err := ctx.BatchReplace(config.DictionaryTbName, update, 5)
	if err == nil {
		rows, err = r.RowsAffected()
	}
	return int(rows), err
}

func delDictionaries(ctx *gdb.Tx, typeId int, ids []int) (driver.Result, error) {
	var sql *gdb.Model
	if len(ids) > 1 {
		for index, id := range ids {
			if index == 0 {
				sql = ctx.Table(config.DictionaryTbName).Where("id=?", id)
			} else {
				sql.Or("id=?", id)
			}
		}
	}
	if len(ids) == 1 {
		sql = ctx.Table(config.DictionaryTbName).Where("id=?", ids[0])
	}
	// 先全部软删除
	r, err := ctx.Table(config.DictionaryTbName).Where("type_id=?", typeId).Data(g.Map{"delete": 1}).Update()
	if len(ids) > 0 {
		// 再还原保留的数据
		r, err = sql.Data(g.Map{"delete": 0}).Update()
	}
	return r, err
}
