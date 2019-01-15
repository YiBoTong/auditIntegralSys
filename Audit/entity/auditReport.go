package entity

type AuditReportListItem struct {
	Id              int    `db:"id" json:"id" field:"id"`
	ProgrammeId     int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	ProgrammeTitle  string `db:"programme_title" json:"programmeTitle" field:"programme_title"`
	StartTime       string `db:"start_time" json:"startTime" field:"start_time"`
	EndTime         string `db:"end_time" json:"endTime" field:"end_time"`
	QueryStartTime  string `db:"query_start_time" json:"queryStartTime" field:"query_start_time"`
	QueryEndTime    string `db:"query_end_time" json:"queryEndTime" field:"query_end_time"`
	UpdateTime      string `db:"update_time" json:"updateTime" field:"update_time"`
	DraftId         int    `db:"draft_id" json:"draftId" field:"draft_id"`
	ProjectName     string `db:"project_name" json:"projectName" field:"project_name"`
	Year            int    `db:"year" json:"year" field:"year"`
	Number          int    `db:"number" json:"number" field:"number"`
	ConfirmationId  int    `db:"confirmation_id" json:"confirmationId" field:"confirmation_id"`
	RectifyReportId int    `db:"rectify_report_id" json:"rectifyReportId" field:"rectify_report_id"`
	AuthorId        int    `db:"author_id" json:"authorId" field:"author_id"`
	State           string `db:"state" json:"state" field:"state"`
}

type AuditReportItem struct {
	Id              int    `db:"id" json:"id" field:"id"`
	ProgrammeId     int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	DraftId         int    `db:"draft_id" json:"draftId" field:"draft_id"`
	ConfirmationId  int    `db:"confirmation_id" json:"confirmationId" field:"confirmation_id"`
	RectifyReportId int    `db:"rectify_report_id" json:"rectifyReportId" field:"rectify_report_id"`
	Year            int    `db:"year" json:"year" field:"year"`
	Number          int    `db:"number" json:"number" field:"number"`
	State           string `db:"state" json:"state" field:"state"`
}

type AuditReportContent struct {
	Id      int    `db:"id" json:"id" field:"id"`
	Content string `db:"content" json:"content" field:"content"`
}

type AuditReport struct {
	AuditReportItem
	BasicInfo string `db:"basic_info" json:"basicInfo" field:"basic_info"`
	Reason    string `db:"reason" json:"reason" field:"reason"`
	Plan      string `db:"plan" json:"plan" field:"plan"`
}
