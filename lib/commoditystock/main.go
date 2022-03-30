package Commodity_Stock

import (
	//"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type stock struct {
	gorm.Model
	CommoditiesId     uint
	CommoditiesName   string
	CommoditiesNumber int
	CommoditiesPrice  float32
}

func NewStock(id uint, name string, number int,price float32) (err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/commodity_stock.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(stock{})
	stock := stock{CommoditiesId: id, CommoditiesName: name, CommoditiesNumber: number,CommoditiesPrice: price}
	db.Create(&stock)
	return nil
}

func RetrieveStock(Id uint) (result stock, err error) {
	var stock stock
	db, err := gorm.Open("sqlite3", "lib/lib_files/commodity_stock.db")
	if err != nil {
		panic("failed to retrieve")
	} else {
		println("数据库已经连接")
	}
	defer db.Close()

	db.First(&stock, "id = ?", Id)

	return stock, nil
}

func DeleteStock(Id uint) (err error) {
	var stock stock
	db, err := gorm.Open("sqlite3", "lib/lib_files/commodity_stock.db")
	if err != nil {
		panic("failed to delete")
	}
	defer db.Close()

	db.Where("Id = ?", Id).Delete(&stock)
	return nil
}

func UpdateNumber(id uint, number int) (err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/commodity_stock.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()

	db.Model(&stock{}).Where("CommoditiesId  = ?", id).Update("CommoditiesNumber", number)
	return nil
}
func UpdatePrice(id uint, price float32) (err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/Commodity_Stock.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()

	db.Model(&stock{}).Where("CommoditiesId  = ?", id).Update("CommoditiesPrice", price)
	return nil
}

