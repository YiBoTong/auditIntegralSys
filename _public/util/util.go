package util

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
	"gitee.com/johng/gf/g/util/gconv"
	"reflect"
	"strings"
	"time"
)

var UrlMsgStr map[string]string

func init() {
	initUrlMsgStr()
}

func initUrlMsgStr() {
	var temp = make(map[string]string)
	temp["list"] = "获取$列表"
	temp["add"] = "添加"
	temp["edit"] = "编辑"
	temp["get"] = "获取"
	temp["delete"] = "删除"
	temp["tree"] = "获取"
	temp["is-use"] = "变更$状态"
	temp["login"] = "登录系统"
	temp["logout"] = "退出系统"
	temp["upload"] = "上传"
	temp["password"] = "密码修改"
	temp["systemSetup/dictionaries"] = "字典"
	temp["systemSetup/log"] = "日志"
	temp["systemSetup/login"] = "人员"
	temp["worker/user"] = ""
	temp["worker/file"] = "文件"
	temp["org/department"] = "部门/结构/网点"
	temp["org/notice"] = "通知公告"
	temp["org/user"] = "人员"
	UrlMsgStr = temp
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// 获取当前时间字符串
func GetLocalNowTimeStr() string {
	localTime := time.Now().Format("2006-01-02 15:04:05")
	return localTime
}

func GetSqlLogURLByRequest(r *ghttp.Request) string {
	return strings.Split(r.RequestURI, "?")[0]
}

func GetSqlLogMsgByRequest(r *ghttp.Request) string {
	urlArr := strings.Split(r.RequestURI, "/")
	urlLen := len(urlArr)
	msg := UrlMsgStr[strings.Split(urlArr[len(urlArr)-1], "?")[0]]
	if strings.Index(msg, "$") > -1 {
		msg = strings.Replace(msg, "$", UrlMsgStr[urlArr[urlLen-3]+"/"+urlArr[urlLen-2]], 1)
	} else {
		msg += UrlMsgStr[urlArr[urlLen-3]+"/"+urlArr[urlLen-2]]
	}
	return msg
}

func GetSqlLogDataByRequest(r *ghttp.Request) string {
	method := r.Request.Method
	data := ""
	switch method {
	case "POST", "PUT":
		data = "-"
	default:
		data = gconv.String(strings.Split(r.Request.RequestURI, "?")[1])
	}
	return data
}

func GetUserIdByRequest(r *ghttp.Cookie) int {
	userId := 0
	token := r.Get(config.CookieIdName)
	hasToken, _ := g.Redis().Do("EXISTS", token)
	if gconv.Bool(hasToken) {
		res, _ := g.Redis().Do("GET", token)
		userId = gconv.Int(res)
	}
	return userId
}
