package db_department

import (
	"auditIntegralSys/_public/config"
	"database/sql/driver"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func AddDepartmentUser(users []g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	// 批次5条数据写入
	r, err := db.BatchInsert(config.DepartmentUserTbName, users, 5)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func GetDepartmentUser(departmentId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.DepartmentUserTbName + " du")
	sql.InnerJoin(config.UserTbName+" u", "du.user_id=u.user_id")
	sql.Fields("du.*,u.user_name")
	sql.Where("du.department_id=?", departmentId)
	sql.And("u.delete=?", 0)
	sql.And("du.delete=?", 0)
	r, err := sql.OrderBy("du.id desc").All()
	return r.ToList(), err
}

func UpdateDepartmentUser(departmentId int, add []g.Map, update []g.Map, updateIds []int) (bool, error) {
	db := g.DB()
	ctx, err := db.Begin()
	if err == nil && len(updateIds) > 0 {
		_, err = delDepartmentUser(ctx, departmentId, updateIds)
	}
	if err == nil && len(add) > 0 {
		_, err = addDepartmentUser(ctx, add)
	}
	if err == nil && len(update) > 0 {
		_, err = updateDepartmentUser(ctx, update)
	}
	if err == nil {
		err = ctx.Commit()
	} else {
		err = ctx.Rollback()
	}
	return err == nil, err
}

func addDepartmentUser(ctx *gdb.TX, add []g.Map) (int, error) {
	var rows int64 = 0
	// 批次5条数据写入
	r, err := ctx.BatchInsert(config.DepartmentUserTbName, add, 5)
	if err == nil {
		rows, err = r.RowsAffected()
	}
	return int(rows), err
}

func updateDepartmentUser(ctx *gdb.TX, update []g.Map) (int, error) {
	var rows int64 = 0
	// 批次5条数据写入
	r, err := ctx.BatchReplace(config.DepartmentUserTbName, update, 5)
	if err == nil {
		rows, err = r.RowsAffected()
	}
	return int(rows), err
}

func delDepartmentUser(ctx *gdb.TX, departmentId int, ids []int) (driver.Result, error) {
	var sql *gdb.Model
	if len(ids) > 1 {
		for index, id := range ids {
			if index == 0 {
				sql = ctx.Table(config.DepartmentUserTbName).Where("id=?", id)
			} else {
				sql.Or("id=?", id)
			}
		}
	}
	if len(ids) == 1 {
		sql = ctx.Table(config.DepartmentUserTbName).Where("id=?", ids[0])
	}
	// 先全部软删除
	r, err := ctx.Table(config.DepartmentUserTbName).Where("department_id=?", departmentId).Data(g.Map{"delete": 1}).Update()
	if len(ids) > 0 {
		// 再还原保留的数据
		r, err = sql.Data(g.Map{"delete": 0}).Update()
	}
	return r, err
}
