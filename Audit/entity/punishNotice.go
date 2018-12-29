package entity

type PunishNoticeItem struct {
	Id                  int    `db:"id" json:"id" field:"id"`
	UserId              int    `db:"user_id" json:"userId" field:"user_id"`       // 被通知人员id
	UserName            string `db:"user_name" json:"userName" field:"user_name"` // 被通知人员姓名
	ConfirmationId      int    `db:"confirmation_id" json:"confirmationId" field:"confirmation_id"`
	DraftId             int    `db:"draft_id" json:"draftId" field:"draft_id"`
	IntegralId          int    `db:"integral_id" json:"integralId" field:"integral_id"`
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
	UpdateTime          string `db:"update_time" json:"updateTime" field:"update_time"`
	Time                string `db:"time" json:"time" field:"time"`
	State               string `db:"state" json:"state" field:"state"`
}

type PunishNoticeBasisItem struct {
	Id               int    `db:"id" json:"id" field:"id"`
	ProgrammeBasisId int    `db:"programme_basis_id" json:"programmeBasisId" field:"programme_basis_id"` // 方案依据ID
	Content          string `db:"content" json:"content" field:"content"`                                // 方案依据内容
}

// 违规行为
type PunishNoticeBasisBehaviorItem struct {
	Id         int    `db:"id" json:"id" field:"id"`
	UserId     int    `db:"user_id" json:"userId" field:"user_id"`
	UserName   string `db:"user_name" json:"userName" field:"user_name"`
	BehaviorId int    `db:"behavior_id" json:"behaviorId" field:"behavior_id"`
	Content    string `db:"content" json:"content" field:"content"`
	UpdateTime string `db:"update_time" json:"updateTime" field:"update_time"`
}

type PunishNoticeScore struct {
	Score int `db:"score" json:"score" field:"score"` // 本次扣分
}

type PunishNotice struct {
	PunishNoticeItem
	PunishNoticeScore
	SumScore     int                             `db:"sum_score" json:"sumScore" field:"sum_score"` // 总扣分（不含本次扣分）
	BasisList    []PunishNoticeBasisItem         `db:"basis_list" json:"basisList" field:"basis_list"`
	BehaviorList []PunishNoticeBasisBehaviorItem `db:"behavior_list" json:"behaviorList" field:"behavior_list"`
}
