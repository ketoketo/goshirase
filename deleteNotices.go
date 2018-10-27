package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func deleteNotices(client *twitter.Client, upperLimit int) {
	db := getConnection()
	defer db.Close()

	// フォロワーと上限値比較処理
	nowRecord := 0
	db.Model(&Notice{}).Count(&nowRecord)
	delCount := nowRecord - upperLimit
	log.Printf("delCount: %d, upperLimit: %d, nowRecord: %d", delCount, upperLimit, nowRecord)

	// 上限を超えている場合、フォロワー以外で越えた分削除
	if delCount > 0 {
		sql := strings.Replace(noticesDeleteTmplate, "##COUNT##", strconv.Itoa(delCount), -1)
		log.Println(sql)
		db.Exec(sql)
	}
}
