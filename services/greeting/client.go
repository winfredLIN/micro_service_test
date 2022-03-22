package greeting

import (
	"context"
	"fmt"
	pb "api/protobuf/greeting"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Host string
	Port string
}

// 实现新建Client的方法，这可以服务于很多不同的Client，可复用
func NewClient(config Config) (pb.GreetingServiceClient, error) {
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

func NewRegistrationClient(config Config) (pb.RegistrationServiceClient, error) {
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
func NewLoginClient(config Config)(pb.LoginServiceClient,error){
	//Client需要在没有回应的时候自己结束，因为很多服务没有交互界面让用户自行关闭请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println(config.Host + config.Port)
	connection, err := grpc.DialContext(ctx, config.Host+config.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}
	return pb.NewLoginServiceClient(connection), nil
}

