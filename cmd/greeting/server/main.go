package main

import (
	"log"
	pb "micro-service-test/api/protobuf/greeting"
	"micro-service-test/services/greeting"
	"net"
	"os"
	"os/signal"
	"syscall"

	cfg "micro-service-test/config"

	"google.golang.org/grpc"
)

//应该创建一个配置文件服务于所有的文件的配置
type Config struct {
	Host string
	Port string
}

var config = greeting.Config{
	Host: cfg.GetConfig().Server.Host,
	Port: ":" + cfg.GetConfig().Server.Port,
}

func main() {

	// 建立服务并且注册
	greetingServiceServer := greeting.NewServer()
	greetingRegistrationServer := greeting.NewRegistrationServer()

	grpcServer := grpc.NewServer()
	// 注册需要注册gRPC的服务以及在protoc定义的服务
	pb.RegisterGreetingServiceServer(grpcServer, greetingServiceServer)
	pb.RegisterRegistrationServiceServer(grpcServer, greetingRegistrationServer)
	// 监听信道是否拥挤？
	listener, err := net.Listen("tcp", config.Port)

	if err != nil {
		log.Fatalln(err)
	}
	// 需要实现服务端长久运行服务需要运行goroutine gorountine有需要channel
	ch := make(chan os.Signal, 1)

	go func() {
		if err = grpcServer.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
	// 这里是一个优雅退出
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	<-ch

}
