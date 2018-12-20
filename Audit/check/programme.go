package check

import "auditIntegralSys/_public/config"

type ProgrammeState string
type ProgrammeUserRole string

const (
	p_draft        ProgrammeState = "draft"        // 草稿
	p_report       ProgrammeState = "report"       // 上报
	p_dep_reject   ProgrammeState = "dep_reject"   // 部门负责人驳回
	p_dep_adopt    ProgrammeState = "dep_adopt"    // 部门负责人通过
	p_admin_reject ProgrammeState = "admin_reject" // 分管领导驳回
	p_admin_adopt  ProgrammeState = "publish"      // 分管领导通过

	//p_author    ProgrammeUserRole = "author"    // 创建人
	p_detUser   ProgrammeUserRole = "detUser"   // 部门负责人
	p_adminUser ProgrammeUserRole = "adminUser" // 分管领导
)

func (this ProgrammeState) Has() (bool, string) {
	msg := ""
	hasState := false
	switch this {
	case p_draft, p_report, p_dep_adopt, p_dep_reject, p_admin_adopt, p_admin_reject:
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
	case p_detUser, p_adminUser:
		hasState = true
	default:
		hasState = false
		msg = config.RoleStr + config.NoHad
	}
	return hasState, msg
}
