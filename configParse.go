package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	ConsumerKey    string `json:"TWITTER_CONSUMER_KEY"`
	ConsumerSecret string `json:"TWITTER_CONSUMER_SECRET"`
	AccessToken    string `json:"TWITTER_ACCESS_TOKEN"`
	AccessSecret   string `json:"TWITTER_ACCESS_SECRET"`
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
