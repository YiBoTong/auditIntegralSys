package app

type Status struct {
	Code  int    `json:"code"`
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
}

type Pager struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

type Response struct {
	Data   interface{} `json:"data"`
	Status Status      `json:"status"`
}

type ListResponse struct {
	Data   interface{} `json:"data"`
	Page   Pager       `json:"page"`
	Status Status      `json:"status"`
}
