package entities

// Webhook contains all the info
type Webhook struct {
	ID      int64   `json:"update_id"`
	Message Message `json:"message"`
}
