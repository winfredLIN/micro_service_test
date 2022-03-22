package main

import (
	"fmt"
	"lib/dbcontext"
)

func main() {
	dbcontext.Create_Account("haha", "123456")
	dbcontext.Retrieve_Account(1)
	infor, _ := dbcontext.Retrieve_Account(1)
	fmt.Println("user name : ", infor.Name)
}
