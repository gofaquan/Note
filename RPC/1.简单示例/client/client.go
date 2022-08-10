package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 然后通过client.Call调用具体的RPC方法
	// 在调用client.Call时:
	// 		第一个参数是用点号链接的RPC服务名字和方法名字，
	// 		第二个参数是 请求参数
	//      第三个是请求响应, 必须是一个指针, 有底层rpc服务帮你赋值

	// 建立socket连接
	client, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	// 执行RPC方法
	var response string
	request := "我是请求消息"

	client.Call("HelloSvc.Hello", request, &response)

	// 打印调用响应
	fmt.Println(response)
}
