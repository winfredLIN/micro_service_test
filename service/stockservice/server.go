package stockservice

import (
	pb "api/protobuf/stock"
	"context"
	stock "lib/commoditystock"
)

type Server struct {
	pb.StockServiceServer
}

func NewStockServer() pb.StockServiceServer {
	return &Server{}
}

//更改库存数量 更改价格
func (s *Server) StockChange(ctx context.Context, req *pb.ChangeRequest) (*pb.ChangeResponse, error) {
	product, _ := stock.RetrieveStock(uint(req.CommodityId))
	if req.ChangeCall == "number" {
		number := product.CommoditiesNumber + int(req.CommodityNumber)
		stock.UpdateNumber(uint(req.CommodityId), number)
	}
	if req.ChangeCall == "price" {
		stock.UpdatePrice(uint(req.CommodityId), req.CommodityPrice)
	}
	return &pb.ChangeResponse{ChangeAnswer: true}, nil
}

func (s *Server) StockShow(ctx context.Context, req *pb.ShowRequest) (*pb.ShowResponse, error) {
	product, _ := stock.RetrieveStock(uint(req.CommodityId))

	return &pb.ShowResponse{CommodityPrice: product.CommoditiesPrice, CommodityId: uint32(product.CommoditiesId), CommodityNumber: uint32(product.CommoditiesNumber), CommodityName: product.CommoditiesName}, nil
}
