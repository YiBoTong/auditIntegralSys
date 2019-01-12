package db_introduction

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func getListSql(db gdb.DB, authorId int, where g.Map) *gdb.Model {
	sql := db.Table(table.Introduction + " i")
	sql.LeftJoin(table.Draft+" d", "i.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "d.programme_id=p.id")
	sql.LeftJoin(table.Department+" dd", "d.department_id=dd.id")
	sql.LeftJoin(table.Department+" qdd", "d.query_department_id=qdd.id")
	sql.Where("d.delete=?", 0)
	// 项目名称模糊查询
	if where["project_name"] != nil && where["project_name"] != "" {
		sql.And("d.project_name like ?", strings.Replace("%?%", "?", gconv.String(where["project_name"]), 1))
		delete(where, "project_name")
	}
	return sql
}

func GetCount(authorId int, where g.Map) (int, error) {
	db := g.DB()
	sql := getListSql(db, authorId, where)
	r, err := sql.Count()
	return r, err
}

func GetList(authorId, offset, limit int, where g.Map) (g.List, error) {
	db := g.DB()
	sql := getListSql(db, authorId, where)
	fields := []string{
		"d.*,i.id",
		"i.number",
		"p.title as programme_title",
		"dd.name as department_name",
		"qdd.name as query_department_name",
		"i.id as introduction_id",
	}
	sql.Fields(strings.Join(fields, ","))
	r, err := sql.Limit(offset, limit).OrderBy("i.id desc").Select()
	return r.ToList(), err
}

func Add(data g.Map) (int, error) {
	db := g.DB()
	res, err := db.Table(table.Introduction).Data(data).Insert()
	id, _ := res.LastInsertId()
	return int(id), err
}

func GetByDraftId(draftId int) (entity.IntroductionItem, error) {
	db := g.DB()
	res, err := db.Table(table.Introduction).Where("draft_id=? AND `delete`=?", draftId, 0).One()
	introductionItem := entity.IntroductionItem{}
	_ = res.ToStruct(&introductionItem)
	return introductionItem, err
}

func Get(id int) (entity.IntroductionItem, error) {
	db := g.DB()
	res, err := db.Table(table.Introduction).Where("id=? AND `delete`=?", id, 0).One()
	introductionItem := entity.IntroductionItem{}
	_ = res.ToStruct(&introductionItem)
	return introductionItem, err
}
