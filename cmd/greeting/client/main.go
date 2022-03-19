package main

import (
	pb "api/protobuf/greeting"
	"context"
	"fmt"
	"log"
	cfg "config"
	"services/greeting"
)

type Config struct {
	Host string
	Port string
}

var config = greeting.Config{
	Host: cfg.GetConfig().Client.Host,
	Port: ":" + cfg.GetConfig().Client.Port,
}

func main() {
	//创建并配置Client
	client1, err := greeting.NewRegistrationClient(config)
	if err != nil {
		log.Fatalln(err)
	}

	res1, err := client1.Register(context.Background(), &pb.InformationRequest{Request: true})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("answer is", res1.Answer)

	client, err := greeting.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}
	//发送SayHelloRequest，包括内容Name
	res, err := client.SayHello(context.Background(), &pb.SayHelloRequest{Name: "vincent"})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.Message)

}

//ctrl+shift reload 可以重启vscode
