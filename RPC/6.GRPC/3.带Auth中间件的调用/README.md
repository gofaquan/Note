# Hello world Grpc

```sh
> cd go-course-projects/go7/rpc-g7/grpc/simple/server
# 生成 serivce pb编译文件
> protoc -I=. --go_out=. --go_opt=module="gitee.com/go-course/rpc-g7/grpc/simple/server" pb/hello.proto



# 补充rpc 接口定义protobuf的代码生成
protoc -I=. --go_out=. --go_opt=module="gitee.com/go-course/rpc-g7/grpc/simple/server" --go-grpc_out=. --go-grpc_opt=module="gitee.com/go-course/rpc-g7/grpc/simple/server" pb/hello.proto
```