package dbcontext

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Users_Info struct {
	//gorm.Model // uid、创建、更新和删除表的时间
	Id uint16 `gorm:"primaryKey"`
	//CreatedAt time.Time `gorm:"<-:create"`// 只允许创建和读取
	//UpdatedAt time.Time `gorm:"<-:create";gorm:"autoUpdateTime:nano"` // 使用时间戳填纳秒数充更新时间`
	//DeletedAt time.Time `gorm:"<-:create"` //gorm.DeletedAt `gorm:"index"`设置索引
	Name            string
	Account_Balance float32 //账户余额，先充钱再买
	Premium         bool    // 会员打9折
	Password        string  //
}

func Create_Account(name string, password string) (err error) {
	db, err := gorm.Open("sqlite3", "lib/lib_files/User_Account.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(Users_Info{})
	user := Users_Info{Name: name, Password: password, Account_Balance: 0, Premium: false}
	db.Create(&user)
	return nil
}

func Retrieve_Account(id uint16) (result Users_Info, err error) {
	var user Users_Info
	db, err := gorm.Open("sqlite3", "lib/lib_files/User_Account.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()

	db.First(&user, "Id = ?", id) //查询某个对象
	return user, nil
}

// 查询表中是否有同样的昵称，如果有，则返回false
func Retrieve_UserName(name string) (result Users_Info) {
	var user Users_Info
	db, err := gorm.Open("sqlite3", "lib/lib_files/User_Account.db")
	if err != nil {
		panic("failed to retrieve")
	}
	defer db.Close()
	db.Where("name = ?", name).First(&user)
	return user
}

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
