package server

import (
	pb "api/protobuf/payment"
	cfg "config"
	"log"
	"net"
	"os"
	"os/signal"
	"services/paymentService"
	"syscall"

	"google.golang.org/grpc"
)

var config = paymentService.Config{
	Host: cfg.GetConfig().Server.Host,
	Port: ":" + cfg.GetConfig().Server.Port,
}

func LaunchPayServer() {
	// 建立服务并且注册
	PayServer := paymentService.NewPayServer()

	grpcServer := grpc.NewServer()
	// 注册需要注册gRPC的服务以及在protoc定义的服务
	pb.RegisterPaymentServiceServer(grpcServer, PayServer)
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
