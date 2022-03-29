package main

import (
	Commodity "LaunchService/LaunchCommodity/server"
	User "LaunchService/LaunchUser/server"
)

func main() {

	User.LaunchLoginServer()
	User.LaunchRegistrationServer()

	Commodity.LaunchCommodityServer()
}
