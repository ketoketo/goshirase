package main

import (
	"log"
	"time"

	"github.com/dghubble/go-twitter/twitter"
)

func registerFollower(client *twitter.Client) {
	db := getConnection()
	defer db.Close()
	noticeMigrate(db)

	followers, _, err := client.Followers.IDs(&twitter.FollowerIDParams{})
	if err != nil {
		panic(err.Error)
	}

	for _, followerID := range followers.IDs {
		count := 0
		db.Model(&Notice{}).Where("user_id = ?", followerID).Count(&count)
		// フォロワーを登録
		if count == 0 {
			user, _, err := client.Users.Show(&twitter.UserShowParams{
				UserID: followerID,
			})
			if err != nil {
				return
			}
			log.Printf("Follower count is %d.", user.FollowersCount)
			notice := &Notice{
				UserID:         followerID,
				FollowCount:    user.FollowersCount,
				FollowFlag:     1,
				RegisteredTime: time.Now(),
			}
			db.Create(&notice)
		} else {
			log.Printf("%d is already registered.", followerID)
			notice := Notice{
				UserID:     followerID,
				FollowFlag: 1,
			}
			db.Model(&notice).Updates(Notice{FollowFlag: 1})
		}
	}
}
