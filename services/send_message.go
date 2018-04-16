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
func SendMessage(chat *entities.Chat, message string) (err error) {
	if chat.ID == 0 || message == "" {
		err = errors.New("Missing Data (Re-check your chat Id and Message)")
		return err
	}

	req, err := http.NewRequest("GET", config.Telegram.FullURL("sendMessage"), nil)
	if err != nil {
		log.Println(err)
		return err
	}

	query := req.URL.Query()
	query.Add("chat_id", strconv.Itoa(int(chat.ID)))
	query.Add("text", message)
	req.URL.RawQuery = query.Encode()

	//TODO: Add response check here!!!
	http.Get(req.URL.String())

	return err
}

// SendTypingAction lets user know, that bot is working on it
func SendTypingAction(chat *entities.Chat) (err error) {
	if chat.ID == 0 {
		err = errors.New("Missing Data (Re-check your chat Id and Message)")
		return err
	}

	req, err := http.NewRequest("GET", config.Telegram.FullURL("sendChatAction"), nil)
	if err != nil {
		log.Println(err)
		return err
	}

	query := req.URL.Query()
	query.Add("chat_id", strconv.Itoa(int(chat.ID)))
	query.Add("action", "typing")
	req.URL.RawQuery = query.Encode()

	_, err = http.Get(req.URL.String())
	return err
}
