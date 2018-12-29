package util

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/encoding/gjson"
	"gitee.com/johng/gf/g/net/ghttp"
	"gitee.com/johng/gf/g/util/gconv"
	"reflect"
	"regexp"
	"strings"
	"time"
)

var UrlMsgStr map[string]string

// 转驼峰命名法正则规则
var camelCaseReg = regexp.MustCompile(`_\w`)
// sql中的关键字
var sqlKeyReg = regexp.MustCompile(`^(key|order)$`)

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
	temp["state"] = "变更$状态"
	temp["admin_examine"] = "分管领导审核$"
	temp["dep_examine"] = "部门负责人审核$"
	temp["select"] = "获取$选择列表"
	temp["login"] = "登录系统"
	temp["logout"] = "退出系统"
	temp["upload"] = "上传"
	temp["all"] = "获取全部$"
	temp["search"] = "$搜索"
	temp["title"] = "$搜索"
	temp["read"] = "$已读"
	temp["read"] = "$已读"
	temp["edit_number"] = "$填写文件编号"
	temp["edit_score"] = "$填写稽核积分"
	temp["edit_author"] = "$签署"
	temp["systemSetup/dictionaries"] = "字典"
	temp["systemSetup/log"] = "日志"
	temp["systemSetup/login"] = "人员"
	temp["systemSetup/menu"] = "菜单"
	temp["systemSetup/rbac"] = "权限"
	temp["worker/user"] = "登录员工"
	temp["worker/file"] = "文件"
	temp["worker/menu"] = "菜单"
	temp["org/department"] = "部门/结构/网点"
	temp["org/notice"] = "通知公告"
	temp["org/user"] = "人员"
	temp["org/clause"] = "管理办法"
	temp["audit/programme"] = "审核方案"
	temp["audit/draft"] = "工作底稿"
	temp["audit/confirmation"] = "事实确认书"
	temp["audit/punishNotice"] = "惩罚通知书"
	temp["audit/rectify"] = "整改通知"
	temp["audit/rectifyReport"] = "整改报告"
	temp["audit/integral"] = "积分"
	temp["audit/statistical"] = "统计分析"
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

// 获取请求URL
func GetSqlLogURLByRequest(r *ghttp.Request) string {
	return strings.Split(r.RequestURI, "?")[0]
}

// 获取请求方式
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

// 获取get请求参数
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

// 根据请求获取userId
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

// 转换为驼峰命名
func CamelCase(str string) string {
	return camelCaseReg.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Split(s, "_")[1])
	})
}

// 请求json转换为sql语句map
// gMap 传引用
// key 冒号右侧映射给左侧
// typeStr 冒号右侧的类型转换为左侧类型（nowTime为当前时间字符串）
func GetSqlMapByReqJson(gMap g.Map, reqData gjson.Json, key string, typeStr string) {
	srcKey := strings.Split(key, ":")      // 冒号左边是sql对应的字段，右边是json字段
	srcType := strings.Split(typeStr, ":") // 冒号左边是目标类型，右边是json传的类型
	sqlKey := srcKey[0]
	sqlType := srcType[0]
	jsonKey := sqlKey
	jsonType := sqlType
	var val interface{}
	if sqlType == "nowTime" { // 当前字符串
		val = GetLocalNowTimeStr()
	} else {
		if len(srcKey) > 1 {
			jsonKey = srcKey[1] // json中的键
		}
		val = reqData.Get(CamelCase(jsonKey)) // 获取到json中对应的值
		if len(srcType) > 1 {
			jsonType = srcType[1]
			val = gconv.Convert(val, jsonType) // json中传递的数据类型
		}
		val = gconv.Convert(val, sqlType) // sql保存时的数据类型
	}
	if sqlKey != "id" || (sqlKey == "id" && val != 0) { // 过滤id为0
		gMap[sqlKey] = val
	}
}

// 请求json转换为搜索语句map
// gMap 传引用
// key 冒号右侧映射给左侧
// typeStr 冒号右侧的类型转换为左侧类型（nowTime为当前时间字符串）
func GetSearchMapByReqJson(gMap g.Map, reqData gjson.Json, key string, typeStr string) {
	srcKey := strings.Split(key, ":")      // 冒号左边是sql对应的字段，右边是json字段
	srcType := strings.Split(typeStr, ":") // 冒号左边是目标类型，右边是json传的类型
	sqlKey := srcKey[0]
	sqlType := srcType[0]
	jsonKey := sqlKey
	jsonType := sqlType
	var val interface{}
	if len(srcKey) > 1 {
		jsonKey = srcKey[1] // json中的键
	}
	val = reqData.Get(CamelCase(jsonKey)) // 获取到json中对应的值
	if val != nil && val != "" {
		if len(srcType) > 1 {
			jsonType = srcType[1]
			val = gconv.Convert(val, jsonType) // json中传递的数据类型
		}
		if strings.Index(sqlKey, " like ") > 0 {
			val = strings.Replace("%?%", "?", gconv.String(val), 1)
		}
		val = gconv.Convert(val, sqlType) // sql保存时的数据类型
		gMap[sqlKeyReg.ReplaceAllString(sqlKey, "`$1`")] = val
	}
}

// 获取sql map
func GetSqlMap(reqJson gjson.Json, reqDataMap map[string]interface{}, sqlMap g.Map) {
	for k, v := range reqDataMap {
		GetSqlMapByReqJson(sqlMap, reqJson, k, gconv.String(v))
	}
}

func GetSqlMapItemFun(json gjson.Json, reqDataMap map[string]interface{}, appendFun func(itemMap g.Map)) {
	if dataLen := len(json.ToArray()); dataLen > 0 {
		for index := 0; index < dataLen; index++ {
			item := g.Map{}
			reqBasisItem := json.GetJson(gconv.String(index))
			for k, v := range reqDataMap {
				GetSqlMapByReqJson(item, *reqBasisItem, k, gconv.String(v))
			}
			appendFun(item)
		}
	}
}
