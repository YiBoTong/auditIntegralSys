package entity

type Rbac struct {
	IsRead  bool `db:"is_read" json:"isRead" field:"is_read"`
	IsWrite bool `db:"is_write" json:"isWrite" field:"is_write"`
}

type RbacMenu struct {
	Menu
	Rbac
}

type RbacMenus struct {
	RbacMenu
	Children []RbacMenu `db:"children" json:"children" field:"children"`
}
