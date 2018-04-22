package handlers

import (
	"log"
	"strings"

	"github.com/vtl-pol/youtube-bot-api/controllers"
	"github.com/vtl-pol/youtube-bot-api/entities"
	"github.com/vtl-pol/youtube-bot-api/services"
)

// ProcessMessage analyzes text and calls according service
func ProcessMessage(message *entities.Message) (err error) {
	err = services.SendTypingAction(&message.Chat)
	if err != nil {
		log.Println(err)
		return err
	}

	switch {
	case strings.Contains(message.Text, "/start"):
		err = controllers.Start(message)
	case strings.Contains(message.Text, "/channels"):
		err = controllers.ListChannels(message)
	case strings.Contains(message.Text, "/subscribe"):
		err = controllers.Subscribe(message)
	case strings.Contains(message.Text, "/unsubscribe"):
		err = controllers.Unsubscribe(message)
	default:
		services.SendMessage(&message.Chat, "Yeah... About that")
		services.SendTypingAction(&message.Chat)
		services.SendMessage(&message.Chat, "I'dont have the code to do that thing, yet...")
		services.SendMessage(&message.Chat, "But soon...")

		err = services.SendMessage(&message.Chat, "/start - Back to menu")
	}

	return err
}
