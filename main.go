package main

import (
	"log"
	"os"
	"sort"

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
			Name:    "add",
			Aliases: []string{"g"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
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

	// flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	// consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	// consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
	// accessToken := flags.String("access-token", "", "Twitter Access Token")
	// accessSecret := flags.String("access-secret", "", "Twitter Access Secret")
	// flags.Parse(os.Args[1:])
	// flagutil.SetFlagsFromEnv(flags, "TWITTER")

	// if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
	// 	log.Fatal("Consumer key/secret and Access token/secret required")
	// }

	// config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	// token := oauth1.NewToken(*accessToken, *accessSecret)
	// httpClient := config.Client(oauth1.NoContext, token)

	// client := twitter.NewClient(httpClient)
	// fmt.Println(client)

	// :TODO CLI化
	// registerAll(client)
	// registerFollower(client)
	// goshirase(client)
	// deleteNotices(client, 1)
}

func selectConfigure() {
	env := envParse()
	println(env)
}
