package main

import (
	"fmt"
	"github.com/gofaquan/service"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 约束客户端接口的实现
var _ service.HelloService = (*HelloServiceClient)(nil)

func NewHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	// 1. 建立socket连接
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	// 客户端采用Json 格式来编解码
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	if err != nil {
		return nil, err
	}

	return &HelloServiceClient{
		client: client,
	}, nil
}

type HelloServiceClient struct {
	client *rpc.Client
}

func (c *HelloServiceClient) Hello(request string, response *string) error {
	// 执行RPC方法
	return c.client.Call(fmt.Sprintf("%s.Hello", service.SERVICE_NAME), request, response)
}

func main() {
	// 创建客户端
	c, err := NewHelloServiceClient("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	var resp string
	if err := c.Hello("我是请求内容", &resp); err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
