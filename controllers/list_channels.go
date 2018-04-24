package controllers

import (
	"log"

	"github.com/vtl-pol/youtube-bot-api/services/msgutils"

	"github.com/vtl-pol/youtube-bot-api/config"
	"github.com/vtl-pol/youtube-bot-api/entities"
)

// ListChannels shows list of current subscriptions
func ListChannels(m *entities.Message) (err error) {
	config.DB.Connect()
	defer config.DB.Connection.Close()
	channels := []entities.Subscription{}
	err = config.DB.Connection.Collection("subscriptions").Find("UserID = ?", m.FromID).All(&channels)
	if err != nil {
		log.Println(err)
		return err
	}

	if len(channels) == 0 {

		_, err = msgutils.SendMessage(&m.Chat, "You haven't subscribed to any channels yet")
		if err != nil {
			log.Println(err)
			return err
		}

		_, err = msgutils.SendMessage(&m.Chat, "Let's just hit /subscribe and get the first one \U0001F609")
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return err
}
