package entity

import "auditIntegralSys/Worker/entity"

// 单条工作底稿
type DraftItem struct {
	Id                  int    `db:"id" json:"id" field:"id"`
	ProgrammeId         int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	ProgrammeName       string `db:"programme_name" json:"programmeName" field:"programme_name"`
	QueryDepartmentId   string `db:"query_department_id" json:"queryDepartmentId" field:"query_department_id"`
	QueryDepartmentName string `db:"query_department_name" json:"queryDepartmentName" field:"query_department_name"`
	DepartmentId        int    `db:"department_id" json:"departmentId" field:"department_id"`
	DepartmentName      string `db:"department_name" json:"departmentName" field:"department_name"`
	Number              string `db:"number" json:"number" field:"number"`
	ProjectName         string `db:"project_name" json:"projectName" field:"project_name"`
	Public              uint8  `db:"public" json:"public" field:"public"`
	Time                string `db:"time" json:"time" field:"time"`
	UpdateTime          string `db:"update_time" json:"updateTime" field:"update_time"`
	State               string `db:"state" json:"state" field:"state"`
}

// 工作底稿内容
type DraftContent struct {
	Id              int    `db:"id" json:"id" field:"id"`
	DraftId         int    `db:"draft_id" json:"draftId" field:"draft_id"`
	Type            string `db:"type" json:"type" field:"type"`
	BehaviorId      int    `db:"behavior_id" json:"behaviorId" field:"behavior_id"`
	BehaviorContent string `db:"behavior_content" json:"behaviorContent" field:"behavior_content"`
}

// 复查人
type DraftAdminUser struct {
	Id       int    `db:"id" json:"id" field:"id"`
	DraftId  int    `db:"draft_id" json:"draftId" field:"draft_id"`
	UserId   int    `db:"user_id" json:"userId" field:"user_id"`
	UserName string `db:"user_name" json:"userName" field:"user_name"`
}

// 被检查人
type DraftInspectUser struct {
	Id       int    `db:"id" json:"id" field:"id"`
	DraftId  int    `db:"draft_id" json:"draftId" field:"draft_id"`
	UserId   int    `db:"user_id" json:"userId" field:"user_id"`
	UserName string `db:"user_name" json:"userName" field:"user_name"`
}

// 检查人
type DraftQueryUser struct {
	Id       int    `db:"id" json:"id" field:"id"`
	DraftId  int    `db:"draft_id" json:"draftId" field:"draft_id"`
	UserId   int    `db:"user_id" json:"userId" field:"user_id"`
	UserName string `db:"user_name" json:"userName" field:"user_name"`
}

// 负责人
type DraftReviewUser struct {
	Id       int    `db:"id" json:"id" field:"id"`
	DraftId  int    `db:"draft_id" json:"draftId" field:"draft_id"`
	UserId   int    `db:"user_id" json:"userId" field:"user_id"`
	UserName string `db:"user_name" json:"userName" field:"user_name"`
}

type Draft struct {
	DraftItem
	ContentList []DraftContent     `db:"content_list" json:"contentList" field:"content_list"` // 内容
	AdminUser   DraftAdminUser     `db:"admin_user" json:"adminUser" field:"admin_user"`       // 复查人
	InspectUser []DraftInspectUser `db:"inspect_user" json:"inspectUser" field:"inspect_user"` // 被检查人
	QueryUser   []DraftQueryUser   `db:"query_user" json:"queryUser" field:"query_user"`       // 检查人
	ReviewUser  DraftReviewUser    `db:"review_user" json:"reviewUser" field:"review_user"`    // 负责人
	FileList    []entity.File      `db:"file_list" json:"fileList" field:"file_list"`          // 附件
}
