package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func goshirase(client *twitter.Client) {
	db := getConnection()
	defer db.Close()
	noticeMigrate(db)

	notices := []Notice{}

	// 1日経過している場合、対象
	db.Where("registered_time < (NOW() - INTERVAL 1 DAY)").Limit(900).Find(&notices)
	// db.Where("registered_time < NOW()").Limit(900).Find(&notices) // debug

	for _, notice := range notices {
		user, _, err := client.Users.Show(&twitter.UserShowParams{
			UserID: notice.UserID,
		})
		if err != nil {
			panic(err.Error)
		}

		log.Printf("before: %d, after: %d", notice.FollowCount, user.FollowersCount)
		// 前回よりフォロワーが減っている場合お知らせしてあげる
		diff := notice.FollowCount - user.FollowersCount
		if diff > 0 {
			message := strings.Replace(messageTmplate, "##UPTIME##", notice.RegisteredTime.Format("15時04分05秒"), -1)
			message = strings.Replace(message, "##FOL##", strconv.Itoa(diff), -1)
			log.Println("@" + user.ScreenName + " " + message)
			// client.Statuses.Update("@"+user.ScreenName+" "+message, nil)
		}
		// :TODO DB更新処理　フォロワーとdate
	}

}

var messageTmplate = `前日の##UPTIME##に比べ、フォロワーが##FOL##人減少しています。

フォローしていただければ、このお知らせをずっと受け取れます。
#goshirase
`
