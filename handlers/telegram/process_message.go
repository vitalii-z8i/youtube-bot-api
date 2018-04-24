package telegram

import (
	"log"
	"strings"

	"github.com/vtl-pol/youtube-bot-api/controllers"
	"github.com/vtl-pol/youtube-bot-api/entities"
	"github.com/vtl-pol/youtube-bot-api/services/msgutils"
)

// ProcessMessage analyzes text and calls according service
func ProcessMessage(message *entities.Message) (err error) {
	err = msgutils.SendTypingAction(&message.Chat)
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
		msgutils.SendMessage(&message.Chat, "Yeah... About that")
		msgutils.SendTypingAction(&message.Chat)
		msgutils.SendMessage(&message.Chat, "I'dont have the code to do that thing, yet...")
		msgutils.SendMessage(&message.Chat, "But soon...")

		err = msgutils.SendMessage(&message.Chat, "/start - Back to menu")
	}

	return err
}
