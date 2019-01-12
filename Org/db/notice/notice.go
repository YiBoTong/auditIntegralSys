package db_notice

import (
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func getListSql(db gdb.DB, authorId int, where g.Map) *gdb.Model {
	sql := db.Table(table.Notice).Where("`delete`=?", 0)
	// 部门数据（包含全部范围数据）
	if where["department_id"] != nil && where["department_id"] != 0 {
		sql.And("(department_id=? OR `range`=?)", where["department_id"], 1)
		delete(where, "department_id")
	} else {
		sql.And("`range`=?", 1)
	}
	// 标题模糊查询
	if where["title"] != nil && where["title"] != "" {
		sql.And("title like ?", strings.Replace("%?%", "?", gconv.String(where["title"]), 1))
		delete(where, "title")
	}
	// 查询自己和别人已发布的数据
	sql.And("(author_id=? OR (author_id!=? AND state=?))", authorId, authorId, state.Publish)
	if len(where) > 0 {
		sql.And(where)
	}
	return sql
}

func GetNoticeCount(authorId int, where g.Map) (int, error) {
	db := g.DB()
	r, err := getListSql(db, authorId, where).Count()
	return r, err
}

func GetNotices(authorId, offset, limit int, where g.Map) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := getListSql(db, authorId, where)
	r, err := sql.Limit(offset, limit).OrderBy("id desc").Select()
	return r.ToList(), err
}

func GetNotice(id int) (entity.Notice, error) {
	var Notice entity.Notice
	db := g.DB()
	sql := db.Table(table.Notice + " n")
	sql.LeftJoin(table.Department+" d", "n.department_id=d.id")
	sql.Fields("n.*,d.name as department_name")
	sql.Where("n.id=?", id).And("`delete`=?", 0)
	r, err := sql.One()
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
