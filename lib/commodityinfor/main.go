package Commodity_Infor

import (
	//"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type commodity struct {
	gorm.Model
	CommodityName        string
	CommodityDiscription string
	Price                float32
}

func NewCommodity(Name string, Discription string, Price float32) (restlt commodity, err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/commodity_infor.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(commodity{})
	product := commodity{CommodityName: Name, CommodityDiscription: Discription, Price: Price}
	db.Create(&product)
	return product, nil
}

func RetrieveCommodity(id uint) (result commodity, err error) {
	var commodity commodity
	db, err := gorm.Open("sqlite3", "lib/lib_files/commodity_infor.db")
	if err != nil {
		panic("failed to retrieve")
	} else {
		println("数据库已经连接")
	}
	defer db.Close()

	db.First(&commodity, "id = ?", id)
	return commodity, nil
}

func UpdateCommodity(id uint, newName string) (err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/commodity_infor.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()

	db.Model(&commodity{}).Where("id = ?", id).Update("CommodityName", newName)
	return nil
}

func DeleteCommodity(id uint) (err error) {
	var commodity commodity
	db, err := gorm.Open("sqlite3", "lib/lib_files/commodity_infor.db")
	if err != nil {
		panic("failed to delete")
	}
	defer db.Close()

	db.Where("Id = ?", id).Delete(&commodity)
	return nil
}
