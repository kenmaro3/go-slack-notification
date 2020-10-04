package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	BotAccessToken string
	ClientSecret   string
	SigningSecret  string
	TargetChannel  string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		BotAccessToken: cfg.Section("slack").Key("bot_user_auth_access_token").String(),
		ClientSecret:   cfg.Section("slack").Key("client_secret").String(),
		SigningSecret:  cfg.Section("slack").Key("signing_secret").String(),
		TargetChannel:  cfg.Section("slack").Key("target_channel").String(),
	}
}
