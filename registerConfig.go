package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Config struct {
	ConsumerKey    string `json:"TWITTER_CONSUMER_KEY"`
	ConsumerSecret string `json:"TWITTER_CONSUMER_SECRET"`
	AccessToken    string `json:"TWITTER_ACCESS_TOKEN"`
	AccessSecret   string `json:"TWITTER_ACCESS_SECRET"`
}

func parse(filePath string) (*Config, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &Config{}, err
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return &Config{}, err
	}
	return &config, nil
}

func makeConfigJson(config *Config) ([]byte, error) {
	jsonBytes, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func writeConfig(path string, data []byte) {
	ioutil.WriteFile(path, data, 0644)
}

func stdReader(stdin io.Reader, message string, param string) string {
	fmt.Printf("%s [%s]: ", message, param)
	scanner := bufio.NewScanner(stdin)
	scanner.Scan()
	return scanner.Text()
}

func registerConfig(path string) error {
	// var config *Config
	config := &Config{}
	// ファイルが存在する場合、parse
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		config, err = parse(path)
		if err != nil {
			return err
		}
	}
	// 入力・登録(空文字の場合登録しない)
	tmpConsumerKey := stdReader(os.Stdin, "TWITTER_CONSUMER_KEY", config.ConsumerKey)
	tmpConsumerSecret := stdReader(os.Stdin, "TWITTER_CONSUMER_SECRET", config.ConsumerSecret)
	tmpAccessToken := stdReader(os.Stdin, "TWITTER_ACCESS_TOKEN", config.AccessToken)
	tmpAccessSecret := stdReader(os.Stdin, "TWITTER_ACCESS_SECRET", config.AccessSecret)
	if tmpConsumerKey != "" {
		config.ConsumerKey = tmpConsumerKey
	}
	if tmpConsumerSecret != "" {
		config.ConsumerSecret = tmpConsumerSecret
	}
	if tmpAccessToken != "" {
		config.AccessToken = tmpAccessToken
	}
	if tmpAccessSecret != "" {
		config.AccessSecret = tmpAccessSecret
	}

	data, err := makeConfigJson(config)
	if err != nil {
		return err
	}
	writeConfig(path, data)
	return nil
}
