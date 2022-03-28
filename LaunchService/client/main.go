package client

import (
	pb "api/protobuf/user"
	"context"
	"fmt"
	"log"
	login "services/userService"
)

// 开放用户名密码的接口给gin
func LaunchLoginClient(username string, password string) (answer string) {
	loginClient, err := login.NewLoginClient()
	if err != nil {
		log.Fatalln(err)
	}
	loginResponse, err := loginClient.Login(context.Background(), &pb.LoginRequest{LoginCall: true, Username: username, Password: password})
	if err != nil {
		log.Fatalln(err)
	}

	return loginResponse.Password
}

func LaunchRegistrationClient() bool {
	// 客户端：注册
	client, err := login.NewRegistrationClient()
	if err != nil {
		log.Fatalln(err)
	}

	res1, err := client.Register(context.Background(), &pb.RegisterRequest{RegisterCall: true})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("answer is", res1.AccountNumber)
	return true
}
