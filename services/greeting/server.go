package greeting

import (
	pb "api/protobuf/greeting"
	"context"
	"fmt"

	"lib/User_Account"
)

type Server struct {
	pb.GreetingServiceServer
	pb.RegistrationServiceServer
	pb.LoginServiceServer
}

// 实现问候服务
func NewServer() pb.GreetingServiceServer {
	return &Server{}
}

func (s *Server) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{
		SayHelloAnswer: "hello " + req.SayHelloName,
	}, nil
}

// 实现注册服务
func NewRegistrationServer() pb.RegistrationServiceServer {
	return &Server{}
}
func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var username string
	var userpassword1 string
	var userpassword2 string

	for {
		fmt.Printf("创建账号 \n 请输入用户名：")
		fmt.Scan(&username)
		name := User_Account.Retrieve_UserName(username).Name
		fmt.Printf("创建的名字 %s ，查询的名字 %s", username, name)
		if name == username {
			fmt.Println("用户名重复，请重新输入")
		} else {
			fmt.Println("用户名可用")
			break
		}

	}

	for {
		fmt.Printf("\n 请输入密码：")
		fmt.Scan(&userpassword1)
		fmt.Printf("\n 请确认密码：")
		fmt.Scan(&userpassword2)
		if userpassword1 == userpassword2 {
			User_Account.Create_Account(username, userpassword1)
			fmt.Print("成功创建账号")
			break
		} else {
			fmt.Print("两次输入的密码不一致,请重新输入")
		}
	}
	return &pb.RegisterResponse{
		RegisterAnswer: "这是账号注册服务",
	}, nil
}