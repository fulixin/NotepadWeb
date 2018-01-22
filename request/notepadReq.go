package request

import (
	"net/http"
	"fmt"
	dataUtil "../util"
	logmy "../log"
	"../appinface/register"
	"../appinface/note"
	"encoding/json"
)

//定义路由器结构类型
type Routers struct {
}

/**
注册用户
 */
func (r *Routers) Register(i map[string]string, writer http.ResponseWriter) {
	result := make(map[string]interface{})
	u := register.Register{
	}
	u.Username = i["username"]
	u.Password = i["password"]
	err := register.UserRegister(&u)
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
获取日志列表
 */
func (r *Routers) QueryNoteList(i map[string]string, writer http.ResponseWriter) {
	result := make(map[string]interface{})
	notes, err := note.QueryNoteList()
	if err != nil {
		writer.WriteHeader(dataUtil.Loginerror)
		result["result"] = err.Error()
		result["resultMap"] = ""
	} else {
		writer.WriteHeader(dataUtil.Success)
		resultMap := make(map[string]interface{})
		resultMap["List"] = dataUtil.StructsToSlisp(notes)
		result["result"] = "success"
		result["resultMap"] = resultMap
		logmy.Trace.Println(resultMap)
	}
	s, errs := json.Marshal(result)
	if errs != nil {
		fmt.Println(errs)
		return
	}
	writer.Write(s)
}

/**
保存日志
 */
func (r *Routers) SaveNote(i map[string]string, writer http.ResponseWriter) {
	result := make(map[string]interface{})
	n := note.Note{}
	n.NoteDate = i["NoteDate"]
	n.NoteContext = i["NoteContext"]
	n.NoteTitle = i["NoteTitle"]
	err := note.SaveNote(&n)
	if err != nil {
		writer.WriteHeader(dataUtil.Loginerror)
		result["result"] = err.Error()
		result["resultMap"] = ""
	} else {
		writer.WriteHeader(dataUtil.Success)
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
