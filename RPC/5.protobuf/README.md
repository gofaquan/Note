# 如何生成代码
### 下载 protoc
https://github.com/protocolbuffers/protobuf
### proto-gen-go
https://grpc.io/docs/languages/go/quickstart/

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
```sh
# 在 protobuf/pb 文件夹下面对应
$ protoc -I=. --go_out=. --go_opt=module="github.com/gofaquan/protobuf/pb"   ./hello.proto

# 在 protobuf 文件夹下面对应
$ protoc -I=./pb/ --go_out=./pb/ --go_opt=module="github.com/gofaquan/protobuf/"  hello.proto
```