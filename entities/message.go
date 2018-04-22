package entities

import (
	"database/sql"
	"strings"

	"github.com/vtl-pol/youtube-bot-api/config"
	db "upper.io/db.v3"
)

// Message contains info about received message, autor and chat ID
type Message struct {
	ID            int64         `json:"message_id" db:"ID"`
	FromID        int64         `db:"FromID,omitempty"`
	ChatID        int64         `db:"ChatID,omitempty"`
	Chat          Chat          `json:"chat" db:"-"`
	From          User          `json:"from" db:"-"`
	Date          int64         `json:"date" db:"-"`
	Text          string        `json:"text" db:"Text,omitempty"`
	PrevID        sql.NullInt64 `db:"PrevID,omitempty"`
	NextID        sql.NullInt64 `db:"NextID,omitempty"`
	ActionTrigger string        `db:"ActionTriger,omitempty"`
}

// ChatLastMessage Grabs a last message from current Chat
func (m *Message) ChatLastMessage() db.Result {
	return config.DB.Connection.Collection("messages").Find("ChatID = ? AND NextID IS NULL", m.ChatID)
}

// ArgsAfterCommand - returns text after command. If a message contains one
func (m *Message) ArgsAfterCommand(command string) (result string) {

	if !strings.Contains(m.Text, command) {
		return ""
	}
	parts := strings.Split(m.Text, command)
	result = strings.TrimSpace(parts[len(parts)-1])

	return result
}

// SetActionTrigger - Marks a message as an action trigger to know, what to do with the next one
func (m *Message) SetActionTrigger(action string) {
	m.ActionTrigger = action
	config.DB.Connection.Collection("messages").UpdateReturning(m)
}
