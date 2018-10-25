package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func deleteNotices(client *twitter.Client) {
	db := getConnection()
	defer db.Close()

	// :TODO フォロワーと上限値比較処理

	sql := strings.Replace(noticesDeleteTmplate, "##COUNT##", strconv.Itoa(30), -1)
	log.Println(sql)
	db.Exec(sql)
}
