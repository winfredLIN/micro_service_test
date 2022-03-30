package userservice

import (
	pb "api/protobuf/user"
	"context"
	account "lib/useraccount"
)

type Server struct {
	pb.LoginServiceServer
	pb.RegistrationServiceServer
}

func NewLoginServer() pb.LoginServiceServer {
	return &Server{}
}

// 输入：请求bool 用户名string 密码string，若用户名不存在返回false，err string“用户名不存在” 若正确则返回对应的账号密码
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, _ := account.RetrieveUserName(req.Username)

	if user.Name != req.Username {
		return &pb.LoginResponse{NameCorrect: false}, nil
	} else if user.Password != req.Password {
		return &pb.LoginResponse{NameCorrect: true,PasswordCorrect: false}, nil
	}
	return &pb.LoginResponse{NameCorrect: true, PasswordCorrect: true}, nil
}

func NewRegistrationServer() pb.RegistrationServiceServer {
	return &Server{}
}

// 注册账户；若用户名存在返回true，应重新输入。成功注册则返回账号
func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	_, err := account.RetrieveUserName(req.Username)
	if err != nil {
		return &pb.RegisterResponse{NameExist: true}, err
	}
	user, err := account.CreateAccount(req.Username, req.Password)
	return &pb.RegisterResponse{NameExist: false, AccountNumber: int32(user.ID)}, nil
}
