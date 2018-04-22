package controllers

import (
	"fmt"
	"log"
	"regexp"

	"github.com/vtl-pol/youtube-bot-api/config"
	"github.com/vtl-pol/youtube-bot-api/entities"
	"github.com/vtl-pol/youtube-bot-api/services"
)

// Subscribe searches a channel on YouTube (if anything provided) or shows a search hint
func Subscribe(m *entities.Message) (err error) {
	config.DB.Connect()
	defer config.DB.Connection.Close()
	if channelName := m.ArgsAfterCommand("/subscribe"); channelName != "" {
		// Remove extra data from possible channel name
		reg, err := regexp.Compile("(to )|(for )")
		if err != nil {
			log.Fatal(err)
			return err
		}
		channelName = reg.ReplaceAllString(channelName, "")

		messageText := fmt.Sprintf("Got it. Let me see if I can find something by the \"%s\" name", channelName)
		services.SendMessage(&m.Chat, messageText)
		// call channel search
		services.SendMessage(&m.Chat, "Yeah... I have no YouTube API at the moment. So, come back later and it'll be ready. I promise)))")
	}

	m.SetActionTrigger("channel_search")
	services.SendMessage(&m.Chat, "You got it! Enter a name of a channel, you'd like to subscribe to and I'll do the job:")
	return err
}
