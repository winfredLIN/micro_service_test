package client

import (
	pb "api/protobuf/commodity"
	"context"
	"log"
	commodity "service/commodityservice"
)

func LaunchCommodityAddClient(Id uint32, Name string, Price float32, Discription string) (id uint32) {

	client, err := commodity.NewCommodityClient()
	if err != nil {
		log.Fatalln(err)
	}
	Response, err := client.CommodityAdd(context.Background(), &pb.AddRequest{CommodityId: Id, CommodityName: Name, CommodityPrice: Price, CommodityDiscription: Discription})
	if err != nil {
		log.Fatalln(err)
	}
	return Response.IdResponse
}

func LaunchCommodityShowClient(id uint32) (Name string, Price float32, Discription string) {
	client, err := commodity.NewCommodityClient()
	if err != nil {
		log.Fatalln(err)
	}
	Response, err := client.CommodityShow(context.Background(), &pb.ShowRequest{CommodityId: id})
	if err != nil {
		log.Fatalln(err)
	}
	return Response.NameResponse, Response.PriceResponse, Response.CommodityDiscription
}
