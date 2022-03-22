package main

import (
	pb "api/protobuf/greeting"
	"context"
	"fmt"
	"log"
	"services/greeting"
	"services/login"
)

func main() {

	// 客户端：注册
	client1, err := greeting.NewRegistrationClient()
	if err != nil {
		log.Fatalln(err)
	}

	res1, err := client1.Register(context.Background(), &pb.RegisterRequest{RegisterCall: true})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("answer is", res1.RegisterAnswer)
	//客户端：登陆
	loginClient, err := login.NewLoginClient()
	if err != nil {
		log.Fatalln(err)
	}
	loginResponse, err := loginClient.Login(context.Background(), &pb.LoginRequest{LoginCall: true})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("answer is", loginResponse.LoginAnswer)

	//客户端：问候
	client, err := greeting.NewClient()
	if err != nil {
		log.Fatalln(err)
	}
	//发送SayHelloRequest，包括内容Name
	res, err := client.SayHello(context.Background(), &pb.SayHelloRequest{SayHelloName: "vincent"})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.SayHelloAnswer)

}
