package entity

type ConfirmationListItem struct {
	Id                    int    `db:"id" json:"id" field:"id"`
	ProgrammeId           int    `db:"programme_id" json:"programmeId" field:"programme_id"`
	ProgrammeTitle        string `db:"programme_title" json:"programmeTitle" field:"programme_title"`
	QueryDepartmentId     int    `db:"query_department_id" json:"queryDepartmentId" field:"query_department_id"`
	QueryDepartmentName   string `db:"query_department_name" json:"queryDepartmentName" field:"query_department_name"`
	DepartmentId          int    `db:"department_id" json:"departmentId" field:"department_id"`
	DepartmentName        string `db:"department_name" json:"departmentName" field:"department_name"`
	Number                string `db:"number" json:"number" field:"number"`
	ProjectName           string `db:"project_name" json:"projectName" field:"project_name"`
	Public                uint8  `db:"public" json:"public" field:"public"`
	Time                  string `db:"time" json:"time" field:"time"`
	UpdateTime            string `db:"update_time" json:"updateTime" field:"update_time"`
	State                 string `db:"state" json:"state" field:"state"`
	ConfirmationReceiptId int    `db:"confirmation_receipt_id" json:"confirmationReceiptId" field:"confirmation_receipt_id"` // 回执ID
	HasRead               uint8  `db:"has_read" json:"hasRead" field:"has_read"`                                             // 是否已读
	HasReadTime           string `db:"has_read_time" json:"hasReadTime" field:"has_read_time"`                               // 已读时间
}

type ConfirmationItem struct {
	Id                    int    `db:"id" json:"id" field:"id"`
	DraftId               int    `db:"draft_id" json:"draftId" field:"draft_id"`                                             // 工作底稿ID
	ConfirmationReceiptId int    `db:"confirmation_receipt_id" json:"confirmationReceiptId" field:"confirmation_receipt_id"` // 回执ID
	HasRead               uint8  `db:"has_read" json:"hasRead" field:"has_read"`                                             // 是否已读
	HasReadTime           string `db:"has_read_time" json:"hasReadTime" field:"has_read_time"`                               // 已读时间
	State                 string `db:"state" json:"state" field:"state"`
}

type Confirmation struct {
	ConfirmationItem
	Draft        DraftItem        `db:"draft" json:"draft" field:"draft"`                        // 工作底稿
	Programme    ProgrammeItem    `db:"programme" json:"programme" field:"programme"`            // 方案
	BasisList    []ProgrammeBasis `db:"basis_list" json:"basisList" field:"basis_list"`          // 依据
	DraftContent []DraftContent   `db:"draft_content" json:"draftContent" field:"draft_content"` // 工作底稿违规内容
}
