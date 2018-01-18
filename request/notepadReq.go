package request

import (
	"net/http"
	"fmt"
	"net/url"
	dataUtil "../util"
	logmy "../log"
	"../appinface/register"
	"../appinface/login"
	"reflect"
	"strings"
	"encoding/json"
)

//定义控制器函数Map类型，便于后续快捷使用
type ControllerMapsType map[string]reflect.Value

//声明控制器函数Map类型变量
var ControllerMaps ControllerMapsType

//定义路由器结构类型
type Routers struct {
}

func Notepad(writer http.ResponseWriter, request *http.Request) {
	dataMap, err := parameUtil(request.URL)
	if err != nil {
		logmy.Error.Println("解析发送参数错误:", err)
	}
	var ruTest Routers
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

/**
获取调用方法名称
 */
func getMethodName(str string) string {
	str = strings.Replace(str, "/notepad/", "", 1)
	str = strings.Replace(str, ".action", "", 1)
	return dataUtil.StrFirstToUpper(str)
}

/**
注册用户
 */
func (r *Routers) Register(i map[string]string, writer http.ResponseWriter) {
	result := make(map[string]interface{})
	u := register.Register{
	}
	u.SetUsername(i["username"])
	u.SetPassword(i["password"])
	err := u.Register()
	if err != nil {
		writer.WriteHeader(dataUtil.Registererror)
		result["result"] = err.Error()
		result["resultMap"] = ""
	} else {
		writer.WriteHeader(dataUtil.Success)
		writer.Write([]byte("注册用户成功"))
		result["result"] = "success"
		result["resultMap"] = ""
	}
	s, errs := json.Marshal(result)
	if errs != nil {
		fmt.Println(errs)
		return
	}
	writer.Write(s)
}

/**
登录
 */
func (r *Routers) Login(i map[string]string, writer http.ResponseWriter) {
	result := make(map[string]interface{})
	u := login.Login{}
	u.SetUsername(i["username"])
	u.SetPassword(i["password"])
	err := u.VerifyAccount()
	if err != nil {
		writer.WriteHeader(dataUtil.Loginerror)
		result["result"] = err.Error()
		result["resultMap"] = ""
	} else {
		err := u.GetUser()
		if err != nil {
			writer.WriteHeader(dataUtil.Loginerror)
			result["result"] = err.Error()
			result["resultMap"] = ""
		}
		writer.WriteHeader(dataUtil.Success)
		resultMap := make(map[string]string)
		resultMap["userName"] = u.GetUsername()
		resultMap["userId"] = u.GetUserid()
		result["result"] = "success"
		result["resultMap"] = resultMap
	}
	s, errs := json.Marshal(result)
	if errs != nil {
		fmt.Println(errs)
		return
	}
	writer.Write(s)
}

func parameUtil(u *url.URL) (map[string]string, error) {
	fmt.Println(u.RawQuery)
	m1 := u.Query()
	fmt.Println(m1)
	requestMap := make(map[string]string)
	for k, v := range m1 {
		requestMap[k] = v[0]
	}
	fmt.Println(requestMap["data"])
	dataMap, err := dataUtil.Json2map(requestMap["data"])
	if err != nil {
		println(err)
		return nil, err
	}
	return dataMap, nil
}
