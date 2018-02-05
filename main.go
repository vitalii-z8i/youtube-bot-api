package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/vtl-pol/youtube-bot-api/entities"
	"github.com/vtl-pol/youtube-bot-api/services"
)

func processWebhook(respose http.ResponseWriter, request *http.Request) {
	fmt.Println("-------------REQUEST-------------")
	rawBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		fmt.Fprintf(respose, "%s", err)
	}

	fmt.Printf("%s", rawBody)
	fmt.Printf("\n")
	var webhook entities.Webhook

	err = json.Unmarshal(rawBody, &webhook)

	if err != nil {
		fmt.Println(err)
		io.WriteString(respose, "All bad")
	} else {
		defaultMessage := "Hey There! I'm a new bot and can't do much stuff for now. But stay tuned and maybe some day I'll learn something"
		result, err := services.SendMessage(webhook.Message.Chat, defaultMessage)
		if !result || err != nil {
			log.Println(err)
			io.WriteString(respose, "All bad again")
		} else {
			io.WriteString(respose, "All ok here")
		}
	}

	fmt.Println("---------------------------------")
}

// our main function
func main() {
	http.HandleFunc("/", processWebhook)

	fmt.Println("Started a web server on 8000 port (http://127.0.0.1:8000/)")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
