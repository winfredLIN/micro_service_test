package main

import (
	Commodity "LaunchService/LaunchCommodity/server"
	Payment "LaunchService/LaunchPayment/server"
	Stock "LaunchService/LaunchStock/server"
	User "LaunchService/LaunchUser/server"
)

func main() {
	User.LaunchLoginServer()
	User.LaunchRegistrationServer()

	Commodity.LaunchCommodityServer()
	Stock.LaunchStockServer()
	Payment.LaunchPayServer()
}
