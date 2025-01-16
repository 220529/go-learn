package services

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CmsApp struct {
	db *gorm.DB
}

func NewCmsApp() *CmsApp {
	app := &CmsApp{}
	connectDB(app)
	return app
}

func connectDB(app *CmsApp) *CmsApp {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connect database success", mysqlDB)
	db, err := mysqlDB.DB()
	if err != nil {
		panic("failed to connect database, 设置失败")
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)
	mysqlDB = mysqlDB.Debug()
	app.db = mysqlDB
	return app
}
