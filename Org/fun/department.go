package fun

import (
	"auditIntegralSys/Org/db/department"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/util/gconv"
)

func UpdateDepartmentHasChild(parentId int) error {
	hadChild, err := db_department.HadChildByParentId(parentId)
	hasChild := false
	if hadChild > 0 {
		hasChild = true
	}
	if err == nil {
		_, err = db_department.UpdateDepartment(parentId, g.Map{"has_child": gconv.Int(hasChild)})
	}
	return err
}

func UpdateDepartmentHasChildById(id int) error {
	hadChild :=0
	hasChild := false
	department,err := db_department.GetDepartment(id)
	if err == nil {
		hadChild, err = db_department.HadChildByParentId(department.ParentId)
		if hadChild > 1 {
			hasChild = true
		}
		if err == nil {
			_, err = db_department.UpdateDepartment(department.ParentId, g.Map{"has_child": gconv.Int(hasChild)})
		}
	}
	return err
}