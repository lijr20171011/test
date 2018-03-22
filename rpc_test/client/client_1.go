package client

import (
	"net/rpc"

	"../models"

	"../../utils"
)

// 客户端
func MyClientTest() {
	myClient, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if utils.IsErr(err, "dialing:") {
		return
	}
	args := &models.Args{7, 8}
	// 同步调用
	// reply := 0
	// err = myClient.Call("Arith.Multiply", args, &reply)
	// if utils.IsErr(err, "arith error:") {
	// 	return
	// }
	// utils.Infof("Arith: %d*%d=%d", args.A, args.B, reply)

	// 异步调用
	quotient := new(models.Quotient)
	divCall := myClient.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done
	utils.Info(replyCall.Args)
	utils.Info(replyCall.Reply)
}
