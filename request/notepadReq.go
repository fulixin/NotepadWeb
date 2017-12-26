package request

import "net/http"

func Notepad(writer http.ResponseWriter, request *http.Request) {
	var code=200
	url:=request.URL.Path
	seei:=request.Header.Get("")
	writer.WriteHeader(code)
	writer.Write([]byte(url))
}
