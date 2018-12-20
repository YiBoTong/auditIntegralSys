package check

import "auditIntegralSys/_public/config"

type ProgrammeState string
type ProgrammeUserRole string

const (
	P_draft        string = "draft"        // 草稿
	P_report       string = "report"       // 上报
	P_adopt        string = "adopt"        // 通过
	P_reject       string = "reject"       // 驳回
	P_publish      string = "publish"      // 发布
	P_dep_reject   string = "dep_reject"   // 部门负责人驳回
	P_dep_adopt    string = "dep_adopt"    // 部门负责人通过
	P_admin_reject string = "admin_reject" // 分管领导驳回
	P_admin_adopt  string = "admin_adopt"  // 分管领导驳回

	p_draft        ProgrammeState = "draft"        // 草稿
	p_report       ProgrammeState = "report"       // 上报
	p_adopt        ProgrammeState = "adopt"        // 通过
	p_reject       ProgrammeState = "reject"       // 驳回

	//P_author    ProgrammeUserRole = "author"    // 创建人
	P_detUser   ProgrammeUserRole = "detUser"   // 部门负责人
	P_adminUser ProgrammeUserRole = "adminUser" // 分管领导
)

func (this ProgrammeState) Has() (bool, string) {
	msg := ""
	hasState := false
	switch this {
	// 草稿、上报、通过、驳回
	case p_draft, p_report, p_adopt, p_reject:
		hasState = true
	default:
		hasState = false
		msg = config.StateStr + config.NoHad
	}
	return hasState, msg
}

func (this ProgrammeUserRole) Has() (bool, string) {
	msg := ""
	hasState := false
	switch this {
	case P_detUser, P_adminUser:
		hasState = true
	default:
		hasState = false
		msg = config.RoleStr + config.NoHad
	}
	return hasState, msg
}
