package server

import (
	"errors"
	"net"
	"net/http"
	"net/rpc"

	"../../utils"

	"../models"
)

// 辅助结构
type Arith int

// 乘法
func (t *Arith) Multiply(args *models.Args, reply *int) (err error) {
	*reply = args.A * args.B
	return
}

// 除法
func (t *Arith) Divide(args *models.Args, reply *models.Quotient) (err error) {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	reply.Quo = args.A / args.B
	reply.Rem = args.A % args.B
	return
}

// 服务器
func MyServerTest() {
	arith := new(Arith)
	rpc.Register(arith)                                // 注册服务
	rpc.HandleHTTP()                                   // 通过http暴露
	listen, err := net.Listen("tcp", "127.0.0.1:1234") // 监听端口
	if utils.IsErr(err, "listen error:") {
		return
	}
	go http.Serve(listen, nil) // 启动服务
}
