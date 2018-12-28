package config

const (
	DepartmentMsgStr  = "部门"
	ListStr           = "获取列表"
	DelStr            = "删除"
	GetStr            = "获取"
	ReadStr           = "已读"
	AddStr            = "添加"
	EditStr           = "编辑"
	UploadStr         = "上传"
	SuccessStr        = "成功"
	ErrorStr          = "失败"
	UserCode          = "员工号"
	StateStr          = "状态"
	RoleStr           = "人员角色"
	Had               = "已存在"
	NoHad             = "不存在"
	ChangeState       = "状态变更"
	LoginStr          = "登录"
	LogoutStr         = "退出"
	UserInfoStr       = "人员信息"
	LoginTipStr       = "登录超时，请重新登录"
	PasswordErrStr    = "密码错误"
	ChangePasswordStr = "修改密码"
	PasswordLenErrStr = "密码长度为6-18位"
	RbacStr           = "权限"
	MastHasOneStr     = "至少要有一个"
	ExamineStr        = "审核"
)

// 获取操作提示
func GetTodoResMsg(str string, error bool) string {
	if error {
		return str + ErrorStr
	}
	return str + SuccessStr
}
