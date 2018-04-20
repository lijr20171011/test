package main

import (
	"my_project/git_projects/test/utils"
)

func main() {
	utils.Info("start")
	mylog.Println("test", "测试")
	utils.Info("end")
}

var mylog *utils.MyLog

func init() {
	mylog = utils.LogInit()
}
