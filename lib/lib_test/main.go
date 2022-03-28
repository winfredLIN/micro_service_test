package main

import (
	"fmt"
	"lib/User_Account"
)

func main() {
	user,_ :=User_Account.Retrieve_UserName("haha")
	fmt.Println(user.Password)
	user1,err :=User_Account.Retrieve_UserName("name string")
	fmt.Println(user1.Name)
	fmt.Println(err)
}
