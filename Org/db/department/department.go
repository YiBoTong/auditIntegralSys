package db_department

import (
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
)

func GetDepartmentsByParentId(parentId int, search g.Map) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(table.Department).Where("`delete`=?", 0)
	sql.And("parent_id=?", parentId)
	if len(search) > 0 {
		sql.And(search)
	}
	sql.OrderBy("id desc")
	r, err := sql.All()
	return r.ToList(), err
}

func HasDepartment(departmentId int) (bool, error) {
	db := g.DB()
	sql := db.Table(table.Department).Where("`delete`=?", 0).And("id=?", departmentId)
	count, err := sql.Count()
	return count > 0, err
}

func HadChildByParentId(parintId int) (int, error) {
	db := g.DB()
	sql := db.Table(table.Department).Where("`delete`=?", 0).And("parent_id=?", parintId)
	count, err := sql.Count()
	return count, err
}

func AddDepartment(department g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	r, err := db.Insert(table.Department, department)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func GetDepartment(id int) (entity.Department, error) {
	var department entity.Department
	db := g.DB()
	sql := db.Table(table.Department+" d")
	sql.LeftJoin(table.Department+" dd","d.parent_id=dd.id")
	sql.Fields("d.*,dd.name as parent_dep_name")
	sql.Where("d.id=?", id)
	sql.And("d.delete=?", 0)
	r, err := sql.One()
	_ = r.ToStruct(&department)
	return department, err
}

func UpdateDepartment(id int, department g.Map) (int, error) {
	var rows int64 = 0
	db := g.DB()
	r, err := db.Table(table.Department).Where("id=?", id).Data(department).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}

func DelDepartment(departmentId int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	r, err := db.Table(table.Department).Where("id=?", departmentId).Data(g.Map{"delete": 1}).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}
