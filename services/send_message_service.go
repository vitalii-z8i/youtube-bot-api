package services

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/vtl-pol/youtube-bot-api/config"
	"github.com/vtl-pol/youtube-bot-api/entities"
)

// SendMessage sends a reply to a telegram chat
func SendMessage(chat entities.Chat, message string) (status bool, err error) {
	if chat.ID == 0 || message == "" {
		err = errors.New("Missing Data (Re-check your chat Id and Message)")
		return false, err
	}

	req, err := http.NewRequest("GET", config.Telegram.FullURL("sendMessage"), nil)
	if err != nil {
		log.Println(err)
		return false, err
	}

	query := req.URL.Query()
	query.Add("chat_id", strconv.Itoa(chat.ID))
	query.Add("text", message)
	req.URL.RawQuery = query.Encode()

	//TODO: Add response check here!!!
	http.Get(req.URL.String())

	return true, nil
}
