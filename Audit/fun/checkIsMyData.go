package fun

import (
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func CheckIsMyData(sql gdb.Model, authorInfo entity.User, where g.Map) *gdb.Model {
	sql.LeftJoin(table.DraftAdminUser+" dau", "dau.draft_id=d.id")
	sql.LeftJoin(table.DraftReviewUser+" dru", "dru.draft_id=d.id")
	sql.LeftJoin(table.DraftAdminUser+" dqu", "dqu.draft_id=d.id")

	departmentId := authorInfo.DepartmentId

	// 项目名称模糊查询
	if where["project_name"] != nil && where["project_name"] != "" {
		sql.And("d.project_name like ?", strings.Replace("%?%", "?", gconv.String(where["project_name"]), 1))
		delete(where, "project_name")
	}

	// 项目名称模糊查询
	if where["query_department_id"] != nil && where["query_department_id"] != 0 {
		departmentId = gconv.Int(where["query_department_id"])
		delete(where, "query_department_id")
	}

	sql.And(
		"("+
		// 查询自己
			"d.author_id=?"+
			" OR "+
		// 别人已发布并且被检查部门是自己的部门的数据
			"(d.author_id!=? AND d.query_department_id=? AND d.public=? AND d.state=?)"+
			" OR "+
		// 查询自己是参与者的数据
			"(d.author_id!=? AND (dau.user_id=? OR dru.user_id=? OR dqu.user_id=?))"+
			")",
		authorInfo.UserId,
		authorInfo.UserId, departmentId, 1, state.Publish,
		authorInfo.UserId, authorInfo.UserId, authorInfo.UserId, authorInfo.UserId,
	)
	return &sql
}
