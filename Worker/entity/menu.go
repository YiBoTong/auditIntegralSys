package entity

type Menu struct {
	Id       int    `db:"id" json:"id" field:"id"`
	Path     string `db:"path" json:"path" field:"path"`
	Title    string `db:"title" json:"title" field:"title"`
	Icon     string `db:"icon" json:"icon" field:"icon"`
	NoCache  bool   `db:"no_cache" json:"noCache" field:"no_cache"`
	Order    int    `db:"order" json:"order" field:"order"`
	ParentId int    `db:"parent_id" json:"parentId" field:"parent_id"`
	HasChild bool   `db:"has_child" json:"hasChild" field:"has_child"`
	Time     string `db:"time" json:"time" field:"time"`
}

type Menus struct {
	Menu
	Children map[Menu]interface{} `db:"children" json:"children" field:"children"`
}
