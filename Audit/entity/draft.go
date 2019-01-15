package entity

import "auditIntegralSys/Worker/entity"

// 单条工作底稿
type DraftItem struct {
	Id                  int    `db:"id" json:"id" field:"id"`
	ProgrammeId         int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	ProgrammeTitle      string `db:"programme_title" json:"programmeTitle" field:"programme_title"`
	QueryDepartmentId   int    `db:"query_department_id" json:"queryDepartmentId" field:"query_department_id"`
	QueryDepartmentName string `db:"query_department_name" json:"queryDepartmentName" field:"query_department_name"`
	DepartmentId        int    `db:"department_id" json:"departmentId" field:"department_id"`
	DepartmentName      string `db:"department_name" json:"departmentName" field:"department_name"`
	IntroductionId      int    `db:"introduction_id" json:"introductionId" field:"introduction_id"`
	Year                int    `db:"year" json:"year" field:"year"`
	Number              int    `db:"number" json:"number" field:"number"`
	ProjectName         string `db:"project_name" json:"projectName" field:"project_name"`
	Public              uint8  `db:"public" json:"public" field:"public"`
	QueryStartTime      string `db:"query_start_time" json:"queryStartTime" field:"query_start_time"`
	QueryEndTime        string `db:"query_end_time" json:"queryEndTime" field:"query_end_time"`
	UpdateTime          string `db:"update_time" json:"updateTime" field:"update_time"`
	AuthorId            int    `db:"author_id" json:"authorId" field:"author_id"`
	State               string `db:"state" json:"state" field:"state"`
}

// 工作底稿内容
type DraftContent struct {
	Id              int    `db:"id" json:"id" field:"id"`
	DraftId         int    `db:"draft_id" json:"draftId" field:"draft_id"`
	Order           int    `db:"order" json:"order" field:"order"`
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

// 检查人
type DraftQueryUser struct {
	Id       int    `db:"id" json:"id" field:"id"`
	DraftId  int    `db:"draft_id" json:"draftId" field:"draft_id"`
	UserId   int    `db:"user_id" json:"userId" field:"user_id"`
	UserName string `db:"user_name" json:"userName" field:"user_name"`
	IsLeader int    `db:"is_leader" json:"isLeader" field:"is_leader"`
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
	ContentList    []DraftContent    `db:"content_list" json:"contentList" field:"content_list"`            // 内容
	AdminUserList  []DraftAdminUser  `db:"admin_user_list" json:"adminUserList" field:"admin_user_list"`    // 复查人
	QueryUserList  []DraftQueryUser  `db:"query_user_list" json:"queryUserList" field:"query_user_list"`    // 检查人
	ReviewUserList []DraftReviewUser `db:"review_user_list" json:"reviewUserList" field:"review_user_list"` // 负责人
	FileList       []entity.File     `db:"file_list" json:"fileList" field:"file_list"`                     // 附件
}
