package db_dictionaries

import (
	"auditIntegralSys/_public/config"
	"database/sql/driver"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

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
	if err == nil && len(updateIds) > 0 {
		_, err = delDictionaries(ctx, typeId, updateIds)
	}
	if err == nil && len(add) > 0 {
		_, err = addDictionaries(ctx, add)
	}
	if err == nil && len(update) > 0 {
		_, err = updateDictionaries(ctx, update)
	}
	if err == nil {
		err = ctx.Commit()
	} else {
		err = ctx.Rollback()
	}
	return err == nil, err
}

func addDictionaries(ctx *gdb.TX, add []g.Map) (int, error) {
	var rows int64 = 0
	// 批次5条数据写入
	r, err := ctx.BatchInsert(config.DictionaryTbName, add, 5)
	if err == nil {
		rows, err = r.RowsAffected()
	}
	return int(rows), err
}

func updateDictionaries(ctx *gdb.TX, update []g.Map) (int, error) {
	var rows int64 = 0
	// 批次5条数据写入
	r, err := ctx.BatchReplace(config.DictionaryTbName, update, 5)
	if err == nil {
		rows, err = r.RowsAffected()
	}
	return int(rows), err
}

func delDictionaries(ctx *gdb.TX, typeId int, ids []int) (driver.Result, error) {
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
