package main

import (
	"net/http"
	"./request"
)

func main() {
	http.HandleFunc("/notepad/", request.Notepad)
	http.ListenAndServe(":4000",nil)
}
