package client

import (
	pb "api/protobuf/stock"
	"context"
	"log"
	stock "service/stockservice"
)

// 暂时只允许更改数量和价格
func StockChangeClient(call string, price float32, number uint32, id uint32, name string) (ChangeAnswer bool) {
	Client, err := stock.NewStockServiceClient()
	if err != nil {
		log.Fatalln(err)
	}
	Response, err := Client.StockChange(context.Background(), &pb.ChangeRequest{ChangeCall: call, CommodityPrice: price, CommodityNumber: number, CommodityId: id})
	if err != nil {
		log.Fatalln(err)
	}

	return Response.ChangeAnswer
}

// 注册
func StockShowClient(id uint32) (price float32, Id uint32, number uint32, name string) {
	// 客户端：注册
	client, err := stock.NewStockServiceClient()
	if err != nil {
		log.Fatalln(err)
	}
	res, err := client.StockShow(context.Background(), &pb.ShowRequest{CommodityId: id})

	if err != nil {
		log.Fatalln(err)
	}
	return res.CommodityPrice, res.CommodityId, res.CommodityNumber, res.CommodityName
}
