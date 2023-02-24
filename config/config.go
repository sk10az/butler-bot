package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	ApiId           int    `yaml:"api_id"`
	ApiHash         string `yaml:"api_hash"`
	BotToken        string `yaml:"bot_token"`
	BearerToken     string `yaml:"bearer_token"`
	WhitelistChatId int64  `yaml:"whitelist_chat_id"`
}

type DBConfig struct {
	Host          string   `json:"host"`
	ReadOnlyHosts []string `json:"read_only_hosts"`
	Port          uint     `json:"port"`
	User          string   `json:"user"`
	Password      string   `json:"password"`
	DBName        string   `json:"db_name"`
	TZ            string   `json:"tz"`
}

func loadConfig() (*Config, error) {
	var config Config
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
