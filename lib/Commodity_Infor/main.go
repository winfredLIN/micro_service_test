package Commodity_Infor

import (
	//"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Commodity struct {
	gorm.Model
	CommodityName        string
	CommodityDiscription string
	Price                float32
}

func New_Commodity(Name string, Discription string, Price float32) (restlt Commodity, err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/Commodity_Infor.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(Commodity{})
	product := Commodity{CommodityName: Name, CommodityDiscription: Discription, Price: Price}
	db.Create(&product)
	return product, nil
}

func Retrieve_Commodity(id uint) (result Commodity, err error) {
	var commodity Commodity
	db, err := gorm.Open("sqlite3", "lib/lib_files/Commodity_Infor.db")
	if err != nil {
		panic("failed to retrieve")
	} else {
		println("数据库已经连接")
	}
	defer db.Close()

	db.First(&commodity, "id = ?", id)
	return commodity, nil
}

func Update_Commodity(id uint, newName string) (err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/Commodity_Infor.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()

	db.Model(&Commodity{}).Where("id = ?", id).Update("CommodityName", newName)
	return nil
}

func Delete_Commodity(id uint) (err error) {
	var commodity Commodity
	db, err := gorm.Open("sqlite3", "lib/lib_files/Commodity_Infor.db")
	if err != nil {
		panic("failed to delete")
	}
	defer db.Close()

	db.Where("Id = ?", id).Delete(&commodity)
	return nil
}
