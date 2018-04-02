package entities

import (
	"github.com/vtl-pol/youtube-bot-api/config"
	"upper.io/db.v3"
)

// Chat contains an ID of a chat with a bot and its messages
type Chat struct {
	ID       int64 `json:"ID" db:"ID"`
	Messages []Message
}

// LastMessage Grabs a last message from current Chat
func (c *Chat) LastMessage() db.Result {
	return config.DB.Connection.Collection("messages").Find("ChatID = ? AND NextID IS NULL", c.ID)
}
