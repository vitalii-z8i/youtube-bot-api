package entities

// Message contains info about received message, autor and chat ID
type Message struct {
	ID   int32  `json:"message_id"`
	From User   `json:"from"`
	Chat Chat   `json:"chat"`
	Date int32  `json:"date"`
	Text string `json:"text"`
}
