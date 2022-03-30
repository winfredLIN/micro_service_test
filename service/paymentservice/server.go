package paymentservice

import (
	pb "api/protobuf/payment"
	"context"
	user "lib/userpayment"
)

type Server struct {
	pb.PaymentServiceServer
}

func NewPayServer() pb.PaymentServiceServer {
	return &Server{}
}

// 只支付，在BFF层再对身份进行验证
func (s *Server) Pay(ctx context.Context, req *pb.PayRequest) (*pb.PayResponse, error) {
	userinfo, _ := user.RetrievePayment(uint(req.UserId))

	if userinfo.Account_Balance < req.TotalPrice {
		return &pb.PayResponse{AccountBalance: userinfo.Account_Balance, Success: false}, nil
	}

	balance := userinfo.Account_Balance - req.TotalPrice
	user.ChangeAccountBalance(uint(req.UserId), balance)
	return &pb.PayResponse{AccountBalance: balance, Success: true}, nil
}

func (s *Server) Refund(ctx context.Context, req *pb.RefundRequest) (*pb.RefundResponse, error) {
	userinfo, _ := user.RetrievePayment(uint(req.UserId))

	balance := userinfo.Account_Balance + req.TotalPrice
	user.ChangeAccountBalance(uint(req.UserId), balance)
	return &pb.RefundResponse{AccountBalance: balance, Success: true}, nil
}
