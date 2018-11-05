package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	consumerKey    string `json:"TWITTER_CONSUMER_KEY"`
	consumerSecret string `json:"TWITTER_CONSUMER_SECRET"`
	accessToken    string `json:"TWITTER_ACCESS_TOKEN"`
	accessSecret   string `json:"TWITTER_ACCESS_SECRET"`
}

func parse(filePath string) (Config, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		log.Fatal(err)
		return Config{}, err
	}
	return config, nil
}
