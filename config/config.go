package config

import (
	"log"
	"os"

	"github.com/jinzhu/configor"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
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

// DatabaseConfig contains all DB-connection creds/options
type DatabaseConfig struct {
	sqlite.ConnectionURL
	Connection sqlbuilder.Database
}

// Connect establishes DB connection and runs an SQL-query
func (db *DatabaseConfig) Connect() {
	sess, err := sqlite.Open(db)
	if err != nil {
		log.Panicf("db.Open(): %q\n", err)
	}
	sess.SetLogging(true)
	db.Connection = sess
}

// BotConfig contains bot-related configuration
type BotConfig struct {
	Commands []string
}

// Telegram is an instance of TelegramConfig
var Telegram TelegramConfig

// DB connection instance
var DB DatabaseConfig

// BotSettings list of supported commands and maybe other bot-related stuff
var BotSettings BotConfig

func init() {
	var tgConfigFile = "config/telegram_config.json"
	if _, err := os.Stat(tgConfigFile); os.IsNotExist(err) {
		log.Panicf("Missing Config: file %s was not found. \n", tgConfigFile)
	}
	if err := configor.Load(&Telegram, tgConfigFile); err != nil {
		log.Panicln(err)
	}

	var dbConfigFile = "config/database_config.json"
	if _, err := os.Stat(dbConfigFile); os.IsNotExist(err) {
		log.Panicf("Missing Config: file %s was not found. \n", dbConfigFile)
	}
	if err := configor.Load(&DB, dbConfigFile); err != nil {
		log.Panicln(err)
	}

	var botConfFile = "config/bot_config.json"
	if _, err := os.Stat(botConfFile); os.IsNotExist(err) {
		log.Panicf("Missing Config: file %s was not found. \n", botConfFile)
	}
	if err := configor.Load(&BotSettings, botConfFile); err != nil {
		log.Panicln(err)
	}
}
