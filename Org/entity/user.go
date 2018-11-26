package entity

type User struct {
	UserId       int    `db:"user_id" json:"userId" field:"user_id"`
	DepartmentId int    `db:"department_id" json:"departmentId" field:"department_id"`
	UserName     string `db:"user_name" json:"userName" field:"user_name"`
	UserCode     int    `db:"user_code" json:"userCode" field:"user_code"`
	Class        string `db:"class" json:"class" field:"class"`
	Phone        string `db:"phone" json:"phone" field:"phone"`
	IdCard       string `db:"id_card" json:"idCard" field:"id_card"`
	UpdateTime   string `db:"update_time" json:"updateTime" field:"update_time"`
}
