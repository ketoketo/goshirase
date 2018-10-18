package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Target struct {
	UserID      int64 `sql:"type:bigint(20)" gorm:"primary_key"`
	FollowerNum int64 `gorm:"not null"`
}

type TargetDetail struct {
	UserID   int64 `sql:"type:bigint(20)" gorm:"primary_key"`
	Follower int64 `sql:"type:bigint(20)" gorm:"primary_key"`
}

func getConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:mysql@/go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&Target{})
	db.AutoMigrate(&TargetDetail{})
}
