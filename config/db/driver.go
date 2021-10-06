package db // Package db file driver ini untuk setup ke database

import (
	"boilerplate/config/env"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	var err error
	var configDb = env.Config

	if DB != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println("recovery error")
		}
	}()

	urlDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",configDb.Username,configDb.Password,configDb.Host,configDb.Port,configDb.Name)

	DB, err = gorm.Open(mysql.Open(urlDB), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to DB")
	}

	log.Print("success connect to db")

	AutoMigrate(DB)

}