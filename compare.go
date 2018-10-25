package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func compare(client *twitter.Client) {
	db := getConnection()
	defer db.Close()
	compareMigrate(db)

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
		fmt.Println(targetUser)
		// 変数初期化
		count := -1
		db.Model(&Target{}).Where("user_id = ?", targetUser).Count(&count)
		fmt.Println(count)
		fmt.Println("@" + user.ScreenName + " you are not registered")
		if count == 0 {
			client.Statuses.Update("@"+user.ScreenName+" you are not registered", &twitter.StatusUpdateParams{
				InReplyToStatusID: tweet.ID,
			})
		} else if count == 1 {
			db.Exec("Create table IF NOT EXISTS " + user.ScreenName + " like target_details")
			db.Exec("truncate table " + user.ScreenName)
			followers, _, err := client.Followers.List(&twitter.FollowerListParams{
				UserID: tweet.User.ID,
				Count:  500,
			})
			if err != nil {
				return
			}
			
			for _, follower := range followers.Users {
				db.Exec("insert into " + user.ScreenName + "(user_id,follower)values(" + strconv.FormatInt(targetUser, 10) + "," + strconv.FormatInt(follower.ID, 10) + ")")
			}
			fmt.Println(replaceSelectSql(CompareNewOldSQL, REPLACE1, user.ScreenName))
			rows, err := db.Raw(replaceSelectSql(CompareNewOldSQL, REPLACE1, user.ScreenName)).Rows()
			if err != nil {
				panic(err.Error())
			}

			rise := []string{}
			reduction := []string{}
			for rows.Next() {
				var result CompareResult
				db.ScanRows(rows, &result)
				tmp, _, err := client.Users.Show(&twitter.UserShowParams{
					UserID: result.Follower,
				})
				if err != nil {
					continue
				}
				if result.NewOldFlag == 0 {
					reduction = append(reduction, tmp.Name)
				} else {
					rise = append(rise, tmp.Name)
				}
			}
			rows.Close()
			var buffer bytes.Buffer
			buffer.WriteString("-増加-\r\n")
			if cap(rise) != 0 {
				buffer.WriteString(strings.Join(rise, ",") + "\r\n")
			} else {
				buffer.WriteString("なし\r\n")
			}
			buffer.WriteString("-減少-\r\n")
			if cap(reduction) != 0 {
				buffer.WriteString(strings.Join(reduction, ",") + "\r\n")
			} else {
				buffer.WriteString("なし\r\n")
			}
			// データ入れ替え
			db.Delete(TargetDetail{}, "user_id = ?", targetUser)
			db.Exec("insert into target_details select * from " + user.ScreenName)
			db.Exec("Drop table " + user.ScreenName)

			client.Statuses.Update("@"+user.ScreenName+" "+buffer.String(), &twitter.StatusUpdateParams{
				InReplyToStatusID: tweet.ID,
			})
		} else {
			client.Statuses.Update("@"+user.ScreenName+" error occered", &twitter.StatusUpdateParams{
				InReplyToStatusID: tweet.ID,
			})
		}
	}

	for message := range stream.Messages {
		demux.Handle(message)
	}
}
