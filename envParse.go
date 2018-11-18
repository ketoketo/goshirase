package main

import (
	"flag"
	"log"

	"github.com/coreos/pkg/flagutil"
)

func envParse() *Config {

	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "", "Twitter Access Token")
	accessSecret := flags.String("access-secret", "", "Twitter Access Secret")
	flagutil.SetFlagsFromEnv(flags, "TWITTER")

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Println("no env")
		return nil
	}

	config := &Config{
		ConsumerKey:    *consumerKey,
		ConsumerSecret: *consumerSecret,
		AccessToken:    *accessToken,
		AccessSecret:   *accessSecret,
	}
	return config
}
