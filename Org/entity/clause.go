package entity

import "auditIntegralSys/Worker/entity"

type ClauseRes struct {
	Clause
	Content  []ClauseContent `db:"content" json:"content" field:"content"`
	FileList []entity.File   `db:"file_list" json:"fileList" field:"file_list"`
}

type Clause struct {
	Id             int    `db:"id" json:"id" field:"id"`
	DepartmentId   int    `db:"department_id" json:"departmentId" field:"department_id"`
	DepartmentName string `db:"department_name" json:"departmentName" field:"department_name"`
	Title          string `db:"title" json:"title" field:"title"`
	From           string `db:"from" json:"from" field:"from"`
	Type           string `db:"type" json:"type" field:"type"`
	AuthorId       int    `db:"author_id" json:"authorId" field:"author_id"`
	AuthorName     string `db:"author_name" json:"authorName" field:"author_name"`
	Number         string `db:"number" json:"number" field:"number"`
	UpdateTime     string `db:"update_time" json:"updateTime" field:"update_time"`
	State          string `db:"state" json:"state" field:"state"`
}

type ClauseTitle struct {
	Id             int    `db:"id" json:"id" field:"id"`
	DepartmentId   int    `db:"department_id" json:"departmentId" field:"department_id"`
	DepartmentName string `db:"department_name" json:"departmentName" field:"department_name"`
	Title          string `db:"title" json:"title" field:"title"`
	From           string `db:"from" json:"from" field:"from"`
	Type           string `db:"type" json:"type" field:"type"`
	Number         string `db:"number" json:"number" field:"number"`
}

type ClauseContent struct {
	Id         int    `db:"id" json:"id" field:"id"`
	ClauseId   int    `db:"clause_id" json:"clauseId" field:"clause_id"`
	IsTitle    int8   `db:"is_title" json:"isTitle" field:"is_title"`
	TitleLevel string `db:"title_level" json:"titleLevel" field:"title_level"`
	Content    string `db:"content" json:"content" field:"content"`
	Order      int    `db:"order" json:"order" field:"order"`
}
