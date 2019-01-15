package db_rbac

import (
	"auditIntegralSys/_public/table"
	"fmt"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/util/gconv"
)

func Get(key string, menuParentId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(table.Menu + " m")
	sql.LeftJoin(table.Rbac+" r", "r.key='"+key+"' AND r.menu_id=m.id")
	sql.Where("m.delete=?", 0)
	sql.And("m.is_use=?", 1)
	sql.And("m.parent_id=?", menuParentId)
	res, err := sql.OrderBy("m.order asc").All()
	return res.ToList(), err
}

func Del(key string) (int, error) {
	var row int64 = 0
	db := g.DB()
	res, err := db.Table(table.Rbac).Where("`key`=?", key).Delete()
	if err == nil {
		row, _ = res.RowsAffected()
	}
	return int(row), err
}

func Add(data []g.Map) (int, error) {
	var row int64 = 0
	db := g.DB()
	res, err := db.BatchInsert(table.Rbac, data, 5)
	if err == nil {
		row, _ = res.RowsAffected()
	}
	return int(row), err
}

func RemoveOldData()  {
	db := g.DB()
	tableArr := []string{
		"rbac", "menu", "dictionary_type", "dictionary", "users", "user_job", "login", "logs", "department", "department_user", "notice", "notice_file", "notice_inform", "clause", "clause_content", "clause_file", "files", "programme", "programme_basis", "programme_content", "programme_step", "programme_business", "programme_emphases", "programme_user", "programme_examine_dep", "programme_examine_admin", "draft", "draft_content", "draft_admin_user", "draft_inspect_user", "draft_query_user", "draft_review_user", "draft_file", "confirmation", "confirmation_basis", "confirmation_receipt", "confirmation_receipt_content", "punish_notice", "punish_notice_behavior", "punish_notice_score", "integral", "integral_edit", "rectify", "rectify_report",
	}
	for _, v := range tableArr {
		n := gconv.String(v)
		r, _ := db.Table(n).Where("`delete`=?", 1).Delete()
		row, _ := r.RowsAffected()
		fmt.Println("删除数据表：", n, " --> ", row)
	}
}