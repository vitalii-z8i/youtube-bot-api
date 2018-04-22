package controllers

import (
	"log"

	"github.com/vtl-pol/youtube-bot-api/config"
	"github.com/vtl-pol/youtube-bot-api/entities"
	"github.com/vtl-pol/youtube-bot-api/services"
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
		err = services.SendMessage(&m.Chat, "You haven't subscribed to any channels, so... Job's done, I guess \U0001f604")
		if err != nil {
			log.Println(err)
			return err
		}

		err = services.SendMessage(&m.Chat, "Anyways, I got a better idea. Let's /subscribe to a new channel instead \U0001F609")
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		m.SetActionTrigger("channel_search")
		err = services.SendMessage(&m.Chat, "You got it! Enter a name of a channel, you'd like to subscribe to and I'll do the job:")
	}

	return err
}
