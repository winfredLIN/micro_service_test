package client

import (
	pb "api/protobuf/user"
	"context"
	"log"
	login "services/userService"
)

// 登陆
func LaunchLoginClient(username string, password string) (nameCrrect bool, answer string) {
	Client, err := login.NewLoginClient()
	if err != nil {
		log.Fatalln(err)
	}
	loginResponse, err := Client.Login(context.Background(), &pb.LoginRequest{LoginCall: true, Username: username, Password: password})
	if err != nil {
		log.Fatalln(err)
	}

	return loginResponse.NameCorrect, loginResponse.Password
}

// 注册
func LaunchRegistrationClient(username string, password string) (NameExist bool, AccountNumber int32) {
	// 客户端：注册
	client, err := login.NewRegistrationClient()
	if err != nil {
		log.Fatalln(err)
	}
	res, err := client.Register(context.Background(), &pb.RegisterRequest{RegisterCall: true, Username: username, Password: password})

	if err != nil {
		log.Fatalln(err)
	}
	return res.NameExist, res.AccountNumber
}
