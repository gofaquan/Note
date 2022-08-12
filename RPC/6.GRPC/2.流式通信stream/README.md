# Hello world Grpc

```sh
> cd 6.GRPC/2.流式通信stream/server
# 生成 serivce pb编译文件
> protoc -I=. --go_out=. --go_opt=module="github.com/gofaquan/server" pb/hello.proto



# 补充rpc 接口定义protobuf的代码生成
protoc -I=. --go_out=. --go_opt=module="github.com/gofaquan/server" --go-grpc_out=. --go-grpc_opt=module="github.com/gofaquan/server" pb/hello.proto
```