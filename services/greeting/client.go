package greeting

import (
	"context"
	"fmt"
	pb "api/protobuf/greeting"
	"time"
	cfg "config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
// 配置客户端
type Config struct {
	Host string
	Port string
}
var config = Config{
	Host: cfg.GetConfig().Client.Host,
	Port: ":" + cfg.GetConfig().Client.Port,
}

// 实现新建问候的客户端服务
func NewClient() (pb.GreetingServiceClient, error) {
	//Client需要在没有回应的时候自己结束，因为很多服务没有交互界面让用户自行关闭请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println(config.Host + config.Port)
	connection, err := grpc.DialContext(ctx, config.Host+config.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}
	return pb.NewGreetingServiceClient(connection), nil
}
//新建注册的客户端服务
func NewRegistrationClient() (pb.RegistrationServiceClient, error) {
	//Client需要在没有回应的时候自己结束，因为很多服务没有交互界面让用户自行关闭请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println(config.Host + config.Port)
	connection, err := grpc.DialContext(ctx, config.Host+config.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}
	return pb.NewRegistrationServiceClient(connection), nil
}
