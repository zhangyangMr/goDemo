# grpc入门教程

## 1. grpc官网网址

github: https://github.com/grpc/grpc-go

说明文档：https://grpc.io/docs/languages/go/quickstart/

## 2. grpc相关工具的安装

* golang的安装，自行解决

* grpc的安装

  ```shell
  $ export GO111MODULE=on  # Enable module mode
  $ go get google.golang.org/protobuf/cmd/protoc-gen-go \
           google.golang.org/grpc/cmd/protoc-gen-go-grpc
  # 查看到gopath/bin下面已有grpc相关二进制工具       
  $ export PATH="$PATH:$(go env GOPATH)/bin"
  ```

## 3. 测试代码

参考：grpc-go/examples/helloworld即可

## 4. proto文件编写

```shell
syntax = "proto3";

option go_package = "helloworld/helloworld";

package helloworld;

// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
  // Sends another greeting
  rpc SayHelloAgain (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```

## 5. 生成pb.go文件

```shell
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./helloworld.proto
source_relative：当前proto所在的目录
```

## 6. grpc server代码方法重写介绍

```shell
// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

# 说明：当前函数是使用的方法重写，server结构体是一个父类(里面包含了匿名变量	pb.UnimplementedGreeterServer)，pb.go中的UnimplementedGreeterServer是server的一个子类，父类和子类都有SayHello的方法实现，调用父类肯定是先使用父类，父类没有调用子类
```

