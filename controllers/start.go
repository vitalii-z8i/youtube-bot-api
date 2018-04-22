package controllers

import (
	"fmt"

	"github.com/vtl-pol/youtube-bot-api/entities"
	"github.com/vtl-pol/youtube-bot-api/services"
)

// Start handles initial "/start" command
func Start(m *entities.Message) (err error) {
	err = services.SendMessage(&m.Chat, fmt.Sprintf("Welcome %s! I'll help you to keep track on and manage your YouTube stuff. Lets get started:", m.From.FirstName))
	if err == nil {
		menuText := "/channels    - A list of channels, you're subcribed;\n"
		menuText += "/subscribe   - Subscribe to a channel (duh)\n"
		menuText += "/unsubscribe - Cancel a subscription to one of your channels (duh)"
		err = services.SendMessage(&m.Chat, menuText)
	}
	return err
}
