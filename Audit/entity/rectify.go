package entity

type RectifyListItem struct {
	Id                  int    `db:"id" json:"id" field:"id"`
	Year                int    `db:"year" json:"year" field:"year"`
	Number              int    `db:"number" json:"number" field:"number"`
	RectifyReportId     int    `db:"rectify_report_id" json:"rectifyReportId" field:"rectify_report_id"`
	DraftId             int    `db:"draft_id" json:"draftId" field:"draft_id"`
	ConfirmationId      int    `db:"confirmation_id" json:"confirmationId" field:"confirmation_id"`
	ProgrammeId         int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	ProgrammeTitle      string `db:"programme_title" json:"programmeTitle" field:"programme_title"`
	QueryDepartmentId   int    `db:"query_department_id" json:"queryDepartmentId" field:"query_department_id"`
	QueryDepartmentName string `db:"query_department_name" json:"queryDepartmentName" field:"query_department_name"`
	DepartmentId        int    `db:"department_id" json:"departmentId" field:"department_id"`
	DepartmentName      string `db:"department_name" json:"departmentName" field:"department_name"`
	ProjectName         string `db:"project_name" json:"projectName" field:"project_name"`
	StartTime           string `db:"start_time" json:"startTime" field:"start_time"`
	EndTime             string `db:"end_time" json:"endTime" field:"end_time"`
	PlanStartTime       string `db:"plan_start_time" json:"planStartTime" field:"plan_start_time"`
	PlanEndTime         string `db:"plan_end_time" json:"planEndTime" field:"plan_end_time"`
	QueryStartTime      string `db:"query_start_time" json:"queryStartTime" field:"query_start_time"`
	QueryEndTime        string `db:"query_end_time" json:"queryEndTime" field:"query_end_time"`
	UpdateTime          string `db:"update_time" json:"updateTime" field:"update_time"`
	State               string `db:"state" json:"state" field:"state"`
	AuthorId            int    `db:"author_id" json:"authorId" field:"author_id"`
	ReportState         string `db:"report_state" json:"reportState" field:"report_state"`
}

type RectifyItem struct {
	Id              int    `db:"id" json:"id" field:"id"`
	Year            int    `db:"year" json:"year" field:"year"`
	Number          int    `db:"number" json:"number" field:"number"`
	RectifyReportId int    `db:"rectify_report_id" json:"rectifyReportId" field:"rectify_report_id"`
	DraftId         int    `db:"draft_id" json:"draftId" field:"draft_id"`
	ConfirmationId  int    `db:"confirmation_id" json:"confirmationId" field:"confirmation_id"`
	ProgrammeId     int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	UserId          int    `db:"user_id" json:"userId" field:"user_id"`
	UserName        string `db:"user_name" json:"userName" field:"user_name"`
	LastTime        string `db:"last_time" json:"lastTime" field:"last_time"`
	UpdateTime      string `db:"update_time" json:"updateTime" field:"update_time"`
	State           string `db:"state" json:"state" field:"state"`
}

type RectifyContent struct {
	Content string `db:"content" json:"content" field:"content"`
}

type Rectify struct {
	RectifyItem
	Demand              string                `db:"demand" json:"demand" field:"demand"`
	Suggest             string                `db:"suggest" json:"suggest" field:"suggest"`
	Programme           ProgrammeItem         `db:"programme" json:"programme" field:"programme"`                                 // 方案
	Draft               DraftItem             `db:"draft" json:"draft" field:"draft"`                                             // 工作底稿
	ConfirmationContent []ConfirmationContent `db:"confirmation_content" json:"confirmationContent" field:"confirmation_content"` // 事实确认书违规内容
	ProgrammeBusiness   []ProgrammeBusiness   `db:"programme_business" json:"programmeBusiness" field:"programme_business"`       // 方案业务范围
}
