package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gofaquan/server/pb"
	"google.golang.org/grpc"
)

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
// type HelloServiceServer interface {
// 	Hello(context.Context, *Request) (*Response, error)
// 	mustEmbedUnimplementedHelloServiceServer()
// }

// HelloServiceServer must be embedded to have forward compatible implementations.
type HelloServiceServer struct {
	pb.UnimplementedHelloServiceServer
}

// alice  <-- hello, alice

func (s *HelloServiceServer) Hello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Value: fmt.Sprintf("hello, %s", req.Value)}, nil
}

func main() {
	// s grpc.ServiceRegistrar  grpc Server
	// srv HelloServiceServer   实现类
	server := grpc.NewServer()

	// 把实现类注册给Grpc Server
	pb.RegisterHelloServiceServer(server, new(HelloServiceServer))

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	log.Printf("gprc listen addr: 127.0.0.1:8888")
	// 监听Socket, HTTP2内置
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
