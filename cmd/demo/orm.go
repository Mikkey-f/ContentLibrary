package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

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
