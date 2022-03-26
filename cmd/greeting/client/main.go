package main

import (
	"LaunchService/client"
	//"fmt"
	//"GinService/getparam"
	ginlogin "GinService/LoginService"
)

func main() {
	username, passowrd := ginlogin.GinLogin()
	answer := client.LaunchLoginClient(username, passowrd)
	ginlogin.GinHello(username, answer)

	//client.LaunchGreetingClient(getparam.Getparam())

	// for choose := "xuanze"; choose != "exit"; {
	// 	fmt.Printf("请选择注册(regist)或登录(login),退出请输入exit")
	// 	fmt.Scan(&choose)
	// 	if choose == "login" {
	// 		client.LaunchLoginClient()
	// 	}
	// 	if choose == "regist" {
	// 		client.LaunchRegistrationClient()
	// 	}
	// }
}
