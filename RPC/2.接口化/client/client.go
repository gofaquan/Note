package main

import (
	"fmt"
	"net/rpc"

	service "github.com/gofaquan/Interface"
)

// 约束客户端接口的实现
var _ service.HelloSvcInterface = (*HelloSvcClient)(nil)

func NewHelloServiceClient() (*HelloSvcClient, error) {
	// 1. 建立socket连接
	client, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		return nil, err
	}
	return &HelloSvcClient{
		client: client,
	}, nil
}

type HelloSvcClient struct {
	client *rpc.Client
}

func (c *HelloSvcClient) Hello(request string, response *string) error {
	// 执行RPC方法
	return c.client.Call(fmt.Sprintf("%s.Hello", service.SvcName), request, response)
}

func main() {
	// 创建客户端
	c, err := NewHelloServiceClient()
	if err != nil {
		panic(err)
	}

	var resp string
	if err := c.Hello("gofaquan", &resp); err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
