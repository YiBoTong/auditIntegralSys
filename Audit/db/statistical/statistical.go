package db_statistical

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/Audit/fun"
	entity2 "auditIntegralSys/Org/entity"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"strings"
)

func getListSql(db gdb.DB, authorInfo entity2.User, where g.Map) *gdb.Model {
	sql := db.Table(table.AuditReport + " ar")
	sql.LeftJoin(table.Draft+" d", "ar.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "ar.programme_id=p.id")
	sql.LeftJoin(table.Department+" dd", "d.department_id=dd.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")
	sql.Where("ar.delete=? AND ar.state=?", 0, state.Publish)

	sql.GroupBy("ar.id")

	sql = fun.CheckIsMyData(*sql, authorInfo, where)

	if len(where) > 0 {
		sql.And(where)
	}
	return sql
}

func Count(userInfo entity2.User, where g.Map) (int, error) {
	db := g.DB()
	r, err := getListSql(db, userInfo, where).Count()
	return r, err
}

// 审计报告上报了的才能出现统计
func List(userInfo entity2.User, offset int, limit int, where g.Map) (g.List, error) {
	db := g.DB()
	sql := getListSql(db, userInfo, where)
	fields := []string{
		"d.*",
		"ar.*",
		"p.start_time",
		"p.end_time",
		"d.query_start_time,d.query_end_time",
		"dd.name as department_name",
		"dq.name as query_department_name",
		"p.title as programme_title",
	}
	sql.Fields(strings.Join(fields, ","))
	r, err := sql.Limit(offset, limit).OrderBy("ar.id desc").Select()
	return r.ToList(), err
}

func Get(id int, where ...g.Map) (entity.StatisticalListItem, error) {
	db := g.DB()
	fields := []string{
		"d.*",
		"ar.*",
		"p.start_time",
		"p.end_time",
		"d.query_start_time,d.query_end_time",
		"dd.name as department_name",
		"dq.name as query_department_name",
		"p.title as programme_title",
	}
	sql := db.Table(table.AuditReport + " ar")
	sql.LeftJoin(table.Draft+" d", "ar.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "ar.programme_id=p.id")
	sql.LeftJoin(table.Department+" dd", "d.department_id=dd.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")
	sql.Fields(strings.Join(fields, ","))
	sql.Where("ar.id=? AND ar.delete=? AND ar.state=?", id, 0, state.Publish)
	if len(where) > 0 {
		sql.And(where[0])
	}
	item := entity.StatisticalListItem{}
	r, err := sql.One()
	_ = r.ToStruct(&item)
	return item, err
}

func GetOneStatisticalUser(draftId int) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.Integral + " i")
	sql.LeftJoin(table.User+" u", "i.user_id=u.user_id")
	sql.LeftJoin(table.PunishNoticeBehavior+" pb", "i.punish_notice_id=pb.punish_notice_id")

	sql.Where("d.id=? AND d.delete=?", draftId, 0)


	fields := []string{
		"i.id",
		"SUM(i.score) AS `sum_score`",
		"SUM(i.money) AS `sum_money`",
	}
	sql.Fields(strings.Join(fields, ","))

	r, err := sql.All()
	return r.ToList(), err

}
