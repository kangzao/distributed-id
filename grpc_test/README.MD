go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/hellogrpc.proto


这将生成 hellogrpc.pb.go 和 hellogrpc_grpc.pb.go 文件，它们包含：
用于填充、序列化和检索 HelloRequest 和 HelloReply 消息类型的代码。
生成的客户端和服务端代码。

运行服务端,在server目录下运行：
go run main.go

从另外一个终端，运行客户端
go run main.go --name=Tom


将看到如下输出：
Greeting: Hello Tom
Greeting: Hello again Tom

hellogrpc.pb.go:用于填充、序列化、获取 message HelloRequest 和 message HelloReply 的代码。

hellogrpc_grpc.pb.go:生成的 client 端和 server 端代码