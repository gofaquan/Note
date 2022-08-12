package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/gofaquan/middleware/server"
	"github.com/gofaquan/server/pb"
)

// go 主go, 主线程
func main() {
	// 第一步, 建立网络连接
	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	// gprc为我们生成一个客户端调用服务端的SDK
	client := pb.NewHelloServiceClient(conn)

	// req <--> resp
	// 添加认证凭证信息
	crendential := server.NewClientCredential("admin", "123456")
	ctx := metadata.NewOutgoingContext(context.Background(), crendential)
	resp, err := client.Hello(ctx, &pb.Request{Value: "bob"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	// stream
	stream, err := client.Channel(ctx)
	if err != nil {
		panic(err)
	}

	// 启用一个Goroutine来发送请求
	go func() {
		for {
			err := stream.Send(&pb.Request{Value: "alice"})
			if err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// 主循环 负责接收服务端响应
	for {
		resp, err = stream.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	}
}
