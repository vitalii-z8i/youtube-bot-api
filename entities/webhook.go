package entities

import (
	"database/sql"
	"log"

	"github.com/vtl-pol/youtube-bot-api/config"
)

// Webhook contains all the info
type Webhook struct {
	ID            int64   `json:"update_id"`
	Message       Message `json:"message,omitempty"`
	EditedMessage Message `json:"edited_message,omitempty"`
}

// StoreWebhookInfo saves information from webhooks into DB
func (wh *Webhook) StoreWebhookInfo() (Message, error) {
	config.DB.Connect()
	defer config.DB.Connection.Close()
	var err error
	var existingUser User
	var user = wh.Message.From
	var chat = wh.Message.Chat
	var message = wh.Message
	var lastMessage = Message{}

	config.DB.Connection.Collection("users").Find("ID", user.ID).One(&existingUser)
	if existingUser.ID == 0 {
		_, err = config.DB.Connection.Collection("users").Insert(&user)
	} else if existingUser.FirstName != user.FirstName {
		err = config.DB.Connection.Collection("users").UpdateReturning(&user)
	}
	if err != nil {
		log.Println(err)
		return wh.Message, err
	}

	if chatExists, _ := config.DB.Connection.Collection("chats").Find("ID", chat.ID).Exists(); !chatExists {
		_, err = config.DB.Connection.Collection("chats").Insert(&chat)
	}

	tx, _ := config.DB.Connection.NewTx(nil)
	chat.LastMessage().One(&lastMessage)
	message.FromID = user.ID
	message.ChatID = chat.ID
	message.PrevID = sql.NullInt64{Int64: lastMessage.ID, Valid: (lastMessage.ID != 0)}

	_, err = tx.Collection("messages").Insert(&message)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return message, err
	}
	if lastMessage.ID != 0 {
		lastMessage.NextID = sql.NullInt64{Int64: message.ID, Valid: (message.ID != 0)}
		err = tx.Collection("messages").UpdateReturning(&lastMessage)
	}
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return message, err
	}
	tx.Commit()

	return message, nil
}
