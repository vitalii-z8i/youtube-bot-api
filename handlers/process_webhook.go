package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/vtl-pol/youtube-bot-api/config"
	"github.com/vtl-pol/youtube-bot-api/entities"
	"github.com/vtl-pol/youtube-bot-api/services"
)

// ProcessWebhook parses request body from Telegram into a system Webhook entity
func ProcessWebhook(respose http.ResponseWriter, request *http.Request) {
	var webhook entities.Webhook
	var errorMessage = "Sorry pal, there's some error on my side. Gotta fix it. But feel free to drop by once I clean up this mess"

	log.Println("-------------REQUEST-------------")
	defer log.Println("---------------------------------")
	rawBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(respose, "%s", "Error has occured")

		return
	}
	log.Printf("PARAMETERS: %s", rawBody)
	if err = json.Unmarshal(rawBody, &webhook); err != nil {
		services.SendMessage(&webhook.Message.Chat, errorMessage)
		log.Println(err)
		fmt.Fprintf(respose, "%s", "Error has occured")

		return
	}
	fmt.Fprintf(respose, "%s", "Accepted")

	if webhook.Message.ID != 0 {
		services.SendTypingAction(&webhook.Message.Chat)
		var message entities.Message
		var err error
		if message, err = webhook.StoreWebhookInfo(); err != nil {
			services.SendMessage(&webhook.Message.Chat, errorMessage)
		} else {
			ProcessMessage(&message)
		}
		if err != nil {
			log.Println(err)
		}
		return
	}

	if webhook.EditedMessage.ID != 0 {
		config.DB.Connect()
		defer config.DB.Connection.Close()
		_, err := config.DB.Connection.Update("messages").Set("Text", webhook.EditedMessage.Text).Where("ID", webhook.EditedMessage.ID).Exec()
		if err != nil {
			services.SendMessage(&webhook.EditedMessage.Chat, "Sorry pal, I'm unable to update your message. Guess, it's out there for ever")
			log.Println(err)
			return
		}

		return
	}
}
