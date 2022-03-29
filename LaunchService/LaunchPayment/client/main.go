package client

import (
	pb "api/protobuf/payment"
	"context"
	"log"
	payment "services/paymentService"
)

func LaunchPayClient(id uint32, totalPrice float32) (AccountBalance float32, success bool) {
	Client, err := payment.NewPaymentClient()
	if err != nil {
		log.Fatalln(err)
	}
	Response, err := Client.Pay(context.Background(), &pb.PayRequest{UserId: id, TotalPrice: totalPrice})

	if err != nil {
		log.Fatalln(err)
	}

	return Response.AccountBalance, Response.Success
}

func LaunchRefundClient(id uint32, totalPrice float32) (AccountBalance float32, success bool) {
	Client, err := payment.NewPaymentClient()
	if err != nil {
		log.Fatalln(err)
	}
	Response, err := Client.Refund(context.Background(), &pb.RefundRequest{})

	if err != nil {
		log.Fatalln(err)
	}

	return Response.AccountBalance, Response.Success

}
