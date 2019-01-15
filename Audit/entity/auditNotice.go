package entity

type AuditNoticeListItem struct {
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

type AuditNoticeItem struct {
	Id      int `db:"id" json:"id" field:"id"`
	DraftId int `db:"draft_id" json:"draftId" field:"draft_id"`
	Year    int `db:"year" json:"year" field:"year"`
	Number  int `db:"number" json:"number" field:"number"`
}

type AuditNotice struct {
	AuditNoticeItem
	Draft    DraftItem           `db:"draft" json:"draft" field:"draft"`
	UserList []DraftQueryUser    `db:"user_list" json:"userList" field:"user_list"`
	Business []ProgrammeBusiness `db:"business" json:"business" field:"business"`
}
