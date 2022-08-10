package main

import (
	"fmt"
	service "github.com/gofaquan/Interface"
	"log"
	"net"
	"net/rpc"
)

// 约束服务端接口的实现
var _ service.HelloSvcInterface = (*HelloSvcServer)(nil)

// HelloSvcServer service handler
type HelloSvcServer struct {
}

// Hello
// request: 请求
// response 响应
// request  -->   name
// response <--  hello, name
func (s *HelloSvcServer) Hello(request string, response *string) error {
	*response = fmt.Sprintf("hello, %s", request)
	return nil
}

// main 里面编写Server
func main() {
	// 把rpc对外暴露的对象注册到rpc框架内部
	rpc.RegisterName(service.SvcName, &HelloSvcServer{})

	// 准备socket
	// 然后我们建立一个唯一的TCP链接，
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	// 获取连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		// 每个客户端单独启用一个routine来处理
		go rpc.ServeConn(conn)
	}

}
