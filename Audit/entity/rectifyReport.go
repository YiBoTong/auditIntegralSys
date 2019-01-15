package entity

import "auditIntegralSys/Worker/entity"

type RectifyReportItem struct {
	Id        int    `db:"id" json:"id" field:"id"`
	RectifyId int    `db:"rectify_id" json:"rectifyId" field:"rectify_id"`
	State     string `db:"state" json:"state" field:"state"`
}

type RectifyReportContentUserItem struct {
	UserId   int    `db:"user_id" json:"userId" field:"user_id"`
	UserName string `db:"user_name" json:"userName" field:"user_name"`
}

type RectifyReportContentItem struct {
	Id             int                            `db:"id" json:"id" field:"id"`
	DraftContentId int                            `db:"draft_content_id" json:"draftContentId" field:"draft_content_id"`
	Content        string                         `db:"content" json:"content" field:"content"`
	Time           string                         `db:"time" json:"time" field:"time"`
	UserList       []RectifyReportContentUserItem `db:"user_list" json:"userList" field:"user_list"`
}

type RectifyReport struct {
	RectifyReportItem
	ContentList []RectifyReportContentItem `db:"content_list" json:"contentList" field:"content_list"`
	FileList    []entity.File              `db:"file_list" json:"fileList" field:"file_list"`
}
