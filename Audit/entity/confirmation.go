package entity

type ConfirmationItem struct {
	Id                    int    `db:"id" json:"id" field:"id"`
	DraftId               int    `db:"draft_id" json:"draftId" field:"draft_id"`                                             // 工作底稿ID
	ProgrammeId           int    `db:"programme_id" json:"programmeId" field:"programme_id"`                                 // 方案ID（工作底稿表）
	ConfirmationReceiptId int    `db:"confirmation_receipt_id" json:"confirmationReceiptId" field:"confirmation_receipt_id"` // 回执ID
	HasRead               uint8  `db:"has_read" json:"hasRead" field:"has_read"`                                             // 是否已读
	HasReadTime           string `db:"has_read_time" json:"hasReadTime" field:"has_read_time"`                               // 已读时间
}

type Confirmation struct {
	ConfirmationItem
	Draft        DraftItem     `db:"draft" json:"draft" field:"draft"`                        // 工作底稿
	Programme    ProgrammeItem `db:"programme" json:"programme" field:"programme"`            // 方案
	DraftContent DraftContent  `db:"draft_content" json:"draftContent" field:"draft_content"` // 工作底稿违规内容
}
