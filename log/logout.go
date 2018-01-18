package log

import (
	"log"
	"os"
	"io"
)

var (
	Trace *log.Logger //说明日志
	Out   *log.Logger //数据打印
	Error *log.Logger //错误提示
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(io.MultiWriter(file, os.Stdout), "Trace:", log.Ldate|log.Ltime|log.Lshortfile)
	Out = log.New(io.MultiWriter(file, os.Stdout), "Out:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stdout), "Error:", log.Ldate|log.Ltime|log.Lshortfile)
}
