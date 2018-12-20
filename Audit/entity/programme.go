package entity

type AddProgramme struct {
	Title             string `db:"title" json:"title" field:"title"`
	QueryDepartmentId int    `db:"query_department_id" json:"queryDepartmentId" field:"query_department_id"`
	UserId            int    `db:"user_id" json:"userId" field:"user_id"`
	QueryPointId      int    `db:"query_point_id" json:"queryPointId" field:"query_point_id"`
	Purpose           string `db:"purpose" json:"purpose" field:"purpose"`
	Type              string `db:"type" json:"type" field:"type"`
	StartTime         string `db:"start_time" json:"startTime" field:"start_time"`
	EndTime           string `db:"end_time" json:"endTime" field:"end_time"`
	PlanStartTime     string `db:"plan_start_time" json:"planStartTime" field:"plan_start_time"`
	PlanEndTime       string `db:"plan_end_time" json:"planEndTime" field:"plan_end_time"`
	State             string `db:"state" json:"state" field:"state"`
}

type ProgrammeSelectItem struct {
	Id                  int    `db:"id" json:"id" field:"id"`
	Title               string `db:"title" json:"title" field:"title"`
	Key                 string `db:"key" json:"key" field:"key"`
	QueryDepartmentId   int    `db:"query_department_id" json:"queryDepartmentId" field:"query_department_id"`
	QueryDepartmentName string `db:"query_department_name" json:"queryDepartmentName" field:"query_department_name"`
	UserId              int    `db:"user_id" json:"userId" field:"user_id"`
	QueryPointId        int    `db:"query_point_id" json:"queryPointId" field:"query_point_id"`
	QueryPointName      string `db:"query_point_name" json:"queryPointName" field:"query_point_name"`
	Purpose             string `db:"purpose" json:"purpose" field:"purpose"`
	Type                string `db:"type" json:"type" field:"type"`
	StartTime           string `db:"start_time" json:"startTime" field:"start_time"`
	EndTime             string `db:"end_time" json:"endTime" field:"end_time"`
	PlanStartTime       string `db:"plan_start_time" json:"planStartTime" field:"plan_start_time"`
	PlanEndTime         string `db:"plan_end_time" json:"planEndTime" field:"plan_end_time"`
}

type ProgrammeItem struct {
	Id                  int    `db:"id" json:"id" field:"id"`
	Title               string `db:"title" json:"title" field:"title"`
	Key                 string `db:"key" json:"key" field:"key"`
	QueryDepartmentId   int    `db:"query_department_id" json:"queryDepartmentId" field:"query_department_id"`
	QueryDepartmentName string `db:"query_department_name" json:"queryDepartmentName" field:"query_department_name"`
	UserId              int    `db:"user_id" json:"userId" field:"user_id"`
	QueryPointId        int    `db:"query_point_id" json:"queryPointId" field:"query_point_id"`
	QueryPointName      string `db:"query_point_name" json:"queryPointName" field:"query_point_name"`
	Purpose             string `db:"purpose" json:"purpose" field:"purpose"`
	Type                string `db:"type" json:"type" field:"type"`
	StartTime           string `db:"start_time" json:"startTime" field:"start_time"`
	EndTime             string `db:"end_time" json:"endTime" field:"end_time"`
	PlanStartTime       string `db:"plan_start_time" json:"planStartTime" field:"plan_start_time"`
	PlanEndTime         string `db:"plan_end_time" json:"planEndTime" field:"plan_end_time"`
	State               string `db:"state" json:"state" field:"state"`
}

type ProgrammeDepExamine struct {
	Id       int    `db:"id" json:"id" field:"id"`
	UserId   int    `db:"user_id" json:"userId" field:"user_id"`
	UserName string `db:"user_name" json:"userName" field:"user_name"`
	Content  string `db:"content" json:"content" field:"content"`
	Time     string `db:"time" json:"time" field:"time"`
	State    string `db:"state" json:"state" field:"state"`
}

type ProgrammeAdminExamine struct {
	Id       int    `db:"id" json:"id" field:"id"`
	UserId   int    `db:"user_id" json:"userId" field:"user_id"`
	UserName string `db:"user_name" json:"userName" field:"user_name"`
	Content  string `db:"content" json:"content" field:"content"`
	Time     string `db:"time" json:"time" field:"time"`
	State    string `db:"state" json:"state" field:"state"`
}

type ProgrammeBasis struct {
	Id          int    `db:"id" json:"id" field:"id"`
	ProgrammeId int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	ClauseId    int    `db:"clause_id" json:"clauseId" field:"clause_id"`
	Content     string `db:"content" json:"content" field:"content"`
	Order       int    `db:"order" json:"order" field:"order"`
}

type ProgrammeBusiness struct {
	Id          int    `db:"id" json:"id" field:"id"`
	ProgrammeId int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	Type        string `db:"type" json:"type" field:"type"`
	Content     string `db:"content" json:"content" field:"content"`
	Order       int    `db:"order" json:"order" field:"order"`
}

type ProgrammeContent struct {
	Id          int    `db:"id" json:"id" field:"id"`
	ProgrammeId int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	Content     string `db:"content" json:"content" field:"content"`
	Order       int    `db:"order" json:"order" field:"order"`
}

type ProgrammeEmphases struct {
	Id          int    `db:"id" json:"id" field:"id"`
	ProgrammeId int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	Content     string `db:"content" json:"content" field:"content"`
	Order       int    `db:"order" json:"order" field:"order"`
}

type ProgrammeStep struct {
	Id          int    `db:"id" json:"id" field:"id"`
	ProgrammeId int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	Type        string `db:"type" json:"type" field:"type"`
	Content     string `db:"content" json:"content" field:"content"`
	Order       int    `db:"order" json:"order" field:"order"`
}

type ProgrammeUser struct {
	Id          int    `db:"id" json:"id" field:"id"`
	ProgrammeId int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	UserId      int    `db:"user_id" json:"userId" field:"user_id"`
	Job         string `db:"job" json:"job" field:"job"`
	JobName     string `db:"job_name" json:"jobName" field:"job_name"`
	Title       string `db:"title" json:"title" field:"title"`
	Task        string `db:"task" json:"task" field:"task"`
	Order       int    `db:"order" json:"order" field:"order"`
}

type Programme struct {
	ProgrammeItem
	Basis         []ProgrammeBasis        `db:"basis" json:"basis" field:"basis"`
	Content       []ProgrammeContent      `db:"content" json:"content" field:"content"`
	Step          []ProgrammeStep         `db:"step" json:"step" field:"step"`
	Business      []ProgrammeBusiness     `db:"business" json:"business" field:"business"`
	Emphases      []ProgrammeEmphases     `db:"emphases" json:"emphases" field:"emphases"`
	UserList      []ProgrammeUser         `db:"user_list" json:"userList" field:"userList"`
	DepExamines   []ProgrammeDepExamine   `db:"dep_examines" json:"depExamines" field:"dep_examines"`
	AdminExamines []ProgrammeAdminExamine `db:"admin_examine" json:"adminExamine" field:"admin_examine"`
}
