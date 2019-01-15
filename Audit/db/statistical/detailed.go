package db_statistical

import (
	"auditIntegralSys/Audit/fun"
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func getDetailedListSql(db gdb.DB, authorInfo entity.User, where g.Map) *gdb.Model {
	sql := db.Table(table.Integral + " i")
	// 底稿排除人员
	sql.LeftJoin(table.Draft+" d", "i.draft_id=d.id")
	// 人员信息
	sql.LeftJoin(table.User+" u", "i.user_id=u.user_id")
	sql.LeftJoin(table.User+" cu", "i.cognizance_user_id=cu.user_id")
	// 部门相关信息
	sql.LeftJoin(table.Department+" ud", "u.department_id=ud.id")
	sql.LeftJoin(table.Department+" cud", "cu.department_id=cud.id")
	// 惩罚通知书获取依据信息
	sql.LeftJoin(table.PunishNotice+" pn", "i.punish_notice_id=pn.id")
	sql.LeftJoin(table.Clause+" c", "pn.basis_clause_id=c.id")
	// 在被检查单位担任职务
	sql.LeftJoin(table.DepartmentUser+" du", "du.department_id=u.department_id AND du.user_id=i.user_id")
	sql.LeftJoin(table.Dictionary+" dd", "du.type=dd.key AND dd.type_id=-2")

	sql.Where("i.delete=?", 0)

	sql.GroupBy("i.id")

	sql = fun.CheckIsMyData(*sql, authorInfo, where)

	// 姓名模糊查询
	if where["user_name"] != nil && where["user_name"] != "" {
		sql.And("u.user_name like ?", strings.Replace("%?%", "?", gconv.String(where["user_name"]), 1))
		delete(where, "user_name")
	}

	if len(where) > 0 {
		sql.And(where)
	}
	return sql
}

func DetailedCount(userInfo entity.User, where g.Map) (int, error) {
	db := g.DB()
	r, err := getDetailedListSql(db, userInfo, where).Count()
	return r, err
}

// 审计报告上报了的才能出现统计
func DetailedList(userInfo entity.User, offset int, limit int, where g.Map) (g.List, error) {
	db := g.DB()
	sql := getDetailedListSql(db, userInfo, where)
	fields := []string{
		"i.*",
		"u.user_name",
		"ud.name as department_name",
		"cud.name as punish_department_name",
		"c.title,c.number",
		"dd.value as role",
	}
	sql.Fields(strings.Join(fields, ","))
	r, err := sql.Limit(offset, limit).OrderBy("i.id desc").Select()
	return r.ToList(), err
}

func DetailedSum(userInfo entity.User, where g.Map) (g.List, error) {
	db := g.DB()
	sql := getDetailedListSql(db, userInfo, where)
	fields := []string{
		"i.id",
		"SUM(i.score) AS `sum_score`",
		"SUM(i.money) AS `sum_money`",
	}
	sql.Fields(strings.Join(fields, ","))
	r, err := sql.All()
	return r.ToList(), err
}
