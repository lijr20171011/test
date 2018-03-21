package utils

import (
	"fmt"
	"path"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func Info(v ...interface{}) {
	format := strings.Repeat("%v ", len(v))
	msg := fmt.Sprintf(format, v...)
	info(2, msg)
}

func ExitWithErr(err error) {
	if err != nil {
		msg := fmt.Sprintf("错误信息: %v ", err)
		info(2, msg)
		syscall.Exit(557)
	}
}

func PrintWithErr(err error) {
	if err != nil {
		msg := fmt.Sprintf("错误信息: %v ", err)
		info(2, msg)
	}
}

func info(step int, msg string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	_, filePath, line, ok := runtime.Caller(step)
	if !ok {
		filePath = "????"
		line = 0
	}
	//获取文件名
	_, file := path.Split(filePath)
	fmt.Println(now + " " + "[" + file + ":" + strconv.Itoa(line) + "]" + msg)
}
