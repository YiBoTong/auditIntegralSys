package entity

import "auditIntegralSys/Worker/entity"

type NoticeList struct {
	Id           int    `db:"id" json:"id" field:"id"`
	DepartmentId int    `db:"department_id" json:"departmentId" field:"department_id"`
	Title        string `db:"title" json:"title" field:"title"`
	Time         string `db:"time" json:"time" field:"time"`
	Range        int    `db:"range" json:"range" field:"range"`
	State        string `db:"state" json:"state" field:"state"`
}

type Notice struct {
	Id           int    `db:"id" json:"id" field:"id"`
	DepartmentId int    `db:"department_id" json:"departmentId" field:"department_id"`
	Title        string `db:"title" json:"title" field:"title"`
	Content      string `db:"content" json:"content" field:"content"`
	Time         string `db:"time" json:"time" field:"time"`
	Range        int    `db:"range" json:"range" field:"range"`
	State        string `db:"state" json:"state" field:"state"`
}

type NoticeRes struct {
	Notice
	FileList []entity.File        `db:"file_list" json:"fileList" field:"file_list"`
	Informs  []DepartmentTreeInfo `db:"informs" json:"informs" field:"informs"`
}
