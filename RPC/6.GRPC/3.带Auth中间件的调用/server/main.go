package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/gofaquan/middleware/server"
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

func (s *HelloServiceServer) Channel(stream pb.HelloService_ChannelServer) error {
	for {
		// 接收请求
		req, err := stream.Recv()
		if err != nil {
			// 当前客户端退出
			log.Printf("recv error, %s", err)
			if err == io.EOF {
				log.Printf("client closed")
				return nil
			}
			return err
		}

		resp := &pb.Response{Value: fmt.Sprintf("hello, %s", req.Value)}
		// 响应请求
		err = stream.Send(resp)
		if err != nil {
			if err == io.EOF {
				log.Printf("client closed")
				return nil
			}
			return err
		}
	}
}

func main() {
	// s grpc.ServiceRegistrar  grpc Server
	// srv HelloServiceServer   实现类

	// 添加认证中间件
	reqAuth := server.NewAuthUnaryServerInterceptor()
	streamAuth := server.NewAuthStreamServerInterceptor()
	server := grpc.NewServer(
		grpc.UnaryInterceptor(reqAuth),
		grpc.StreamInterceptor(streamAuth),
	)

	// 把实现类注册给Grpc Server
	pb.RegisterHelloServiceServer(server, new(HelloServiceServer))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	log.Printf("gprc listen addr: 127.0.0.1:1234")
	// 监听Socket, HTTP2内置
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
