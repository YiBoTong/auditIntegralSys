package entity

type Log struct {
	Id       int    `db:"" json:"" field:""`
	Url      string `db:"url" json:"url" field:"url"`
	Server   string `db:"server" json:"server" field:"server"`
	UserId   int    `db:"user_id" json:"userId" field:"user_id"`
	UserName string `db:"user_name" json:"userName" field:"user_name"`
	Method   string `db:"method" json:"method" field:"method"`
	Msg      string `db:"msg" json:"msg" field:"msg"`
	Time     string `db:"time" json:"time" field:"time"`
}
