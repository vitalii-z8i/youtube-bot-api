package entities

import (
	"database/sql"
	"strings"

	"github.com/vtl-pol/youtube-bot-api/config"
)

// Message contains info about received message, autor and chat ID
type Message struct {
	ID     int64         `json:"message_id" db:"ID"`
	FromID int64         `db:"FromID,omitempty"`
	From   User          `json:"from" db:"-"`
	ChatID int64         `db:"ChatID,omitempty"`
	Chat   Chat          `json:"chat" db:"-"`
	Date   int64         `json:"date" db:"-"`
	Text   string        `json:"text" db:"Text,omitempty"`
	PrevID sql.NullInt64 `db:"PrevID,omitempty"`
	NextID sql.NullInt64 `db:"NextID,omitempty"`
}

// FetchCommand checks if a message contains a bot command
func (m *Message) FetchCommand() string {
	for _, command := range config.BotSettings.Commands {
		if strings.Contains(m.Text, command) {
			return command
		}
	}
	return ""
}
