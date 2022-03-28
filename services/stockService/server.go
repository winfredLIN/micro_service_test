package stockService

import (
	pb "api/protobuf/stock"
	"context"
	stock "lib/Commodity_Stock"
)

type Server struct {
	pb.StockServiceServer
}

func NewStockServer() pb.StockServiceServer {
	return &Server{}
}

func (s *Server) StockChange(ctx context.Context, req *pb.ChangeRequest) (*pb.ChangeResponse, error) {
	product, _ := stock.Retrieve_Stock(uint(req.CommodityId))
	number := product.CommoditiesNumber + int(req.CommodityNumber)
	stock.Update_Number(uint(req.CommodityId), number)
	return &pb.ChangeResponse{ChangeAnswer: true}, nil
}
