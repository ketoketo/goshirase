package main

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func getConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:mysql@/go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("can not open database!")
		panic(err.Error())
	}
	return db
}
