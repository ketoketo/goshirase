package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dghubble/go-twitter/twitter"
)

func notice(client *twitter.Client) {
	db := getConnection()
	defer db.Close()
	noticeMigrate(db)

	params := &twitter.StreamFilterParams{
		Track:         []string{"Docker"},
		StallWarnings: twitter.Bool(true),
	}
	stream, _ := client.Streams.Filter(params)
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		targetUser := tweet.User.ID
		user, _, err := client.Users.Show(&twitter.UserShowParams{
			UserID: targetUser,
		})
		if err != nil {
			panic(err.Error())
		}
		log.Printf("Follower count is %d.", user.FollowersCount)
		log.Println(user.Name)
		log.Println(tweet.Text)
		notice := &Notice{
			UserID:         targetUser,
			FollowCount:    user.FollowersCount,
			FollowFlag:     0,
			RegisteredTime: time.Now(),
		}
		if user.FollowersCount >= 100 {
			db.Where(&notice).FirstOrCreate(&notice)
		}
	}

	go demux.HandleChan(stream.Messages)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	stream.Stop()
}
