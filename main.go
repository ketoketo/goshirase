package main

import (
	"log"
	"os"
	"sort"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Goshirase"
	app.Usage = "test usage"
	app.Version = "0.0.1"

	// flags
	configName := ".goshirase/config"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "profile, p",
			Value:       "config",
			Usage:       "config file name",
			Destination: &configName,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "configure",
			Aliases: []string{"c"},
			Usage:   "set config file",
			Action: func(c *cli.Context) error {
				err := registerConfig(configName)
				return err
			},
		},
		{
			Name:    "registerall",
			Aliases: []string{"ra"},
			Usage:   "register all acounts",
			Action: func(c *cli.Context) error {
				client := createTwitterClient(configName)
				registerAll(client)
				return nil
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// :TODO CLI化
	// registerAll(client)
	// registerFollower(client)
	// goshirase(client)
	// deleteNotices(client, 1)
}

func createTwitterClient(configName string) *twitter.Client {
	env := envParse()
	// envに設定されていない場合、configファイルから取得する
	if env == nil {
		var err error
		env, err = parse(configName)
		if err != nil {
			log.Fatal(err)
			panic(err.Error)
		}
	}

	config := oauth1.NewConfig(env.ConsumerKey, env.ConsumerSecret)
	token := oauth1.NewToken(env.AccessToken, env.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	return client
}
