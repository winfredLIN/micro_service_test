package login

import (
	pb "api/protobuf/greeting"
	"context"
	"fmt"

	"lib/dbcontext"
)

type Server struct {
	pb.LoginServiceServer
}

// 实现服务端登陆的服务
func NewLoginServer()pb.LoginServiceServer{
	return &Server{}
}
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest)(*pb.LoginResponse,error){
	// 1检查账号是否存在2检查账号密码是否正确
	var username string
	var userpassword string
	for {
		fmt.Printf("登陆 \n 请输入用户名：")
		fmt.Scan(&username)
		name := dbcontext.Retrieve_UserName(username).Name
		fmt.Printf("创建的名字 %s ，查询的名字 %s", username, name)
		if name == username {
			break
		} else {
			fmt.Println("用户名不存在")
		}

	}

	for {
		fmt.Printf("\n 请输入密码：")
		fmt.Scan(&userpassword)
		if userpassword == dbcontext.Retrieve_UserName(username).Password {
			fmt.Print("登陆成功")
			break
		} else {
			fmt.Print("密码错误,请重新输入")
		}
	}
	return &pb.LoginResponse{LoginAnswer: "这是登陆服务"},nil
}