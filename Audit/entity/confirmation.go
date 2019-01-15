package entity

import "auditIntegralSys/Worker/entity"

type ConfirmationListItem struct {
	Id                    int    `db:"id" json:"id" field:"id"`
	ProgrammeId           int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	ProgrammeTitle        string `db:"programme_title" json:"programmeTitle" field:"programme_title"`
	QueryDepartmentId     int    `db:"query_department_id" json:"queryDepartmentId" field:"query_department_id"`
	QueryDepartmentName   string `db:"query_department_name" json:"queryDepartmentName" field:"query_department_name"`
	DepartmentId          int    `db:"department_id" json:"departmentId" field:"department_id"`
	DepartmentName        string `db:"department_name" json:"departmentName" field:"department_name"`
	Year                  int    `db:"year" json:"year" field:"year"`
	Number                int    `db:"number" json:"number" field:"number"`
	ProjectName           string `db:"project_name" json:"projectName" field:"project_name"`
	Public                uint8  `db:"public" json:"public" field:"public"`
	QueryStartTime        string `db:"query_start_time" json:"queryStartTime" field:"query_start_time"`
	QueryEndTime          string `db:"query_end_time" json:"queryEndTime" field:"query_end_time"`
	UpdateTime            string `db:"update_time" json:"updateTime" field:"update_time"`
	State                 string `db:"state" json:"state" field:"state"`
	ConfirmationReceiptId int    `db:"confirmation_receipt_id" json:"confirmationReceiptId" field:"confirmation_receipt_id"` // 回执ID
	HasRead               uint8  `db:"has_read" json:"hasRead" field:"has_read"`                                             // 是否已读
	HasReadTime           string `db:"has_read_time" json:"hasReadTime" field:"has_read_time"`                               // 已读时间
	AuthorId              int    `db:"author_id" json:"authorId" field:"author_id"`
}

type ConfirmationItem struct {
	Id                    int    `db:"id" json:"id" field:"id"`
	DraftId               int    `db:"draft_id" json:"draftId" field:"draft_id"`                                             // 工作底稿ID
	ConfirmationReceiptId int    `db:"confirmation_receipt_id" json:"confirmationReceiptId" field:"confirmation_receipt_id"` // 回执ID
	HasRead               uint8  `db:"has_read" json:"hasRead" field:"has_read"`                                             // 是否已读
	HasReadTime           string `db:"has_read_time" json:"hasReadTime" field:"has_read_time"`                               // 已读时间
	State                 string `db:"state" json:"state" field:"state"`
	Year                  int    `db:"year" json:"year" field:"year"`
	Number                int    `db:"number" json:"number" field:"number"`
}

// 稽核发现的问题内容
type ConfirmationContent struct {
	Id              int    `db:"id" json:"id" field:"id"`
	DraftId         int    `db:"draft_id" json:"draftId" field:"draft_id"`
	Order           int    `db:"order" json:"order" field:"order"`
	Type            string `db:"type" json:"type" field:"type"`
	BehaviorId      int    `db:"behavior_id" json:"behaviorId" field:"behavior_id"`
	BehaviorContent string `db:"behavior_content" json:"behaviorContent" field:"behavior_content"`
}

// 被检查人
type ConfirmationUser struct {
	Id       int    `db:"id" json:"id" field:"id"`
	DraftId  int    `db:"draft_id" json:"draftId" field:"draft_id"`
	UserId   int    `db:"user_id" json:"userId" field:"user_id"`
	UserName string `db:"user_name" json:"userName" field:"user_name"`
}

type Confirmation struct {
	ConfirmationItem
	Draft       DraftItem             `db:"draft" json:"draft" field:"draft"`               // 工作底稿
	Programme   ProgrammeItem         `db:"programme" json:"programme" field:"programme"`   // 方案
	BasisList   []ProgrammeBasis      `db:"basis_list" json:"basisList" field:"basis_list"` // 依据
	ContentList []ConfirmationContent `db:"content_list" json:"contentList" field:"content_list"`
	UserList    []ConfirmationUser    `db:"user_list" json:"userList" field:"user_list"`
	FileList    []entity.File         `db:"file_list" json:"fileList" field:"file_list"`
}
