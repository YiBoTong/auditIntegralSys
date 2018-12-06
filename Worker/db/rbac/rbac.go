package db_rbac

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
)

func Get(key string, menuParentId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.MenuTbName + " m")
	sql.InnerJoin(config.RbacTbName+" r", "r.menu_id=m.id")
	sql.Fields("m.*,r.key,r.is_read,r.is_write")
	sql.Where("m.delete=?", 0)
	sql.And("m.parent_id=?", menuParentId)
	sql.And("r.key=?", key)
	sql.OrderBy("m.order asc")
	res, err := sql.All()
	return res.ToList(), err
}