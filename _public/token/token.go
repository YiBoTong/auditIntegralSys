package token

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
	"gitee.com/johng/gf/g/os/gtime"
	"gitee.com/johng/gf/g/util/grand"
	"strconv"
	"strings"
)

func Make() string {
	return strings.ToUpper(strconv.FormatInt(gtime.Nanosecond(), 32) + grand.RandStr(5))
}

func Set(userId int, r *ghttp.Request,isLogin bool) {
	token := ""
	if isLogin {
		token = Make()
	}else {
		token = r.Cookie.Get(config.CookieIdName)
	}
	_, _ = g.Redis().Do("Set", token, userId)
	_, _ = g.Redis().Do("Expire", token, config.CookieMaxAge)
	r.Cookie.SetCookie(config.CookieIdName, token, "", "/", config.CookieMaxAge, true)
}

func Del(r *ghttp.Request) {
	_, _ = g.Redis().Do("Del", r.Cookie.Get(config.CookieIdName))
	r.Cookie.Remove(config.CookieIdName, "", "/")
}
