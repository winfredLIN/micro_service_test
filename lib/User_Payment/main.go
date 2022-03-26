package User_Payment

import (
	//"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Users_Payment struct {
	gorm.Model
	Account_Balance  float32 //账户余额
	Premium          bool    // 会员
	Receving_Address string
}

// 创建
func Create_Payment(recevingAddress string) (result Users_Payment, err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/User_Payment.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(Users_Payment{})
	userPayment := Users_Payment{Account_Balance: 0, Premium: false, Receving_Address: recevingAddress}
	db.Create(&userPayment)
	return userPayment, nil
}

//找到支付信息
func Retrieve_Payment(id uint) (result Users_Payment, err error) {
	var userPayment Users_Payment
	db, err := gorm.Open("sqlite3", "lib/lib_files/User_Payment.db")
	if err != nil {
		panic("failed to retrieve")
	} else {
		println("数据库已经连接")
	}
	defer db.Close()

	//db.Where("id = ?", id).First(&user)
	db.First(&userPayment, "Id = ?", id) //查询某个对象
	return userPayment, nil
}

//账户金额变化
func Change_Account_Balance(id uint, money float32) (result Users_Payment, err error) {
	var userPayment Users_Payment
	db, err := gorm.Open("sqlite3", "lib/lib_files/User_Payment.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()
	db.First(&userPayment, "Id = ?", id)
	accountBalance := userPayment.Account_Balance + money
	db.Model(&Users_Payment{}).Where("Id = ?", id).Update("Account_Balance", accountBalance)
	return userPayment, nil
}

//会员变化
func Change_Premium(id uint, change bool) (result Users_Payment, err error) {
	var userPayment Users_Payment
	db, err := gorm.Open("sqlite3", "lib/lib_files/User_Payment.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()
	db.Model(&Users_Payment{}).Where("Id = ?", id).Update("Premium ", change)
	return userPayment, nil
}

// 地址变化
func Change_Adress(id uint, adress string) (result Users_Payment, err error) {
	var userPayment Users_Payment
	db, err := gorm.Open("sqlite3", "lib/lib_files/User_Payment.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()
	db.Model(&Users_Payment{}).Where("Id = ?", id).Update("Receving_Address", adress)
	return userPayment, nil
}
