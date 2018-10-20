package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "", "Twitter Access Token")
	accessSecret := flags.String("access-secret", "", "Twitter Access Secret")
	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	db := getConnection()
	defer db.Close()
	migrate(db)

	client := twitter.NewClient(httpClient)

	params := &twitter.StreamFilterParams{
		Track:         []string{"@GoShirase test"},
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
		fmt.Println(user.Name)
		// 変数初期化
		count := -1
		db.Model(&Target{}).Where("user_id = ?", targetUser).Count(&count)
		fmt.Println(count)
		fmt.Println("@"+user.ScreenName+" you are not registered")
		if count != 1 {
			client.Statuses.Update("@"+user.ScreenName+" you are not registered", &twitter.StatusUpdateParams{
				InReplyToStatusID: tweet.ID,
			})
		}
	}

	for message := range stream.Messages {
		demux.Handle(message)
	}
}
