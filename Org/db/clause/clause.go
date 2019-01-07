package db_clause

import (
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/Worker/db/file"
	"auditIntegralSys/_public/table"
	"database/sql"
	"gitee.com/johng/gf/g"
)

func GetClauseCount(where g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(table.Clause).Where("`delete`=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Count()
	return r, err
}

func GetClauses(offset int, limit int, where g.Map) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.Clause + " c")
	sql.LeftJoin(table.User+" u", "c.author_id=u.user_id")
	sql.Fields("c.*,u.user_name as author_name")
	sql.Where("c.delete=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Limit(offset, limit).OrderBy("id desc").Select()
	return r.ToList(), err
}

func GetClauseTitle(offset int, limit int, departmentId int, title string) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.Clause)
	if departmentId != 0 {
		sql.Where("`delete`=? AND (department_id=? OR department_id=-1)", 0, departmentId)
	} else {
		sql.Where("`delete`=? AND department_id=-1", 0)
	}
	sql.And("state=?", "publish")
	sql.And("title like ?", "%"+title+"%")
	r, err := sql.Limit(offset, limit).OrderBy("id desc").Select()
	return r.ToList(), err
}

func GetClause(id int) (entity.Clause, error) {
	var Clause entity.Clause
	db := g.DB()
	sql := db.Table(table.Clause + " c")
	sql.LeftJoin(table.User+" u", "c.author_id=u.user_id")
	sql.Fields("c.*,u.user_name as author_name")
	sql.Where("c.delete=?", 0)
	sql.And("c.id=?", id)
	r, err := sql.One()
	_ = r.ToStruct(&Clause)
	return Clause, err
}

func AddByTX(Clause g.Map, content g.List, fileId int) (int, error) {
	db := g.DB()
	tx, err := db.Begin()
	id := 0
	if err == nil {
		var res sql.Result
		res, err = tx.Insert(table.Clause, Clause)
		addId, _ := res.LastInsertId()
		id = int(addId)
	}
	if err == nil {
		for index, v := range content {
			v["clause_id"] = id
			v["order"] = index
		}
		_, err = AddClauseContentsTX(*tx, content)
	}
	if err == nil {
		_, err = db_file.UpdateFile(fileId, g.Map{"form_id": id, "form": table.Clause}, tx)
	}
	if err == nil {
		_ = tx.Commit()
	} else {
		_ = tx.Rollback()
	}
	return int(id), err
}

func AddClause(Clause g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	r, err := db.Insert(table.Clause, Clause)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func UpdateClause(id int, Clause g.Map) (int, error) {
	var rows int64 = 0
	db := g.DB()
	r, err := db.Table(table.Clause).Where("id=?", id).Data(Clause).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}

func DelClause(id int) (int, error) {
	db := g.DB()
	rows := 0
	tx, err := db.Begin()
	if err == nil {
		var res sql.Result
		res, err = tx.Table(table.Clause).Where("id=?", id).Data(g.Map{"delete": 1}).Update()
		row, _ := res.RowsAffected()
		rows = int(row)
	}
	if err == nil && rows != 0 {
		_, err = DelClauseContentByTx(*tx, id)
	}
	if err == nil {
		_ = tx.Commit()
	} else {
		_ = tx.Rollback()
	}
	return rows, err
}
