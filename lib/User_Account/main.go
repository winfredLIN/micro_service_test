package User_Account

import (
	//"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Users_Info struct {
	//gorm.Model // uid、创建、更新和删除表的时间
	//Id uint16 `gorm:"primaryKey"`
	gorm.Model

	Name     string
	Password string //
}

// 新建用户 输入：昵称 密码 输出：用户，err
func Create_Account(name string, password string) (result Users_Info, err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/User_Account.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(Users_Info{})
	user := Users_Info{Name: name, Password: password}
	db.Create(&user)
	return user, nil
}

// 查找用户信息 输入：用户id 输出：用户信息
func Retrieve_Account(id uint16) (result Users_Info, err error) {
	var User Users_Info
	db, err := gorm.Open("sqlite3", "lib/lib_files/User_Account.db")
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
func Retrieve_UserName(name string) (result Users_Info, err error) {
	var user Users_Info
	db, err := gorm.Open("sqlite3", "lib/lib_files/User_Account.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()
	db.Where("name = ?", name).First(&user)
	return user, err
}

// 更改用户名称 输入：用户id ，新的名称 输出：用户信息，err
func Update_Account(id uint16, name string) (result Users_Info, err error) {
	var user Users_Info
	db, err := gorm.Open("sqlite3", "lib/lib_files/User_Account.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()

	db.Model(&Users_Info{}).Where("Id = ?", id).Update("Name", name)
	return user, nil
}

// 删除账户 输入：用户id 输出：err
func Delete_Account(id uint16) (err error) {
	var user Users_Info
	db, err := gorm.Open("sqlite3", "lib/lib_files/User_Account.db")
	if err != nil {
		panic("failed to delete")
	}
	defer db.Close()

	db.Where("Id = ?", id).Delete(&user)
	return nil
}
