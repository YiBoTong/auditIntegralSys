package check

import "auditIntegralSys/_public/config"

type ClauseState string

const (
	c_draft   ClauseState = "draft"   // 草稿
	c_publish ClauseState = "publish" // 发布
)

func (this ClauseState) HasState() (bool, string) {
	msg := ""
	hasState := false
	switch this {
	case c_draft, c_publish:
		hasState = true
	default:
		hasState = false
		msg = config.StateStr + config.NoHad
	}
	return hasState, msg
}
