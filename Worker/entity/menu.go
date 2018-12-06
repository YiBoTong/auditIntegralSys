package entity

type Menu struct {
	Id       int    `db:"id" json:"id" field:"id"`
	Path     string `db:"path" json:"path" field:"path"`
	Order    int    `db:"order" json:"order" field:"order"`
	ParentId int    `db:"parent_id" json:"parentId" field:"parent_id"`
	HasChild bool   `db:"has_child" json:"hasChild" field:"has_child"`
	Time     string `db:"time" json:"time" field:"time"`
	IsUse    bool   `db:"is_use" json:"isUse" field:"is_use"`
	Meta     Meta   `db:"meta" json:"meta" field:"meta"`
}

type Meta struct {
	Icon    string `db:"icon" json:"icon" field:"icon"`
	NoCache bool   `db:"no_cache" json:"noCache" field:"no_cache"`
	Title   string `db:"title" json:"title" field:"title"`
}

type Menus struct {
	Menu
	Children []Menu `db:"children" json:"children" field:"children"`
}
