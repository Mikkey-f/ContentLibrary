package services

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CmsApp struct {
	db *gorm.DB
}

func NewCmsApp() *CmsApp {
	app := &CmsApp{}
	connDB(app)
	return app
}

func connDB(app *CmsApp) {
	mysqlDB, err := gorm.Open(mysql.Open("root:123456@tcp(172.30.210.158:3306)/?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)
	app.db = mysqlDB
}
