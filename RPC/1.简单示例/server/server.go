package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type HelloSvc struct{}

// request: 请求
// response 响应
// request  发送 name
// response 返回 hello, name

// Hello 将对方发送的消息前面添加一个Hello 然后返还给对方
// 由于我们是一个rpc服务, 因此参数上面还是有约束：
// 		第一个参数是请求
// 		第二个参数是响应
// 		返回参数需要一个 error
// 可以类比Http handler
func (s *HelloSvc) Hello(request string, response *string) error {
	*response = fmt.Sprintf("hello, %s", request)
	return nil
	// 如果没有 返参数 error，会出现类型不匹配的 warning
	// 2022/08/10 11:31:14 rpc.Register: type HelloSvc has no exported methods of suitable type
}

func main() {
	// 把rpc对外暴露的对象注册到rpc框架内部
	rpc.RegisterName("HelloSvc", &HelloSvc{})

	// 准备socket 然后我们建立一个唯一的TCP链接
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	// 获取连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		// 每个客户端单独启用一个routine来处理
		go rpc.ServeConn(conn)
	}
}
