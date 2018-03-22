package main

import (
	"time"

	"./client"

	"./server"

	"../utils"
)

func main() {
	server.MyServerTest()
	utils.Info("服务器启动")
	time.Sleep(time.Second)
	utils.Info("客户端启动")
	client.MyClientTest()
	utils.Info("end")
}
