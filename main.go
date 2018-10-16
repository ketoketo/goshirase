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
		Track:         []string{"golang", "Golang"},
		StallWarnings: twitter.Bool(true),
	}
	stream, _ := client.Streams.Filter(params)
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.Text)
		// client.Statuses.Retweet(tweet.ID, nil)
		fmt.Println(tweet.ID)
		fmt.Println(tweet.User.ID)
		fmt.Println(tweet.User.FollowersCount)
		user, _, err := client.Users.Show(&twitter.UserShowParams{
			UserID: tweet.User.ID,
		})
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.Name)
		fmt.Println(user.FollowersCount)
		// user.Protected
	}

	for message := range stream.Messages {
		demux.Handle(message)
	}
}
