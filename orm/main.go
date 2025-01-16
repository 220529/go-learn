package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID       int64     `gorm:"column:id;primary_key"`
	UserID   string    `gorm:"column:user_id"`
	Password string    `gorm:"column:password"`
	Nickname string    `gorm:"column:nickname"`
	Ct       time.Time `gorm:"column:created_at"`
	Ut       time.Time `gorm:"column:updated_at"`
}

func (a Account) TableName() string {
	table := "account"
	return table
}

func connectDB() *gorm.DB {
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
	return mysqlDB
}

func main() {
	db := connectDB()
	var accounts []Account
	if err := db.Find(&accounts).Error; err != nil {
		panic(err)
	}
	var account Account
	if err := db.Where("user_id = ?", 1).Find(&account).Error; err != nil {
		panic(err)
	}
	fmt.Println("accounts: ", accounts)
	fmt.Println("account: ", account)
}
