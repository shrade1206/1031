package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitMysql() (err error) {
	dsn := "root:root@(127.0.0.1:3306)/todo_list?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf("MySQL Error : %s", err.Error())
		return
	}
	DB.AutoMigrate(&Todo{})
	DB.LogMode(true)
	return
}
