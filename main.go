package main

import (
	"./request"
	"net/http"
)

func main() {
	http.HandleFunc("/notepad/", request.Notepad)
	http.ListenAndServe(":4000", nil)
}