package User_Payment

import (
	//"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type usersPayment struct {
	gorm.Model
	Account_Balance  float32 //账户余额
	Premium          bool    // 会员
	Receving_Address string
}

// 创建
func CreatePayment(recevingAddress string) (result usersPayment, err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/user_payment.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(usersPayment{})
	userPayment := usersPayment{Account_Balance: 0, Premium: false, Receving_Address: recevingAddress}
	db.Create(&userPayment)
	return userPayment, nil
}

//找到支付信息
func RetrievePayment(id uint) (result usersPayment, err error) {
	var userPayment usersPayment
	db, err := gorm.Open("sqlite3", "lib/lib_files/user_payment.db")
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
func ChangeAccountBalance(id uint, money float32) (result usersPayment, err error) {
	var userPayment usersPayment
	db, err := gorm.Open("sqlite3", "lib/lib_files/user_payment.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()
	db.First(&userPayment, "Id = ?", id)
	accountBalance := userPayment.Account_Balance + money
	db.Model(&usersPayment{}).Where("Id = ?", id).Update("Account_Balance", accountBalance)
	return userPayment, nil
}

//会员变化
func ChangePremium(id uint, change bool) (result usersPayment, err error) {
	var userPayment usersPayment
	db, err := gorm.Open("sqlite3", "lib/lib_files/user_payment.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()
	db.Model(&usersPayment{}).Where("Id = ?", id).Update("Premium ", change)
	return userPayment, nil
}

// 地址变化
func ChangeAdress(id uint, adress string) (result usersPayment, err error) {
	var userPayment usersPayment
	db, err := gorm.Open("sqlite3", "lib/lib_files/user_payment.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()
	db.Model(&usersPayment{}).Where("Id = ?", id).Update("Receving_Address", adress)
	return userPayment, nil
}
