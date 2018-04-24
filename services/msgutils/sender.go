package msgutils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/vtl-pol/youtube-bot-api/config"
	"github.com/vtl-pol/youtube-bot-api/entities"
)

// Keyboard - A struct to append an inline keyboard to a message
type Keyboard struct {
	KeyboardType string
	Keys         []byte
}

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
	_, err = http.Get(req.URL.String())

	return err
}

// SendMessageWithKeyboard is just like SendMessage only with the keyboard
func SendMessageWithKeyboard(chat *entities.Chat, message string, replyMarkup Keyboard) (err error) {
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
	markupString := fmt.Sprintf("{\"%s\":%s}", "inline_keyboard", replyMarkup.Keys)
	query.Add("reply_markup", markupString)
	req.URL.RawQuery = query.Encode()

	log.Println("==========", req.URL.String(), "==========")
	_, err = http.Get(req.URL.String())

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

// GenerateKeyboard prepares a keyboard for a message
func GenerateKeyboard(ktype string, keys []byte) Keyboard {
	return Keyboard{KeyboardType: ktype, Keys: keys}
}
