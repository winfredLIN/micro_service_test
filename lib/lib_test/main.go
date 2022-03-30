package main

import (
	"fmt"
	"lib/useraccount"
)

func main() {
	user, _ := useraccount.RetrieveUserName("haha")
	fmt.Println(user.Password)
	user1, err := useraccount.RetrieveUserName("name string")
	fmt.Println(user1.Name)
	fmt.Println(err)
}
