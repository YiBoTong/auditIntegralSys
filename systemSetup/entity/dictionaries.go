package entity

// 搜索字典类型
type SearchDictionaryType struct {
	Title  string `db:"title" json:"title" field:"title"`
	Key    string `db:"key" json:"key" field:"key"`
	UserId int    `db:"user_id" json:"userId" field:"user_id"`
}

// 字典类型
type DictionaryType struct {
	Id         int    `db:"id" json:"id" field:"id"`
	TypeId     int    `db:"type_id" json:"typeId" field:"type_id"`
	Key        string `db:"key" json:"key" field:"key"`
	Title      string `db:"title" json:"title" field:"title"`
	IsUse      bool   `db:"is_use" json:"isUse" field:"is_use"`
	UpdateTime string `db:"update_time" json:"updateTime" field:"update_time"`
	UserId     int    `db:"user_id" json:"userId" field:"user_id"`
	UserName   string `db:"user_name" json:"userName" field:"user_name"`
	Describe   string `db:"describe" json:"describe" field:"describe"`
}

type DelDictionaryType struct {
	Id     int  `db:"id" json:"id"`
	Delete bool `db:"delete" json:"-" field:"delete"`
}

type AddDictionaryType struct {
	TypeId     int    `db:"type_id" json:"typeId" field:"type_id" gvalid:"type_id" @integer#字典类型必须选择`
	Key        string `db:"key" json:"key" field:"key" gvalid:"key" @string|max-length:20#类型最多为20个字符串`
	Title      string `db:"title" json:"title" field:"title" gvalid:"title" @string|max-length:20#类型最多为20个字符串`
	IsUse      bool   `db:"is_use" json:"isUse" field:"is_use" gvalid:"is_use" @boolean#是否使用为布尔值`
	UpdateTime string `db:"update_time" json:"updateTime" field:"update_time"`
	UserId     int    `db:"user_id" json:"userId" field:"user_id"`
	Describe   string `db:"describe" json:"describe" field:"describe" gvalid:"describe" @string|max-length:250#描述最多输入250个字符`
}

// 字典
type Dictionary struct {
	Id       int    `db:"id" json:"id" field:"id"`
	TypeId   int    `db:"type_id" json:"typeId" field:"type_id"`
	Key      string `db:"key" json:"key" field:"key"`
	Value    string `db:"value" json:"value" field:"value"`
	Order    int    `db:"order" json:"order" field:"order"`
	Describe string `db:"describe" json:"describe" field:"describe"`
}

type DelDictionary struct {
	Id     int  `db:"id" json:"id"`
	Delete bool `db:"delete" json:"-" field:"delete"`
}

type AddDictionary struct {
	TypeId   int    `db:"type_id" json:"typeId" field:"type_id"`
	Key      string `db:"key" json:"key" field:"key"`
	Value    string `db:"value" json:"value" field:"value"`
	Order    int    `db:"order" json:"order" field:"order"`
	Describe string `db:"describe" json:"describe" field:"describe"`
}
