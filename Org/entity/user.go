package entity

import "auditIntegralSys/Worker/entity"

type User struct {
	UserId         int    `db:"user_id" json:"userId" field:"user_id"`
	DepartmentId   int    `db:"department_id" json:"departmentId" field:"department_id"`
	DepartmentName string `db:"department_name" json:"departmentName" field:"department_name"`
	UserName       string `db:"user_name" json:"userName" field:"user_name"`
	UserCode       string `db:"user_code" json:"userCode" field:"user_code"`
	Sex            int    `db:"sex" json:"sex" field:"sex"`
	Class          string `db:"class" json:"class" field:"class"`
	Phone          string `db:"phone" json:"phone" field:"phone"`
	PortraitId     int    `db:"portrait_id" json:"portraitId" field:"portrait_id"`
	IdCard         string `db:"id_card" json:"idCard" field:"id_card"`
	UpdateTime     string `db:"update_time" json:"updateTime" field:"update_time"`
}

type ImportUserRes struct {
	Id          int    `db:"id" json:"id" field:"id"`
	HadUserList []User `db:"had_user_list" json:"hadUserList" field:"had_user_list"`
}

type LoginUserDepartmentItem struct {
	DepartmentId   int    `db:"department_id" json:"departmentId" field:"department_id"`
	DepartmentName string `db:"department_name" json:"departmentName" field:"department_name"`
	Type           string `db:"type" json:"type" field:"type"`
}

type LoginUserInfo struct {
	User
	RbacList     []entity.RbacListItem     `db:"rbac_list" json:"rbacList" field:"rbac_list"`
	Departments  []LoginUserDepartmentItem `db:"departments" json:"departments" field:"departments"`
	PortraitFile entity.File               `db:"portrait_file" json:"portraitFile" field:"portrait_file"`
}
