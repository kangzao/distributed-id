package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"distributed-id/global"
	"distributed-id/handler"
	"distributed-id/initialize"
	"distributed-id/proto"
	"distributed-id/utils"
)

func main() {
	//redisUtil.Set("test", []byte("777"))
	//result, _ := redisUtil.Get("test")
	//fmt.Println(string(result))
	//go build main.go  然后本地执行,手动设置端口号和ip  ./main -ip 127.0.0.1  -port 50051

	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 50051, "端口号")

	initialize.InitLogger()
	initialize.InitConfig()           //配置文件初始化
	zap.S().Info(global.ServerConfig) //打印服务端配置信息

	flag.Parse()
	zap.S().Info("ip: ", *IP)
	if *Port == 0 {
		*Port, _ = utils.GetFreePort()
	}
	zap.S().Info("port: ", *Port)
	//creates a gRPC server  没有服务 也没有请求被处理
	server := grpc.NewServer()
	proto.RegisterGenerateIdServer(server, &handler.GenerateIDServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	//服务注册
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host,
		global.ServerConfig.ConsulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		//满足GRPC健康检查规范后，这里不需要指定服务了，直接写明地址和端口即可
		GRPC:                           fmt.Sprintf("10.108.130.82:%d", *Port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	serviceID := fmt.Sprintf("%s", uuid.New())
	registration.ID = serviceID
	registration.Port = 8500
	registration.Tags = []string{"distributed-id", "foundation"}
	registration.Address = "10.108.130.82"
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = client.Agent().ServiceDeregister(serviceID); err != nil {
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销成功")
	//client, err := api.NewClient(cfg)
	//if err != nil {
	//	panic(err)
	//}
}
