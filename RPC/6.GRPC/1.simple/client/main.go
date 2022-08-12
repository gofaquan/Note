package main

import (
	"context"
	"fmt"

	"github.com/gofaquan/server/pb"
	"google.golang.org/grpc"
)

func main() {
	// 第一步, 建立网络连接
	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	// gprc为我们生成一个客户端调用服务端的SDK
	client := pb.NewHelloServiceClient(conn)
	resp, err := client.Hello(context.Background(), &pb.Request{Value: "bob"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
