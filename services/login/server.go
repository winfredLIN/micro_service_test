package login

import (
	pb "api/protobuf/greeting"
	"context"

	"lib/User_Account"
)

type Server struct {
	pb.LoginServiceServer
}

// 实现服务端登陆的服务
func NewLoginServer() pb.LoginServiceServer {
	return &Server{}
}

// 接收gin传来的用户名密码，访问数据库判断用户名密码是否正确
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	if req.Username != User_Account.Retrieve_UserName(req.Username).Name {
		return &pb.LoginResponse{LoginAnswer: "用户名不存在"}, nil
	}
	if req.Password == User_Account.Retrieve_UserName(req.Username).Password {
		return &pb.LoginResponse{LoginAnswer: "登陆成功"}, nil
	} else {
		return &pb.LoginResponse{LoginAnswer: "密码错误请重新登陆"}, nil
	}
}
