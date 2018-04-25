package entities

import (
	"log"

	"github.com/vtl-pol/youtube-bot-api/config"
	"github.com/vtl-pol/youtube-bot-api/handlers/ytupdates"
)

// Subscription contains an info of subscribed channels
type Subscription struct {
	ID          int64  `db:"ID"`
	UserID      int64  `db:"UserID"`
	ChannelID   string `db:"ChannelID"`
	ChannelName string `db:"ChannelName"`
	ChannelInfo string `db:"ChannelInfo"`
}

// YTChannel contains basic info of YT channel itself
type YTChannel struct {
	ChannelID   string `json:"callback_data" db:"YTChannelID"`
	ChannelName string `json:"text" db:"ChannelName"`
	ChannelInfo string `db:"ChannelInfo" json:"-"`
}

// Subscribe current channel for new vids
func (sub *Subscription) Subscribe() (err error) {
	channelURL, err := config.SubConf.TopicURLFor(sub.ChannelID)
	if err != nil {
		log.Println(err)
		return err
	}
	config.SubClient.Subscribe(config.SubConf.HubURL, channelURL, ytupdates.ProcessUpdate)
	return err
}

// Unsubscribe current channel for new vids
func (sub *Subscription) Unsubscribe() (err error) {
	config.DB.Connect()
	defer config.DB.Connection.Close()

	// Need to make sure, that no one else is subscribed to this channel before we proceed
	otherSubscriptionsExist, err := config.DB.Connection.Collection("subscriptions").Find("ID <> ? AND ChannedID == ?", sub.ID, sub.ChannelID).Exists()
	if err != nil {
		log.Println(err)
		return err
	}
	if !otherSubscriptionsExist {
		channelURL, err := config.SubConf.TopicURLFor(sub.ChannelID)
		if err != nil {
			log.Println(err)
			return err
		}
		config.SubClient.Unsubscribe(channelURL)
	}
	return err
}
