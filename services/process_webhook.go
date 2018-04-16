package services

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/vtl-pol/youtube-bot-api/config"
	"github.com/vtl-pol/youtube-bot-api/entities"
)

// ProcessWebhook parses request body from Telegram into a system Webhook entity
func ProcessWebhook(respose http.ResponseWriter, request *http.Request) {
	var webhook entities.Webhook
	var errorMessage = "Sorry pal, there's some error on my side. Gotta fix it. But feel free to drop by once I clean up this mess"

	fmt.Println("-------------REQUEST-------------")
	rawBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		SendMessage(&webhook.Message.Chat, errorMessage)
		fmt.Fprintf(respose, "%s", err)
	}

	fmt.Printf("%s", rawBody)
	fmt.Printf("\n")

	err = json.Unmarshal(rawBody, &webhook)

	if err != nil {
		SendMessage(&webhook.Message.Chat, errorMessage)
		fmt.Println(err)
		io.WriteString(respose, "All bad")
	} else if webhook.Message.ID != 0 {
		SendTypingAction(&webhook.Message.Chat)
		var message entities.Message
		var err error
		if message, err = webhook.StoreWebhookInfo(); err != nil {
			SendMessage(&webhook.Message.Chat, errorMessage)
		} else {
			ParseMessage(&message)
			// defaultMessage := fmt.Sprintf("Hey There, %s! I'm a new bot and can't do much stuff for now. But stay tuned and maybe some day I'll learn something", webhook.Message.From.FirstName)
			// err = SendMessage(&webhook.Message.Chat, defaultMessage)
		}
		if err != nil {
			log.Println(err)
			io.WriteString(respose, "All bad again")
		} else {
			io.WriteString(respose, "All ok here")
		}
	} else if webhook.EditedMessage.ID != 0 {
		config.DB.Connect()
		defer config.DB.Connection.Close()
		_, err := config.DB.Connection.Update("messages").Set("Text", webhook.EditedMessage.Text).Where("ID", webhook.EditedMessage.ID).Exec()
		if err != nil {
			SendMessage(&webhook.EditedMessage.Chat, "Sorry pal, I'm unable to update your message. Guess, it's out there for ever")
		}
	}

	fmt.Println("---------------------------------")
}
