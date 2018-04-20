package main

import (
	"my_project/test/utils"
	"time"

	"github.com/astaxie/beego/toolbox"
)

func main() {
	// taskTest()
	// time.Sleep(1 * time.Minute)
}

func taskTest() {
	func1 := toolbox.NewTask("func1", "*/1 * * * * *", func1)
	toolbox.AddTask("func1", func1)
	toolbox.StartTask()
}

func func1() error {
	utils.Info("func1====")
	time.Sleep(2 * time.Second)
	utils.Info("sleep over")
	return nil
}
