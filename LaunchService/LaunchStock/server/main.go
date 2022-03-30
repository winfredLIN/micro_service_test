package server

import (
	pb "api/protobuf/stock"
	cfg "config"
	"log"
	"net"
	"os"
	"os/signal"
	"service/stockservice"
	"syscall"

	"google.golang.org/grpc"
)

var config = stockservice.Config{
	Host: cfg.GetConfig().Server.Host,
	Port: ":" + cfg.GetConfig().Server.Port,
}

func LaunchStockServer() {
	// 建立服务并且注册
	StockServer := stockservice.NewStockServer()

	grpcServer := grpc.NewServer()
	// 注册需要注册gRPC的服务以及在protoc定义的服务
	pb.RegisterStockServiceServer(grpcServer, StockServer)
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
