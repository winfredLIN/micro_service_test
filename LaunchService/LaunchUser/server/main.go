package server

import (
	pb "api/protobuf/user"
	cfg "config"
	"log"
	"net"
	"os"
	"os/signal"
	"service/userservice"
	"syscall"

	"google.golang.org/grpc"
)

var config = userservice.Config{
	Host: cfg.GetConfig().Server.Host,
	Port: ":" + cfg.GetConfig().Server.Port,
}

func LaunchLoginServer() {
	// 建立服务并且注册
	LoginServer := userservice.NewLoginServer()

	grpcServer := grpc.NewServer()
	// 注册需要注册gRPC的服务以及在protoc定义的服务
	pb.RegisterLoginServiceServer(grpcServer, LoginServer)
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

func LaunchRegistrationServer() {
	// 建立服务并且注册
	RegistrationServer := userservice.NewRegistrationServer()

	grpcServer := grpc.NewServer()
	// 注册需要注册gRPC的服务以及在protoc定义的服务
	pb.RegisterRegistrationServiceServer(grpcServer, RegistrationServer)
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
