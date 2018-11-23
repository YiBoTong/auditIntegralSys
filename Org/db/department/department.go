package db_department

import (
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/_public/config"
	"database/sql/driver"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func GetDepartmentsByParentId(parentId int, search g.Map) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.DepartmentTbName).Where("`delete`=?", 0)
	sql.And("parent_id=?", parentId)
	if len(search) > 0 {
		sql.And(search)
	}
	sql.OrderBy("id desc")
	r, err := sql.All()
	return r.ToList(), err
}

func AddDepartment(department g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	r, err := db.Insert(config.DepartmentTbName, department)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func GetDepartment(id int) (entity.Department, error) {
	var department entity.Department
	db := g.DB()
	r, err := db.Table(config.DepartmentTbName).Where("id=?", id).And("`delete`=?", 0).One()
	_ = r.ToStruct(&department)
	return department, err
}

func UpdateDepartment(id int, department g.Map) (int, error) {
	var rows int64 = 0
	db := g.DB()
	r, err := db.Table(config.DepartmentTbName).Where("id=?", id).Data(department).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}

func DelDepartment(departmentId int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	r, err := db.Table(config.DepartmentTbName).Where("id=?", departmentId).Data(g.Map{"delete": 1}).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}

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
	sql.OrderBy("du.id asc")
	r, err := sql.All()
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

func addDepartmentUser(ctx *gdb.Tx, add []g.Map) (int, error) {
	var rows int64 = 0
	// 批次5条数据写入
	r, err := ctx.BatchInsert(config.DepartmentUserTbName, add, 5)
	if err == nil {
		rows, err = r.RowsAffected()
	}
	return int(rows), err
}

func updateDepartmentUser(ctx *gdb.Tx, update []g.Map) (int, error) {
	var rows int64 = 0
	// 批次5条数据写入
	r, err := ctx.BatchReplace(config.DepartmentUserTbName, update, 5)
	if err == nil {
		rows, err = r.RowsAffected()
	}
	return int(rows), err
}

func delDepartmentUser(ctx *gdb.Tx, departmentId int, ids []int) (driver.Result, error) {
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
