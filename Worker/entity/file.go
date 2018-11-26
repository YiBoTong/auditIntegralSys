package entity

type File struct {
	Id       int    `db:"id" json:"id" field:"id"`
	Name     string `db:"name" json:"name" field:"name"`
	Suffix   string `db:"suffix" json:"suffix" field:"suffix"`
	Time     string `db:"time" json:"time" field:"time"`
	FileName string `db:"file_name" json:"fileName" field:"file_name"`
	Path     string `db:"path" json:"path" field:"path"`
}
