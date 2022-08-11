package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/gofaquan/service"
)

// 约束服务端接口的实现
var _ service.HelloService = (*HelloService)(nil)

type HelloService struct {
}

func (s *HelloService) Hello(request string, response *string) error {
	*response = fmt.Sprintf("hello, %s", request)
	return nil
}

func (s *HelloService) Calc(req *service.CalcRequest, response *int) error {
	*response = req.A + req.B
	return nil
}

func NewRPCReadWriteCloser(w http.ResponseWriter, r *http.Request) *RPCReadWriteCloser {
	return &RPCReadWriteCloser{w, r.Body}
}

type RPCReadWriteCloser struct {
	io.Writer
	io.ReadCloser
}

// main 里面编写Server
func main() {
	// 把rpc对外暴露的对象注册到rpc框架内部
	rpc.RegisterName(service.SERVICE_NAME, &HelloService{})

	// 通过 jsonrpc 这个Path 来处理所有的请求
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		rpc.ServeCodec(jsonrpc.NewServerCodec(NewRPCReadWriteCloser(w, r)))
	})

	// 通过HTTP协议接收rpc请求
	http.ListenAndServe(":8080", nil)
}

// POST localhost:8080/jsonrpc
//{
//    "method": "HelloService.Calc",
//    "params": [{"a":1, "b": 2}]
//}
