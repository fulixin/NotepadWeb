package main

import (
	"net/http"
	"./helper"
	"reflect"
	"strings"
	"net/url"
	"fmt"
	dataUtil "./util"
	logmy "./log"
	routers "./request"
	"encoding/json"
	"./appinface/login"
)

//定义控制器函数Map类型，便于后续快捷使用
type ControllerMapsType map[string]reflect.Value

//声明控制器函数Map类型变量
var ControllerMaps ControllerMapsType

var sessionMgr *helper.SessionMgr = nil //session管理器

func main() {
	//创建session管理器,”TestCookieName”是浏览器中cookie的名字，3600是浏览器cookie的有效时间（秒）
	sessionMgr = helper.NewSessionMgr("TestCookieName", 3600)
	http.HandleFunc("/notepad/", Notepad)
	http.ListenAndServe(":4000", nil)
}

func Notepad(writer http.ResponseWriter, request *http.Request) {
	dataMap, err := parameUtil(request.URL)
	if err != nil {
		logmy.Error.Println("解析发送参数错误:", err)
	}
	if getMethodName(request.URL.Path) == "Login" {
		Login(dataMap, writer, request)
	} else if getMethodName(request.URL.Path) == "Logout" {
		Logout(writer, request)
	} else {
		var sessionID = sessionMgr.CheckCookieValid(writer, request)
		if sessionID == "" {
			//http.Redirect(writer, request, "/notepad/login.action", http.StatusFound)
			writer.Write([]byte("请登录"))
			return
		}
		var ruTest routers.Routers
		crMap := make(ControllerMapsType, 0)
		//创建反射变量，注意这里需要传入ruTest变量的地址；
		//不传入地址就只能反射Routers静态定义的方法
		vf := reflect.ValueOf(&ruTest)
		vft := vf.Type()
		//读取方法数量
		mNum := vf.NumMethod()
		//遍历路由器的方法，并将其存入控制器映射变量中
		for i := 0; i < mNum; i++ {
			mName := vft.Method(i).Name
			crMap[mName] = vf.Method(i)
		}
		parms := []reflect.Value{reflect.ValueOf(dataMap), reflect.ValueOf(writer)}
		//使用方法名字符串调用指定方法
		crMap[getMethodName(request.URL.Path)].Call(parms)
	}
}

/**
登录
 */
func Login(i map[string]string, writer http.ResponseWriter, request *http.Request) {
	result := make(map[string]interface{})
	u := login.Login{}
	u.Account = i["username"]
	u.Password = i["password"]
	err := login.VerifyAccount(&u)
	if err != nil {
		writer.WriteHeader(dataUtil.Loginerror)
		result["result"] = err.Error()
		result["resultMap"] = ""
	} else {
		err := login.GetUser(&u)
		if err != nil {
			writer.WriteHeader(dataUtil.Loginerror)
			result["result"] = err.Error()
			result["resultMap"] = ""
		}

		//创建客户端对应cookie以及在服务器中进行记录
		var sessionID = sessionMgr.StartSession(writer, request)
		//踢除重复登录的
		var onlineSessionIDList = sessionMgr.GetSessionIDList()
		for _, onlineSessionID := range onlineSessionIDList {
			if userInfo, ok := sessionMgr.GetSessionVal(onlineSessionID, "UserInfo"); ok {
				if value, ok := userInfo.(login.Login); ok {
					if value.Userid == u.Userid {
						sessionMgr.EndSessionBy(onlineSessionID)
					}
				}
			}
		}

		//设置变量值
		sessionMgr.SetSessionVal(sessionID, "UserInfo", u)
		writer.WriteHeader(dataUtil.Success)
		u.Password = ""
		resultMap := dataUtil.StructToMap(u)
		logmy.Trace.Println(resultMap)
		result["resultMap"] = resultMap
		logmy.Trace.Println(resultMap)
	}
	s, _ := json.Marshal(result)
	writer.Write(s)
}

//处理退出
func Logout(writer http.ResponseWriter, r *http.Request) {
	sessionMgr.EndSession(writer, r) //用户退出时删除对应session
	writer.Write([]byte("退出成功"))
	return
}

/**
获取调用方法名称
 */
func getMethodName(str string) string {
	str = strings.Replace(str, "/notepad/", "", 1)
	str = strings.Replace(str, ".action", "", 1)
	return dataUtil.StrFirstToUpper(str)
}

func parameUtil(u *url.URL) (map[string]string, error) {
	fmt.Println(u.RawQuery)
	m1 := u.Query()
	fmt.Println(m1)
	requestMap := make(map[string]string)
	for k, v := range m1 {
		requestMap[k] = v[0]
	}
	logmy.Trace.Println(requestMap["data"])
	dataMap, err := dataUtil.Json2map(requestMap["data"])
	if err != nil {
		println(err)
		return nil, err
	}
	return dataMap, nil
}
