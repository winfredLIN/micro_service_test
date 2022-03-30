package useraccount

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type usersInfo struct {
	//gorm.Model // uid、创建、更新和删除表的时间
	//Id uint16 `gorm:"primaryKey"`
	gorm.Model

	Name     string
	Password string //
}

// 新建用户 输入：昵称 密码 输出：用户，err
func CreateAccount(name string, password string) (result usersInfo, err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/user_account.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(usersInfo{})
	user := usersInfo{Name: name, Password: password}
	db.Create(&user)
	return user, nil
}

// 查找用户信息 输入：用户id 输出：用户信息
func RetrieveAccount(id uint16) (result usersInfo, err error) {
	var User usersInfo
	db, err := gorm.Open("sqlite3", "lib/lib_files/user_account.db")
	if err != nil {
		panic("failed to retrieve")
	} else {
		println("数据库已经连接")
	}
	defer db.Close()

	//db.Where("id = ?", id).First(&user)
	db.First(&User, "Id = ?", id) //查询某个对象
	println("haha:", User.Name)
	return User, nil
}

// 通过昵称查找用户信息 输入：用户昵称 输出：用户信息
func RetrieveUserName(name string) (result usersInfo, err error) {
	var user usersInfo
	db, err := gorm.Open("sqlite3", "lib/lib_files/user_account.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()
	db.Where("name = ?", name).First(&user)
	return user, nil
}

// 更改用户名称 输入：用户id ，新的名称 输出：用户信息，err
func UpdateAccount(id uint16, name string) (result usersInfo, err error) {
	var user usersInfo
	db, err := gorm.Open("sqlite3", "lib/lib_files/user_account.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()

	db.Model(&usersInfo{}).Where("Id = ?", id).Update("Name", name)
	return user, nil
}

// 删除账户 输入：用户id 输出：err
func DeleteAccount(id uint16) (err error) {
	var user usersInfo
	db, err := gorm.Open("sqlite3", "lib/lib_files/user_account.db")
	if err != nil {
		panic("failed to delete")
	}
	defer db.Close()

	db.Where("Id = ?", id).Delete(&user)
	return nil
}
