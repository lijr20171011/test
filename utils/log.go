package utils

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	defaultFormat = "060102"
	defaultPath   = "./lijrlog/"
)

type MyLog struct {
	fileFormat      string // 日志文件时间格式
	fileName        string // 日志文件名称
	dirPath         string // 日志文件所在目录
	hasPrefix       bool   // 是否需要前缀信息(默认需要)
	isParamsNewLine bool   // 参数之间是否换行(默认不换行)
	mutex           sync.Mutex
	file            *os.File
}

func LogInit(v ...string) *MyLog {
	Info("log init")
	// 获取日志文件格式
	fileFormat := defaultFormat
	if len(v) > 0 {
		fileFormat = v[0]
	}
	// 获取文件所在目录
	dirPath := defaultPath
	if len(v) > 1 {
		dirPath = v[1]
	}
	// 获取是否需要前缀(时间,文件,行号)信息
	hasPrefix := true
	if len(v) > 2 && (v[2] == "false" || v[2] == "False" || v[2] == "f" || v[2] == "F" || v[2] == "0") {
		hasPrefix = false
	}
	isParamsNewLine := false
	if len(v) > 3 && (v[3] == "true" || v[3] == "True" || v[3] == "t" || v[3] == "T" || v[3] == "1") {
		isParamsNewLine = true
	}
	// 获取日志文件名称
	fileName := time.Now().Format(fileFormat)
	// 打开日志文件(不存在则新建)
	file := createOrOpenFile(dirPath, fileName)
	return &MyLog{
		fileFormat:      fileFormat,
		fileName:        fileName,
		file:            file,
		dirPath:         dirPath,
		hasPrefix:       hasPrefix,
		isParamsNewLine: isParamsNewLine,
	}
}

// 打开日志文件(不存在则新建)
func createOrOpenFile(dirPath, fileName string) *os.File {
	dirPath += time.Now().Format("200601/")
	filePath := dirPath + fileName + ".log"
	os.MkdirAll(dirPath, os.ModePerm)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil && os.IsNotExist(err) {
		file, err = os.Create(filePath)
	}
	IsErr(err)
	return file
}

// 将日志内容写入日志文件
func (l *MyLog) Println(v ...interface{}) {
	// 检查文件与日期是否符合
	if l.fileName != time.Now().Format(l.fileFormat) {
		// 修正文件信息
		l.checkFileDate()
	}
	// 输出到文件
	l.println(2, v...)
}

func (l *MyLog) checkFileDate() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	// 判断日期是否已处理
	if l.fileName == time.Now().Format(l.fileFormat) {
		return
	}
	l.file.Close()
	l.fileName = time.Now().Format(l.fileFormat)
	l.file = createOrOpenFile(l.dirPath, l.fileName)
}

// 将日志内容写入日志文件
func (l *MyLog) println(step int, v ...interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	data := &[]byte{}
	// 是否需要前缀
	if l.hasPrefix {
		now := time.Now().Format("2006-01-02 15:04:05")
		_, filePath, line, ok := runtime.Caller(step)
		if !ok {
			filePath = "????"
			line = 0
		}
		//获取文件名
		_, file := path.Split(filePath)
		// 整理前缀信息
		*data = append(*data, (now + " " + "[" + file + ":" + strconv.Itoa(line) + "]\n")...)
	}
	// 日志内容格式
	format := ""
	// 是否需要换行
	if l.isParamsNewLine {
		format = strings.Repeat("%v\n", len(v))
	} else {
		format = strings.Repeat("%v ", len(v))
	}
	*data = append(*data, fmt.Sprintf(format, v...)...)
	*data = append(*data, []byte{'\n', '\n'}...)
	_, err := l.file.Write(*data)
	IsErr(err)
}
