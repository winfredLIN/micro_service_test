package greeting

import (
	"context"
	"fmt"
	pb "api/protobuf/greeting"

	"lib/dbcontext"
)

type Server struct {
	pb.GreetingServiceServer
	pb.RegistrationServiceServer
}

func NewServer() pb.GreetingServiceServer {
	return &Server{}
}

func (s *Server) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{
		Message: "hello " + req.Name,
	}, nil
}

func NewRegistrationServer() pb.RegistrationServiceServer{
	return &Server{}
}
// 暂时没有设计不想创建账号的退出入口。（进来就别走了）

func (s *Server) Register(ctx context.Context,req *pb.InformationRequest)(*pb.AnswerResponse, error){
	var username string
	var userpassword1 string
	var userpassword2 string
	fmt.Printf("创建账号 \n 请输入用户名：")
	fmt.Scan(&username)

	for {
		fmt.Printf("\n 请输入密码：")
		fmt.Scan(&userpassword1)
		fmt.Printf("\n 请确认密码：")
		fmt.Scan(&userpassword2)
		if userpassword1 == userpassword2 {
			dbcontext.Create_Account(username, userpassword1)
			fmt.Print("成功创建账号")
			break
		}else{fmt.Print("两次输入的密码不一致,请重新输入")}
	}
	return &pb.AnswerResponse{
		Answer: "哈哈",
	},nil
}
