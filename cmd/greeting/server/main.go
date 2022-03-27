package main

import (
	"log"
	pb "api/protobuf/user"
	"services/greeting"
	"services/login"
	"net"
	"os"
	"os/signal"
	"syscall"

	cfg "config"

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

	// "建立"服务并且注册
	greetingRegistrationServer := greeting.NewRegistrationServer()
	LoginServiceServer := login.NewLoginServer()

	grpcServer := grpc.NewServer()
	// "注册"需要注册gRPC的服务以及在protoc定义的服务
	pb.RegisterRegistrationServiceServer(grpcServer, greetingRegistrationServer)
	pb.RegisterLoginServiceServer(grpcServer, LoginServiceServer)
	
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
