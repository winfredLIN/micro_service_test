package Commodity_Stock

import (
	//"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Stock struct {
	gorm.Model
	CommoditiesId     uint
	CommoditiesName   string
	CommoditiesNumber int
	CommoditiesPrice  float32
}

func New_Stock(id uint, name string, number int,price float32) (err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/Commodity_Stock.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(Stock{})
	stock := Stock{CommoditiesId: id, CommoditiesName: name, CommoditiesNumber: number,CommoditiesPrice: price}
	db.Create(&stock)
	return nil
}

func Retrieve_Stock(Id uint) (result Stock, err error) {
	var stock Stock
	db, err := gorm.Open("sqlite3", "lib/lib_files/Commodity_Stock.db")
	if err != nil {
		panic("failed to retrieve")
	} else {
		println("数据库已经连接")
	}
	defer db.Close()

	db.First(&stock, "id = ?", Id)

	return stock, nil
}

func Delete_Stock(Id uint) (err error) {
	var stock Stock
	db, err := gorm.Open("sqlite3", "lib/lib_files/Commodity_Stock.db")
	if err != nil {
		panic("failed to delete")
	}
	defer db.Close()

	db.Where("Id = ?", Id).Delete(&stock)
	return nil
}

func Update_Number(id uint, number int) (err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/Commodity_Stock.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()

	db.Model(&Stock{}).Where("CommoditiesId  = ?", id).Update("CommoditiesNumber", number)
	return nil
}
func Update_Price(id uint, price float32) (err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/Commodity_Stock.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()

	db.Model(&Stock{}).Where("CommoditiesId  = ?", id).Update("CommoditiesPrice", price)
	return nil
}

