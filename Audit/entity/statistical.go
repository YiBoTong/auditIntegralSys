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

type StatisticalDetailedTotal struct {
	SumScore int `db:"sum_score" json:"sumScore" field:"sum_score"`
	SumMoney int `db:"sum_money" json:"sumMoney" field:"sum_money"`
}

type StatisticalDetailedItem struct {
	UserId               int    `db:"user_id" json:"userId" field:"user_id"`
	UserName             string `db:"user_name" json:"userName" field:"user_name"`
	Sex                  int    `db:"sex" json:"sex" field:"sex"`
	Money                int    `db:"money" json:"money" field:"money"`
	Score                int    `db:"score" json:"score" field:"score"`
	DepartmentName       string `db:"department_name" json:"departmentName" field:"department_name"`
	Title                string `db:"title" json:"title" field:"title"`
	Number               string `db:"number" json:"number" field:"number"`
	PunishDepartmentName string `db:"punish_department_name" json:"punishDepartmentName" field:"punish_department_name"`
	Role                 string `db:"role" json:"role" field:"role"`
}

type StatisticalMyNumByYearItem struct {
	Month int `db:"month" json:"month" field:"month"`
	Num   int `db:"num" json:"num" field:"num"`
}

type StatisticalMyNumByYear struct {
	Year int                          `db:"year" json:"year" field:"year"`
	Info []StatisticalMyNumByYearItem `db:"info" json:"info" field:"info"`
}

type StatisticalTopBehaviorContentItem struct {
	Content string `db:"content" json:"content" field:"content"`
	Sum     int    `db:"sum" json:"sum" field:"sum"`
}

type StatisticalTopBehaviorDepartmentItem struct {
	Name string `db:"name" json:"name" field:"name"`
	Sum  int    `db:"sum" json:"sum" field:"sum"`
}

type StatisticalTopBehaviorAndDepartment struct {
	Behavior   []StatisticalTopBehaviorContentItem    `db:"behavior" json:"behavior" field:"behavior"`
	Department []StatisticalTopBehaviorDepartmentItem `db:"department" json:"department" field:"department"`
	Score      []StatisticalMyNumByYearItem           `db:"score" json:"score" field:"score"`
}

type StatisticalViolationTopDepartment struct {
	Year []StatisticalMyNumByYearItem `db:"year" json:"year" field:"year"`
}

type StatisticalOneUserScoreItem struct {
	UserName string `db:"user_name" json:"userName" field:"user_name"`
	UserId   int    `db:"user_id" json:"userId" field:"user_id"`
	Money    int    `db:"money" json:"money" field:"money"`
	Score    int    `db:"score" json:"score" field:"score"`
}

type StatisticalOneUserBehaviorItem struct {
	UserName string `db:"user_name" json:"userName" field:"user_name"`
	UserId   int    `db:"user_id" json:"userId" field:"user_id"`
	Num      int    `db:"num" json:"num" field:"num"`
}

type StatisticalOneUser struct {
	Score    []StatisticalOneUserScoreItem    `db:"score" json:"score" field:"score"`
	Behavior []StatisticalOneUserBehaviorItem `db:"behavior" json:"behavior" field:"behavior"`
}
