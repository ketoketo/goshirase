package main

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Target struct {
	UserID     int64 `sql:"type:bigint(20)" gorm:"primary_key"`
	UpdateTime time.Time
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

func replaceSelectSql(sql string, targetVal string, replaceVal string) string {
	return strings.Replace(sql, targetVal, replaceVal, -1)
}

var REPLACE1 = "###REPLACE###"

var CompareNewOldSQL = `
SELECT 0 as new_old_flag,a.follower as follower FROM target_details a
LEFT JOIN ###REPLACE### b ON a.user_id=b.user_id and a.follower=b.follower
where b.user_id is null
UNION
SELECT 1 as new_old_flag,b.follower as follower FROM target_details a
RIGHT JOIN ###REPLACE### b ON a.user_id=b.user_id and a.follower=b.follower
where a.user_id is null
`

type CompareResult struct {
	NewOldFlag int
	Follower   int64
}
