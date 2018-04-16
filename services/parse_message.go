package services

import (
	"fmt"

	"github.com/vtl-pol/youtube-bot-api/entities"
)

// ParseMessage analyzes text and calls according service
func ParseMessage(message *entities.Message) (err error) {
	if command := message.FetchCommand(); command != "" {
		// err = SendMessage(message.Chat, "Command")
		fmt.Println("COMMAND: ", command)
		switch command {
		case "/start":
			err = SendMessage(&message.Chat, fmt.Sprintf("Welcome %s! I'll help you to keep track on and manage your YouTube stuff. Lets get started:", message.From.FirstName))
			if err == nil {
				menuText := "/channels    - A list of channels, you're subcribed;\n"
				menuText += "/subscribe   - Subscribe to a channel (duh)\n"
				menuText += "/unsubscribe - Cancel a subscription to one of your channels (duh)"
				err = SendMessage(&message.Chat, menuText)
			}
		default:
			SendMessage(&message.Chat, "Yeah... About that")
			SendTypingAction(&message.Chat)
			SendMessage(&message.Chat, "I'dont have the code to do that thing, yet...")
			SendMessage(&message.Chat, "But soon...")

			err = SendMessage(&message.Chat, "/start - Back to menu")
		}
	} else {
		err = SendMessage(&message.Chat, "Whoa! Take it easy, pal \U0001f604 I'm not even sure how to process the commands, I supposed to do. Not to talk about a custom message")
		err = SendMessage(&message.Chat, "Let's just /start over, shal we? \U0001f604")
	}
	return err
}
