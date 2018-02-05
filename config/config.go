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
	if _, err := os.Stat("config/telegram_config.json"); os.IsNotExist(err) {
		log.Panicln(err)
	}
	if err := configor.Load(&Telegram, "config/telegram_config.json"); err != nil {
		log.Panicln(err)
	}
}
