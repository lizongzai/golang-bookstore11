package mysql

import (
	"bookstore/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	dsn := "root:password@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Connect database failed.")
	} else {
		fmt.Println("Connected database successfully.")
	}

	DB = db
	fmt.Println("Database Name: ", DB.Name())

	if errMsg := DB.AutoMigrate(model.User{}, model.Book{}); errMsg != nil {
		log.Fatal(errMsg)
	}
}
