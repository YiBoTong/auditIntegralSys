package db_clause

import (
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
)

func GetClauseContents(clauseId int, offset int, limit int, where g.Map) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(table.ClauseContent).Where("`delete`=?", 0)
	sql.And("clause_id=?", clauseId)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Limit(offset, limit).OrderBy("`order` asc").Select()
	return r.ToList(), err
}

func SearchClauseContents(contentStr string) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.ClauseContent).Where("`delete`=?", 0)
	sql.And("content like ?", "%"+contentStr+"%")
	r, err := sql.Limit(0, 20).OrderBy("id asc").Select()
	return r.ToList(), err
}

func GetClauseContent(id int) (entity.ClauseContent, error) {
	var ClauseContent entity.ClauseContent
	db := g.DB()
	r, err := db.Table(table.ClauseContent).Where("id=?", id).And("`delete`=?", 0).One()
	_ = r.ToStruct(&ClauseContent)
	return ClauseContent, err
}

func AddClauseContent(ClauseContent g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	r, err := db.Insert(table.ClauseContent, ClauseContent)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func AddClauseContents(ClauseContentList []g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	r, err := db.BatchInsert(table.ClauseContent, ClauseContentList, 5)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func UpdateClauseContent(id int, ClauseContent g.Map) (int, error) {
	var rows int64 = 0
	db := g.DB()
	r, err := db.Table(table.ClauseContent).Where("id=?", id).Data(ClauseContent).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}

func UpdateClauseContents(ClauseContentList []g.Map) (int, error) {
	var rows int64 = 0
	db := g.DB()
	r, err := db.BatchReplace(table.ClauseContent, ClauseContentList, 5)
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}

func DelClauseContent(id int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	r, err := db.Table(table.ClauseContent).Where("id=?", id).Data(g.Map{"delete": 1}).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}

func DelClauseContentByClauseId(ClauseId int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	r, err := db.Table(table.ClauseContent).Where("clause_id=?", ClauseId).Data(g.Map{"delete": 1}).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}
