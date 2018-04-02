package main

import (
	"log"
	"net/http"

	"github.com/vtl-pol/youtube-bot-api/services"
)

// our main function
func main() {
	http.HandleFunc("/tg", services.ProcessWebhook)

	log.Println("Started a web server on 8000 port (http://127.0.0.1:8000/)")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
