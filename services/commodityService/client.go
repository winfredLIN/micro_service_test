package commodityService

import (
	pb "api/protobuf/commodity"
	cfg "config"
	"context"
	"time"

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

func NewCommodityClient() (pb.CommodityServiceClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	connection, err := grpc.DialContext(ctx, config.Host+config.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}
	return pb.NewCommodityServiceClient(connection), nil

}
