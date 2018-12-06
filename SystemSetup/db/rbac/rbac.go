package db_rbac

import (
	"auditIntegralSys/_public/config"
	"fmt"
	"gitee.com/johng/gf/g"
)

func Get(key string, menuParentId int) ([]map[string]interface{}, error) {
	db := g.DB()
	str := fmt.Sprintf("SELECT m.*,r.* FROM %v m LEFT JOIN %v r ON (r.menu_id=m.id) WHERE m.delete=0 AND m.parent_id=%v", config.MenuTbName, config.RbacTbName, menuParentId)
	sql := str
	sql += " UNION "
	sql += fmt.Sprintf("%v AND r.key='%v'", str, key)
	sql += " ORDER BY `order` ASC"
	res, err := db.GetAll(sql)
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
