package db_rbac

import (
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
)

func GetRbacMenu(key g.Slice, menuParentId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(table.Menu)
	sql.Where("`delete`=?", 0)
	sql.And("parent_id=?", menuParentId)
	sql.And("is_use=?", 1)
	sql.OrderBy("`order` asc")
	res, err := sql.All()
	return res.ToList(), err
}

func GetUserRbacByKeys(userKeys g.Slice) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.Rbac)
	sql.Where("`key` IN (?)", userKeys)
	sql.And("(is_read=1 OR is_write=1)")
	sql.And("`delete`=0")
	sql.OrderBy("is_write DESC")
	r, err := sql.All()
	return r.ToList(), err
}
