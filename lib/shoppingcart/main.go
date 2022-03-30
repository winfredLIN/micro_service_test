package Shopping_Cart

import (
	//"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// 关联数据库 建立外键
type shoppingCart struct {
	gorm.Model
	userId          uint
	CommoditiesId   uint
	CommoditiesName string
	CommodityPrice  float32
}

// 购物车加入新的商品 输入：用户id 商品信息，返回：购物车信息
func NewCart(userId uint, commodityId uint, commodityName string, commodityPrice float32) (result shoppingCart, err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/shopping_cart.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(shoppingCart{})
	Cart := shoppingCart{userId: userId, CommoditiesId: commodityId, CommoditiesName: commodityName, CommodityPrice: commodityPrice}
	db.Create(&Cart)
	return Cart, nil
}

// 找到购物车内容，输入：购物车id 返回：购物车信息
func RetrieveCart(cartId uint) (result shoppingCart, err error) {
	var cart shoppingCart
	db, err := gorm.Open("sqlite3", "lib/lib_files/shopping_cart.db")
	if err != nil {
		panic("failed to retrieve")
	} else {
		println("数据库已经连接")
	}
	defer db.Close()

	db.First(&cart, "id = ?", cartId)

	return cart, nil
}

// 删除购物车内容，输入：购物车id 返回：购物车信息
func DeleteCart(cartId uint) (err error) {
	var cart shoppingCart
	db, err := gorm.Open("sqlite3", "lib/lib_files/shopping_cart.db")
	if err != nil {
		panic("failed to delete")
	}
	defer db.Close()

	db.Where("Id = ?", cartId).Delete(&cart)
	return nil
}
