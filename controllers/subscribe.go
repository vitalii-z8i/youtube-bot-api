package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"

	"github.com/vtl-pol/youtube-bot-api/config"
	"github.com/vtl-pol/youtube-bot-api/entities"
	"github.com/vtl-pol/youtube-bot-api/services/msgutils"
	"github.com/vtl-pol/youtube-bot-api/services/ytutils"
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
		err = msgutils.SendMessage(&m.Chat, messageText)
		if err != nil {
			log.Println(err)
			return err
		}
		err = msgutils.SendTypingAction(&m.Chat)
		if err != nil {
			log.Println(err)
			return err
		}

		foundChannels, err := ytutils.ChannelsSearch(messageText)
		if err != nil {
			log.Println(err)
			return err
		}
		channelsJSON, err := json.Marshal(foundChannels)

		if err != nil {
			log.Println(err)
			return err
		}
		if string(channelsJSON) != "null" {
			messageText = "Ok! Here's what I was able to find:\nPick a channel to confirm subscription or hit /subscribe to start a new search"
			err = msgutils.SendMessageWithKeyboard(&m.Chat, messageText, msgutils.GenerateKeyboard("inline_keyboard", channelsJSON))
		} else {
			messageText = "Sorry pal. Looks like there's nothing under that name.\nHit /subscribe to start a new search"
			err = msgutils.SendMessage(&m.Chat, messageText)
		}
		return err
	}

	m.SetActionTrigger("channel_search")
	msgutils.SendMessage(&m.Chat, "You got it! Enter a name of a channel, you'd like to subscribe to and I'll do the job:")
	return err
}
