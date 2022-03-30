package main

import (
	Commodity "launchservice/launchcommodity/server"
	Payment "launchservice/launchpayment/server"
	Stock "launchservice/launchstock/server"
	User "launchservice/launchuser/server"
)

func main() {
	User.LaunchLoginServer()
	User.LaunchRegistrationServer()

	Commodity.LaunchCommodityServer()
	Stock.LaunchStockServer()
	Payment.LaunchPayServer()
}
