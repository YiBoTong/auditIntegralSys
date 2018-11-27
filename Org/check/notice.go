package check

import (
	"auditIntegralSys/_public/config"
)

type NoticeState string

const (
	Draft   NoticeState = "draft"   // 草稿
	Publish NoticeState = "publish" // 发布
)

func (this NoticeState) HasState() (bool, string) {
	msg := ""
	hasState := false
	switch this {
	case Draft, Publish:
		hasState = true
	default:
		hasState = false
		msg = config.StateStr + config.NoHad
	}
	return hasState, msg
}
