package Commodity_Order

import (
	//"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// 关联数据库 建立外键
type Order struct {
	gorm.Model
	CommoditiesId   []uint
	CommoditiesName []string
}
type Commodity struct {
	gorm.Model
	CommodityName        string
	CommodityType        string
	CommodityDiscription string
	Price                float32
}

func New_Order(commodity ...Commodity) (err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/Commodity_Order.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	commodityId := make([]uint, 0)
	commodityName := make([]string, 0)

	for _, i := range commodity {
		commodityId = append(commodityId, i.ID)
		commodityName = append(commodityName, i.CommodityName)
	}
	db.AutoMigrate(Order{})
	order := Order{CommoditiesId: commodityId, CommoditiesName: commodityName}
	db.Create(&order)
	return nil
}

func Retrieve_Order(orderId uint) (result Order, err error) {
	var order Order
	db, err := gorm.Open("sqlite3", "lib/lib_files/Commodity_Order.db")
	if err != nil {
		panic("failed to retrieve")
	} else {
		println("数据库已经连接")
	}
	defer db.Close()

	db.First(&order, "id = ?", orderId)

	return order, nil
}

func Delete_Order(orderId uint) (err error) {
	var order Order
	db, err := gorm.Open("sqlite3", "lib/lib_files/Commodity_Order.db")
	if err != nil {
		panic("failed to delete")
	}
	defer db.Close()

	db.Where("Id = ?", orderId).Delete(&order)
	return nil
}
