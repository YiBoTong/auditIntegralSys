package entity

type Log struct {
	Id        int    `db:"" json:"" field:""`
	Type      string `db:"type" json:"type" field:"type"`
	TypeTitle string `db:"type_title" json:"typeTitle" field:"type_title"`
	UserId    int    `db:"user_id" json:"userId" field:"user_id"`
	UserName  string `db:"user_name" json:"userName" field:"user_name"`
	Method    string `db:"method" json:"method" field:"method"`
	Msg       string `db:"msg" json:"msg" field:"msg"`
	Time      string `db:"time" json:"time" field:"time"`
}
