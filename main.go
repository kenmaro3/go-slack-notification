package main

import (
	"fmt"

	"example.com/m/v2/config"
	"github.com/line/line-bot-sdk-go/linebot"
)

func postLineMessage(message string) {
	bot, err := linebot.New(config.Config.ChannelSecret, config.Config.ChannelToken)
	if err != nil {
		fmt.Println(err)
	}
	if _, err := bot.PushMessage(config.Config.UserId, linebot.NewTextMessage(message)).Do(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	// fmt.Println(config.Config.BotAccessToken)
	// slack.Message()

	postLineMessage("what uppppp")
}
