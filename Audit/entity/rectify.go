package entity

type RectifyListItem struct {
	Id                  int    `db:"id" json:"id" field:"id"`
	RectifyReportId     int    `db:"rectify_report_id" json:"rectifyReportId" field:"rectify_report_id"`
	DraftId             int    `db:"draft_id" json:"draftId" field:"draft_id"`
	ConfirmationId      int    `db:"confirmation_id" json:"confirmationId" field:"confirmation_id"`
	ProgrammeId         int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	ProgrammeTitle      string `db:"programme_title" json:"programmeTitle" field:"programme_title"`
	QueryDepartmentId   int    `db:"query_department_id" json:"queryDepartmentId" field:"query_department_id"`
	QueryDepartmentName string `db:"query_department_name" json:"queryDepartmentName" field:"query_department_name"`
	DepartmentId        int    `db:"department_id" json:"departmentId" field:"department_id"`
	DepartmentName      string `db:"department_name" json:"departmentName" field:"department_name"`
	Number              string `db:"number" json:"number" field:"number"`
	ProjectName         string `db:"project_name" json:"projectName" field:"project_name"`
	StartTime           string `db:"start_time" json:"startTime" field:"start_time"`
	EndTime             string `db:"end_time" json:"endTime" field:"end_time"`
	PlanStartTime       string `db:"plan_start_time" json:"planStartTime" field:"plan_start_time"`
	PlanEndTime         string `db:"plan_end_time" json:"planEndTime" field:"plan_end_time"`
	Time                string `db:"time" json:"time" field:"time"`
	UpdateTime          string `db:"update_time" json:"updateTime" field:"update_time"`
	State               string `db:"state" json:"state" field:"state"`
	ReportState         string `db:"report_state" json:"reportState" field:"report_state"`
}

type RectifyItem struct {
	Id             int    `db:"id" json:"id" field:"id"`
	DraftId        int    `db:"draft_id" json:"draftId" field:"draft_id"`
	ConfirmationId int    `db:"confirmation_id" json:"confirmationId" field:"confirmation_id"`
	ProgrammeId    int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	UserId         int    `db:"user_id" json:"userId" field:"user_id"`
	UserName       string `db:"user_name" json:"userName" field:"user_name"`
	Suggest        string `db:"suggest" json:"suggest" field:"suggest"`
	UpdateTime     string `db:"update_time" json:"updateTime" field:"update_time"`
	State          string `db:"state" json:"state" field:"state"`
}

type Rectify struct {
	RectifyItem
	Programme         ProgrammeItem       `db:"programme" json:"programme" field:"programme"`                           // 方案
	Draft             DraftItem           `db:"draft" json:"draft" field:"draft"`                                       // 工作底稿
	DraftContent      []DraftContent      `db:"draft_content" json:"draftContent" field:"draft_content"`                // 工作底稿违规内容
	ProgrammeBusiness []ProgrammeBusiness `db:"programme_business" json:"programmeBusiness" field:"programme_business"` // 方案业务范围
}
