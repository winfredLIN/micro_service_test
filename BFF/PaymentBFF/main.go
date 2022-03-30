package paymentbff

import (
	payment "launchservice/launchpayment/client"
	stock "launchservice/launchstock/client"
	user "launchservice/launchuser/client"
	"fmt"
)

// 选择购物车中的商品（commodityID），支付（userID，UserAccountBalance）库存充足（CommodityStock）
func Payment(commodityId uint32,userName string,Password string)(result bool,reason string){
	nameCorrect, PasswordCorrect := user.LaunchLoginClient(userName, Password)
	if nameCorrect != true || PasswordCorrect != true{
		return false,"密码错误"
	}
	price,_,number,_:=stock.StockShowClient(commodityId)
	if number <= 0{
		return false,"商品已售空"
	}
	accountBalance,success := payment.LaunchPayClient(1, price)
	if success==true{
		stock.StockChangeClient("number", 1, number-1, commodityId, "1")
		return true,"购买成功"
	}else if success==false{
		return false,fmt.Sprintf("余额不足，你的余额为：%f",accountBalance)
	}
	return
}