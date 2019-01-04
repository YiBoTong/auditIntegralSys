package entity

type StatisticalListItem struct {
	Id              int    `db:"id" json:"id" field:"id"`
	ProgrammeId     int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	ProgrammeTitle  string `db:"programme_title" json:"programmeTitle" field:"programme_title"`
	StartTime       string `db:"start_time" json:"startTime" field:"start_time"`
	EndTime         string `db:"end_time" json:"endTime" field:"end_time"`
	DraftId         int    `db:"draft_id" json:"draftId" field:"draft_id"`
	ProjectName     string `db:"project_name" json:"projectName" field:"project_name"`
	Number          string `db:"number" json:"number" field:"number"`
	ConfirmationId  int    `db:"confirmation_id" json:"confirmationId" field:"confirmation_id"`
	RectifyReportId int    `db:"rectify_report_id" json:"rectifyReportId" field:"rectify_report_id"`
}

type Statistical struct {
	StatisticalListItem
	BusinessList []ProgrammeBusiness `db:"business_list" json:"businessList" field:"business_list"`
}
