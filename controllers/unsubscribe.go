package controllers

import (
	"log"

	"github.com/vtl-pol/youtube-bot-api/services/msgutils"

	"github.com/vtl-pol/youtube-bot-api/config"
	"github.com/vtl-pol/youtube-bot-api/entities"
)

// Unsubscribe removes a channel from users subscriptions
func Unsubscribe(m *entities.Message) (err error) {
	config.DB.Connect()
	defer config.DB.Connection.Close()
	channels := []entities.Subscription{}
	err = config.DB.Connection.Collection("subscriptions").Find("UserID = ?", m.FromID).All(&channels)
	if err != nil {
		log.Println(err)
		return err
	}

	if len(channels) == 0 {
		_, err = msgutils.SendMessage(&m.Chat, "You haven't subscribed to any channels, so... Job's done, I guess \U0001f604")
		if err != nil {
			log.Println(err)
			return err
		}

		_, err = msgutils.SendMessage(&m.Chat, "Anyways, I have a better idea. Let's /subscribe to a new channel instead \U0001F609")
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		_, err = msgutils.SendMessage(&m.Chat, "You got it! Enter a name of a channel, you'd like to subscribe to and I'll do the job:")
	}

	return err
}
