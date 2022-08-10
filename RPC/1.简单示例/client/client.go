package main

import (
	"fmt"
	"net/rpc"
)

func main() {
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
