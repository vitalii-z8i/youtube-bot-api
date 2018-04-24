package main

import (
	"log"
	"net/http"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"

	"github.com/vtl-pol/youtube-bot-api/config"
	"github.com/vtl-pol/youtube-bot-api/handlers/telegram"
)

func main() {
	m, err := migrate.New(
		"file://db/migrations",
		"sqlite3://"+config.DB.ConnectionURL.Database)
	if err != nil {
		log.Panicln(err)
	}
	m.Steps(5)

	http.HandleFunc("/tg", telegram.ProcessWebhook)

	log.Println("Started a web server on 8000 port (http://127.0.0.1:8000/)")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
