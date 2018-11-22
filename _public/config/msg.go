package config

const (
	ListStr        = "获取列表"
	DelStr         = "删除"
	GetStr         = "获取"
	AddStr         = "添加"
	EditStr        = "编辑"
	SuccessStr     = "成功"
	ErrorStr       = "失败"
	ParameterError = "，参数错误"
	UserCode       = "员工号"
	Had            = "已存在"
	NoHad          = "不存在"
)

// 获取操作提示
func GetTodoResMsg(str string, error bool) string {
	if error {
		return str + ErrorStr
	}
	return str + SuccessStr
}
