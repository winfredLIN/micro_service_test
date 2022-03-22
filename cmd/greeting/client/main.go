package main

import (
	"LaunchService/client"
	"fmt"
)

func main() {
	client.LaunchGreetingClient()
	for choose := "xuanze"; choose != "exit"; {
		fmt.Printf("请选择注册(login)或登录(regist),退出请输入exit")
		fmt.Scan(&choose)
		if choose == "login" {
			client.LaunchLoginClient()
		}
		if choose == "regist" {
			client.LaunchRegistrationClient()
		}
	}
}
