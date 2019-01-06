package entity

type User struct {
	UserId     int    `db:"user_id" json:"userId" field:"user_id"`
	UserName   string `db:"user_name" json:"userName" field:"user_name"`
	AuthorName string `db:"author_name" json:"authorName" field:"author_name"`
	LoginInfo
}

type LoginInfo struct {
	LoginId      int    `db:"login_id" json:"login_id" field:"login_id"`
	UserCode     string `db:"user_code" json:"userCode" field:"user_code"`
	UserId       int    `db:"user_id" json:"userId" field:"user_id"`
	Password     string `db:"password" json:"-" field:"password"`
	AuthorId     int    `db:"author_id" json:"authorId" field:"author_id"`
	IsUse        bool   `db:"is_use" json:"isUse" field:"is_use"`
	LoginTime    string `db:"login_time" json:"loginTime" field:"login_time"`
	LoginNum     int    `db:"login_num" json:"loginNum" field:"login_num"`
	ChangePdTime string `db:"change_pd_time" json:"changePdTime" field:"change_pd_time"`
	Delete       bool   `db:"delete" json:"-" field:"delete"`
}
