package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID        int64     `gorm:"column:id;primary_key"`
	UserId    string    `gorm:"column:user_id"`
	Password  string    `gorm:"column:password"`
	Nickname  string    `gorm:"column:nickname"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// 显示声明
//func (Account) TableName() string {
//	return "account"
//}

func main() {
	db := connDB()
	var account []Account
	//if err := db.Find(&account).Error; err != nil {
	//	fmt.Println(err)
	//	return
	//}
	if err := db.Where("id = ?", 1).Find(&account).Error; err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(account)
}
func connDB() *gorm.DB {
	mysqlDB, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/contentlibrary?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)
	mysqlDB = mysqlDB.Debug()

	return mysqlDB
}
