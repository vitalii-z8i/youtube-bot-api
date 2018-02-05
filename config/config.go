package config

import (
	"log"
	"os"

	"github.com/jinzhu/configor"
)

// TelegramConfig contains all Telegram-related configuration
type TelegramConfig struct {
	APIKey string
	APIUrl string
}

// FullURL Get APi url for specified action
func (t *TelegramConfig) FullURL(uri string) string {
	return t.APIUrl + t.APIKey + "/" + uri
}

// Telegram is an instance of TelegramConfig
var Telegram TelegramConfig

func init() {
	var configFile = "config/telegram_config.json"
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Panicf("Missing Config: file %s was not found. \n", configFile)
	}
	if err := configor.Load(&Telegram, configFile); err != nil {
		log.Panicln(err)
	}
}
