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

var noticesDeleteTmplate = `
DELETE FROM notices
WHERE 
	user_id IN (
		SELECT user_id FROM (
			SELECT user_id FROM notices WHERE follow_flag = 0 LIMIT ##COUNT##
		) AS tmp
	)
`