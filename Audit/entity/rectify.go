package entity

type RectifyItem struct {
	Id                  int    `db:"id" json:"id" field:"id"`
	ProgrammeId         int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	ProgrammeTitle      string `db:"programme_title" json:"programmeTitle" field:"programme_title"`
	QueryDepartmentId   int    `db:"query_department_id" json:"queryDepartmentId" field:"query_department_id"`
	QueryDepartmentName string `db:"query_department_name" json:"queryDepartmentName" field:"query_department_name"`
	DepartmentId        int    `db:"department_id" json:"departmentId" field:"department_id"`
	DepartmentName      string `db:"department_name" json:"departmentName" field:"department_name"`
	Number              string `db:"number" json:"number" field:"number"`
	ProjectName         string `db:"project_name" json:"projectName" field:"project_name"`
	Time                string `db:"time" json:"time" field:"time"`
	UpdateTime          string `db:"update_time" json:"updateTime" field:"update_time"`
}

type Rectify struct {
	RectifyItem
	Draft        DraftItem      `db:"draft" json:"draft" field:"draft"`                        // 工作底稿
	DraftContent []DraftContent `db:"draft_content" json:"draftContent" field:"draft_content"` // 工作底稿违规内容
}
