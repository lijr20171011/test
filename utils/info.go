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

// 打印参数
func Info(v ...interface{}) {
	format := strings.Repeat("%v ", len(v))
	msg := fmt.Sprintf(format, v...)
	_info(2, msg)
}

// 根据格式打印参数
func Infof(format string, v ...interface{}) {
	// format := strings.Repeat("%v ", len(v))
	msg := fmt.Sprintf(format, v...)
	_info(2, msg)
}

// // 检查错误,如果err不为空则打印返回true
// func IsErr(err error) bool {
// 	if err != nil {
// 		msg := fmt.Sprintf("错误信息: %v ", err)
// 		_info(2, msg)
// 		return true
// 	}
// 	return false
// }

// 检查错误,如果err不为空则打印返回true
// errInfo : 附加错误信息
func IsErr(err error, errInfo ...interface{}) bool {
	if err != nil {
		format := "错误信息:"
		if len(errInfo) > 0 {
			format += strings.Repeat(" %v >>>", len(errInfo))
		}
		errInfo = append(errInfo, err)
		msg := fmt.Sprintf(format+" %v ", errInfo...)
		_info(2, msg)
		return true
	}
	return false
}

// 检查错误,如果err不为空则打印并退出
func ExitWithErr(err error) bool {
	if err != nil {
		msg := fmt.Sprintf("错误信息: %v ", err)
		_info(2, msg)
		syscall.Exit(557)
		return true
	}
	return false
}

func _info(step int, msg string) {
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
