package msgutils

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

// MessageRequest - service struct to handle TG API response as a message
type MessageRequest struct {
	Result entities.Message `json:"result"`
}

// SendMessage sends a reply to a telegram chat
func SendMessage(chat *entities.Chat, message string) (storedMsg entities.Message, err error) {
	if chat.ID == 0 || message == "" {
		err = errors.New("Missing Data (Re-check your chat Id and Message)")
		return storedMsg, err
	}

	req, err := http.NewRequest("GET", config.Telegram.FullURL("sendMessage"), nil)
	if err != nil {
		log.Println(err)
		return storedMsg, err
	}
	query := req.URL.Query()
	query.Add("chat_id", strconv.Itoa(int(chat.ID)))
	query.Add("text", message)

	req.URL.RawQuery = query.Encode()
	response, err := http.Get(req.URL.String())
	if err != nil {
		log.Println(err)
		return storedMsg, err
	}

	respBody, err := ioutil.ReadAll(response.Body)

	botMsg := MessageRequest{}
	err = json.Unmarshal(respBody, &botMsg)
	if err != nil {
		log.Println(err)
		return storedMsg, err
	}

	storedMsg, err = StoreBotMessage(&botMsg.Result)

	return storedMsg, err
}

// SendMessageWithKeyboard is just like SendMessage only with the keyboard
func SendMessageWithKeyboard(chat *entities.Chat, message string, replyMarkup Keyboard) (storedMsg entities.Message, err error) {
	if chat.ID == 0 || message == "" {
		err = errors.New("Missing Data (Re-check your chat Id and Message)")
		return storedMsg, err
	}

	req, err := http.NewRequest("GET", config.Telegram.FullURL("sendMessage"), nil)
	if err != nil {
		log.Println(err)
		return storedMsg, err
	}

	query := req.URL.Query()
	query.Add("chat_id", strconv.Itoa(int(chat.ID)))
	query.Add("text", message)
	markupString := fmt.Sprintf("{\"%s\":%s}", "inline_keyboard", replyMarkup.Keys)
	query.Add("reply_markup", markupString)
	req.URL.RawQuery = query.Encode()

	response, err := http.Get(req.URL.String())
	if err != nil {
		log.Println(err)
		return storedMsg, err
	}

	respBody, err := ioutil.ReadAll(response.Body)

	botMsg := MessageRequest{}
	err = json.Unmarshal(respBody, &botMsg)
	if err != nil {
		log.Println(err)
		return storedMsg, err
	}

	storedMsg, err = StoreBotMessage(&botMsg.Result)

	return storedMsg, err
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

// StoreBotMessage saves a bot message to DB before send
func StoreBotMessage(msg *entities.Message) (entities.Message, error) {
	config.DB.Connect()
	defer config.DB.Connection.Close()
	lastMessage := entities.Message{}
	tx, _ := config.DB.Connection.NewTx(nil)
	// Don't really care for the userID for a bot message
	msg.FromID = msg.From.ID
	msg.ChatID = msg.Chat.ID

	msg.ChatLastMessage().One(&lastMessage)
	msg.PrevID = sql.NullInt64{Int64: lastMessage.ID, Valid: (lastMessage.ID != 0)}

	_, err := tx.Collection("messages").Insert(msg)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return *msg, err
	}
	if lastMessage.ID != 0 {
		lastMessage.NextID = sql.NullInt64{Int64: msg.ID, Valid: (msg.ID != 0)}
		err = tx.Collection("messages").UpdateReturning(&lastMessage)
	}
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return *msg, err
	}
	tx.Commit()
	return *msg, err
}
