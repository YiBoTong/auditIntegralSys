package entity

type Department struct {
	Id         int    `db:"id" json:"id" field:"id"`
	Name       string `db:"name" json:"name" field:"name"`
	ParentId   int    `db:"parent_id" json:"parentId" field:"parent_id"`
	Code       string `db:"code" json:"code" field:"code"`
	Level      int    `db:"level" json:"level" field:"level"`
	Address    string `db:"address" json:"address" field:"address"`
	Phone      string `db:"phone" json:"phone" field:"phone"`
	UpdateTime string `db:"update_time" json:"updateTime" field:"update_time"`
}

type DepartmentTreeInfo struct {
	Id       int    `db:"id" json:"id" field:"id"`
	Name     string `db:"name" json:"name" field:"name"`
	ParentId int    `db:"parent_id" json:"parentId" field:"parent_id"`
	Code     string `db:"code" json:"code" field:"code"`
	Level    int    `db:"level" json:"level" field:"level"`
}

type AddDepartment struct {
	Name       string `db:"name" json:"name" field:"name"`
	ParentId   int    `db:"parent_id" json:"parentId" field:"parent_id"`
	Code       string `db:"code" json:"code" field:"code"`
	Level      int    `db:"level" json:"level" field:"level"`
	Address    string `db:"address" json:"address" field:"address"`
	Phone      string `db:"phone" json:"phone" field:"phone"`
	UpdateTime string `db:"update_time" json:"updateTime" field:"update_time"`
}

type DepartmentRes struct {
	Department
	UserList []DepUser `db:"user_list" json:"userList" field:"user_list"`
}

type DepUser struct {
	Id       int    `db:"id" json:"id" field:"id"`
	UserId   int    `db:"user_id" json:"userId" field:"user_id"`
	UserName string `db:"user_name" json:"userName" field:"user_name"`
	Type     string `db:"type" json:"type" field:"type"`
	TypeName string `db:"type_name" json:"typeName" field:"type_name"`
}
