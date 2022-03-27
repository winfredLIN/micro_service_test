package Commodity_Stock

import (
	//"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// 关联数据库 建立外键
type Stock struct {
	gorm.Model
	CommoditiesId     uint
	CommoditiesName   string
	CommoditiesNumber int
}

func New_Stock(id uint, name string, number int) (err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/Commodity_Stock.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(Stock{CommoditiesId: id, CommoditiesName: name, CommoditiesNumber: number})
	stock := Stock{}
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
