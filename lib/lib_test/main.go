package main

import (
	"fmt"
	"lib/User_Account"
)

func main() {
	User_Account.Create_Account("haha", "123456")
	User_Account.Retrieve_Account(1)
	infor, _ := User_Account.Retrieve_Account(1)
	fmt.Println("user name : ", infor.Name)
}
