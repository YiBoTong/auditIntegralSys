package entity

type Rbac struct {
	Id       int    `db:"id" json:"id" field:"id"`
	Key      string `db:"key" json:"key" field:"key"`
	Title    string `db:"title" json:"title" field:"title"`
	MenuId   int    `db:"menu_id" json:"menuId" field:"menu_id"`
	HasChild bool   `db:"has_child" json:"hasChild" field:"has_child"`
	IsRead   bool   `db:"is_read" json:"isRead" field:"is_read"`
	IsWrite  bool   `db:"is_write" json:"isWrite" field:"is_write"`
}

type Rbacs struct {
	Rbac
	Children []Rbac `db:"children" json:"children" field:"children"`
}
