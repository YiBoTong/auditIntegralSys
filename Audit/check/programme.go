package check

import "auditIntegralSys/_public/config"

type DraftState string
type ProgrammeState string
type ProgrammeUserRole string
type ConfirmationState string

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

	p_draft  ProgrammeState = "draft"  // 草稿
	p_report ProgrammeState = "report" // 上报
	p_adopt  ProgrammeState = "adopt"  // 通过
	p_reject ProgrammeState = "reject" // 驳回
)

const (
	d_draft   DraftState = "draft"   // 草稿
	d_report  DraftState = "report"  // 上报
	d_adopt   DraftState = "adopt"   // 通过
	d_reject  DraftState = "reject"  // 驳回
	d_publish DraftState = "publish" // 发布

	D_draft   string = "draft"   // 草稿
	D_report  string = "report"  // 上报
	D_adopt   string = "adopt"   // 通过
	D_reject  string = "reject"  // 驳回
	D_publish string = "publish" // 发布
)

const (
	c_draft   ConfirmationState = "draft"   // 草稿
	c_publish ConfirmationState = "publish" // 发布
)

const (
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

func (this DraftState) Has() (bool, string) {
	msg := ""
	hasState := false
	switch this {
	// 草稿、发布
	case d_draft, d_publish:
		hasState = true
	default:
		hasState = false
		msg = config.StateStr + config.NoHad
	}
	return hasState, msg
}

func (this ConfirmationState) Has() (bool, string) {
	msg := ""
	hasState := false
	switch this {
	// 草稿、发布
	case c_draft, c_publish:
		hasState = true
	default:
		hasState = false
		msg = config.StateStr + config.NoHad
	}
	return hasState, msg
}