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
		_, err = msgutils.SendMessage(&m.Chat, messageText)
		if err != nil {
			log.Println(err)
			return err
		}
		err = msgutils.SendTypingAction(&m.Chat)
		if err != nil {
			log.Println(err)
			return err
		}

		err = DisplayChannelsSearch(messageText, m)
		return err
	}

	botMsg, err := msgutils.SendMessage(&m.Chat, "You got it! Enter a name of a channel, you'd like to subscribe to and I'll do the job:")
	if err != nil {
		log.Println(err)
		return err
	}
	botMsg.SetActionTrigger("channel_search")

	return err
}

// DisplayChannelsSearch gets possible channels from YT and displays them for pick
func DisplayChannelsSearch(messageText string, m *entities.Message) (err error) {
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
		botMsg, err := msgutils.SendMessageWithKeyboard(&m.Chat, messageText, msgutils.GenerateKeyboard("inline_keyboard", channelsJSON))
		if err != nil {
			log.Println(err)
			return err
		}
		botMsg.SetActionTrigger("confirm_subscription")
	} else {
		messageText = "Sorry pal. Looks like there's nothing under that name.\nHit /subscribe to start a new search"
		_, err = msgutils.SendMessage(&m.Chat, messageText)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return err
}

// ConfirmSubscription creates a subscription for selected channel
func ConfirmSubscription(m *entities.Message, channelID string) (newSub entities.Subscription, err error) {
	config.DB.Connect()
	defer config.DB.Connection.Close()
	channel, err := ytutils.FindChannel(channelID)
	if err != nil {
		log.Println(err)
		return newSub, err
	}

	newSub = entities.Subscription{UserID: m.FromID, ChannelID: channel.ChannelID, ChannelName: channel.ChannelName, ChannelInfo: channel.ChannelInfo}
	err = config.DB.Connection.Collection("subscriptions").InsertReturning(&newSub)
	if err != nil {
		log.Println(err)
		return newSub, err
	}
	err = newSub.Subscribe()
	if err != nil {
		log.Println(err)
		return newSub, err
	}

	return newSub, err
}
