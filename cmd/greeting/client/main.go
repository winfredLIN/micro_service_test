package main

import (
	"context"
	"fmt"
	"log"
	pb "micro-service-test/api/protobuf/greeting"
	"micro-service-test/services/greeting"
)

//应该创建一个配置文件服务于所有的文件的配置
var config = greeting.Config{
	Host: "192.168.1.104",
	Port: ":50051",
}

func main() {
	//创建并配置Client
	client, err := greeting.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}
	//发送SayHelloRequest，包括内容Name
	res, err := client.SayHello(context.Background(), &pb.SayHelloRequest{Name: "vincent"})
	// 输入iferr回车即可
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.Message)

	client1,err := greeting.NewRegistrationClient(config)
	if err != nil {
		log.Fatalln(err)
	}
	res1, err := client1.Register(context.Background(), &pb.InformationRequest{Request: true})
	
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("answer is",res1.Answer)

}
//ctrl+shift reload 可以重启vscode