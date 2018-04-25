package telegram

import (
	"fmt"
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
		msgutils.SendMessage(&message.Chat, "I'm not really good with non-standart messages, so...\nLet's just keep it simple, OK?\U0001F609")
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
		msgutils.SendMessage(&message.Chat, "I'm not really good with non-standart messages, so...\nLet's just keep it simple, OK?\U0001F609")
	case "confirm_subscription":
		newSub, err := controllers.ConfirmSubscription(message, callBackData)
		if err != nil {
			log.Println(err)
			return err
		}
		msgutils.SendMessage(&message.Chat, fmt.Sprintf("All done! You've subscribed to \"%s\" channel.\n", newSub.ChannelName))
		msgutils.SendMessage(&message.Chat, "You'll receive new videos as they're published")
	}

	return err
}
