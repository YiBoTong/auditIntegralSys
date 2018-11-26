package check

import (
	"auditIntegralSys/Org/db/department"
	"auditIntegralSys/_public/config"
)

func HasDepartment(departmentId int) (bool, string, error) {
	msg := ""
	hasDepartment, err := db_department.HasDepartment(departmentId)
	if err == nil && !hasDepartment {
		msg = config.DepartmentMsgStr + config.NoHad
	}
	return hasDepartment, msg, err
}