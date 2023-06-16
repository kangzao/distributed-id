package main

import (
	"context"
	"fmt"

	"distributed-id/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var generateIdClient proto.GenerateIdClient
var conn *grpc.ClientConn

func Init() {
	var err error
	//这里不能用冒号，要给conn这个全局变量赋值
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	generateIdClient = proto.NewGenerateIdClient(conn)
}

func TestGenerateId() {
	rsp, err := generateIdClient.GetIds(context.Background(), &proto.GetIdRequest{Stub: "C"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", rsp)
}

func main() {
	Init()
	TestGenerateId()
	err := conn.Close()
	if err != nil {
		return
	}

}
