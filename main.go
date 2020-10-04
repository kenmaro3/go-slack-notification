package main

import (
	"fmt"

	"go-slack-notification/config"
	"go-slack-notification/slack"
)

func main() {
	fmt.Println(config.Config.BotAccessToken)
	slack.Message()
}
