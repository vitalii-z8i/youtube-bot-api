package telegram

import (
	"log"
	"strings"

	"github.com/vtl-pol/youtube-bot-api/config"
	"github.com/vtl-pol/youtube-bot-api/controllers"
	"github.com/vtl-pol/youtube-bot-api/entities"
	"github.com/vtl-pol/youtube-bot-api/services/msgutils"
)

// ProcessMessage analyzes text and calls according service
func ProcessMessage(message *entities.Message) (err error) {
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
		ProcessActionTrigger(message)

		_, err = msgutils.SendMessage(&message.Chat, "/start - Back to menu")
	}

	return err
}

// ProcessActionTrigger process message as a followup for previous one (if a trigger is set)
func ProcessActionTrigger(message *entities.Message) (err error) {
	var prevMsg = entities.Message{}
	config.DB.Connect()
	defer config.DB.Connection.Close()
	prevID, err := message.PrevID.Value()
	if err != nil {
		log.Println(err)
		return err
	}
	err = config.DB.Connection.Collection("messages").Find("ID = ?", prevID).One(&prevMsg)
	if err != nil {
		log.Println(err)
		return err
	}
	switch prevMsg.ActionTrigger {
	case "":
		msgutils.SendMessage(&message.Chat, "Yeah... About that")
		msgutils.SendTypingAction(&message.Chat)
		msgutils.SendMessage(&message.Chat, "I'dont have the code to do that thing, yet...\nBut soon...")
	case "channel_search":
		err = controllers.DisplayChannelsSearch(message.Text, message)
	}

	return err
}

// ProcessMessageCallback - process a callback from a message (Basically, a button click)
func ProcessMessageCallback(message *entities.Message, callBackData string) (err error) {
	config.DB.Connect()
	defer config.DB.Connection.Close()
	switch message.ActionTrigger {
	case "":
		msgutils.SendMessage(&message.Chat, "Yeah... About that")
		msgutils.SendTypingAction(&message.Chat)
		msgutils.SendMessage(&message.Chat, "I'dont have the code to do that thing, yet...\nBut soon...")
	case "confirm_subscription":
		msgutils.SendMessage(&message.Chat, "I can TOTALLY do that!")
		msgutils.SendMessage(&message.Chat, "But in like two hours or so...")
	}

	return err
}
