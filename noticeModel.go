package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Notice struct {
	UserID         int64 `sql:"type:bigint(20)" gorm:"primary_key"`
	FollowCount    int
	FollowFlag     int
	RegisteredTime time.Time
}

func noticeMigrate(db *gorm.DB) {
	db.AutoMigrate(&Notice{})
}
