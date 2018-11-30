package check

import (
	"auditIntegralSys/_public/config"
)

type NoticeState string

const (
	n_draft   NoticeState = "draft"   // 草稿
	n_publish NoticeState = "publish" // 发布
)

func (this NoticeState) HasState() (bool, string) {
	msg := ""
	hasState := false
	switch this {
	case n_draft, n_publish:
		hasState = true
	default:
		hasState = false
		msg = config.StateStr + config.NoHad
	}
	return hasState, msg
}
