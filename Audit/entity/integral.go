package entity

type IntegralListItem struct {
	Id                 int    `db:"id" json:"id" field:"id"`
	CognizanceUserId   int    `db:"cognizance_user_id" json:"cognizanceUserId" field:"cognizance_user_id"`
	CognizanceUserName string `db:"cognizance_user_name" json:"cognizanceUserName" field:"cognizance_user_name"`
	UserId             int    `db:"user_id" json:"userId" field:"user_id"`
	UserName           string `db:"user_name" json:"userName" field:"user_name"`
	DraftId            int    `db:"draft_id" json:"draftId" field:"draft_id"`
	PunishNoticeId     int    `db:"punish_notice_id" json:"punishNoticeId" field:"punish_notice_id"`
	ProgrammeId        int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	//ProgrammeTitle      string `db:"programme_title" json:"programmeTitle" field:"programme_title"`
	QueryDepartmentId   int    `db:"query_department_id" json:"queryDepartmentId" field:"query_department_id"`
	QueryDepartmentName string `db:"query_department_name" json:"queryDepartmentName" field:"query_department_name"`
	DepartmentId        int    `db:"department_id" json:"departmentId" field:"department_id"`
	DepartmentName      string `db:"department_name" json:"departmentName" field:"department_name"`
	Score               int    `db:"score" json:"score" field:"score"`
	Number              string `db:"number" json:"number" field:"number"`
	ProjectName         string `db:"project_name" json:"projectName" field:"project_name"`
	Time                string `db:"time" json:"time" field:"time"`
}

type IntegralItem struct {
	Id                 int    `db:"id" json:"id" field:"id"`
	CognizanceUserId   int    `db:"cognizance_user_id" json:"cognizanceUserId" field:"cognizance_user_id"`
	CognizanceUserName string `db:"cognizance_user_name" json:"cognizanceUserName" field:"cognizance_user_name"`
	UserId             int    `db:"user_id" json:"userId" field:"user_id"`
	UserName           string `db:"user_name" json:"userName" field:"user_name"`
	DraftId            int    `db:"draft_id" json:"draftId" field:"draft_id"` // 工作底稿ID
	PunishNoticeId     int    `db:"punish_notice_id" json:"punishNoticeId" field:"punish_notice_id"`
	Score              int    `db:"score" json:"score" field:"score"`
	Time               string `db:"time" json:"time" field:"time"`
}

type IntegralChangeScore struct {
	Id         int    `db:"id" json:"id" field:"id"`
	IntegralId int    `db:"integral_id" json:"integralId" field:"integral_id"`
	Score      int    `db:"score" json:"score" field:"score"`
	UserId     int    `db:"user_id" json:"userId" field:"user_id"`
	Describe   string `db:"describe" json:"describe" field:"describe"`
	UpdateTime string `db:"update_time" json:"updateTime" field:"update_time"`
}

type Integral struct {
	IntegralItem
	SumScore     int                             `db:"sum_score" json:"sumScore" field:"sum_score"`
	BehaviorList []PunishNoticeBasisBehaviorItem `db:"behavior_list" json:"behaviorList" field:"behavior_list"`
}
