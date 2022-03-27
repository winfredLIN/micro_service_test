package main

import (
	//"fmt"
	"lib/Commodity_Infor"
)

func main() {
	//Commodity_Infor.New_Commodity("法国Rouje online衫上衣女", "上衣", "品牌: ROUJE ONLINE", 521.99)
	//Commodity_Infor.Delete_Commodity(2)
	comodity, _ := Commodity_Infor.Retrieve_Commodity(3)

	print(comodity.CommodityDiscription)

}
