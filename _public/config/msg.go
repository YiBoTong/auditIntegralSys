package config

const (
	DepartmentMsgStr = "部门"
	ListStr          = "获取列表"
	DelStr           = "删除"
	GetStr           = "获取"
	AddStr           = "添加"
	EditStr          = "编辑"
	SuccessStr       = "成功"
	ErrorStr         = "失败"
	UserCode         = "员工号"
	StateStr         = "状态"
	Had              = "已存在"
	NoHad            = "不存在"
	ChangeState      = "状态变更"
	LoginStr         = "登录"
	LogoutStr        = "退出"
	UserInfoStr      = "人员信息"
	LoginTispStr     = "登录超时，请重新登录"
)

// 获取操作提示
func GetTodoResMsg(str string, error bool) string {
	if error {
		return str + ErrorStr
	}
	return str + SuccessStr
}
