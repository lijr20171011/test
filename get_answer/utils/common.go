package utils

import (
	"fmt"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//打印(文件,行号,时间)
func Info(v ...interface{}) {
	format := strings.Repeat("%v ", len(v))
	msg := fmt.Sprintf(format, v...)
	info(msg, "[I]")
}

func info(msg string, printType string) {
	when := time.Now().Format("2006-01-02 15:04:05")
	_, filePath, line, ok := runtime.Caller(2)
	if !ok {
		filePath = "????"
		line = 0
	}
	//获取文件名
	_, file := path.Split(filePath)
	fmt.Println(when + " " + printType + " [" + file + ":" + strconv.Itoa(line) + "] " + msg)
}
