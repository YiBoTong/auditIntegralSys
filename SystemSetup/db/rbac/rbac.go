package db_rbac

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
)

func Get(key string, menuParentId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.MenuTbName + " m")
	sql.LeftJoin(config.RbacTbName+" r", "r.key='"+key+"' AND r.menu_id=m.id")
	sql.Where("m.delete=?", 0)
	sql.And("m.parent_id=?", menuParentId)
	res, err := sql.All()
	return res.ToList(), err
}

func Del(key string) (int, error) {
	var row int64 = 0
	db := g.DB()
	res, err := db.Table(config.RbacTbName).Where("`key`=?", key).Delete()
	if err == nil {
		row, _ = res.RowsAffected()
	}
	return int(row), err
}

func Add(data []g.Map) (int, error) {
	var row int64 = 0
	db := g.DB()
	res, err := db.BatchInsert(config.RbacTbName, data, 5)
	if err == nil {
		row, _ = res.RowsAffected()
	}
	return int(row), err
}
