package commodityservice

import (
	pb "api/protobuf/commodity"
	"context"
	commodity "lib/commodityinfor"
)

type Server struct {
	pb.CommodityServiceServer
}

func NewCommodityAddServer() pb.CommodityServiceServer {
	return &Server{}
}

// 向商品库添加一个商品，返回商品ID
func (s *Server) CommodityAdd(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	product, _ := commodity.NewCommodity(req.CommodityName, req.CommodityDiscription, req.CommodityPrice)
	return &pb.AddResponse{IdResponse: uint32(product.ID)}, nil
}

// 查询查询商品的信息。输入id 输出信息
func (s *Server) CommodityShow(ctx context.Context, req *pb.ShowRequest) (*pb.ShowResponse, error) {
	product, _ := commodity.RetrieveCommodity(uint(req.CommodityId))
	return &pb.ShowResponse{NameResponse: product.CommodityName, PriceResponse: product.Price, CommodityDiscription: product.CommodityDiscription}, nil
}
