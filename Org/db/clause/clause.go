package db_clause

import (
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/Worker/db/file"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/table"
	"database/sql"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func getListSql(db gdb.DB, authorId int, where g.Map) *gdb.Model {
	sql := db.Table(table.Clause + " c")
	sql.LeftJoin(table.User+" u", "c.author_id=u.user_id")
	sql.LeftJoin(table.Department+" d", "c.department_id=d.id")
	sql.Where("c.delete=?", 0)
	// 部门数据（包含全部范围数据）
	if where["department_id"] != nil && where["department_id"] != 0 {
		sql.And("(c.department_id=? OR c.department_id=?)", where["department_id"], -1)
		delete(where, "department_id")
	} else {
		sql.And("c.department_id=?", -1)
	}
	// 标题模糊查询
	if where["title"] != nil && where["title"] != "" {
		sql.And("c.title like ?", strings.Replace("%?%", "?", gconv.String(where["title"]), 1))
		delete(where, "title")
	}
	// 查询自己和别人已发布的数据
	sql.And("(c.author_id=? OR (c.author_id!=? AND c.state=?))", authorId, authorId, state.Publish)
	if len(where) > 0 {
		sql.And(where)
	}
	return sql
}

func GetClauseCount(authorId int, where g.Map) (int, error) {
	db := g.DB()
	sql := getListSql(db, authorId, where)
	r, err := sql.Count()
	return r, err
}

func GetClauses(authorId, offset, limit int, where g.Map) (g.List, error) {
	db := g.DB()
	sql := getListSql(db, authorId, where)
	sql.Fields("c.*,u.user_name as author_name,d.name as department_name")
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
	sql.LeftJoin(table.Department+" d", "c.department_id=d.id")
	sql.Fields("c.*,u.user_name as author_name,d.name as department_name")
	sql.Where("c.delete=?", 0)
	sql.And("c.id=?", id)
	r, err := sql.One()
	_ = r.ToStruct(&Clause)
	return Clause, err
}

func AddByTX(Clause g.Map, content g.List, fileIds string) (int, error) {
	db := g.DB()
	tx, err := db.Begin()
	id := 0
	if err == nil {
		var res sql.Result
		res, err = tx.Insert(table.Clause, Clause)
		addId, _ := res.LastInsertId()
		id = int(addId)
	}
	if err == nil && len(content) > 0 {
		for _, v := range content {
			v["clause_id"] = id
		}
		_, err = AddClauseContentsTX(*tx, content)
	}
	if err == nil && len(fileIds) > 0 {
		_, err = db_file.UpdateFileByIds(table.Clause, fileIds, id, tx)
	}
	if err == nil {
		_ = tx.Commit()
	} else {
		_ = tx.Rollback()
	}
	return int(id), err
}

func ImportByTX(Clause g.Map, content g.List, fileId int) (int, error) {
	db := g.DB()
	tx, err := db.Begin()
	id := 0
	if err == nil {
		var res sql.Result
		res, err = tx.Insert(table.Clause, Clause)
		addId, _ := res.LastInsertId()
		id = int(addId)
	}
	if err == nil && len(content) > 0 {
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
