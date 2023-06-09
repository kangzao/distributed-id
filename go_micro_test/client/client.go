package main

import (
	"context"
	pb "distributed-id/go_micro_test/protos"
	"fmt"
	"go-micro.dev/v4"
)

func main() {
	service := micro.NewService(micro.Name("greeter"))
	service.Init()

	greeter := pb.NewGreeterService("greeter", service.Client())
	rsp, err := greeter.Hello(context.TODO(), &pb.HelloRequest{Name: "pingye"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Greeting)
}
